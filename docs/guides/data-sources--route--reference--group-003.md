---
page_title: "xcsh_route reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_route reference."
---

# xcsh_route reference

<a id="canonical-0221033032231121-1102232113320301-1322100011103323-2122313313033132-3123200232332301-2111303232332233-1112312030302022-3132113313201310"></a>

## replace_params property — route_redirect / 003323003320 / 8

Type: `"string"`. Computed.

Exclusive with \[remove\_all\_params retain\_all\_params\].

Upstream description:

Exclusive with \[remove\_all\_params retain\_all\_params\]

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
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2102333210131021-1012211112331000-0212111313222313-0122033201002133-1332113322023102-3102030000211222-3223233301203022-2121222332120001"></a>

<a id="canonical-2312330030203331-2001333312010330-0212032010023003-3100211321233230-0012213002303322-3223333320101210-1003033013102322-1323320330101111"></a>

## response_code property — route_redirect / 003323003320 / 9

Type: `"number"`. Computed.

The HTTP status code to use in the redirect response.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 599,
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
    "ves.io.schema.rules.uint32.lte": "599"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "599"
  }
}
```

- [retain_all_params](data-sources--route--reference--group-003.md#canonical-3001133033010312-3112110313130000-0103031123012010-2010002101212033-1201320120320022-3223103111033311-0023212131113320-3031130220030310): complete subsection reference.

<a id="canonical-2222232130233300-1303233011220003-3122110230213321-1012122203111321-3203100120303012-0033231233200310-1033321311003122-1020302101023300"></a>

## Next pages — route_redirect / 003323003320 / 10

- [routes.route_redirect.remove_all_params](data-sources--route--reference--group-003.md#canonical-3212312301222212-3230213212111021-2210000103300211-2132001001121130-0122030121113210-2220232312012230-0232002211300323-3033133132030320)
- [routes.route_redirect.retain_all_params](data-sources--route--reference--group-003.md#canonical-3001133033010312-3112110313130000-0103031123012010-2010002101212033-1201320120320022-3223103111033311-0023212131113320-3031130220030310)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-3212312301222212-3230213212111021-2210000103300211-2132001001121130-0122030121113210-2220232312012230-0232002211300323-3033133132030320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200201020212033-3003101333000101-0133133233013210-1020020013231023-2130331012301130-0001233200201031-2031000312013021-1211320020223121"></a>

## routes.route_redirect.remove_all_params — remove_all_params / 332100221002 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_redirect](data-sources--route--reference--group-002.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103)
- routes.route_redirect.remove_all_params

<a id="canonical-2032120322132110-1000112332320112-0310031222110120-2202130021203033-2122032200112200-2233012000301212-2303221033021232-1011122110332110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for remove all params.

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

<a id="canonical-2303011321102210-1102100311332221-3211300200312210-0302320112033201-1211003002001002-1223132222320302-0211133033202233-0103200211200113"></a>

## Direct properties — remove_all_params / 332100221002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211212131221021-3121130031230130-1332020131223300-3120323202301231-2032131113013311-2313311321010122-2213213011133002-0303220023133100"></a>

## Next pages — remove_all_params / 332100221002 / 4

- [routes.route_redirect](data-sources--route--reference--group-002.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-3001133033010312-3112110313130000-0103031123012010-2010002101212033-1201320120320022-3223103111033311-0023212131113320-3031130220030310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320220331003103-0301311012220101-3310111212120200-3201121011023302-2232210123121032-2330101233333320-1210303310212303-2030101111302110"></a>

## routes.route_redirect.retain_all_params — retain_all_params / 310033120130 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.route_redirect](data-sources--route--reference--group-002.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103)
- routes.route_redirect.retain_all_params

<a id="canonical-2033321200101322-1032113232103113-3000112022122030-3230131201303010-1010030033011233-1121222123100323-1133313020021332-3231311103130002"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for retain all params.

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

<a id="canonical-0201222001203030-0310001222111122-0013200010333230-1321002211132201-1001100201021021-0231220020102112-1031311232033110-0312200113023203"></a>

## Direct properties — retain_all_params / 310033120130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121000301031323-2221322010303120-0200203200010012-3113023022223032-1020222223033233-0211311131100323-0013232011102133-3003123213021001"></a>

## Next pages — retain_all_params / 310033120130 / 4

- [routes.route_redirect](data-sources--route--reference--group-002.md#canonical-3031003311233213-2113111021311003-2111031332300130-1111311130312100-2120221111012231-3331330300223001-2023020031123100-0201231213001103)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-0122003031122022-2312121320313011-3232203000312111-2300213223121003-3333300031112101-3031102112131312-1230313311203332-3230301311131320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323133310312121-3322323123322013-3201212310020031-0323300022202200-3222123201032003-0201230103231213-2003202012000331-3312303132121011"></a>

## routes.service_policy — service_policy / 020201031032 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.service_policy

<a id="canonical-0120211110302033-0112203233032231-3310003010012113-3123123220311133-2311331120231020-2023310002002100-1322033322100030-2211011113112032"></a>

Type: `"single"`. Computed.

ServicePolicy configuration details at route level.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-service_policy_choice": "[\"disable\"]"
}
```

