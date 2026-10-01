---
page_title: "xcsh_dns_zone reference"
subcategory: "DNS"
description: "Complete grouped canonical reference for xcsh_dns_zone reference."
---

# xcsh_dns_zone reference

<a id="canonical-2300010310330233-1112033021221223-0200300032223212-2301320321023223-1220130310033010-2013121013133321-1110123313323011-3011103101122230"></a>

## primary.rr_set_group.rr_set.cname_record — cname_record / 113202310213 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.cname_record

<a id="canonical-1030012213303033-0002303312011021-3111130311232311-0020011033100021-0332202101323301-3133002320230131-2222332200333102-0102302112003002"></a>

Type: `"single"`. Computed.

DNSCNAMEResourceRecord.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0232032012201123-2313212102030001-1302002131331131-3313002002203112-0120230232033110-3112210212320103-2023323011213031-0133001111010111"></a>

## Direct properties — cname_record / 113202310213 / 3

<a id="canonical-2003100302220033-3122311233212001-0113022310332003-3200210312121032-0000313000310133-0001200132032130-1301203111221020-0103323320203321"></a>

<a id="canonical-1022020333301013-1310321331201112-1132121020020332-3323112013032233-3123213332310323-0233112020033013-0321010020102002-1332332131233100"></a>

## name property — cname_record / 113202310213 / 4

Type: `"string"`. Computed.

CName Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3330130313321200-3212011323333300-3023021223312031-3210212311310232-2232213012131223-3301030123122211-2032112121113302-3021112323212230"></a>

<a id="canonical-0123203321010222-3230110110123003-0033210121111101-3222312102333033-3103011330101201-1001113112110132-3331301121003221-3230231220312211"></a>

## value property — cname_record / 113202310213 / 5

Type: `"string"`. Computed.

Domain. Configuration parameter for value

Upstream description:

Configuration parameter for value

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 255,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-2222222000222103-0331111213210010-0321120120101022-1302300123133133-2012223310033232-3211310133312310-0330213320010111-2233010231210301"></a>

## Next pages — cname_record / 113202310213 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010122323310300-2113311301211120-0321201223303021-1331320120130210-0102203200113321-3313123113201213-1013312233220231-0311101222203121"></a>

## primary.rr_set_group.rr_set.ds_record — ds_record / 012332003220 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.ds_record

<a id="canonical-3323123231232232-3121122132113033-2010032032232020-0223223101013113-1320002203003203-0021121122213032-0110113000320310-3131320002200321"></a>

Type: `"single"`. Computed.

DNS DS Record. DNS DS Record.

Upstream description:

DNS DS Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1003333020210111-1031200230320122-3323130320300122-0113230230011320-1030002320203033-1213213312032012-1310022010003011-1221303023310322"></a>

## Direct properties — ds_record / 012332003220 / 3

<a id="canonical-2120010001021333-2330312330231221-3013231302312000-1012330233333101-2112330010102121-3230301012220311-2211313001301132-3201002122332000"></a>

<a id="canonical-0213002313312322-0131130033301221-2103333203233132-2120321222220010-1131212211233313-1311232311011213-1220131123201320-1013010332121121"></a>

## name property — ds_record / 012332003220 / 4

Type: `"string"`. Computed.

DS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130): complete subsection reference.

<a id="canonical-3133221232321320-3312011212012033-1122231102212322-3112313321100121-2113213201301313-3020210323120202-2303031322321013-1221323111202012"></a>

## Next pages — ds_record / 012332003220 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002122012010222-2122222033221221-2133221330300201-2233020003322320-1302323103313213-0330100300132013-1202310301102021-3033113112332031"></a>

## primary.rr_set_group.rr_set.ds_record.values — values / 220011021233 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- primary.rr_set_group.rr_set.ds_record.values

<a id="canonical-2002210212313213-0103320021303300-3103231001322333-2111012132002320-3231020112223333-1320100103000123-0230020232113323-1030130320013033"></a>

Type: `"list"`. Computed.

DS Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2033032023103301-0211133333311103-2200011030132331-0300110203111212-0010330023130321-0332000220113210-3000133001222202-3322011011231000"></a>

## Direct properties — values / 220011021233 / 3

<a id="canonical-2000112311312320-0220020111313121-0200111002020332-1212301331001213-1313232133301232-1332113323012120-1122200133112113-2122201122320213"></a>

<a id="canonical-1113133213210232-0011113013210133-2311323130122100-3111033300212100-1001312202310213-0132312021312132-3031113123000001-0320221231101221"></a>

## ds_key_algorithm property — values / 220011021233 / 4

Type: `"string"`. Computed.

\[Enum:
UNSPECIFIED|RSASHA1|RSASHA1NSEC3SHA1|RSASHA256|RSASHA512|ECDSAP256SHA256|ECDSAP384SHA384|ED25519|ED448\]
DS key-value must be compatible with the specified algorithm. - UNSPECIFIED: UNSPECIFIED - RSASHA1:
RSASHA1 - RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1 - RSASHA256: RSASHA256 - RSASHA512: RSASHA512 -
ECDSAP256SHA256: ECDSAP256SHA256 - ECDSAP384SHA384: ECDSAP384SHA384 - ED25519: ED25519 - ED448:
ED448. Possible values are \`UNSPECIFIED\`, \`RSASHA1\`, \`RSASHA1NSEC3SHA1\`, \`RSASHA256\`,
\`RSASHA512\`, \`ECDSAP256SHA256\`, \`ECDSAP384SHA384\`, \`ED25519\`, \`ED448\`.

Upstream description:

DS key-value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIED: UNSPECIFIED

&#8203;- RSASHA1: RSASHA1

&#8203;- RSASHA1NSEC3SHA1: RSASHA1-NSEC3-SHA1

&#8203;- RSASHA256: RSASHA256

&#8203;- RSASHA512: RSASHA512

&#8203;- ECDSAP256SHA256: ECDSAP256SHA256

&#8203;- ECDSAP384SHA384: ECDSAP384SHA384

&#8203;- ED25519: ED25519

&#8203;- ED448: ED448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIED",
  "enum": [
    "UNSPECIFIED",
    "RSASHA1",
    "RSASHA1NSEC3SHA1",
    "RSASHA256",
    "RSASHA512",
    "ECDSAP256SHA256",
    "ECDSAP384SHA384",
    "ED25519",
    "ED448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2230223230000300-1232232312112300-2032233211113100-0022102120111130-1120222113201100-2330003332012331-1323220230211111-0301333031331021"></a>

<a id="canonical-3310331001301130-0100221002003013-2211131100112301-1022222230223130-1202202013032101-0013032300121032-0212103213322213-1023003133110031"></a>

## key_tag property — values / 220011021233 / 5

Type: `"number"`. Computed.

Short numeric value which can help quickly identify the referenced DNSKEY-record.

Upstream description:

A short numeric value which can help quickly identify the referenced DNSKEY-record.

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

- [sha1_digest](data-sources--dns_zone--reference--group-003.md#canonical-3203221312201213-2031031221030023-3333333202131301-1331021012331001-2011030223330302-0223330221021222-1331321201021333-0102222111201331): complete subsection reference.

- [sha256_digest](data-sources--dns_zone--reference--group-003.md#canonical-2311203331210120-3312112321300113-1022020101003301-2131013301131322-0221033031112002-1122131211322321-0110130013320321-2121331211133023): complete subsection reference.

- [sha384_digest](data-sources--dns_zone--reference--group-003.md#canonical-3010013011001122-2212200033031212-2021022133102303-0230220332102103-1202131231131311-1000102210210112-2302010302202021-3000222221223232): complete subsection reference.

<a id="canonical-2323232222300013-2103001210032201-1130210302221223-2003300111020330-0310233010202002-1321033203320331-1300013013303210-1003330323121320"></a>

## Next pages — values / 220011021233 / 6

- [primary.rr_set_group.rr_set.ds_record.values.sha1_digest](data-sources--dns_zone--reference--group-003.md#canonical-3203221312201213-2031031221030023-3333333202131301-1331021012331001-2011030223330302-0223330221021222-1331321201021333-0102222111201331)
- [primary.rr_set_group.rr_set.ds_record.values.sha256_digest](data-sources--dns_zone--reference--group-003.md#canonical-2311203331210120-3312112321300113-1022020101003301-2131013301131322-0221033031112002-1122131211322321-0110130013320321-2121331211133023)
- [primary.rr_set_group.rr_set.ds_record.values.sha384_digest](data-sources--dns_zone--reference--group-003.md#canonical-3010013011001122-2212200033031212-2021022133102303-0230220332102103-1202131231131311-1000102210210112-2302010302202021-3000222221223232)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3203221312201213-2031031221030023-3333333202131301-1331021012331001-2011030223330302-0223330221021222-1331321201021333-0102222111201331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301131310331213-3311122233111100-1301021000323100-2000330131033131-2221310310332012-3313013300101131-3322303033111220-3202102310210023"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha1_digest — sha1_digest / 301223013032 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- primary.rr_set_group.rr_set.ds_record.values.sha1_digest

<a id="canonical-1311001023013010-2120332201020300-0330303232223012-0003321111113220-3233132311300323-0331100222123203-0010100312122300-0202111211311011"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 digest.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0231200122031133-1111201122330112-1120120201032311-3203032123233001-0122031300011012-3220212003120313-2333322013012221-1320013310013023"></a>

## Direct properties — sha1_digest / 301223013032 / 3

<a id="canonical-0211022311323301-2333203131012220-1031110301003033-1132331011133230-3030133201011313-2100123333112000-1122031102003013-1032211121233101"></a>

<a id="canonical-0201323221022200-3323320111012111-1122332330101232-0021313120000300-1003322201002132-2122132203113331-0003321103300222-1320033110233012"></a>

## digest property — sha1_digest / 301223013032 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-2331332022222011-0310331223032031-1303010323033231-1012132033301112-0132211232321101-0303023300032130-3100123322033330-3132103113221232"></a>

## Next pages — sha1_digest / 301223013032 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2311203331210120-3312112321300113-1022020101003301-2131013301131322-0221033031112002-1122131211322321-0110130013320321-2121331211133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010013121113201-1110123203311002-3200232310030132-0203331133310101-0000133312202312-2022212133320121-3220202112230112-1131301102021231"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha256_digest — sha256_digest / 133120222302 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- primary.rr_set_group.rr_set.ds_record.values.sha256_digest

<a id="canonical-3120333302320123-1021333022223330-3302111213102300-2111313030123132-0303031303033232-0210010310002012-0212000111133301-2111312021021231"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 digest.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2200130222322030-1311322213333013-1012202023230323-1301123231100202-3133030003321223-3013323211122203-1232212122110123-1033020013322223"></a>

## Direct properties — sha256_digest / 133120222302 / 3

<a id="canonical-0000202333011311-2211122320113121-2311200101111101-3200233123201201-2222202310123300-3133121112121013-1213110333131131-3032030033221030"></a>

<a id="canonical-2110321133213100-2131023103320113-3122022020023220-2303001202231222-3131330232230222-0320011031011023-2213023102113211-1000332223012303"></a>

## digest property — sha256_digest / 133120222302 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-2322221311221300-2200132113020211-3221230032302100-3303321113010332-2031110302323011-0003012323123021-1333001230111221-3121300303232332"></a>

## Next pages — sha256_digest / 133120222302 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3010013011001122-2212200033031212-2021022133102303-0230220332102103-1202131231131311-1000102210210112-2302010302202021-3000222221223232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103020311002013-0013111103223123-3023313320333313-0202202111002100-3310313211131112-0300001131321331-3111130012300132-3032102103003303"></a>

## primary.rr_set_group.rr_set.ds_record.values.sha384_digest — sha384_digest / 113213322311 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.ds_record](data-sources--dns_zone--reference--group-003.md#canonical-0003002010030223-3031302330310202-3322320001011330-0003010123133202-3023120301200122-3120120121330033-2203311132332022-2031233200112331)
- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- primary.rr_set_group.rr_set.ds_record.values.sha384_digest

<a id="canonical-3003231321331023-2130302302202211-1301303101130013-1003130313231020-1111332313312312-1033021122012133-3101033031223101-2333020010213003"></a>

Type: `"single"`. Computed.

Configuration parameter for sha384 digest.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1210211233200012-1321321022322010-2331332233222120-2001011123032011-2330313112133302-2311021132200100-0131200303123101-0100000220023102"></a>

## Direct properties — sha384_digest / 113213322311 / 3

<a id="canonical-1101221320221223-1231332032211112-3231113020303321-0300321103023000-1133300001133003-3132032231213321-0001011320020302-2303232101212231"></a>

<a id="canonical-3121221101131132-0103010213330121-0322122322031323-1313323233011130-0010033213102323-3230020130123323-3331330010311332-2113232022312122"></a>

## digest property — sha384_digest / 113213322311 / 4

Type: `"string"`. Computed.

The 'digest' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 96,
  "minLength": 96,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 96,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 96
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "96",
    "ves.io.schema.rules.string.min_len": "96"
  }
}
```

