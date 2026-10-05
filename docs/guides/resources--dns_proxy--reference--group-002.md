---
page_title: "xcsh_dns_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_proxy reference."
---

# xcsh_dns_proxy reference

<a id="canonical-0003300001221301-0233101223313210-3231101323220312-0031212222232301-0200111001212212-3223022213013032-2212123220222032-1011213202110001"></a>

## proxy_advertisement.advertise_custom.advertise_where — advertise_where / 020221123112 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-3003223223300310-3011211031123323-1201331300020311-1200221300002310-1212102011313022-0133233033320023-2032300332031222-1201323323221333"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_dualstack_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "advertise_v6_on_public"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("advertise_v6_on_public",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("port",
    "port_ranges"),
  validators.ConflictingListObjectAttributes("port",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("port_ranges",
    "use_default_port"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_network",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "virtual_site_with_vip"),
  validators.ConflictingListObjectAttributes("virtual_site",
    "vk8s_service"),
  validators.ConflictingListObjectAttributes("virtual_site_with_vip",
    "vk8s_service")}
```

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

Terraform syntax:

```terraform
advertise_where {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232313222320012-1321122021101330-0001321200300000-3020003130223002-1231113322312122-3100232301330032-3213202022230100-0123231312031113"></a>

## Direct properties — advertise_where / 020221123112 / 3

- [advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313): complete subsection reference.

- [advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303): complete subsection reference.

- [advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323): complete subsection reference.

<a id="canonical-1130132230110332-0321221103110123-1232002310000133-0231111301020130-3313133233232222-0200023030020230-3220013132333001-3132312310322112"></a>

<a id="canonical-2030303033030210-1133030010223321-2022001320100011-2133331120321210-1032103200001001-3022200323301310-3131300101022310-1313333201123013"></a>

## port property — advertise_where / 020221123112 / 4

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Upstream description:

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3033003112303210-0110122122110233-1123031230332003-0211010200133120-0221230301003030-0311033232101012-1101200131002012-1311022033000223"></a>

<a id="canonical-3213222313000313-0102312321101231-3013031312123000-0021332001123331-1232020002333130-1112310223121203-0212113010011030-2323113212131003"></a>

## port_ranges property — advertise_where / 020221123112 / 5

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by "-".

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 512),
}
```

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

- [site](resources--dns_proxy--reference--group-002.md#canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023): complete subsection reference.

- [use_default_port](resources--dns_proxy--reference--group-002.md#canonical-2112301032313320-2121103122013203-2030031013320313-0002330002202003-2330120231030121-1130322300222121-3013201211233313-2002013122031333): complete subsection reference.

- [virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232): complete subsection reference.

- [virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210): complete subsection reference.

- [vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333): complete subsection reference.

<a id="canonical-0023120312133212-0331133221220322-0030333020313210-3020311013302322-2211113123130023-2210233131233030-0111203100310331-2221221133133102"></a>

## Next pages — advertise_where / 020221123112 / 6

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023)
- [proxy_advertisement.advertise_custom.advertise_where.use_default_port](resources--dns_proxy--reference--group-002.md#canonical-2112301032313320-2121103122013203-2030031013320313-0002330002202003-2330120231030121-1130322300222121-3013201211233313-2002013122031333)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021132320102121-2133133330111110-3102323120312002-2102310132023132-2330113120212132-3222133010311003-0101001110323112-3313102221222033"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public — advertise_dualstack_on_public / 131213312213 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-2213121001221203-2023133202223113-3100312111202303-3220233321331230-2213113221011002-2310132211312121-2201212333230202-0130311213022110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232123312131001-3121030030120212-0311110000301321-3013030013332303-1003003301232000-2331123003202130-1022020321122331-2033001312000123"></a>

## Direct properties — advertise_dualstack_on_public / 131213312213 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-0113000330020233-1322203022320013-2100203102010001-0130233121233201-1202030312212132-0201103131322201-0032031332323000-3111001223011203): complete subsection reference.

<a id="canonical-1330102101120303-3322231213010103-3310100230011203-1301130013202202-0202132103302120-1202112020212300-1021123330133303-2132203213033211"></a>

## Next pages — advertise_dualstack_on_public / 131213312213 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-0113000330020233-1322203022320013-2100203102010001-0130233121233201-1202030312212132-0201103131322201-0032031332323000-3111001223011203)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-0113000330020233-1322203022320013-2100203102010001-0130233121233201-1202030312212132-0201103131322201-0032031332323000-3111001223011203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113233313011331-0231302211321003-2301230122130010-3031223301132131-0332233200000102-3232232123121120-2132233010332101-0103313221102000"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip — public_ip / 132121322100 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-0101002010331332-0033103031331232-1312121201100222-0020210021311301-1233320003333001-3301011112310231-3203012033213210-0013123330300321"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003200301322103-0211302232311233-3220111100102332-2103111323110312-0220223112003301-1103130300012313-0223023210012212-2311203033330231"></a>

## Direct properties — public_ip / 132121322100 / 3

<a id="canonical-0101002110311020-1001021310031220-2320202330103030-2101100111223322-2302031312221303-3011230113120211-2211130020003201-1101230030001103"></a>

<a id="canonical-1213303002121203-2331330301232330-2133002022322213-2121323023313122-3201221013332331-1322000231330131-0022213100303321-0302203312133120"></a>

## name property — public_ip / 132121322100 / 4

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

<a id="canonical-1030323332322320-2021200012012203-3332111303333123-0210301101133002-2001231002300022-1013311323022303-2303210021010023-1202323333302122"></a>

<a id="canonical-1310023020313230-2013232033121033-0002020022220300-3313031112231212-2131130121130330-0310112022321111-3023001311000123-0010010200100003"></a>

## namespace property — public_ip / 132121322100 / 5

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

<a id="canonical-2302112001123302-0121212112132102-1031121001313002-0002321223032133-3031121103001002-3330300130223121-0230300003212123-0120332130322332"></a>

<a id="canonical-1201302231012010-0021310220331210-2123031112333302-1111112023122011-1123312112102003-2100122010313302-3310321023200120-3222111222322020"></a>

## tenant property — public_ip / 132121322100 / 6

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

<a id="canonical-1113233331030200-0202201032331022-0221123330233201-3122332201230012-1200330321020003-0003103323223021-3330210331103132-0131333200132110"></a>

## Next pages — public_ip / 132121322100 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-1210302231200223-3200232113022130-0220223102302232-3000323111130131-1101121322121130-1100201211301122-3202031033020112-3030021320132313)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021011223112020-1301123111131013-3103212211120113-3030210330003100-0112322221333112-0131003320031012-0112131312212010-3233321320110100"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public — advertise_on_public / 012100203102 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-2301300312103302-0312310200132131-1313010010301011-1323023113302020-1110311312303132-2221122200022310-2032020303211003-3122233120202031"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1330203303300203-1120222113113133-0302302133222121-3131001232321023-1132021323121223-3331120202333311-0123311113303010-2303202111303332"></a>

## Direct properties — advertise_on_public / 012100203102 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-3113231113312211-0320001121010222-2001203011233030-0333120020303233-3302101011020303-0012012331100231-2222310323021102-2233313122321223): complete subsection reference.

<a id="canonical-2203023011123012-2023100010320231-2103031330203001-3112002331310232-3001013200321121-1301312010012233-3013201323022113-1013020120202201"></a>

## Next pages — advertise_on_public / 012100203102 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-3113231113312211-0320001121010222-2001203011233030-0333120020303233-3302101011020303-0012012331100231-2222310323021102-2233313122321223)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3113231113312211-0320001121010222-2001203011233030-0333120020303233-3302101011020303-0012012331100231-2222310323021102-2233313122321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303102332202021-0312122010030333-0130110231313022-0333013311113133-3120003123131310-2002031320112231-0320331112011032-3230113211101102"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip — public_ip / 303100210233 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-1111130223301320-1101111331101331-1300133030201233-2212331321111303-0133100220231212-0023302213121000-0302213313022322-3211021320203112"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203012031232300-3231103001001101-3002201023103130-3133313322112203-2213120030222002-3330233311221331-3333203200121121-1302012000100111"></a>

## Direct properties — public_ip / 303100210233 / 3

<a id="canonical-3231101301213200-2231022302002202-0230230233232333-3202200331023123-2030111222313233-3120210002203011-3001313330003001-3313003202001113"></a>

<a id="canonical-1233032112203203-1020233233132223-0210332312210211-1020211000131023-3023222010200110-0313133231322020-2103320322120310-1200102212011203"></a>

## name property — public_ip / 303100210233 / 4

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

<a id="canonical-2202202230001320-1332000303213111-0132213132310300-3231222033100212-3211220121220011-2003013311113103-1221022330300122-0313233220011333"></a>

<a id="canonical-2012211122230021-0010100220203030-1100222302112003-3100031233202223-0312333003131312-3013320303033112-1332303030322213-1221330223302013"></a>

## namespace property — public_ip / 303100210233 / 5

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

<a id="canonical-3323202010223220-1302223012100233-0120231330303101-3222232213210133-3102023123031323-3021332112001022-1210332332311232-1013113131111231"></a>

<a id="canonical-1131213322322312-0003113011220030-2121330120130131-1300223000203201-1222100023300021-0223211330110332-2320101203310033-3003000313023232"></a>

## tenant property — public_ip / 303100210233 / 6

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

<a id="canonical-0001111110233330-0313133111001310-1301001101031002-0310012120211322-1233223022223020-3332222213131020-3321100002020010-2001323130023311"></a>

## Next pages — public_ip / 303100210233 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2330201123331113-3300203000032202-3113132212311201-3233330303301201-0121303030001130-3030122132211110-3013222131331201-0131020100202303)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302310120211210-0130210230333003-1300102230331003-3103333031330301-3000113210230212-0221021011220333-1303322001303230-3100013100111320"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public — advertise_v6_on_public / 312322303002 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-2212301320331322-0123000231200203-1111121232133032-1032121112310320-0330020003110130-3112110310131020-1200311003003121-0231322301330021"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132113331120101-0333120021031233-0100132313303323-3030030203112220-0220301312023230-1123133100233020-2332223011011012-2331112022033132"></a>

## Direct properties — advertise_v6_on_public / 312322303002 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-3201202033131101-0031331311311313-0310023021330133-2101300100231130-0221321121001222-1011011310332110-3030033111232320-2010233012213031): complete subsection reference.

<a id="canonical-1101331130132332-3312033023301220-0122031232210231-3233113111003230-1111010211011122-1323300320311103-3001001002101000-2123120230122130"></a>

## Next pages — advertise_v6_on_public / 312322303002 / 4

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-3201202033131101-0031331311311313-0310023021330133-2101300100231130-0221321121001222-1011011310332110-3030033111232320-2010233012213031)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3201202033131101-0031331311311313-0310023021330133-2101300100231130-0221321121001222-1011011310332110-3030033111232320-2010233012213031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3123000130332032-0312333201123103-3320002121020201-0032122321220120-3132032013122220-1321130323213221-3111331112331333-0232303100033333"></a>

## proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip — public_ip / 103231100211 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-2221201302323300-3333022111320103-2211320000312123-1102131012110121-0033020020232232-2213320101321122-2101111211211011-2323310003012000"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313221121222121-3203113310101300-1113202012330101-1023021100030022-0000102113020003-1210011321200112-0231112233101300-3332103121120232"></a>

## Direct properties — public_ip / 103231100211 / 3

<a id="canonical-3201010212011110-1101233311202011-0202000000123211-2133332101113121-0121222122221120-1333201021002302-0133122020322231-3211001033233100"></a>

<a id="canonical-3220213310231303-2122320113320003-0301032301132131-0220223200333330-3010130103332133-1210211113123122-0202311313110200-3133310130100313"></a>

## name property — public_ip / 103231100211 / 4

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

<a id="canonical-3102122002113122-3300300131120122-1113122333230200-2302010010011131-1131221322111010-2310330201003212-3332202133022003-1321301132100023"></a>

<a id="canonical-2103020102220333-2133030320133021-1000201013103110-3200032213110033-0333330010001223-3330210012010000-1131311013331310-0131111220022111"></a>

## namespace property — public_ip / 103231100211 / 5

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

<a id="canonical-2020300013132010-0233030301132132-3121212223130211-1000323302303121-3002233022302032-2311201203111032-2030010310310003-1022031331113003"></a>

<a id="canonical-1020100202031002-1023020101102211-1221022113130020-0323010021022203-3303000332202331-2031223112130023-0232033313122103-2233300000211113"></a>

## tenant property — public_ip / 103231100211 / 6

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

<a id="canonical-0223020123302023-1023211111000130-1303132121203022-3220132213313003-0210032230323203-2312013200013122-0302100100023112-3233311003233230"></a>

## Next pages — public_ip / 103231100211 / 7

- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-1120333322121220-3021032102210113-0133220302303212-3220302231211022-1231132230302013-3013301002003330-1210223200002302-3213023211201323)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0033313133220203-2213030131322303-2013022321102122-0031121002003232-0100023020033301-3003302101022212-2321310031130220-1322111131201230"></a>

## proxy_advertisement.advertise_custom.advertise_where.site — site / 133313220320 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-2313301120110301-3332023233032012-1201123231210211-2012022203130103-3213311223123101-0212222202132033-2000330003321103-2221302003022032"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0323303221220220-2332102110000131-2121123031312311-2102210202130213-2302221223223022-3232113001203103-0113320132011022-1332102021202010"></a>

## Direct properties — site / 133313220320 / 3

<a id="canonical-3201111200013122-0102203133302121-3032230313013310-1123302131020131-0030211101103103-0132213022132003-3220202302223332-3010122021002122"></a>

<a id="canonical-2031111010101010-1111023003303121-2121021023211300-3320023013231231-0033320311123013-0300330230333031-0213120311021012-2101033313203332"></a>

## ip property — site / 133313220320 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-1101031132033003-1110311021030231-1320231223211013-2002220113303210-3211201230003320-1301320032012110-3001310132330301-0033120312332011"></a>

<a id="canonical-3331312201021033-0103220120313302-0113330222013213-2032323330332211-2221302211112311-1030001000002202-1123112333032132-2130202113120003"></a>

## network property — site / 133313220320 / 5

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [site](resources--dns_proxy--reference--group-002.md#canonical-1121030032133221-0203020103230201-2311002322003103-2203232022211100-1133333333013223-2220010301221110-1323320012020221-0123303220133132): complete subsection reference.

<a id="canonical-1113021203102222-2131012300001312-0221233323202202-3221022100331100-3111233211011130-2120310001033303-0032231001022000-3030210202200310"></a>

## Next pages — site / 133313220320 / 6

- [proxy_advertisement.advertise_custom.advertise_where.site.site](resources--dns_proxy--reference--group-002.md#canonical-1121030032133221-0203020103230201-2311002322003103-2203232022211100-1133333333013223-2220010301221110-1323320012020221-0123303220133132)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1121030032133221-0203020103230201-2311002322003103-2203232022211100-1133333333013223-2220010301221110-1323320012020221-0123303220133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021313232332332-0200023032111022-2332210212100010-1123010003001230-0200001111320122-2130223101031123-1331100232113102-2211201312103330"></a>

## proxy_advertisement.advertise_custom.advertise_where.site.site — site / 213011331301 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-3121012021321022-1030012101200203-3130002231220300-1110031221000032-2203321213002333-1011121123133032-0232203123300102-2233132331232310"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023211221103031-3002020332220330-2101223020303031-0122012230221303-3120110220102102-2020313003212101-0203010021200230-3320221021310323"></a>

## Direct properties — site / 213011331301 / 3

<a id="canonical-1313223333010210-2012223303213231-3023002210222033-3201203101313220-3232223323323312-1211020201033010-1103102132333020-0313212220230210"></a>

<a id="canonical-3013123110022201-1130233133302310-1212320302333311-1103223100112303-0210303000321312-2223033103203020-2113320213300232-0102213222022033"></a>

## name property — site / 213011331301 / 4

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

<a id="canonical-1231323310120323-1232220000320223-1111022232022230-1321321310220332-2310202310032330-3322310010001023-0001323201023231-2312110103212232"></a>

<a id="canonical-0113331210100313-0013220230311001-3223123001003322-0231303210100120-1113203103122213-1031101003323011-1213300110003113-0313202210321033"></a>

## namespace property — site / 213011331301 / 5

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

<a id="canonical-0100000320101122-0000210011223003-2133220103122002-1202320200133322-1011221010112223-2303312121101122-0102211031033322-3030130323231133"></a>

<a id="canonical-2230101021113332-0123313212332333-0223032112332320-1111222021100231-3133001102132100-1332120103202202-3102222310032211-3011021001003220"></a>

## tenant property — site / 213011331301 / 6

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

<a id="canonical-3223022123311301-2022031121011212-0002002222000301-1203001232220010-2011311210300120-1111213301201010-0212332010321113-3231333301333032"></a>

## Next pages — site / 213011331301 / 7

- [proxy_advertisement.advertise_custom.advertise_where.site](resources--dns_proxy--reference--group-002.md#canonical-0333220002313122-1230023033012223-0011103000031311-1102313210121022-3120111101003332-0300211233200311-2023302120013213-3202030012302023)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2112301032313320-2121103122013203-2030031013320313-0002330002202003-2330120231030121-1130322300222121-3013201211233313-2002013122031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220120113201310-0323332203303113-2312202221300022-0201123130310300-2211112032203130-1220332021121120-1013101022210133-3013232130021111"></a>

## proxy_advertisement.advertise_custom.advertise_where.use_default_port — use_default_port / 120320221222 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-1133113321203103-3101030303002133-1022220023012200-1012022331030232-0211212130232033-2332001201010201-3320012112230023-2210110123321301"></a>

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
use_default_port = {}
```

