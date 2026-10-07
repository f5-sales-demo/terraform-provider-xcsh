---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-0333133020303332-1323011201230012-2311303101010232-3211213101223311-2201123212232313-3233030020211010-1233332002232021-1030310120101212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-1023223131320110-3212102023100121-0122222113001321-0213311110012111-0101011000033303-2300310133222303-0011321011233213-0302113212031103"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313232102200321-2230213320101001-1220102231021303-3100122302303321-3201301301231020-3322030103000323-2101131303100000-1113100132201333"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4`

<a id="canonical-3022002020030030-3223123333233231-0211003221310213-1320310113012102-1230020100132023-3110222000223300-0222301223230320-2220300132202021"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1310330231031000-2130012120001210-1123023320201133-2223030113212011-1030333232311221-1102112023213320-3222211112103002-1323032211012122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-003.md#canonical-1030323103122102-2110303300011113-3112123030312202-0023322132323301-0313231010200312-2121232001212230-2223322131302001-0131030112133200)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-003.md#canonical-0331313103011311-0321301101021012-2132310111131201-1300130021202012-1002323132031021-3011110300030222-2113123321201131-1222120203213301)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-1012312122322110-3131320000323011-0310211130210030-0233303323323212-1313122333231303-2200021000011202-1320201132211202-0302300130101022"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1230313231031023-2022230122120020-3200310322222201-3001202310001102-2033123102110123-2013032231030202-2111123203130332-0212332232020220"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6`

<a id="canonical-2231132120223200-2231102322020202-2022121320201212-2301302333012123-2121213022233211-2232033222200330-0101131233110022-3333200001122130"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2312102122133200-0120302223320132-0133013231100200-0201333332120223-3022330101331300-0311113223000130-0123020123002012-1012331331330010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3303000132001100-3220230222001113-3131003032113202-2032120200312223-2103031100133223-0320203032011222-0023212102301010-2211202123031011"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000003013021311-2201312102022121-3200011231110110-0313223311011011-2323212000000321-3021010023021120-2312132210111113-2020213111223303"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets`

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-2232020032130012-2331121230102331-0330122310203233-3221321201111230-0232120112200210-2231332303103313-3032213100132332-3203023100000023): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-3231311301112302-0100331021113132-0200202202131020-1033322331012222-1132103203013001-1222003232023001-2112013133003001-1211333133330100): complete subsection reference.

<a id="canonical-2232020032130012-2331121230102331-0330122310203233-3221321201111230-0232120112200210-2231332303103313-3032213100132332-3203023100000023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2312102122133200-0120302223320132-0133013231100200-0201333332120223-3022330101331300-0311113223000130-0123020123002012-1012331331330010)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0230122321233130-3113323101122330-1300201300113231-2221103020302302-1311320033131020-1022220033113303-2033201310321010-2121320033002131"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-2032210033111020-1201201222203022-2102121333312330-3233233213011023-3000201020022113-2012312033333003-0213131022023112-1033033013313231"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4`