<a id="canonical-1011200232120313-1232032000332021-0302003103211032-0301221312001123-3001133030220133-0021202013021310-0322130131301310-2212023201101210"></a>

## Next pages — sha384_digest / 113213322311 / 5

- [primary.rr_set_group.rr_set.ds_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0321330012321211-3123322210321000-1031221223211332-3030233112100300-0330320222303223-1200332131011201-3211322103013330-3032103002302130)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1300212020003230-1113311332101220-1132330003331301-3333111020100102-3003313132131220-0122230031303231-2203031223032120-2332332003333102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133202001213321-3230312233020033-0313003320203001-2120220100221101-0220133021131100-2111121301133122-2232203322323330-3202133311301323"></a>

## primary.rr_set_group.rr_set.eui48_record — eui48_record / 233122121321 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.eui48_record

<a id="canonical-2001000003012201-0103211033302211-3210331303310200-0132323302313312-2131011003203221-3101321110100030-1212121121002232-2010002212211130"></a>

Type: `"single"`. Computed.

Configuration parameter for eui48 record.

Upstream description:

DNS EUI48 Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0013213102312113-3230003232201320-1121231111323033-0301312002021103-1022233231333123-3310311320331002-1033310031112000-3220320301320112"></a>

## Direct properties — eui48_record / 233122121321 / 3

<a id="canonical-3130212232120323-1122323212113010-0001012120200023-0201130300120321-2230111012303020-1110303000001301-0000030020001230-3011110120212102"></a>

<a id="canonical-2021223210330232-3331032230313203-3131103332021313-3031202212332200-2323321133112221-3202111020303131-2233002223212203-2131302101233122"></a>

## name property — eui48_record / 233122121321 / 4

Type: `"string"`. Computed.

EUI48 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-1121220222210001-2202001311302222-3222323232333211-2020031120112200-2111201311232301-1330201121301331-3302310300220021-1123133321323333"></a>

<a id="canonical-3311303202001100-0233201023102123-3322302103200221-2310223210020103-2111220101223033-3113103322200310-1123222233112222-0222002333122313"></a>

## value property — eui48_record / 233122121321 / 5

Type: `"string"`. Computed.

EUI48 Identifier. A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Upstream description:

A valid eui48 identifier, for example: 01-23-45-67-89-ab.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 17,
  "minLength": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 17,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 17,
    "pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "17",
    "ves.io.schema.rules.string.min_len": "17",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){5}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-0000321220101113-0232303100202211-2032333103313321-3321022323122203-1011220103331323-0101332013200320-2021322030202102-0130020301210232"></a>

## Next pages — eui48_record / 233122121321 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0003303213311201-3003200303313223-1121023300112121-2002211020100233-1112212221032122-1331101211011320-2003332313321333-2110200232323101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212322311103022-2000003102111333-0202110300302330-3122202321312313-0132120003322113-1001300030101230-0221003031302312-2010312321310321"></a>

## primary.rr_set_group.rr_set.eui64_record — eui64_record / 301120213330 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.eui64_record

<a id="canonical-1212201231123200-0110310002232221-1000113123120321-1200121211201322-1123100112113002-1230103230003231-3122120313201201-3133223022301123"></a>

Type: `"single"`. Computed.

Configuration parameter for eui64 record.

Upstream description:

DNS EUI64 Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2020233101313132-2332121100223221-2300321300003001-0022120130233301-2311330120212133-2011133201130101-2332031222303123-2003320110230100"></a>

## Direct properties — eui64_record / 301120213330 / 3

<a id="canonical-3223031003000022-1301323031312210-1302322221021013-1133310210310132-1220023121201013-0220301331002032-0330103110030000-1211221101132120"></a>

<a id="canonical-2111213100302213-2303233233011320-0013131133002103-0000203020120310-1112100312332330-0202221313302300-1331202313302033-0310330302303130"></a>

## name property — eui64_record / 301120213330 / 4

Type: `"string"`. Computed.

EUI64 Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

<a id="canonical-2331200101303131-2201302132103211-1310001230130200-3321132230320223-3133322100001223-1011002231031201-3112211203302113-1311103100303112"></a>

<a id="canonical-2101230123323203-3330211103203101-1200023031202203-2113013232210221-0321122213213232-2200122130102131-1203032023222030-2113322113000123"></a>

## value property — eui64_record / 301120213330 / 5

Type: `"string"`. Computed.

EUI64 Identifier. A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Upstream description:

A valid EUI64 identifier, for example: 01-23-45-67-89-ab-cd-ef.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 23,
  "minLength": 23,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 23,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 23,
    "pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "23",
    "ves.io.schema.rules.string.min_len": "23",
    "ves.io.schema.rules.string.pattern": "^([0-9A-Fa-f]{2}-){7}([0-9A-Fa-f]{2})$"
  }
}
```

<a id="canonical-0302020030333132-0033313323201113-1133331020310100-0222132301010012-3322230132230301-2221200331030332-3020130303332312-1331033222011211"></a>

## Next pages — eui64_record / 301120213330 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021132233131202-2102113103011130-3323222100111222-1330133102130113-0202003222002112-1110301113202330-3031001022210301-3202313002221102"></a>

## primary.rr_set_group.rr_set.lb_record — lb_record / 121313110023 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.lb_record

<a id="canonical-1010002011111230-0030000301202311-2013002322300101-3033112001200130-3300230103231010-2032221112102132-3233310201023023-1110101302322332"></a>

Type: `"single"`. Computed.

DNS Load Balancer Record. DNS Load Balancer Record.

Upstream description:

DNS Load Balancer Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1020233312022332-1300121312312301-3210300202222000-0012112300001032-1020221113220201-1312001103330103-3120003023033102-1222231112333021"></a>

## Direct properties — lb_record / 121313110023 / 3

<a id="canonical-0000321020211220-0303010000002220-0212332120131033-2232210021012112-3213233323302222-2022012000131232-0211310122022312-0213101122333323"></a>

<a id="canonical-2321102220330032-0120223010330022-0011230211221110-1132230131331010-1303033322032210-2320001002100123-2031010223310000-2220033332211220"></a>

## name property — lb_record / 121313110023 / 4

Type: `"string"`. Computed.

Load Balancer record name (except for SRV DNS Load balancer record) should be a simple record name
and not a subdomain of a subdomain.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
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
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

- [value](data-sources--dns_zone--reference--group-003.md#canonical-0121010122012212-1330211002022212-1012203201000101-3332330100222032-3201223332222113-0021233010201332-1112313320320122-2313332311002133): complete subsection reference.

<a id="canonical-2311111131301120-1223323303220102-0211101012010200-0223211311322331-1122333200300122-1032222311302013-3033021020033313-0202331113223023"></a>

## Next pages — lb_record / 121313110023 / 5

- [primary.rr_set_group.rr_set.lb_record.value](data-sources--dns_zone--reference--group-003.md#canonical-0121010122012212-1330211002022212-1012203201000101-3332330100222032-3201223332222113-0021233010201332-1112313320320122-2313332311002133)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0121010122012212-1330211002022212-1012203201000101-3332330100222032-3201223332222113-0021233010201332-1112313320320122-2313332311002133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222111222333302-1001022313120120-1003002220320132-1111023003222223-1131232210332022-0213103321031310-1132122213122222-1132301200210200"></a>

## primary.rr_set_group.rr_set.lb_record.value — value / 110102100011 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121)
- primary.rr_set_group.rr_set.lb_record.value

<a id="canonical-3121320223110223-1332123103002222-2313010003220230-2012231231333302-1010223011211013-2321221233000212-0331033121313101-0023233010031111"></a>

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

<a id="canonical-3131320133330101-2000310232213010-3313230330112000-0011002002132211-2113300231203312-0100220212022101-0102231032131233-3320130021022021"></a>

## Direct properties — value / 110102100011 / 3

<a id="canonical-0000202220212313-3111321311123302-2131021122203101-3000230310120101-1112231002303110-1303023323321232-0013211232021121-2031223113331022"></a>

<a id="canonical-0230113230031230-3023112302213101-0123313133133301-0302033203202101-1033120230202232-1333201320111013-2123332110032230-2222233013233212"></a>

## name property — value / 110102100011 / 4

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

<a id="canonical-3313322000321000-3113213123012211-2032012223101033-2231033201320333-1310023100033020-0100030313121201-0210333130312303-2221022213230133"></a>

<a id="canonical-2133121133023313-3111321310320002-1013023303020311-2323022200013022-1321300021022323-1102200211210303-1012333213001012-1032022322022231"></a>

## namespace property — value / 110102100011 / 5

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

<a id="canonical-2330303013222230-0011223112102330-0022222023310020-2032231332020121-1232102130201023-1202203103200000-0203222211233301-3332202201032101"></a>

<a id="canonical-3100122222222223-2002203301133220-3122211132233001-0201202332102213-3101111121201221-1200000213031033-0313133133121221-1201011233221332"></a>

## tenant property — value / 110102100011 / 6

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

<a id="canonical-1033330321013013-0030332211232120-3211022113233232-1210113130201230-2121132112332021-0201120130313010-2121310033001022-3111213310233211"></a>

## Next pages — value / 110102100011 / 7

- [primary.rr_set_group.rr_set.lb_record](data-sources--dns_zone--reference--group-003.md#canonical-1230020033200103-1330330300123130-3220203300231002-1211332020202310-3103022120112000-2013013012212211-2200123220201232-2330030320121121)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120220103321202-1231313333023123-2122203122100011-0311313033210132-2210330202212130-1233022132330231-0033032220323311-0202101302311333"></a>

## primary.rr_set_group.rr_set.loc_record — loc_record / 301310302222 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.loc_record

<a id="canonical-2310233030001321-2302022120200120-2011323023012232-3330132222012231-3121102131100201-0332131311020001-2001331033212200-1023211110112202"></a>

Type: `"single"`. Computed.

DNS LOC Record. DNS LOC Record.

Upstream description:

DNS LOC Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3203201302330220-1200031311322132-3230110122233323-3233301232300303-1213112331321323-3201132303100111-2033102111331020-1320203033310202"></a>

## Direct properties — loc_record / 301310302222 / 3

<a id="canonical-3132112120333211-2233232033231211-1000331213201310-1101220230213201-0011032113331230-1030330122200111-1031020013333310-2333122001003212"></a>

<a id="canonical-0323211321200221-0321333121221000-3010210323302233-0310013322030312-0302131313002120-1321010123210312-2331201022001312-0230122201030113"></a>

## name property — loc_record / 301310302222 / 4

Type: `"string"`. Computed.

LOC Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-2330323032333122-2133122122112213-0010023311010233-1013321030002213-1010011203220211-1333111330011310-0110333330333322-2302320200003101): complete subsection reference.

<a id="canonical-0332120013310331-0133100033110313-0122032032223030-0130230020310123-1222321122121201-3032030033021021-3332220313202203-1111122202132113"></a>

## Next pages — loc_record / 301310302222 / 5

- [primary.rr_set_group.rr_set.loc_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2330323032333122-2133122122112213-0010023311010233-1013321030002213-1010011203220211-1333111330011310-0110333330333322-2302320200003101)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2330323032333122-2133122122112213-0010023311010233-1013321030002213-1010011203220211-1333111330011310-0110333330333322-2302320200003101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201230022021200-3311221003223120-1002321322011113-1121131101330233-1132200003323032-2133120113001102-0023302223211013-1312220233010310"></a>

## primary.rr_set_group.rr_set.loc_record.values — values / 232122301303 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010)
- primary.rr_set_group.rr_set.loc_record.values

<a id="canonical-1211112130033313-0001333210220013-1012303312003132-1302203221002201-1011022323200210-2103131232211132-1112322032113322-2101333131202112"></a>

Type: `"list"`. Computed.

LOC Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3111212122100220-1031111201201300-0131123022000332-2023023301220111-3010113111001312-1003110321110002-1020121013203121-0112103033211310"></a>

## Direct properties — values / 232122301303 / 3

<a id="canonical-2222132112000030-0310330123330132-0321123120201130-0111123130111021-2322113203022132-3221101203323322-3122223211202021-1220121231133330"></a>

<a id="canonical-1211012113031210-3203212331221313-0200011321133301-0013011120231301-3113003130023032-2211323102313200-1123310030022021-3101010302110133"></a>

## altitude property — values / 232122301303 / 4

Type: `"number"`. Computed.

Altitude. Altitude in meters.

Upstream description:

Altitude in meters.

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
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "-100000.00",
    "ves.io.schema.rules.float.lte": "42849672.95",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3022330112300003-2101102102123232-2033210031113203-3002322120120220-2213030031112110-2311010213002330-3012311300031232-3021121331102103"></a>

<a id="canonical-1220003102021113-3030212231120232-2230310100322001-3220323201211200-2211321031230002-0113333020031002-0313103100313120-1202323033302010"></a>

## horizontal_precision property — values / 232122301303 / 5

Type: `"number"`. Computed.

Horizontal Precision. Horizontal Precision in meters.

Upstream description:

Horizontal Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-3110222123202312-3023302020220102-3303211313030132-2333010102201203-3221232320101031-0213101110130202-2121323110220213-1300013223030123"></a>

<a id="canonical-2303122232303031-3122221323031103-0133110123033101-0020000013221310-1020010131023123-1123000323000223-2322200000111333-1031220031212321"></a>

## latitude_degree property — values / 232122301303 / 6

Type: `"number"`. Computed.

Latitude degree, an integer between 0 and 90, including 0 and 90.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 90,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "90",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0111330302313111-2332211212300211-3101302322000200-1233233202312011-0210232101131002-1303121021013232-3223312011021132-3201300323021120"></a>

<a id="canonical-1002001120133010-3033321303311222-2012123202202123-3201200333303210-0203322130212031-1021231223003330-0121220023013110-3232020131122120"></a>

## latitude_hemisphere property — values / 232122301303 / 7

Type: `"string"`. Computed.

\[Enum: N|S\] Latitude hemisphere can only be N or S - N: North Hemisphere - S: South Hemisphere.
Possible values are \`N\`, \`S\`. Defaults to \`N\`.

Upstream description:

Latitude hemisphere can only be N or S

&#8203;- N: North Hemisphere

&#8203;- S: South Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "N",
  "enum": [
    "N",
    "S"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3200123113211231-3332120110000101-0110320132310200-0001222030032322-0310122331020230-2311203101032201-2301033033103122-2321133131203232"></a>

<a id="canonical-3223330333210201-2133023013211011-3010131001233332-0131121232132102-0320112322232232-2020311201011310-1222202121232100-0132222301200012"></a>

## latitude_minute property — values / 232122301303 / 8

Type: `"number"`. Computed.

Latitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-2210102333110102-0031330021011331-3020033131333202-1033111233002220-2110211233321231-0032032001331021-3103002220111222-0202321222020020"></a>

<a id="canonical-1021231002330213-0023130330220223-0300130103122013-3313321333321311-3330113101130232-3310022103323123-3120310312202301-1332323031103222"></a>

## latitude_second property — values / 232122301303 / 9

Type: `"number"`. Computed.

Latitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-1212030312111130-1312210022310221-1010201220330032-0220302120333312-2222332330230100-0032301211312102-1011212222221011-0112212201020311"></a>

<a id="canonical-2330122110211022-1212332002303110-1313330322200121-1101331000321130-3131330012312230-3013221101122233-2101112123223111-2113310133223330"></a>

## location_diameter property — values / 232122301303 / 10

Type: `"number"`. Computed.

Diameter of a sphere enclosing the described entity, in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-2111001200013213-0110120330110322-0331221330203303-3102132333302121-3110020033013120-2310032301031003-2312013123323222-0010323001321223"></a>

<a id="canonical-0022323111011031-1330001321100231-0112233002103102-2321320013321212-1101100221031130-0211320103023333-3331331112312310-2301310301330233"></a>

## longitude_degree property — values / 232122301303 / 11

Type: `"number"`. Computed.

Longitude degree, an integer between 0 and 180, including 0 and 180.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 180,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "180",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-3111201232032021-1220021122201120-2033010222130000-0022210021203103-3101002332121313-1311331012330212-3000003302112111-2123010131000321"></a>

<a id="canonical-3131023102111323-3220020220332231-0021132122232011-1032103320232221-0320113310220320-1001211113131101-3133023033223320-3112003020203313"></a>

## longitude_hemisphere property — values / 232122301303 / 12

Type: `"string"`. Computed.

\[Enum: E|W\] Longitude hemisphere can only be E or W - E: East Hemisphere - W: West Hemisphere.
Possible values are \`E\`, \`W\`. Defaults to \`E\`.

Upstream description:

Longitude hemisphere can only be E or W

&#8203;- E: East Hemisphere

&#8203;- W: West Hemisphere.

Receipt-pinned upstream constraints:

```json
{
  "default": "E",
  "enum": [
    "E",
    "W"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1030211331000333-2301033202230211-1103101233121032-1032131212102230-1332223222132232-0312010111001312-3211332012212022-3020023001300003"></a>

<a id="canonical-0212032203122131-3110202033122133-3203303303110103-1130022023211303-1030031203033132-0111000023023202-3020000021303233-2231331220132232"></a>

## longitude_minute property — values / 232122301303 / 13

Type: `"number"`. Computed.

Longitude minute, an integer between 0 and 59, including 0 and 59.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 59,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.int32.gte": "0",
    "ves.io.schema.rules.int32.lte": "59"
  }
}
```

<a id="canonical-2331223312002332-2222001221322123-0102102132203110-3321230121322300-1032102220010030-0211012221133102-1011212332200333-1031221332101102"></a>

<a id="canonical-1000233013220223-2332213322313222-3323020233131200-3100032302302230-3103133012313300-3322201311131220-3101100213123320-3221222013100012"></a>

## longitude_second property — values / 232122301303 / 14

Type: `"number"`. Computed.

Longitude second, an decimal between 0 and 59.999, including 0 and 59.999.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "59.999"
  }
}
```

<a id="canonical-2200331331113020-0321300301110030-1002113022231312-0123033233113033-2101130132120012-3012022111201103-2230222003012022-0332210312312211"></a>

<a id="canonical-2000313113003111-1100132311330230-3211311012312001-2233230312230121-0303223013203102-1100033033202301-0200301001201000-3323230203203323"></a>

## vertical_precision property — values / 232122301303 / 15

Type: `"number"`. Computed.

Vertical Precision. Vertical Precision in meters.

Upstream description:

Vertical Precision in meters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.float.gte": "0.0",
    "ves.io.schema.rules.float.lte": "90000000.00"
  }
}
```

<a id="canonical-2221303322200130-1322023212003022-0303332022331020-3122211331111010-3113000221313212-0033200212333112-1001012113221202-3113232031023130"></a>

## Next pages — values / 232122301303 / 16

- [primary.rr_set_group.rr_set.loc_record](data-sources--dns_zone--reference--group-003.md#canonical-3103002133110011-3300030122231322-1033330212002032-1120030230101113-1221213310123022-1300023322200233-0112012020111212-1320210023313010)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330323212233023-0020010233021011-1100131101003002-2220012233233202-3330101123332113-3201130320331130-3033003201330003-3232330130210310"></a>

## primary.rr_set_group.rr_set.mx_record — mx_record / 201032313300 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.mx_record

<a id="canonical-1032301230121302-2230303233101300-3200131020003011-3223231131120111-3211000013130303-3113013032123030-2313121333332032-0030212331102010"></a>

Type: `"single"`. Computed.

DNSMXResourceRecord.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3332120200002001-1133221020133300-1300123102031230-3001001322312320-1002032033203111-2031030103001203-0301320002231230-1311001013131000"></a>

## Direct properties — mx_record / 201032313300 / 3

<a id="canonical-3311033132001020-0102022220030111-1232031110121310-2231022232031233-1133110222101113-1132313012310030-2313023133102322-1333311331220330"></a>

<a id="canonical-1132223313022012-2233112021132031-3123212231130302-0102110031221222-0210331233200010-2113220121133212-3230222001011103-3110202001200323"></a>

## name property — mx_record / 201032313300 / 4

Type: `"string"`. Computed.

MX Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-3312021201221301-3030330303323212-1020333131110122-1133200100220220-3301010123311000-1023002001023200-3221330220300001-3323322023301322): complete subsection reference.

<a id="canonical-2202033222221003-2232021312233300-3031133323310133-0013300021212221-1031132011211321-2332321301031013-3213301303101120-3210032133031003"></a>

## Next pages — mx_record / 201032313300 / 5

- [primary.rr_set_group.rr_set.mx_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3312021201221301-3030330303323212-1020333131110122-1133200100220220-3301010123311000-1023002001023200-3221330220300001-3323322023301322)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3312021201221301-3030330303323212-1020333131110122-1133200100220220-3301010123311000-1023002001023200-3221330220300001-3323322023301322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332100132001211-2130302333133300-2120132003332010-0123310010102013-2133303301313022-2102212010032013-2331222201020132-1021013120110013"></a>

## primary.rr_set_group.rr_set.mx_record.values — values / 111121301203 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022)
- primary.rr_set_group.rr_set.mx_record.values

<a id="canonical-2010113303312310-1321103120212023-2203002233330103-3133021300012001-0313201120310203-2332220233012323-3310013003130201-2321220322121321"></a>

Type: `"list"`. Computed.

MX Record Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100"
  }
}
```

<a id="canonical-0023312212102102-2312103132332311-0233211303230103-0221221221133311-2023110220311200-3102202100030332-1103132332222131-2300020100320022"></a>

## Direct properties — values / 111121301203 / 3

<a id="canonical-2212112331311011-1102003102311101-3303301010021132-1132200332102233-1303330021322112-2102312121323021-3020103332131030-3211011300301002"></a>

<a id="canonical-3111023131333100-3101032310021012-0132013110121022-0202110223212101-0231022030110233-2010021202030223-2110201303020132-1221122232212132"></a>

## domain property — values / 111121301203 / 4

Type: `"string"`. Computed.

Mail exchanger domain name, please provide the full hostname, for.

Upstream description:

Mail exchanger domain name, please provide the full hostname, for example: mail.example.com.

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="canonical-1102132003010311-2221313113001020-3201012012000233-0212222330213313-0203330131320113-1211222100232300-3202231313221013-3313300300310020"></a>

<a id="canonical-1121200121212202-2212331210031030-1323321301130102-1012301120221013-0333020323132232-2021113131221113-0231032102101310-2132233131120333"></a>

## priority property — values / 111121301203 / 5

Type: `"number"`. Computed.

Priority. Mail exchanger priority code.

Upstream description:

Mail exchanger priority code.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3030223302000132-2213021313222200-1331331232300022-0031231311202012-3220221102222031-0012012202121023-1132302113230100-0012103222101130"></a>

## Next pages — values / 111121301203 / 6

- [primary.rr_set_group.rr_set.mx_record](data-sources--dns_zone--reference--group-003.md#canonical-2232020013010222-3011001301213203-0001030100231211-0233002301012011-0102021001102022-2003330212132112-3220310312321133-1330303033320022)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220032000133032-2312300312030100-2020012011232112-0123230202321020-1302030333022113-2112333301010111-0300322023230302-0123201003012311"></a>

## primary.rr_set_group.rr_set.naptr_record — naptr_record / 032211112100 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.naptr_record

<a id="canonical-0133232032210101-0001200232000323-1130202303131203-3331333233303332-0012131002202331-0320323131233332-0000212013233331-1011012221020233"></a>

Type: `"single"`. Computed.

Configuration parameter for naptr record.

Upstream description:

DNS NAPTR Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1121111212130313-0232031223032231-0312111001212010-1211201312033230-0213221322031123-0312032333330333-0301113300203211-0323031000032133"></a>

## Direct properties — naptr_record / 032211112100 / 3

<a id="canonical-1212011312322322-0312203331301013-3322121111210101-2200211220000111-3032112033200333-0300322110002031-3202232010012220-1112222133303213"></a>

<a id="canonical-1112133321100102-0030101102311303-2200010231333333-3223320023123220-3232322002313211-1022203011011022-3200322211320331-2101201200123101"></a>

## name property — naptr_record / 032211112100 / 4

Type: `"string"`. Computed.

NAPTR Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-1023231202220220-2030321031113022-3303221030012233-3132133203303320-1021013000001222-3211013210110322-2121322002201003-3010321130013113): complete subsection reference.

<a id="canonical-3331131121101202-0101023203303323-0012231333113010-0321322022033120-3001200113322020-2312131101020201-0221111020231111-0302233332131222"></a>

## Next pages — naptr_record / 032211112100 / 5

- [primary.rr_set_group.rr_set.naptr_record.values](data-sources--dns_zone--reference--group-003.md#canonical-1023231202220220-2030321031113022-3303221030012233-3132133203303320-1021013000001222-3211013210110322-2121322002201003-3010321130013113)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1023231202220220-2030321031113022-3303221030012233-3132133203303320-1021013000001222-3211013210110322-2121322002201003-3010321130013113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031333102012010-3321210310013333-3110203222322203-1100321221321302-1131020103003231-2300211110220303-2321232130220112-1203131113330331"></a>

## primary.rr_set_group.rr_set.naptr_record.values — values / 200322301222 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010)
- primary.rr_set_group.rr_set.naptr_record.values

<a id="canonical-2200231101020302-0013202003222102-3000130322230201-1032213200122031-0022233212211123-0021232220010203-3020201223101211-0110020030232033"></a>

Type: `"list"`. Computed.

NAPTR Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2121100211210300-3102022302031020-0301130030333011-2321012301332102-3131321100321032-2231103133131222-2322223121231103-2130313102130311"></a>

## Direct properties — values / 200322301222 / 3

<a id="canonical-1031012120012113-1313312322333103-2032103222123012-2121333312111012-3031130302031121-2202230212033331-0203301113002231-2312323013113013"></a>

<a id="canonical-1101202023201303-3012320000110322-0231030111332020-1222322221313303-3203132221220330-1022330321320010-1311002321023020-0000113013023033"></a>

## flags property — values / 200322301222 / 4

Type: `"string"`. Computed.

Flag to control aspects of the rewriting and interpretation of the fields in the record. At this
time only four flags, S/A/U/P, are defined.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^(S|s|A|a|U|u|P|p)$"
  }
}
```

<a id="canonical-2311202231301323-3121333022301101-2223030001121022-0022120312113201-0303102103310032-2220312031222320-0312013323332111-2000122301101021"></a>

<a id="canonical-2320031020301320-2310002100233032-0202003110012023-3001323030330220-1012302322101302-2310120320212110-0210101333021201-1122212303000122"></a>

## order property — values / 200322301222 / 5

Type: `"number"`. Computed.

Order in which the NAPTR records must be processed. A lower number indicates a higher preference.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0032220230321232-0210110100300333-0111130031311110-3013012130120210-0211233310211013-1131033010331121-2223031133003030-2013111232122032"></a>

<a id="canonical-1111202123332021-0112013030123220-1303131113300321-1320312012220301-1110133033322000-3212033200202330-3303101220001132-1310202222131000"></a>

## preference property — values / 200322301222 / 6

Type: `"number"`. Computed.

Preference when records have the same order. A lower number indicates a higher preference.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-0101230230333232-1332200302210321-3130112222202020-0102100120101330-0121133023313321-3112011001110013-2323021301022023-0003332310132202"></a>

<a id="canonical-2322300330031211-2021031130032233-1213012231002303-2210330132313322-1023313210111332-2123032221010332-1222023300221302-0321121220101001"></a>

## regexp property — values / 200322301222 / 7

Type: `"string"`. Computed.

Regular expression to construct the next domain name to lookup.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
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
    "ves.io.schema.rules.string.max_len": "255"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255"
  }
}
```

<a id="canonical-2200200231122010-1201013102211212-2301011122201231-2223033312113200-2311302021012310-3030120312101310-0121001121111000-0222110200010111"></a>

<a id="canonical-1222301333222130-3230313130223002-1303201201233331-1122010233102230-0101202211200331-2320011233021220-0311300223120332-3311030031332011"></a>

## replacement property — values / 200322301222 / 8

Type: `"string"`. Computed.

The next NAME to query for NAPTR, SRV, or address records depending on the value of the flags field.

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

<a id="canonical-0222322101131033-2330013302030033-2322120201123121-2313311332010213-2210100232003131-3301000021132131-1001003200202201-0313032032003033"></a>

<a id="canonical-3102322031310232-1031312021001332-3111132122312101-1303310201113310-0111001201230011-0011022203112320-1130012321021121-0230201333211232"></a>

## service property — values / 200322301222 / 9

Type: `"string"`. Computed.