<a id="canonical-3010300011303001-3230231021113212-0001033323220020-0301011300230103-0112323332233310-1111212233031313-1131233122102110-1222121102022221"></a>

## Direct properties — use_default_port / 120320221222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300302000021023-2120111233010321-0121303210332130-3111301312212021-0012120111222122-0311113301111323-2021121021312211-0231312321031032"></a>

## Next pages — use_default_port / 120320221222 / 4

- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312123332210011-2123321111303033-2211013112322122-1200003210132303-1031103133030000-1112011121120211-2330022301320032-3201030310320212"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network — virtual_network / 331013032222 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-3022012000222210-1003230222202323-3003020331002222-2220021330120300-0113301010130111-0100211103333212-0303320203012333-1121320112020023"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_v6_vip",
    "specific_v6_vip"),
  validators.ConflictingObjectAttributes("default_vip",
    "specific_vip")}
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
  "x-ves-oneof-field-v6_vip_choice": "[\"default_v6_vip\",\"specific_v6_vip\"]",
  "x-ves-oneof-field-vip_choice": "[\"default_vip\",\"specific_vip\"]"
}
```

Terraform syntax:

```terraform
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110010012123031-0101000102210021-1133320300100222-1023222301321211-0001332211211031-2302233011310000-3100022231133033-0121032213313120"></a>