<a id="canonical-3310100210201301-3120103212121020-1303321123232212-0203211312132030-1022002221020212-1013321303230222-1313101112102211-2301312000133213"></a>

## Direct properties — service_policy / 020201031032 / 3

<a id="canonical-0103020030002230-3110133023123232-1123201313310303-1223212303110230-0012032210122132-3301011300031101-1201321010100311-2022202212311012"></a>

<a id="canonical-0333030103323200-1321033231321132-3312012130032102-0000031032231001-1102003202222100-0300212032101333-0321303013131203-0022100203122000"></a>

## disable_spec property — service_policy / 020201031032 / 4

Type: `"bool"`. Computed.

Exclusive with \[\] disable service policy at route level, if it is configured at virtual-host
level.

<a id="canonical-3310130230031233-1231321110120222-0030313023012020-3122322301101320-3030121203320223-3012123333231330-1301302023003222-3103321302321313"></a>

## Next pages — service_policy / 020201031032 / 5

- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-2130023122111302-3312131310223330-3111030320001222-1300131123332221-1300033010300132-3331301213032123-3013023033221232-0223231010031303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323213331032302-2113131003213323-2003312221021201-3101022231113230-3220023100320030-1200312331120310-1100203331333333-2203022031113102"></a>

## routes.waf_exclusion_policy — waf_exclusion_policy / 030323202023 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.waf_exclusion_policy

<a id="canonical-0102021320210111-2132312231202321-3013023220302010-0003020210112022-2020200210101302-2333300233103113-1131321320110012-3110003001331232"></a>

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

<a id="canonical-1321130301212031-3110233303012101-1311232222112121-2012203233111212-0322001231300122-3022112002032303-1000103221301122-1013310103020313"></a>

## Direct properties — waf_exclusion_policy / 030323202023 / 3

<a id="canonical-1330111020203003-3313231120300213-0133202300130220-1233231202100320-1133331011211112-3122010203130312-2113030030332303-0002123321223002"></a>

<a id="canonical-2211321001323232-0312033310211133-0022320321013110-3232320221010031-0232123102200310-2322230132300001-1021000232310311-3011203031210210"></a>

## name property — waf_exclusion_policy / 030323202023 / 4

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

<a id="canonical-1230223032103331-2131013110202030-1010203010332101-3321203033111002-3113222200013311-1210031301322010-1233201322223302-1330022223222001"></a>

<a id="canonical-0213300320232021-2201212232120131-0112220223203211-0131013010222330-1223121130312010-1311033231012232-2020132001313311-3112322300033131"></a>

## namespace property — waf_exclusion_policy / 030323202023 / 5

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

<a id="canonical-0001332303212101-0002322132021222-2310330202221010-1221121012011120-1103321233231200-1120033111330313-2231210122233322-1123101111220111"></a>

<a id="canonical-3302213302310003-0013101232320120-0331233030201102-3322230202112222-2032022011323301-3123101022021331-1222201010101133-1130222022130000"></a>

## tenant property — waf_exclusion_policy / 030323202023 / 6

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

<a id="canonical-1220003110211333-2313231111112332-0301323112310321-3111300301133210-3132100323103032-3233110231010122-2131002002031000-0022000220131000"></a>

## Next pages — waf_exclusion_policy / 030323202023 / 7

- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232122022211110-0333010202033102-0220223213100022-0231030012102130-3030110211320002-1021130203002300-1033003000211211-1102232313232232"></a>

## routes.waf_type — waf_type / 000202301002 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- routes.waf_type

<a id="canonical-0000203232210111-1312212321222112-1133232323000121-0223302032303213-3021033121313200-1330222203301112-2110020213030231-1033301211332320"></a>

Type: `"single"`. Computed.