Specifies the service(s) available down this rewrite path.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 255,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 255,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "255",
    "ves.io.schema.rules.string.pattern": "^([A-Za-z][A-Za-z0-9]{0,31}(\\\\+[A-Za-z][A-Za-z0-9]{0,31})*$|^$)"
  }
}
```

<a id="canonical-3301202213033020-1223013103230333-1302230103111030-0101322012211131-1210233120303202-1100031200323133-3232132000200221-2212323030033233"></a>

## Next pages — values / 200322301222 / 10

- [primary.rr_set_group.rr_set.naptr_record](data-sources--dns_zone--reference--group-003.md#canonical-2130200302222231-0201323023000030-2231013210302021-3320130221100012-2012002232303013-1033323023120320-3201231132301223-2120220121110010)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1301020100313003-2113130112303333-0222200332023331-1011211333020200-3333320312030230-2202333330201232-2032210300303202-3111001112202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221333330213002-2312322233200133-1003313023122330-2303323032333130-3133112211110003-3103122121213222-2322021310223002-0122332223330130"></a>

## primary.rr_set_group.rr_set.ns_record — ns_record / 111331131331 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.ns_record

<a id="canonical-0330312201100120-2112312301203220-3230003000130221-0023232000030332-3220220102110023-1132123131100021-3103230321221112-3322223320131103"></a>

Type: `"single"`. Computed.

DNSNSResourceRecord.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113130113223110-3323033030220131-1201202232200301-1200103322233102-1032031102023133-1231323002012021-3021203311231033-1313303002011023"></a>

## Direct properties — ns_record / 111331131331 / 3

<a id="canonical-0200211233030201-3122121200101333-3220202313311001-1213312221301101-2333121220001202-3100112100122322-0301211011133220-3031313202011301"></a>

<a id="canonical-3332313131222230-2202131231211012-2210003131132003-3110133031113210-3210231220013010-1332320320321333-0321303122322110-0310303113131021"></a>

## name property — ns_record / 111331131331 / 4

Type: `"string"`. Computed.

NS Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2121333000030022-1111112311202130-1321222312321100-0231012012113003-0022310110111233-3320111233223231-1230011103023030-3130102330101301"></a>

<a id="canonical-1313230210001320-0223020112100123-1221330032020113-2012003300123320-0323120023033002-0033321003103220-1020131222100030-2120123302300033"></a>

## values property — ns_record / 111331131331 / 5

Type: `["list", "string"]`. Computed.

Name Servers. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0122222313110310-2312312000100231-0000110220332120-3211021022211311-3002213032013033-1010301000122011-3130323320333220-1223110011001122"></a>

## Next pages — ns_record / 111331131331 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2211303020222311-3310323332010203-0232301233301110-3121123122223302-3103033101210322-2020123011002111-3032313210101120-2001120323033330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303322130222013-0123333021021123-2020321331201002-2302303310302123-0233332113101022-1103111132020313-1101221203320131-3331211121322000"></a>

## primary.rr_set_group.rr_set.ptr_record — ptr_record / 023032011110 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.ptr_record

<a id="canonical-3030302130012103-1031200320000313-3213320132112032-3011133211333100-2211102023313110-0000310203120132-2233122200311201-0012110010030112"></a>

Type: `"single"`. Computed.

DNSPTRResourceRecord.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3321211200233022-2111101000233023-1010100221302102-2031011300311122-3022203002300020-1201311330213102-1203201213211313-3030100301312202"></a>

## Direct properties — ptr_record / 023032011110 / 3

<a id="canonical-2320112321012331-2211032133102013-2012112213232302-2200312233332133-2030122012311323-3000122321100123-2301210012312211-0232130132131110"></a>

<a id="canonical-1231321031011000-2200232022303113-0020323103030210-0020033213313032-2121232230332300-0112023220213202-0233000110030300-1301121223332000"></a>

## name property — ptr_record / 023032011110 / 4

Type: `"string"`. Computed.

PTR Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-3022000221212222-1223011132033230-1001032003202132-3013221033030122-0030010331211121-3002101102030231-2032222322202102-0122110132221302"></a>

<a id="canonical-3112232013031213-2320003332322123-1202001330332333-3232113333101202-0303303033011221-3121200203130030-1132000210311013-0010020020203223"></a>

## values property — ptr_record / 023032011110 / 5

Type: `["list", "string"]`. Computed.

Domain Name. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.hostname": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0233002003202300-1021022111121213-1031122003011111-3003331223223101-2230102211301300-0311202031002000-3220332121123011-0323002031132020"></a>

## Next pages — ptr_record / 023032011110 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221020320210113-3111111212221203-3312213101001001-1203330120130132-1023301320301213-1223113203331001-3122231020223011-2121313320100111"></a>

## primary.rr_set_group.rr_set.srv_record — srv_record / 122311100221 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.srv_record

<a id="canonical-1000221220030112-2010100232132001-3022002120021130-1103323000130332-1102212231213330-0122231210232320-2020130132322232-3121222113332300"></a>

Type: `"single"`. Computed.

DNSSRVResourceRecord.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0120231000001300-0230221133212022-1030323221220021-2032332021332120-0233111013202123-2303110320120012-2023102213022020-0310323102033112"></a>

## Direct properties — srv_record / 122311100221 / 3

<a id="canonical-2302322211300310-3222300011022310-0210330303222112-1321232112133201-1223132210322111-1213310002320301-2223302100310113-0202212113031003"></a>

<a id="canonical-3030110203123030-2203000033321122-0310230311020121-0310311003121223-1322002132221120-2300231200000112-3230233313021303-3110012230201311"></a>

## name property — srv_record / 122311100221 / 4

Type: `"string"`. Computed.

SRV Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^([*]|[a-zA-Z0-9-_]{1,63})([.][a-zA-Z0-9-_]{1,63})*$"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-0211103321112302-0300323031310021-2302311310202221-3120013221210130-1001211300111011-2322132002232311-3113032000333100-1312130101032030): complete subsection reference.

<a id="canonical-3101330323332223-2131122332132133-1111223221200123-0313232300210333-0122101132310301-1030013103010221-2113220111323000-1220322102130122"></a>

## Next pages — srv_record / 122311100221 / 5

- [primary.rr_set_group.rr_set.srv_record.values](data-sources--dns_zone--reference--group-003.md#canonical-0211103321112302-0300323031310021-2302311310202221-3120013221210130-1001211300111011-2322132002232311-3113032000333100-1312130101032030)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0211103321112302-0300323031310021-2302311310202221-3120013221210130-1001211300111011-2322132002232311-3113032000333100-1312130101032030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110131200123110-0030000232120230-2300100013112130-0310110223121310-1301120231312003-1121012113311233-1112110112002232-3302122100233230"></a>

## primary.rr_set_group.rr_set.srv_record.values — values / 232101301133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312)
- primary.rr_set_group.rr_set.srv_record.values

<a id="canonical-2212233103113113-0311022000300322-0310231201320310-1100021111230311-2303130003123100-2223332202103301-0111000113330012-1221313232200032"></a>

Type: `"list"`. Computed.

SRV Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3323213110310313-0202013222021022-3302200321313210-2220220301313103-2021103120013021-2203231310310111-2210223102301323-0103202133102122"></a>

## Direct properties — values / 232101301133 / 3

<a id="canonical-3003201220200130-1021330232001031-2321011021133032-0002310022312232-2102213113331110-0132012022323033-1223031023100212-0321202210233230"></a>

<a id="canonical-2012310322211332-3103303300132123-3111211232101310-3231320230120112-2203313020310323-0021213302023112-0111233233320020-0222031001301110"></a>

## port property — values / 232101301133 / 4

Type: `"number"`. Computed.

Port. Port on which the service can be found.

Upstream description:

Port on which the service can be found.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3210220321011103-0221123113220123-3101321201123011-3132321110103301-2232212102010012-3320012113021320-3201132330203101-1003103012001021"></a>

<a id="canonical-1203101200323110-2202210220333221-0120102111011303-1202103212310022-2100223001111112-1323233032200110-3312122023223033-1112203223133123"></a>

## priority property — values / 232101301133 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2222300302002330-0331121232303121-1323233320131120-2101102203212132-2213100223130100-1131220022202331-2100132311031101-0210020030000301"></a>

<a id="canonical-0233230330131120-3102010101101102-1032303213112113-3201203211123322-2023321312333320-2121111003333223-1313312220333222-2303023112003031"></a>

## target property — values / 232101301133 / 6

Type: `"string"`. Computed.

Hostname of the machine providing the service.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^[.]$|^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\\\\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\\\\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\\\\.?$"
  }
}
```

<a id="canonical-0313220012311033-3131021312031102-2010003303132000-2313020302003312-3301221231001320-0300103203001322-1232303020123102-3030021033312210"></a>

<a id="canonical-0200020301211000-0332220131331202-0312200332212222-3212022222100120-0330231302221233-2322223010312332-0231203101320221-1121201313011000"></a>

## weight property — values / 232101301133 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0301313000003032-0222312010303010-3301012122130233-3120311300003032-2231022201331020-2031133122321022-3310022020233303-2201331101030323"></a>

## Next pages — values / 232101301133 / 8

- [primary.rr_set_group.rr_set.srv_record](data-sources--dns_zone--reference--group-003.md#canonical-0223103023112010-0122311332230213-0020010203033121-2323203031213310-1321013110303211-1310102220022031-3130133320120313-1333011010122312)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022300333003101-3123022303110313-3120103131232130-3131022333100222-0021313311202132-0131321100110312-3231203111233231-1012102233030221"></a>

## primary.rr_set_group.rr_set.sshfp_record — sshfp_record / 221132031200 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.sshfp_record

<a id="canonical-2121012320102322-0110331111102013-2132110133203320-3100031022001321-0032102321220212-3221111323222212-3100312000000230-0120322113132122"></a>

Type: `"single"`. Computed.

Configuration parameter for sshfp record.

Upstream description:

DNS SSHFP Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0200211113221220-0312320122333312-3102232133212332-1023233313213133-2300330011020212-2000033003231100-1313222012333332-0311233203030332"></a>

## Direct properties — sshfp_record / 221132031200 / 3

<a id="canonical-1231031000113201-3321323320010331-3323322201120313-0222231133212021-1013111310313302-1232032001323032-1233021231002103-3113110111201021"></a>

<a id="canonical-2030233210132302-1003302103302233-2333333331130013-3210123210222123-0131103332123112-3303310012022102-3312201100223302-3122231102030011"></a>

## name property — sshfp_record / 221132031200 / 4

Type: `"string"`. Computed.

SSHFP Record name, please provide only the specific subdomain or record name without the base
domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110): complete subsection reference.

<a id="canonical-0332133033002322-0201323033221313-2311320212313021-3211013002132030-0133122131121023-3301233312303032-0222000332223231-0330033031033313"></a>

## Next pages — sshfp_record / 221132031200 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012021110032011-2120321301320300-1000113333030101-2302232213320312-2103003332301213-3121223302320310-1321012011130022-2202311301301232"></a>