## Direct properties — virtual_network / 331013032222 / 3

- [default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-1220122210013202-3110010111110212-1322100322101333-0033320123132330-2301100013223132-0313233301201313-2001210202121310-0211123203323300): complete subsection reference.

- [default_vip](resources--dns_proxy--reference--group-002.md#canonical-2203223331301230-3330320130220010-1011311223200330-0230333310220212-1123301033122322-1311022221221232-2302221010123022-3321230233223303): complete subsection reference.

<a id="canonical-3110212102132200-3120232310321203-2322303302230101-3130230123100321-1000222030302202-3201130033020301-2003010211213011-0003112111133110"></a>

<a id="canonical-0031201321211112-2213002000120122-3032012130022233-0223203101113310-0003002332123332-1222211213000323-0032011002300022-1200113033221110"></a>

## specific_v6_vip property — virtual_network / 331013032222 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv6Validator(),
}
```

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

<a id="canonical-1230000013311031-2131303130201300-1003313310130233-2121111000210002-0102300220101132-2320122211322330-1231310112321213-2001310200000012"></a>

<a id="canonical-1213311120201013-3212003102213301-0303130021313221-1221033302233021-3002233013033331-3202202320113313-1003111321113001-3300123101203333"></a>

## specific_vip property — virtual_network / 331013032222 / 5

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Upstream description:

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

- [virtual_network](resources--dns_proxy--reference--group-002.md#canonical-1110121310200022-1331323031201213-0110022323020313-2000103120103200-2003221011311312-3231202332120111-1312320311132020-2321212031032320): complete subsection reference.

<a id="canonical-3333202131202113-2032223122102331-0010133001302210-0301133012320023-1011020133333001-2230030100020312-2201213312303200-1123011000300000"></a>

## Next pages — virtual_network / 331013032222 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip](resources--dns_proxy--reference--group-002.md#canonical-1220122210013202-3110010111110212-1322100322101333-0033320123132330-2301100013223132-0313233301201313-2001210202121310-0211123203323300)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip](resources--dns_proxy--reference--group-002.md#canonical-2203223331301230-3330320130220010-1011311223200330-0230333310220212-1123301033122322-1311022221221232-2302221010123022-3321230233223303)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-1110121310200022-1331323031201213-0110022323020313-2000103120103200-2003221011311312-3231202332120111-1312320311132020-2321212031032320)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1220122210013202-3110010111110212-1322100322101333-0033320123132330-2301100013223132-0313233301201313-2001210202121310-0211123203323300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303313300031321-1233112231230332-1301230223023300-0113102023311033-2313013233123023-3013301003022000-1122333132120031-0301131001210302"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip — default_v6_vip / 111100012033 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-2203220020003122-1223331223330131-1213210021002211-1331300220113020-2320333000310102-3131200201120111-2213013021320012-2310133311103010"></a>

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
default_v6_vip = {}
```

<a id="canonical-0030033200220031-3210212121310002-1223230030121203-3021000133313103-0130012222232222-3131130011021301-3233120202122000-2333221002302030"></a>

## Direct properties — default_v6_vip / 111100012033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233222100120030-3312121031120331-1320212032132031-1010201030210320-1203120033321213-0030120301200030-3201201231031331-3220120111131213"></a>

## Next pages — default_v6_vip / 111100012033 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2203223331301230-3330320130220010-1011311223200330-0230333310220212-1123301033122322-1311022221221232-2302221010123022-3321230233223303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333322131332030-2301212333002103-3030232023211003-0300231231030222-1101000131210211-2101300321103001-1023220000311130-3022210122013012"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip — default_vip / 132110010221 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_vip