WAF instance will be pointing to an app\_firewall object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-ref_type": "[\"app_firewall\",\"disable_waf\",\"inherit_waf\"]"
}
```

<a id="canonical-1321212130120230-3211213133302012-0200122303211123-3320113202113333-0221111110120001-2111230100301110-1210002301011132-3000222032013333"></a>

## Direct properties — waf_type / 000202301002 / 3

- [app_firewall](data-sources--route--reference--group-003.md#canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332): complete subsection reference.

- [disable_waf](data-sources--route--reference--group-003.md#canonical-1320201023113001-3121120311102013-3030312122111212-1202223212111100-3103200212220332-1223330020120011-3320011033130113-2103131322333023): complete subsection reference.

- [inherit_waf](data-sources--route--reference--group-003.md#canonical-2101320330322032-3311333110132211-3223101103110030-0020132323000112-0110211120000303-3223313302010332-2211030221221303-2131323020032203): complete subsection reference.

<a id="canonical-0321313133001301-3310001023132301-2020130212222113-3112002221321313-1132000010011020-3103300030103101-3122021323222200-3302011300302023"></a>

## Next pages — waf_type / 000202301002 / 4

- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332)
- [routes.waf_type.disable_waf](data-sources--route--reference--group-003.md#canonical-1320201023113001-3121120311102013-3030312122111212-1202223212111100-3103200212220332-1223330020120011-3320011033130113-2103131322333023)
- [routes.waf_type.inherit_waf](data-sources--route--reference--group-003.md#canonical-2101320330322032-3311333110132211-3223101103110030-0020132323000112-0110211120000303-3223313302010332-2211030221221303-2131323020032203)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211223333032003-1222213131103301-0000000133000333-1021112230220030-2032321103223302-2311231311311032-2113023301123112-2000331013122322"></a>

## routes.waf_type.app_firewall — app_firewall / 121312210331 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- routes.waf_type.app_firewall

<a id="canonical-0020231331220323-3031331033100200-1120333103130222-1211103132113331-3010320123323121-0313201322312212-2010031112310211-0330130310103330"></a>

Type: `"single"`. Computed.

List of references to the app\_firewall configuration objects.

Upstream description:

A list of references to the app\_firewall configuration objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3131103113213211-1222313222211331-0301201202131311-3322310323303332-2013121002312023-2233321230021222-2320013133310133-1103221303220000"></a>

## Direct properties — app_firewall / 121312210331 / 3

- [app_firewall](data-sources--route--reference--group-003.md#canonical-1131313111311232-3120310130010230-3312112323200010-3011033332233100-0032331002023211-0101101233000110-0120203210121320-0222031013231001): complete subsection reference.

<a id="canonical-3002130032003233-2100201101110102-1032312211332120-0210301131223211-0330101023012301-1303211210122211-3212133110021133-1311223121103301"></a>

## Next pages — app_firewall / 121312210331 / 4

- [routes.waf_type.app_firewall.app_firewall](data-sources--route--reference--group-003.md#canonical-1131313111311232-3120310130010230-3312112323200010-3011033332233100-0032331002023211-0101101233000110-0120203210121320-0222031013231001)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-1131313111311232-3120310130010230-3312112323200010-3011033332233100-0032331002023211-0101101233000110-0120203210121320-0222031013231001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122222012302220-1012223020231323-1112302002323003-2121131200320303-2111003301333023-1130130221003001-2013202333203300-0032001221100110"></a>

## routes.waf_type.app_firewall.app_firewall — app_firewall / 301013023223 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332)
- routes.waf_type.app_firewall.app_firewall

<a id="canonical-3333133223101013-2213111111313231-0012020331031330-0223031321232131-3303022200200131-0330201332023012-1100001222032202-2202330011023331"></a>

Type: `"list"`. Computed.

References to an Application Firewall configuration object.

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
    "ves.io.schema.rules.repeated.num_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1"
  }
}
```

<a id="canonical-0123103021321102-1332022201111122-2231131200322022-3031302000001201-0323031233332122-2111101321122033-1321310010111010-1112301200103221"></a>

## Direct properties — app_firewall / 301013023223 / 3

<a id="canonical-2121213330030331-2200332111100100-3121131311211211-3010133023032333-1101300010220200-3232333101212131-1302231301032323-1131031110231221"></a>

<a id="canonical-2232313020012230-2102013101013321-0200023202130220-3230012113321033-2103331012000322-3323223220131221-0110003013112110-3011222112233111"></a>

## kind property — app_firewall / 301013023223 / 4

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

<a id="canonical-3110120212000211-0221110303020000-0330132003200110-1302313121021130-1011330120000230-2121122333010111-1021321310322203-1120333002303213"></a>

<a id="canonical-1310222233010113-0323321222003312-0222200133033110-2223211203111223-3121010202123021-3303233200111322-0232021123202300-0122033031022320"></a>

## name property — app_firewall / 301013023223 / 5

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

