---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-1300210103203222-3133322322012212-1212032213001303-3323113301003322-1033112201011300-2010300221132121-1200032301300233-3001012331021110"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_dualstack_on_public / 312110333330 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-3333011233220233-0132031303033103-1200130311003002-1220223133333222-3210331022303333-1313100030202130-0022003323210300-2312312312013212"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3233212033101130-2331213322323213-3230300302210302-0332011131030123-2201111033123213-2232011300022010-0332233321330110-1223032120022101"></a>

## Direct properties — advertise_dualstack_on_public / 312110333330 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0331232212113320-3221210122132321-1320231221130212-2112121003311112-0100320231023313-0002203332012013-0203220331112332-0331003102002122): complete subsection reference.

<a id="canonical-3233321333111313-3202213112021302-3023201022022000-3223230311321332-3033210122033322-0120303000321131-1122313103320120-3203100103013323"></a>

## Next pages — advertise_dualstack_on_public / 312110333330 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0331232212113320-3221210122132321-1320231221130212-2112121003311112-0100320231023313-0002203332012013-0203220331112332-0331003102002122)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0331232212113320-3221210122132321-1320231221130212-2112121003311112-0100320231023313-0002203332012013-0203220331112332-0331003102002122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333312022023332-1211012121113002-0210201101210232-3202332331230012-3211313232012132-2132230321030021-1132332100021121-1122301231120333"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — public_ip / 110200330210 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-0132123303300300-0302300023320120-1302232100300133-2220222110112113-0223322131221330-1300221223321322-0321220131002313-3321231233002031)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-3221332131020000-2223023303302233-3122330110113213-1300333121320221-1113203323233300-2012331110002132-0323300112000101-2122301211032031"></a>

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

<a id="canonical-3010111013221021-3103320211120020-2312022213003133-1001332033300112-0113233301100021-1100110002111232-0203023131031203-0002121023312021"></a>

## Direct properties — public_ip / 110200330210 / 3

<a id="canonical-3113202300333111-2031000321111301-3312303333333301-3003230321223202-0213333202203211-1231130123032130-2320233233211132-2020113231123030"></a>

<a id="canonical-2230321132223210-0030101121111110-3233232313012031-1113302223323210-3000010311322020-1132233311311132-1213120100120010-2311300332230011"></a>

## name property — public_ip / 110200330210 / 4

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

<a id="canonical-0121011210201103-1332310323122222-2002020310113220-1210122023132033-2021212303232202-2223212333202002-3003232123032303-3302303223000023"></a>

<a id="canonical-1020021333132110-0102113300322020-3223101223131310-3303030101111332-2131021132132311-1111311323322212-2321031021112332-1223200222022213"></a>

## namespace property — public_ip / 110200330210 / 5

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

<a id="canonical-2222312310301200-0031310330232132-3132332100202313-2023213103030200-2022233003320012-0330232100133130-2013221112112212-2030331100122132"></a>

<a id="canonical-3231232332221211-0020213131023102-3123323030203211-2010233100111320-3123202302300322-0222131222012313-0202011312020220-1101131101322120"></a>

## tenant property — public_ip / 110200330210 / 6

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

<a id="canonical-3011313323313032-1010332301112300-1322203002113030-1331032012232203-2200023001212230-0220123100233113-3033331033223131-1321200300110031"></a>

## Next pages — public_ip / 110200330210 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-001.md#canonical-0132123303300300-0302300023320120-1302232100300133-2220222110112113-0223322131221330-1300221223321322-0321220131002313-3321231233002031)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3333022132212222-3111201032102213-1013002313023331-0213231031023210-2222300310332023-1013233020000121-0221011301231012-2133231303223103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111103132032012-0231233100231332-2323110313202102-0223323033221331-1321011012233133-3112023311323132-0010021103103001-1212003130001322"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public — advertise_on_public / 100101311002 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-0102320013321202-2003132020213231-1031032130321121-3320101220023021-1010301112022212-1002333223131300-0133013302020110-1330220221310202"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1002122333203230-2201323221122121-2021203132030322-2220333333322032-3222321230003200-0121312311133032-1132031020032212-3022302310031202"></a>

## Direct properties — advertise_on_public / 100101311002 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-2233322310033003-0100313113220303-1331230033113231-0011210331322311-3300003302223121-3323111221123031-3323113321331021-3223100111032121): complete subsection reference.

<a id="canonical-0102133203211231-0132200133201311-3213322233332123-1010030031323231-1110113303132330-3331120002113103-2002100232002012-0103023002303203"></a>

## Next pages — advertise_on_public / 100101311002 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-2233322310033003-0100313113220303-1331230033113231-0011210331322311-3300003302223121-3323111221123031-3323113321331021-3223100111032121)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2233322310033003-0100313113220303-1331230033113231-0011210331322311-3300003302223121-3323111221123031-3323113321331021-3223100111032121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332312223131101-1211103223001223-3120323201202122-0323303020020211-3221311213011213-0033022001003302-1203201001321203-3110013002022331"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip — public_ip / 020012130321 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3333022132212222-3111201032102213-1013002313023331-0213231031023210-2222300310332023-1013233020000121-0221011301231012-2133231303223103)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-0331033300133222-3020331123101331-1103000323331302-3222021321231122-0123110301202120-1310113330311111-2300030212122203-3231023323323312"></a>

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

<a id="canonical-3100123102311132-3323330030330302-3022100302212313-0220131001331112-1001211123210133-3311300221110331-3320230001221000-2323033301220111"></a>

## Direct properties — public_ip / 020012130321 / 3

<a id="canonical-0120230132131222-2111030003302231-1330300131330110-1130033330321100-1010131021012030-2322222013312210-0121331011203003-1112220110320001"></a>

<a id="canonical-1223332100111131-2302201332003000-2010202020233121-0311111313102120-1210301032221212-3223001223123111-3130223223223230-2212331331312310"></a>

## name property — public_ip / 020012130321 / 4

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

<a id="canonical-2020110113012221-1200033231000212-0020322112333100-2221123221311100-2130330002220312-3020333203210032-2333210102222032-1300121220200332"></a>

<a id="canonical-2020031333031222-0320320033203231-3222011213323313-3111013330030102-3001123032231312-3000100210330323-1112030230031332-1132210122303303"></a>

## namespace property — public_ip / 020012130321 / 5

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

<a id="canonical-2033100113103210-2022101030230022-0002202212131000-2203013131002330-0233101230332301-2122013330202222-2333201030002300-1332312212322303"></a>

<a id="canonical-3010332310022102-2030331322322003-2322223011000123-2321002202322103-1012322300000332-2331121011033120-1210121301313022-2100110002003111"></a>

## tenant property — public_ip / 020012130321 / 6

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

<a id="canonical-1311323103211131-2102323120112023-3113233010102310-0201211233320001-0220013312131021-2201132321220120-1333310222012000-3300002023203310"></a>