<a id="canonical-3002013123313002-2002221302033222-2120010212131001-1113202112223310-1223122033121211-2003231332122221-2010302100230332-2020102010213031"></a>

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
default_vip = {}
```

<a id="canonical-2232203122301301-2110021101312121-0210011233220111-3000320212023200-3210020132013200-2313023101212021-3212113132011020-1322120011002233"></a>

## Direct properties — default_vip / 132110010221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002110231200012-3300301121101102-1232021113200210-1211233301220230-3023300302211223-1211132031213311-2123133212000033-1213110110101031"></a>

## Next pages — default_vip / 132110010221 / 4

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1110121310200022-1331323031201213-0110022323020313-2000103120103200-2003221011311312-3231202332120111-1312320311132020-2321212031032320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120021031031203-3221033331230022-1003032112231222-1131221320302233-2103013301301203-2012212023023200-2201123032120113-1021300030332332"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network — virtual_network / 120310001131 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.virtual_network

<a id="canonical-0222213132202020-1121330003202230-2031323101023201-0012200222211231-2201122200130000-2031211101010131-2121002312020001-2010333323321223"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-1011203132203230-2002212233213023-3111003312202203-3130031313203133-3211231030101202-2310311020132333-2332220121100133-0131312301030113"></a>

## Direct properties — virtual_network / 120310001131 / 3

<a id="canonical-3000030120100023-2221123130313031-2233311332003232-0002031013010222-3312213033231011-3011100011222330-2101100112231230-1230121003332212"></a>

<a id="canonical-0331013220202331-0213111233011002-0021301121002212-3322200330031201-3121023122003211-3333133212301021-0202213331321112-1212300231221331"></a>

## name property — virtual_network / 120310001131 / 4

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

<a id="canonical-1213201320010032-2323020212020103-0021021333011123-0011330010320222-1331222122213122-1110132200221330-1303303033121313-2201310210311213"></a>

<a id="canonical-2301311133023200-1102300222211010-0300200032003230-3201000230100230-2320211120323312-2312003233032133-0010231220211203-3030120111213131"></a>

## namespace property — virtual_network / 120310001131 / 5

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

<a id="canonical-2202003331322310-2223003132110232-3000010131230310-1332011033030300-1223232303013021-1030010201221032-3223001101011003-2131113113313200"></a>

<a id="canonical-3302030332213301-3212323310010222-3323012323110121-0220330223332013-0103320220201122-2132211321112210-1131223330322213-1313323311001232"></a>

## tenant property — virtual_network / 120310001131 / 6

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

<a id="canonical-1332032122121220-1002233221010011-2313011213033133-1231210133001101-0332203321221312-3301103303212310-2331212022020210-2111121130010113"></a>

## Next pages — virtual_network / 120310001131 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--dns_proxy--reference--group-002.md#canonical-2310102012101212-1313110121233032-0222012221233332-0332003311302302-3001031232200013-1030223210030220-3330302311103220-0202103201323322)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333002001323311-2120101300003321-3220221232102210-3300103132213213-0122221321020202-2303203110310131-3223032033323310-1333023211020201"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site — virtual_site / 201011302322 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site

<a id="canonical-2000331302020031-0312102230303301-0303023121001132-0321022210120003-3313010333322232-0213101331032233-1332020130010321-2122023120111232"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301112210331131-2232102332212022-0103130111133031-0033031100231130-2311230010220310-1223011221003230-3213003110221211-0230033203103112"></a>

## Direct properties — virtual_site / 201011302322 / 3

<a id="canonical-1113203110221003-1131202132321021-0233311312113320-3101003021013022-2002201311231113-3323232113113300-3203323133303200-2003333221311120"></a>

<a id="canonical-1022101013332003-0012021231232110-3000123003233211-2002221311011112-1211232120022013-3332131220302121-3203110312321320-3313222113001203"></a>

## network property — virtual_site / 201011302322 / 4

Type: `"string"`. Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_INSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE","SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_IP_FABRIC","SITE_NETWORK_OUTSIDE","SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP","SITE_NETWORK_SERVICE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_INSIDE_AND_OUTSIDE",
    "SITE_NETWORK_INSIDE",
    "SITE_NETWORK_OUTSIDE",
    "SITE_NETWORK_SERVICE",
    "SITE_NETWORK_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_INSIDE_AND_OUTSIDE_WITH_INTERNET_VIP",
    "SITE_NETWORK_IP_FABRIC"),
}
```

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

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2030312330101200-1000002022223202-2111300132122030-3332310200013100-0313000330011321-2023221300331202-2230322112122203-3111000001231200): complete subsection reference.

<a id="canonical-0230130320333330-1000013233322122-2232132031031300-1212220112001211-1301103133013303-2112222000300201-2111023121131233-0131231022202310"></a>

## Next pages — virtual_site / 201011302322 / 5

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2030312330101200-1000002022223202-2111300132122030-3332310200013100-0313000330011321-2023221300331202-2230322112122203-3111000001231200)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2030312330101200-1000002022223202-2111300132122030-3332310200013100-0313000330011321-2023221300331202-2230322112122203-3111000001231200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100100001303013-1010311313222211-1103311123012202-3031022202001220-0323200001111023-0133222110232203-1210213032133022-1011302133011212"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site — virtual_site / 102312323021 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site.virtual_site

<a id="canonical-0021100002030003-2312230001232001-1122222010311010-0330101233331201-2212002313313212-1022301200032232-0321210320221323-0200300030302020"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300123321002111-2023002300110302-1300210200331300-1313131320320320-0322111120222222-1213231122131103-2123023213320302-3323323302012000"></a>

## Direct properties — virtual_site / 102312323021 / 3

<a id="canonical-3020311221030122-1202003213021230-3131032021021123-1023202122213202-0112122213201330-1200010010221332-0211121231213322-3111110322011133"></a>

<a id="canonical-3311113113112300-3113030001220121-2110003323112231-1202320111120110-2222130102002000-3000212113001101-2022133331011303-2112313301012012"></a>

## name property — virtual_site / 102312323021 / 4

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

<a id="canonical-0013220031130313-3013001031010213-2221032300333201-2012030233223110-2301003302233122-3333203233303012-1322210120302322-1132001132021111"></a>

<a id="canonical-3121002030110111-1022011002103310-2221322300331022-2120013132111212-2110101013131112-1003303303311230-1310111112231213-1023320031130012"></a>

## namespace property — virtual_site / 102312323021 / 5

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

<a id="canonical-0103203210211310-2021030133010203-1121110111130102-0333113032013112-0232113332232031-0303302210312110-0312013303012330-0302203133303123"></a>

<a id="canonical-2131131320221021-0033003001201200-3121312203012110-2101121112103312-2001002102331011-0330001023321213-0200121010320123-0012323220120023"></a>

## tenant property — virtual_site / 102312323021 / 6

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

<a id="canonical-1303032222303020-2322120120023333-3102130202330120-3110223123301231-2112200030212322-2233232321010223-0100312113021122-3022230302332111"></a>

## Next pages — virtual_site / 102312323021 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-2121010310201023-0012033031302111-2200021322330221-2310000013030030-0102122021131301-1220332330110022-3322322113031011-0121112303002232)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120231210231310-1302030132202223-1312333300131023-1132122021123123-1312323101000200-0121120233233133-1233133300032310-3311130020022131"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip — virtual_site_with_vip / 301221201203 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip

<a id="canonical-2322000121112300-0011031322121120-3021203331300001-3022021012123321-3132222132112302-0123301231010300-3031032201200212-0333213323230201"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site_with_vip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3031003200212101-1002210012031031-0101033133323223-3130031331332202-3301013321300320-3023302301303102-1330111310302113-1200223231021121"></a>

## Direct properties — virtual_site_with_vip / 301221201203 / 3

<a id="canonical-2132210203323022-2002010100023111-3312332103331320-1131203323222103-2312132320020230-0111211230011121-2230222232010302-1001133313323231"></a>