## primary.rr_set_group.rr_set.sshfp_record.values — values / 000202222123 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- primary.rr_set_group.rr_set.sshfp_record.values

<a id="canonical-2331300330310133-0031203323201200-0020002201201200-0030033113010311-3330221213323321-1321210003322111-3011230300321203-0332322021001321"></a>

Type: `"list"`. Computed.

SSHFP Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-3011322301113232-2032012211013313-3203223100223131-2011103230101003-2300103210033113-1032103113023302-2122122203203013-2120300300232103"></a>

## Direct properties — values / 000202222123 / 3

<a id="canonical-1300023031211311-3313232011221303-0122212033232122-0212102021130121-1002320332020233-2310132330313320-1231033230020221-1213001232222300"></a>

<a id="canonical-0323312100022310-2023010112010120-0103211230311300-0110230213323032-2023223323030011-3233113300031021-3223202220312023-2202112231203300"></a>

## algorithm property — values / 000202222123 / 4

Type: `"string"`. Computed.

\[Enum: UNSPECIFIEDALGORITHM|RSA|DSA|ECDSA|Ed25519|Ed448\] SSHFP algorithm value must be compatible
with the specified algorithm. - UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM - RSA: RSA - DSA: DSA -
ECDSA: ECDSA - Ed25519: Ed25519 - Ed448: Ed448. Possible values are \`UNSPECIFIEDALGORITHM\`,
\`RSA\`, \`DSA\`, \`ECDSA\`, \`Ed25519\`, \`Ed448\`. Defaults to \`UNSPECIFIEDALGORITHM\`.

Upstream description:

SSHFP algorithm value must be compatible with the specified algorithm.

&#8203;- UNSPECIFIEDALGORITHM: UNSPECIFIEDALGORITHM

&#8203;- RSA: RSA

&#8203;- DSA: DSA

&#8203;- ECDSA: ECDSA

&#8203;- Ed25519: Ed25519

&#8203;- Ed448: Ed448.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNSPECIFIEDALGORITHM",
  "enum": [
    "UNSPECIFIEDALGORITHM",
    "RSA",
    "DSA",
    "ECDSA",
    "Ed25519",
    "Ed448"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-1212033231013101-1211210110120033-2201033211111131-0322310331021203-2200012100333231-0333231210321010-3113021122010322-1233011230030312): complete subsection reference.

- [sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-0321003213330333-0032202123013222-1320212310111100-0222320102323300-1220220323000302-1020321111210022-3020001101001020-3133201013032301): complete subsection reference.

<a id="canonical-3022321211113210-1122302230220000-3202131332302213-2213101203230122-3323211103302201-3100121332223310-2120101121202322-0112020023201303"></a>

## Next pages — values / 000202222123 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-1212033231013101-1211210110120033-2201033211111131-0322310331021203-2200012100333231-0333231210321010-3113021122010322-1233011230030312)
- [primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint](data-sources--dns_zone--reference--group-003.md#canonical-0321003213330333-0032202123013222-1320212310111100-0222320102323300-1220220323000302-1020321111210022-3020001101001020-3133201013032301)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1212033231013101-1211210110120033-2201033211111131-0322310331021203-2200012100333231-0333231210321010-3113021122010322-1233011230030312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120020033303301-2121222331032011-2310132003021012-3003323130230022-3121302303222233-1013133220113302-3222110010033010-3230310132120213"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint — sha1_fingerprint / 202000202330 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- primary.rr_set_group.rr_set.sshfp_record.values.sha1_fingerprint

<a id="canonical-1310013002022232-3123100310320020-3302132323103132-3010312231002020-0322211133101202-0330130210120012-0023132001322001-0221131011332221"></a>

Type: `"single"`. Computed.

Configuration parameter for sha1 fingerprint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0022330203311300-0000203301022000-1022103210122122-0303223113223032-0222032113333333-2322311332223201-3001203201230031-3102131023133131"></a>

## Direct properties — sha1_fingerprint / 202000202330 / 3

<a id="canonical-0232000100001222-2300203103302233-3121333113330031-3221212301303111-2210323011131110-1130103021323111-0102312333130231-1103232131000330"></a>

<a id="canonical-2100333120111100-2012233310103031-2301212302313103-2233123133323001-0021102031232322-3320330130022202-2320030332121010-0320300002213001"></a>

## fingerprint property — sha1_fingerprint / 202000202330 / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 40,
  "minLength": 40,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 40,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 40,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "40",
    "ves.io.schema.rules.string.min_len": "40"
  }
}
```

<a id="canonical-0230122023121233-2112221132133210-0332201211030011-3230210231212211-1030032110002031-1121023103120302-2212222232131021-0021312120031203"></a>

## Next pages — sha1_fingerprint / 202000202330 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0321003213330333-0032202123013222-1320212310111100-0222320102323300-1220220323000302-1020321111210022-3020001101001020-3133201013032301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203321130112220-0010121310031111-3033232202321133-1102121332021222-2333033220321011-2321101130233212-1322012101121121-0021213313210232"></a>

## primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint — sha256_fingerprint / 301121100132 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.sshfp_record](data-sources--dns_zone--reference--group-003.md#canonical-1320201233022013-2010302020223210-3223212010011202-0302303001331232-1022311200333221-3030011233031331-2201220131002312-2233122232100332)
- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- primary.rr_set_group.rr_set.sshfp_record.values.sha256_fingerprint

<a id="canonical-2021231231131011-2002010300220312-2013311332002322-2323103231333301-3032102030233320-0321333100133023-2223003113133021-3222311112110112"></a>

Type: `"single"`. Computed.

Configuration parameter for sha256 fingerprint.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0211310100002322-1121112033312210-2021122112222120-0122301112011300-3111001323000311-0023200020030000-1023132310030321-3012213300122022"></a>

## Direct properties — sha256_fingerprint / 301121100132 / 3

<a id="canonical-3201203001133212-0131133232032213-2011311202003222-2021102032301211-1201033223000230-2302323122302213-2013030113111010-0021122211020113"></a>

<a id="canonical-0223313213210032-2023313011112213-2331013113003133-3110303211220030-1221213332210013-0331333313003330-3130310001002322-1021033232201311"></a>

## fingerprint property — sha256_fingerprint / 301121100132 / 4

Type: `"string"`. Computed.

The 'fingerprint' is the DS key and the actual contents of the DS record.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[0-9a-fA-F]",
      "description": "Hexadecimal characters only"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 64,
    "pattern": "^[0-9a-fA-F]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "64"
  }
}
```

<a id="canonical-0322220311032121-0031012223313123-2233331323300110-1022132312211330-2332302001103123-0330330320201001-0031123203000201-3212220011000022"></a>

## Next pages — sha256_fingerprint / 301121100132 / 5

- [primary.rr_set_group.rr_set.sshfp_record.values](data-sources--dns_zone--reference--group-003.md#canonical-3201121033301222-3001030332303100-3001330110033123-2300301233132233-2230221111201220-3211111320003312-0210112201112113-1100200100032110)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303003102112123-2332323333222123-3030023233200231-0223112213210010-3022103311103221-3222312102031313-2010300101033320-3030212101102023"></a>

## primary.rr_set_group.rr_set.tlsa_record — tlsa_record / 333223331110 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.tlsa_record

<a id="canonical-1101223212221102-3322031333001130-1103322111122133-1033323303213231-3332201223232100-1321231100221231-2130122203112013-0222012110223001"></a>

Type: `"single"`. Computed.

Configuration parameter for tlsa record.

Upstream description:

DNS TLSA Record.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0012001332101321-3301123010102202-2203233030132332-2010113110010220-1002021111121303-0003321003233303-1112230333322001-1112203323221200"></a>

## Direct properties — tlsa_record / 333223331110 / 3

<a id="canonical-3013011123320300-1330002211301302-0010222000332022-0331110332323222-3013010011022032-0310312331031133-3232332231113010-0011100130233301"></a>

<a id="canonical-3330322131213223-3111330301120323-0123100003322112-0021230102300321-1302221020122332-2230033311020032-1302112321011330-0122033222103232"></a>

## name property — tlsa_record / 333223331110 / 4

Type: `"string"`. Computed.

TLSA Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}",
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
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^([a-zA-Z0-9*?]|([a-zA-Z0-9?*]+-[a-zA-Z0-9*?]+)){0,253}"
  }
}
```

