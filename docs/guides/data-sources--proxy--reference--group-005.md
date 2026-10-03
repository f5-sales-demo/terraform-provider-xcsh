---
page_title: "xcsh_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_proxy reference."
---

# xcsh_proxy reference

<a id="canonical-0303302302300102-0233203230000300-3320122303010110-3222301123301320-3001210002013131-1202210233023202-3123031033322033-0320121203201310"></a>

## network property — virtual_site / 031331011302 / 4

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

- [virtual_site](data-sources--proxy--reference--group-005.md#canonical-1302033331002032-1233232303323222-2033003231120320-2232201001323013-0321020302101332-0221333200123101-1123201021133203-1022302201220101): complete subsection reference.

<a id="canonical-1311312102133100-3210223300120131-3120332302211113-2120333132110302-3020232012121112-1112120201023011-3010212013131210-1030001000121010"></a>

## Next pages — virtual_site / 031331011302 / 5

- [site_virtual_sites.advertise_where.virtual_site.virtual_site](data-sources--proxy--reference--group-005.md#canonical-1302033331002032-1233232303323222-2033003231120320-2232201001323013-0321020302101332-0221333200123101-1123201021133203-1022302201220101)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1302033331002032-1233232303323222-2033003231120320-2232201001323013-0321020302101332-0221333200123101-1123201021133203-1022302201220101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2320120130322313-3313020223222002-3003132330011111-3210320222220230-3102223300101002-3100200301321123-2100112033310323-1111023233313020"></a>

## site_virtual_sites.advertise_where.virtual_site.virtual_site — virtual_site / 331013123132 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [site_virtual_sites](data-sources--proxy--reference--group-004.md#canonical-1012201322333030-2201110212121013-2012102101130230-0200303100213333-0202312022013300-2020101021021321-1320201313130100-3300013202221330)
- [site_virtual_sites.advertise_where](data-sources--proxy--reference--group-004.md#canonical-1101311121310312-0200013102222131-0303021112013023-0200203203003231-1221230311001231-0230200233121202-1303302123221200-0111202023011311)
- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-004.md#canonical-0011000332332130-3130012231211212-3202112322222113-3111230132210223-2100313233112001-3112212110232222-2031021001200031-3310310312213001)
- site_virtual_sites.advertise_where.virtual_site.virtual_site

<a id="canonical-3201133000030101-2111222220311312-1223130313233220-0003221131111102-3001333133233332-1111313322322230-2032222020330202-3202301320120331"></a>

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

<a id="canonical-2313330221233120-3003310302211031-3213103300112033-0313301120201031-3230031212002111-3033021322122122-3111213013101110-2322013332021323"></a>

## Direct properties — virtual_site / 331013123132 / 3

<a id="canonical-1201303031032001-1001122221023331-3320003123203113-1132120311222212-0322333312112103-0100123030300021-1300222133300112-3030230332331030"></a>

<a id="canonical-0013212103331330-2000210222102223-0231022110323300-2213323312100102-2212033313012311-2103303010122302-0213301312313002-0211320313320203"></a>

## name property — virtual_site / 331013123132 / 4

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

<a id="canonical-2212033210031212-3310110313112022-3023000013120231-2133003010321322-0311122100133321-3321322031031220-2011330021220100-0011310223203122"></a>

<a id="canonical-1031001323310130-2033313223133102-1332012002332231-1310032332201320-1132123203101232-0130220303111012-0322313212331021-2032110102233010"></a>

## namespace property — virtual_site / 331013123132 / 5

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

<a id="canonical-0332021221300200-0222232011123103-0233323231032102-1212312113123021-1233032122322330-0332003210200322-1333012322002222-3032021030131001"></a>

<a id="canonical-1030001221331112-3313320323203031-3332232223013323-1212301101323110-0201113011202221-0011332222330321-3013122013112131-3333302321103210"></a>

## tenant property — virtual_site / 331013123132 / 6

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

<a id="canonical-3120031033002113-0023121112203132-1231000020000302-2321012023332322-0323001121302322-0112133232223012-0321033133030000-0131103001311112"></a>

## Next pages — virtual_site / 331013123132 / 7

- [site_virtual_sites.advertise_where.virtual_site](data-sources--proxy--reference--group-004.md#canonical-0011000332332130-3130012231211212-3202112322222113-3111230132210223-2100313233112001-3112212110232222-2031021001200031-3310310312213001)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031321233011130-1322000233203311-1102003212023022-3101112132211020-3222222311032110-3200003332220300-2120223320032310-0331222311332311"></a>

## tls_intercept — tls_intercept / 213210331130 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- tls_intercept

<a id="canonical-2121232312132030-1021002130200310-0121113113303233-3010003311301312-3031223213223011-2223320220020023-0033302022021110-1101201020202012"></a>

Type: `"single"`. Computed.

Configuration to enable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-interception_policy_choice": "[\"enable_for_all_domains\",\"policy\"]",
  "x-ves-oneof-field-signing_cert_choice": "[\"custom_certificate\",\"volterra_certificate\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca_url\",\"volterra_trusted_ca\"]"
}
```

<a id="canonical-3312032223302110-0030323133121132-2010302101022133-0121231220321233-2133132203203212-0212303130000210-2123300211000101-2113303013311030"></a>

## Direct properties — tls_intercept / 213210331130 / 3

- [custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220): complete subsection reference.

- [enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-3230230222011313-3333232211102332-1213233310022012-2200313222323113-0332021201200103-0202133211030313-3132201012212010-1321323130131223): complete subsection reference.

- [policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100): complete subsection reference.

<a id="canonical-0102133020010102-2213012022123223-3310030013031013-0203321130301312-2202121010022122-3010200322231312-1221222212020132-2033120022201111"></a>

<a id="canonical-0220213002310002-1212133201023321-0310200222211003-1020330300202313-3122222033230213-1103310001210233-1333001223012212-2322102113121010"></a>

## trusted_ca_url property — tls_intercept / 213210331130 / 4

Type: `"string"`. Computed.

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

Upstream description:

Exclusive with \[volterra\_trusted\_ca\] Custom Root CA Certificate for validating upstream server
certificate.

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
    "maxLength": 131072,
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

- [volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-3020010120123000-0301123030210101-0321130032221323-2123030030303222-2003011230132203-2230223213222003-3100213222023330-0120202213313121): complete subsection reference.

- [volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-0200220011202212-2321233023123120-1202023310102102-0222131001130112-0320121130100111-2103030131332202-1120332223231302-0301113013200231): complete subsection reference.

<a id="canonical-0011003311332000-0012233213222220-2222300210110001-0113221202100133-1013031021102011-0111111330112031-2101323101220113-3030302132213330"></a>

## Next pages — tls_intercept / 213210331130 / 5

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [tls_intercept.enable_for_all_domains](data-sources--proxy--reference--group-005.md#canonical-3230230222011313-3333232211102332-1213233310022012-2200313222323113-0332021201200103-0202133211030313-3132201012212010-1321323130131223)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.volterra_certificate](data-sources--proxy--reference--group-005.md#canonical-3020010120123000-0301123030210101-0321130032221323-2123030030303222-2003011230132203-2230223213222003-3100213222023330-0120202213313121)
- [tls_intercept.volterra_trusted_ca](data-sources--proxy--reference--group-005.md#canonical-0200220011202212-2321233023123120-1202023310102102-0222131001130112-0320121130100111-2103030131332202-1120332223231302-0301113013200231)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311322213113333-1000213013332332-3312010221123112-0031311330022123-1223021103213231-1013132232003220-0301123202001221-3220233011113111"></a>

## tls_intercept.custom_certificate — custom_certificate / 113312031130 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.custom_certificate

<a id="canonical-3232201211233031-3302012031311123-3001022011213013-3030010132310122-1012011300203110-2120201011302001-2312333002221232-2013321321332302"></a>

Type: `"single"`. Computed.

Configuration parameter for custom certificate.

Upstream description:

Handle to fetch certificate and key.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ocsp_stapling_choice": "[\"custom_hash_algorithms\",\"disable_ocsp_stapling\",\"use_system_defaults\"]"
}
```

<a id="canonical-1331202331313130-3000230123312102-0211323211033223-3031010110032211-3223232220311123-3003121010102131-0022221112333333-0123320130133033"></a>

## Direct properties — custom_certificate / 113312031130 / 3

<a id="canonical-3032113001001113-0020202300333322-2021210112230303-0303003102320003-0302033300021212-0320103321201300-0232333012323230-3310330020032210"></a>

<a id="canonical-2031113220232022-1213220102020322-3232323223211223-1302230210132002-3131302133230111-2302200301220332-1132322321010333-1233032301030022"></a>

## certificate_url property — custom_certificate / 113312031130 / 4

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

- [custom_hash_algorithms](data-sources--proxy--reference--group-005.md#canonical-1023213030332201-2312100033010103-1010203103321123-0111001103030100-0132333110030230-3313200112220113-1202302003003233-1021310321232310): complete subsection reference.

<a id="canonical-2020220001230310-1310333302301032-0112320213131013-2131311131001203-1233030201122032-0221031232311101-2331332332010021-2022033103331101"></a>

<a id="canonical-3310023012113033-3002121002311231-0013123111221030-0311211103301213-2312313233023333-2231121011223022-0112122000201130-2313333212121010"></a>

## description_spec property — custom_certificate / 113312031130 / 5

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--proxy--reference--group-005.md#canonical-1030233231323122-3203323120201203-2300102221222310-3100102220230001-1033122103110300-0110200213132320-3133201033332133-1023222210303130): complete subsection reference.

- [private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013): complete subsection reference.

- [use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-1302232200212121-3021213100303023-2102012001133212-3021120313030310-3122221011012111-2302202111200013-3232133330311313-3213203112122232): complete subsection reference.

<a id="canonical-1332220301321212-1230002111303201-3200121021010300-3301021003320220-0031102013000111-0100011011131232-1131320320031030-3033020302213111"></a>

## Next pages — custom_certificate / 113312031130 / 6

- [tls_intercept.custom_certificate.custom_hash_algorithms](data-sources--proxy--reference--group-005.md#canonical-1023213030332201-2312100033010103-1010203103321123-0111001103030100-0132333110030230-3313200112220113-1202302003003233-1021310321232310)
- [tls_intercept.custom_certificate.disable_ocsp_stapling](data-sources--proxy--reference--group-005.md#canonical-1030233231323122-3203323120201203-2300102221222310-3100102220230001-1033122103110300-0110200213132320-3133201033332133-1023222210303130)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- [tls_intercept.custom_certificate.use_system_defaults](data-sources--proxy--reference--group-005.md#canonical-1302232200212121-3021213100303023-2102012001133212-3021120313030310-3122221011012111-2302202111200013-3232133330311313-3213203112122232)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1023213030332201-2312100033010103-1010203103321123-0111001103030100-0132333110030230-3313200112220113-1202302003003233-1021310321232310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1302030002110213-3303122330002000-1221321232230312-3111232232311313-3100322132310000-2333131320311213-3003000122022101-3223030331211111"></a>

## tls_intercept.custom_certificate.custom_hash_algorithms — custom_hash_algorithms / 032023101210 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.custom_hash_algorithms

<a id="canonical-0130300112022033-0030312103101302-2022332303012121-2223121110330330-1012220323322231-3210232222201113-0223302013033130-0122303212333021"></a>

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

<a id="canonical-3012112102332330-1301021313110231-1323221012223001-0233321333200102-1231130231302130-2232122120011230-2033323002320012-0112322020312123"></a>

## Direct properties — custom_hash_algorithms / 032023101210 / 3

<a id="canonical-3200323301112120-3101303120100111-2021213132221023-3212121111103233-1022010200113323-2130201330103311-2202010233201012-2033110311302111"></a>

<a id="canonical-3223212120201002-3202123010321020-3300332211000131-3113013313320133-2312102200331213-1300101332021201-3303213313232020-3031200221132130"></a>

## hash_algorithms property — custom_hash_algorithms / 032023101210 / 4

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

<a id="canonical-2321002202000011-2212222331221223-0232113332121002-2100301123312133-3013233200221022-3013233120331023-2331201313212312-2111111012202301"></a>

## Next pages — custom_hash_algorithms / 032023101210 / 5

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1030233231323122-3203323120201203-2300102221222310-3100102220230001-1033122103110300-0110200213132320-3133201033332133-1023222210303130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222310221301220-1102110123130320-0223322312011201-3130103001032111-3301303332112321-3001122102131321-2221222332222302-0222213022120213"></a>

## tls_intercept.custom_certificate.disable_ocsp_stapling — disable_ocsp_stapling / 202323023002 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.disable_ocsp_stapling

<a id="canonical-3201322323312302-3101312131312220-2022201233311300-0021122010211331-3201230333320103-1221301010202003-1100011032121030-3221131332121201"></a>

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

<a id="canonical-1321313321011231-1210020212011121-3320100130033103-3001020031000030-3311211010112221-3211302111220003-0033030203211020-2003031211311032"></a>

## Direct properties — disable_ocsp_stapling / 202323023002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131000120033032-2302013023113121-1031321200221122-2011323100021332-1301211202222000-2223321201101310-3203002223203321-1032231030023022"></a>

## Next pages — disable_ocsp_stapling / 202323023002 / 4

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223223323231300-1022213023200011-3311013113301311-3200111312102133-2331111211320231-1002313201032302-0302001230211112-3222232322022100"></a>

## tls_intercept.custom_certificate.private_key — private_key / 303300302322 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.private_key

<a id="canonical-1321110132331221-3323111110320330-3201123022231032-1020320300200320-0200110113213103-1200221012112103-0011002102322123-2311311312011132"></a>

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

<a id="canonical-1300132012003311-2101313030001023-0101110200003121-2120303331300322-3301213321232112-3321021230122103-2313202200103222-2231231311313223"></a>

## Direct properties — private_key / 303300302322 / 3

- [blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-0322103332321130-1213301201221100-2100011323312323-0130121211021000-0130121132113332-2011103223230322-0000212002313013-3212022300212232): complete subsection reference.

- [clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-1322212321230133-2133303132322001-3330322212211033-1232322231000102-3231132220100133-0132221001212132-3321011302213132-2200131333232202): complete subsection reference.

<a id="canonical-3202123323030012-1320220211101303-1133331220333231-0300312212031222-2200300213032312-3100110330220031-3000222023213233-2302210100232200"></a>

## Next pages — private_key / 303300302322 / 4

- [tls_intercept.custom_certificate.private_key.blindfold_secret_info](data-sources--proxy--reference--group-005.md#canonical-0322103332321130-1213301201221100-2100011323312323-0130121211021000-0130121132113332-2011103223230322-0000212002313013-3212022300212232)
- [tls_intercept.custom_certificate.private_key.clear_secret_info](data-sources--proxy--reference--group-005.md#canonical-1322212321230133-2133303132322001-3330322212211033-1232322231000102-3231132220100133-0132221001212132-3321011302213132-2200131333232202)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0322103332321130-1213301201221100-2100011323312323-0130121211021000-0130121132113332-2011103223230322-0000212002313013-3212022300212232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112032113003021-0121213211333022-2331213322200003-2023100130310331-1220300022201101-0002120032031311-2013223023102211-1100133313302320"></a>

## tls_intercept.custom_certificate.private_key.blindfold_secret_info — blindfold_secret_info / 001031111322 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- tls_intercept.custom_certificate.private_key.blindfold_secret_info

<a id="canonical-2130120001211200-2023110233223223-0222210032032110-1233200000232212-2311301002222312-0120313103023031-1000020000022222-1123030132123131"></a>

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

<a id="canonical-0000012120210132-0300323012032101-2130020301020223-0323311123323203-3130120233231300-3103100323120220-0131000331231020-1320031022230022"></a>

## Direct properties — blindfold_secret_info / 001031111322 / 3

<a id="canonical-3133020322322213-2333332200311132-0130310312200321-2122211023321002-1202113031210112-0203031120110211-3301223211131021-1310203132122120"></a>

<a id="canonical-3132100321211030-2320312211102212-2222300230201302-1033211031202023-2002123300003221-0213233000010033-2101303303101113-3213132210320113"></a>

## decryption_provider property — blindfold_secret_info / 001031111322 / 4

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

<a id="canonical-3300030203302003-3021022002320200-2212330120322031-3023010201301310-2123012002331131-2133310031211122-2231200320122121-1002313301032322"></a>

<a id="canonical-0101212122221001-0211122201002202-0132121310310020-2110232221330101-2133233031130201-2121200320332030-0323311203230010-3001212210111020"></a>

## location property — blindfold_secret_info / 001031111322 / 5

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

<a id="canonical-3323231132013021-1201331320010211-0210103100312033-3011001303330321-0132133201011233-1303003000210203-2301001312231223-2130133020133313"></a>

<a id="canonical-2201033122312020-2321312232231010-3120311333322311-1113322330301000-2200320120011321-3221013033132121-3003223311131030-2212300113200130"></a>

## store_provider property — blindfold_secret_info / 001031111322 / 6

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

<a id="canonical-2223321021000002-3033020123033322-2223210013323232-1302131011221012-2312200121331120-2213232120220311-1131131113230130-3023003323122311"></a>

## Next pages — blindfold_secret_info / 001031111322 / 7

- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1322212321230133-2133303132322001-3330322212211033-1232322231000102-3231132220100133-0132221001212132-3321011302213132-2200131333232202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1002011220031132-1330133330213310-3003310230112031-0020230331031101-2232231112031223-2223311111031231-3012113111330033-3012223030332131"></a>

## tls_intercept.custom_certificate.private_key.clear_secret_info — clear_secret_info / 322003213011 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- tls_intercept.custom_certificate.private_key.clear_secret_info

<a id="canonical-1220233030332323-0010003022111130-1221031230130102-3021122123302321-2221222223133323-0133001022232301-2021221101210130-2333211301320002"></a>

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

<a id="canonical-0332310322200131-3211033320133233-0321232010203111-2332123132000123-2111201111210213-1100330332021110-0312022133121303-2131033200013200"></a>

## Direct properties — clear_secret_info / 322003213011 / 3

<a id="canonical-2023012032300020-3130130122203230-0200111221020000-3021113231330001-2233110010301332-0203331020321031-3300300113232133-1203013223223222"></a>

<a id="canonical-0323222133312012-1332330332113311-0011313212233010-1323012310131312-1303100213320021-3203131302013023-2022001331212000-2113220303122213"></a>

## provider_ref property — clear_secret_info / 322003213011 / 4

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3003332223312231-3301323201013300-1201200303001212-0030101103331331-1302011320113303-1203213132132320-2223023212111003-1321313220320020"></a>

<a id="canonical-1101200132222220-3332000002000100-0012223013202321-2313213300311201-3002023113022132-2130033020030121-1301021031122022-1122000310233200"></a>

## URL property — clear_secret_info / 322003213011 / 5

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

<a id="canonical-2210032032021320-0031021120112131-2003201210303110-2033303131123203-3202032311310202-2122231221110221-1021211023310310-2211223003302131"></a>

## Next pages — clear_secret_info / 322003213011 / 6

- [tls_intercept.custom_certificate.private_key](data-sources--proxy--reference--group-005.md#canonical-3110120000322320-0101001302313122-1331202112230322-2202301331030002-0023312212222101-3110002202232021-3023303132133330-0232333301023013)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1302232200212121-3021213100303023-2102012001133212-3021120313030310-3122221011012111-2302202111200013-3232133330311313-3213203112122232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211000102322303-0201200022123112-3103111332333300-3131000130302203-3212111002112012-3131011022130330-1100300322323303-3233003013130330"></a>

## tls_intercept.custom_certificate.use_system_defaults — use_system_defaults / 122031131132 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- tls_intercept.custom_certificate.use_system_defaults

<a id="canonical-3210101010210211-0003110102111231-0313232033230131-1320322121023121-0211233123211003-0103112203310321-3320111020221230-0113222113120201"></a>

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

<a id="canonical-0131313311010120-0010030103321102-0102010313021020-1300012302021110-3132331312200213-0300022203210130-2132000001322130-1323030211323203"></a>

## Direct properties — use_system_defaults / 122031131132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0212323121303031-1200300000303133-2102302100011023-1103102322121132-1312122302301302-3131322121133330-2330210203302221-0311113302213002"></a>

## Next pages — use_system_defaults / 122031131132 / 4

- [tls_intercept.custom_certificate](data-sources--proxy--reference--group-005.md#canonical-1231123100111031-1220000311120232-0330321000211131-1210012123123012-3321122100011131-2031002021100303-2020020312202221-0000213111300220)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3230230222011313-3333232211102332-1213233310022012-2200313222323113-0332021201200103-0202133211030313-3132201012212010-1321323130131223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212302002023220-0131220123121003-1333121210103023-3130200100103312-2033001310100232-2300200301002202-0113210120231031-2303201131122232"></a>

## tls_intercept.enable_for_all_domains — enable_for_all_domains / 011123002311 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.enable_for_all_domains

<a id="canonical-1330022123231330-2322122032123301-0312121302133213-0322313230020103-0033201221231313-1133301231300323-0012010313013202-1302211022331311"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable for all domains.

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

<a id="canonical-1112013122030131-3013310011311300-3120201332211020-3233003313102102-3030323331132322-3312000000013200-1030203320333111-2112031220111030"></a>

## Direct properties — enable_for_all_domains / 011123002311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003210231002103-2200130020311130-1211301001301033-2313223132101023-1011302120030133-3031011311213002-3020332302001221-2001311023001233"></a>

## Next pages — enable_for_all_domains / 011123002311 / 4

- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001313200102311-0021032132311103-0302213323111211-2302001011122300-1232000303323223-1003200301331233-3202200313232200-1003211201300222"></a>

## tls_intercept.policy — policy / 031110323211 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.policy

<a id="canonical-2101031213001101-1213003123113223-1320233222033313-3033132022011020-1230122203212100-3230122133133232-3210213110311030-2201310011021210"></a>

Type: `"single"`. Computed.

Policy to enable or disable TLS interception.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2123030233003102-2131333302222121-3332212213010000-3200201001102313-3100210322200223-3033001230103122-3123201333302020-1012311021333220"></a>

## Direct properties — policy / 031110323211 / 3

- [interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112): complete subsection reference.

<a id="canonical-3021333210110203-1221303302010130-0132020122320023-1213222120322301-3311323132130200-3021333023101331-2230103220101330-3011310101213211"></a>

## Next pages — policy / 031110323211 / 4

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220113111232110-3222131202023203-1010020313221212-2030230101113133-1300032233210332-3010201013230232-0202302211221110-1330213332110102"></a>

## tls_intercept.policy.interception_rules — interception_rules / 032110220000 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- tls_intercept.policy.interception_rules

<a id="canonical-1113032002213112-3322213023132212-3103021213030223-0312300312212233-0303223123321130-3231332133031300-3212332320202121-1301323320211021"></a>

Type: `"list"`. Computed.

List of ordered rules to enable or disable for TLS interception.

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

<a id="canonical-3023321300202321-3333230322313200-0232333213131102-1202223233330123-0113220110133210-0221032213112023-2223302230031003-0323003200311021"></a>

## Direct properties — interception_rules / 032110220000 / 3

- [disable_interception](data-sources--proxy--reference--group-005.md#canonical-2120002331023310-2300002310011311-3023231120133323-3123331131120200-1110110202022323-3020312210233102-1302301033100013-2113011320103113): complete subsection reference.

- [domain_match](data-sources--proxy--reference--group-005.md#canonical-0301211320300332-0131231113212310-0300320322212133-2233333301210212-0130231321210120-1221333032210230-0012022300222310-0303111321201312): complete subsection reference.

- [enable_interception](data-sources--proxy--reference--group-005.md#canonical-1023223231130220-3312032212030130-3011110301100311-0220302013032313-0001011202133010-2031233331000130-3221121321223012-3030210003033123): complete subsection reference.

<a id="canonical-1102201112131011-0310103023030112-3001102230100223-1023000123133133-1310320323220120-2123030213030211-1032023202312323-2220220210013203"></a>

## Next pages — interception_rules / 032110220000 / 4

- [tls_intercept.policy.interception_rules.disable_interception](data-sources--proxy--reference--group-005.md#canonical-2120002331023310-2300002310011311-3023231120133323-3123331131120200-1110110202022323-3020312210233102-1302301033100013-2113011320103113)
- [tls_intercept.policy.interception_rules.domain_match](data-sources--proxy--reference--group-005.md#canonical-0301211320300332-0131231113212310-0300320322212133-2233333301210212-0130231321210120-1221333032210230-0012022300222310-0303111321201312)
- [tls_intercept.policy.interception_rules.enable_interception](data-sources--proxy--reference--group-005.md#canonical-1023223231130220-3312032212030130-3011110301100311-0220302013032313-0001011202133010-2031233331000130-3221121321223012-3030210003033123)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-2120002331023310-2300002310011311-3023231120133323-3123331131120200-1110110202022323-3020312210233102-1302301033100013-2113011320103113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030333201222020-1123122322003120-0011330232000233-3022013013033010-0011332102310013-0010233210002010-3323231100230211-3101303332102211"></a>

## tls_intercept.policy.interception_rules.disable_interception — disable_interception / 213210310320 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- tls_intercept.policy.interception_rules.disable_interception

<a id="canonical-3230310311131200-2202011103231103-1133330322213212-3011112320010120-2313232303213000-1102322223010202-1222012233231310-0213032033000332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable interception.

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

<a id="canonical-1330001313020032-0223121321020313-1021002003100230-3130033022222133-1103003011333121-2330001021311020-2200233300203221-3211333000020102"></a>

## Direct properties — disable_interception / 213210310320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123212221013311-3320230131333230-2100101320233222-0213311210201211-1101021320220123-0002011103302022-0030201030102011-2232300331022311"></a>

## Next pages — disable_interception / 213210310320 / 4

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0301211320300332-0131231113212310-0300320322212133-2233333301210212-0130231321210120-1221333032210230-0012022300222310-0303111321201312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101111313030222-2030202103111210-2032223333000110-1221213112113303-1013023200100132-0221033331021003-2121213200301212-1221320203001011"></a>

## tls_intercept.policy.interception_rules.domain_match — domain_match / 321021231212 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- tls_intercept.policy.interception_rules.domain_match

<a id="canonical-2000211212020220-0323011121211120-3022030001001331-2333103121200000-1100332222323320-2303213003233321-3100201022110311-0132232200311121"></a>

Type: `"single"`. Computed.

Configuration parameter for domain match.

Upstream description:

Domains names.

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

<a id="canonical-2232011233033313-3123020323021313-3210122201033330-0210011012120332-2111331002100221-3030330110032113-0220113103020100-2131001003303322"></a>

## Direct properties — domain_match / 321021231212 / 3

<a id="canonical-3121132113013333-2011011013232313-3020113110312310-0012302303230203-0302013022233011-3210033020331303-2102200101103230-1203301202033122"></a>

<a id="canonical-2332002113133103-2102133221101123-3300232201322332-2331201311330100-3302221003300121-2020102322230120-1110330321013133-0130011200100012"></a>

## exact_value property — domain_match / 321021231212 / 4

Type: `"string"`. Computed.

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[regular expression\_value suffix\_value\] Exact domain name.

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

<a id="canonical-1210110210120023-1103002303201111-1103113101110233-3132303330011010-0203300113020133-3222233222001221-2230310212311113-2100312032012320"></a>

<a id="canonical-2333321301110101-0220221213210202-0221011330332223-1023001113333120-0312211210130032-0213100133200033-3321102022201121-3102013011223231"></a>

## regex_value property — domain_match / 321021231212 / 5

Type: `"string"`. Computed.

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

Upstream description:

Exclusive with \[exact\_value suffix\_value\] Regular Expression value for the domain name.

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

<a id="canonical-1010202121322323-3021203112223202-2113023231112103-2020231202201303-2102203201030300-1013032132130231-1300331012111300-2303021121300211"></a>

<a id="canonical-1020131003022333-2020201120221003-3310212120201021-1013032120032302-2231012012011301-3202131103323033-3110233133310200-2012212001002203"></a>

## suffix_value property — domain_match / 321021231212 / 6

Type: `"string"`. Computed.

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[exact\_value regular expression\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-1023032122332111-3322220133003120-2002131003201230-0021010220111101-2032113223001220-3003222021112230-3023302000212020-0323113202000221"></a>

## Next pages — domain_match / 321021231212 / 7

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-1023223231130220-3312032212030130-3011110301100311-0220302013032313-0001011202133010-2031233331000130-3221121321223012-3030210003033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320323021301313-3330333112221311-1012312231202022-1020301112321321-1012003111020112-3220201101131131-3323122131110120-2333200301322213"></a>

## tls_intercept.policy.interception_rules.enable_interception — enable_interception / 011103111023 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [tls_intercept.policy](data-sources--proxy--reference--group-005.md#canonical-3103321231313310-3112212012231231-2021103203231010-3203131230032202-2311212331112113-2203200112111100-2213330010011221-1023110203010100)
- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- tls_intercept.policy.interception_rules.enable_interception

<a id="canonical-1100020223102300-2322222300101130-2112310310233020-1220231110200301-3020010330133230-1301022323201233-1320122113213301-2310222122330113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable interception.

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

<a id="canonical-3032203212021103-1022332320320310-2113223101333121-0223321130221211-3223221000102123-0133211332222221-0202210331210320-3233211331032103"></a>

## Direct properties — enable_interception / 011103111023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233231303023220-1122303213212320-0013330210122032-3033120122133330-0322223313332131-0133130113021321-3222330223132201-1322012122031001"></a>

## Next pages — enable_interception / 011103111023 / 4

- [tls_intercept.policy.interception_rules](data-sources--proxy--reference--group-005.md#canonical-1023013012011313-3030230203032111-3020023301210312-3201110221202022-1210031331033210-2233033030120113-0233111003200102-0320103010122112)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-3020010120123000-0301123030210101-0321130032221323-2123030030303222-2003011230132203-2230223213222003-3100213222023330-0120202213313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313330001122013-1322302131223311-1302103113202313-2010110003033230-3013122030231233-2021012020020310-3020331032320133-3033202220220201"></a>

## tls_intercept.volterra_certificate — volterra_certificate / 001133200113 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.volterra_certificate

<a id="canonical-3232030101211322-0013110331131102-1232303132112113-0233211103033200-1033002221322313-2220003312130223-3301221323232332-3331032222313221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra certificate.

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

<a id="canonical-0102030313230133-1321301011010002-0312000211112200-0230320302132233-3130013302200203-1120033212003323-1202331311020213-1133131131120103"></a>

## Direct properties — volterra_certificate / 001133200113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121332123303300-0100201213221322-2322203200201313-3121200003100112-1331013211303212-0320120103003211-0022300100232001-3100111333232133"></a>

## Next pages — volterra_certificate / 001133200113 / 4

- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)

<a id="canonical-0200220011202212-2321233023123120-1202023310102102-0222131001130112-0320121130100111-2103030131332202-1120332223231302-0301113013200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311121003033131-0000100213331112-3221133000201200-0101310231203001-1231300010222322-1121121111001013-0213101132310313-1220121311330230"></a>

## tls_intercept.volterra_trusted_ca — volterra_trusted_ca / 032013131032 / 2

Breadcrumbs:

- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
- [Property reference](data-sources--proxy--reference--group-001.md#canonical-2122303110233010-2221322023230131-1322001310220111-0301130320032201-0202331233311132-2000101111011020-3303123003020132-2031333021133303)
- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- tls_intercept.volterra_trusted_ca

<a id="canonical-3310300332010012-0031300302033223-2003332123223232-2201023223220121-2313332013010022-0102102113300211-3023130302230021-3310013030133301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca.

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

<a id="canonical-1102030032132111-0121222202330302-0002202023032110-3133323101311122-0031032120120200-3300203011011330-3033133112111010-2130120200131022"></a>

## Direct properties — volterra_trusted_ca / 032013131032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223032220131223-3032231033020022-2001331220231020-0111030003333233-2211200100301022-0132002220123011-3021320230310111-0221322222222231"></a>

## Next pages — volterra_trusted_ca / 032013131032 / 4

- [tls_intercept](data-sources--proxy--reference--group-005.md#canonical-3230311230022201-3233301113033013-0131313300030111-1331130023003220-0030332220003121-2232222302322321-3210303320132310-2011003233213312)
- [xcsh_proxy](../data-sources/proxy.md#canonical-2122302220103020-0021013113111101-2022122303101203-3201232230333333-3301303113323330-1012000313313332-0200222300230012-1132013130320033)