## Next pages — public_ip / 020012130321 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3333022132212222-3111201032102213-1013002313023331-0213231031023210-2222300310332023-1013233020000121-0221011301231012-2133231303223103)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3011320001333033-3323010222033222-2312310300222031-3102010203100032-0322103312102132-1310220022003211-2201223332202331-3012320333320133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111300100110230-2133303132132123-3101330331212330-2203011302001121-3003223300103101-0101130030020333-0000031303212333-2011110032322102"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public — advertise_v6_on_public / 000331300123 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2311012002011031-1223120322223032-2113131002211203-3012101321121220-2120213001132301-1213102103133223-2122320020322001-0233333321001020"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132120223332022-1213023000301213-3313133212010232-3300013020030202-0302113230123333-3113132222122000-1301110030231213-2010122133211303"></a>

## Direct properties — advertise_v6_on_public / 000331300123 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-2130113210010113-2031300331003133-2332130320133102-1001003131231212-2303001312120330-2033310013100033-1122202011312311-3133132113323230): complete subsection reference.

<a id="canonical-0100021020010111-2131102331133121-0032200331011002-3112111302210123-0033330031313312-0021030213012332-2111100122112320-3223311132213131"></a>

## Next pages — advertise_v6_on_public / 000331300123 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-2130113210010113-2031300331003133-2332130320133102-1001003131231212-2303001312120330-2033310013100033-1122202011312311-3133132113323230)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2130113210010113-2031300331003133-2332130320133102-1001003131231212-2303001312120330-2033310013100033-1122202011312311-3133132113323230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031220232001233-1123230313211130-3200013022001130-2310200011032002-0300130113303010-3300212010202111-0123002013320122-2112231231321302"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip — public_ip / 333120022221 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3011320001333033-3323010222033222-2312310300222031-3102010203100032-0322103312102132-1310220022003211-2201223332202331-3012320333320133)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-3303133212201001-2323201302023212-3031232322031010-2203013022102122-1203213201030223-0022312212200001-1113301031320020-3023003112001122"></a>

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

<a id="canonical-3102030101122112-3100102022300030-2011330123012101-2230003113332022-2110002003332223-0203100022132313-3030002101330220-0332101313013013"></a>

## Direct properties — public_ip / 333120022221 / 3

<a id="canonical-2110331301031230-1022303220022312-3233202030311310-0221223322001203-1003112203332300-2333032030212132-0330221231031202-3112122102102200"></a>

<a id="canonical-0332332020011100-1013131332003201-1130123211133210-2301120212022010-2303320233211112-2330320120303122-1002302030221331-3213203300133322"></a>

## name property — public_ip / 333120022221 / 4

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

<a id="canonical-1231103032103101-2210013100311311-2130302200033133-0112200032321013-2231213003210303-3231332122002102-3330211112320021-2221010310003322"></a>

<a id="canonical-3030302113111213-2231010120232222-0331130203213133-2200010332312322-2211201000002333-0221302312233310-0031302211103231-1002031213102233"></a>

## namespace property — public_ip / 333120022221 / 5

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

<a id="canonical-0233332010022201-2303332230210232-0322022122211313-3131033213111333-0313212032220201-1031033203311331-2102311331332302-0220333210031321"></a>

<a id="canonical-2330231121032030-0210331210232230-3333233202200303-1222112221003322-0003323021333232-3303120203020110-0210303210210010-3020020311021223"></a>

## tenant property — public_ip / 333120022221 / 6

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

<a id="canonical-0310213000311121-1312103313221320-2003311033223001-0103222020011000-1102311222313321-2000032122233012-3000100300112302-0131123031321233"></a>

## Next pages — public_ip / 333120022221 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3011320001333033-3323010222033222-2312310300222031-3102010203100032-0322103312102132-1310220022003211-2201223332202331-3012320333320133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1202013332121321-0333321230221233-0300031321311331-0202212330103032-2001113111023230-2302012002012120-1001301130310311-1312022200123212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210321221113020-2232212020202131-0202121210300033-2212000333212130-2021200330012202-1211313031103133-0310021323202123-2110232211110233"></a>

## proxy_advertisement.advertise_custom.advertise_where.site — site / 331132031330 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-3002230103303213-1130233102321131-2032122030302012-2331322320012303-1022213021202010-2032312130311310-2203301213133312-2032013302233323"></a>

Type: `"single"`. Computed.

Defines a reference to a CE site along with network type and an optional IP address where a load
balancer could be advertised.

Upstream description:

This defines a reference to a CE site along with network type and an optional IP address where a
load balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0003213010230212-3330222130103001-1020121111220102-1202333303220111-2222312110032131-1302032023200001-2300232321233233-1300221100212122"></a>

## Direct properties — site / 331132031330 / 3

<a id="canonical-1132221200331030-2312132310313022-0002223320201022-0222331302200133-1023031332021211-0123232113001021-3003220212100320-1231111013023301"></a>

<a id="canonical-1023332323320321-2202023001132233-0301200223213110-0011012230001031-2313112203010101-0010132230302313-1023123021321333-0131011330013123"></a>

## ip property — site / 331132031330 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-0032001201113330-2220220203311302-0033010021002310-2002231302221221-0130320031332300-0320113300000102-1000300310200012-3220212022220313"></a>

<a id="canonical-0310331210011223-1202131103012313-2033010023110123-2201201132232020-3130231022233102-2012020013212333-0103112233132123-3013311100303232"></a>

## network property — site / 331132031330 / 5

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-0132321301022122-1323110133332132-2203323132312220-0130222212120211-0133113223021101-2333110103222101-3211323000202211-0231121031303233): complete subsection reference.

<a id="canonical-0023321332012321-2130002031121130-3201222203220230-0022333132333103-3233130201021301-3312232113233010-2310222000223000-3212113033203203"></a>

## Next pages — site / 331132031330 / 6

