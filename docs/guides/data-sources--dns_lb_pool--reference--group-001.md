---
page_title: "xcsh_dns_lb_pool reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_lb_pool reference."
---

# xcsh_dns_lb_pool reference

<a id="canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- Property reference

<a id="canonical-3322101301110102-1211003002103132-3002021033133312-0203030303022332-0221323210133122-0032230023333323-2322110220222100-2223303233023211"></a>

### Direct properties for `xcsh_dns_lb_pool`

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300): complete subsection reference.

- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3113200321003123-3322202013011013-3313213000031121-1133010331020321-1313120230301331-1310313133032133-3331032311202021-1201020223023122): complete subsection reference.

<a id="canonical-3200121310013212-1300120032222233-2013000032001212-3223201121202322-1210213333012120-0113020030330303-3011113200303030-3320032230102021"></a>

<a id="canonical-1200232032212023-3323222132120101-3113302300233011-1122032332333223-0103221300031033-3131133201021212-1012333321003022-0113000201231013"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 64,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "64",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
      "minLength": 1,
      "type": "string"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211): complete subsection reference.

<a id="canonical-1231120111310231-2113201131102322-0221313032130132-3002303010000202-2232222212211120-0230002020231210-2320323203132312-0302300221133031"></a>

<a id="canonical-0332201223213301-3133313232031223-2101222300110030-2210011302013212-2133112301213233-1310203130230120-1230311030133322-1313333023132230"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the DNSLBPool.

Additional upstream details:

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-1001203310300032-1212313021011232-2223202232121021-2301312123232032-2020002122301223-3201110133133322-3223122103332213-0031223122232211"></a>

<a id="canonical-2131003031002200-3223022122113100-3111203221200301-2131121320211013-1323332212020202-1233102002112023-3123332032030312-0022113101331211"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-1122220021122122-3301312332030333-1200310203101112-3313330101001120-1030203211113203-1232033210110313-2203303213222233-2013123333333202"></a>

<a id="canonical-0322011323133121-3003200221111321-3210203001130111-1013023122233320-0133201231333233-3113111101300122-0132023323010333-1221030231221321"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113202230333221-0213303332221032-3002003211103231-0221200021330333-1232103132120211-2220322333031031-0021230312001003-0122111012002322"></a>

<a id="canonical-0112102032001212-2031332102010311-3321212301100102-2023013223310201-1210313022022303-1323302211213313-3022110010132212-1121012231021221"></a>

#### `load_balancing_mode` property

Type: `"string"`. Computed.