<a id="canonical-0021013131202132-3300213002230030-1110203220013231-3021130101213330-0332012113233203-3112203010210033-3022230213333222-0013130012113303"></a>

## ip property — virtual_site_with_vip / 301221201203 / 4

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPv4Validator(),
}
```

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

<a id="canonical-2200023110002111-3233330110110202-0310101130032002-1223121223231232-3012003001121201-0231330012333000-1101000133210332-3332221232103223"></a>

<a id="canonical-3200031131102200-1203233221013001-1101332312023003-0211233121131232-3103222312200303-0003013233201030-1223013213131023-1332303033221012"></a>

## network property — virtual_site_with_vip / 301221201203 / 5

Type: `"string"`. Optional.

\[Enum: SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE|SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\] Defines
network types to be used on virtual-site with specified VIP All outside networks. All inside
networks. Possible values are \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`,
\`SITE\_NETWORK\_SPECIFIED\_VIP\_INSIDE\`. Defaults to \`SITE\_NETWORK\_SPECIFIED\_VIP\_OUTSIDE\`.

Upstream description:

This defines network types to be used on virtual-site with specified VIP

All outside networks. All inside networks.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["SITE_NETWORK_SPECIFIED_VIP_INSIDE","SITE_NETWORK_SPECIFIED_VIP_OUTSIDE"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("SITE_NETWORK_SPECIFIED_VIP_OUTSIDE",
    "SITE_NETWORK_SPECIFIED_VIP_INSIDE"),
}
```

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

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-0312003011223001-1303020320021031-3321123013210122-0210122130020001-1123202233010322-3010132221110111-3120202331321130-1223310232330323): complete subsection reference.

<a id="canonical-0013010003232131-2212330312003200-2012221103113312-3201111123113033-3313210223300122-1212330003232003-0033203233020232-2100202102120221"></a>

## Next pages — virtual_site_with_vip / 301221201203 / 6

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-0312003011223001-1303020320021031-3321123013210122-0210122130020001-1123202233010322-3010132221110111-3120202331321130-1223310232330323)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-0312003011223001-1303020320021031-3321123013210122-0210122130020001-1123202233010322-3010132221110111-3120202331321130-1223310232330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011031223010111-2321031112203212-3120333310123212-2010030323210311-3323132013121200-1101012032123211-3012220101201311-0013120222323221"></a>

## proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site — virtual_site / 102001231300 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210)
- proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip.virtual_site

<a id="canonical-3133103110122322-3103021010203231-1100232131203130-0030320020132023-1101302110231321-3311101130112220-0113111310313232-1332312100121212"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320302133201032-0230003213003221-2000013032133103-2213132213321132-1111221000231010-1323200330212202-0033120222030020-1022211123003131"></a>

## Direct properties — virtual_site / 102001231300 / 3

<a id="canonical-2320122331023032-1102021312221201-1022013013303200-3030120130123203-0330203120331201-1102103331203103-2003001032031310-1320202001333232"></a>

<a id="canonical-3101113212101031-0210002201210023-0021200313100302-1310021300212333-3220312032022202-2123201220003003-1233303123033222-1320030031103222"></a>

## name property — virtual_site / 102001231300 / 4

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

<a id="canonical-1011122203221030-1123113212013023-3322120131223133-2123123031031331-0122203112213131-3301123121211030-2223133212310030-2120202120310203"></a>

<a id="canonical-3221031221313123-2202120111002100-1130301001310131-0331233223202022-1330200210210120-3133012230231321-1221311000303033-2132320311101221"></a>

## namespace property — virtual_site / 102001231300 / 5

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

<a id="canonical-1021133303223230-2201001230213313-1112001120332232-1321330111301111-3202320310003300-3023223333232321-0310200333311020-0012303210323222"></a>

<a id="canonical-1330203301131001-1333330121221330-0322303301222332-2001211200033321-1231002000201112-0123003233023023-3330320000031201-1232320000010221"></a>

## tenant property — virtual_site / 102001231300 / 6

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

<a id="canonical-0011133231131332-0031313123022131-2113233033032031-2031333212022233-3101232111331212-3313203020103310-3023123002030300-1210013220223221"></a>

## Next pages — virtual_site / 102001231300 / 7

