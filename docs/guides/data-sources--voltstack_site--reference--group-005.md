---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-3132230200001111-0220021020123331-2020231020103323-0123021101220113-2001303011130313-3001230220002330-3110203202230301-0321333321022211"></a>

## namespace property — tunnel / 212102031102 / 5

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

<a id="canonical-0102213011301103-3133213200202011-2323111212010113-1123200131003312-0111202101122022-3202101020012023-2032202133231030-1022303303230220"></a>

<a id="canonical-1121133230130213-2023231211031121-0020303201023230-1202230100002020-1023312121230133-0332011231021331-0330033333011003-0032221001230120"></a>

## tenant property — tunnel / 212102031102 / 6

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

<a id="canonical-2023032300230102-2211200020013210-0030013320012103-1001033333202110-1202230331001032-1322213201333332-1310322130123320-1222200300030111"></a>

## Next pages — tunnel / 212102031102 / 7

- [custom_network_config.interface_list.interfaces.tunnel_interface](data-sources--voltstack_site--reference--group-004.md#canonical-0021221203301001-0231230123110300-3302201103203212-3011203133010101-0232210102120133-1001311221312311-1010333311201132-3102320330012000)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2002130123201112-1323230122233002-2232111000322023-3031023330332200-1230030123112310-1210031320002122-3232311030303131-1213111120211032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031203011132122-3120031101101022-1201033322311202-1030110333130313-1302103103101333-3101230031131121-3310330330210312-1120111202310113"></a>

## custom_network_config.no_forward_proxy — no_forward_proxy / 200101322333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.no_forward_proxy

<a id="canonical-2323302002313010-0130111023113123-3323311313300300-1212300103331110-2321230301321123-3013201010213312-1232023211120230-3311113310323012"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no forward proxy.

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

<a id="canonical-0002003033110311-1220033213133101-0103333320302322-3132133200310003-1230001130130110-1331113322320132-2313021303012221-0111221023101233"></a>

## Direct properties — no_forward_proxy / 200101322333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3020320032130220-1331022122121031-0113132020011320-0013013123313031-2300012020010332-0102321213130113-1313011221303302-3120101112333321"></a>

## Next pages — no_forward_proxy / 200101322333 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3133132320123111-1320123130213301-2113313200013133-2332102301313031-3003131220210013-0200312221001100-2123121232231203-0022023223133002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011333332000302-0113001030200003-3230332032230001-2332230102233001-3230032213110112-2212220301312002-1312221020102313-0022200320213222"></a>

## custom_network_config.no_global_network — no_global_network / 121322010220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.no_global_network

<a id="canonical-2021120131232310-1020011132012000-2123001133212121-3303321331001321-3132221032220121-1301121122113222-1312330230220313-3303110001103133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no global network.

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

<a id="canonical-2321113121300322-0002130320113033-2231020323302213-3122120310131311-0131020223331021-1012201103231233-3000010033310201-0003000133030010"></a>

## Direct properties — no_global_network / 121322010220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301020112111203-2210101303001322-0312321033021230-1323103330333221-1100211232320311-1012320232211320-0232303130333012-1021033311202313"></a>

## Next pages — no_global_network / 121322010220 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3332203012022302-3300232301331112-1213030100022113-0230021010103310-2002322311123023-0023321010023111-0032030122201311-0223001232101312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303302230130010-2222021310312303-1330130001201333-3333020210333102-0122113032302313-2213010130102132-1231211103223221-3212030302111210"></a>

## custom_network_config.no_network_policy — no_network_policy / 031000203302 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.no_network_policy

<a id="canonical-0113230010210103-2221102100202120-3203202110030003-0333312221130120-0210322211202312-0111332210003223-1211023012300223-3312033022022113"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

<a id="canonical-0321013031212131-2221330012100111-1032331211033103-2211001002103122-0013002122011211-3110111002230011-0321123131132202-3133311230111102"></a>

## Direct properties — no_network_policy / 031000203302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001202221311331-0002300012031030-2202301233233132-0102103300121101-0311003303002302-3212001120023332-2032322103300113-0203110020203110"></a>

## Next pages — no_network_policy / 031000203302 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112212000202131-3230000123102201-0323111210322233-3112110312023322-1210230010323330-0333232023313301-2231000133333221-1011111132022213"></a>

## custom_network_config.sli_config — sli_config / 211000023123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.sli_config

<a id="canonical-0000302013130120-2323312302320211-1200313232310111-0201033203001231-1223020331300200-3100313301210313-1221312301131130-2000011000312102"></a>

Type: `"single"`. Computed.

Site local inside network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-0030322113300330-1321202333233132-3013322330231121-0102323200011102-0310321011331131-0001022101313322-3012032103322033-3133321222302311"></a>

## Direct properties — sli_config / 211000023123 / 3

- [no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0130320100213312-0032323012111120-0313033332121011-0032300111102301-1321233033220100-0033003330133213-1222301330330112-3320130121310220): complete subsection reference.

- [no_v6_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1221101200233321-1330100322232002-3122213320131310-3122103201330330-3203123311203330-0111330331233310-0011101200110000-0113012301112021): complete subsection reference.

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032): complete subsection reference.

- [static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021): complete subsection reference.

<a id="canonical-1023010320330222-2023103220002031-2110311201230332-3332031200221322-1122012130002113-3002331103322203-0013011102132220-1032201311323233"></a>

## Next pages — sli_config / 211000023123 / 4

- [custom_network_config.sli_config.no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0130320100213312-0032323012111120-0313033332121011-0032300111102301-1321233033220100-0033003330133213-1222301330330112-3320130121310220)
- [custom_network_config.sli_config.no_v6_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1221101200233321-1330100322232002-3122213320131310-3122103201330330-3203123311203330-0111330331233310-0011101200110000-0113012301112021)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0130320100213312-0032323012111120-0313033332121011-0032300111102301-1321233033220100-0033003330133213-1222301330330112-3320130121310220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2211123201033311-3300012323001103-1303032102322001-3202002323330300-3203310130032122-2131222031110221-0000003220113303-1110031103202133"></a>

## custom_network_config.sli_config.no_static_routes — no_static_routes / 331023111323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- custom_network_config.sli_config.no_static_routes

<a id="canonical-3100120311210332-0311102133012312-0020030132001002-3110123123221031-1220020121202003-1013311221032100-2132311302101032-1000100323331033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-3333310130203012-0031200122023300-0333123220123333-2011202122302322-3300101021202133-0112210003201132-3102322102001110-1323303022212121"></a>

## Direct properties — no_static_routes / 331023111323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233123001111131-2221000233022233-0131330132323101-1230020202331330-3231130302131223-3303330011100022-2133233310131133-1323221200213111"></a>

## Next pages — no_static_routes / 331023111323 / 4

- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1221101200233321-1330100322232002-3122213320131310-3122103201330330-3203123311203330-0111330331233310-0011101200110000-0113012301112021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020313201320133-0110012233332332-3110012133231132-2221020123000312-2032310222132303-0312111110303222-3212003201222033-1112231330201331"></a>

## custom_network_config.sli_config.no_v6_static_routes — no_v6_static_routes / 023020003111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- custom_network_config.sli_config.no_v6_static_routes

<a id="canonical-3203023221221022-0012000203302011-3023210231223332-3221132330121331-1010132130310122-1202331210031210-3023113032300332-1322222301203312"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no v6 static routes.

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

<a id="canonical-1133120201020302-1332232102012323-0333002322303313-2210331123000033-2101110220302121-2001023332031233-0122222100213233-0203112131333021"></a>

## Direct properties — no_v6_static_routes / 023020003111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223032203212022-3311131213300320-3313313230230131-0310111023210022-2002213100313323-3213321230030300-2121111131222223-0031202013111300"></a>

## Next pages — no_v6_static_routes / 023020003111 / 4

- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110300330323303-1323112000020130-0131303332230332-0121013023101001-1301121032211302-1213203211323223-2310113130033223-2010120213231231"></a>

## custom_network_config.sli_config.static_routes — static_routes / 011311022310 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- custom_network_config.sli_config.static_routes

<a id="canonical-1311110213233013-0213213012312123-3201213321012003-1113311221033212-1231211113211333-3122200313012022-1313010303321033-2210000220230112"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-2231220001001203-3333201332100133-0001320021201202-0310213232010100-3330331332033100-2303230021101211-0110210333210301-0211330012000213"></a>

## Direct properties — static_routes / 011311022310 / 3

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300): complete subsection reference.

<a id="canonical-2133330310032222-3203122113022221-3010200001011110-0312001311323000-1101320212330101-2222223220012102-3010130333011220-0033222323033230"></a>

## Next pages — static_routes / 011311022310 / 4

- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101013033023300-2312223131331111-2231032022002313-1012130012033111-0211332302232022-3033312131112202-2200200330311123-3130110321122220"></a>

## custom_network_config.sli_config.static_routes.static_routes — static_routes / 020001011032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- custom_network_config.sli_config.static_routes.static_routes

<a id="canonical-2320331132303013-2331303223112321-3003121210100302-0030131310111211-0100000221011213-1211212003001012-2131030030011112-1121002301203101"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-3002212110022110-0121033321301211-1323321003101103-1123113132111322-1320102032320133-1320001333230332-0322232331302233-0120301023233113"></a>

## Direct properties — static_routes / 020001011032 / 3

<a id="canonical-3222003131023102-3122213032100223-0313022002103311-2202212031310223-3111332213210331-0122223101310210-3012013313102232-2210233110103320"></a>

<a id="canonical-1120021023011103-1203302122002021-2003120131021111-0103003320020131-0021032233211030-0003100102012130-3312311223002013-3111021130312322"></a>

## attrs property — static_routes / 020001011032 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-2132231013130020-3313223030233122-0200231301221330-2103233030220110-2302003112321212-3203100110113202-3203102031020213-2333220131222313): complete subsection reference.

<a id="canonical-2012320223130023-1110302120312033-2310300312331102-3223103002132032-1131321230210102-3120113033030202-1333023121020223-3130131031132033"></a>

<a id="canonical-0100011022222130-3030001320323112-3332121302013112-2001200312021120-3001223300002131-0010322211133011-0133110323100101-1001231300312322"></a>

## ip_address property — static_routes / 020001011032 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-0202122011210031-1333000302311013-0231320023320032-0201321302333123-0300303121213122-0230222231000032-1113130202333122-3320011123333310"></a>

<a id="canonical-1302000031230320-2110002031110200-0300113123013131-0121232132101323-0300211120222312-0200330322202333-1020110003113301-1000333302102023"></a>

## ip_prefixes property — static_routes / 020001011032 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221): complete subsection reference.

<a id="canonical-3322220322203303-3203202233122310-2000133330111201-1120001233330111-1310100331322010-0210313000013130-0111130311222021-1330033222033223"></a>

## Next pages — static_routes / 020001011032 / 7