<a id="canonical-3322121301321211-3112333220233323-2303333033333120-1032322211112311-1210220300301222-0030120223213222-1323203320301301-0010323011022321"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` property

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-1330331223311222-0110100210213332-0132321013203101-0132002211333031-0202201021201023-2020121321223130-3011110232210132-2232123231033330"></a>

<a id="canonical-2332231122013223-1110111001003112-0202231230331211-2110020331321021-2303000000212122-1323221213323233-1331001323031230-2310102013230030"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` property

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3231311301112302-0100331021113132-0200202202131020-1033322331012222-1132103203013001-1222003232023001-2112013133003001-1211333133330100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.inside_static_routes](resources--aws_tgw_site--reference--group-003.md#canonical-1220102301021103-2101230331200333-1320112002130223-3322021232321002-2300321222300032-1311322021121333-1300213110111200-0122111222312130)
- [vn_config.inside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-003.md#canonical-1111202311002203-0100300223012331-1003213022321123-2132201001333033-2233011311203120-1011300330231230-1132313222230312-2233033213203231)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-003.md#canonical-2032200001011311-1330020102322220-2303133102113110-3011112022233012-3103222000122001-1001203031110230-0100311121203211-3330333223013220)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2312102122133200-0120302223320132-0133013231100200-0201333332120223-3022330101331300-0311113223000130-0123020123002012-1012331331330010)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-3120301021120021-0031000220112232-3130132120311122-2032222031220200-3133100232201231-3231322322331111-1023130212111221-0021223101001003"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320323321113320-2311303010222200-2130032133112002-2222112031310022-2022022211122322-2320221010200011-1331123003223013-2030001210310203"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6`

<a id="canonical-3110021210020211-2003030310021312-1211213011312133-0202220030323210-2132200122100001-3000211233111012-3010231130321321-2120223220022301"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` property

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-1333013233113210-0101211333300001-2033031213212002-0131311230302331-0300312122302122-2223222112020110-2032103022022012-1011000333210310"></a>

<a id="canonical-2313332312203221-3103111112313230-3002300330132201-1022132310122020-0132220113010210-1200303200021312-1020202121221203-0332001003330313"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` property

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Additional upstream details:

IPv6 address must be specified as hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0"
The address can be compacted by suppressing zeros e.g. "2001:db8::2::"

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2011122312231230-3013232300312312-2202102000033103-3030221030323032-1020222321312033-1302222331002000-1132203010312230-3310010210310332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_dc_cluster_group` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.no_dc_cluster_group

<a id="canonical-2113010033121112-2303223233130301-3313211123220002-0002111010030112-3121030222220210-0021003120111210-1133030210223123-3000320312113202"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
no_dc_cluster_group = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220211023313210-3121131032130332-1202222301132133-1313020021300200-3101221300210022-2010132200330031-1332002033100321-3002120301331100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_global_network` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.no_global_network

<a id="canonical-2221201312100210-0212332020333032-0301021012030333-1112120112312021-1332121002322313-1301312030213202-0321232322310130-0023132332110303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no global network.

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

Terraform syntax:

```terraform
no_global_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300132311300213-2213102221212301-1032330303012220-1032121031020202-0231221133312300-0101032232312332-3230133033132033-2301132320222332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_inside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.no_inside_static_routes

<a id="canonical-3203131130312231-2202032330220011-0212221102102311-0031200311002233-2032211232210211-1321133312003003-1331112322132223-2111211310002331"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no inside static routes.

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

Terraform syntax:

```terraform
no_inside_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0121330320232133-2321110122302131-1013232000113230-1212100032312110-2013322033302110-1203131120200211-3202133212002123-0302122002202023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_outside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.no_outside_static_routes

<a id="canonical-3002313302203313-0222330331130322-1322211332322222-0003010030200020-3232312220101331-0021113203201123-3021200203112210-3111221033321222"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no outside static routes.

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

Terraform syntax:

```terraform
no_outside_static_routes = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.outside_static_routes

<a id="canonical-2232131130132012-1110233200102200-1023231320300232-0102211112120011-2203120022123122-0100112130222200-2013320312201221-1302013232132330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for outside static routes.

Additional upstream details:

List of static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("static_route_list")}
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
outside_static_routes {
  # Configure direct properties listed below.
}
```

<a id="canonical-1021001210200011-1100320011231123-3121333000101203-0213322112233002-2331103223130032-3010202313013023-3012130313030211-3030023112133012"></a>

### Direct properties for `vn_config.outside_static_routes`

- [static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323): complete subsection reference.

<a id="canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- vn_config.outside_static_routes.static_route_list

<a id="canonical-0211131030323112-0203000102023232-3021021030112021-1301323311222102-2232032113032012-1222001211303303-1131202022123103-3011103131213023"></a>

Type: `"object"`. list nested block, Optional.

List of Static Routes. List of Static routes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("custom_static_route",
    "simple_static_route")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
static_route_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130302232102332-0102213332233300-1200330220031120-3320332222111232-1130123021202112-0221210102302223-2322320330320330-0031303202322033"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list`

- [custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310): complete subsection reference.

<a id="canonical-3220220133332031-3332312230201232-1022333330022231-0022233011233103-0123030111003201-1120322231111012-2332113123233033-0200231110322021"></a>

<a id="canonical-3232013212222023-1331312331001202-2020332132200331-0120112310222223-2120201202020110-0010133032302223-0202013003301323-2333333221233221"></a>

#### `vn_config.outside_static_routes.static_route_list.simple_static_route` property

Type: `"string"`. Optional.

Exclusive with \[custom\_static\_route\] Use simple static route for prefix pointing to single
interface in the network.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- vn_config.outside_static_routes.static_route_list.custom_static_route

<a id="canonical-3021211111133303-2232303013132321-1033221111103233-3300000111212030-3232002222011003-0121100300201210-1303330122212330-2302010303331010"></a>

Type: `"object"`. single nested block, Optional.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("subnets")}
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
custom_static_route {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210020300220023-3023320233333113-3101210011333313-3311111231002221-0120221310003010-3300331300103223-0123013200233333-1013223101033102"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route`

<a id="canonical-0003111212031130-3012322110322011-0022100122001201-0333331231032330-1332322223333021-3202213331001030-1213103031202031-1010322221233313"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.attrs` property

Type: `["list", "string"]`. Optional.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(4),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

- [labels](resources--aws_tgw_site--reference--group-004.md#canonical-0232323103201002-3021320032123221-3323211230201312-0100012302023310-1113210210101213-0113122013210000-0023222300300010-0010310300201201): complete subsection reference.

- [nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030): complete subsection reference.

- [subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321): complete subsection reference.

<a id="canonical-0232323103201002-3021320032123221-3323211230201312-0100012302023310-1113210210101213-0113122013210000-0023222300300010-0010310300201201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.labels` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- vn_config.outside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-1332222301001012-3002002103331232-2022023010200300-3332311000301121-2000211010102231-3223322022211312-2020000311131021-3300100301222210"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Static Route, these labels can be used in network policy.

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
labels {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-3321022100210320-1212022111321120-0230301231312311-0321200032023332-2322200100232033-0003233003212131-3000221212301130-2333322010112131"></a>

Type: `"object"`. single nested block, Optional.

Nexthop. Identifies the next-hop for a route.

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
nexthop {
  # Configure direct properties listed below.
}
```

<a id="canonical-3101003223021122-2020220030332003-0220021111301313-3102300222330011-2221111320030020-1123330031013220-1123203102023313-0313011110003132"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop`

- [interface](resources--aws_tgw_site--reference--group-004.md#canonical-2023012210202202-2022212003110122-1001120330030311-1120311120221320-3110200110001203-3310210332003212-2213112020121223-3323102031323000): complete subsection reference.

- [nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301): complete subsection reference.

<a id="canonical-1122023323030231-0202321033332110-0233110120110231-3300111310130000-0312312010130310-0203123120222111-0323303021121100-2113212203322203"></a>

<a id="canonical-3122201333012300-3113220210311220-0332020002130022-1132220132310103-3302032213201322-1012100320323110-3111211103201121-3033201132201321"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.type` property

Type: `"string"`. Optional.

\[Enum: NEXT\_HOP\_DEFAULT\_GATEWAY|NEXT\_HOP\_USE\_CONFIGURED|NEXT\_HOP\_NETWORK\_INTERFACE\]
Defines types of next-hop Use default gateway on the local interface as gateway for route. Assumes
there is only one local interface on the virtual network. Use the specified address as nexthop Use
the network interface as nexthop Discard nexthop, used when attr type is Advertise Used in VoltADN..
Possible values are \`NEXT\_HOP\_DEFAULT\_GATEWAY\`, \`NEXT\_HOP\_USE\_CONFIGURED\`,
\`NEXT\_HOP\_NETWORK\_INTERFACE\`. Defaults to \`NEXT\_HOP\_DEFAULT\_GATEWAY\`.

Additional upstream details:

Defines types of next-hop

Use default gateway on the local interface as gateway for route. Use the specified address as
nexthop Use the network interface as nexthop Discard nexthop, used when attr type is Advertise Used
in VoltADN private virtual network.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["NEXT_HOP_DEFAULT_GATEWAY","NEXT_HOP_NETWORK_INTERFACE","NEXT_HOP_USE_CONFIGURED"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NEXT_HOP_DEFAULT_GATEWAY",
  "enum": [
    "NEXT_HOP_DEFAULT_GATEWAY",
    "NEXT_HOP_USE_CONFIGURED",
    "NEXT_HOP_NETWORK_INTERFACE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2023012210202202-2022212003110122-1001120330030311-1120311120221320-3110200110001203-3310210332003212-2213112020121223-3323102031323000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-0121332302020121-2331131002030011-1303021001222310-0003221212303213-3211032122100200-2212003230000111-2212200030012310-0020120021130111"></a>

Type: `"object"`. list nested block, Optional.

Nexthop is network interface when type is 'Network-Interface'.

Additional upstream details:

Nexthop is network interface when type is "Network-Interface"

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

Terraform syntax:

```terraform
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0111032012003320-2122200220311123-0120230030021011-0211030101020010-2121313123002312-3330132110012333-2302320130021232-0331230103012101"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface`

<a id="canonical-3221002323313332-0123332021033331-1031030121231132-3212013301031102-2022100020231333-2332033001000201-3023023000030312-2100220031212033"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0030013322112311-2312010033031031-3033023303011122-0201001101332133-2213310312011132-3201210031020222-1212300310300211-0233010220110132"></a>

<a id="canonical-2312303300231100-3223301012123021-3111031131333231-0010130230322332-0310333110111033-2212112032103231-2000321021012010-2331133032031223"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0330022121012303-1030113301212300-0310322132332030-3313201123000222-0220000310011233-3010211221011333-2121033302032333-3230021133223032"></a>

<a id="canonical-3323233211233323-3000012213230232-3230130200021113-3333102023103233-0123202031003122-2123132102132001-2302122001133312-3311322130300311"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1002201013231301-1130112132333120-1032203321003211-0123320211111111-0003311010212020-0030103102220223-0031233301330332-1233013323103101"></a>

<a id="canonical-3103331213033231-3323022213010322-2220032231123103-0012333013210020-1022103120321101-2100310212003001-3223213120113230-3011320312213111"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0323201233133020-3310131130331131-2131113223113303-0111121121312302-2200232301220112-1101131320133330-0013312033113321-2023111023033232"></a>

<a id="canonical-0020302120310102-2313013313123321-2313332333101223-0202202333203023-2210222202100112-0013111330132230-3130030323301212-1100103310312113"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-3211232202201232-1031300221322332-3332100103100122-2313323303000230-3321130303320310-0121203121030022-1301200213222033-0120003101331032"></a>

Type: `"object"`. single nested block, Optional.

IP Address used to specify an IPv4 or IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("dual_stack",
    "ipv4"),
  validators.ConflictingObjectAttributes("dual_stack",
    "ipv6"),
  validators.ConflictingObjectAttributes("ipv4",
    "ipv6")}
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
  "x-ves-oneof-field-ver": "[\"dual_stack\",\"ipv4\",\"ipv6\"]"
}
```

Terraform syntax:

```terraform
nexthop_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022302033130123-0031311121222213-3023330000333111-2023320112322230-0120320011331133-3211100023112210-2033102001011302-1122230223220101"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address`

- [dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323): complete subsection reference.

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-3012220220211011-0101111003200313-3103010211022000-1102023331221330-2310230232020330-3222202000331100-2020121310121311-3031332020021120): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-0102212312123232-2331013133111011-1130302133131120-2000330101011122-2020303120011212-3032111211221211-0002122112000302-2133201133212311): complete subsection reference.