\[Enum: ROUND\_ROBIN|RATIO\_MEMBER|STATIC\_PERSIST|PRIORITY\] - ROUND\_ROBIN: Round-Robin Round
Robin will ensure random equal distribution of requests among all pool members in a pool. -
RATIO\_MEMBER: Ratio-Member Ratio-Member performs load balancing of requests across the pool members
based on the ratio assigned to each pool member - STATIC\_PERSIST.. Possible values are
\`ROUND\_ROBIN\`, \`RATIO\_MEMBER\`, \`STATIC\_PERSIST\`, \`PRIORITY\`. Defaults to
\`ROUND\_ROBIN\`.

Additional upstream details:

&#8203;- ROUND\_ROBIN: Round-Robin

Round Robin will ensure random equal distribution of requests among all pool members in a pool.
&#8203;- RATIO\_MEMBER: Ratio-Member

Ratio-Member performs load balancing of requests across the pool members based on the ratio assigned
to each pool member &#8203;- STATIC\_PERSIST: Static-Persist

The Static Persist load balancing method uses the persist mask, with the source IP address of the
Local Domain Name Server (LDNS), in a deterministic algorithm to send requests to a specific pool
member. If the DNS resolver passes ECS (EDNS-Client-Subnet) information, then a hash of it will be
used, to send the client to the same pool member &#8203;- PRIORITY: Priority

The Priority load balancing method returns all available endpoints in a pool with the highest
priority. Pool Members have a priority value, starting from zero, where a lower value means a higher
priority.

Receipt-pinned upstream constraints:

```json
{
  "default": "ROUND_ROBIN",
  "enum": [
    "ROUND_ROBIN",
    "RATIO_MEMBER",
    "STATIC_PERSIST",
    "PRIORITY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-0222213301012121-1331313201223331-3230012302232130-1013220221311321-3332133002323312-1320012231332301-2220233021131232-2230101013222001): complete subsection reference.

<a id="canonical-1223220120030300-3020312223222101-3332101103323003-1221202311310212-3013103301311323-3323300010330220-2323131221231301-2201331203110112"></a>

<a id="canonical-2333031222302330-3031200233002333-0223302131330322-3200313333030332-0333021033033110-2211203202213213-3020322120013200-1121313012131131"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DNSLBPool.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2003132103112012-0330120120003200-0101132020011310-3111202011222231-2310010333120122-1212323130222323-1231300320001032-0023221020022301"></a>

<a id="canonical-2101231300233310-3003111130010001-0113221020333301-1021211311032210-2023300113203300-2013301122230111-0322333131001003-1112222301022032"></a>

#### `namespace` property

Type: `"string"`. Optional, Computed.

Namespace where the DNSLBPool exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1033101312200110-3320312032020203-1322213322022011-2111022211202232-2101111013330231-1000030101312201-0203211113230302-3310203000232311): complete subsection reference.

<a id="canonical-0033113021233003-0133213121121032-2023332223201320-3311101012103213-3101120133312321-0333012203223311-3130103111312300-1132310301220201"></a>

<a id="canonical-0001102201001000-2300010322013001-2222212302213123-0220212332211101-1233300021321033-0323211102221310-1020013031310233-3112032330310312"></a>

#### `ttl` property

Type: `"number"`. Computed.

\[OneOf: TTL, use\_rrset\_ttl\] Exclusive with \[use\_rrset\_ttl\] Custom TTL in seconds (default
&#8203;30) for responses from this pool.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 2147483647,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

OneOf alternatives in this subsection:

- [TTL](data-sources--dns_lb_pool--reference--group-001.md#canonical-0033113021233003-0133213121121032-2023332223201320-3311101012103213-3101120133312321-0333012203223311-3130103111312300-1132310301220201)
- [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-3001211201122102-3330232120003112-3211312103210012-0112200000110020-1332112311031231-0211010213230111-1231320310212332-2222030020130330)

Select alternatives according to the provider validators above.

- [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-3103203103023110-1213233013101233-0020122011003001-2311022001103201-0020331202131321-1010221101300002-1311010322110100-3230130000033321): complete subsection reference.

<a id="canonical-1231333223210133-0221123230033231-2201132022220210-2213211033010211-0031203110120212-0103310131311022-2132132101230221-1132133131130310"></a>

### All schema paths for `xcsh_dns_lb_pool`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `a_pool` | [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3110133003130222-1021132100113311-3331330212300301-3100002212022112-3200011232323111-2332211301132101-1213013312312012-1102311230111130) |
| `a_pool.disable_health_check` | [a_pool.disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-0031100231010012-0002120013110023-0100121003101213-0321213201023033-1111213200000113-2300132333000101-2123232201033201-1012311332230311) |
| `a_pool.health_check` | [a_pool.health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-3120111310311332-0303010031201131-0112120012000033-3212131233300002-2221111200032202-0222102300202101-0321120332022212-3232303300111120) |
| `a_pool.health_check.name` | [a_pool.health_check.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-1331312230123221-2120232130120002-0101103013012112-0130322212210200-2323210313120213-2321321211010100-0223100323230302-0030020233231231) |
| `a_pool.health_check.namespace` | [a_pool.health_check.namespace](data-sources--dns_lb_pool--reference--group-001.md#canonical-1232300231120212-1332100231200330-0100210001000233-0111310123111131-1233213301333200-2022223313130300-3003132113333110-0120303120103131) |
| `a_pool.health_check.tenant` | [a_pool.health_check.tenant](data-sources--dns_lb_pool--reference--group-001.md#canonical-2121323013301023-2210033222112232-1213110031233310-1111123222032312-3000013233333023-0330321213013131-0231120333102212-2102130221333120) |
| `a_pool.max_answers` | [a_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201303130211332-2110010323220111-1022101300110102-2101130312323322-2312213120022220-2023020231202121-3232331302302330-1220310221321003) |
| `a_pool.members` | [a_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201221002100222-0333330320222310-1211323130302203-1231002203320132-2323301303121031-2203221120113032-0332121201000322-0031232133311231) |
| `a_pool.members.disable_spec` | [a_pool.members.disable_spec](data-sources--dns_lb_pool--reference--group-001.md#canonical-2212033021101300-2231222101130002-1201131200302103-0112002320123111-2003123331233300-2133200310320233-2303000111323003-3003220110131312) |
| `a_pool.members.ip_endpoint` | [a_pool.members.ip_endpoint](data-sources--dns_lb_pool--reference--group-001.md#canonical-3131100313223120-3302033132030323-1130332221022120-3210020020101111-0232213013123100-2302213323231330-1031330100121031-3300320121010132) |
| `a_pool.members.name` | [a_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-3000312003313230-2212130320131123-3202123211301332-2011030213320102-1033030200031202-1322210332032011-2101031113010133-1120223331031010) |
| `a_pool.members.priority` | [a_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-3022300302032000-3102033102113103-3333310203013132-0100032022103223-0320101310031102-2323031101132003-0300131331132131-0131221332002330) |
| `a_pool.members.ratio` | [a_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-2310323210012100-1232203022301022-0212332022121132-0300013123200022-3101231311200021-3300001322232302-3022333310130112-2233103330203122) |
| `aaaa_pool` | [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3331333332133320-0322212213320112-0112131123311322-1123122031020212-3310020321110213-2301023312223320-2113010122302330-0223202221012201) |
| `aaaa_pool.max_answers` | [aaaa_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-0033310331011313-2011313132121013-0223120232130310-1010110022220232-2101113021123121-1110200123100212-0003112310033311-3122122223300112) |
| `aaaa_pool.members` | [aaaa_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-2032221233203020-3103321331030012-2203312132302232-3322330123120233-2012023123023021-2310222321301323-1132033310100021-0033212023011312) |
| `aaaa_pool.members.disable_spec` | [aaaa_pool.members.disable_spec](data-sources--dns_lb_pool--reference--group-001.md#canonical-1113312303322233-1203112112322120-1033332021122303-2313330303131230-1321321121310121-2322220313020030-0121022302000010-0332320112313313) |
| `aaaa_pool.members.ip_endpoint` | [aaaa_pool.members.ip_endpoint](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201103221131233-0333032132233123-1321332112203310-1022101210331023-0033011231032233-2123113322131301-1130332233211233-1022102302113321) |
| `aaaa_pool.members.name` | [aaaa_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-3032013130013031-0121333120331020-3213200002323303-3013032011012112-2333120220213110-0330233121012031-2011112230313123-2200023323110023) |
| `aaaa_pool.members.priority` | [aaaa_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-0121303210012132-3303122021110232-0323210113121332-2332001033230133-1311023222203121-3123320101120201-2012320003113330-1233221231011301) |
| `aaaa_pool.members.ratio` | [aaaa_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-0122113021312102-0212132213100210-0000131232122333-0010312133312120-1011002203003323-3111121322212000-0031121013102133-2002332033301213) |
| `annotations` | [annotations](data-sources--dns_lb_pool--reference--group-001.md#canonical-3200121310013212-1300120032222233-2013000032001212-3223201121202322-1210213333012120-0113020030330303-3011113200303030-3320032230102021) |
| `cname_pool` | [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2010322303230133-3010012212113232-3200112321212311-2223020101312321-0013120001231120-1221003303133011-1213100012111320-0100311321020333) |
| `cname_pool.disable_health_check` | [cname_pool.disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-1313122120131121-0010020131220232-3323033130211013-2100011133213110-0103203020030221-3303323012333021-3021321031231000-1101003010233002) |
| `cname_pool.health_check` | [cname_pool.health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-3202031120111001-2210233122231022-0313033123112102-2112332332200211-0132312010203122-1312111012120233-3211202322030131-3302132122232222) |
| `cname_pool.health_check.name` | [cname_pool.health_check.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-2012210332202320-0111133313002133-0313310212203231-0321011103321323-2020321102131003-3012220331112112-0031312131302022-1323112312010302) |
| `cname_pool.health_check.namespace` | [cname_pool.health_check.namespace](data-sources--dns_lb_pool--reference--group-001.md#canonical-0021021323201120-3331101333020020-1003222000232131-1211213113032000-2123331333322023-1000230312331111-3312220030330301-0101223333111101) |
| `cname_pool.health_check.tenant` | [cname_pool.health_check.tenant](data-sources--dns_lb_pool--reference--group-001.md#canonical-3312201130333321-2331012101330321-2110011033233233-2233032222112332-3031032012001330-0121202200210222-2102000033202222-1110231012201220) |
| `cname_pool.members` | [cname_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-3022200313100231-1100233233103313-2132231111031000-1332320021231230-1131002132100103-0113022223030231-2012212323323023-3003212210323303) |
| `cname_pool.members.domain` | [cname_pool.members.domain](data-sources--dns_lb_pool--reference--group-001.md#canonical-3302313311020030-2312002331121230-2222022033213113-0001003122033033-2201311312112011-0232301230122011-2023010332230310-0211303331212103) |
| `cname_pool.members.final_translation` | [cname_pool.members.final_translation](data-sources--dns_lb_pool--reference--group-001.md#canonical-2232333022232010-3022231010122301-2020121021230030-0002320301332001-2313233002201032-2120001331133231-1110002012103322-1200202020201310) |
| `cname_pool.members.name` | [cname_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-2212132012332023-3123020002300213-0321322202203322-2123101233230022-1330320021313010-1302012332120011-2001103232030232-1200123231130133) |
| `cname_pool.members.priority` | [cname_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-3201213022213111-2122302110010310-0321130202023303-2310212102000313-0130231031323100-0131211120210132-3302113102200203-0112003110200102) |
| `cname_pool.members.ratio` | [cname_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-3200133000313201-1330013333303231-1302303310211123-1033011332312323-0322132230310123-1003311223122223-0301312121003123-1230122112122313) |
| `description` | [description](data-sources--dns_lb_pool--reference--group-001.md#canonical-1231120111310231-2113201131102322-0221313032130132-3002303010000202-2232222212211120-0230002020231210-2320323203132312-0302300221133031) |
| `id` | [ID](data-sources--dns_lb_pool--reference--group-001.md#canonical-1001203310300032-1212313021011232-2223202232121021-2301312123232032-2020002122301223-3201110133133322-3223122103332213-0031223122232211) |
| `labels` | [labels](data-sources--dns_lb_pool--reference--group-001.md#canonical-1122220021122122-3301312332030333-1200310203101112-3313330101001120-1030203211113203-1232033210110313-2203303213222233-2013123333333202) |
| `load_balancing_mode` | [load_balancing_mode](data-sources--dns_lb_pool--reference--group-001.md#canonical-2113202230333221-0213303332221032-3002003211103231-0221200021330333-1232103132120211-2220322333031031-0021230312001003-0122111012002322) |
| `mx_pool` | [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1103013221030002-1110231321000100-0130202221021310-1310122223103021-1300123233202211-3330122203300012-3013122300220031-3031002220132331) |
| `mx_pool.max_answers` | [mx_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-3113231200203220-1311022201302020-0012012111320011-1003210021001221-2202023231310332-3130003320113213-1202003031220113-3322112103212123) |
| `mx_pool.members` | [mx_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-2112210003220031-1022213020322003-3022110131313310-2222000201202332-0120003220200322-2002123210303312-3001220001100212-2011213002323033) |
| `mx_pool.members.domain` | [mx_pool.members.domain](data-sources--dns_lb_pool--reference--group-001.md#canonical-2110210001002220-0303133112203020-0033203100233203-3131211023222121-3223031121300121-2322322302301322-2123322102211212-1003300132311220) |
| `mx_pool.members.name` | [mx_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-0101232032320202-3103001101100223-0100021100020201-0223123332102123-3312303101020232-1121223132331222-3210320313121130-3231300231213131) |
| `mx_pool.members.priority` | [mx_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-3000213031101300-1220020011030300-0332203233123211-1331012212032002-0102111023111232-0021020233032322-3201101002120011-3213113321200212) |
| `mx_pool.members.ratio` | [mx_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-3232222003000233-2321303102103331-2133030322303322-2213303233300103-3230020011333103-1210320231112133-0013001013331011-1331132201023112) |
| `name` | [name](data-sources--dns_lb_pool--reference--group-001.md#canonical-1223220120030300-3020312223222101-3332101103323003-1221202311310212-3013103301311323-3323300010330220-2323131221231301-2201331203110112) |
| `namespace` | [namespace](data-sources--dns_lb_pool--reference--group-001.md#canonical-2003132103112012-0330120120003200-0101132020011310-3111202011222231-2310010333120122-1212323130222323-1231300320001032-0023221020022301) |
| `srv_pool` | [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-0233230131013131-3113222100133202-0320130302200322-1223023300133003-0123301120002202-0233201320130300-3311320020232022-3223111032021133) |
| `srv_pool.max_answers` | [srv_pool.max_answers](data-sources--dns_lb_pool--reference--group-001.md#canonical-1211032202013310-0200303023131203-0230333222110122-2222122232302210-1331001112200030-3232330231030301-3231012133131011-1220312311010021) |
| `srv_pool.members` | [srv_pool.members](data-sources--dns_lb_pool--reference--group-001.md#canonical-3213123332003032-1111331012012202-0022210301122202-3301110100223111-2020220213031330-1223010000121202-1203023232313010-1032112030210112) |
| `srv_pool.members.final_translation` | [srv_pool.members.final_translation](data-sources--dns_lb_pool--reference--group-001.md#canonical-1100310100133200-0103312022233130-1222322320121123-3301222013112110-0310310003322103-1231111021323022-3232301230300031-2110211123312111) |
| `srv_pool.members.name` | [srv_pool.members.name](data-sources--dns_lb_pool--reference--group-001.md#canonical-2302221332300312-2300221133302313-0301020003310111-0122031320023123-2301023220031202-3123202030212010-1100012100021231-3021132232302230) |
| `srv_pool.members.port` | [srv_pool.members.port](data-sources--dns_lb_pool--reference--group-001.md#canonical-2311310102100120-2232003021301133-0313311232103323-3210310301121331-0333111323212213-3112103303011223-2022032210032113-2210222030301001) |
| `srv_pool.members.priority` | [srv_pool.members.priority](data-sources--dns_lb_pool--reference--group-001.md#canonical-1132111011022311-1030013110301333-3302120233002201-0212100212212310-0000113030310002-0313001230222222-0213100103301010-1203030100233211) |
| `srv_pool.members.ratio` | [srv_pool.members.ratio](data-sources--dns_lb_pool--reference--group-001.md#canonical-2033101022102200-0133122323301202-3002301223233310-1000201121102210-3312320233311011-1100013312313200-3303333011013222-1113032113313133) |
| `srv_pool.members.target` | [srv_pool.members.target](data-sources--dns_lb_pool--reference--group-001.md#canonical-2021220132122122-1332101011230213-0103222303311130-0200130210113211-1121123100313322-2222002302332131-3330233022213031-3301003133312300) |
| `srv_pool.members.weight` | [srv_pool.members.weight](data-sources--dns_lb_pool--reference--group-001.md#canonical-1111101332011212-3230212133010323-3212222113013110-1233213103101310-0230213120322012-1233230011210223-1220002202331223-0002103232012213) |
| `ttl` | [TTL](data-sources--dns_lb_pool--reference--group-001.md#canonical-0033113021233003-0133213121121032-2023332223201320-3311101012103213-3101120133312321-0333012203223311-3130103111312300-1132310301220201) |
| `use_rrset_ttl` | [use_rrset_ttl](data-sources--dns_lb_pool--reference--group-001.md#canonical-3001211201122102-3330232120003112-3211312103210012-0112200000110020-1332112311031231-0211010213230111-1231320310212332-2222030020130330) |

<a id="canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- a_pool

<a id="canonical-3110133003130222-1021132100113311-3331330212300301-3100002212022112-3200011232323111-2332211301132101-1213013312312012-1102311230111130"></a>

Type: `"single"`. Computed.

\[OneOf: a\_pool, aaaa\_pool, cname\_pool, mx\_pool, srv\_pool\] Pool for A Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

OneOf alternatives in this subsection:

- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3110133003130222-1021132100113311-3331330212300301-3100002212022112-3200011232323111-2332211301132101-1213013312312012-1102311230111130)
- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3331333332133320-0322212213320112-0112131123311322-1123122031020212-3310020321110213-2301023312223320-2113010122302330-0223202221012201)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-2010322303230133-3010012212113232-3200112321212311-2223020101312321-0013120001231120-1221003303133011-1213100012111320-0100311321020333)
- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1103013221030002-1110231321000100-0130202221021310-1310122223103021-1300123233202211-3330122203300012-3013122300220031-3031002220132331)
- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-0233230131013131-3113222100133202-0320130302200322-1223023300133003-0123301120002202-0233201320130300-3311320020232022-3223111032021133)

Select alternatives according to the provider validators above.

<a id="canonical-0323331231330101-0313030220211303-1101032130032023-2230231131101220-0221133111123121-2000311121033223-3122122230122230-0132120113110122"></a>

### Direct properties for `a_pool`

- [disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-1133112221100130-3123112132301222-2303322322303130-0030030211133022-1231000222212312-3133310100120330-1111120022331203-2222133333033123): complete subsection reference.

- [health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-3213113202220221-2310313301223132-3323301010323001-3310333132133013-3202203112020310-1332210002012201-3232312312030121-1132003233221320): complete subsection reference.

<a id="canonical-2201303130211332-2110010323220111-1022101300110102-2101130312323322-2312213120022220-2023020231202121-3232331302302330-1220310221321003"></a>

<a id="canonical-2032133310033112-1032102223220331-0131113222303312-2120210212130023-2332010102123002-1130002102021101-3130102200231111-3213012223131100"></a>

#### `a_pool.max_answers` property

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-1323202033230000-3122132123112000-3103332322202001-0113201131202111-2233131212313310-1001323310300031-3232331210122223-3310310233321203): complete subsection reference.

<a id="canonical-1133112221100130-3123112132301222-2303322322303130-0030030211133022-1231000222212312-3133310100120330-1111120022331203-2222133333033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool.disable_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300)
- a_pool.disable_health_check

<a id="canonical-0031100231010012-0002120013110023-0100121003101213-0321213201023033-1111213200000113-2300132333000101-2123232201033201-1012311332230311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable health check.

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

<a id="canonical-3213113202220221-2310313301223132-3323301010323001-3310333132133013-3202203112020310-1332210002012201-3232312312030121-1132003233221320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool.health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300)
- a_pool.health_check

<a id="canonical-3120111310311332-0303010031201131-0112120012000033-3212131233300002-2221111200032202-0222102300202101-0321120332022212-3232303300111120"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3120210013210030-3230010231011331-3022023032010313-1023000203300332-0013301230332311-2331021030223232-1223100003110032-0030230323023201"></a>

### Direct properties for `a_pool.health_check`

<a id="canonical-1331312230123221-2120232130120002-0101103013012112-0130322212210200-2323210313120213-2321321211010100-0223100323230302-0030020233231231"></a>

#### `a_pool.health_check.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1232300231120212-1332100231200330-0100210001000233-0111310123111131-1233213301333200-2022223313130300-3003132113333110-0120303120103131"></a>

<a id="canonical-3103030222013101-2013322230210323-2200013333103111-1222321203222201-3230003223313322-2132122323032310-3022302321111203-2111323111032101"></a>

#### `a_pool.health_check.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2121323013301023-2210033222112232-1213110031233310-1111123222032312-3000013233333023-0330321213013131-0231120333102212-2102130221333120"></a>

<a id="canonical-0131132003230311-2230032031102313-3302311003101111-1133302111111202-0011110010313321-1131130311121103-2310301131012102-0100100332022102"></a>

#### `a_pool.health_check.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1323202033230000-3122132123112000-3103332322202001-0113201131202111-2233131212313310-1001323310300031-3232331210122223-3310310233321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `a_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [a_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3111123132023110-0033301321322121-0320203222211000-2010033200211202-3101101202310110-0310100033023000-1133121301131210-3310320300121300)
- a_pool.members

<a id="canonical-2201221002100222-0333330320222310-1211323130302203-1231002203320132-2323301303121031-2203221120113032-0332121201000322-0031232133311231"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2131313011221201-1100233020020131-0321100012031011-3302021131311013-1101131011220032-1120212230310310-0301032022222031-0213102332210133"></a>

### Direct properties for `a_pool.members`

<a id="canonical-2212033021101300-2231222101130002-1201131200302103-0112002320123111-2003123331233300-2133200310320233-2303000111323003-3003220110131312"></a>

#### `a_pool.members.disable_spec` property

Type: `"bool"`. Computed.

Value of true will disable the pool-member.

<a id="canonical-3131100313223120-3302033132030323-1130332221022120-3210020020101111-0232213013123100-2302213323231330-1031330100121031-3300320121010132"></a>

<a id="canonical-1132230031233230-1212101103320213-1313301032111000-0123221100231131-0000013032130023-3222301232323302-2103313332223121-0231022122121020"></a>

#### `a_pool.members.ip_endpoint` property

Type: `"string"`. Computed.

Public IP. Public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3000312003313230-2212130320131123-3202123211301332-2011030213320102-1033030200031202-1322210332032011-2101031113010133-1120223331031010"></a>

<a id="canonical-3131102032111122-1121320333333233-3123020200123023-1203022003302002-1232022203002013-3023210022012033-3300121303023232-2313032003012322"></a>

#### `a_pool.members.name` property

Type: `"string"`. Computed.

Name. Pool member name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3022300302032000-3102033102113103-3333310203013132-0100032022103223-0320101310031102-2323031101132003-0300131331132131-0131221332002330"></a>

<a id="canonical-1302220023332233-1132001133322223-1322103010122333-1121120223022132-2102301313010013-0001312020120231-0031302203101133-1223311113131203"></a>

#### `a_pool.members.priority` property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-2310323210012100-1232203022301022-0212332022121132-0300013123200022-3101231311200021-3300001322232302-3022333310130112-2233103330203122"></a>

<a id="canonical-2121020032123032-0130022001010313-3010003133001020-1033103110131310-1230313023031332-1311122002102223-0202033111122333-3232330123333310"></a>

#### `a_pool.members.ratio` property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-3113200321003123-3322202013011013-3313213000031121-1133010331020321-1313120230301331-1310313133032133-3331032311202021-1201020223023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aaaa_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- aaaa_pool

<a id="canonical-3331333332133320-0322212213320112-0112131123311322-1123122031020212-3310020321110213-2301023312223320-2113010122302330-0223202221012201"></a>

Type: `"single"`. Computed.

Pool for AAAA Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3310230020313210-0002112111122100-2231302001011021-0113303301300113-1001332130012000-0212233301300112-2203012323212300-3122012000010133"></a>

### Direct properties for `aaaa_pool`

<a id="canonical-0033310331011313-2011313132121013-0223120232130310-1010110022220232-2101113021123121-1110200123100212-0003112310033311-3122122223300112"></a>

#### `aaaa_pool.max_answers` property

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-3330231313131112-2230323222113321-0232330201121330-1021130323222322-3001220022312012-1323112332011303-2003132101303123-0233222033231313): complete subsection reference.

<a id="canonical-3330231313131112-2230323222113321-0232330201121330-1021130323222322-3001220022312012-1323112332011303-2003132101303123-0233222033231313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `aaaa_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [aaaa_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-3113200321003123-3322202013011013-3313213000031121-1133010331020321-1313120230301331-1310313133032133-3331032311202021-1201020223023122)
- aaaa_pool.members

<a id="canonical-2032221233203020-3103321331030012-2203312132302232-3322330123120233-2012023123023021-2310222321301323-1132033310100021-0033212023011312"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3312211200203322-0233200323333223-3322313203111130-1212122230323331-2001223302003231-3323231321213103-0322032312220302-1223313003301303"></a>

### Direct properties for `aaaa_pool.members`

<a id="canonical-1113312303322233-1203112112322120-1033332021122303-2313330303131230-1321321121310121-2322220313020030-0121022302000010-0332320112313313"></a>

#### `aaaa_pool.members.disable_spec` property

Type: `"bool"`. Computed.

Value of true will disable the pool-member.

<a id="canonical-2201103221131233-0333032132233123-1321332112203310-1022101210331023-0033011231032233-2123113322131301-1130332233211233-1022102302113321"></a>

<a id="canonical-1031103030331310-2002022112121302-1221230033310022-1301113132232031-0211200320120321-1202230020103322-2233312003130030-1110320310120022"></a>

#### `aaaa_pool.members.ip_endpoint` property

Type: `"string"`. Computed.

Public IP. Public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3032013130013031-0121333120331020-3213200002323303-3013032011012112-2333120220213110-0330233121012031-2011112230313123-2200023323110023"></a>

<a id="canonical-0030201311032123-3202212212231122-3331023202131101-0322033322113010-1012211332112311-2230211212022202-2200303211321202-0231131122312212"></a>

#### `aaaa_pool.members.name` property

Type: `"string"`. Computed.

Name. Pool member name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-0121303210012132-3303122021110232-0323210113121332-2332001033230133-1311023222203121-3123320101120201-2012320003113330-1233221231011301"></a>

<a id="canonical-3321330323331023-1313003110001010-2122000200130031-3100103213101122-0001100132232212-0111320213320022-1230302311211301-1311233320132322"></a>

#### `aaaa_pool.members.priority` property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-0122113021312102-0212132213100210-0000131232122333-0010312133312120-1011002203003323-3111121322212000-0031121013102133-2002332033301213"></a>

<a id="canonical-0201002303102213-1102231100302020-1232102121010201-3113230210333331-1330122023133102-1023323303330213-0203303303311103-3102132033110122"></a>

#### `aaaa_pool.members.ratio` property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- cname_pool

<a id="canonical-2010322303230133-3010012212113232-3200112321212311-2223020101312321-0013120001231120-1221003303133011-1213100012111320-0100311321020333"></a>

Type: `"single"`. Computed.

Pool for CNAME Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-health_check_choice": "[\"disable_health_check\",\"health_check\"]"
}
```

<a id="canonical-2331310023330213-0321133200002101-1230311210000223-3213112110023100-3121130113210221-2021033021133030-3201030030223013-0311332331100023"></a>

### Direct properties for `cname_pool`

- [disable_health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-1201123103102100-3223010111202320-3230321311103223-2223101311201330-0321230130122011-2223213000201203-1031320101032012-3213220012012233): complete subsection reference.

- [health_check](data-sources--dns_lb_pool--reference--group-001.md#canonical-0133233310322030-1131110213032322-0111112003223311-2110013221112111-1022112221030212-0321230012302011-1030111302231313-3313013022133022): complete subsection reference.

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-0000213312220302-1011113232232320-0133221000321100-0121330130211212-0322202131221000-2013230022132131-1000102103303232-3021200300131033): complete subsection reference.

<a id="canonical-1201123103102100-3223010111202320-3230321311103223-2223101311201330-0321230130122011-2223213000201203-1031320101032012-3213220012012233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool.disable_health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211)
- cname_pool.disable_health_check

<a id="canonical-1313122120131121-0010020131220232-3323033130211013-2100011133213110-0103203020030221-3303323012333021-3021321031231000-1101003010233002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable health check.

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

<a id="canonical-0133233310322030-1131110213032322-0111112003223311-2110013221112111-1022112221030212-0321230012302011-1030111302231313-3313013022133022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool.health_check` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211)
- cname_pool.health_check

<a id="canonical-3202031120111001-2210233122231022-0313033123112102-2112332332200211-0132312010203122-1312111012120233-3211202322030131-3302132122232222"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1003300122322021-3022001300310322-2010123210203121-1121112312322210-1321320222111332-3313110020203002-0302333103001100-3311001313233212"></a>

### Direct properties for `cname_pool.health_check`

<a id="canonical-2012210332202320-0111133313002133-0313310212203231-0321011103321323-2020321102131003-3012220331112112-0031312131302022-1323112312010302"></a>

#### `cname_pool.health_check.name` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0021021323201120-3331101333020020-1003222000232131-1211213113032000-2123331333322023-1000230312331111-3312220030330301-0101223333111101"></a>

<a id="canonical-0231320302011022-1332032132330323-3320011001020011-0032231233231303-3300120012122103-3002103331302213-1032131113031130-0012132032332130"></a>

#### `cname_pool.health_check.namespace` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3312201130333321-2331012101330321-2110011033233233-2233032222112332-3031032012001330-0121202200210222-2102000033202222-1110231012201220"></a>

<a id="canonical-1220233101103331-0233131012110110-0212102131330122-3213221230112313-3101020033120212-2213032300103203-0100301200110131-0012120003111003"></a>

#### `cname_pool.health_check.tenant` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0000213312220302-1011113232232320-0133221000321100-0121330130211212-0322202131221000-2013230022132131-1000102103303232-3021200300131033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `cname_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [cname_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1221101101101203-1232231130303113-3233111001012331-0223200211330133-0000031100033023-1121001003332323-0000203310233120-0300332003111211)
- cname_pool.members

<a id="canonical-3022200313100231-1100233233103313-2132231111031000-1332320021231230-1131002132100103-0113022223030231-2012212323323023-3003212210323303"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3220031003203023-2001030000022100-2321032321032200-1033223001111331-1100031222232300-0330031201300031-1100220110031200-1333212330332231"></a>

### Direct properties for `cname_pool.members`

<a id="canonical-3302313311020030-2312002331121230-2222022033213113-0001003122033033-2201311312112011-0232301230122011-2023010332230310-0211303331212103"></a>

#### `cname_pool.members.domain` property

Type: `"string"`. Computed.

Specifies the fully qualified domain name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-2232333022232010-3022231010122301-2020121021230030-0002320301332001-2313233002201032-2120001331133231-1110002012103322-1200202020201310"></a>

<a id="canonical-0323332001310300-1303123322120101-1020321101330213-1333320100130032-2233203312222012-2132221223213330-0003212322330031-3301030020012023"></a>

#### `cname_pool.members.final_translation` property

Type: `"bool"`. Computed.

If this flag is true, the CNAME record will not be translated further.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2212132012332023-3123020002300213-0321322202203322-2123101233230022-1330320021313010-1302012332120011-2001103232030232-1200123231130133"></a>

<a id="canonical-1311302300210300-0100032120112123-1303223211113311-1020320031032322-1212233311321120-2331302013101100-3213221303333030-0032031013023010"></a>

#### `cname_pool.members.name` property

Type: `"string"`. Computed.

Name. Pool member name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3201213022213111-2122302110010310-0321130202023303-2310212102000313-0130231031323100-0131211120210132-3302113102200203-0112003110200102"></a>

<a id="canonical-0323310023000202-0020311323023121-3322033002202230-2233001020310330-2303310013223211-1320122301002021-3213302003333201-2002302230001131"></a>

#### `cname_pool.members.priority` property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Priority. Determines the order in which traffic is
routed to pool members. The lower the number, the higher the priority, making those members active
while higher-numbered members act as backups.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "255"
  }
}
```

<a id="canonical-3200133000313201-1330013333303231-1302303310211123-1033011332312323-0322132230310123-1003311223122223-0301312121003123-1230122112122313"></a>

<a id="canonical-0123111303221333-0033232011332130-2300123002132201-0211001221103211-1213200032223011-1302112031322100-1112223201133233-3112002122321211"></a>

#### `cname_pool.members.ratio` property

Type: `"number"`. Computed.

Used if the pool’s load balancing mode is set to Ratio-Member.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-0222213301012121-1331313201223331-3230012302232130-1013220221311321-3332133002323312-1320012231332301-2220233021131232-2230101013222001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mx_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- mx_pool

<a id="canonical-1103013221030002-1110231321000100-0130202221021310-1310122223103021-1300123233202211-3330122203300012-3013122300220031-3031002220132331"></a>

Type: `"single"`. Computed.

Pool for MX Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2003203013200132-3231301031221330-1200011103220122-1032203203132323-2121001320000020-2320300323332303-3301311020232300-2202131023001330"></a>

### Direct properties for `mx_pool`

<a id="canonical-3113231200203220-1311022201302020-0012012111320011-1003210021001221-2202023231310332-3130003320113213-1202003031220113-3322112103212123"></a>

#### `mx_pool.max_answers` property

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-3023221002033333-3232100131120122-0301000012212322-0003001123112213-3212210200012321-0031220002200321-2100301333101023-1032233113001220): complete subsection reference.

<a id="canonical-3023221002033333-3232100131120122-0301000012212322-0003001123112213-3212210200012321-0031220002200321-2100301333101023-1032233113001220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mx_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [mx_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-0222213301012121-1331313201223331-3230012302232130-1013220221311321-3332133002323312-1320012231332301-2220233021131232-2230101013222001)
- mx_pool.members

<a id="canonical-2112210003220031-1022213020322003-3022110131313310-2222000201202332-0120003220200322-2002123210303312-3001220001100212-2011213002323033"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1010302310121323-1323021012221123-3022210310010223-1202212000321032-2103312111221213-0011300300120012-0321210031223102-1330330010303311"></a>

### Direct properties for `mx_pool.members`

<a id="canonical-2110210001002220-0303133112203020-0033203100233203-3131211023222121-3223031121300121-2322322302301322-2123322102211212-1003300132311220"></a>

#### `mx_pool.members.domain` property

Type: `"string"`. Computed.

Domain name for routing and identification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-0101232032320202-3103001101100223-0100021100020201-0223123332102123-3312303101020232-1121223132331222-3210320313121130-3231300231213131"></a>

<a id="canonical-2132300312321111-2101301101111032-1320120000110033-0223301113302301-3311033122303102-0232330100311121-1131111203300333-3303213000000313"></a>

#### `mx_pool.members.name` property

Type: `"string"`. Computed.

Name. Pool member name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-3000213031101300-1220020011030300-0332203233123211-1331012212032002-0102111023111232-0021020233032322-3201101002120011-3213113321200212"></a>

<a id="canonical-2330233001023220-0223311231010332-0103200232103213-3233113021121212-0003131212233022-1230303301100223-2103212233121133-3021200033121222"></a>

#### `mx_pool.members.priority` property

Type: `"number"`. Computed.

MX Record Priority. MX Record priority.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3232222003000233-2321303102103331-2133030322303322-2213303233300103-3230020011333103-1210320231112133-0013001013331011-1331132201023112"></a>

<a id="canonical-3021023333103013-1100002201012230-2233033202202110-2022200032213333-2201302000320103-3202120023101321-3023221232232020-2031130303121301"></a>

#### `mx_pool.members.ratio` property

Type: `"number"`. Computed.

Load Balancing Ratio. Load Balancing Ratio.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-1033101312200110-3320312032020203-1322213322022011-2111022211202232-2101111013330231-1000030101312201-0203211113230302-3310203000232311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `srv_pool` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- srv_pool

<a id="canonical-0233230131013131-3113222100133202-0320130302200322-1223023300133003-0123301120002202-0233201320130300-3311320020232022-3223111032021133"></a>

Type: `"single"`. Computed.

Pool for SRV Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0113320121211121-1223001112003133-1122002110031212-1132300123000102-3001231011231013-2132231312111313-1333030210011332-1320102203310022"></a>

### Direct properties for `srv_pool`

<a id="canonical-1211032202013310-0200303023131203-0230333222110122-2222122232302210-1331001112200030-3232330231030301-3231012133131011-1220312311010021"></a>

#### `srv_pool.max_answers` property

Type: `"number"`. Computed.

Limit on number of Resource Records to be included in the response to query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

- [members](data-sources--dns_lb_pool--reference--group-001.md#canonical-3031023030220301-3022123033323001-0301300001121231-3000313222022100-1023111122233311-1332221313300122-0223023303212221-1030323202331332): complete subsection reference.

<a id="canonical-3031023030220301-3022123033323001-0301300001121231-3000313222022100-1023111122233311-1332221313300122-0223023303212221-1030323202331332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `srv_pool.members` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- [srv_pool](data-sources--dns_lb_pool--reference--group-001.md#canonical-1033101312200110-3320312032020203-1322213322022011-2111022211202232-2101111013330231-1000030101312201-0203211113230302-3310203000232311)
- srv_pool.members

<a id="canonical-3213123332003032-1111331012012202-0022210301122202-3301110100223111-2020220213031330-1223010000121202-1203023232313010-1032112030210112"></a>

Type: `"list"`. Computed.

Pool Members. Configuration parameter for members

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1323113203213002-2210311203020113-1231213133202331-3002320313220101-1031230002320313-0131111120312013-1201031312303033-2323331221200323"></a>

### Direct properties for `srv_pool.members`

<a id="canonical-1100310100133200-0103312022233130-1222322320121123-3301222013112110-0310310003322103-1231111021323022-3232301230300031-2110211123312111"></a>

#### `srv_pool.members.final_translation` property

Type: `"bool"`. Computed.

If this flag is true, the SRV record will not be translated further.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2302221332300312-2300221133302313-0301020003310111-0122031320023123-2301023220031202-3123202030212010-1100012100021231-3021132232302230"></a>

<a id="canonical-0302111332201010-0300321321333311-3302223313323130-2013010121213213-2313323323110031-1311320113302131-3011320230213031-2121303321321132"></a>

#### `srv_pool.members.name` property

Type: `"string"`. Computed.

Name. Pool member name.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2311310102100120-2232003021301133-0313311232103323-3210310301121331-0333111323212213-3112103303011223-2022032210032113-2210222030301001"></a>

<a id="canonical-0302213300321313-2311211001330112-3110131231203212-0302232132300201-0100102231021323-2312123021312302-1201212100132102-3031212203023033"></a>

#### `srv_pool.members.port` property

Type: `"number"`. Computed.

Port. Port on which the service can be found.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1132111011022311-1030013110301333-3302120233002201-0212100212212310-0000113030310002-0313001230222222-0213100103301010-1203030100233211"></a>

<a id="canonical-0221320012211123-1223012202011032-3230330213221221-3303030303132111-3323302121120011-0221222110331222-2222133012321123-1011022302132303"></a>

#### `srv_pool.members.priority` property

Type: `"number"`. Computed.

Priority of the target. A lower number indicates a higher preference.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-2033101022102200-0133122323301202-3002301223233310-1000201121102210-3312320233311011-1100013312313200-3303333011013222-1113032113313133"></a>

<a id="canonical-1231120221330323-1020112113310301-2102102020311313-1332302003103103-0232123221312302-2313012332213003-3010003310323032-2022200202213223"></a>

#### `srv_pool.members.ratio` property

Type: `"number"`. Computed.

Load Balancing Ratio. Configuration parameter for ratio

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "100"
  }
}
```

<a id="canonical-2021220132122122-1332101011230213-0103222303311130-0200130210113211-1121123100313322-2222002302332131-3330233022213031-3301003133312300"></a>

<a id="canonical-0012002322222232-2333211001111303-0030112030031130-1201300133311012-0210012220202131-3133030321213113-3320222332220200-0120202132102302"></a>

#### `srv_pool.members.target` property

Type: `"string"`. Computed.

Domain name of the machine providing the service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-1111101332011212-3230212133010323-3212222113013110-1233213103101310-0230213120322012-1233230011210223-1220002202331223-0002103232012213"></a>

<a id="canonical-2311032103222221-3333011312021201-2012000022310222-2322200211330133-0210103213221200-0213321320332023-1001012111012100-0212011020300002"></a>

#### `srv_pool.members.weight` property

Type: `"number"`. Computed.

Weight of the target. A higher number indicates a higher preference.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minimum": 0,
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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3103203103023110-1213233013101233-0020122011003001-2311022001103201-0020331202131321-1010221101300002-1311010322110100-3230130000033321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_rrset_ttl` properties

Breadcrumbs:

- [xcsh_dns_lb_pool](../data-sources/dns_lb_pool.md#canonical-3331103211112201-2233210311233301-1133131300233210-0231310011012223-1221000112222301-0013203111112121-2202031122031100-2323323031110010)
- [Property reference](data-sources--dns_lb_pool--reference--group-001.md#canonical-2201003233003102-0220100013322230-3020001003132311-2102010200113333-2212332101121122-3000221210320102-0202303311010221-1213321012030033)
- use_rrset_ttl

<a id="canonical-3001211201122102-3330232120003112-3211312103210012-0112200000110020-1332112311031231-0211010213230111-1231320310212332-2222030020130330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for use rrset TTL.

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