- [custom_network_config.sli_config.static_routes.static_routes.default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-2132231013130020-3313223030233122-0200231301221330-2103233030220110-2302003112321212-3203100110113202-3203102031020213-2333220131222313)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2132231013130020-3313223030233122-0200231301221330-2103233030220110-2302003112321212-3203100110113202-3203102031020213-2333220131222313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312131302223333-2122300112030123-1130200021212003-2212223003020003-0100132223300031-0302202323230210-2323333130000311-0312222213122221"></a>

## custom_network_config.sli_config.static_routes.static_routes.default_gateway — default_gateway / 232232221121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- custom_network_config.sli_config.static_routes.static_routes.default_gateway

<a id="canonical-1023020113012320-0111321320210221-1222220321133312-0331233321212301-1013222102033310-3101330122022220-2033130122103113-0011232032031301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-1102131210221132-3220201112312113-2220323013223231-2230223102310100-0101202001300110-1323303103103103-2003102113230100-3001213120213221"></a>

## Direct properties — default_gateway / 232232221121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222110320313111-1321113022003331-3111020310131332-2022212131121223-0333111231030003-3101013100311030-2311301212300210-0230301310102312"></a>

## Next pages — default_gateway / 232232221121 / 4

- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210303113323031-2220211312323321-1002310203231130-0132322133210130-1100012223210211-2202021232211320-2030132110221120-3113030100301031"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface — node_interface / 133202120230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- custom_network_config.sli_config.static_routes.static_routes.node_interface

<a id="canonical-0013120010033031-2030203002033303-0131302103202233-3102021013232012-2011110110210020-2101313310302021-3011330332212221-0003223323113101"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-0223302132223202-0200230330120033-3312231020122221-0112213133131031-3032322001223223-2132203131201131-0110133033100321-1233122000200023"></a>

## Direct properties — node_interface / 133202120230 / 3

- [list](data-sources--voltstack_site--reference--group-005.md#canonical-3312312102231012-1332210310202000-0101130100002313-1003332201010111-3223223221113332-1210122022011202-2022320311333332-1131102313311211): complete subsection reference.

<a id="canonical-1203011223102013-0323311021000030-2132303011021331-0211211121112203-0001103011302332-1001300201320223-3130011102323130-2022223112101000"></a>

## Next pages — node_interface / 133202120230 / 4

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-3312312102231012-1332210310202000-0101130100002313-1003332201010111-3223223221113332-1210122022011202-2022320311333332-1131102313311211)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3312312102231012-1332210310202000-0101130100002313-1003332201010111-3223223221113332-1210122022011202-2022320311333332-1131102313311211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320310333023223-3303001210210121-2332320231223023-3203300021203031-1331131322132332-2110032230020101-0323321110320003-0222122020233333"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list — list / 112110201111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list

<a id="canonical-2201132212321123-1301033122013132-3111210010202203-0102121010230222-1031232100230031-3310212211031011-0133201000200223-3200301121201100"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-1320101112110211-0001110013030113-0121110101223120-2011202220112311-0003131113031111-2223013220102221-0320123003132022-0331011022022112"></a>

## Direct properties — list / 112110201111 / 3

- [interface](data-sources--voltstack_site--reference--group-005.md#canonical-1333111311030222-1001233112332121-1232102200031202-1120232233101000-2322002301002001-2030323122213020-0211100003233131-2121132300332131): complete subsection reference.

<a id="canonical-1001201221113222-3233312011213300-1230210200010323-0011121001020022-0301130313131323-3323230022303220-0320001021102221-3033301122223002"></a>

<a id="canonical-3321311203322321-3201003033000212-0113120022010212-1312221210312133-0210301102020211-1200203122120303-3202313002223012-0311123122132202"></a>

## node property — list / 112110201111 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-2231200010102232-3230212212130010-3123133130302013-1221220012112213-2133023233303021-1233023023031203-3022212112233021-3001100122213231"></a>

## Next pages — list / 112110201111 / 5

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface](data-sources--voltstack_site--reference--group-005.md#canonical-1333111311030222-1001233112332121-1232102200031202-1120232233101000-2322002301002001-2030323122213020-0211100003233131-2121132300332131)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1333111311030222-1001233112332121-1232102200031202-1120232233101000-2322002301002001-2030323122213020-0211100003233131-2121132300332131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130032231301221-1322221202221121-0301320103303030-1212302101302331-1213002102303033-1202010320022013-1010222300013131-1321100103211320"></a>

## custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface — interface / 110211011032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0311202221311012-2130232012030203-3131300210013031-1320322130312022-1332122230001120-2103121021123220-2102220300222312-3123311022310032)
- [custom_network_config.sli_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2302230123020111-3300031131001213-0220210201012101-2232323300101320-1031112232112103-1030112303213202-0122032220310102-3310103231002300)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1020322103333333-0201000202013023-2032212303332131-0331330302011230-3003010111231231-1133311000333203-0331110210121331-1303312000330221)
- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-3312312102231012-1332210310202000-0101130100002313-1003332201010111-3223223221113332-1210122022011202-2022320311333332-1131102313311211)
- custom_network_config.sli_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-3012012100332011-3321132031210012-2331331232302011-2212331303213323-0331132333323213-1113101231211202-1310030011003233-1303000100003013"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2223323022232203-2103332221000200-3003120013200023-0320101013020100-1023333003032311-1220020112230202-0321011110200022-3203222331331112"></a>

## Direct properties — interface / 110211011032 / 3

<a id="canonical-3012213310012010-1132013100312033-0102200011300121-3310022231022210-3222310133213123-0021001033311332-3222130331113132-1123301100302221"></a>

<a id="canonical-2000030121233223-1000023002203130-0223012131200320-2221111301120223-0030110100032310-3010211202323233-0013003023121001-2220202112233230"></a>

## kind property — interface / 110211011032 / 4

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

<a id="canonical-2133322332323211-3010022320321300-3132211010012020-3323203311000323-1113232223003130-2332120221310100-3112332231311333-1221123200033301"></a>

<a id="canonical-2330100121221033-1210011301333322-1303221031210300-0333121002102010-0230212303100101-1210200321031113-1313013313222221-0101311032112110"></a>

## name property — interface / 110211011032 / 5

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

<a id="canonical-1030133213033302-3303003130030233-2233231331033032-1021313020021311-2311031002233233-0223110101313000-3213132123122110-0111310310120202"></a>

<a id="canonical-1301100220230231-1023120231031321-1333223101101231-2003020322301102-0303031023022101-2021312230311333-1120120311230120-3013113213223103"></a>

## namespace property — interface / 110211011032 / 6

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

<a id="canonical-3112103220121321-3112023131011112-1223022202021213-2011013113213030-3033000222330113-0232203112101021-2100003112233211-0132210130223300"></a>

<a id="canonical-3121300100303000-2021010101323311-0031023012211302-2221333002132003-1301013210113100-3321011021310302-2101100031312021-3201033023032133"></a>

## tenant property — interface / 110211011032 / 7

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

<a id="canonical-1001120130301313-3131311230203112-3113103211121220-0031030010103101-0110231112123001-1103320031331111-2033012110221300-3131121200302311"></a>

<a id="canonical-0320310112120231-3302211221212221-1301320312111311-2131133002223321-3203303230303110-3331312321230130-1120020001012132-3030231233331322"></a>

## uid property — interface / 110211011032 / 8

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

<a id="canonical-2212133113222122-3001313011100220-3032201222013031-2333322123301203-3301200210222202-3132111110023021-0220121003111322-0322222232100301"></a>

## Next pages — interface / 110211011032 / 9

- [custom_network_config.sli_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-3312312102231012-1332210310202000-0101130100002313-1003332201010111-3223223221113332-1210122022011202-2022320311333332-1131102313311211)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310010310302222-2112313201010330-0002123120021231-0112300201310230-1012302020013203-0332301210312330-1220330201322310-0120121020103203"></a>

## custom_network_config.sli_config.static_v6_routes — static_v6_routes / 020233122132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- custom_network_config.sli_config.static_v6_routes

<a id="canonical-2313212110012113-0130131030031230-2203310322103302-0111310032132013-1130212322130102-3103111302310133-3101123031022032-3211000300311321"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-0003220021102313-1130111131210322-3312101112112303-1100223231101313-2133230000220020-0112322313032010-2120122120211322-0310221131112120"></a>

## Direct properties — static_v6_routes / 020233122132 / 3

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121): complete subsection reference.

<a id="canonical-2311000300233002-3100222010310210-1130312133121233-1300003003332111-3233331313000333-0303312030320322-0130032201212120-1301003030322331"></a>

## Next pages — static_v6_routes / 020233122132 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312221022313101-1311230331233233-0303232303312010-1313123313130331-2232023122220020-2030001001110313-0311310333111203-3112110220330323"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes — static_routes / 011033212121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- custom_network_config.sli_config.static_v6_routes.static_routes

<a id="canonical-1221233031130300-3213233332100031-0101102231001020-1222002030330101-0320231132213211-1330002301100000-3030132221102013-0312333101101321"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-3110210221131022-1100003100023232-0020320132133331-1101002111313132-3120301123122011-2002232301220311-0233133003232233-0130103003020223"></a>

## Direct properties — static_routes / 011033212121 / 3

<a id="canonical-0213210113233132-0003333030120033-3310011212212200-1333213002313102-0013320310100122-0000301210110210-1021301110212301-1313223112321002"></a>

<a id="canonical-2221103013333112-2033133210020201-1302121031011302-3002023321131232-3232313010011321-0222333133020301-0201021331130300-2222221021032330"></a>

## attrs property — static_routes / 011033212121 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-0220211220210103-2131321010010213-0321023113130331-1302202201220110-3220113110122220-1101023022313103-1220113011310322-0333031312220121): complete subsection reference.

<a id="canonical-1101210331011022-0303212310123101-3323112220300313-3001003233112033-1121020003022121-0011110233212312-3322302131130111-2202021300000001"></a>

<a id="canonical-3222310212312021-0200113002111001-1233332003133313-0110123320303101-2000203222303020-2112101101333023-0201230233031300-2213003121203330"></a>

## ip_address property — static_routes / 011033212121 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-1012133313032330-3030102132011203-3103033321130323-1213013133031201-0103102003223033-0200203312321030-3210033332222313-1202012033101121"></a>

<a id="canonical-0303210323111212-1311312011023301-1311323332333102-3013133211113130-2123212130321011-2313231012330223-3220112002222120-3221201111222322"></a>

## ip_prefixes property — static_routes / 011033212121 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-3301031032302113-0030001103003012-2030232203120102-3300212312303320-0230202220321000-3220302303311031-3112020030221132-1301131232312212): complete subsection reference.

<a id="canonical-0022211131103010-0000213120021303-3013331200111302-3332013102010030-1233232222010010-0323123320120132-1211032223202100-1130213122232131"></a>

## Next pages — static_routes / 011033212121 / 7