<a id="canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-2010303032211013-1130002131031332-0110231002123323-1121003031220200-0120130301132312-1220212303220102-1133330300211303-2312201313101110"></a>

Type: `"object"`. single nested block, Optional.

DualStackAddressType represents both IPv4 and IPv6 together.

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
dual_stack {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201211302123323-0101133203330221-1102223222320013-3013031213323303-1113332010113202-3023032121021023-2000311103022231-0003320322131122"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack`

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-2301101203002112-2201022211203301-3113213310122323-3331220132331233-0010103232211230-2303222131033301-0002102123130031-3122112332230233): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-2200211323110000-3103123202121323-1102320321331321-0123013210110321-0111211003020121-3202001130111203-3321211030002101-1033330203213223): complete subsection reference.

<a id="canonical-2301101203002112-2201022211203301-3113213310122323-3331220132331233-0010103232211230-2303222131033301-0002102123130031-3122112332230233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-0122332203100202-1323121013131003-1120331030123202-1033001023032020-1210212323133000-1031313230332330-3221130333020102-2303003332132000"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210020212221211-1223030213201010-1031202230032012-1102211012103222-3331020313332110-0021013011200130-3030021331233101-2333033332013022"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4`

<a id="canonical-0132030300320113-0110132132331220-2022302221133112-3303312013130320-2332123312301202-0302212311213000-2303003230011013-3110233211133113"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2200211323110000-3103123202121323-1102320321331321-0123013210110321-0111211003020121-3202001130111203-3321211030002101-1033330203213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](resources--aws_tgw_site--reference--group-004.md#canonical-3000130022202130-1231231312311331-3012111201122220-1211233110223010-1330300000301110-1322113202102323-3222211232032111-1233002320311323)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-3103021333233110-1011210131322000-2123321202132102-1033132312123001-1332332020200203-3113112232100312-2202200210302311-0130210312320211"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230322123103001-1113033010021020-1012103313031310-2321032231220010-1113032103300221-1011032213311112-3133133020033021-2120133133130030"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6`

<a id="canonical-3000222210112033-0313011023210122-2101122210233303-1102323223021113-3010010101022002-2302113201110111-3130020112301230-2020301312032030"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3012220220211011-0101111003200313-3103010211022000-1102023331221330-2310230232020330-3222202000331100-2020121310121311-3031332020021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-0022102320330031-1301233021001322-0133021130031201-1313122222013012-1133113200001310-2320302220302023-3312133032223203-0233122221012212"></a>

Type: `"object"`. single nested block, Optional.

IPv4 address in dotted decimal notation (e.g., 192.0.2.1).

Additional upstream details:

IPv4 Address in dot-decimal notation.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300030231221200-3032331021210320-0132131133031301-0022320320301001-2303021013210203-0103110302122102-0010200322233322-1222223121013200"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4`

<a id="canonical-3012032120033230-1002320131101230-0302310110031322-0222123111033310-1130011333121033-2331211001320323-2021223223010332-1200133300013111"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` property

Type: `"string"`. Optional.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0102212312123232-2331013133111011-1130302133131120-2000330101011122-2020303120011212-3032111211221211-0002122112000302-2133201133212311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop](resources--aws_tgw_site--reference--group-004.md#canonical-0003030130021111-3312333313002332-2112222111320300-1102302011333013-1321333031231001-0002003300230301-1030331331223202-0121210011003030)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](resources--aws_tgw_site--reference--group-004.md#canonical-2022301320110121-3133021101100200-1312331200213023-0120233323001100-3023323112033212-3001021103301221-2120102230133113-3221030222230301)
- vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-3110230121213301-2031211111333120-1223001300021321-3210231131001300-2222113113003332-2221101013021333-3132203201131120-2001231132120122"></a>

Type: `"object"`. single nested block, Optional.

IPv6 Address specified as hexadecimal numbers separated by ':'.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220120121100330-2232323100222330-2120220032330220-3110302203013332-3230230221312123-0031123220122012-0131222130133223-2303231222333033"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6`

<a id="canonical-3122323132320201-3302232002130000-0110111200113013-1022010232320320-2003321102313023-2131003022312220-0033023222122033-1221100021023201"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` property

Type: `"string"`. Optional.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3033333132110230-0101321032301313-2200333332220220-3330322331100021-3221102322011032-2113022010001113-2220330301031031-3232030221323301"></a>

Type: `"object"`. list nested block, Optional.

Subnets. List of route prefixes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.ConflictingListObjectAttributes("ipv4",
    "ipv6")}
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256"
  }
}
```

Terraform syntax:

```terraform
subnets {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020233300302330-1020312123122231-2120103213131200-3003030222233311-0113021321113333-1331123030310013-0031312213003012-0110333312203111"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets`

- [IPv4](resources--aws_tgw_site--reference--group-004.md#canonical-1100132202210331-2030031320112010-2231332212322120-3223102330101330-0301101201211030-0312223003210130-0033321303211333-0223103020031213): complete subsection reference.

- [IPv6](resources--aws_tgw_site--reference--group-004.md#canonical-1022033101311020-1133132121221020-0101021213303113-2332100001221011-1311202021023223-0303103210313230-3102301121021330-3202323332110100): complete subsection reference.

<a id="canonical-1100132202210331-2030031320112010-2231332212322120-3223102330101330-0301101201211030-0312223003210130-0033321303211333-0223103020031213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-0323231030102312-2031232301232010-0220222200330232-0231223122131211-3213023310322223-0313313013310030-2301132112303002-1120303112133022"></a>

Type: `"object"`. single nested block, Optional.

IPv4 subnets specified as prefix and prefix-length. Prefix length must be &lt;= 32.

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
ipv4 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210023012222010-3103101010301311-1212031211330212-3133002122000100-2313213102301133-2023230002021120-0012030000002100-3011002200013332"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4`

<a id="canonical-1232010211121032-0232021230333032-1022300220202133-3323230101203020-3011111101320201-0201313213303000-1333310333033223-0310100330212221"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` property

Type: `"number"`. Optional.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(32),
}
```

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "32"
  }
}
```

<a id="canonical-3103022313302202-0002120301011210-3010302030010020-1012013123113010-0123300332121311-2300003102102120-3320212010212200-0013310312101123"></a>

<a id="canonical-0033213033330233-0321223210100031-0320103211110000-3010123213210303-3013013200301122-2131202333003013-2130020201333211-0200222321130021"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` property

Type: `"string"`. Optional.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1022033101311020-1133132121221020-0101021213303113-2332100001221011-1311202021023223-0303103210313230-3102301121021330-3202323332110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- [vn_config.outside_static_routes](resources--aws_tgw_site--reference--group-004.md#canonical-2301113011113011-1112320230210202-0313100212313032-1131333013212023-3211313333212000-2110001321132102-1131010130012311-2101312102132133)
- [vn_config.outside_static_routes.static_route_list](resources--aws_tgw_site--reference--group-004.md#canonical-0012330200333200-1003330212132300-0033033221012201-3330110232112210-0013230010001331-0003332013203201-0013211233213031-0201233331322323)
- [vn_config.outside_static_routes.static_route_list.custom_static_route](resources--aws_tgw_site--reference--group-004.md#canonical-3112202231210201-1221331101020010-2311032322110333-3203310310200232-2230321032201221-3313011300211333-0131031220003312-2233310210310310)
- [vn_config.outside_static_routes.static_route_list.custom_static_route.subnets](resources--aws_tgw_site--reference--group-004.md#canonical-2003220033322023-1112302212322313-3102223321203121-3022310110102112-1302132203022131-3321131111012001-0221030023311130-3232233232102321)
- vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1023232322123303-2120112113232100-0103223303310332-1313030210220223-0032302220033330-2330203030211331-3333100021200202-3112121010123032"></a>

Type: `"object"`. single nested block, Optional.

IPv6 subnets specified as prefix and prefix-length. Prefix-legnth must be &lt;= 128.

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
ipv6 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3310130030322333-1221112332220221-2230122210100312-3113203032300232-3032001311321133-2222313311130110-3111112033331222-1200123030202011"></a>

### Direct properties for `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6`

<a id="canonical-0132122202103201-0211030311300202-1132011203322003-2220101232103030-0110020121313133-2210002100223302-1203103120023210-0300123021322323"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` property

Type: `"number"`. Optional.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "128"
  }
}
```

<a id="canonical-3111311323101023-3311301333110110-2330200200323033-2120232102212013-3333301203210010-0223001013321213-2111032000323001-0320313323013221"></a>

<a id="canonical-2020301221123221-0321110011212331-0223011121232000-2110301000032322-3133200033131123-2222012031001011-0222013202111211-3011000210003313"></a>

#### `vn_config.outside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` property

Type: `"string"`. Optional.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Additional upstream details:

IPv6 address must be specified as hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0"
The address can be compacted by suppressing zeros e.g. "2001:db8::2::"

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3303020033221031-0121001232113232-0331101030300303-3303101011322302-1303033333020021-2321013112020223-0312203301333130-3001033131210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.sm_connection_public_ip` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.sm_connection_public_ip

<a id="canonical-1202210310030101-1012103013201211-1002001200202022-3330223100321212-1113102330110213-1102313323310013-3212020313210310-2211230323011322"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
sm_connection_public_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021012122100010-3002323221023231-2002011020202100-1001130223230002-1100232013102230-1031022221333310-1130233020000232-2303021330003011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.sm_connection_pvt_ip` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vn_config](resources--aws_tgw_site--reference--group-003.md#canonical-0101002010113120-1121300322102322-2320133222211303-0011102031313120-3202320223313220-3220002032331202-0331231132033323-1222203320002031)
- vn_config.sm_connection_pvt_ip

<a id="canonical-3001222103030002-2113002333332213-3101301320231220-2101231001300030-3330022330113332-2130013113113300-2122332310131203-3002112231320021"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
sm_connection_pvt_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vpc_attachments` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- vpc_attachments

<a id="canonical-2330320031031130-0223131233013102-0031030323103312-1323110303231202-0032131311202231-2300322232210321-2211310022213121-2021200310311132"></a>

Type: `"object"`. single nested block, Optional.

Spoke VPCs to be attached to the AWS TGW Site.

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
vpc_attachments {
  # Configure direct properties listed below.
}
```

<a id="canonical-3322031103121310-3300011133312033-1021021201032323-0323320232220313-3123301130001223-1211021030313112-1120110021101111-0022012023130102"></a>

### Direct properties for `vpc_attachments`

- [vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302): complete subsection reference.

<a id="canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vpc_attachments.vpc_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303)
- vpc_attachments.vpc_list

<a id="canonical-1210321101230210-2200013012202212-0211221230322001-2121112201021310-3211233213301020-2002023033112131-3133321132020220-0222221111321220"></a>

Type: `"object"`. list nested block, Optional.

List of VPC attachments to transit gateway.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128"
  }
}
```

Terraform syntax:

```terraform
vpc_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021213000033012-1301012201022010-3302032312233212-1320121131032233-3023313030222310-0130121130230123-1213332121333202-0013110012112301"></a>

### Direct properties for `vpc_attachments.vpc_list`

- [labels](resources--aws_tgw_site--reference--group-004.md#canonical-3030120202031300-1103002110210310-2320233211301011-3102302231331323-3201002123312313-2032312122232301-3012212022301011-3313030103300002): complete subsection reference.

<a id="canonical-3200001003301011-3223122010122023-1300221302312301-2031311121211021-1112132231112213-0300212231022323-3302203123022121-1010223002332223"></a>

<a id="canonical-3012312123331212-1231231111302003-3010113110312201-1022033202030222-3231123232311323-1213000222210231-2120212011020122-0102230231301310"></a>

#### `vpc_attachments.vpc_list.vpc_id` property

Type: `"string"`. Optional.

VPC ID. Information about existing VPC.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
    },
    "pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(vpc-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-3030120202031300-1103002110210310-2320233211301011-3102302231331323-3201002123312313-2032312122232301-3012212022301011-3313030103300002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vpc_attachments.vpc_list.labels` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [vpc_attachments](resources--aws_tgw_site--reference--group-004.md#canonical-1222010312230111-1313132323202221-2013100312110003-1213330213121030-1101221323132002-2333001031120331-0310220231323322-3332320031313303)
- [vpc_attachments.vpc_list](resources--aws_tgw_site--reference--group-004.md#canonical-2222113021122020-0012033011131223-1212101133221212-1133313010011213-2032103312331322-1002312232020110-2300122030323031-0113000122111302)
- vpc_attachments.vpc_list.labels

<a id="canonical-3210011200110111-1123303222023031-3121103133003212-3301201031101331-2302321202100113-1102100130033303-2110013002113232-1233213302200202"></a>

Type: `"object"`. single nested block, Optional.

Add labels for the VPC attachment. These labels can then be used in policies such as enhanced
firewall.

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
labels {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_signatures` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- waf_signatures

<a id="canonical-3222213203010330-0133312022232213-1321000200323331-1122003320120230-1011033320302333-3322120000321332-0231130203100112-2321011322233231"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-2120232220013203-2121020233033000-2201313220222131-2011013230110022-0013033103101222-1201032130311223-1102133013321233-3110010010112123"></a>

### Direct properties for `waf_signatures`

- [automatic](resources--aws_tgw_site--reference--group-004.md#canonical-3232131113232120-1030302320302203-0212112203333322-1200310110032030-1232102131102032-1333003102012002-2202203111032311-0020021100203011): complete subsection reference.

- [manual](resources--aws_tgw_site--reference--group-004.md#canonical-3131311212203203-0323201303302230-3101003023223233-1001210013131211-2110130201001003-3302131020220333-2201200202321013-1020212313120130): complete subsection reference.

<a id="canonical-3232131113232120-1030302320302203-0212112203333322-1200310110032030-1232102131102032-1333003102012002-2202203111032311-0020021100203011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_signatures.automatic` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032)
- waf_signatures.automatic

<a id="canonical-0313330010000201-2032202213233101-2201101031123001-1023232213222320-2013230313222013-2101323121320132-3031123323200121-2030232013223213"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
automatic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131311212203203-0323201303302230-3101003023223233-1001210013131211-2110130201001003-3302131020220333-2201200202321013-1020212313120130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `waf_signatures.manual` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [waf_signatures](resources--aws_tgw_site--reference--group-004.md#canonical-1231203221223333-2112012121012221-1133120301121100-2302021133012202-1310000132122112-1230101332320210-1011033201120021-0322100013102032)
- waf_signatures.manual

<a id="canonical-0012101302001113-2310030110220330-1112010323013232-0102111022123101-1032033001110320-1222021331023312-3302323013130322-0211021220212203"></a>

Type: `["object", {}]`. Optional.

Enable this option

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

Terraform syntax:

```terraform
manual = {}
```

This is an empty object or choice marker. It has no direct properties.