- [values](data-sources--dns_zone--reference--group-003.md#canonical-2300322000010100-3330131011030130-0002231310231000-0231233002101022-1131210322121011-2330020130222313-3123203000313102-2213110333130232): complete subsection reference.

<a id="canonical-0132332321120301-1230231330322012-2202312320232011-1223020133332233-0032331302103012-0023022121022113-2032011021032111-0000002033201111"></a>

## Next pages — tlsa_record / 333223331110 / 5

- [primary.rr_set_group.rr_set.tlsa_record.values](data-sources--dns_zone--reference--group-003.md#canonical-2300322000010100-3330131011030130-0002231310231000-0231233002101022-1131210322121011-2330020130222313-3123203000313102-2213110333130232)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2300322000010100-3330131011030130-0002231310231000-0231233002101022-1131210322121011-2330020130222313-3123203000313102-2213110333130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232100322033103-0301321000102231-2230312332221101-3120133331222020-2223310122011200-0301111320131331-0221302033332223-3021212102320211"></a>

## primary.rr_set_group.rr_set.tlsa_record.values — values / 222123010030 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121)
- primary.rr_set_group.rr_set.tlsa_record.values

<a id="canonical-2312332012210332-2122213103112133-1132301230330200-0223000200211331-1121010310320232-3020220333130203-3313313323312102-2330321021132333"></a>

Type: `"list"`. Computed.

TLSA Value. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-2031012221303131-3332323320112021-1203331231220032-3223222203332033-0022130020222223-0012021213131222-1100033133100203-0110210010323132"></a>

## Direct properties — values / 222123010030 / 3

<a id="canonical-2022331111110123-2100103133330003-2011211220002120-2200222331011202-0233133231121211-2100130131020002-0020133032101201-1120100312130101"></a>

<a id="canonical-3000003322113333-0020313230002223-0220011110113103-2131231113012212-3133302321010321-3202110012003201-2121033100200333-3231032002121032"></a>

## certificate_association_data property — values / 222123010030 / 4

Type: `"string"`. Computed.

The actual data to be matched given the settings of the other fields.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hex",
    "maxLength": 4096,
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
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hex": "true",
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-1032033121302201-0313000230112110-1201323010101122-2203310201303130-2302113130330111-3201013302201122-1023203000012300-0202130302300122"></a>

<a id="canonical-3112210202330221-3023222021312330-0022032102002303-3313302323202122-1302232030010022-3321211130301103-1030132201121300-1313300100033311"></a>

## certificate_usage property — values / 222123010030 / 5

Type: `"string"`. Computed.

\[Enum:
CertificateAuthorityConstraint|ServiceCertificateConstraint|TrustAnchorAssertion|DomainIssuedCertificate\]
&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint - ServiceCertificateConstraint:
Service Certificate Constraint - TrustAnchorAssertion: Trust Anchor Assertion -
DomainIssuedCertificate: Domain Issued Certificate. Possible values are
\`CertificateAuthorityConstraint\`, \`ServiceCertificateConstraint\`, \`TrustAnchorAssertion\`,
\`DomainIssuedCertificate\`. Defaults to \`CertificateAuthorityConstraint\`.

Upstream description:

&#8203;- CertificateAuthorityConstraint: Certificate Authority Constraint

&#8203;- ServiceCertificateConstraint: Service Certificate Constraint

&#8203;- TrustAnchorAssertion: Trust Anchor Assertion

&#8203;- DomainIssuedCertificate: Domain Issued Certificate.

Receipt-pinned upstream constraints:

```json
{
  "default": "CertificateAuthorityConstraint",
  "enum": [
    "CertificateAuthorityConstraint",
    "ServiceCertificateConstraint",
    "TrustAnchorAssertion",
    "DomainIssuedCertificate"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2231300020001110-2202032131312021-3013103200231130-1122101031311230-3002123203312100-1023230201233010-3121330111100032-2221122330021322"></a>

<a id="canonical-3010213023220122-3132200000200210-1131310112031023-3300312030233022-0231111012121030-3113222002103111-1021120111211330-1131302210032101"></a>

## matching_type property — values / 222123010030 / 6

Type: `"string"`. Computed.

\[Enum: NoHash|SHA256|SHA512\] - NoHash: No Hash - SHA256: SHA-256 - SHA512: SHA-512. Possible
values are \`NoHash\`, \`SHA256\`, \`SHA512\`. Defaults to \`NoHash\`.

Upstream description:

&#8203;- NoHash: No Hash

&#8203;- SHA256: SHA-256

&#8203;- SHA512: SHA-512.

Receipt-pinned upstream constraints:

```json
{
  "default": "NoHash",
  "enum": [
    "NoHash",
    "SHA256",
    "SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0301100322000222-2210310111030100-0120200130323031-0313300323302321-2020023000003023-2230202232131031-2302031133313230-0222231030103131"></a>

<a id="canonical-3013232211112233-1201131332220122-1300003123012032-0300303201310033-3012333030102322-2113133312032003-3223313302330112-1102300111202103"></a>

## selector property — values / 222123010030 / 7

Type: `"string"`. Computed.

\[Enum: FullCertificate|UseSubjectPublicKey\] - FullCertificate: Full Certificate -
UseSubjectPublicKey: Use Subject Public Key. Possible values are \`FullCertificate\`,
\`UseSubjectPublicKey\`. Defaults to \`FullCertificate\`.

Upstream description:

&#8203;- FullCertificate: Full Certificate

&#8203;- UseSubjectPublicKey: Use Subject Public Key.

Receipt-pinned upstream constraints:

```json
{
  "default": "FullCertificate",
  "enum": [
    "FullCertificate",
    "UseSubjectPublicKey"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2222212311223131-3200330312233023-3232232331110133-2111013233331122-1011332210320121-3312122002111113-0020012231323332-2301311113013221"></a>

## Next pages — values / 222123010030 / 8

- [primary.rr_set_group.rr_set.tlsa_record](data-sources--dns_zone--reference--group-003.md#canonical-1120111220133012-1033321320120331-3101013303033222-2303001213103323-1230312203202132-1302200000132221-0213201322100310-2310212123010121)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-2112000100003211-0010122023020330-2302300331333210-3123120333012012-1313313003312213-2120120330311231-1033321003121122-0030131112120013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110221123332000-0100113202301020-3233333232112311-1233123230323030-3212130223322320-0322323120001321-0032321131231313-2011201012233331"></a>

## primary.rr_set_group.rr_set.txt_record — txt_record / 323110100133 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [primary.rr_set_group](data-sources--dns_zone--reference--group-002.md#canonical-1123222033113212-3322310132230022-0003203213303322-3032311213113203-3003230323212030-2021000103002312-2303100332113121-0122303013013202)
- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- primary.rr_set_group.rr_set.txt_record

<a id="canonical-3120013321330311-3132233230321201-0111310300333022-1333330321012022-3022221032332233-0201113210031120-0121111121022320-2322222212030010"></a>

Type: `"single"`. Computed.

DNSTXTResourceRecord.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2210311300233332-0021220212320021-0112110201010231-3201010033211232-3102001111113122-3012332133313202-2331210200000131-1220203123112323"></a>

## Direct properties — txt_record / 323110100133 / 3

<a id="canonical-1230123002302333-0320100331032333-0332032213020213-2313010123113222-3130032100232102-0001031233212030-3111013300000210-2232023222320232"></a>

<a id="canonical-3332233202213003-0200103201323033-2003100102022003-1030222200201202-2220011001120123-1210221103310102-1200120132332130-2120100203313033"></a>

## name property — txt_record / 323110100133 / 4

Type: `"string"`. Computed.

TXT Record name, please provide only the specific subdomain or record name without the base domain.

Receipt-pinned upstream constraints:

```json
{
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
    "pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$",
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
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.pattern": "^$|^([*]|[a-zA-Z0-9-/_]{1,63})([.][a-zA-Z0-9-/_]{1,63})*$"
  }
}
```

<a id="canonical-2313020101101202-1330213132331132-0000110133301131-0012232110110003-3303020000303202-3130122103003331-0120203331232122-3001322231031230"></a>

<a id="canonical-0220023321103030-1023001110123002-1110222202311231-2322210013012322-3321103222100333-0103213010332331-1000223310323232-2223113021301003"></a>

## values property — txt_record / 323110100133 / 5

Type: `["list", "string"]`. Computed.

Text. Configuration parameter for values

Upstream description:

Configuration parameter for values

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4000",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1032202133011310-1003310200110332-3310313131133103-1221311213023133-3323202131331103-0100032313332013-2211310322131001-1312030021121012"></a>

## Next pages — txt_record / 323110100133 / 6

- [primary.rr_set_group.rr_set](data-sources--dns_zone--reference--group-002.md#canonical-0321130120130123-0321021202222021-1111221100311313-3310203033333202-2100230030032013-2312210231011130-2202301302333100-1330111100320302)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-3111301200310120-2012211111103233-2001110322021102-1132112033133212-1100110333112313-0210200023120120-3233033130322010-1032311322222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000323312202313-3313311222210220-3311022310213231-2000013112010100-2312312330112331-0222133303131330-2302003002323012-3312201321222033"></a>

## primary.soa_parameters — soa_parameters / 213111032122 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- primary.soa_parameters

<a id="canonical-3332200032033121-0332301033131232-0222202113000111-3302133332331313-1302101233232222-2003033200120003-1030130033320323-1003113213310120"></a>

Type: `"single"`. Computed.

Configuration parameter for soa parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1100321231311131-3102132121012021-3233333120031020-3303103130023223-2011303231111331-1312100303231312-0131332103000003-1102010312031121"></a>

## Direct properties — soa_parameters / 213111032122 / 3

<a id="canonical-1201033000303111-2223303120002331-0021033000232021-3123122112002012-1123232021130110-3023211303113021-3311312133110210-0002231111232302"></a>

<a id="canonical-0121020012321202-1003302201222223-2122020113220131-3122333230111202-1100212003322102-2132231032233003-3033301323321213-1200333101230201"></a>

## expire property — soa_parameters / 213111032122 / 4

Type: `"number"`. Computed.

Expire value indicates when secondary nameservers should stop answering request for this zone if
primary does not respond.

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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3100223313012320-2031021332301022-2331201032203202-2002023020101122-2110133031121231-3201313011302023-0120221133100100-3220113013302122"></a>

<a id="canonical-2231012012022110-0121030302321322-3221000133200312-1131022023123301-0300333233000310-3133131233132321-2331231203033103-2132113112121022"></a>

## negative_ttl property — soa_parameters / 213111032122 / 5

Type: `"number"`. Computed.

Negative TTL value indicates how long to cache non-existent resource record for this zone.

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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3120102023112030-3232121120322003-0021302112212101-2201012221110322-0220003032222203-2310310313302020-1202332331112213-0131133131130230"></a>

<a id="canonical-0223101302023122-0221001203313233-1110100130123232-0121001010323032-2322123301223102-0310021020001122-0333201221111021-1103231011232322"></a>

## refresh property — soa_parameters / 213111032122 / 6

Type: `"number"`. Computed.

Refresh value indicates when secondary nameservers should query for the SOA record to detect zone
changes.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 3600
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "3600",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-0132113021222221-2013001303010022-3221122332130011-0033113000132220-1303321123221002-0302320220223033-2122331011201321-2323110332310111"></a>

<a id="canonical-3122023011313033-1021331201200130-1330311333212211-0213313112201233-0201310003102133-0321320003031000-2022123313113322-3123130131012300"></a>

## retry property — soa_parameters / 213111032122 / 7

Type: `"number"`. Computed.

Retry value indicates when secondary nameservers should retry to request the serial number if
primary does not respond.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 60
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "60",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3213120213122202-0213310131301121-3202133221121323-0102313213023020-0023231121222111-1321320022330200-1010321303222311-0122010110212101"></a>

<a id="canonical-0012320212020203-3101112311220300-1101330213330311-2212331332230003-2101120013111312-3030310230103012-1300111310011011-2112003120230023"></a>

## ttl property — soa_parameters / 213111032122 / 8

Type: `"number"`. Computed.

TTL. SOA record time to live (in seconds)

Upstream description:

SOA record time to live (in seconds)

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
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "2147483647"
  }
}
```

<a id="canonical-3112330330210031-2010102312003301-3010331223132030-2223111212003322-0321033102030212-0203022311232102-3032332022011110-3120201121120230"></a>

## Next pages — soa_parameters / 213111032122 / 9

- [primary](data-sources--dns_zone--reference--group-001.md#canonical-3202300020132202-1031131002222003-2332211013000101-2012120102112221-3101102031020211-2011220131110101-1030320013331200-1333313321323331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011000233321030-1210221220011022-0223012133332002-1203301022213112-3103313021322001-2131210023301331-2232120020222311-3321133331111001"></a>

## secondary — secondary / 023233211202 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- secondary

<a id="canonical-3313001100221131-0333113010030320-1300031010002223-1231331333010010-0122201331100023-3022321102020300-0200013201131111-3320301233010323"></a>

Type: `"single"`. Computed.

SecondaryDNSCreateSpecType.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0313111032123333-1200023110023203-3332230001333222-2023223000210310-3312032012300232-0110003202103332-3301323033313203-3122032201033333"></a>

## Direct properties — secondary / 023233211202 / 3

<a id="canonical-2231231030010330-2312101322001132-3330301120301212-2132112131013123-3301303301023113-1220032023023021-1321133312132220-0110301231200022"></a>

<a id="canonical-3231302330132132-2132302322010212-1022032300313220-2020110012323213-3030223300300323-0331333100311010-2302300121003000-0021223200001333"></a>

## primary_servers property — secondary / 023233211202 / 4

Type: `["list", "string"]`. Computed.

Configuration parameter for primary servers.

Upstream description:

Configuration parameter for primary servers

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3330223310100001-0110311112003331-3013103213122000-2103321322103210-2012323002310000-1321203300233211-2200023201100210-2022222311131132"></a>

<a id="canonical-2230003322202212-1300013111332311-2110032323312103-0222012131223232-0110131103332111-0203221111021113-1223133131100120-0121231133131121"></a>

## tsig_key_algorithm property — secondary / 023233211202 / 5

Type: `"string"`. Computed.

\[Enum: HMAC\_MD5|UNDEFINED|HMAC\_SHA1|HMAC\_SHA224|HMAC\_SHA256|HMAC\_SHA384|HMAC\_SHA512\] TSIG
key-value must be compatible with the specified algorithm - UNDEFINED: UNDEFINED - HMAC\_MD5:
HMAC\_MD5 - HMAC\_SHA1: HMAC\_SHA1 - HMAC\_SHA224: HMAC\_SHA224 - HMAC\_SHA256: HMAC\_SHA256 -
HMAC\_SHA384: HMAC\_SHA384 - HMAC\_SHA512: HMAC\_SHA512. Possible values are \`HMAC\_MD5\`,
\`UNDEFINED\`, \`HMAC\_SHA1\`, \`HMAC\_SHA224\`, \`HMAC\_SHA256\`, \`HMAC\_SHA384\`,
\`HMAC\_SHA512\`. Defaults to \`UNDEFINED\`.

Upstream description:

TSIG key-value must be compatible with the specified algorithm

&#8203;- UNDEFINED: UNDEFINED

&#8203;- HMAC\_MD5: HMAC\_MD5

&#8203;- HMAC\_SHA1: HMAC\_SHA1

&#8203;- HMAC\_SHA224: HMAC\_SHA224

&#8203;- HMAC\_SHA256: HMAC\_SHA256

&#8203;- HMAC\_SHA384: HMAC\_SHA384

&#8203;- HMAC\_SHA512: HMAC\_SHA512.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNDEFINED",
  "enum": [
    "HMAC_MD5",
    "UNDEFINED",
    "HMAC_SHA1",
    "HMAC_SHA224",
    "HMAC_SHA256",
    "HMAC_SHA384",
    "HMAC_SHA512"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2032021232222221-3331232103113031-0222131211003030-0111213310012111-0120232201320103-3220323322013310-1223033222120332-1100022012211321"></a>

<a id="canonical-2032031032013120-3020303032001121-2322103101313020-2033032202013200-3311022211010102-2233231203122003-2023002210221132-0003321311002131"></a>

## tsig_key_name property — secondary / 023233211202 / 6

Type: `"string"`. Computed.

TSIG key name as used in TSIG protocol extension.

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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331): complete subsection reference.

<a id="canonical-0333001333012303-1202200311300111-1222102010220130-3023123203222030-3312333233113012-2101302232133033-2302113230031210-3330311032121332"></a>

## Next pages — secondary / 023233211202 / 7

- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122300113320102-0210103231020302-2301330303101212-0212002120201123-0301313010211222-0200331103010110-1203212031011322-0233201233020303"></a>

## secondary.tsig_key_value — tsig_key_value / 331221012223 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- secondary.tsig_key_value

<a id="canonical-3101233323332301-0312030330010020-0112200003233111-3030302033213132-1231033132112202-1221323011331233-3132101020311220-0322033320313121"></a>

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

<a id="canonical-0032123320231030-3033301313022311-1222030332021233-3132030331113112-0112331221132021-0101303323130013-1101222100331033-0122230031232032"></a>

## Direct properties — tsig_key_value / 331221012223 / 3

- [blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-1132333021301033-3331233222032213-2213011102012210-1120300021003122-3020030301333202-2313113001320103-1101220211020023-0103200133333003): complete subsection reference.

- [clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-0323113331312310-3323000001233012-3330320133313030-1133231201132300-2233133131302311-2333301023012213-1023201101130113-2221321221301002): complete subsection reference.

<a id="canonical-3333322111101322-3212122201030323-3213030101311023-2101302312203020-3000321001032232-0331222110011130-3223213232112122-3331203022110311"></a>

## Next pages — tsig_key_value / 331221012223 / 4

- [secondary.tsig_key_value.blindfold_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-1132333021301033-3331233222032213-2213011102012210-1120300021003122-3020030301333202-2313113001320103-1101220211020023-0103200133333003)
- [secondary.tsig_key_value.clear_secret_info](data-sources--dns_zone--reference--group-003.md#canonical-0323113331312310-3323000001233012-3330320133313030-1133231201132300-2233133131302311-2333301023012213-1023201101130113-2221321221301002)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-1132333021301033-3331233222032213-2213011102012210-1120300021003122-3020030301333202-2313113001320103-1101220211020023-0103200133333003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001210210011020-2133130230033120-0221331023212100-1223031123232031-2331201302011030-2301111231332002-0130301233222103-2330122332123102"></a>

## secondary.tsig_key_value.blindfold_secret_info — blindfold_secret_info / 201023331100 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- secondary.tsig_key_value.blindfold_secret_info

<a id="canonical-3322021301003221-1011000011013122-3221210313313322-1010211101010230-0021322012320102-3301222202013311-0330313211203121-3310320323300313"></a>

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

<a id="canonical-0223303311032112-3221001213310101-0031103302022200-0202230312100213-2101210202203320-0200332311210320-2023222110232020-1303030311001302"></a>

## Direct properties — blindfold_secret_info / 201023331100 / 3

<a id="canonical-2210020032131012-1313301020210112-2223100111000300-1320110301123013-0121113322102313-0110122232020112-2301213112313002-3122000002222200"></a>

<a id="canonical-1123301212202101-2222211320200032-2211311020120220-3312001131122313-3020333122331111-1100212220111300-0202121031121112-3123031130122000"></a>

## decryption_provider property — blindfold_secret_info / 201023331100 / 4

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

<a id="canonical-1300313031010331-1211322202132302-0323221133333301-2120210322212120-2320303010232122-3302202122310030-0301313211010021-1100333111010221"></a>

<a id="canonical-2112133331103300-0031302033023300-3220211122100202-0201021320023022-0332230102010132-1002232131301033-0031123112021312-0102120102012132"></a>

## location property — blindfold_secret_info / 201023331100 / 5

Type: `"string"`. Computed, Sensitive.

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the uri\_ref. It could be in URL format for string:/// Or it could be a path if the
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1332321201022032-0211212221300332-2022310122013230-2211233030113203-0220013000023200-1133221300112003-2122103012323023-0322031022312012"></a>

<a id="canonical-0123301310033221-0300003222221103-0130122110013111-1020022211002022-2233031120332022-2221223321332130-3223103021020303-3313013310203031"></a>

## store_provider property — blindfold_secret_info / 201023331100 / 6

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

<a id="canonical-0023321011203330-3100133122332303-2033132231210312-2030201002120010-2110200111111000-2332223032323012-3321133201010130-1130220123023033"></a>

## Next pages — blindfold_secret_info / 201023331100 / 7

- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)

<a id="canonical-0323113331312310-3323000001233012-3330320133313030-1133231201132300-2233133131302311-2333301023012213-1023201101130113-2221321221301002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301231120200213-2332023121211221-0003213122231212-3033112032213220-3023223220302203-0121301201203310-3002231220303113-2320303213120311"></a>

## secondary.tsig_key_value.clear_secret_info — clear_secret_info / 113022323331 / 2

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
- [Property reference](data-sources--dns_zone--reference--group-001.md#canonical-1320000120201111-2203011021030331-2320302132022301-3231311102321211-2320003331323333-0100003331333313-2321331201112230-0113213102020011)
- [secondary](data-sources--dns_zone--reference--group-003.md#canonical-0110011233113032-0002112311333113-1210112333111230-3033232211331300-2300022322300333-2131100301000032-2320031122203212-3103312010003023)
- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- secondary.tsig_key_value.clear_secret_info

<a id="canonical-0002213001331323-1200100121010123-2301120030233222-2203330130212332-1331302213010022-1330213231323202-2010321231203310-0110223012100222"></a>

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

<a id="canonical-3221021301023132-1033233313110022-0130112312131310-0322122022100033-2230330112212321-1232221210332110-0201220013310012-2013030120303000"></a>

## Direct properties — clear_secret_info / 113022323331 / 3

<a id="canonical-3012021300301121-2020233020332311-3000200220320123-2000003303023023-3202331302310133-2231133310330133-3302321211300300-3300030030031111"></a>

<a id="canonical-2231130232302313-2322102213322132-1031221130102213-2111022102120032-2030220232021320-1220211220021232-2202011123223330-0131332323123123"></a>

## provider_ref property — clear_secret_info / 113022323331 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1330123220130302-2101200121133331-3010022021113122-3331023211020110-2230132031312100-0101030030100311-3213331132110111-3103103113303113"></a>

<a id="canonical-2301223101110301-2103112323102233-0300033111212211-1312031202110223-2311030001202121-3223023020000122-1322212231212121-2303011013211103"></a>

## URL property — clear_secret_info / 113022323331 / 5

Type: `"string"`. Computed, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded Base64 format. When asked for this secret, caller will GET Secret bytes after
Base64 decoding.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1330032000332213-1211313210332132-0232322010033103-0032231222311322-1122200311121230-1313201130132022-2310102313233001-2001300320210310"></a>

## Next pages — clear_secret_info / 113022323331 / 6

- [secondary.tsig_key_value](data-sources--dns_zone--reference--group-003.md#canonical-0200300300112212-2013221101122013-0311222300002000-3123120223132313-1103101120032321-1012100100013213-3111132232013321-0100110220213331)
- [xcsh_dns_zone](../data-sources/dns_zone.md#canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131)