- [custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-0220211220210103-2131321010010213-0321023113130331-1302202201220110-3220113110122220-1101023022313103-1220113011310322-0333031312220121)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-3301031032302113-0030001103003012-2030232203120102-3300212312303320-0230202220321000-3220302303311031-3112020030221132-1301131232312212)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0220211220210103-2131321010010213-0321023113130331-1302202201220110-3220113110122220-1101023022313103-1220113011310322-0333031312220121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203113201021221-2202303223101110-1001320113110032-3120020232302300-1111110320303213-1113303210200222-2323200101021030-3120300122000111"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway — default_gateway / 303100001312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- custom_network_config.sli_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-0231322033000233-3223013322123111-0011111213111331-3002022232102223-2002132011113032-3312112000232212-0000222321321310-1122202231021113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-2202232033221012-0232122022233311-1102100222230220-3300310211123222-1132033312203122-0320322201103112-0113201333233111-1221122212200223"></a>

## Direct properties — default_gateway / 303100001312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1210002303311030-1332111113000333-1130113103321032-3021212001123230-1201320003201103-2013110020133000-2013023123001102-0002200221323103"></a>

## Next pages — default_gateway / 303100001312 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3301031032302113-0030001103003012-2030232203120102-3300212312303320-0230202220321000-3220302303311031-3112020030221132-1301131232312212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232030132300110-3032000200123013-2232100022010302-2313112312002212-0103330132321210-0312002313302013-3131321132001231-1323313131330231"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface — node_interface / 100233231113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface

<a id="canonical-0331120013111211-2313202021211001-2022322223233031-0202110121002131-0220020120011222-2121232322302311-3021311132030321-1023032121031012"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-0233320233303330-0013332100123121-3121001033321201-0312203112323303-1212321132332203-3033010002320302-3223130300122110-1023133100010231"></a>

## Direct properties — node_interface / 100233231113 / 3

- [list](data-sources--voltstack_site--reference--group-005.md#canonical-2311212213320201-3120222300311232-1312100012031320-0313121113203332-3330013333321120-0130220103131130-1300210233333210-0211001111001231): complete subsection reference.

<a id="canonical-2033001120220312-1012302033013103-2102000010201330-0333332133321200-2213330000302011-0003221111111102-2032021330211232-0213023203022302"></a>

## Next pages — node_interface / 100233231113 / 4

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-2311212213320201-3120222300311232-1312100012031320-0313121113203332-3330013333321120-0130220103131130-1300210233333210-0211001111001231)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2311212213320201-3120222300311232-1312100012031320-0313121113203332-3330013333321120-0130220103131130-1300210233333210-0211001111001231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311031201132131-3330300332310302-3103001022200323-0323131222213013-1230310333312230-0202021003100230-2123120100113333-1102200133300203"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list — list / 131330012000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-3301031032302113-0030001103003012-2030232203120102-3300212312303320-0230202220321000-3220302303311031-3112020030221132-1301131232312212)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-1210001331000120-1301331310113302-0010101210031320-3013102230220300-1111002030203100-0333012020120212-3011233301311200-2103301211033221"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-0010223022312303-1020100012001231-1220100222132020-0233010303021011-0322222012122313-3301130100013102-3103032003223220-2132131201321330"></a>

## Direct properties — list / 131330012000 / 3

- [interface](data-sources--voltstack_site--reference--group-005.md#canonical-2323221302002331-3103133231211311-1112321220022112-2321112133111022-2202030011003002-2200322300023333-1212222102120133-1223023100211301): complete subsection reference.

<a id="canonical-3002320231111312-3212030303232220-3210220201233200-2111332022323003-3022122121102031-1130113212320133-2330331211020011-1113223131002011"></a>

<a id="canonical-1120312232020130-2033120210301212-3120331202331022-1131023023331300-2120220111333032-3012332121012222-0123231002223311-3002133133222223"></a>

## node property — list / 131330012000 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-3010121100200221-0303201101100000-1300300320322301-3020222222011221-3000003012021021-3330031211123210-2030312023031120-2030122223332233"></a>

## Next pages — list / 131330012000 / 5

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--voltstack_site--reference--group-005.md#canonical-2323221302002331-3103133231211311-1112321220022112-2321112133111022-2202030011003002-2200322300023333-1212222102120133-1223023100211301)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-3301031032302113-0030001103003012-2030232203120102-3300212312303320-0230202220321000-3220302303311031-3112020030221132-1301131232312212)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2323221302002331-3103133231211311-1112321220022112-2321112133111022-2202030011003002-2200322300023333-1212222102120133-1223023100211301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102223330131201-1213123120100022-0022001010020122-1133113230033001-0220020022200331-0001200033223003-3201320203130103-0213110101330233"></a>

## custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 233220013220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.sli_config](data-sources--voltstack_site--reference--group-005.md#canonical-0103212222312220-0001131112200333-0121313002131020-0023010000132011-1000200021230333-1112123331022122-2110203303220333-2102132113002333)
- [custom_network_config.sli_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3320121130202022-3001002310030011-2223320112323303-2103013223023011-2323000001203210-3003133310101131-0312011030231121-2131313132220021)
- [custom_network_config.sli_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0302302303032122-0222123131221303-2023102123101121-0303230111102120-3023301030031003-0233032111020123-0123231220232322-1201220313200121)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-3301031032302113-0030001103003012-2030232203120102-3300212312303320-0230202220321000-3220302303311031-3112020030221132-1301131232312212)
- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-2311212213320201-3120222300311232-1312100012031320-0313121113203332-3330013333321120-0130220103131130-1300210233333210-0211001111001231)
- custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-2133003133131213-2332013023300032-0010202322302200-3010010101300013-1123320302301323-2321222332200003-3332010213232003-0321000200122023"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2010113021210112-3022010300013103-0303310322320123-2003332200030030-2310100101112030-3101233322303131-3303002032013220-0013021231231132"></a>

## Direct properties — interface / 233220013220 / 3

<a id="canonical-3012010202320012-2033131231230003-0011332000210112-0310132323123320-3030110011023101-0001333301003101-2030301102221110-2223121130313220"></a>

<a id="canonical-2022320123030202-3222121103212221-1310111312223032-0303100110122010-2021311211121333-0123222321302132-3022222301223100-3020231031102133"></a>

## kind property — interface / 233220013220 / 4

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

<a id="canonical-3210101232331313-0031112001300123-0012221233133302-1132013321201223-1231132110032123-1123122001211031-3113121123212030-0130100312310302"></a>

<a id="canonical-0313312030103103-1222302303032330-1233203102330303-2021303030232013-1220302122313122-0223121131002232-1112221030220121-2011010211331130"></a>

## name property — interface / 233220013220 / 5

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

<a id="canonical-3212113020333132-2020132133303220-1333101233333121-0003123032302132-1121222200111302-0221200211212221-1130032230020021-1113211302121222"></a>

<a id="canonical-2231000103032312-0120232132120022-1332120013323300-3212231112002021-0231101323210101-2302312212011120-3133132121023132-1023012311312001"></a>

## namespace property — interface / 233220013220 / 6

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

<a id="canonical-1101330202112013-0011200122201323-1120203203023110-2333030020123213-1333222001213020-3230011331132331-3202311031221200-3011003320000001"></a>

<a id="canonical-2313131200023103-2022111333213233-2111212012102013-2211123322033322-0210120321203000-3220320231003213-2032331030023103-3003123210100223"></a>

## tenant property — interface / 233220013220 / 7

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

<a id="canonical-3330321131133013-3212011201311221-3030003230133222-0020233033011132-0212001323231233-2233002112002300-0012333202330121-0212320232302220"></a>

<a id="canonical-0320222031012130-1122013003300001-2010131311322322-2213033033302320-2100031002022002-3221202103120012-0001220003301323-2331231233231311"></a>

## uid property — interface / 233220013220 / 8

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

<a id="canonical-2103111202210102-3120221012010013-0210103123121233-3033112231013031-1330211120010300-1011023002232213-1333201012000033-2132011111332223"></a>

## Next pages — interface / 233220013220 / 9

- [custom_network_config.sli_config.static_v6_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-2311212213320201-3120222300311232-1312100012031320-0313121113203332-3330013333321120-0130220103131130-1300210233333210-0211001111001231)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213112323303113-0030032210023122-3203122332200133-1331333211233232-0210000101100103-0310121003222110-2032033110102013-3113212312000212"></a>

## custom_network_config.slo_config — slo_config / 002112130102 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.slo_config

<a id="canonical-2000321010333023-3332320113003013-3332212001022320-3200301102332301-2002102103303311-3221201021031221-1101232230113023-3023223101102031"></a>