<a id="canonical-2112212222101111-2302011323133332-1030330011000032-3312203323222110-0322230312033001-1333032110211100-1132322323130120-2132302230320110"></a>

<a id="canonical-2301212020301102-3323321000213332-3112123010002213-3211201213333211-2312330110201132-0100020331031231-1321302000200331-1330020000132023"></a>

## namespace property — app_firewall / 301013023223 / 6

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

<a id="canonical-0232130132102221-1230300033232101-1300023002221300-3223300100330332-0230020201323120-0331232021101021-1102021222311321-3211122332022312"></a>

<a id="canonical-2020000022302000-2121110032003000-3112003001023232-3122312013210011-2113111330210023-0313122111133300-2030111100022230-1232210233030003"></a>

## tenant property — app_firewall / 301013023223 / 7

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

<a id="canonical-2300302123331103-0220313323130323-2210230032300213-3123223001320032-1213003303311121-1132023111130221-2200123202333313-2320300033301002"></a>

<a id="canonical-2030132110132131-3121011010012030-3213332103020233-2323210200311010-3121020011033322-2021213300232001-0200012100121310-2123332233332313"></a>

## uid property — app_firewall / 301013023223 / 8

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

<a id="canonical-3302333211210322-1132021231020111-0321230003021000-3031231102022101-3330103302132002-1122111202021010-2313232100001201-2002032200321031"></a>

## Next pages — app_firewall / 301013023223 / 9

- [routes.waf_type.app_firewall](data-sources--route--reference--group-003.md#canonical-2020000300332202-2313022220012333-0031313002211002-1203003321210130-2133210011020131-2022331332130112-3122033002213300-3020310233103332)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-1320201023113001-3121120311102013-3030312122111212-1202223212111100-3103200212220332-1223330020120011-3320011033130113-2103131322333023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121013331031103-3301312132032202-0131023230333332-2313330101311230-2021210110210203-1330303032330203-2113120203200033-0112330120331313"></a>

## routes.waf_type.disable_waf — disable_waf / 321211221102 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- routes.waf_type.disable_waf

<a id="canonical-1311301330230310-3322000023323223-0231222322120003-0133003210322311-1030011031130121-2212310322320133-3203333030201113-1313023003101221"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable waf.

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

<a id="canonical-2121120202331313-1030111101030033-1032103122021310-1032121331321321-1002002100213031-3031333201310302-0020333031110211-1113222012201333"></a>

## Direct properties — disable_waf / 321211221102 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0020021230021212-2320002202300101-0022211013000303-3203313210110033-3330032313102313-3333202033232111-2033200132033122-2321211132101011"></a>

## Next pages — disable_waf / 321211221102 / 4

- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)

<a id="canonical-2101320330322032-3311333110132211-3223101103110030-0020132323000112-0110211120000303-3223313302010332-2211030221221303-2131323020032203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033222133111212-3332201010331330-0131113131030011-3332033132002131-2003223101002023-0101012102200210-2221103201001202-0001030220222201"></a>

## routes.waf_type.inherit_waf — inherit_waf / 010230221020 / 2

Breadcrumbs:

- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
- [Property reference](data-sources--route--reference--group-001.md#canonical-2212301313311221-2313200310322120-2131100211322003-0220011200012202-1102120133130302-0220302222322112-3232312213201301-3123012131332333)
- [routes](data-sources--route--reference--group-001.md#canonical-0121210202130211-2311023332311110-1102321013231023-2132030003031313-1021032033031212-3122211102213330-2203313311310212-3201022220331021)
- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- routes.waf_type.inherit_waf

<a id="canonical-0113313123212030-3112020123111331-0123333101202012-1330102303203232-1011233123333203-2330313020102010-0120222220011000-0213232212330100"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inherit waf.

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

<a id="canonical-3121123311010322-1122020211121211-0322013122233323-1122212310220032-2333332331202102-2310322311220303-2211020112102200-3120111021221130"></a>

## Direct properties — inherit_waf / 010230221020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120330123321321-1303303211220132-3103301211111120-2212030312112001-3021131002230113-1332322023122132-1133131330011001-2213133233102023"></a>

## Next pages — inherit_waf / 010230221020 / 4

- [routes.waf_type](data-sources--route--reference--group-003.md#canonical-1303120232330332-3020313221101000-0130321002101012-1313102131100000-3013131220122231-1130022213201220-0013210022302201-0010200133301120)
- [xcsh_route](../data-sources/route.md#canonical-3222100232013222-1030330100002202-0031320130331110-1011111201002221-2030012120011111-3301321100122112-3131030011011213-1132121023303013)