- [proxy_advertisement.advertise_custom.advertise_where.virtual_site_with_vip](resources--dns_proxy--reference--group-002.md#canonical-3120210322311313-1031233122013201-3010103131230133-1220113202032033-0022231001022221-1121230121113102-2203301120120222-1130010230032210)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331221331131300-0233311021021333-1033003210312233-0211130102023103-1113311011103233-0301132111130031-0312103010330322-2311133101131003"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service — vk8s_service / 000020300231 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service

<a id="canonical-0130112030213102-0210233031303222-3021132321111132-0112031102232131-3010302133123322-0202211020200131-2330222113123311-0032222201313303"></a>

Type: `"object"`. single nested block, Optional.

Defines a reference to a RE site or virtual site where a load balancer could be advertised in the
vK8s service network.

Upstream description:

This defines a reference to a RE site or virtual site where a load balancer could be advertised in
the vK8s service network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_site")}
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
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
vk8s_service {
  # Configure direct properties listed below.
}
```

<a id="canonical-0213333123112112-1021011220102111-0230111321013213-2203323003221301-2302130200333322-2021321323113011-3211331220100033-0120112220033332"></a>

## Direct properties — vk8s_service / 000020300231 / 3

- [site](resources--dns_proxy--reference--group-002.md#canonical-3320333020331332-2201010102332312-0320303013333121-1132002321120201-2300021202011330-3010011302033130-0213133123001033-2020303200012230): complete subsection reference.

- [virtual_site](resources--dns_proxy--reference--group-002.md#canonical-3103101011303203-2022103310012021-3001313311211213-3322321020321310-0220133030010133-1112333000112311-3301302213102023-0231221312201301): complete subsection reference.

<a id="canonical-2222301022320130-1333131230011101-2320131301302311-2333100330301030-2100121023220200-2222112121311120-1323100112313320-0222112230012331"></a>

## Next pages — vk8s_service / 000020300231 / 4

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site](resources--dns_proxy--reference--group-002.md#canonical-3320333020331332-2201010102332312-0320303013333121-1132002321120201-2300021202011330-3010011302033130-0213133123001033-2020303200012230)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site](resources--dns_proxy--reference--group-002.md#canonical-3103101011303203-2022103310012021-3001313311211213-3322321020321310-0220133030010133-1112333000112311-3301302213102023-0231221312201301)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3320333020331332-2201010102332312-0320303013333121-1132002321120201-2300021202011330-3010011302033130-0213133123001033-2020303200012230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021223112032133-0302230102303310-3100030032320321-1103011011332131-3013130300210120-0323330022132333-0202313330330211-2102212011000102"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site — site / 213212233010 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.site

<a id="canonical-2001123130003223-3102010022300233-1000113102130232-3002030030222122-2333113202212310-2233230001012130-1131333012200223-3011013031031102"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032311113132200-2230111103110201-3120203220110123-2220023102302230-2101321300220113-3231211020101201-2333322021132211-0313233031232202"></a>

## Direct properties — site / 213212233010 / 3

<a id="canonical-0001221132022202-1222203001212323-3211203320130021-2131031222033322-1212101201101313-0012112101013012-1032032122330303-0302313133212012"></a>

<a id="canonical-1000211130232130-1121332200301103-1100121021100223-0100221323200211-1210111021202301-3112112313211210-1212112330102312-3001230112020212"></a>

## name property — site / 213212233010 / 4

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

<a id="canonical-1030111032030311-1010221012023113-3010030331100301-3110133120232013-3332000133321222-3213321122302233-1003311131301231-0320010323102123"></a>

<a id="canonical-2232300112200012-0331011301220001-1200213132033102-3331212122230132-3101120101320111-0111111000133310-2113020231232222-1330320111020302"></a>

## namespace property — site / 213212233010 / 5

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

<a id="canonical-1330200132333031-0302212303001322-1303022012212031-0313203120323233-3210211020321100-2300031121332202-2113331302021211-2301122231333021"></a>

<a id="canonical-1031310100112302-1122202113320102-2001332131100323-2112200210233121-2103131310111133-0301231230133113-1332300033122222-0121201202311211"></a>

## tenant property — site / 213212233010 / 6

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

<a id="canonical-3220210310021010-2202323212203332-1100103303211322-1013120311322221-0102313030321301-1202011302123031-0210022102211100-2000312103231233"></a>

## Next pages — site / 213212233010 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3103101011303203-2022103310012021-3001313311211213-3322321020321310-0220133030010133-1112333000112311-3301302213102023-0231221312201301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222003223332312-1022331203312321-0322113231310000-0212221113121200-0123331330111331-1033013002003000-2003313202221320-1201211200013130"></a>

## proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site — virtual_site / 221002200111 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_custom](resources--dns_proxy--reference--group-001.md#canonical-3230203022210102-3133112122202023-3302231132330322-2301230313021012-3230302133100123-2111102202222022-1313331122103301-3232330002222101)
- [proxy_advertisement.advertise_custom.advertise_where](resources--dns_proxy--reference--group-001.md#canonical-0212102131211322-1322200121200133-0333010223030332-3311012120012333-3022033203310011-3231023331133130-3312020033312232-0222321301332220)
- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- proxy_advertisement.advertise_custom.advertise_where.vk8s_service.virtual_site

<a id="canonical-0220022230213123-0211200020100231-3333333112201112-3000110013002000-2203230312011032-0311013030310231-2212331121000202-1330233300102112"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-3033010011300331-2123322022123001-1221001320112233-3123023000030111-3232333113021032-1132232213301121-1012122331030233-0132311232310330"></a>

## Direct properties — virtual_site / 221002200111 / 3

<a id="canonical-0020112322220221-3202032303112203-0202110203223033-2032020031200100-0120202232310132-1122110130303202-2031212203011100-2300112220022021"></a>

<a id="canonical-2012302310303302-3100011031300130-3231202302120002-1322133333320222-1221101203010110-1210312100122300-2302033322032233-3213320023133131"></a>

## name property — virtual_site / 221002200111 / 4

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

<a id="canonical-0232011111001203-2020010031301132-1212112131222022-2313313221321230-0331111331103203-3003203131313330-2033212132120232-3021223233113030"></a>

<a id="canonical-3133331130211233-0230203103100212-1212012220113220-1203102030310013-0033210331301001-0003031222100322-3102021111100021-1030210321210323"></a>

## namespace property — virtual_site / 221002200111 / 5

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

<a id="canonical-0003331312302013-3110210300130033-2303030030031110-1210200012231303-3013110103323223-1003020133322202-1123101122331232-3213101110333123"></a>

<a id="canonical-1120021201212100-1233132122331111-2230013320000022-3033200101022323-1001312233132331-3132010123031100-2100223102011320-1332221201112023"></a>

## tenant property — virtual_site / 221002200111 / 6

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

<a id="canonical-0333312000100011-3230302212013001-0010321230201002-0301002002322223-2131112231300101-2312330023112221-3200312023032010-1223211330020011"></a>

## Next pages — virtual_site / 221002200111 / 7

- [proxy_advertisement.advertise_custom.advertise_where.vk8s_service](resources--dns_proxy--reference--group-002.md#canonical-3331020230023233-2321211121222102-3023033122212123-2101133133002010-1023311023330020-2220002321331011-0202230202222210-2111222102020333)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-0312321001033003-2000102210132110-2013301302220000-1320023122223111-1132300310003200-1203011203201332-2211111321321133-0310222130312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331230202133022-0010212013011332-2001110222010200-1120021020023222-2000222022011201-1212200320001113-0033112000010323-0310023300033032"></a>

## proxy_advertisement.advertise_dualstack_on_public — advertise_dualstack_on_public / 201210320122 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_dualstack_on_public

<a id="canonical-3020020013121233-3031311033231313-0030211030221130-1123003331111033-0111302233303300-1221131132300211-2232120103231112-3122030121232213"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_dualstack_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2213132230233011-3002300232002121-1021220013321130-3231222012211113-3022332202321022-2302011112030230-1012303011012111-2231121332133023"></a>

## Direct properties — advertise_dualstack_on_public / 201210320122 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-2232002212321132-3033031231032131-3032131102301100-2222201012313302-1013332023321100-0010330303230101-0100001000300303-2231202322201223): complete subsection reference.

<a id="canonical-0022132302330220-1203033001120223-0332001102020232-3112021300201133-3001223112100000-2103220220231221-0223033213022021-2232012211013000"></a>

## Next pages — advertise_dualstack_on_public / 201210320122 / 4

- [proxy_advertisement.advertise_dualstack_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-2232002212321132-3033031231032131-3032131102301100-2222201012313302-1013332023321100-0010330303230101-0100001000300303-2231202322201223)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2232002212321132-3033031231032131-3032131102301100-2222201012313302-1013332023321100-0010330303230101-0100001000300303-2231202322201223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012010132112233-0222030112010330-0203013331211301-1022010123001021-1331123113130223-1101101223022013-1233032020223303-2333120331221100"></a>

## proxy_advertisement.advertise_dualstack_on_public.public_ip — public_ip / 030321122310 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-0312321001033003-2000102210132110-2013301302220000-1320023122223111-1132300310003200-1203011203201332-2211111321321133-0310222130312302)
- proxy_advertisement.advertise_dualstack_on_public.public_ip

<a id="canonical-2203212130210320-2323202330103120-0202202120201120-3320313010330321-1301111200321033-1010121211013230-3210212103332222-2002311303202022"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333322301020231-0011333112323210-2322101213032031-0301030310220010-2320300111101302-1233321220232323-3212133331120110-1132231021001312"></a>

## Direct properties — public_ip / 030321122310 / 3

<a id="canonical-2221313111301231-0111330101312110-3211320220030333-2000131202113000-1031330103121303-3321320231010021-0230310222220113-1220131023312000"></a>

<a id="canonical-3001231111322111-2323023312223230-2102133113211103-2023112302131112-2323130031022310-0300310002113021-0312023213100200-3003112132120111"></a>

## name property — public_ip / 030321122310 / 4

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

<a id="canonical-2300101312112133-3031003200211203-2100223222111003-2133032111210030-3120201122220132-1230102223320023-1113203231300323-1302301113211133"></a>

<a id="canonical-3123323321321301-0312323022111230-0023102031111221-0323331223301003-0110313221011331-3021302220321131-1202103201001223-1132233133302221"></a>

## namespace property — public_ip / 030321122310 / 5

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

<a id="canonical-3022011221103222-3320122312311232-2222231213323312-1023320232231113-0212213020022200-3300200113133311-1120320021023333-0000030333122013"></a>

<a id="canonical-0030303302022011-2110323111303233-2001301201020220-1023110002331222-3121020232312333-1002122311102002-2223122133302302-2312320332323231"></a>

## tenant property — public_ip / 030321122310 / 6

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

<a id="canonical-3230132002002311-1021013101220201-0131100012030022-1223202233330302-2302011231111311-3010023210013203-2113233222300110-0210210113020221"></a>

## Next pages — public_ip / 030321122310 / 7

- [proxy_advertisement.advertise_dualstack_on_public](resources--dns_proxy--reference--group-002.md#canonical-0312321001033003-2000102210132110-2013301302220000-1320023122223111-1132300310003200-1203011203201332-2211111321321133-0310222130312302)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2220303323213103-2110110021231302-1232003120023021-3003122123002120-0021020320000023-3232200110132031-0222010121032220-2103301122232110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102110221103131-0033101323212213-2111111232230231-2123130221121103-0332320132110033-3221112003203131-3230122200232003-0121210012032000"></a>

## proxy_advertisement.advertise_on_public — advertise_on_public / 103333123333 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public

<a id="canonical-1110322033101122-3003022311212312-0022132232200202-1221003100230212-2223131233220202-0031012212233122-3122230130002102-0322200120311221"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-2220213321300220-0333033301201330-3120010322202133-3102331302220013-1122302031201121-1101110103213202-2210231022200123-1212210001211113"></a>

## Direct properties — advertise_on_public / 103333123333 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-1211001132230302-3112320133313021-3310213102223200-3233220301232120-0223012021000110-2023102323120002-3220210033113312-1112323130102201): complete subsection reference.

<a id="canonical-2321210311023101-2010211132120002-3232232232303303-3022300110333033-3200023310121211-3021222022323103-2330212121302122-0321202210232011"></a>

## Next pages — advertise_on_public / 103333123333 / 4

- [proxy_advertisement.advertise_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-1211001132230302-3112320133313021-3310213102223200-3233220301232120-0223012021000110-2023102323120002-3220210033113312-1112323130102201)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1211001132230302-3112320133313021-3310213102223200-3233220301232120-0223012021000110-2023102323120002-3220210033113312-1112323130102201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001002122012120-3221020100012311-1211211113223110-3311230311210233-3020112000210121-2033021023322012-2023132001003122-3122221113113011"></a>

## proxy_advertisement.advertise_on_public.public_ip — public_ip / 110202123103 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2220303323213103-2110110021231302-1232003120023021-3003122123002120-0021020320000023-3232200110132031-0222010121032220-2103301122232110)
- proxy_advertisement.advertise_on_public.public_ip

<a id="canonical-3001021030310113-2101133213311130-3331323200230030-1121232033133013-3220300001300301-3321313313311032-3132220030033003-1333201122131330"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2012102202101201-0331003020010120-1301210331320010-1123001120200110-0013212322321213-2231021012303303-3110011022023303-0120131321313301"></a>

## Direct properties — public_ip / 110202123103 / 3

<a id="canonical-2313202322003112-1102301011313203-2311320032022130-2221103210232303-3023232120302031-3223220122231203-1123113301121000-2230020203320313"></a>

<a id="canonical-2131330111112002-3130330130112001-1023122200110121-2321230022032310-3000233132103133-2020122332301013-3332002213313030-3221002033110310"></a>

## name property — public_ip / 110202123103 / 4

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

<a id="canonical-3100121110222121-0102320103011120-0003323132001220-0321302022202230-1320111201310311-2023332123332223-2330313023331301-3211001112033302"></a>

<a id="canonical-1033210320132103-3133100221112122-0213123202022333-0312201231020030-0223012320103223-3021002110023323-1203202121031302-3320312100032311"></a>

## namespace property — public_ip / 110202123103 / 5

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

<a id="canonical-1102230123021113-3230233331221020-1313003133033123-0033330311122023-0230211001222132-3032131130013123-3133023020123313-3111320003231111"></a>

<a id="canonical-0222130220100203-1230313021303213-2220031020100232-1313300012211203-1120120331033000-3222220003313020-3021230323232231-1211010130021231"></a>

## tenant property — public_ip / 110202123103 / 6

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

<a id="canonical-3021032313023231-1031011131100322-0033330200000222-2300120110232013-1033130130332020-0001313303220332-2300212000101020-0213302131012221"></a>

## Next pages — public_ip / 110202123103 / 7

- [proxy_advertisement.advertise_on_public](resources--dns_proxy--reference--group-002.md#canonical-2220303323213103-2110110021231302-1232003120023021-3003122123002120-0021020320000023-3232200110132031-0222010121032220-2103301122232110)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3121000120202100-0310113221300320-2102133012003031-0222233000032201-3001032020133122-2001303133003202-3101011113003301-0002310313213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212033210021220-0323221032213333-3200002133330122-3203333101012110-2030120023112213-1223302232122003-2021030133231321-2303103032032232"></a>

## proxy_advertisement.advertise_on_public_default_dualstack_vip — advertise_on_public_default_dualstack_vip / 003010010000 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public_default_dualstack_vip

<a id="canonical-2010310231010331-1113120123100230-1130133211101302-1323300011022312-1313011332032302-2231232320233211-0231121223022222-1212133303212022"></a>

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
advertise_on_public_default_dualstack_vip = {}
```

<a id="canonical-1101203322220122-0313203212330132-0222013320002113-3320012000231202-1203303323033123-3332303322302033-2310221032202331-1021210330030210"></a>

## Direct properties — advertise_on_public_default_dualstack_vip / 003010010000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322200322113230-3230011003203230-2103111122021021-0011010332200123-2321231011031123-0112022203010202-3110121122202310-0120010003030133"></a>

## Next pages — advertise_on_public_default_dualstack_vip / 003010010000 / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-2220312223313100-0021103312020302-1331000023323330-1112111302001112-1233121002000332-3311110001302010-1233221013011120-2221310113023300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122223311213323-2030111020212311-1100032232100013-3013113303122121-3223233133013310-2230003310311011-0121111233312200-2031232301222133"></a>

## proxy_advertisement.advertise_on_public_default_ipv6_vip — advertise_on_public_default_ipv6_vip / 232133223212 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public_default_ipv6_vip

<a id="canonical-1332102031200011-1133321222023301-3310333101100033-2011031112313130-1001012102311000-0022332003011103-1313122030303011-3210312303100033"></a>

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
advertise_on_public_default_ipv6_vip = {}
```

<a id="canonical-0310303313023000-2110302033120001-1210222222221330-2103210011010023-0230203130321300-1031021220312102-3012202211121031-3310213233021011"></a>

## Direct properties — advertise_on_public_default_ipv6_vip / 232133223212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233312332321313-3312213023330331-3231330130303111-2300231021302013-1303330023003313-1121331030222322-3333133333212222-0210311220330221"></a>

## Next pages — advertise_on_public_default_ipv6_vip / 232133223212 / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-1222203232301300-1301320120330311-0033333301311202-3100200223211001-3002302013120130-1132130033200011-3033130130111121-3021122101310213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212001011210233-3031001203222323-0020110303002101-2023230213330003-1212011322313231-3223022130313332-3210123223010100-3322220231001330"></a>

## proxy_advertisement.advertise_on_public_default_vip — advertise_on_public_default_vip / 310200131123 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_on_public_default_vip

<a id="canonical-2302022130021210-0120233122130103-0221120000331100-0112232110212303-2120231312000302-1323232301203220-2030203122001030-2130001213223123"></a>

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
advertise_on_public_default_vip = {}
```

<a id="canonical-2200020333112213-2121110121110120-2022221312011303-1103300322220013-2112031113033023-3111121030020210-0020123113212001-1013313210203310"></a>

## Direct properties — advertise_on_public_default_vip / 310200131123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021101102113110-3112002023232222-1223313023232130-0310213230312023-3130120112133210-2211312330312330-3020310010000203-0231023103100212"></a>

## Next pages — advertise_on_public_default_vip / 310200131123 / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231313103131012-0203101212002321-0031200031133131-0011123111310101-0330302030221232-0012232100002333-3202012332110210-2032023321101121"></a>

## proxy_advertisement.advertise_v6_on_public — advertise_v6_on_public / 000131330223 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.advertise_v6_on_public

<a id="canonical-1102121313013002-0112123033211102-0303312200120302-0221001121202211-3103001303023120-1123311230231321-2011302033231102-0320313103333302"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
advertise_v6_on_public {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210303202331331-3000030000131331-2310013112022231-2223021111200022-1331222132033233-0221301032332012-3130000313321123-2011123313111323"></a>

## Direct properties — advertise_v6_on_public / 000131330223 / 3

- [public_ip](resources--dns_proxy--reference--group-002.md#canonical-3220333302012002-3032210132331233-1313301000022101-2121121121101312-2121023102010221-2222230202210002-2232301122301321-3232121311223223): complete subsection reference.

<a id="canonical-1303302012132210-0101002332032321-2021103231130331-1102121033113321-3311132332113101-2312112121333010-3021113100131121-0032111333022322"></a>

## Next pages — advertise_v6_on_public / 000131330223 / 4

- [proxy_advertisement.advertise_v6_on_public.public_ip](resources--dns_proxy--reference--group-002.md#canonical-3220333302012002-3032210132331233-1313301000022101-2121121121101312-2121023102010221-2222230202210002-2232301122301321-3232121311223223)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3220333302012002-3032210132331233-1313301000022101-2121121121101312-2121023102010221-2222230202210002-2232301122301321-3232121311223223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232023101000223-0102301312313330-1231300332331123-3201332120031021-1023230133130120-2322001000233330-2321120323312103-3302122132000120"></a>

## proxy_advertisement.advertise_v6_on_public.public_ip — public_ip / 101301031221 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123)
- proxy_advertisement.advertise_v6_on_public.public_ip

<a id="canonical-3300021010030330-3100133331003023-0121323203322311-1312201312012223-1301023321100203-0330230103120023-1103011002032102-2001032110001221"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202113330223133-3232303223101212-2221020130300202-0110220010030001-2132120132202310-3220121100121310-0123123330331101-2223203201130311"></a>

## Direct properties — public_ip / 101301031221 / 3

<a id="canonical-0133333023031101-0000101021012011-0111120102110222-0202003232121211-0200200300303112-3332020210223220-1113031231320200-1111331330003330"></a>

<a id="canonical-2202202332233323-2010320310220211-2120302000322211-2202210120112111-1232021023210201-3101232221010222-1020300002320300-2103100311203030"></a>

## name property — public_ip / 101301031221 / 4

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

<a id="canonical-3203131020011022-1121133111110032-3000123031131033-0013111301132233-1001022103323120-2112020112030122-1303311302322230-0333122021200121"></a>

<a id="canonical-1021110211313020-1112003323120110-3103022220333202-0211321031330310-1113023231232312-1323111311012023-2321320333000312-3310100302101122"></a>

## namespace property — public_ip / 101301031221 / 5

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

<a id="canonical-3100002311122231-3102221233003330-1110030121121230-3122111111200211-1102313031220033-0032102122101122-3001100320020231-1123030231103222"></a>

<a id="canonical-3022303312111102-3003233303233231-1011033223333131-1323132013232302-1030110013223023-2120023312122322-2212300210122300-0313030313331300"></a>

## tenant property — public_ip / 101301031221 / 6

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

<a id="canonical-0301223112031103-1001113033121212-0030131300232213-2303323310021312-2233302323300200-3230210102310102-2331122330101331-0002323310311322"></a>

## Next pages — public_ip / 101301031221 / 7

- [proxy_advertisement.advertise_v6_on_public](resources--dns_proxy--reference--group-002.md#canonical-0010212203201301-0332200213312212-1220010333030212-2201021203301200-0030321311132322-0001222133122102-1212100023322133-0113000303320123)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-0323231032132032-3103023010222030-3122002213313020-2200330021120102-0010211000032013-0320332021031131-0020201322001221-1333331311212313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311002111320103-1222323000103232-2303101221120001-0131113001012010-2333322203113302-0211022333032103-0212231101103321-3231102212333223"></a>

## proxy_advertisement.do_not_advertise — do_not_advertise / 220110211122 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- proxy_advertisement.do_not_advertise

<a id="canonical-3011330200112223-3321223312121211-1330323001012213-3130112113101221-0301013010310310-2213000202032202-0211201223130122-2310021302333213"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
do_not_advertise = {}
```

<a id="canonical-1212202300111311-0012013012311022-3112022210213132-3310013102033032-2302303323222013-2033130132202130-3021012032233220-2013311333310023"></a>

## Direct properties — do_not_advertise / 220110211122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311231021323330-1102130310100113-1311020031220133-2210120303110320-0012302231032012-2032112300032131-3321301333320212-1100120213011322"></a>

## Next pages — do_not_advertise / 220110211122 / 4

- [proxy_advertisement](resources--dns_proxy--reference--group-001.md#canonical-1102110013302122-1201101111110200-1313002300123133-1332121102302131-1020303100012301-3122233213101333-1103023120020201-2312232123301230)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)

<a id="canonical-3323333303000202-2222033001303102-1032113131211303-1311033002122012-1002203323202112-3112220230111220-3311101211320012-3313320002112211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110232102323221-1022101332310313-3020103111310033-0122000311303222-1313223013120103-1203122101113023-1130201312000131-0000313103130003"></a>

## timeouts — timeouts / 103202232011 / 2

Breadcrumbs:

- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- timeouts

<a id="canonical-2121133002213012-3323330102321310-3111120012100112-0102120022101130-3121022013002130-1103013032222110-1232313023220331-0300331230201033"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0003031302020022-2012203212220033-2200201311113223-3132022212200011-3111020012113302-0130033202231201-0332020031200312-3101231132110102"></a>

## Direct properties — timeouts / 103202232011 / 3

<a id="canonical-0200202011322323-3301012301121211-0332101003323232-3020211300301023-2002321300210213-0313220111310002-2301323013213202-3110021122030001"></a>

<a id="canonical-1310321331120221-1330010331130011-2331031012212322-3133033230102333-3021200203200331-3320212210232012-3101213133231023-3200231221320123"></a>

## create property — timeouts / 103202232011 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1022130011203113-2100212123031332-2322101203002233-3211030031322230-2232023030102023-3303322301132131-0230201111033022-3010310001203220"></a>

<a id="canonical-0333220101322001-3233303231222031-1230103132101203-3020230220212120-1320232223132213-2231113322101201-1313011020233211-2132230000022222"></a>

## delete property — timeouts / 103202232011 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2011132030031233-3212321333011301-0031131131013303-1022211311020122-1312022132030312-1020120131003010-2203321131321232-1020100130020332"></a>

<a id="canonical-1021011222333323-2123032322301302-0332000201311320-3132110020333330-3310121223032011-2301003221033323-1213233211132312-2013202112103030"></a>

## read property — timeouts / 103202232011 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-3231113033232332-3300021010010210-2011120231102311-1202222300233212-2123111000230301-3220131110112002-2301021133203033-0222020322033210"></a>

<a id="canonical-0222110023203311-1332123022301311-2333110032012121-2330110303111210-1013110233300011-3101032032102321-0201321102000021-0001333212221033"></a>

## update property — timeouts / 103202232011 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3101110112030122-2310000311211323-2100120313231303-3101023112133103-3000221231233122-3320302221222023-3323223210221322-2001132301233010"></a>

## Next pages — timeouts / 103202232011 / 8

- [Property reference](resources--dns_proxy--reference--group-001.md#canonical-3303213222031312-1332303332303100-2111101031113012-1202300313121231-0133112331330001-0002220333121132-1331120002013332-3233312011312210)
- [xcsh_dns_proxy](../resources/dns_proxy.md#canonical-1233100033001301-0003110332210023-3123330320331003-3311310010021030-2113023300031313-2100021000202101-1031321333010001-3220022131100122)