- [proxy_advertisement.advertise_custom.advertise_where.site.site](data-sources--dns_proxy--reference--group-002.md#canonical-0132321301022122-1323110133332132-2203323132312220-0130222212120211-0133113223021101-2333110103222101-3211323000202211-0231121031303233)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0132321301022122-1323110133332132-2203323132312220-0130222212120211-0133113223021101-2333110103222101-3211323000202211-0231121031303233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202232220320213-0001323310101332-1030212213312121-1013033203331333-2123203211310302-0211330033030121-1212301132020021-0223131132003103"></a>

## proxy_advertisement.advertise_custom.advertise_where.site.site — site / 302333012110 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-1202013332121321-0333321230221233-0300031321311331-0202212330103032-2001113111023230-2302012002012120-1001301130310311-1312022200123212)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-3330302210303232-3122012001122102-0123311121100203-2332212012030202-2312333303223021-3021222200321200-1131233203320131-2120032102303301"></a>

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

<a id="canonical-3000302002103132-2000232022200011-3322323233323030-0012233033221131-3023020300301221-3333030123333000-3000232230203201-0000120032231302"></a>

## Direct properties — site / 302333012110 / 3

<a id="canonical-1201113331101030-2320323211103110-0231103302300221-0133303112002000-2302021133303031-2223320201000213-1212200200203311-1130122013231321"></a>

<a id="canonical-0031323313021033-2121313232221220-2313123102131310-3123312102312200-0111002002202120-1012210202201130-0011123203102232-3130212112200212"></a>

## name property — site / 302333012110 / 4

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

<a id="canonical-0301323222203222-1103223311303200-3312030032012003-3223303113030110-0032022003132210-2120200033010031-1000021120322323-0131121012113103"></a>

<a id="canonical-2313101012023122-3011212012311033-3322112211322012-1101230021320231-3101221020011210-0210133331103103-3330330222221310-0113231211210033"></a>

## namespace property — site / 302333012110 / 5

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

<a id="canonical-0002010213113313-0201133320121003-1023202001331320-2011300332211331-3002331023210213-2203031103221102-2223201120113102-0223132313332003"></a>

<a id="canonical-2232002200112211-1123121213123220-2211202312113331-0300030001003223-2332302033111102-2031202122132123-3100201112233210-3022330122212003"></a>

## tenant property — site / 302333012110 / 6

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

<a id="canonical-2331303213230313-0032322100013033-2212222022301303-1100021230133302-3231012013300003-0313012011032122-0122001301012202-3100303002231021"></a>

## Next pages — site / 302333012110 / 7

- [proxy_advertisement.advertise_custom.advertise_where.site](data-sources--dns_proxy--reference--group-002.md#canonical-1202013332121321-0333321230221233-0300031321311331-0202212330103032-2001113111023230-2302012002012120-1001301130310311-1312022200123212)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1322010222011223-3333222022032231-0303010321312313-1213220133011033-0310121331102331-1013011301322221-0033302111332320-1211310333310013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102130330112301-2131300203113330-3211313201112012-0231202123212132-3131202221230110-2330130330022310-3010011212103033-1223331031223032"></a>

## proxy_advertisement.advertise_custom.advertise_where.use_default_port — use_default_port / 103201222010 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-0030223121030210-1312320332032211-2132220111210333-3120012113201313-1201011211211312-0230323203332331-2100131112231211-2213300021132021"></a>

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

<a id="canonical-0203211233322002-3300033333210133-2230010311130300-1030220131110112-1122112020130031-1223112002032200-0131301113323110-3101111222303321"></a>

## Direct properties — use_default_port / 103201222010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331023222022331-3123212021103111-1030130020103333-2320233131122320-1203110031211222-1222301111120221-2233212023303320-2120320313023322"></a>

## Next pages — use_default_port / 103201222010 / 4

- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213221101323330-2031122100002333-0113211011132122-1302111320031100-0323331220320112-3001123030300203-3311003232023010-0233320332121110"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network — virtual_network / 222032013320 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-2331322221302011-3011100222100032-2112223213132132-0130233310221230-1333323303333101-3320202323320212-3232021230101321-0333331200220302"></a>

Type: `"single"`. Computed.

Parameters to advertise on a given virtual network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

<a id="canonical-2100332311033101-2322120200131231-3320232023112121-0001321000001301-3232200320321013-3213302312030222-3012320011011130-3332212301131023"></a>

## Direct properties — virtual_network / 222032013320 / 3

- [default_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1110011301133202-3202122320022121-3331110010002210-1201113010313003-1301220222002001-0300033113231011-0312222330030332-2212223221022233): complete subsection reference.

- [default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-3033103102332201-0132311333110201-3201011033310103-3333202313112223-1110300130132201-3103031310311222-3320212311022311-2123122102133211): complete subsection reference.

<a id="canonical-1203001012231023-2200120113320101-0022301103303333-2130322003311132-1022001233211220-2310333023012330-2221331123001310-3121003323033121"></a>

<a id="canonical-1000202313221033-2301202103013300-0231120130133200-0202300313103301-1301110311113223-1330213031212320-1012123222030113-2132313003223121"></a>

## specific_v6_vip property — virtual_network / 222032013320 / 4

Type: `"string"`. Computed.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6": "true"
  }
}
```

<a id="canonical-0010203030001200-2130022020310002-3022300332011011-0202232131123321-2203321211320321-0012002102131312-2320030323231322-0033103130222112"></a>

<a id="canonical-2313230320210232-0011113123233200-1121220111320103-3132312331030212-1012331231131031-2121031130221130-2212202233322222-2032023032203312"></a>

## specific_vip property — virtual_network / 222032013320 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-3012021201022203-2013212000330112-1131122233120112-1211221310102330-3112322003321031-2333203233233301-2303130033111032-0000021012112310): complete subsection reference.

<a id="canonical-2302231222122230-3301320201030302-3302131312113300-2230122213010111-3311213131313131-2013210130011033-3300311102232221-0100101100002102"></a>

## Next pages — virtual_network / 222032013320 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](data-sources--dns_proxy--reference--group-002.md#canonical-1110011301133202-3202122320022121-3331110010002210-1201113010313003-1301220222002001-0300033113231011-0312222330030332-2212223221022233)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](data-sources--dns_proxy--reference--group-002.md#canonical-3033103102332201-0132311333110201-3201011033310103-3333202313112223-1110300130132201-3103031310311222-3320212311022311-2123122102133211)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-3012021201022203-2013212000330112-1131122233120112-1211221310102330-3112322003321031-2333203233233301-2303130033111032-0000021012112310)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1110011301133202-3202122320022121-3331110010002210-1201113010313003-1301220222002001-0300033113231011-0312222330030332-2212223221022233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231312233032010-2102131302123303-1300031331100322-1301323110002033-0021211201202302-1203300313210031-0011312013103013-0210232330021012"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip — default_v6_vip / 021022011123 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2120300033223332-0212003101300300-2101201012221121-0113103100130113-1211203203003233-2203010222330231-1332032200130301-1332101203131313"></a>

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

<a id="canonical-1201223321222022-1122222132121021-2200311230332023-0212122222231210-2332333311322322-0133211132001101-0310301103301022-2002231102221102"></a>

## Direct properties — default_v6_vip / 021022011123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320110301312123-1302113211002223-1201122322111012-3220133011111320-1000300200133130-0331111121331331-1103233330100121-0311031202012312"></a>

## Next pages — default_v6_vip / 021022011123 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3033103102332201-0132311333110201-3201011033310103-3333202313112223-1110300130132201-3103031310311222-3320212311022311-2123122102133211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212013313312122-0121331312033300-1100231221232010-3130223131210232-1232130333221001-3002031123121201-3223302002332012-2112131011303020"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip — default_vip / 222302121203 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-3331302000200313-1002221312100232-0120212033130322-1031103310303210-1320222122131212-2321123302123101-3330102010230003-0032130313133111"></a>

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

<a id="canonical-0201023203211022-0203212331302320-1331030210023013-1310031202313201-0212323212001300-2012330033102000-3212001123200103-2200103322321313"></a>

## Direct properties — default_vip / 222302121203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120333323030302-2333133303131231-2101322101203302-2332120030003301-0200310312013102-2331100103131201-3220023333330112-3022213213212003"></a>

## Next pages — default_vip / 222302121203 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3012021201022203-2013212000330112-1131122233120112-1211221310102330-3112322003321031-2333203233233301-2303130033111032-0000021012112310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231100011032333-1303103100123123-1003131233013123-2123332012311213-3123130031333101-0113033122312013-0213111102110122-1310223002302311"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network — virtual_network / 313120300310 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-0203010212131123-2330300232303102-1103313022101132-0211010211130121-2133200333110230-3130203022023103-1021321322223332-1200333210210223"></a>

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

<a id="canonical-3032222213311231-2323020023010220-0000111110003310-3100221321300101-1121320110313223-0010033222110123-1220022313003012-1003213230210113"></a>

## Direct properties — virtual_network / 313120300310 / 3

<a id="canonical-2311133313101213-2021133103021123-2323002033030120-3011003123030130-3313111232103231-0321312010000003-2213313223213021-3100300010323302"></a>

<a id="canonical-2022223123110321-2320000000033222-3213011022302132-0232020020202210-0011013223213212-2033100211002021-2132003101332320-2113123012312300"></a>

## name property — virtual_network / 313120300310 / 4

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

<a id="canonical-3003031010233103-2013001132101032-2222222133113222-0230211101222311-0112213112000232-3330112110102200-0220313012221330-1020011323123023"></a>

<a id="canonical-3113301002130310-3030330313321203-1121220023013000-1323023232303030-3333330021013201-1000331212320122-2103201300223121-0130333222101012"></a>

## namespace property — virtual_network / 313120300310 / 5

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

<a id="canonical-0210211020220132-3200312320131231-1032011233300200-2022322110033032-1103100001303112-2210132020222110-1333000221101310-3233022122212110"></a>

<a id="canonical-2130011333202011-3222013130002130-0311011321113013-2031330012000300-3312001303311322-0022231133310103-0222213302013133-1031111020333120"></a>

## tenant property — virtual_network / 313120300310 / 6

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

<a id="canonical-3012030211310112-1212102020221321-1020230110321202-1230000132301000-1302210011313120-3232131231221231-1330222130220310-3313213122100320"></a>

## Next pages — virtual_network / 313120300310 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](data-sources--dns_proxy--reference--group-002.md#canonical-1210321133310320-2021112303102212-2320133003123233-1211100103020321-2023323111103133-3210003112302123-0003332023102101-3033302113213130)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0032203302302013-3111220213021010-0303003102222033-2220212203222302-1212122000221101-3333223131032300-0332333012122110-1013233212322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211122331313122-2210032001000130-1233112330000231-3332132220302101-2130212230003003-0132203130333103-2201210033301231-3323123103323301"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site — virtual_site / 221010013121 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-2030102301102310-3322020220231032-3033002123231113-1030022231331202-3231312301032213-2301023233022303-3031131201132213-2220032200033123"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type where a load balancer
could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1130000100201021-2200322210212210-1230201210313030-1121030020331121-3131301303112030-2120223023000333-2032000222212333-3110231130032020"></a>

## Direct properties — virtual_site / 221010013121 / 3

<a id="canonical-0332203300131313-0233111331002231-3030103103333230-3301031331232012-0110031302332022-1122001011312131-0222331321203132-3213132101002323"></a>

<a id="canonical-0001122032303120-0001203011220200-3331202303110233-1303332033322003-0332210103020012-3112321321111232-2033203232012010-2210021113210002"></a>

## network property — virtual_site / 221010013121 / 4

Type: `"string"`. Computed.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Upstream description:

This defines network types to be used on site

All inside and outside networks. All inside and outside networks with internet VIP support. All
inside networks. All outside networks. All outside networks with internet VIP support. VK8s service
network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_INSIDE_AND_OUTSIDE",
  "enum": [
    "SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2012022001232103-2333011021301010-1022323013200032-2132200313111213-1210231313331331-1101121323103202-3111221320333330-1122310023121203): complete subsection reference.

<a id="canonical-0212002301000331-1113112131120010-1303121232102113-0111202213223123-3002233033330200-1101313111002133-2212013021320211-3302130312133333"></a>

## Next pages — virtual_site / 221010013121 / 5

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2012022001232103-2333011021301010-1022323013200032-2132200313111213-1210231313331331-1101121323103202-3111221320333330-1122310023121203)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2012022001232103-2333011021301010-1022323013200032-2132200313111213-1210231313331331-1101121323103202-3111221320333330-1122310023121203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232100003322213-0310121110032122-0011301321112133-2033232011203201-3300100032333223-2012310210200010-0112320332320023-3310221032301002"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 321023210011 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0032203302302013-3111220213021010-0303003102222033-2220212203222302-1212122000221101-3333223131032300-0332333012122110-1013233212322320)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-1230302220330002-1201213102013312-2333020322031000-2303231223301033-0013101310221231-1101133203012323-3230020110101223-0100203131023232"></a>

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

<a id="canonical-0220212111111221-2232033223032310-2112112221110032-3223321030231323-0030231021222202-1203003323123303-1231332221201133-3202100332312002"></a>

## Direct properties — virtual_site / 321023210011 / 3

<a id="canonical-2032030110020013-0012111102233320-0011233121202322-2122220311313123-3130121333000231-0312013310231300-0322312213221201-3330311113212023"></a>

<a id="canonical-0012000131333130-3321130133202213-1233230211321031-3013320003011313-2032320320231231-2331003001312313-1233222021002121-2112301112102331"></a>

## name property — virtual_site / 321023210011 / 4

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

<a id="canonical-0222020302101131-2211101131110232-2011221000010132-2010313321230320-3002133310331212-1231320312202221-2322200203033110-3002310210312122"></a>

<a id="canonical-0111110030122123-3321020231312120-0110310203201203-2302303230100010-0022222120221233-1320311032201032-1230213210123011-0000120002031211"></a>

## namespace property — virtual_site / 321023210011 / 5

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

<a id="canonical-2032333332030030-3202030323230123-3321221002101010-0222003313202013-0333031223032002-3123332021313110-1032111230203200-3203303323322201"></a>

<a id="canonical-3222001101330200-0230233000111232-3113133011123133-1230301232101221-3231002320333332-2301333231123123-0220311122231210-3230130323110122"></a>

## tenant property — virtual_site / 321023210011 / 6

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

<a id="canonical-2020131011003033-1332102303100131-1210010121101023-0023300010333010-1300223203302333-3330122232100032-2220101303102011-1010300203103111"></a>

## Next pages — virtual_site / 321023210011 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-0032203302302013-3111220213021010-0303003102222033-2220212203222302-1212122000221101-3333223131032300-0332333012122110-1013233212322320)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000013213200133-2220101231332203-3032113330331211-0322123000002220-0220123123221220-0210103031102010-0032131023220110-1332300330302022"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip — virtual_site_with_vip / 323212120121 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-1332100020122100-2213112023010000-2322331010010130-1113001101313131-1132321130101011-2103010110333013-2011103123233303-1312022022102233"></a>

Type: `"single"`. Computed.

Defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Upstream description:

This defines a reference to a customer site virtual site along with network type and IP where a load
balancer could be advertised.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1212320101303112-0001323333333121-2103311211201201-1232323222323021-2023110100221232-0332232030323330-2022213221222110-0323122100202213"></a>

## Direct properties — virtual_site_with_vip / 323212120121 / 3

<a id="canonical-0003331001023312-1213212033132123-0232221022122222-0030231331013012-3011001033100033-0203223231300133-3000130212310332-1001002213132201"></a>

<a id="canonical-2000203002311022-2332312332120230-0110000023002133-0300201232222300-3123101003112021-2012313013001310-3112300132021113-1121330201000132"></a>

## ip property — virtual_site_with_vip / 323212120121 / 4

Type: `"string"`. Computed.

Use given IP address as VIP on the site.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3021000131331310-2113113231220321-2210201320223000-1232321320122223-1233200320332133-0021021221323331-1033302111310302-1003013012333332"></a>

<a id="canonical-0222010002000323-1230320321202211-0012002211001230-1332120010320300-2300110322302110-0133310203000112-2220231023231311-2223220030012203"></a>

## network property — virtual_site_with_vip / 323212120121 / 5

Type: `"string"`. Computed.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Receipt-pinned upstream constraints:

```json
{
  "default": "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
  "enum": [
    "SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2113023311201301-3102123023111311-0101233331301221-2012003322331012-3323013312013221-1311303011302101-2132201230221112-1321203031122231): complete subsection reference.

<a id="canonical-2232010020323322-1300233321102100-1331101201003111-1332131330012213-0003321300122201-2032003013103002-2300220201131102-1000221132210300"></a>

## Next pages — virtual_site_with_vip / 323212120121 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-2113023311201301-3102123023111311-0101233331301221-2012003322331012-3323013312013221-1311303011302101-2132201230221112-1321203031122231)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2113023311201301-3102123023111311-0101233331301221-2012003322331012-3323013312013221-1311303011302101-2132201230221112-1321203031122231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101113311013022-0000223202320231-2222201122333013-3302132100331010-0020301111330210-3102211112113323-3120131132302112-1100331022120233"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — virtual_site / 030101213301 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-0321113323330103-1222322133010120-0233132221011111-1321122110102210-1112231321310303-2203102111213301-1113310000132302-0330231123301021"></a>

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

<a id="canonical-1110212202121012-3132122312211211-2012210021303331-3332332010302300-1022033022330123-1103300120222303-2200323032302031-1322301030131220"></a>

## Direct properties — virtual_site / 030101213301 / 3

<a id="canonical-2212102111221101-3100300312221123-1100020321303032-1200231221013030-1321303222003303-2333111201022010-0333320001000010-0032101301212023"></a>

<a id="canonical-0032202213103032-3202213221220111-1120300202123022-0203132103003223-3000032001322010-2330320301231113-0021221223302133-1303002321332013"></a>

## name property — virtual_site / 030101213301 / 4

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

<a id="canonical-2030310102002022-3102331323011030-2132221222201003-1131330011201313-2132011213023202-2021120300220101-1220022133311221-0101033321231013"></a>

<a id="canonical-1101001311302131-1001022122112221-0021111111310132-1131322211020031-0301033332201212-3002103323001113-0212230103100310-1000132313011230"></a>

## namespace property — virtual_site / 030101213301 / 5

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

<a id="canonical-1101023220003033-2313312100323330-3210321223312300-2220203003302132-1322130313323033-2310101211311121-0033313022201100-3201003212012233"></a>

<a id="canonical-2223332003200030-1032001210201020-2200102010103012-0113101001010030-1211013221330013-2033101211102122-0112303231011222-3010020200233122"></a>

## tenant property — virtual_site / 030101213301 / 6

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

<a id="canonical-3200032202132310-0310110222113313-0310303323320133-2101120200200032-3221003230013233-1322310110330112-0212332233021333-1221123203301331"></a>

## Next pages — virtual_site / 030101213301 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](data-sources--dns_proxy--reference--group-002.md#canonical-0033131110232301-2002233322320212-3111233001121322-2023001311103132-3232133223101330-2202030133010212-2111220111013323-0233220313313133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1223200003011000-3030323000030023-0013320021311211-3113300323220122-2033331232211213-2233031122033330-0023013023222110-0120220303213103"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service — vk8s_service / 023312302213 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-1033112100123332-2013010212202232-1333211020033213-0120032111021023-0233033021033101-1203002012112211-2110220212023321-1230300000122231"></a>

Type: `"single"`. Computed.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-1201310222312312-1211030222202012-3213013023311313-1200211010001320-0110003123012200-1231211332312011-0102333230322210-3022030012103202"></a>

## Direct properties — vk8s_service / 023312302213 / 3

- [site](data-sources--dns_proxy--reference--group-002.md#canonical-1333333123130202-1310021311112302-3233033330332110-2002110010032233-1031113301212123-3033003233023030-3221321013303232-0000131332301200): complete subsection reference.

- [virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-3311332103031002-3020231101022322-1133023003222201-3101010220323231-1023320312310130-3003302013311030-0221110133012303-2000220023100331): complete subsection reference.

<a id="canonical-3230021311130203-2203213330233101-1010312331012111-3130223212111000-1211001131023230-1221231002111133-0022131111133222-3202121311220020"></a>

## Next pages — vk8s_service / 023312302213 / 4

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](data-sources--dns_proxy--reference--group-002.md#canonical-1333333123130202-1310021311112302-3233033330332110-2002110010032233-1031113301212123-3033003233023030-3221321013303232-0000131332301200)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](data-sources--dns_proxy--reference--group-002.md#canonical-3311332103031002-3020231101022322-1133023003222201-3101010220323231-1023320312310130-3003302013311030-0221110133012303-2000220023100331)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1333333123130202-1310021311112302-3233033330332110-2002110010032233-1031113301212123-3033003233023030-3221321013303232-0000131332301200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210300021310221-0120212321302120-3323121101313021-0201031010323102-2210303322033211-1002133200301230-0001320003013000-0220010312121313"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site — site / 321023112302 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-3102211002111302-2222211312331122-0031322132221331-3012231000123103-1223231203330221-1102003221333021-3333310232323222-2000302312212130"></a>

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

<a id="canonical-2230331111123231-3000223332311230-2111230101331010-2111201233320212-1111222322132000-1113323212331023-0212200332303313-2312212131310302"></a>

## Direct properties — site / 321023112302 / 3

<a id="canonical-3200323101031203-3211102110021310-3323303132003123-3330210011000220-0033302310103312-1303323001113212-3003132232323003-0202010103203123"></a>

<a id="canonical-3031223210220002-0232203000332210-2310113303313022-1221303132113223-0320033100322313-1121221023220321-3120321200222100-3133101332102111"></a>

## name property — site / 321023112302 / 4

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

<a id="canonical-3100230232202330-1213233121303210-2223212101222022-1121321000111203-3331303222203323-0213332021210010-1011010212003012-1011031203100111"></a>

<a id="canonical-2213332231032312-3020331130200201-0111123223330100-0300100223010130-1222121110230330-3132011030312310-2221223032003213-1123223100011010"></a>

## namespace property — site / 321023112302 / 5

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

<a id="canonical-3211311003012203-3300013122331133-3101231222011223-3011321212301330-3232313200113220-2313223323230203-3133130301001221-1201320200223222"></a>

<a id="canonical-1201002320003331-0132130303120230-3132320232101001-3122321311110033-2213321333201121-0212222010232203-0210310302022111-3313120312010332"></a>

## tenant property — site / 321023112302 / 6

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

<a id="canonical-3300212131311210-0033223301233113-2030100210333011-3022023133312100-3011330302201201-1333211013300122-1113223312223103-0331210203210002"></a>

## Next pages — site / 321023112302 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3311332103031002-3020231101022322-1133023003222201-3101010220323231-1023320312310130-3003302013311030-0221110133012303-2000220023100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321222131201122-1330233313033330-3013333203223130-0133110013212210-1101021021310100-0320000003311332-2022203000100303-1230313111311323"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 130333322022 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_custom](data-sources--dns_proxy--reference--group-001.md#canonical-0121320103123000-0212221121020201-1332101102231101-0230012010231220-1120001321320132-3232322222212322-3211300112301312-1201311231022101)
- [proxy_advertisement.advertise_custom.advertise_where](data-sources--dns_proxy--reference--group-001.md#canonical-0310210230001331-3033021211312312-1002120201130001-1301322301121003-1122112322223303-1331210131110332-2313010230213100-2302012312313220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-1020330201012223-1013003233110010-2120313231101200-1033232133013211-2312320323231303-2213320102131223-2111313033313313-1310100101202113"></a>

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

<a id="canonical-2332103012333012-2231220222200311-1323112010102302-2330222202311030-0222232011012333-3302130133231022-0333122220101200-0221301013031203"></a>

## Direct properties — virtual_site / 130333322022 / 3

<a id="canonical-1020210112011102-0032103032012201-2110310110332023-0130203302120322-2032231222110020-2231132203132311-2303301320302331-2132130011300122"></a>

<a id="canonical-1330230233001032-3231231122311331-0103010131111133-1310001121121112-0020203110223223-1030301313112332-2032030201132112-2001332323311102"></a>

## name property — virtual_site / 130333322022 / 4

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

<a id="canonical-2022123211100103-2301211113323202-1321001223021030-0201333100112322-1313311001231231-2030202002000312-0213023132231232-2031212113302010"></a>

<a id="canonical-1222100320331232-3321333210113132-3010212311203230-0320030331313113-3011102202233323-2012032233312100-1323321231122120-1133001120010302"></a>

## namespace property — virtual_site / 130333322022 / 5

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

<a id="canonical-1333222213132301-0301212022320032-3101201223022103-3122312203300303-3323130120131210-3233001122120021-3221202230311121-3222313030303110"></a>

<a id="canonical-1312012311113032-1010220023320102-3320120320130321-0013301310302313-3223222221333231-3033011003333300-3032033201101202-0002322131201010"></a>

## tenant property — virtual_site / 130333322022 / 6

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

<a id="canonical-2113112310013023-2002022333331313-3021012131211110-1130213222203232-3033212232110033-3332023013120220-2023100001303221-1222311123010113"></a>

## Next pages — virtual_site / 130333322022 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](data-sources--dns_proxy--reference--group-002.md#canonical-2033011233111310-3303121203100330-1200131323233001-0011132112303132-2021223312002302-0232310330222110-2120123123021030-3030010023303130)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002331320302033-3202303313220012-1103302131210031-2010023322330233-3032033130001120-1011030013221210-3013011333231013-2213222002203303"></a>

## proxy_advertisement.advertise_dualstack_on_public — advertise_dualstack_on_public / 110110223123 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_dualstack_on_public

<a id="canonical-0321120322221213-3230322211221020-3033020113111230-3330130321232023-3212302102201312-3202323102023302-1111220132323122-2323200203010022"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3330211013203010-2223210033110112-0122213313321322-3221030311312012-0202200232132303-3012121103100312-0202120203232003-0133013311231100"></a>

## Direct properties — advertise_dualstack_on_public / 110110223123 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-1010032223032131-2222112020111232-2030210133123303-2023212111030102-3303322102211213-0212330211213013-2101323101230100-1013202012100331): complete subsection reference.

<a id="canonical-0230303230322330-1120010213221222-0301000103210021-2020002010130213-1332012320221110-1220100310303210-0100031301023311-1010031232022001"></a>

## Next pages — advertise_dualstack_on_public / 110110223123 / 4

- [proxy_advertisement.advertise_dualstack_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-1010032223032131-2222112020111232-2030210133123303-2023212111030102-3303322102211213-0212330211213013-2101323101230100-1013202012100331)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1010032223032131-2222112020111232-2030210133123303-2023212111030102-3303322102211213-0212330211213013-2101323101230100-1013202012100331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130321010031311-3310202021102202-1322222123322330-1331123301032110-0122212130221011-1031232210102312-2202301122202223-2123320103033123"></a>

## proxy_advertisement.advertise_dualstack_on_public.public_ip — public_ip / 322323000210 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121)
- proxy_advertisement.advertise_dualstack_on_public.public_ip

<a id="canonical-3132031011321222-1222130130003221-1011323212132333-1200311333232201-3033211212101232-2321002132302213-3200121312022003-0132113311131320"></a>

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

<a id="canonical-2303020322000031-2111231321032212-3223131132232003-2312332003023323-0012102320223133-2113023002100210-1332010310120312-2331210131312210"></a>

## Direct properties — public_ip / 322323000210 / 3

<a id="canonical-0330211001301212-3011302010302213-0211330220211020-2232022321302301-0022201023032033-0011301201223232-2102133012102013-2331002302100123"></a>

<a id="canonical-3121132001110020-2011222100001002-1122132012202111-0311333022222102-1123300221022311-1301032121203300-2213013000233300-1202020221312000"></a>

## name property — public_ip / 322323000210 / 4

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

<a id="canonical-0101313111332023-1323012133113303-2222132230032020-3233000220233120-3200232111321012-0301111221323132-2211112331313203-3203303110132323"></a>

<a id="canonical-0320001122110233-1221201100232213-0303323203201321-3001333231231301-0332012123100300-2113330023201020-2112100300010133-1110100233333331"></a>

## namespace property — public_ip / 322323000210 / 5

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

<a id="canonical-2300201110000333-1212300211222212-0112132022012221-0131003301103013-0002300202130132-0323311133210122-0200221131201102-2220111021322302"></a>

<a id="canonical-3321220211333030-1332011133321000-0012003112013301-3033300030000131-3112203302022231-0030102002131323-3221212133111303-0023000011033123"></a>

## tenant property — public_ip / 322323000210 / 6

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

<a id="canonical-2232211212203303-2012221321101322-0120002101010132-2311002013132110-1230311311132132-3110103113303210-1111120200021122-1111301211123133"></a>

## Next pages — public_ip / 322323000210 / 7

- [proxy_advertisement.advertise_dualstack_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-0201333130332013-0112020301131201-2021022313321201-3013233130311223-3022300202201021-2131301231003023-3121323331203223-2002102000332121)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233102100000012-1120110002213123-2112320302230022-3100123030131223-2320000232223312-1102020132100121-3212311131102210-3200133122023032"></a>

## proxy_advertisement.advertise_on_public — advertise_on_public / 120122232323 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public

<a id="canonical-2201111330223031-3100333012303300-1322103101012023-3113003022102223-2113331212030300-2101002211303002-2312100311212232-3220203201203203"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1111233312102023-2231110200131122-3010123110320332-0121233333311003-2303012230331003-0231230013221033-2013332333011023-3200103231210211"></a>

## Direct properties — advertise_on_public / 120122232323 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0303001021330002-0331320211123221-2320033130031302-3230331111330111-0012312332123311-3112122330203003-1222112033022233-1103301022002213): complete subsection reference.

<a id="canonical-3230000033012222-0013003301102301-1301121001000102-1031000030311202-1320032022321000-3213201202033231-0022023103111021-3130201313020332"></a>

## Next pages — advertise_on_public / 120122232323 / 4

- [proxy_advertisement.advertise_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0303001021330002-0331320211123221-2320033130031302-3230331111330111-0012312332123311-3112122330203003-1222112033022233-1103301022002213)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0303001021330002-0331320211123221-2320033130031302-3230331111330111-0012312332123311-3112122330203003-1222112033022233-1103301022002213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131202012313113-2020320202230232-0232220300030312-2123131223200012-3020301113100222-2213322033223300-3103010330031323-2012223111213233"></a>

## proxy_advertisement.advertise_on_public.public_ip — public_ip / 231222121010 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130)
- proxy_advertisement.advertise_on_public.public_ip

<a id="canonical-0311110300000112-2122110131010333-2302313301012321-1233101230312222-2131300313320323-3223230201323122-0132320200003303-1112203021210312"></a>

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

<a id="canonical-1321230331200023-3030332312231002-1101010210123213-2300210210221301-0023221102210113-1212023022233301-0321033100032302-2011310222330012"></a>

## Direct properties — public_ip / 231222121010 / 3

<a id="canonical-0313301030330320-3212022003311130-0111130200010033-1332003012230031-1020333003200220-1300010113332202-0300200322201132-2002103000132133"></a>

<a id="canonical-1223133033211332-2221031311302211-3232011233101032-1101200020312013-0201000131122011-1223203102131233-2002233200021111-1220113030203302"></a>

## name property — public_ip / 231222121010 / 4

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

<a id="canonical-2100010122300000-1011033203333301-1010010023022330-2031030321120320-3120031123031132-1221130123230321-3110112301221131-2301322220130300"></a>

<a id="canonical-0231300022213002-0322032001112312-0130103111133121-3313211003131132-2010030323311312-2000211033312131-0202111001322011-0110013102132322"></a>

## namespace property — public_ip / 231222121010 / 5

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

<a id="canonical-3010310230103100-1331012023321322-2300303011213312-2113121322022121-3103013300132103-3033310200021301-3230022310013123-1301311321113030"></a>

<a id="canonical-3020313023320233-2132311313202010-3113322130033322-1223021321121103-3100222302003111-1223232113311112-0320301001212303-3311123212322210"></a>

## tenant property — public_ip / 231222121010 / 6

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

<a id="canonical-1132320201002000-3311110330223113-3220200310011331-0330033202023330-0131333110300000-2002123213003010-2221012220221220-1310111013211200"></a>

## Next pages — public_ip / 231222121010 / 7

- [proxy_advertisement.advertise_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3310110220212000-2133311323002011-3330302201310131-3311331023300100-3031000200023101-0321033223132322-1202213132212101-2303002101010130)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1131222002311300-2222212030001300-0012300103101232-0031023232101211-3302302310331123-2301120201231323-1120203233013132-2310220201330210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203023131013131-0020012310120103-1110131201033023-3110032323131333-0031002220221211-0332113021023302-3010013231231223-3111222103023321"></a>

## proxy_advertisement.advertise_on_public_default_dualstack_vip — advertise_on_public_default_dualstack_vip / 110333103121 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public_default_dualstack_vip

<a id="canonical-0002131110333021-0323213313130330-0303222033120113-3121123012310001-0310022123203211-3132333211233222-0021222122221202-0311333213020313"></a>

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

<a id="canonical-1211002233323211-2030133201132113-3200130220020313-2111001210332203-3332303003023201-3101003221123102-0130103202302230-3130002020332003"></a>

## Direct properties — advertise_on_public_default_dualstack_vip / 110333103121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032200330310203-1222132331033000-1301130311201213-0110323230330303-2101010321122030-1123210332111111-1020010102230012-0332010212221312"></a>

## Next pages — advertise_on_public_default_dualstack_vip / 110333103121 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0010111020311300-1202110002311312-3111100031020231-3233312101103231-2321101300233120-2111220013030021-0203011312322031-2032301213311120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003332210102303-3132321320202122-0311231031100323-3013321001320220-1100012330101213-3211121122123333-0220001121110321-2103300222133202"></a>

## proxy_advertisement.advertise_on_public_default_ipv6_vip — advertise_on_public_default_ipv6_vip / 322033100200 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public_default_ipv6_vip

<a id="canonical-2100231102122202-1003100230223330-0201110101011120-0310333333020133-0210311033230211-0033222310022021-0001122310313100-0202211021231331"></a>

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

<a id="canonical-2310301031222013-3001211210133022-3032211202021123-3223322123203103-1323222003310013-1132301200331132-0020231300322213-0102013101132321"></a>

## Direct properties — advertise_on_public_default_ipv6_vip / 322033100200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312303123200320-2013010113203000-1120121333133220-3201123032303202-1213031200003103-3210330302132222-1323011020033033-3312133202223312"></a>

## Next pages — advertise_on_public_default_ipv6_vip / 322033100200 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1022322020332313-1201203221223320-0201230131212313-1100002123112013-2313010311202123-1200121102113312-1212301113321232-0123230022132320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331100000010122-0023332022300102-2222333231232331-0131032000101320-0013303300001302-1233300130102011-1233213202131321-1300210230000201"></a>

## proxy_advertisement.advertise_on_public_default_vip — advertise_on_public_default_vip / 001002033023 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_on_public_default_vip

<a id="canonical-0302120000223013-1311100103032323-1130112221333032-2022213100113122-3233203233211300-2000133011000311-1310023331233323-2121130232011303"></a>

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

<a id="canonical-1133030031322020-0233332320001030-3321131031122111-2222130233202312-1133133222313303-3000132222220210-2103120110321203-0230020123323013"></a>

## Direct properties — advertise_on_public_default_vip / 001002033023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230211320220010-3003133210132102-1010221100212333-0222101320111221-2231012220130321-3333123122113020-1210010122121122-3100201311110303"></a>

## Next pages — advertise_on_public_default_vip / 001002033023 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303332002033131-0112310103013222-3020123330203331-2011212311120211-2132100313021321-3312333133233332-1221331333113032-3111022033221213"></a>

## proxy_advertisement.advertise_v6_on_public — advertise_v6_on_public / 103103112022 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.advertise_v6_on_public

<a id="canonical-1101300330101121-0301312132001130-1100120010001300-3010321212103232-1010000210322330-2323121332332221-0032001133121212-2233330130210210"></a>

Type: `"single"`. Computed.

Defines a way to advertise a load balancer on public. If optional public\_ip is provided, it will
only be advertised on RE sites where that public\_ip is available.

Upstream description:

This defines a way to advertise a load balancer on public. If optional public\_ip is provided, it
will only be advertised on RE sites where that public\_ip is available.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2212322200101121-1210202232211133-2000120133013110-1321333130211102-0120330302103303-1230203230211122-0311132133002132-3002020231210021"></a>

## Direct properties — advertise_v6_on_public / 103103112022 / 3

- [public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0000322003033311-1021323203112330-0200122313332003-3020332323121223-3311020323221022-3203121320010213-2323130323121130-0230011313330231): complete subsection reference.

<a id="canonical-0320322033202313-1103232123020020-2001222002101130-1130312232233132-3210231010130001-3032032211302021-0013033003102110-0101320313301212"></a>

## Next pages — advertise_v6_on_public / 103103112022 / 4

- [proxy_advertisement.advertise_v6_on_public.public_ip](data-sources--dns_proxy--reference--group-002.md#canonical-0000322003033311-1021323203112330-0200122313332003-3020332323121223-3311020323221022-3203121320010213-2323130323121130-0230011313330231)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-0000322003033311-1021323203112330-0200122313332003-3020332323121223-3311020323221022-3203121320010213-2323130323121130-0230011313330231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221233331210222-2020021113202033-2333230003203203-2110012230031001-2221232022033221-1322210023210010-3302131030121130-3333331123333110"></a>

## proxy_advertisement.advertise_v6_on_public.public_ip — public_ip / 021211101232 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213)
- proxy_advertisement.advertise_v6_on_public.public_ip

<a id="canonical-3100220323023013-2123101123301222-3300213013002132-2211313010132301-0100201000131223-0332010232210323-3031111311321013-0110011322011300"></a>

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

<a id="canonical-0311211210131023-3210202001313130-1000113003323112-2001000001031302-1001203011313321-1301002231102213-1121012020000323-1113101123233101"></a>

## Direct properties — public_ip / 021211101232 / 3

<a id="canonical-0031200223103212-2302232323030101-2332323010002310-0323212023001321-1321231202110001-1133313121312031-3331022002321302-2001033210120233"></a>

<a id="canonical-0223000010031131-3013102200030023-1100122212201110-3232131130010022-3312222332031121-0012331133333031-3130322212000301-0121130102011123"></a>

## name property — public_ip / 021211101232 / 4

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

<a id="canonical-3323103212220131-1332012031332222-2001313132200200-0300213113223102-3323201103321333-1320112020033002-1020212203333020-2013133310021331"></a>

<a id="canonical-0022323200011113-2110332120223211-1033132302131200-3003120213000212-0030300211030302-2002222331113210-1212232202001311-1303233211130222"></a>

## namespace property — public_ip / 021211101232 / 5

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

<a id="canonical-2201022200321133-0211213110211113-1230210323023132-2302223133020011-0233112121221232-3220120032022220-0000211013202012-2313232110222222"></a>

<a id="canonical-3320001310320120-2011003001023321-0322303112300010-3123020012320313-1102012103201220-0222101220222103-2200102031002022-3112202223333010"></a>

## tenant property — public_ip / 021211101232 / 6

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

<a id="canonical-3331330123210320-1223101331332110-0333210130231210-1012002313032321-2220102230120303-0000111120222330-2033121301210121-3122120013022111"></a>

## Next pages — public_ip / 021211101232 / 7

- [proxy_advertisement.advertise_v6_on_public](data-sources--dns_proxy--reference--group-002.md#canonical-3023010033002312-1211102200100233-0013330212203322-0033130332302222-1020020000213233-0021211213221212-1221003001332210-3122311210303213)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)

<a id="canonical-1012133330232223-3320102120220023-3002110212112031-0330112322100132-3202033312322102-3301100130033303-3200110031011233-3223231020123322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013301001122331-1222023032030030-3002132321012132-2202113100122111-0103133213333032-0011221101023332-0032133303333233-2321103112222232"></a>

## proxy_advertisement.do_not_advertise — do_not_advertise / 002320222100 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
- [Property reference](data-sources--dns_proxy--reference--group-001.md#canonical-1133330113022100-3110333211021121-3111232100100000-3311232321023320-3202221331130133-0013122201302330-1103000312200021-1203111112100020)
- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- proxy_advertisement.do_not_advertise

<a id="canonical-1200102101022133-0203120300303311-2321030000103133-0230002201020300-1210232213020300-2023300023002312-3221333213021123-1111131110311202"></a>

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

<a id="canonical-3320300232223013-0330212201313322-1020013322020002-1113033230132211-3000121111002101-1233322302133231-0201030022010131-0313001123201201"></a>

## Direct properties — do_not_advertise / 002320222100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100302111023132-2122103031003121-1100310332132112-1022022032332023-3202233211121103-3010313301213213-2202311002300023-3110021322313201"></a>

## Next pages — do_not_advertise / 002320222100 / 4

- [proxy_advertisement](data-sources--dns_proxy--reference--group-001.md#canonical-1233321320320303-1032110303233312-2131113231101333-3032323203221001-0110213212120103-3322121223310111-0201121112313233-2310233121220133)
- [xcsh_dns_proxy](../data-sources/dns_proxy.md#canonical-1122013202133033-3201201122310010-3303231012021212-2123120003213300-3130001133310203-2011013122222333-0130313333122120-0323311233013112)