Type: `"single"`. Computed.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_static_v6_routes\",\"static_v6_routes\"]"
}
```

<a id="canonical-2313112311101210-3033123202200302-3320202122221133-3231003300303123-2033032221103212-0313110113100033-3003003110210331-3012313110331010"></a>

## Direct properties — slo_config / 002112130102 / 3

- [dc_cluster_group](data-sources--voltstack_site--reference--group-005.md#canonical-0102101220201001-0202112010130003-3331000022332333-3312000323310111-2132000002032331-3100322011203232-1310201311012122-2322223130202111): complete subsection reference.

- [labels](data-sources--voltstack_site--reference--group-005.md#canonical-3230322110332033-1201222202202310-2023100213012313-2311221120201000-1302333312010112-2132310233311200-0201321010333031-0000200232332203): complete subsection reference.

- [no_dc_cluster_group](data-sources--voltstack_site--reference--group-005.md#canonical-1313300111132002-2023230011132230-1021122320010130-0221023233223101-0231121131231110-3003102032001103-3213211210232212-0110010121012013): complete subsection reference.

- [no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3322330012003100-1232010122333131-2301310121013300-0330301030030033-0301011032222203-3301113110111301-0321312113013213-1232211313000131): complete subsection reference.

- [no_static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2112231113001223-3222033310123231-1312200332033322-2220102102211200-0213333322332101-0123003333022330-3000012000322203-2122033203223201): complete subsection reference.

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123): complete subsection reference.

- [static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001): complete subsection reference.

<a id="canonical-1132300202012011-0110033013012111-3220012101310002-1020103011221211-3311112002211012-0333121200201233-2030021303121303-3332133231120212"></a>

## Next pages — slo_config / 002112130102 / 4

- [custom_network_config.slo_config.dc_cluster_group](data-sources--voltstack_site--reference--group-005.md#canonical-0102101220201001-0202112010130003-3331000022332333-3312000323310111-2132000002032331-3100322011203232-1310201311012122-2322223130202111)
- [custom_network_config.slo_config.labels](data-sources--voltstack_site--reference--group-005.md#canonical-3230322110332033-1201222202202310-2023100213012313-2311221120201000-1302333312010112-2132310233311200-0201321010333031-0000200232332203)
- [custom_network_config.slo_config.no_dc_cluster_group](data-sources--voltstack_site--reference--group-005.md#canonical-1313300111132002-2023230011132230-1021122320010130-0221023233223101-0231121131231110-3003102032001103-3213211210232212-0110010121012013)
- [custom_network_config.slo_config.no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-3322330012003100-1232010122333131-2301310121013300-0330301030030033-0301011032222203-3301113110111301-0321312113013213-1232211313000131)
- [custom_network_config.slo_config.no_static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2112231113001223-3222033310123231-1312200332033322-2220102102211200-0213333322332101-0123003333022330-3000012000322203-2122033203223201)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0102101220201001-0202112010130003-3331000022332333-3312000323310111-2132000002032331-3100322011203232-1310201311012122-2322223130202111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311133121021300-0203220313103301-0030000331020101-0032000223021203-1012202121102131-2023120123311010-2202101231302311-1033120230232211"></a>

## custom_network_config.slo_config.dc_cluster_group — dc_cluster_group / 333323323103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.dc_cluster_group

<a id="canonical-3212130221333231-0030131120123002-0333033100113131-1203131011011022-1030123101213321-1331320013110323-1112213210220013-0221132011000310"></a>

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

<a id="canonical-3201233102103100-2220232102011031-3003200002330322-1211211003132023-2021331333010123-0201020001331022-0203000001220322-3130323200101022"></a>

## Direct properties — dc_cluster_group / 333323323103 / 3

<a id="canonical-3003121000201213-3003131110011001-2012301113202233-2132322211023001-2113310302330001-2003101202213102-2330321313123222-1220233030131232"></a>

<a id="canonical-0331122123330312-0021210221212112-3212321330002231-2110000202221301-0002222313001013-1002110012031120-1103212212021301-1331302202011011"></a>

## name property — dc_cluster_group / 333323323103 / 4

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

<a id="canonical-1233311132132001-3330100123010303-1231003012131303-3113330323132011-0203003321033133-1221232111102303-0332030030103211-0103321320002212"></a>

<a id="canonical-3312303210332211-2123111232213312-3001133213113121-0220122113302130-3021031111332231-3220220003300222-2222330323022223-1102002101122331"></a>

## namespace property — dc_cluster_group / 333323323103 / 5

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

<a id="canonical-1201212333321230-1121010222030300-1003022131321211-1223102301022232-2320112223202300-2003001330312000-1102233321021020-0102113322323113"></a>

<a id="canonical-1200021310211230-1021222210111102-1212300121210000-0220210101201011-2332220310110123-3222313301223233-3120232012010002-2032102301030122"></a>

## tenant property — dc_cluster_group / 333323323103 / 6

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

<a id="canonical-1313301311021330-0020223301122121-2323001231230300-2223030012200201-1330211001211012-1023210021023211-1012100003103010-3333331312203211"></a>

## Next pages — dc_cluster_group / 333323323103 / 7

- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3230322110332033-1201222202202310-2023100213012313-2311221120201000-1302333312010112-2132310233311200-0201321010333031-0000200232332203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303200202210133-2312310030310320-1132222213213131-1311210021031302-3320023003223203-2111101223221203-1112302321200033-2221323320300103"></a>

## custom_network_config.slo_config.labels — labels / 201110010110 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.labels

<a id="canonical-1301203330011002-3213300333101030-0302112313310202-1200133011002230-1001302122213301-3300000001101022-3021012233230313-2330130331300230"></a>

Type: `"single"`. Computed.

Add Labels for this network, these labels can be used in firewall policy.

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

<a id="canonical-2313301200223113-2300003011220111-1012212323333010-3022323021111331-2021011120321311-1211112211131302-2302312331320123-3020233110003231"></a>

## Direct properties — labels / 201110010110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130302101330230-0201000330100111-3323213023303200-3020012030322231-0021003011122323-3101121231031010-1311101120013322-2211023021331332"></a>

## Next pages — labels / 201110010110 / 4

- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1313300111132002-2023230011132230-1021122320010130-0221023233223101-0231121131231110-3003102032001103-3213211210232212-0110010121012013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012022331300313-3001212212113132-0111100011330123-3221112323001011-1333331203030032-0113223020103033-3301121113223100-3021132321012222"></a>

## custom_network_config.slo_config.no_dc_cluster_group — no_dc_cluster_group / 001211222103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.no_dc_cluster_group

<a id="canonical-2032031202002102-1030133130321213-0120102021012321-0030220032121231-1322011323232331-3022230222010222-3312121133021120-3133303333121333"></a>

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

<a id="canonical-0132220303301210-0311032320312001-0110231232002211-0301021120221230-0121220113103303-0232222222311010-2030000300231320-1230131110131211"></a>

## Direct properties — no_dc_cluster_group / 001211222103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0311122223232011-0113000230033331-3231223331210013-1312222330101032-1021213011120200-3323123113110213-3112013323200302-3231321311010311"></a>

## Next pages — no_dc_cluster_group / 001211222103 / 4

- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3322330012003100-1232010122333131-2301310121013300-0330301030030033-0301011032222203-3301113110111301-0321312113013213-1232211313000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223031022132321-0021110012000322-0133020103301030-3201033001030000-2001321321032101-1003210310013121-3031230213320222-2231301123310020"></a>

## custom_network_config.slo_config.no_static_routes — no_static_routes / 023120333311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.no_static_routes

<a id="canonical-1321121113130310-0300133300333311-1012112133021332-0211330033111220-1321001321130111-2103131020333311-2020130031211220-2302000321111213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-1322301212312022-0203031130300130-1022101223102202-1000113002323333-2011312302003031-0220101222132233-0232023032000031-0332232302302022"></a>

## Direct properties — no_static_routes / 023120333311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102102022312313-0031000100233332-2203321132331000-0000013323010133-0010332001132031-2012211000230032-0022113130011123-1122213010213332"></a>

## Next pages — no_static_routes / 023120333311 / 4

- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2112231113001223-3222033310123231-1312200332033322-2220102102211200-0213333322332101-0123003333022330-3000012000322203-2122033203223201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112031123202130-3302131022300130-3300202202321332-0103012021100312-0221321100222002-2031012312102030-3003103121200013-2320101003202231"></a>

## custom_network_config.slo_config.no_static_v6_routes — no_static_v6_routes / 313210002032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.no_static_v6_routes

<a id="canonical-2213130213002132-0020120221130022-2222310031311210-0000231220211311-1011332133110211-2200010303230310-0320223302011330-3330022010112222"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static v6 routes.

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

<a id="canonical-2320101232220122-3021113231301211-2010301123022130-2310133200030210-0232302320113213-1010230023232303-0303200200003303-0111231311022211"></a>

## Direct properties — no_static_v6_routes / 313210002032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200302211102321-3110132210231121-0013102112030302-3320320003010031-2021210033110320-2032211112223330-3303222031010123-0111112212023230"></a>

## Next pages — no_static_v6_routes / 313210002032 / 4

- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303113012113210-1133303220120032-2320030032133210-3002130121031001-2102130113320313-3333211110332333-2112132120013033-0212102331313200"></a>

## custom_network_config.slo_config.static_routes — static_routes / 033002220020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.static_routes

<a id="canonical-3202203131102000-2213102211001030-3213122323131001-0003002313130121-1100011233120231-3020222112023132-3231012233321013-0133121102033100"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-1201212002001000-3211032200112021-2322302132022312-1230121230333330-1031031211111112-3033310013113201-0331200120123210-3313212020013222"></a>

## Direct properties — static_routes / 033002220020 / 3

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310): complete subsection reference.

<a id="canonical-1011032010332333-3321012133123023-0210020112102031-1122000213012122-3122123220133211-2012333200311221-0313030101301122-1321220301122000"></a>

## Next pages — static_routes / 033002220020 / 4

- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123310313212320-2033131202221231-2311220213101112-1130000300020231-3030333203310122-3113230023213002-2333110121232030-1230233200310101"></a>

## custom_network_config.slo_config.static_routes.static_routes — static_routes / 210310201202 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- custom_network_config.slo_config.static_routes.static_routes

<a id="canonical-1102113122223232-2130101102321120-3223220002201102-2033233101012311-3011330222220013-3123012102032212-2122130013232220-3132002200113120"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-2332120013033002-0302131310211101-0222032322313320-0200132212103311-0332203131010123-1231111003220331-1321230100001301-0111213011111020"></a>

## Direct properties — static_routes / 210310201202 / 3

<a id="canonical-0232303100203302-1020220202332021-2023033002030221-0312220320120202-1301013223110102-3111331332300332-1123010333201001-3233210100332222"></a>

<a id="canonical-3003101120102011-2212210111000303-2120021121030303-3302021133310101-0102313021031131-3303210110200130-0111322211022303-3321200202003131"></a>

## attrs property — static_routes / 210310201202 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-0301230012231110-1133312312013201-3113201330232230-3303030122002231-3230310200303213-3112101120322313-0310221003211311-2001210231302212): complete subsection reference.

<a id="canonical-0131010310201113-2220330311221111-3212102121032311-2230123031023020-2132331012023212-2233121003333033-1233333120312213-2220322233113002"></a>

<a id="canonical-1320003221311210-2231231211132112-0202313113001002-2230022120301101-3023101222223133-0212000301010300-2022032220231323-3230321003002021"></a>

## ip_address property — static_routes / 210310201202 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-3310322221101210-1032300121223232-2231022300300212-2223303113212102-2032011020303200-1111023122311333-1003002331020130-1102202211300100"></a>

<a id="canonical-3020033221011223-2322202210310120-2222210130230131-3203131222130031-0101001011330330-2210232323231001-2000322032022212-1211203012113101"></a>

## ip_prefixes property — static_routes / 210310201202 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2233100312123320-3121010003200320-2303030312121210-1102331203110013-3333110322110123-2301110222133111-1311220133323113-0110103100220303): complete subsection reference.

<a id="canonical-1000023310010213-1333120203210132-0123121231021121-2002120212310010-2323000322333133-1102102323213111-3023123131321331-1232023302330311"></a>

## Next pages — static_routes / 210310201202 / 7

- [custom_network_config.slo_config.static_routes.static_routes.default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-0301230012231110-1133312312013201-3113201330232230-3303030122002231-3230310200303213-3112101120322313-0310221003211311-2001210231302212)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2233100312123320-3121010003200320-2303030312121210-1102331203110013-3333110322110123-2301110222133111-1311220133323113-0110103100220303)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0301230012231110-1133312312013201-3113201330232230-3303030122002231-3230310200303213-3112101120322313-0310221003211311-2001210231302212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002320203202011-0312302021032210-1331001221113111-2033020033122232-0101121002231220-2320123113100312-0132123033133023-2113120303310331"></a>

## custom_network_config.slo_config.static_routes.static_routes.default_gateway — default_gateway / 202232030001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- custom_network_config.slo_config.static_routes.static_routes.default_gateway

<a id="canonical-2200230300223110-1133133100100202-3322022122330210-3321301123301113-1203030102201210-1120213223300001-3000003003101312-0203322130003321"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-0121222030223203-1331212110202012-0330110032202122-3021233023021123-0200201331331200-2003130321330132-3011233230030131-0121022133112300"></a>

## Direct properties — default_gateway / 202232030001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1202302323103020-1110322201332000-3020011301133213-0303001330132100-1010011000121300-0202230101121210-1130003202233323-2020331222010302"></a>

## Next pages — default_gateway / 202232030001 / 4

- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2233100312123320-3121010003200320-2303030312121210-1102331203110013-3333110322110123-2301110222133111-1311220133323113-0110103100220303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220230332031202-2023230322000023-3212011321333201-0202010021020202-1311021311323032-1110302103332313-2211332302233210-0222110330101320"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface — node_interface / 021333031010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- custom_network_config.slo_config.static_routes.static_routes.node_interface

<a id="canonical-3233112013103033-3102032332211212-2200031320120331-0120100322303303-1000312113321301-1000313303113210-3222321130123200-3013032333021203"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-3213303012022332-2231220301210332-1201122101111010-3223331112201022-0223022032013102-0122220023330030-2111030203020000-2123033132131332"></a>

## Direct properties — node_interface / 021333031010 / 3

- [list](data-sources--voltstack_site--reference--group-005.md#canonical-1123332121103023-1202322121212103-2233122231320313-3103113030311121-1201311030321303-2120131101110323-2110002323233312-0203022323333302): complete subsection reference.

<a id="canonical-2110113312321032-3032323121122122-1221013122003330-3132222320220322-1322223301101212-3021202223302101-0230013223133023-0222320113011332"></a>

## Next pages — node_interface / 021333031010 / 4

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-1123332121103023-1202322121212103-2233122231320313-3103113030311121-1201311030321303-2120131101110323-2110002323233312-0203022323333302)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1123332121103023-1202322121212103-2233122231320313-3103113030311121-1201311030321303-2120131101110323-2110002323233312-0203022323333302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202310202212100-0211221131301110-1023003030033031-0312211322200232-2113302221130310-2132312330001233-1213213013111111-2123212203222132"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list — list / 131100122100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2233100312123320-3121010003200320-2303030312121210-1102331203110013-3333110322110123-2301110222133111-1311220133323113-0110103100220303)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list

<a id="canonical-3232312011220322-0102312111221032-1023111333021230-0202223022312022-2010133131130013-0200231100303201-3011101220212021-1222100002213211"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-1310333310133233-2221003113212212-2230311011210122-0111012210121002-0301123103200130-1131112033003212-1132331210100131-3212203232203301"></a>

## Direct properties — list / 131100122100 / 3

- [interface](data-sources--voltstack_site--reference--group-005.md#canonical-0212211012232213-2212033003232212-0110312122231133-0130232201323333-3021310221212003-0100133033003002-0120302020211220-2010220300102300): complete subsection reference.

<a id="canonical-2303000303323213-2230020123322310-3232221021222101-0301120201331013-2023303211322203-3122121231320222-0313223110303200-0222023002333331"></a>

<a id="canonical-3012013233221113-3313310120222012-3022122201110231-0300102210330321-0123322313131100-3231211011113333-3030312030131010-1212300322031032"></a>

## node property — list / 131100122100 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-1200222221031330-0013130300321312-1102203202332221-1120320032222100-2120133221221202-2132021112020300-1323003133213312-3021103103302030"></a>

## Next pages — list / 131100122100 / 5

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface](data-sources--voltstack_site--reference--group-005.md#canonical-0212211012232213-2212033003232212-0110312122231133-0130232201323333-3021310221212003-0100133033003002-0120302020211220-2010220300102300)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2233100312123320-3121010003200320-2303030312121210-1102331203110013-3333110322110123-2301110222133111-1311220133323113-0110103100220303)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0212211012232213-2212033003232212-0110312122231133-0130232201323333-3021310221212003-0100133033003002-0120302020211220-2010220300102300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011101231110002-2203300210303311-1122223002103202-3111010121233132-1313233310013200-3120033130322312-3031032012222211-2120110322012021"></a>

## custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface — interface / 132033020021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1232310221332121-0310032100333333-0113322003123333-3113230313301332-0120322222001113-3132313001230111-2103012311033231-1022333212220123)
- [custom_network_config.slo_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2201213230123330-1233011123210333-2332020000101220-1222132122330021-3011310132313231-1220033133333001-0103310023121332-1123021223213310)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2233100312123320-3121010003200320-2303030312121210-1102331203110013-3333110322110123-2301110222133111-1311220133323113-0110103100220303)
- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-1123332121103023-1202322121212103-2233122231320313-3103113030311121-1201311030321303-2120131101110323-2110002323233312-0203022323333302)
- custom_network_config.slo_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-1320232101303320-0221212222311121-0331003012222132-3030111023222212-1121203300312233-2233311033022122-0230020000102023-0210221020321213"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3230332223011233-3013111022332313-1330121310301210-3222332232023120-2021002100113012-1311012211333310-3012121211333101-0020300310203003"></a>

## Direct properties — interface / 132033020021 / 3

<a id="canonical-2123123310101001-0301122000302102-1302311303332223-0313021133222011-1212033021223211-0121220220323112-1212200333023032-1110101020202003"></a>

<a id="canonical-3123021331233302-2201311132001323-3330100102330013-0000102211133001-1102021233212030-1021300210020332-2201011312322211-2131302022123110"></a>

## kind property — interface / 132033020021 / 4

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

<a id="canonical-3102032131300003-3222112233213010-1131333303220321-0102131032112311-0102121030031300-3213003210010031-0120331312232213-0123023301310320"></a>

<a id="canonical-1010223210130021-3020221031201311-1201113203212211-0131230221113030-3132200323201102-2032121203330131-1310333310133100-0212213021020033"></a>

## name property — interface / 132033020021 / 5

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

<a id="canonical-2002321302200000-0202311232033312-2212020022332022-1123132002133310-1220211323301001-1223110133120300-2110030123012312-0202300102203303"></a>

<a id="canonical-1103110202303302-3011032333003321-3102231303033102-0001103202003111-3020221213321300-3333212131203233-2100012020123100-1220113112323103"></a>

## namespace property — interface / 132033020021 / 6

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

<a id="canonical-1101331102032023-2320020302220311-3131032113212330-3300223000220331-2332201101113210-3132303312312101-1323330310322330-2111121121030022"></a>

<a id="canonical-1303321310102032-2131120220331212-0302231312321322-2023200233100021-3100122112000103-1110030220312322-2033310010311022-3212212332312333"></a>

## tenant property — interface / 132033020021 / 7

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

<a id="canonical-3030301331022223-3210031100102332-2120321131333223-1321010003303101-2212220003303122-2212322212022332-0100301023310231-0302010333111003"></a>

<a id="canonical-1211313012211233-1302221223020120-1330311120232033-2213123300223223-0001213231110202-3001111301013232-1022030232002312-2213022221232331"></a>

## uid property — interface / 132033020021 / 8

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

<a id="canonical-0030321130110122-3302303321231001-0330302010303202-0331020000321200-2132003333011021-0111332213213122-2203233333021210-3211232111232012"></a>

## Next pages — interface / 132033020021 / 9

- [custom_network_config.slo_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-1123332121103023-1202322121212103-2233122231320313-3103113030311121-1201311030321303-2120131101110323-2110002323233312-0203022323333302)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030030321323201-0133331111233220-1212133113112002-1030121021202030-2202320310300122-3103311221033213-2212102222213201-0210121102222110"></a>

## custom_network_config.slo_config.static_v6_routes — static_v6_routes / 221321221033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- custom_network_config.slo_config.static_v6_routes

<a id="canonical-2301103311021311-0101032213311013-2302301311310103-2021101320223330-3220301233313013-2122101300103033-2201232012003203-3303112001210301"></a>

Type: `"single"`. Computed.

Configuration parameter for static v6 routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-1113022202201130-1100311112120330-2122323332103233-1213333120230013-0321103131103311-0331330000101321-2320023033330023-2013303310203011"></a>

## Direct properties — static_v6_routes / 221321221033 / 3

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202): complete subsection reference.

<a id="canonical-2232310300103303-3131032303303323-2213113121213332-0131111022311230-2320231101200321-0002120111133113-0113001013322023-2100323133221013"></a>

## Next pages — static_v6_routes / 221321221033 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102032111312221-0110020011101300-0321123103122232-0303233231031333-2002000002120203-1100010331232220-0022022233321322-1111113033211201"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes — static_routes / 130131300123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- custom_network_config.slo_config.static_v6_routes.static_routes

<a id="canonical-2103302030231313-2133303321203133-0113102010220331-0022021131301002-3012332121310111-3002212310322210-0003011211023000-2112023213321231"></a>

Type: `"list"`. Computed.

Static IPv6 Routes. List of IPv6 static routes.

Upstream description:

List of IPv6 static routes.

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

<a id="canonical-2200030232011231-2223011220111231-3122111330310213-0322032332123013-1330301100132023-2201310233301032-2130221322101230-1001130301213310"></a>

## Direct properties — static_routes / 130131300123 / 3

<a id="canonical-1323002220320023-1102000001030201-1100013100001111-0331232101012021-1122101130011030-3201032310001020-0001102301330233-1110312332232321"></a>

<a id="canonical-3320001310001220-0232223033133321-3123110332000020-3332202233003321-1120011112232222-2002130132012322-3132022013322130-2022222223210202"></a>

## attrs property — static_routes / 130131300123 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-2121112232122201-3121103112221101-2030310331311012-0231212312310333-0223123202313031-1032223121310000-1200100003312310-2302230300003011): complete subsection reference.

<a id="canonical-1133111033101011-0230112130232100-1300301023122200-1213021010011120-2113301011301010-0013111332003200-0312311223211121-3300201313203221"></a>

<a id="canonical-1233010211013131-3032123131020012-1201102003012330-3322323211031330-3102213331113030-0111302030210213-3001031100320032-1323312332320213"></a>

## ip_address property — static_routes / 130131300123 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv6",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-1013321112031100-3331313323210101-3131310221110300-1332302212020032-0101312132200100-1013130132333222-1230303113121300-0321110111223001"></a>

<a id="canonical-0231210120323022-1233130210330033-2230333012302032-0111133113011023-3322202200110112-2120120100132131-2132310121101031-3022230323120122"></a>

## ip_prefixes property — static_routes / 130131300123 / 6

Type: `["list", "string"]`. Computed.

List of IPv6 route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313): complete subsection reference.

<a id="canonical-2312001320031021-3212320122302213-0103231220121210-3312200231100333-1021031310323020-1220221033313311-1323313032211332-3322133122133031"></a>

## Next pages — static_routes / 130131300123 / 7

- [custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-2121112232122201-3121103112221101-2030310331311012-0231212312310333-0223123202313031-1032223121310000-1200100003312310-2302230300003011)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2121112232122201-3121103112221101-2030310331311012-0231212312310333-0223123202313031-1032223121310000-1200100003312310-2302230300003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212020330320012-2130100301020213-2103100103311122-0221133022120300-2111233112023102-0303323122210002-3000212201000302-3203001030110100"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway — default_gateway / 332200233312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- custom_network_config.slo_config.static_v6_routes.static_routes.default_gateway

<a id="canonical-2113111203020311-1230013331210001-0101002212023002-3103323112120232-1021003331203123-3201110022111003-3110131022121130-0113101301021013"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-2331301311221131-0001211022200011-1221030101231210-0321030012032000-2312322000300221-0331233323230320-0220232122012311-2123020110012201"></a>

## Direct properties — default_gateway / 332200233312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013223110011301-0301330202230002-3322211002201101-3100321322222020-2230210003100210-1331301332123020-3010221123323312-1210132122020313"></a>

## Next pages — default_gateway / 332200233312 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231333333023012-3323131003320112-0211333132121212-0110321230221331-0113100220113201-2323321122223210-2113020200323111-0223300102112021"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface — node_interface / 330320213323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface

<a id="canonical-2012300310302213-2301033331123222-2321112122301300-0201211312200103-0221001013230323-0030311223222111-1302332330101001-1012213111010210"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-3203101231020223-3212132232312313-2112002302323310-2012313112312202-0001322321100030-1223232122102231-1312211130213332-2231333201320000"></a>

## Direct properties — node_interface / 330320213323 / 3

- [list](data-sources--voltstack_site--reference--group-005.md#canonical-1110312012212333-0022233200220321-0223330203133133-1300302330003113-3211011210100220-1010133232303031-1201331131110023-2111131312332122): complete subsection reference.

<a id="canonical-3122213003113132-3003211223113311-0310112221001112-3112232212120323-1300030113033301-3123323302133121-2321200012033301-0000231301101001"></a>

## Next pages — node_interface / 330320213323 / 4

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-1110312012212333-0022233200220321-0223330203133133-1300302330003113-3211011210100220-1010133232303031-1201331131110023-2111131312332122)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1110312012212333-0022233200220321-0223330203133133-1300302330003113-3211011210100220-1010133232303031-1201331131110023-2111131312332122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211221230001200-0332321323021312-3001013130112010-0000231230001130-1223103212103222-2022223300001200-0013103011302221-0301022130220103"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list — list / 213100331310 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list

<a id="canonical-1130202031013030-3303122131232013-1031312211212223-2230001222233303-3033320030030311-2201111112133221-3120110201013213-0122301002303122"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-1220013133310332-1103110333022310-2322011232031100-0011003122000121-1132120210023130-0121302222230223-0221220210210032-2302213330213001"></a>

## Direct properties — list / 213100331310 / 3

- [interface](data-sources--voltstack_site--reference--group-005.md#canonical-2030000031301002-0223312303122312-1112202211103331-0231112012210231-0103330021102203-2221133223310330-0320110230300321-2122302223233010): complete subsection reference.

<a id="canonical-0210013302102033-2213302213113302-0023310200230101-2023322201332321-0023012233201000-2330210030210203-2231321101122133-3110002302103333"></a>

<a id="canonical-1121101002231000-3130121121131001-1310023300303220-2222300222230220-0122003103102202-2000322032211103-2221121310303330-2112112121210123"></a>

## node property — list / 213100331310 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-1100101022123311-2313203102330021-2302002131013332-2203002223333310-0201102332132001-2321311002320320-1331231231032232-1022322233132111"></a>

## Next pages — list / 213100331310 / 5

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface](data-sources--voltstack_site--reference--group-005.md#canonical-2030000031301002-0223312303122312-1112202211103331-0231112012210231-0103330021102203-2221133223310330-0320110230300321-2122302223233010)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2030000031301002-0223312303122312-1112202211103331-0231112012210231-0103330021102203-2221133223310330-0320110230300321-2122302223233010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232323103010003-3322200120011213-2030120333222003-1221223221332031-2031112112320222-3130033221212221-2233020133223101-3301103100223123"></a>

## custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface — interface / 331232103123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [custom_network_config.slo_config](data-sources--voltstack_site--reference--group-005.md#canonical-2012302033131313-0112313333203013-0222312103011303-1202231112003210-0030001133223102-1233101312030000-1112113001223201-3231100201223111)
- [custom_network_config.slo_config.static_v6_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2033023022123002-0002320322031031-2232120013012223-1123311123300231-1022120130322222-3212301020100031-3203310122003212-2131011230121001)
- [custom_network_config.slo_config.static_v6_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-0300130033323031-2021231023213311-2130223301130100-0000031013033001-1010221223110123-0200120320210311-2211132123332213-1013123300223202)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-1011030001223203-3330312030333133-3130223022301223-3232100220110012-1323220113313213-0231130211130312-0310033011031221-3001112032300313)
- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-1110312012212333-0022233200220321-0223330203133133-1300302330003113-3211011210100220-1010133232303031-1201331131110023-2111131312332122)
- custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-3133000130012020-3032201133220211-1032311123231030-2231001210120032-2122130122133023-0101320322212232-3311102112011320-3131333312212022"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-1031333313131221-0313232213122311-1232111213001210-2102321003132202-3123021310212012-2011030112022221-1102302233001100-0020022111101231"></a>

## Direct properties — interface / 331232103123 / 3

<a id="canonical-1120210020202303-3011013110032010-3211203211102102-2033100102310132-2103231321232321-0112123322323330-3221212302021210-1210100001120333"></a>

<a id="canonical-0222301203331033-1113103200321122-1022303031312110-1002313132230103-2322332113020133-0100020032212011-2110303111031111-0010003011300232"></a>

## kind property — interface / 331232103123 / 4

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

<a id="canonical-2030300230313102-0323312112213130-1122310212311113-0002322221123201-0332231202121312-3000203212003311-2322022213220103-2323010001033030"></a>

<a id="canonical-1310111110222200-2011003120033023-1023201122200121-1220211033120003-0230320331112300-1222032310213011-0331130302210132-1130000302000302"></a>

## name property — interface / 331232103123 / 5

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

<a id="canonical-2011330133001200-1122102203022212-2332023302131123-3310211002100303-2113321331300212-3011230211200331-3021301001022300-0022101132011231"></a>

<a id="canonical-1130020030030211-2212032310310120-2022200002011202-3310222320212132-1202122313321313-3223101303330202-2232112000110333-0230312131203100"></a>

## namespace property — interface / 331232103123 / 6

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

<a id="canonical-1231310231032022-1112022131233121-0013313100321122-2201031223303302-1203233212111320-1303133113220001-2003301130102302-2333323030203200"></a>

<a id="canonical-1320121123112103-3321032233121302-1112300032021021-1122301011001022-2121112311002313-2033203110222013-0302001323132230-1310323003101102"></a>

## tenant property — interface / 331232103123 / 7

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

<a id="canonical-3103330200311132-2210110122103310-3312011132021103-1312122213212212-2303233101202212-2022322210111212-0220101021021232-3133321022112212"></a>

<a id="canonical-3121321111111230-0130220212222021-3321203200301022-0321322221302213-0221002213213201-2223202022210030-3110311102312302-3133022030220011"></a>

## uid property — interface / 331232103123 / 8

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

<a id="canonical-0111313020031230-0020030121020232-0232231113210203-1202323112301320-1000211122111330-0210210300002123-2002021330310233-2230010310133120"></a>

## Next pages — interface / 331232103123 / 9

- [custom_network_config.slo_config.static_v6_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-1110312012212333-0022233200220321-0223330203133133-1300302330003113-3211011210100220-1010133232303031-1201331131110023-2111131312332122)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2223213210321030-3311100222230231-2223220012312331-2003322023011310-1023223030223302-1211203202321312-1221021000032010-1203110122102320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031032223321220-1112013032100233-0300233121301112-0130100121131011-2220110000131101-2212023020133213-0131200320232120-2211012221002220"></a>

## custom_network_config.sm_connection_public_ip — sm_connection_public_ip / 333323232223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.sm_connection_public_ip

<a id="canonical-2002033200302012-2123211323133213-1310331333102101-3003230321011202-1111000113213103-2113031100110013-2303012231202213-1002021113320001"></a>

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

<a id="canonical-3201212133021313-2312202121102100-1122312322302133-1121102030012032-2132301332321123-3231100202312230-2332022012230033-0311233011200120"></a>

## Direct properties — sm_connection_public_ip / 333323232223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011203330323121-0131213301211031-0121230010312131-2332003323123311-3203100220331133-3033000000321131-1230110113221213-1132003031030310"></a>

## Next pages — sm_connection_public_ip / 333323232223 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2232310101033032-2012321030300210-2121332101130011-2230300001102330-3302302011123313-0233022301231212-2001000302003323-0233331011133203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210003220112130-3332223133102231-1220303133121133-2230230013120031-1300100020213230-2112110032031112-3033323300121020-0330133023201201"></a>

## custom_network_config.sm_connection_pvt_ip — sm_connection_pvt_ip / 300201333211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- custom_network_config.sm_connection_pvt_ip

<a id="canonical-1101302021300000-2003312223122301-3321201210331120-0113001312021103-2102121100311033-2123111101302033-3312012212123322-2112122021100013"></a>

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

<a id="canonical-1231211310233312-1201311321011002-1212021133302103-1223203212030012-3232220201011223-0031323230320331-1112101033121221-2311003131122102"></a>

## Direct properties — sm_connection_pvt_ip / 300201333211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020103021202010-3313331132323313-2333213322031230-3200313132113303-1211032202033222-2321122112021001-2101332021320012-3312031123023333"></a>

## Next pages — sm_connection_pvt_ip / 300201333211 / 4

- [custom_network_config](data-sources--voltstack_site--reference--group-003.md#canonical-2202232331010233-2312320200121023-3031112203021033-0001003233122021-0021333010311103-1331323331330321-3223202331033020-2123032021001201)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022331210133010-0301000301110312-3012011110221331-0001122030320222-3033323232213202-1112122301330000-0111220003332232-1120321212111333"></a>

## custom_storage_config — custom_storage_config / 000033011200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- custom_storage_config

<a id="canonical-0113000123033200-2201323020302003-2322121132022120-3232320122300213-1323223010031212-3233110301321113-1301110001121000-3131301220021021"></a>

Type: `"single"`. Computed.

\[OneOf: custom\_storage\_config, default\_storage\_config; Default: default\_storage\_config\]
VssStorageConfiguration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage_class\",\"storage_class_list\"]",
  "x-ves-oneof-field-storage_device_choice": "[\"no_storage_device\",\"storage_device_list\"]",
  "x-ves-oneof-field-storage_interface_choice": "[\"no_storage_interfaces\",\"storage_interface_list\"]"
}
```

OneOf alternatives in this subsection:

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-0113000123033200-2201323020302003-2322121132022120-3232320122300213-1323223010031212-3233110301321113-1301110001121000-3131301220021021)
- [default_storage_config](data-sources--voltstack_site--reference--group-008.md#canonical-3112201001132100-1130202120221200-0220303013130011-0011331213301220-0331123030130031-2321023013233332-0300222320113013-0023020100230102)

Select alternatives according to the provider validators above.

<a id="canonical-2101003332301113-0102013001313321-3133001122121023-0123313202311100-2211020003222223-3120030112322301-1130213322331030-1121323313010201"></a>

## Direct properties — custom_storage_config / 000033011200 / 3

- [default_storage_class](data-sources--voltstack_site--reference--group-005.md#canonical-2002213331002100-1221130112230210-1211130213230220-3333211320001323-2302011010210000-1022012300000333-1032211323212331-2111212332300132): complete subsection reference.

- [no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1000012300213203-2230120322230320-2321210333320100-3320011301132330-2223003002201311-0303120020130013-1023200030331212-3220301331122023): complete subsection reference.

- [no_storage_device](data-sources--voltstack_site--reference--group-005.md#canonical-0113131211212100-1103200300333023-3030221301123012-2112133101001313-3031110310313012-2120022030012113-2132213321120100-0132032110333011): complete subsection reference.

- [no_storage_interfaces](data-sources--voltstack_site--reference--group-005.md#canonical-3232202032332123-1300031002223103-2101320001023110-2210313132001113-0103122301012011-3013000021302310-3102312123332311-0212023122330213): complete subsection reference.

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010): complete subsection reference.

- [storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333): complete subsection reference.

- [storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301): complete subsection reference.

- [storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321): complete subsection reference.

<a id="canonical-0130122332322022-2313312100000311-0233203101323323-1213333031303223-1231202230131002-3300033023203231-2322133012120032-0022113310333102"></a>

## Next pages — custom_storage_config / 000033011200 / 4

- [custom_storage_config.default_storage_class](data-sources--voltstack_site--reference--group-005.md#canonical-2002213331002100-1221130112230210-1211130213230220-3333211320001323-2302011010210000-1022012300000333-1032211323212331-2111212332300132)
- [custom_storage_config.no_static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1000012300213203-2230120322230320-2321210333320100-3320011301132330-2223003002201311-0303120020130013-1023200030331212-3220301331122023)
- [custom_storage_config.no_storage_device](data-sources--voltstack_site--reference--group-005.md#canonical-0113131211212100-1103200300333023-3030221301123012-2112133101001313-3031110310313012-2120022030012113-2132213321120100-0132032110333011)
- [custom_storage_config.no_storage_interfaces](data-sources--voltstack_site--reference--group-005.md#canonical-3232202032332123-1300031002223103-2101320001023110-2210313132001113-0103122301012011-3013000021302310-3102312123332311-0212023122330213)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- [custom_storage_config.storage_class_list](data-sources--voltstack_site--reference--group-005.md#canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333)
- [custom_storage_config.storage_device_list](data-sources--voltstack_site--reference--group-006.md#canonical-2202321233200123-3222211221031213-2000231121201303-1330131320331102-1021321302201033-3231012021212032-1123021213103321-1221301223221301)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2002213331002100-1221130112230210-1211130213230220-3333211320001323-2302011010210000-1022012300000333-1032211323212331-2111212332300132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303120001323120-2020103022233021-0230201023320212-3322322132202101-3121213331200110-1033222201103131-2300001103301321-0323121120231213"></a>

## custom_storage_config.default_storage_class — default_storage_class / 123121312123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.default_storage_class

<a id="canonical-0002003213311333-2133003301222332-0000222020003131-1003332003013233-0202001210311332-0310232322200100-3312301031111203-0202303331300232"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default storage class.

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

<a id="canonical-2121323121201033-1122113213230031-3012122002231302-2332230120032222-0311100322231100-0333312001230211-1020123323230332-2100331132222011"></a>

## Direct properties — default_storage_class / 123121312123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000012321123203-2110101303112231-0202122333133312-0030200031110003-2101212223302220-1130320231220132-2323113331112011-2003233221233132"></a>

## Next pages — default_storage_class / 123121312123 / 4

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1000012300213203-2230120322230320-2321210333320100-3320011301132330-2223003002201311-0303120020130013-1023200030331212-3220301331122023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212131021232222-3220000310030302-2031111133002031-0011220102033012-3030203321320010-2333323302010020-0103123020132031-1103120230100032"></a>

## custom_storage_config.no_static_routes — no_static_routes / 331302230220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.no_static_routes

<a id="canonical-3110112322322303-3122010330320320-1100313313203013-0001101212332010-3110313312203113-1230310112002010-3221233012010113-2331103212012102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no static routes.

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

<a id="canonical-0233022221233302-3300301011100312-3031102032321332-0300031101322102-3111131210230132-3231322113020323-1023032131003310-3110323110202101"></a>

## Direct properties — no_static_routes / 331302230220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233000111000320-1303321122002103-3111133002030222-3012313222103002-1210112011113003-2211032320132200-1022030201110312-0230321231113102"></a>

## Next pages — no_static_routes / 331302230220 / 4

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0113131211212100-1103200300333023-3030221301123012-2112133101001313-3031110310313012-2120022030012113-2132213321120100-0132032110333011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311330033012202-2210330103110302-0131230231300320-1200212102123020-2023103121002100-1321332103130200-0033020100302100-0021100030230232"></a>

## custom_storage_config.no_storage_device — no_storage_device / 221331200123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.no_storage_device

<a id="canonical-3101201022321002-0123132010312213-2113210233222321-1302113123002011-2201210031230223-0320202001120130-1302001113022112-1211213012112131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no storage device.

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

<a id="canonical-0303321301102122-3322013223312013-1020220320121312-0113012221200121-2133012010313133-1333333322233032-2111230303310012-2010232301111301"></a>

## Direct properties — no_storage_device / 221331200123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122323302333321-1133302011123111-1201201231201331-0111221100031231-2231113102300213-0013333313320123-1021101232120332-0002230023332100"></a>

## Next pages — no_storage_device / 221331200123 / 4

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3232202032332123-1300031002223103-2101320001023110-2210313132001113-0103122301012011-3013000021302310-3102312123332311-0212023122330213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121322302230220-2013310213221232-2232000110120000-1111223302013122-3003013012220002-2312302103303310-1333033113022320-1120200323132210"></a>

## custom_storage_config.no_storage_interfaces — no_storage_interfaces / 123002113031 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.no_storage_interfaces

<a id="canonical-0011200313101102-2323132032100333-0321123132133303-0220331230332030-0222103321230201-2223112100200320-3031112010303333-3122011020030001"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no storage interfaces.

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

<a id="canonical-2300300022121310-1002330130132301-3331120212022121-0211230000313121-1112223010211001-1200130033030223-1301211331131103-0331323110111301"></a>

## Direct properties — no_storage_interfaces / 123002113031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230332201003202-1132310033232320-0112203121333012-0330132213012211-3322033101320113-2213212120200010-0311133000133210-1211023121010320"></a>

## Next pages — no_storage_interfaces / 123002113031 / 4

- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120223320223102-2221020230320211-1112010133013120-3101320121130012-1102202032101032-1032301000003112-0110131220131331-1230131030121000"></a>

## custom_storage_config.static_routes — static_routes / 210331032000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.static_routes

<a id="canonical-2132030210331231-1231232010231230-3120332020203333-2102031231211103-0122331232320010-0101002133011333-0320201023223313-2133212031010312"></a>

Type: `"single"`. Computed.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

<a id="canonical-1012103303110312-1223030100333001-1313323223010230-1313000031001013-3230331131111102-0131301112000331-0101012222201333-3132301230003033"></a>

## Direct properties — static_routes / 210331032000 / 3

- [static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010): complete subsection reference.

<a id="canonical-0000111202231012-0123102313020000-2220100032222002-1132303000011203-3012103311310022-3231013233320201-1212112022022012-1203313112022323"></a>

## Next pages — static_routes / 210331032000 / 4

- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201320202000301-0230201312301111-2332022230302200-2220133210031112-3223203133022300-2220001320100101-0333211011031032-1000021203212311"></a>

## custom_storage_config.static_routes.static_routes — static_routes / 201030020311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- custom_storage_config.static_routes.static_routes

<a id="canonical-3223203322032002-3333003223231300-0321300211233321-1033200101233333-3032000002131102-0311120031301103-2300311012203310-1130110131121122"></a>

Type: `"list"`. Computed.

Static Routes. List of static routes.

Upstream description:

List of static routes.

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

<a id="canonical-3031022111203132-3211312113220112-1331220310101212-0111111311220332-0031001133012100-1233210233232103-2312310031031102-2120012030102212"></a>

## Direct properties — static_routes / 201030020311 / 3

<a id="canonical-3021123120033202-3213022131300100-3101131322031000-2212303211002322-1301001220302221-1031330021301202-0022301311100223-0210331101121013"></a>

<a id="canonical-1002110222033301-0213122301202122-3302211030110120-3333121110010032-3012302312300120-3122310023212200-2321221123303322-1333111231111021"></a>

## attrs property — static_routes / 201030020311 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of attributes that control forwarding, dynamic routing and control plane (host) reachability.
Possible values are \`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`,
\`ROUTE\_ATTR\_INSTALL\_HOST\`, \`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`.
Defaults to \`ROUTE\_ATTR\_NO\_OP\`.

Upstream description:

List of attributes that control forwarding, dynamic routing and control plane (host) reachability.

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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-3200233132032130-3313102320333332-0002130310300033-3213012230330211-3012133223320020-0020331312102132-3303113322102232-3123312213320021): complete subsection reference.

<a id="canonical-1310332113332323-0212232000202102-2113001131103000-3123021202200301-3130333301321132-0020222030133120-3111211232003213-3213102211211120"></a>

<a id="canonical-1100210333130131-0023212232000310-1203001011000102-1030223010032331-3010012313323221-1021301300010133-1200010011001232-0321312100300302"></a>

## ip_address property — static_routes / 201030020311 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Upstream description:

Exclusive with \[default\_gateway node\_interface\] Traffic matching the IP prefixes is sent to this
IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
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

<a id="canonical-0101212322013210-0220013012330330-1331222330322012-2200200302022023-2131122112223212-3330000310102312-1310012303103310-1303312023003002"></a>

<a id="canonical-0032103103011311-3110322213233321-0301211130031133-1311330132330030-1320123313110323-1223120203002321-0310003001322112-2200113123020013"></a>

## ip_prefixes property — static_routes / 201030020311 / 6

Type: `["list", "string"]`. Computed.

List of route prefixes that have common next hop and attributes.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2231313331131321-0021030102003011-3332302302320100-2001110113323322-1023333221002031-2313330120232312-0011322110011302-1230201333332211): complete subsection reference.

<a id="canonical-0132003121310331-1102032101211021-3313102100132223-3322121232221030-0230232231003202-0113031332303212-1201100330002101-1232232313302302"></a>

## Next pages — static_routes / 201030020311 / 7

- [custom_storage_config.static_routes.static_routes.default_gateway](data-sources--voltstack_site--reference--group-005.md#canonical-3200233132032130-3313102320333332-0002130310300033-3213012230330211-3012133223320020-0020331312102132-3303113322102232-3123312213320021)
- [custom_storage_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2231313331131321-0021030102003011-3332302302320100-2001110113323322-1023333221002031-2313330120232312-0011322110011302-1230201333332211)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3200233132032130-3313102320333332-0002130310300033-3213012230330211-3012133223320020-0020331312102132-3303113322102232-3123312213320021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233220021130021-0032020000230202-0123100023123313-0010002302102210-0222312230102002-2002223123211110-3313223122120101-3211222120010213"></a>

## custom_storage_config.static_routes.static_routes.default_gateway — default_gateway / 313230313001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- custom_storage_config.static_routes.static_routes.default_gateway

<a id="canonical-2000020221120100-0220231030311202-1302310130002322-0013330312222101-2203330022011210-3001302010231322-0200323110112033-3123302213101201"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default gateway.

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

<a id="canonical-1223322120312222-3123233113021221-3230032001220301-1311203312101100-1231023203303013-3023212120120203-0123220001123010-1103131301223031"></a>

## Direct properties — default_gateway / 313230313001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322120122302030-2213120331203331-1100031121113221-0233031003133332-0200313301302013-1023200103013022-0020213110212221-3231223102003212"></a>

## Next pages — default_gateway / 313230313001 / 4

- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2231313331131321-0021030102003011-3332302302320100-2001110113323322-1023333221002031-2313330120232312-0011322110011302-1230201333332211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202120303130121-0021133133220231-1012312023021133-0200102110130032-1233020223022313-1130132103231002-1012310101202333-0020033022011212"></a>

## custom_storage_config.static_routes.static_routes.node_interface — node_interface / 010032112300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- custom_storage_config.static_routes.static_routes.node_interface

<a id="canonical-1012232331310011-3233131011312133-3022322122123003-0222003132100213-0120211132131230-2201001120223023-2223211231203311-1102023010322300"></a>

Type: `"single"`. Computed.

On multinode site, this type holds the information about per node interfaces.

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

<a id="canonical-3313000101310323-2013133112112010-3232232010133133-1322111210003221-1012301110321101-3332021100333310-2033330123023320-2210013032230001"></a>

## Direct properties — node_interface / 010032112300 / 3

- [list](data-sources--voltstack_site--reference--group-005.md#canonical-3021102210101220-1130310001213321-0221323112122033-1021032302030021-0221322231230000-2010303230000331-0230020313220000-3002223122121322): complete subsection reference.

<a id="canonical-0233012102011300-1210132130012000-1131012103020123-2103230130302103-2320023311330111-0301332222112102-3321101300113113-2100232330002032"></a>

## Next pages — node_interface / 010032112300 / 4

- [custom_storage_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-3021102210101220-1130310001213321-0221323112122033-1021032302030021-0221322231230000-2010303230000331-0230020313220000-3002223122121322)
- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3021102210101220-1130310001213321-0221323112122033-1021032302030021-0221322231230000-2010303230000331-0230020313220000-3002223122121322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111322300303003-0012023213102101-3002120012330321-1023133102301200-2013003310333111-2113011013120120-2123231033112332-1031213220201222"></a>

## custom_storage_config.static_routes.static_routes.node_interface.list — list / 102032201020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- [custom_storage_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2231313331131321-0021030102003011-3332302302320100-2001110113323322-1023333221002031-2313330120232312-0011322110011302-1230201333332211)
- custom_storage_config.static_routes.static_routes.node_interface.list

<a id="canonical-0301231002203303-0201232022103201-0201110101203200-0321112101100012-0312032110223101-0130200020212033-0111021232012301-0303100030301022"></a>

Type: `"list"`. Computed.

On a multinode site, this list holds the nodes and corresponding networking\_interface.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.max_items": "8"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8"
  }
}
```

<a id="canonical-0030313002133023-3231322331210121-3313132233300230-0310112323213103-1312330120302313-2032232101223221-1320120230020333-2223103122303303"></a>

## Direct properties — list / 102032201020 / 3

- [interface](data-sources--voltstack_site--reference--group-005.md#canonical-1312320001331200-2033310130011203-0202223113220232-2203211323111200-2101103000010101-1310203022311331-0311021330110013-1113220000303310): complete subsection reference.

<a id="canonical-0032303330220333-0301132300233312-3310001102030001-2110231132331012-3211103310101132-3012302112102233-2213213032203022-1212213003332223"></a>

<a id="canonical-3222120232313013-1112230030312320-0102032331303221-1110110333003203-2200203131101202-3233310332000301-3020213311210302-0203302000301322"></a>

## node property — list / 102032201020 / 4

Type: `"string"`. Computed.

Node. Node name on this site.

Upstream description:

Node name on this site.

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

<a id="canonical-1000321111123312-1322013122131210-2300312022013021-0210332130132202-3200311322332211-0122103100201132-3210330003012023-3321112202122023"></a>

## Next pages — list / 102032201020 / 5

- [custom_storage_config.static_routes.static_routes.node_interface.list.interface](data-sources--voltstack_site--reference--group-005.md#canonical-1312320001331200-2033310130011203-0202223113220232-2203211323111200-2101103000010101-1310203022311331-0311021330110013-1113220000303310)
- [custom_storage_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2231313331131321-0021030102003011-3332302302320100-2001110113323322-1023333221002031-2313330120232312-0011322110011302-1230201333332211)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1312320001331200-2033310130011203-0202223113220232-2203211323111200-2101103000010101-1310203022311331-0311021330110013-1113220000303310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010331211132202-2003111021302113-3033002220122033-2013322102201211-1020121212120101-2203110002131122-1013132233232200-1101203201232201"></a>

## custom_storage_config.static_routes.static_routes.node_interface.list.interface — interface / 130232121300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-2003103223203212-2220313300312200-1023211013303311-2032230330032230-1300111301203321-3210100100300011-3021120312300020-2203131131033010)
- [custom_storage_config.static_routes.static_routes](data-sources--voltstack_site--reference--group-005.md#canonical-1121300311320330-3203003303301323-1111211012002130-3022132301030301-1321310133020012-1102033310111133-0202120213112020-1132132303100010)
- [custom_storage_config.static_routes.static_routes.node_interface](data-sources--voltstack_site--reference--group-005.md#canonical-2231313331131321-0021030102003011-3332302302320100-2001110113323322-1023333221002031-2313330120232312-0011322110011302-1230201333332211)
- [custom_storage_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-3021102210101220-1130310001213321-0221323112122033-1021032302030021-0221322231230000-2010303230000331-0230020313220000-3002223122121322)
- custom_storage_config.static_routes.static_routes.node_interface.list.interface

<a id="canonical-2210010330011313-2101123330033212-3220102203002211-3113332023201102-3121201223003310-2311021311001101-3310210102030321-1202322131321300"></a>

Type: `"list"`. Computed.

Interface. Interface reference on this node.

Upstream description:

Interface reference on this node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-2111313331022310-0231020121120102-0302202313321033-2011301011103323-2203201303033031-2220112332120103-2000121021010330-1000101021201122"></a>

## Direct properties — interface / 130232121300 / 3

<a id="canonical-0011011321232121-3233211223033221-1130210020223000-1202220311020112-2202030331203200-1332121003130322-3302300012313220-2113013321320030"></a>

<a id="canonical-0311132332230113-1200000221032102-2333221201333230-3130030122313311-1311300303000030-2033133012013230-1121122030131311-3113313103230321"></a>

## kind property — interface / 130232121300 / 4

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

<a id="canonical-2321033320303320-1003212021113303-2002322200331310-0323003202010030-3021222220213212-2332031122201221-0030132001332300-3311321113133221"></a>

<a id="canonical-2203030323022210-3111313232133333-2210220031230001-2202110310332110-1022230031201031-3113313210103131-1120232320122102-3220323031130130"></a>

## name property — interface / 130232121300 / 5

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

<a id="canonical-0213311223201320-2133311101120312-2330320022100011-2033001303231332-3132331013101131-0033301230323032-3313120032320323-2223122311103211"></a>

<a id="canonical-1031002323012203-0222011100201202-0200232000000322-2000030000323200-0331320310113322-0303122032221230-1203200122030033-0213030011022201"></a>

## namespace property — interface / 130232121300 / 6

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

<a id="canonical-2101123102000133-0233123110123110-0133100033323211-2033301210030122-0200311032123333-2102212210103011-1011213101201301-3331103010202112"></a>

<a id="canonical-0000031031033002-2320131303212333-3000210303131320-1230222033233303-3301313020031003-1002013103101101-2211300000321222-0123130202030123"></a>

## tenant property — interface / 130232121300 / 7

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

<a id="canonical-1211012301311301-1320011333033030-1321132202113201-3133032213132211-2002113311103123-1020233031333300-2301302113233132-1022021213313333"></a>

<a id="canonical-3131233302332332-1113202223133113-1232323030203301-1211333312121233-2230300221213132-3132012300012012-1112013110001030-0321233332313021"></a>

## uid property — interface / 130232121300 / 8

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

<a id="canonical-3030201023302203-1111012223322123-0332212222223330-1213020323310110-3113100010302103-1023202121233212-0000102303100101-1332200101123203"></a>

## Next pages — interface / 130232121300 / 9

- [custom_storage_config.static_routes.static_routes.node_interface.list](data-sources--voltstack_site--reference--group-005.md#canonical-3021102210101220-1130310001213321-0221323112122033-1021032302030021-0221322231230000-2010303230000331-0230020313220000-3002223122121322)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0313302010302212-0330323323030201-2223132001112311-1033010323032321-2200130111100323-2131003310231011-1023031012213200-2333011201031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221002110012130-3332301023100121-0301121331202123-1233300110001012-0022222022332221-3131200301011131-3022220312302022-0202112200111230"></a>

## custom_storage_config.storage_class_list — storage_class_list / 331302022300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- custom_storage_config.storage_class_list

<a id="canonical-3333232223232121-3322020001131233-0201112023111313-0230302322031111-1313213331121131-2132230320302002-3223221231203202-0022012310100331"></a>

Type: `"single"`. Computed.

Add additional custom storage classes in Kubernetes for this fleet.

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

<a id="canonical-1102022302331222-1302000022321033-2301133210022031-3220032201033023-1231131133310030-0023032312021131-0110033020130210-3301321233331213"></a>

## Direct properties — storage_class_list / 331302022300 / 3

- [storage_classes](data-sources--voltstack_site--reference--group-006.md#canonical-2310120202333033-1102122022233230-0203312132320232-0120230001210010-1202232103101111-1231121222201002-3123003130120232-1201133030100332): complete subsection reference.
