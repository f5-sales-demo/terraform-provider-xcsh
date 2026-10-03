---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0132121010102103-3012211112132100-0330202301131013-1212233133223320-0130002230332202-0322202223333310-1333322130231201-1213000210021310"></a>

## aws.not_managed.node_list.interface_list.ethernet_interface — ethernet_interface / 132131032302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-2113012110122220-0111133211023101-3100213320122201-0233023312010131-2203031110332032-1301203012210221-3120003313331023-2320020331000233"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("mac")}
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
ethernet_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223330101003103-1211232202102103-0112202321023022-1112221301333220-2312202332221102-0000211320210212-3133110203121012-1213110302032010"></a>

## Direct properties — ethernet_interface / 132131032302 / 3

<a id="canonical-0123200122012023-2020213230122120-3000031313021130-2203000202201202-2122303101133032-1020223200103330-3103233000200300-2313010202000313"></a>

<a id="canonical-0031303012101331-2312012120103232-3133310302113321-1102320321130032-0210213110132101-0231202202320110-3311111131130201-0010001131302312"></a>

## device property — ethernet_interface / 132131032302 / 4

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Upstream description:

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "false",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2113013330312213-2333301322321132-2232011220230333-1303320230200131-0133303223203333-3222323033023310-1132021301112223-3023322113120212"></a>

<a id="canonical-0233021001220302-3330321303233131-3200310011300000-1300303023200130-1130233303300211-1113202323221011-1101311020201121-3313032021001011"></a>

## mac property — ethernet_interface / 132131032302 / 5

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Upstream description:

Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.MACValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "mac-address",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.mac": "true"
  }
}
```

<a id="canonical-3002223303313303-0232133001211300-1111301233132003-3112310323302113-3000000310031332-2102331110121331-0301113102103330-3322113331222210"></a>

## Next pages — ethernet_interface / 132131032302 / 6

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011201032331212-0133020233131003-1312301000222131-0231322333001101-0232131102012311-1130010000302312-1203322323331011-3021201310230013"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 221212233201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-2133112330200311-3030301302232031-0112111301023023-0332121220312120-3003001113123212-3311322211030133-1020030110232210-2100301130023200"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("host",
    "router")}
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
  "x-ves-oneof-field-autoconfig_choice": "[\"host\",\"router\"]"
}
```

Terraform syntax:

```terraform
ipv6_auto_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3201212100013200-0200110310233021-3013102222022212-0130322122230223-0020202113201212-1133333323033120-1003231210023033-1330001120131011"></a>

## Direct properties — ipv6_auto_config / 221212233201 / 3

- [host](resources--securemesh_site_v2--reference--group-004.md#canonical-1113030123022303-3101213222012011-3001021120320221-2033021103231303-2001032223211112-0331321130000110-3101203213310331-1323001120130001): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111): complete subsection reference.

<a id="canonical-0300132320210323-0230211233121220-1022001010223300-2333100012122131-2233213030333021-1232100300031122-0132232220022300-0300131202012311"></a>

## Next pages — ipv6_auto_config / 221212233201 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-004.md#canonical-1113030123022303-3101213222012011-3001021120320221-2033021103231303-2001032223211112-0331321130000110-3101203213310331-1323001120130001)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1113030123022303-3101213222012011-3001021120320221-2033021103231303-2001032223211112-0331321130000110-3101203213310331-1323001120130001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331033313010230-1203132311122232-1313310123033121-1201332212102112-1202002222311122-1300320033222121-0233321130020020-1323002210133101"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 231222231213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-0012022111132300-2333310310223212-0211321003320310-0010330133001210-3121223133323031-1001211101213212-1001310313132112-3222300001023330"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

<a id="canonical-0332310100221233-3113330201230323-3322103131020032-3202102222212300-1122301321303103-3333232230233222-2121033331230003-3200132110302332"></a>

## Direct properties — host / 231222231213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220320112032331-3031101123201101-2020111231033203-3313003002111222-1122113012102210-1112112230310011-0021013211102332-3232323230222132"></a>

## Next pages — host / 231222231213 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200022321131120-1100330100032112-2212233122202311-3200032211310210-0210021320331221-1123302030310033-1132003003210303-3112131103110123"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 322330000321 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1331110200200230-0102030033201003-0200211333322020-0301100222201233-3232300321120013-0202123002303120-0303012303010000-0233030111103110"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("network_prefix",
    "stateful")}
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
  "x-ves-oneof-field-address_choice": "[\"network_prefix\",\"stateful\"]"
}
```

Terraform syntax:

```terraform
router {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033301102220132-0123222013213032-3232021333213233-2002012103131110-1300333230130213-3000232131312122-1202303123000303-2010210012222231"></a>

## Direct properties — router / 322330000321 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030): complete subsection reference.

<a id="canonical-0210310111130121-3320003101223312-0231002303222010-2012003000333000-3203211113013010-0102001223213032-3323133330311120-2102211133102220"></a>

<a id="canonical-1232130010233312-2333313031201113-1022221221032212-2020333003100213-2001001202323123-3201203013003033-2103222213212121-2022220122133123"></a>

## network_prefix property — router / 322330000321 / 4

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": ".*::/64$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv6_prefix": "true",
    "ves.io.schema.rules.string.pattern": ".*::/64$"
  }
}
```

- [stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110): complete subsection reference.

<a id="canonical-2021131030122000-2100231231030131-2220222100131010-0112032310202232-3212131302023211-0032303222221220-1202220332002221-1323313023322211"></a>

## Next pages — router / 322330000321 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113132323221010-1012301101023310-1222013020022120-2020120113203301-0033301003202313-0023123233013010-0110120121000231-3122100122022231"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 212132202213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-2021312003012032-0221133102211110-1023311131330020-3023113011310330-0133203130220132-2201131111201132-1131122030133111-0002232302332301"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_list",
    "local_dns")}
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
  "x-ves-oneof-field-dns_choice": "[\"configured_list\",\"local_dns\"]"
}
```

Terraform syntax:

```terraform
dns_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121231130303022-1201330122323233-3113220230331000-2212110011112311-3331332321101323-2333120010110323-3112301300202311-2333223310213113"></a>

## Direct properties — dns_config / 212132202213 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0132132211230333-2121113302300120-2200302331330301-2332122031230022-3321003121030201-1003221302011133-2323113301133001-3022001010232300): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000): complete subsection reference.

<a id="canonical-3223232100131032-0310210303011023-2003221323121332-2220123002130011-1001222202213312-1311023031230032-1232130232211010-2130000202211023"></a>

## Next pages — dns_config / 212132202213 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0132132211230333-2121113302300120-2200302331330301-2332122031230022-3321003121030201-1003221302011133-2323113301133001-3022001010232300)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0132132211230333-2121113302300120-2200302331330301-2332122031230022-3321003121030201-1003221302011133-2323113301133001-3022001010232300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101121223013001-1113011020010213-1202232220102200-3002113320102100-0020010310313120-3331100220331110-2311323313130211-0011233030121012"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 323333132212 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0221022121000300-1312021002201310-0232030022300111-3232032132110000-0333332032313102-1101030211322321-0103102023013202-0300121202000231"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dns_list")}
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
configured_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320321313022032-1003213002133301-3013232003301310-0203201011121313-3113120312201102-3301031021200333-0010003333212211-2323131321200230"></a>

## Direct properties — configured_list / 323333132212 / 3

<a id="canonical-3122020210103312-1313033310000311-0130120132013000-3200123100230121-3330323003300012-1012033221003121-1121032303021122-3123131010121203"></a>

<a id="canonical-2110212133333211-1331010311231222-3120210233120212-1300313030223012-3222001320012211-2031202132330021-3023131213003022-2010130101211322"></a>

## dns_list property — configured_list / 323333132212 / 4

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 4),
}
```

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
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.ipv6": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0303322231001332-2332011220010020-3320023210123223-2132030102003000-2020230301131032-0122330233121003-0012223101111011-0302223102202011"></a>

## Next pages — configured_list / 323333132212 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033102202132230-2002321213010013-1201211220231133-3303301003002113-1221002331211101-0022022323021301-2312230203223332-0300230130102102"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 113231232312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-0211000302131233-0333222220221301-0320332320133320-3333333332303112-2311013310221031-1113110102323232-0010203332021031-3123101202130123"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("configured_address",
    "first_address"),
  validators.ConflictingObjectAttributes("configured_address",
    "last_address"),
  validators.ConflictingObjectAttributes("first_address",
    "last_address")}
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
  "x-ves-oneof-field-local_dns_choice": "[\"configured_address\",\"first_address\",\"last_address\"]"
}
```

Terraform syntax:

```terraform
local_dns {
  # Configure direct properties listed below.
}
```

<a id="canonical-1022312323032111-1020002233203211-1221011330321302-3301120312120330-3102012202301031-1333203212112322-2110213010101003-3112130313300321"></a>

## Direct properties — local_dns / 113231232312 / 3

<a id="canonical-1011320003222200-2203210022302320-0223112003302103-3312203002301113-3130121111223213-2033103200011213-0010120211311133-1223330323322102"></a>

<a id="canonical-3112211322021121-3213203312300031-2223310221202103-0230132201033021-0303300231332121-1331222222023110-1001012303010123-0312322301302110"></a>

## configured_address property — local_dns / 113231232312 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
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

- [first_address](resources--securemesh_site_v2--reference--group-004.md#canonical-3120121000201311-2332022321121020-2221120100312330-1033212101300212-2101332130302112-0010111201330132-2223302303121120-2222000222021113): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0303321111201131-1003011222031021-1003310002233211-1202203003320310-0220333010321110-3222220113311003-2300213202211332-1133320002030222): complete subsection reference.

<a id="canonical-0302322023022023-1300320212123000-0332231322232221-2231221113133331-1331012323001120-1321220322211232-0333030330133112-1011313130321012"></a>

## Next pages — local_dns / 113231232312 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-004.md#canonical-3120121000201311-2332022321121020-2221120100312330-1033212101300212-2101332130302112-0010111201330132-2223302303121120-2222000222021113)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0303321111201131-1003011222031021-1003310002233211-1202203003320310-0220333010321110-3222220113311003-2300213202211332-1133320002030222)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3120121000201311-2332022321121020-2221120100312330-1033212101300212-2101332130302112-0010111201330132-2223302303121120-2222000222021113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102210012220030-3331300133231332-2103132332130121-3302310003222312-0222303022221130-1030031300230200-2002202000122113-0122330221112030"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 131112123020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0113321112033221-2123002303202203-3133101130111320-0101222131221132-0031102121012003-3230210003331013-2210112210203332-1302320303311330"></a>

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
first_address = {}
```

<a id="canonical-3233020123221212-0201111003112131-2323313322101301-1022212330301200-0012212032231231-2031010112000231-1123103021030210-2222310100301002"></a>

## Direct properties — first_address / 131112123020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010300030230322-1111033122221130-3321200023211223-1331003122100323-1111332233130032-0330213003312211-0320020103032020-1213020310231302"></a>

## Next pages — first_address / 131112123020 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0303321111201131-1003011222031021-1003310002233211-1202203003320310-0220333010321110-3222220113311003-2300213202211332-1133320002030222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020033311313122-3201330330210222-3312033221231111-2221310210230310-0320222131002013-2231033002202213-2212122001222211-0232203112303031"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 122303232223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-004.md#canonical-2203020111310202-2331111320103323-3303101323321032-0000330110021030-2323332213123300-1111202002323002-1323122031213132-0003202130302030)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-3113310320211021-2332320332211321-1333212201132232-2010301130001110-0131003233103011-2302120111221333-0011232311121021-1000111332123220"></a>

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
last_address = {}
```

<a id="canonical-2322320221213010-0303121120002302-1011222202312231-3000102013000330-1212221212303312-1000112121120230-2210202322103103-2012300131023230"></a>

## Direct properties — last_address / 122303232223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213203303030322-0101301031302203-1001202111200223-2022102101213011-2312233323022201-1022303202030010-2100232311311202-3101222022301232"></a>

## Next pages — last_address / 122303232223 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-004.md#canonical-2233133321011121-3111110120213000-3121332012330321-1102232233201200-0121102130021303-0132323121332002-3201113310203313-3020102013321000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023013312302301-0212320320203031-0200031121321123-3011133313201202-0320030120332013-0123221000312013-3200010230313001-2120102020021030"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 333201332101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-2300323301032022-1311231103020311-3120020101333330-2031110010002133-2120233302133333-1123113023123102-3302211122330011-1333223230221130"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
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
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
stateful {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203100133221013-3213010213223223-3230002031231122-3300021302230113-0020333322332233-2020131133203011-2101102020102013-0300311003211332"></a>

## Direct properties — stateful / 333201332101 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-004.md#canonical-0110311100313221-0030311223223302-3032023030021201-3032213032332310-2100223210110101-2322232130001130-2112022320323032-3300132311100123): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-004.md#canonical-0111021313130113-3012102110032021-3301300332200020-1330313222320200-1121030223131221-0231101320012332-0220212002231303-2133323130333012): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020): complete subsection reference.

<a id="canonical-3000001212032321-2113202202003023-0303100231303030-1023111201323002-0010231033311023-1231220020233032-0001302220122010-3123220133112213"></a>

<a id="canonical-3130210003312322-0310012211023031-0231330113300130-1132202310332302-2322101131333003-3132023021122010-1132121031201130-3201012132103311"></a>

## fixed_ip_map property — stateful / 333201332101 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-004.md#canonical-1330001013210123-1120220010020113-3202320223302023-2231201331210030-0101122001201110-3230010032333130-0003121112232331-0031022010210210): complete subsection reference.

<a id="canonical-2123001233320121-0033320111103332-1111220101121330-2103203022320111-1132230131120212-0333123000200121-0323203202102203-1213103111031123"></a>

## Next pages — stateful / 333201332101 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-004.md#canonical-0110311100313221-0030311223223302-3032023030021201-3032213032332310-2100223210110101-2322232130001130-2112022320323032-3300132311100123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-004.md#canonical-0111021313130113-3012102110032021-3301300332200020-1330313222320200-1121030223131221-0231101320012332-0220212002231303-2133323130333012)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-004.md#canonical-1330001013210123-1120220010020113-3202320223302023-2231201331210030-0101122001201110-3230010032333130-0003121112232331-0031022010210210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0110311100313221-0030311223223302-3032023030021201-3032213032332310-2100223210110101-2322232130001130-2112022320323032-3300132311100123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302013221313022-3103111001320011-0012200002303233-2121111100331333-2012210112112002-1203331211223312-0120111033300111-2222102020211222"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 200302222310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2103321310223012-3213003311011211-1222330032233111-2133012132022130-1111123031210022-3133020210032101-3113020111110211-2011303123112230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-0123132111110022-1023213320221103-3123033131103101-1231010230122110-2012303202212303-2120110323220021-1212310133110230-1110321212131022"></a>

## Direct properties — automatic_from_end / 200302222310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233113233013022-3132313023133303-1230011022323330-2203013203132110-1300211233031210-2123120131211113-2101110331301222-1330022122102333"></a>

## Next pages — automatic_from_end / 200302222310 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0111021313130113-3012102110032021-3301300332200020-1330313222320200-1121030223131221-0231101320012332-0220212002231303-2133323130333012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221023301330100-0212013212211023-0013021120013210-0031312333102001-3212013233132230-1200220230332312-3110002003121130-3333230302001233"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 022300211132 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0011131323322333-2222122013230302-3213201211123001-1000111003111012-3131022310021022-3220100013110002-1302031303332320-2303212233323000"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-3302312231131121-0000110333310120-0200122022221330-2021102003300222-3011310132232123-1032322233330301-1002013120132200-2233313223203330"></a>

## Direct properties — automatic_from_start / 022300211132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002131310100300-3222210222313123-0111032103203011-1222020310322111-1102123203021303-0333320130011030-2212313130032130-3012131203033233"></a>

## Next pages — automatic_from_start / 022300211132 / 4

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010123103232320-0222202331020121-1000331231312203-2102023103123231-3112122130112010-1002020001230321-0110331233322122-3033010131111232"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 011132311322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2123322302033321-0111010333321010-0210331221322023-0223232123002333-3003203013231303-3223232033310232-0303330023021222-0121213203032333"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP server can allocate IP addresses.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013213111232330-0223230310013231-2322202010222322-1312300013331003-2331302312022212-1131131002111010-1033002022122222-0302323331023221"></a>

## Direct properties — dhcp_networks / 011132311322 / 3

<a id="canonical-1322210211013111-1223203011222301-1301230200003023-2203320030120320-3232110213000132-3323122031330030-2222211200003313-0203301020230202"></a>

<a id="canonical-2032112130002110-0123322303010233-3123013201203301-1333121333301011-0120330112002110-1200103213022321-1301300233212032-2111331133323100"></a>

## network_prefix property — dhcp_networks / 011132311322 / 4

Type: `"string"`. Optional.

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

Upstream description:

Exclusive with \[\] Network Prefix to be used for IPv6 address auto configuration.

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

<a id="canonical-3020212000003323-1130110030033103-3222002031132132-0211300211003231-1010120033103133-2011323111133323-1303131131112000-3103320320022030"></a>

<a id="canonical-1313103312223322-3312313323113010-1322323212203220-3331223302110122-2120310333020330-2322103221000202-3000001131300123-0312121033023010"></a>

## pool_settings property — dhcp_networks / 011132311322 / 5

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-004.md#canonical-3000223213321230-1022320121030122-1123230031101333-2330302220022310-2220002023033222-3303233133220131-3223232122000010-2013120200333013): complete subsection reference.

<a id="canonical-3211200222101301-2323320030203323-1132220331312330-1020322122100301-2222000333230201-3222013233103320-2132022233123333-2300022133022021"></a>

## Next pages — dhcp_networks / 011132311322 / 6

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-004.md#canonical-3000223213321230-1022320121030122-1123230031101333-2330302220022310-2220002023033222-3303233133220131-3223232122000010-2013120200333013)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3000223213321230-1022320121030122-1123230031101333-2330302220022310-2220002023033222-3303233133220131-3223232122000010-2013120200333013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302230211302102-3001321203313001-0111121123312312-2113020023111310-1132021230221333-1333103031133211-2130120113133010-3202131310030201"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 232021112332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-2012320123203311-0302032203311200-3320100122122123-0202320332220111-0132301203311320-1133201022302233-3323203103030130-0023123031033230"></a>

Type: `"object"`. list nested block, Optional.

List of non overlapping IP address ranges.

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
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
pools {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221331022322111-1002203001031332-2301203102331300-0101120003312313-2330333320002312-3102312211000333-1323123020202233-0011000311233122"></a>

## Direct properties — pools / 232021112332 / 3

<a id="canonical-0022030110131312-2001233213102322-1030011222223322-2232000210222312-1220010111203002-2110313323133323-0101333223133221-3300001330001110"></a>

<a id="canonical-1332222202201223-3332222213332003-2223112201331111-2103320133303233-1220220122303203-3222303220131330-0013010110111233-0311331131300231"></a>

## end_ip property — pools / 232021112332 / 4

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0230210220102313-1033122211021231-0032301003221112-2202332131220323-2212122002221131-1201311333111320-0300230321112333-1220021030031220"></a>

<a id="canonical-3210311313201012-1320203103201123-2033223232133200-0220300113202233-1130233031312022-0033120122203212-2200220200113210-2113032032312001"></a>

## start_ip property — pools / 232021112332 / 5

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2332011202321123-1320013301121202-1010123033021332-1113002031123300-2030321201001230-1221122133103121-2012103311101210-3312321203000202"></a>

## Next pages — pools / 232021112332 / 6

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-3113033113121313-3220101011103010-2000122203133121-2230020013013203-2100330022201102-2230033201033003-3321220110000022-1333330200210020)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1330001013210123-1120220010020113-3202320223302023-2231201331210030-0101122001201110-3230010032333130-0003121112232331-0031022010210210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210233113311132-3331113310032322-2213001323312321-3202001113101123-0123001102303321-0321303133122030-0030332101000112-1202113323231031"></a>

## aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 313330103330 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-004.md#canonical-1321033203010221-3121211103202303-3032120002333100-2101003122202213-2102110102203103-3202210230332001-0313021102012230-1012303131321123)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-004.md#canonical-1322100203220121-1102123310103012-0200001002111031-3231022212100213-0323300231032003-3102211000233311-0311220210212102-2021132131330111)
- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1220231210331112-1320022003131313-0020113131220221-0222232233113031-3020120112030112-2231000332321203-2003230320012020-2302010023123330"></a>

Type: `"object"`. single nested block, Optional.

Map of Interface IPv6 assignments per node.

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
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023202013111221-3101202212120200-3301311302032010-3300303131130123-1002112331003223-3032313321201213-0212332332232213-0103313203010001"></a>

## Direct properties — interface_ip_map / 313330103330 / 3

<a id="canonical-0111120002001103-3312232210112020-2112230032100211-0222032023023310-1023213111102231-0001333313232002-3222203002233033-0031323011201000"></a>

<a id="canonical-1202033130233113-3230302110312030-0011001130303300-2103003302002332-3003031122230323-3213333132032030-2330210203033202-3331301030101022"></a>

## interface_ip_map property — interface_ip_map / 313330103330 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv6\":\"true\"},\"values\":{\"format\":\"ipv6\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 64
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "64",
      "ves.io.schema.rules.map.values.string.ipv6": "true"
    },
    "values": {
      "format": "ipv6",
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv6": "true"
  }
}
```

<a id="canonical-0133222231330330-2322211121000011-3132221200231000-0222330303113011-1231330203031130-3112103211010200-0100101110012230-3321101013132300"></a>

## Next pages — interface_ip_map / 313330103330 / 5

- [aws.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-004.md#canonical-1213102302120202-0322321122022233-3213203103000223-0212202023213131-2220102202132003-0002301030211012-2113320222031020-3333231020103110)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0202022030000330-3101201300011103-3023131130202300-1321001033331112-0130022201000002-0123211120301111-2013310230200001-0330223213013320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213003131220320-0121101102010223-1003213131132301-3202332310113111-0300300000230322-1203313132201203-1100010100121212-3110001220213011"></a>

## aws.not_managed.node_list.interface_list.monitor — monitor / 213002223200 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.monitor

<a id="canonical-2103102133300200-3101200031022323-3311311303020231-1021231333133222-0311223003133233-0233211212023120-2232332212331020-1010023022330123"></a>

Type: `["object", {}]`. Optional.

Link Quality Monitoring configuration for a network interface.

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
monitor = {}
```

<a id="canonical-1303132321331300-1200202101231122-0202223202300223-3301312102203021-0312322333320323-1230031011322113-1323301023232230-3332100230122220"></a>

## Direct properties — monitor / 213002223200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322000203222313-2211230021003313-3023312102213313-3133130333221223-3012012310313320-2210221333211113-1332013103331220-1133110223221323"></a>

## Next pages — monitor / 213002223200 / 4

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0203330022300012-3030233012313303-1203203000030122-0323130220203010-1002312111330030-2202230202112120-0330211133301321-2130311201310200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131302120312011-0311012212211031-3232333130010121-3110132302123200-3110222211330030-3201002023231102-1110132133202202-2202133212112031"></a>

## aws.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 110103220213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1032233321320230-2011232121222233-2220220012303231-0213322111211332-3022322013322131-2211013122112333-2133111113301321-0200130222302210"></a>

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
monitor_disabled = {}
```

<a id="canonical-0112121021313121-1012202112001301-2000023000030300-3133212020312300-1122112312231110-0002213211331112-3313022330030212-0323121010311312"></a>

## Direct properties — monitor_disabled / 110103220213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100332103322211-3022322022301011-0002223102231100-3020001333133211-1312100021031110-1321211032323002-0013002202102230-0202031230300332"></a>

## Next pages — monitor_disabled / 110103220213 / 4

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333120001133033-0002311110230320-3111332213210111-3203313102221011-2130213030331021-0121121200303230-0221221011233020-3203030110220312"></a>

## aws.not_managed.node_list.interface_list.network_option — network_option / 320133220323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.network_option

<a id="canonical-3030030013132130-1202033130300221-0100123011313223-2011300032112311-2121010021221100-2220232212133111-2003023322330303-2330033333210012"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional.

Upstream description:

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network")}
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
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\"]"
}
```

Terraform syntax:

```terraform
network_option {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303213102323332-0103202200021233-3102001312301323-1100233031031021-2231103122122231-3021331330232103-1332101111200012-3201211030301132"></a>

## Direct properties — network_option / 320133220323 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2000120230120320-2021221112131120-1103013231233302-2020123001223230-3121032210320012-3220012201020031-3202202202320132-3011032112233213): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2020011030112311-1110111103312323-0002200102322300-0103131130332123-3030120131130222-1033303132301322-2300010210201223-2213013131003221): complete subsection reference.

<a id="canonical-2133120010200000-2331212101002233-0301010220030203-0303223203000332-3001033202032033-0222202311133120-2101233211312322-1021120331111113"></a>

## Next pages — network_option / 320133220323 / 4

- [aws.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2000120230120320-2021221112131120-1103013231233302-2020123001223230-3121032210320012-3220012201020031-3202202202320132-3011032112233213)
- [aws.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-004.md#canonical-2020011030112311-1110111103312323-0002200102322300-0103131130332123-3030120131130222-1033303132301322-2300010210201223-2213013131003221)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2000120230120320-2021221112131120-1103013231233302-2020123001223230-3121032210320012-3220012201020031-3202202202320132-3011032112233213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0312212310121313-1103330101122211-0310231230102133-1033120002313002-2301122133323301-0231212221231111-3001013022131220-1111133302001310"></a>

## aws.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 200001133203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133)
- aws.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3230010121111030-1131130311330130-1312101311003221-3022333322333203-2110311330231120-0012312033203033-0032330302320113-1203122010202000"></a>

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
site_local_inside_network = {}
```

<a id="canonical-2221332302131221-2120023032331200-0303011322112001-3102012222020132-0110331123320121-2013302230333000-1013001221212301-1111123301021323"></a>

## Direct properties — site_local_inside_network / 200001133203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002212300103023-2102301310021013-1320023322211232-1110033021300131-3303212322030112-0010020201301211-2320330002101233-1010101232321023"></a>

## Next pages — site_local_inside_network / 200001133203 / 4

- [aws.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2020011030112311-1110111103312323-0002200102322300-0103131130332123-3030120131130222-1033303132301322-2300010210201223-2213013131003221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202122312200231-1000101112301221-1301222330323122-0221210113012222-1301212111113002-1232033313112000-1100320210011121-3223001202211123"></a>

## aws.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 001323332012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133)
- aws.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-1122220221012310-0031201003102222-0122020022012211-2001320300033201-3302321322301310-3020001301221013-1132213121213122-3133322212032131"></a>

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
site_local_network = {}
```

<a id="canonical-0303011113013010-2332132312232000-2032013231221102-0311110113201210-0033201111002333-1300311211310302-1220110112011312-3122103302331233"></a>

## Direct properties — site_local_network / 001323332012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131310303012001-1002030312303032-3122313103022321-1101131333300133-2200032301321312-3201020022011030-1313110031231132-0233030221202032"></a>

## Next pages — site_local_network / 001323332012 / 4

- [aws.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-004.md#canonical-2230030020022233-0010033032210111-3202132331232202-3302031010211230-3001230231232310-2023022320321330-3310321102033231-3203203132332133)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2002311221033103-0201212323100003-2322100030302123-2310132032331013-1132300330330102-3220122223323020-2001023033031213-2211312023003220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002121110101321-2012332210011213-0300020202131013-0121203031020322-2230132102300100-0133132312231013-1230222103031213-0112302222113213"></a>

## aws.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 113021330211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-2313232112023130-0012201133313112-3200210311320120-3221010013021030-3101103021100320-3232223232203333-1011322013231130-2013002332230023"></a>

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
no_ipv4_address = {}
```

<a id="canonical-0113002123131102-3013003323101133-2031312003112311-1032132100132222-1332222322322321-1200211303020102-0202200300031101-3210022021000030"></a>

## Direct properties — no_ipv4_address / 113021330211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302103032231323-0020231130311303-1322330001023021-2103110122010012-2222101300303222-2120301200220312-3111021221311320-1131201232020002"></a>

## Next pages — no_ipv4_address / 113021330211 / 4

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1100230100302110-1330113121032003-3031310032013333-1200022322232311-0333211120100102-3031311232002231-0211032031122023-1011233232110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2011132012120013-0320313132131212-0112002320301331-3000331311120020-2223012100101233-0321302230301120-1231130331200113-1222032332320331"></a>

## aws.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 222303220210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0120331102320211-0013213122313101-0320322111330130-3302231130320323-2030130111013032-2113210033202030-3331003222301111-0320311103001202"></a>

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
no_ipv6_address = {}
```

<a id="canonical-2121013020103023-1102221100200012-0122033202003210-0020313020331310-0223123102333311-2022311333130220-3211232220032020-3023210302222010"></a>

## Direct properties — no_ipv6_address / 222303220210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332131301202020-0203211211203213-2222222212002301-2330101001211200-2233203323113313-2113220123203202-2133213211210022-2200221221121330"></a>

## Next pages — no_ipv6_address / 222303220210 / 4

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0133113233101313-1032101200201102-0200310132002132-1201332321002212-0100330020311220-0221003203110220-1313130130120313-2230302313120113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020133010012233-3032120023311202-2230321123101011-3300113132302210-2321330222121310-0131110102101121-2221302023330131-3200332031233200"></a>

## aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 131121101210 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3003233121310103-0022222323031111-2032032313201322-3022101223132132-3012101033010003-1000030320312010-3011323323330301-2220102303311323"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

<a id="canonical-0033132330330122-0223332120013020-2012100023031033-2312330230131011-3321002233122130-2303311210330220-0330030200123220-0200121001122011"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 131121101210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032221133303032-3030213223313131-2033122333232012-2110121333300320-3112111300332323-3210212313311212-2032100300102233-1103131312232230"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 131121101210 / 4

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2313131032101202-3131121321022321-3302102310201112-1000003301331033-3313123101011223-2201030101012232-0001013321233122-2300312230322120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223113212030022-1100122322121312-0120311321201121-1302231222012300-2323323212021021-1332332301323030-1031202013132102-1221112130222121"></a>

## aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 113321311022 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-1312012303113112-0033320032312112-1310311203230313-0212010200012111-0332303322222020-2023210123031203-3013201231211131-0100101202231200"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

<a id="canonical-1312000131222000-2022320023111113-2233022002133132-0101212023213232-1213100103022013-0313333121220023-3000203111332101-2201313122103033"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 113321311022 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2310233102020213-2212323322022113-3212130232121123-3011303303011122-1321111112333300-0310012300011223-1131210220301011-0000301201022130"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 113321311022 / 4

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0231330132121312-3133120010121331-3232322110010102-0123010002112203-0130321130031031-3232300002310113-2330021230332221-1113011321023231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232133200302321-2001330230033233-2331230323121023-1013022202302323-0203023113223302-2332313202212122-2123123331211213-1320331300000201"></a>

## aws.not_managed.node_list.interface_list.static_ip — static_ip / 322301332223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.static_ip

<a id="canonical-1211322112031130-2020031013232111-3111020320330211-0220131332302201-1131311302123111-2023133323122010-0020322331000331-0320003211131312"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210121012121323-2210010303310313-1133110201011332-3321221221233012-3302000101101232-1120133021022020-2231212203020100-2132312131322323"></a>

## Direct properties — static_ip / 322301332223 / 3

<a id="canonical-1321301211011302-1003331311223320-3322012220320200-0202211303113132-2110323201203231-0213310110120012-3023203301233220-3323131330313232"></a>

<a id="canonical-0303210302213211-0230331203220021-3120011230031010-0230011010313200-0202131033111021-1113213003201323-2202101103121302-1230320131321022"></a>

## default_gw property — static_ip / 322301332223 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-2022333123210130-2120022111321301-1022031100103300-0230303122202020-2221223011322212-2013132013020310-0310211203303112-0201103030331220"></a>

<a id="canonical-0032123022131213-3200323020032301-2311320010000031-3123012112200223-3301212120002121-3033120210331321-1021002011200030-0212232032102001"></a>

## dns_server property — static_ip / 322301332223 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2223132013021332-3201232303030220-2113002021232033-1321213233210121-0002300023312030-3231233323120301-1213300020233023-1102312332110330"></a>

<a id="canonical-1330113003132302-2203320310322002-0322313111123302-3003133202123210-3310030221112131-0013310331210201-3211222130310232-1112103103010311"></a>

## ip_address property — static_ip / 322301332223 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-1303222211322213-1223023013003120-3022221212200100-3330331212300101-0200101103331133-3201320330011201-2011111033012003-3131333122201310"></a>

## Next pages — static_ip / 322301332223 / 7

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211221222012231-1031211323130022-3111113023220202-2101223223302200-2321322133120133-1301010220322123-3203302200011331-1100311012223013"></a>

## aws.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 100032031230 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3031322233110102-3112111213002112-2222001121112331-1332011330320131-1201233021011221-3210311303220002-1202213322300210-2011312233303103"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cluster_static_ip",
    "node_static_ip")}
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
  "x-ves-oneof-field-network_prefix_choice": "[\"cluster_static_ip\",\"node_static_ip\"]"
}
```

Terraform syntax:

```terraform
static_ipv6_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300020000111302-2221000103103232-0332013212313322-2303210100012202-0220021302303230-0203033212300000-1132030100101333-2313002320232232"></a>

## Direct properties — static_ipv6_address / 100032031230 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-004.md#canonical-3022220120333102-1313110210000232-1003030211221203-2213022030201202-0130221133131023-2131323010220102-2001223113312331-3013112231301223): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-004.md#canonical-2033311313310121-0130101223302032-0330200200112233-2200300112213220-1220320032323101-1130132100031330-3211210310210330-2132110220320031): complete subsection reference.

<a id="canonical-2111311213321033-2120202221331220-1000031311020102-1033023313232322-2221333200300232-0211101020122121-3121100233010001-1312201021331032"></a>

## Next pages — static_ipv6_address / 100032031230 / 4

- [aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-004.md#canonical-3022220120333102-1313110210000232-1003030211221203-2213022030201202-0130221133131023-2131323010220102-2001223113312331-3013112231301223)
- [aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-004.md#canonical-2033311313310121-0130101223302032-0330200200112233-2200300112213220-1220320032323101-1130132100031330-3211210310210330-2132110220320031)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3022220120333102-1313110210000232-1003030211221203-2213022030201202-0130221133131023-2131323010220102-2001223113312331-3013112231301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120201121003021-1302311312201202-3113333000233030-1130111331131233-1122100223003030-1331323203320122-2332102103221030-3002001212132311"></a>

## aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 121103311333 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303)
- aws.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-0301012030120230-2022110233322201-2021001001233203-0301202020030311-1121031023321022-0203023222001223-0222010112032213-3222323313213101"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for cluster.

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
cluster_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202020033302302-1121212021132000-1311010033220320-1002233131120221-3112112020310022-2303233220213303-0313133313010332-0101203030010032"></a>

## Direct properties — cluster_static_ip / 121103311333 / 3

<a id="canonical-0123203202112210-0130133020301112-2032313220102322-1022011130131230-1310322031230033-2330331233020232-2011313023222200-0322221303230020"></a>

<a id="canonical-1120031132112100-3121302321122312-2123113121120203-1230222213021010-1310113213130302-1003021010233312-1322321023312210-1013102131000232"></a>

## interface_ip_map property — cluster_static_ip / 121103311333 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"128\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "deterministic": true,
    "keys": {
      "maxLength": 128,
      "minLength": 1,
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.max_len": "128",
      "ves.io.schema.rules.map.keys.string.min_len": "1",
      "ves.io.schema.rules.map.max_pairs": "128"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "128"
  }
}
```

<a id="canonical-2212331233032100-3101113121122021-0121001021002002-0133030113230121-0032310222031021-1103120003121011-3323302230222301-3013030323021231"></a>

## Next pages — cluster_static_ip / 121103311333 / 5

- [aws.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2033311313310121-0130101223302032-0330200200112233-2200300112213220-1220320032323101-1130132100031330-3211210310210330-2132110220320031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310311003203020-1202010231110310-0300231022121133-1122232132210022-2121320313133130-1023000320031230-3010000201323032-0121230302230100"></a>

## aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 203103110000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [aws.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303)
- aws.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-3133032011000031-2032212332331133-0011131100330113-0132111011101013-0222012021300113-1012331132312120-3020320033131200-0110300220211030"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_address")}
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
node_static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223020222303322-3012210320111121-2032131001100212-1010120212112001-2120103012311132-2323130030013003-3322213020002013-1013211202122231"></a>

## Direct properties — node_static_ip / 203103110000 / 3

<a id="canonical-3313303102333323-1000033232322130-3221300030222223-0002102230111232-3033312123002321-3232210220002221-2330230013221132-1301003211220213"></a>

<a id="canonical-0230011130331201-0300323221022000-2011230220120113-0031112300330123-3202100320301120-2003030231112212-1311310230210111-1100112002121100"></a>

## default_gw property — node_static_ip / 203103110000 / 4

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
  validators.IPValidator(),
}
```

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
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="canonical-3211301103322020-3201312023012333-1012322302012122-1102000322033300-0020131222121022-0010132312113001-2321233223231011-2121133123010123"></a>

<a id="canonical-2101332320101112-0211233230232301-2032203131111102-1222130132021030-1213213122133121-2131331031031002-2330312302210010-0121120001200332"></a>

## dns_server property — node_static_ip / 203103110000 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3231101210212312-0013033123122113-0013023230011032-1332313330302202-1012301210112023-0231000322022301-3113123100111131-2002323011011100"></a>

<a id="canonical-1230220020111111-2110030110112111-1332010221300112-1320301203300022-3010330032202301-0103132101211233-3002012022332032-2021303211021102"></a>

## ip_address property — node_static_ip / 203103110000 / 6

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.CIDRValidator(),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "cidr",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip_prefix": "true"
  }
}
```

<a id="canonical-1031030030222233-2233122110110012-0330331123013131-3021012010022323-2302231031132311-0023200021111023-0211301233212012-2010312221021010"></a>

## Next pages — node_static_ip / 203103110000 / 7

- [aws.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0022010121023230-1010112111012231-3100032313003023-0321102200222000-0131001132132311-0202202301120022-1131012000222230-3100232030323303)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1021322130033210-3013010303233300-1230322322220200-1030112032003212-0101211221120122-0231232330311100-1012100021331010-1131310012212000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212223323331003-3331021333323233-3102133321130300-0302212233111223-3200123102232120-0333313003002232-3133003200021103-2231111122222202"></a>

## aws.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 010211300020 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [aws](resources--securemesh_site_v2--reference--group-003.md#canonical-2110013303321312-2120002202233001-3022012321110003-1003102021021001-3033100301110222-3213122033202023-1003232312003032-0322102033301223)
- [aws.not_managed](resources--securemesh_site_v2--reference--group-003.md#canonical-2331321022212133-2200013203303300-1200323201323000-2130031003123312-2022021011100233-3312320201110322-3202031200333201-1202113222033310)
- [aws.not_managed.node_list](resources--securemesh_site_v2--reference--group-003.md#canonical-3220113033302221-1033300203131122-0100011132103232-3011112022310122-2211321200333310-0020233221231131-0331000210210021-1000031313120202)
- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- aws.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-2231121102200133-1231331133233320-1010103110021220-3312002310101212-3302330131200103-0010113033123231-1132233031300123-1211320331311110"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device",
    "vlan_id")}
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
vlan_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220130301322111-2232002332021323-3302123230021121-3302312132111101-2102322131101310-1002312320111100-3122321303103230-2021223000131331"></a>

## Direct properties — vlan_interface / 010211300020 / 3

<a id="canonical-0011111023321321-3020311003023223-1101311113101022-0331303310332231-2013103111222230-1020232032322303-3033022110000030-1132023102230021"></a>

<a id="canonical-2231210332010322-1012103322323303-0110213212013210-1131302203113130-0032123232312102-2032023021132031-0000233103102213-2332100110130333"></a>

## device property — vlan_interface / 010211300020 / 4

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-2033222230111311-1200123010002111-1210013123100133-2011030330230133-3211130330133320-3233112120011110-0022220100030321-3123203021133133"></a>

<a id="canonical-0210133022232102-2233010333213103-1303303000200001-0230132003033101-0012313301232231-3203032221311232-1212300201332333-2131210232132120"></a>

## vlan_id property — vlan_interface / 010211300020 / 5

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 4095),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 4095,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "4095"
  }
}
```

<a id="canonical-1211130211120321-1203033111011202-3110110120211120-2231332032020102-3320102232032102-0221013223030032-1213222123031021-1103202303000303"></a>

## Next pages — vlan_interface / 010211300020 / 6

- [aws.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-003.md#canonical-2213113003010131-2020030332110203-1120313022121123-0000013032112123-0201321120010032-1310223132203322-0313231331223132-3333320301013210)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301322312102233-2233003103021331-1111032022012001-1201302320120002-0101313012210022-1001221333110023-0110201203131313-2331132332132302"></a>

## Azure — Azure / 121121021323 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- Azure

<a id="canonical-2113313310322103-1200312233131310-1321033010323123-2021323211323313-1232210302330301-0320012003322032-2101302210230133-0231203211203233"></a>

Type: `"object"`. single nested block, Optional.

Azure Provider Type. Azure Provider Type.

Upstream description:

Azure Provider Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-orchestration_choice": "[\"not_managed\"]"
}
```

Terraform syntax:

```terraform
azure {
  # Configure direct properties listed below.
}
```

<a id="canonical-1320120203112011-0123101013310021-2222132021131021-3001003321320303-2222020213333213-2210300232012131-0031302123302220-0311102110200132"></a>

## Direct properties — Azure / 121121021323 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121): complete subsection reference.

<a id="canonical-3332131203202131-2203121130021221-0322020120113321-2220011312103230-0103332100200010-3232230332102102-1212220230231210-2002002300000002"></a>

## Next pages — Azure / 121121021323 / 4

- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012301322210222-1320332100110110-3121113023111002-1133103311333103-2321231212112313-2121110012031331-1001300020010012-0100221013111013"></a>

## Azure.not_managed — not_managed / 000133211012 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- Azure.not_managed

<a id="canonical-3311323323211312-2031030310212312-3303111233210120-2313313221333333-2322131221311100-2200310021203103-0021210031232003-2003320202133302"></a>

Type: `"object"`. single nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
not_managed {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103312121203220-3122323110313022-2313020231011220-1332030211123210-0132321311233231-2211200221301323-1002300303320020-2302121203213200"></a>

## Direct properties — not_managed / 000133211012 / 3

- [node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300): complete subsection reference.

<a id="canonical-0320200230210303-0333130321011211-3203221113110202-1103102133000301-1101321320320313-2110000320201010-2132332221031031-2333013112203311"></a>

## Next pages — not_managed / 000133211012 / 4

- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2120111333010231-1210202030032113-2330133012203332-3301201132233120-0301120003123320-3013122200223021-1302311332313001-0320003212110332"></a>

## Azure.not_managed.node_list — node_list / 331123223233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- Azure.not_managed.node_list

<a id="canonical-0000201231203121-1020322001201031-3331322333130221-0323030212101200-0220010320211223-0013233002033113-3122123000232003-1000331310301021"></a>

Type: `"object"`. list nested block, Optional.

Section will show nodes associated with this site.

Upstream description:

This section will show nodes associated with this site. Note: For sites that are not orchestrated by
F5XC, create nodes in the chosen provider. Once a node is created and registers with the site, it
will be shown in this section.

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
node_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020300223230021-0100201112010123-1131331210231123-1130331230003122-2230212020310020-2113103121311101-1220331031203000-0313022132121200"></a>

## Direct properties — node_list / 331123223233 / 3

<a id="canonical-1102330002203330-1303323223301011-0120230103213132-1210120321022223-1212231303322030-2301231303111023-1011023323232001-1112112012231303"></a>

<a id="canonical-0103100232201202-0002231131110111-3330032320010233-2231331021302203-2320300332321011-0303332231331210-1021133212110122-2222222321223323"></a>

## hostname property — node_list / 331123223233 / 4

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Upstream description:

Hostname for this Node.

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
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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

- [interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132): complete subsection reference.

<a id="canonical-1313330310032012-3222112332330102-0310230031103130-0222321011133333-0213102322033200-3330002033033313-3233233222233021-1312120301212031"></a>

<a id="canonical-1330322101022132-2213112223320211-1202122210332133-2122111120301301-0013123233300033-2221230332310320-1323211122112312-1113030211012131"></a>

## public_ip property — node_list / 331123223233 / 5

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Upstream description:

Public IP for this Node.

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

<a id="canonical-0302011221132222-1012002221310220-2333200010221222-0010103113203222-1122101133332003-3203223302123123-1023220333323220-3303210013332100"></a>

<a id="canonical-0032011031332220-2133302032101101-2322233312032113-0220010232301201-0333320123221030-1022200113030232-3310121303302212-2022202230003033"></a>

## type property — node_list / 331123223233 / 6

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Upstream description:

Type for this Node, can be Control or Worker.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("Control",
    "Worker"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "Control",
    "Worker"
  ],
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
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"Control\\\",\\\"Worker\\\"]"
  }
}
```

<a id="canonical-2210021021203131-0210102021221320-1212012313110033-1121013322333312-0121000121312202-0222222031100223-2000303000223012-1210311320000032"></a>

## Next pages — node_list / 331123223233 / 7

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022103300200033-1313023230002220-3001032012333102-2003312313131330-1220303123310113-2110123113000212-0203230110111101-3332002010232221"></a>

## Azure.not_managed.node_list.interface_list — interface_list / 322330310001 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- Azure.not_managed.node_list.interface_list

<a id="canonical-3001212221110111-0123102032113101-0331320323300033-0223200131001201-1213011121000031-3121020031202011-3103111030000110-1110002031101302"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("bond_interface",
    "ethernet_interface"),
  validators.ConflictingListObjectAttributes("bond_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "no_ipv4_address"),
  validators.ConflictingListObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingListObjectAttributes("ethernet_interface",
    "vlan_interface"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingListObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingListObjectAttributes("no_ipv4_address",
    "static_ip"),
  validators.ConflictingListObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingListObjectAttributes("site_to_site_connectivity_interface_disabled",
    "site_to_site_connectivity_interface_enabled")}
```

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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131312221200010-2233121110223113-3021333221012000-1003011232323102-0000332220322312-1202302100232303-3332301213121003-3001230010323100"></a>

## Direct properties — interface_list / 322330310001 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013): complete subsection reference.

<a id="canonical-3231020130113203-1320203303002332-0333212130111220-0001112021032131-0103210223203120-1130312010000031-1031103223330012-3101313013123302"></a>

<a id="canonical-0322321220323321-0022001221232101-3100110231101310-1022023013131233-3210301210111221-1200003212200212-1110233101223231-0000010323300130"></a>

## description_spec property — interface_list / 322330310001 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-004.md#canonical-0233210220123210-3110212232020123-2012020010311011-0203333030021302-2120031221013311-0021001010211031-0111131332203331-3112332000122100): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-0122222112012221-1303023111032023-0023302003202230-1000033332201233-2113112000202201-1120210213111000-3210131222311020-3130312203030233): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131): complete subsection reference.

<a id="canonical-2021123301033100-0332030210231320-2321111330230113-3322323310102032-1213132202121020-3112302231003112-2321221313031322-1002001312322220"></a>

<a id="canonical-0301023330300310-1202013101302221-0030030102323312-1310012110120033-3310310101212222-2301313021230022-0222203213220011-3123102233230312"></a>

## is_management property — interface_list / 322330310001 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-0331103302113330-1223121200033031-1123030301131331-0322332233120100-2110222221103202-2331030223322100-2333111111320222-0120000130201130"></a>

<a id="canonical-0033312010103130-0032332032331313-1212133030102313-3320020111332021-2102020000223031-2232313133000311-3020000201221232-2012102002121321"></a>

## is_primary property — interface_list / 322330310001 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-2301232312232313-1200122320002333-1322032213120200-0001133210003233-3102001123100221-2032122130121012-3213011102002013-3122221000032032"></a>

<a id="canonical-3130020022200021-0230303132302033-1330232033031030-1213001120032012-3200223003133101-0222303301022023-1233221102222311-3112212002300121"></a>

## labels property — interface_list / 322330310001 / 7

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":16},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"64\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"16\",\"ves.io.schema.rules.map.values.string.max_len\":\"64\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":64,\"minLength\":1,\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 16
    },
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
      "ves.io.schema.rules.map.max_pairs": "16",
      "ves.io.schema.rules.map.values.string.max_len": "64",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 64,
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
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "16",
    "ves.io.schema.rules.map.values.string.max_len": "64",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

- [monitor](resources--securemesh_site_v2--reference--group-005.md#canonical-0012332332002200-3102111303112310-1332033032331322-2313132132223022-3223201320211202-0030102022000222-1313012110000230-0311203322223203): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-005.md#canonical-3212332310213213-3300033111113300-0101112023103310-2211012333331012-2302013210330030-0233133220330322-0103220102103300-3321310221201212): complete subsection reference.

<a id="canonical-3023222333001302-2203001322233322-1301100023000111-3021321012010221-3201211220330123-2330230023003303-0001110021033311-2123322332303203"></a>

<a id="canonical-0323110000003200-3320333012112033-0003023123120213-3110330113232231-2122210020332311-1131231201122302-2310331221301321-2223000130031000"></a>

## mtu property — interface_list / 322330310001 / 8

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 8000},
  ),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 8000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-8000"
  }
}
```

<a id="canonical-3002222123311230-2313113211100112-2332231330331233-2300321301231003-3121013122003011-2202213030312032-1202120132200323-2111112212003213"></a>

<a id="canonical-3013231221101221-0211021022302200-3112311201132233-3231232001003112-1223100312103302-3231130020132010-3313030331010033-0322323210020001"></a>

## name property — interface_list / 322330310001 / 9

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Upstream description:

Name of this Interface.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [network_option](resources--securemesh_site_v2--reference--group-005.md#canonical-0203123231033121-0031100313212100-1311300120012133-0333010013103231-1101031300033123-1200211110330200-0020223120123010-2311232113313000): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-005.md#canonical-3301000232322233-3022320013231312-0113322101120131-1120100222101201-1312221213013300-0002303113231311-1330310113122223-0023303333221230): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-0012211330030212-0013201021031130-0100233021002233-2232323210033022-0002100312201203-2112113030200010-1012313312221110-0032232310211201): complete subsection reference.

<a id="canonical-2223303000122103-0131012030331000-1211131033300313-1001110031210212-0011023131130210-1123231112322200-2101302310022303-2230021331103212"></a>

<a id="canonical-3330200232200232-3310110020110022-0233323132132100-1030132121102100-3102323103003330-2121102231223030-0312113233112222-0010331100112302"></a>

## priority property — interface_list / 322330310001 / 10

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Upstream description:

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 255),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-005.md#canonical-1113102132323123-2300332322133033-1032231331233123-3200023012010021-0101001300032310-0232233203013222-3021310023000210-3113003022001010): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-005.md#canonical-1311333023133210-0132231022023003-2101022021332322-1133323223020302-1303011201300233-1302230132320223-3232122321102132-2103323232033311): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-2110310311302232-0030010210011223-3113233222301032-3320311330331330-3033323202220023-2012311330121312-0210322021013113-3133120212032121): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-2211320003103001-3111012233312332-3202121123112303-3201112120030102-2311032010301133-1333023311102331-3301130003133313-2123002312121123): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-0331212221203100-2112013330321301-1013303220131023-0300102331331333-2120332302201122-3020013210210312-0111133211021002-3130221213131123): complete subsection reference.

<a id="canonical-2311202132203203-3203022203322213-1011203203313000-0022132233132233-1312203103223010-2220220032302010-3221033122231102-0030021211202021"></a>

## Next pages — interface_list / 322330310001 / 11

- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- [azure.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-004.md#canonical-0233210220123210-3110212232020123-2012020010311011-0203333030021302-2120031221013311-0021001010211031-0111131332203331-3112332000122100)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [azure.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-0122222112012221-1303023111032023-0023302003202230-1000033332201233-2113112000202201-1120210213111000-3210131222311020-3130312203030233)
- [azure.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-005.md#canonical-0332121112210012-0012200002300033-3201102113120313-2331213210112221-1330333101001021-1221011222030012-1012011021312312-3312300133030131)
- [azure.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-005.md#canonical-0012332332002200-3102111303112310-1332033032331322-2313132132223022-3223201320211202-0030102022000222-1313012110000230-0311203322223203)
- [azure.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-005.md#canonical-3212332310213213-3300033111113300-0101112023103310-2211012333331012-2302013210330030-0233133220330322-0103220102103300-3321310221201212)
- [azure.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-005.md#canonical-0203123231033121-0031100313212100-1311300120012133-0333010013103231-1101031300033123-1200211110330200-0020223120123010-2311232113313000)
- [azure.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-005.md#canonical-3301000232322233-3022320013231312-0113322101120131-1120100222101201-1312221213013300-0002303113231311-1330310113122223-0023303333221230)
- [azure.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-0012211330030212-0013201021031130-0100233021002233-2232323210033022-0002100312201203-2112113030200010-1012313312221110-0032232310211201)
- [azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-005.md#canonical-1113102132323123-2300332322133033-1032231331233123-3200023012010021-0101001300032310-0232233203013222-3021310023000210-3113003022001010)
- [azure.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-005.md#canonical-1311333023133210-0132231022023003-2101022021332322-1133323223020302-1303011201300233-1302230132320223-3232122321102132-2103323232033311)
- [azure.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-005.md#canonical-2110310311302232-0030010210011223-3113233222301032-3320311330331330-3033323202220023-2012311330121312-0210322021013113-3133120212032121)
- [azure.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-005.md#canonical-2211320003103001-3111012233312332-3202121123112303-3201112120030102-2311032010301133-1333023311102331-3301130003133313-2123002312121123)
- [azure.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-005.md#canonical-0331212221203100-2112013330321301-1013303220131023-0300102331331333-2120332302201122-3020013210210312-0111133211021002-3130221213131123)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2223012111331220-1223220013212323-0100231102030120-2121021312003230-3211003231121001-2203001312311030-0231310131221113-3112203301321310"></a>

## Azure.not_managed.node_list.interface_list.bond_interface — bond_interface / 210232312203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.bond_interface

<a id="canonical-3101212112023020-3121012313111230-2123130210332220-2020223311300220-3020202001012033-0012220101202113-2230220103322133-2330101100301032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Upstream description:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingObjectAttributes("active_backup",
    "lacp")}
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
  "x-ves-oneof-field-lacp_choice": "[\"active_backup\",\"lacp\"]"
}
```

Terraform syntax:

```terraform
bond_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333033230012012-2303303112202032-1032233303031333-3330022302111133-3131203023330201-2220230012021302-1221201131000033-1122233131102033"></a>

## Direct properties — bond_interface / 210232312203 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-004.md#canonical-1003232300130320-1313033033131113-2333111300321201-2330222311021022-1323130233313000-2133112133002020-3011131200312232-0323010330000301): complete subsection reference.

<a id="canonical-0333200131003123-3121331311202201-1022230301231023-0313313201020122-0331112010131210-0212303323220201-1223213111131003-2113011223210003"></a>

<a id="canonical-3302310030101021-0131122330232330-3130112211212110-2131321021201320-1311033210211302-2201202202112002-0230103321321311-2131123200223013"></a>

## devices property — bond_interface / 210232312203 / 4

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 8),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [lacp](resources--securemesh_site_v2--reference--group-004.md#canonical-1123013312002020-0313331111332212-0330332333312332-1311121123313131-0330220100212300-3321122031311320-1010302101330222-3203301200222322): complete subsection reference.

<a id="canonical-0003203320212103-1332002012020220-0032213131000310-1210132132233132-3103000133031020-0312201020322032-3231333311020320-3010201210032030"></a>

<a id="canonical-2231323301122231-1311202332113223-0000021311120131-0312220022321322-1112211323011302-1210231202030300-3131331303322220-3312331033133303"></a>

## link_polling_interval property — bond_interface / 210232312203 / 5

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Upstream description:

Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(500, 5000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minimum": 500
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "500",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-2000232323311110-1010301123220221-2032222223120010-0020220221133330-3123113322332230-1212311101322223-1111113211212023-0102233303021000"></a>

<a id="canonical-0311103011323011-1230033320330001-3223001123000211-3021231201102033-3312013111333032-3021230012331112-1033132311112302-1221223200310321"></a>

## link_up_delay property — bond_interface / 210232312203 / 6

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 1000),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 1000,
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
    "ves.io.schema.rules.uint32.lte": "1000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "1000"
  }
}
```

<a id="canonical-0320230133003230-0113113322213310-1010233112302231-1022233020100133-2023131312321111-0031202001023030-2203232002223332-3021010110232222"></a>

<a id="canonical-0231110013310111-2001031301032132-0022011130012202-0320212030321312-0021010302331112-2301131012200101-0001122311330230-2121232233003310"></a>

## name property — bond_interface / 210232312203 / 7

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

Upstream description:

Name for the Bond. Ex 'bond0'

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
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
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2021030013111212-2133223132131020-3302031213231103-0021313130033120-1000121311233311-3032232302230121-3000132120003001-2210333321022300"></a>

## Next pages — bond_interface / 210232312203 / 8

- [azure.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-004.md#canonical-1003232300130320-1313033033131113-2333111300321201-2330222311021022-1323130233313000-2133112133002020-3011131200312232-0323010330000301)
- [azure.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-004.md#canonical-1123013312002020-0313331111332212-0330332333312332-1311121123313131-0330220100212300-3321122031311320-1010302101330222-3203301200222322)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1003232300130320-1313033033131113-2333111300321201-2330222311021022-1323130233313000-2133112133002020-3011131200312232-0323010330000301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320301333121003-0022033013223021-1023330321302211-3322010210132112-0110122231030120-1210131112030031-2231231321120322-0332023021300203"></a>

## Azure.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 020331123130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- Azure.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1313330020303331-3000103232120130-1200313103122213-2323323330013121-2020220021300120-2032332203012123-3033233312012032-1202131311121002"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

<a id="canonical-1003110233211111-1331001312101031-2112220320011320-0220321030312003-0130113031131202-0201010100221110-2002131132031221-1133300223123101"></a>

## Direct properties — active_backup / 020331123130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102303232013333-3201020220131032-0113002230231022-0322221003102131-2003231131003002-2000133000311132-1213331301321201-3220211021002123"></a>

## Next pages — active_backup / 020331123130 / 4

- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1123013312002020-0313331111332212-0330332333312332-1311121123313131-0330220100212300-3321122031311320-1010302101330222-3203301200222322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013111132230101-3331221331212032-3022013121012311-1201131313000023-1332302131022130-1031100220030022-2300331003100303-3132010022213020"></a>

## Azure.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 000021110120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- Azure.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-3220123331000210-3021311020321321-0033222302112220-2010320103132213-0223230121000121-2000010222002010-0322232311002013-1301110323001123"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Upstream description:

LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate")}
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
lacp {
  # Configure direct properties listed below.
}
```

<a id="canonical-0223022120003221-1221130022101221-2001210320222320-0201120202320020-3111022032300332-0131321023122131-3320201111111302-2231013212221023"></a>

## Direct properties — lacp / 000021110120 / 3

<a id="canonical-1221131021313203-2301320332031133-3013223202032301-1232332322321323-2110230201311010-1003121031200132-2000211001230212-1021022130302012"></a>

<a id="canonical-0222033011303332-0231333233110310-1030310002012013-3103130233333120-3133330330011111-2322133330130213-0332313011132133-0322001033103120"></a>

## rate property — lacp / 000021110120 / 4

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "30"
  }
}
```

<a id="canonical-1033101022232130-0323333302203123-0021230232111021-3120201120220130-2132330013201230-0223030312203323-2132032211211221-0222002330021030"></a>

## Next pages — lacp / 000021110120 / 5

- [azure.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-004.md#canonical-3210121022123102-0230310100330222-2012212230131130-0221113330300013-2310110130011003-1230113012210110-3330133200022221-0302223132030013)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0233210220123210-3110212232020123-2012020010311011-0203333030021302-2120031221013311-0021001010211031-0111131332203331-3112332000122100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112211120202020-3031002320011023-1302012131310231-0032131003023021-1032312123223223-2220302002123320-3222312302303133-0132200110202320"></a>

## Azure.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 302121121302 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3312123233123023-0320233302331013-1221000332222323-3003021213303100-1003023123131311-3000231120210000-3111330200003320-1302300022210133"></a>

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
dhcp_client = {}
```

<a id="canonical-2022031031133200-3010003232132232-0232011201233323-1113233202220202-2203003001232112-3121320110222323-1131012321210331-0203031111322330"></a>

## Direct properties — dhcp_client / 302121121302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1220013100122322-0223221232212202-2321201300221120-1320102313231100-1232103331332003-2221003231010313-1101201123100123-0010131110000302"></a>

## Next pages — dhcp_client / 302121121302 / 4

- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210302033023302-1010020303101130-3232133210222113-1223321312330123-2101320113201201-0202123221011223-2221232323111233-1012002122232332"></a>

## Azure.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 121320012122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- Azure.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1332211133002300-3300231303222202-2023032201131321-2013102223023233-2112301320031133-1113101120100002-3220111303302110-1031220210103013"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Upstream description:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("dhcp_networks"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "automatic_from_start"),
  validators.ConflictingObjectAttributes("automatic_from_end",
    "interface_ip_map"),
  validators.ConflictingObjectAttributes("automatic_from_start",
    "interface_ip_map")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-ves-oneof-field-interfaces_addressing_choice": "[\"automatic_from_end\",\"automatic_from_start\",\"interface_ip_map\"]"
}
```

Terraform syntax:

```terraform
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322013223102301-1302113121213121-3120222033020011-0132332313100202-1223311300012212-2010230101032301-3110301002011330-3211112330330302"></a>

## Direct properties — dhcp_server / 121320012122 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-004.md#canonical-2233313322033011-0211202300302003-2201100111320103-0111123031303310-2330300033120012-1002331320131300-3021022321331321-1331233100121111): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-004.md#canonical-2103312320321332-3123311030031210-1222033012111322-3331020310201100-2231222131113301-2032000010323230-1321301330000221-2312030220312131): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301): complete subsection reference.

<a id="canonical-0303321320121331-0032222113332301-2222031200120120-2230311222012011-3330021312001310-0110002100213310-1113033332331001-0330213110312033"></a>

<a id="canonical-0310031100222002-3222333210203231-2212123030020310-1203331101330122-0133200111000112-3022323223032130-2013311303220203-2331221001132111"></a>

## dhcp_option82_tag property — dhcp_server / 121320012122 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-2213113310320323-2122221323313022-1133311103302212-1130122323332201-0110001022331122-1202013233003010-1122121013101130-1020012012301111"></a>

<a id="canonical-2102131130202303-2003223023202110-3221300111330303-3231313110010210-0203131100102312-3013201003222132-3101101221233222-1033103311012021"></a>

## fixed_ip_map property — dhcp_server / 121320012122 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":128},\"category\":\"discovery\",\"constraintType\":\"map\",\"crossEntry\":{\"uniqueValues\":true},\"deterministic\":true,\"keys\":{\"format\":\"mac-address\",\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.mac\":\"true\",\"ves.io.schema.rules.map.max_pairs\":\"128\",\"ves.io.schema.rules.map.unique_values\":\"true\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "cardinality": {
      "maxProperties": 128
    },
    "category": "discovery",
    "constraintType": "map",
    "crossEntry": {
      "uniqueValues": true
    },
    "deterministic": true,
    "keys": {
      "format": "mac-address",
      "type": "string"
    },
    "originalRules": {
      "ves.io.schema.rules.map.keys.string.mac": "true",
      "ves.io.schema.rules.map.max_pairs": "128",
      "ves.io.schema.rules.map.unique_values": "true",
      "ves.io.schema.rules.map.values.string.ipv4": "true"
    },
    "values": {
      "format": "ipv4",
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
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.mac": "true",
    "ves.io.schema.rules.map.max_pairs": "128",
    "ves.io.schema.rules.map.unique_values": "true",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

- [interface_ip_map](resources--securemesh_site_v2--reference--group-005.md#canonical-1220020203222101-3221122312100013-2111102322111013-0132220221333302-0010000311221302-0303001000030231-3332010302330313-2100010101223221): complete subsection reference.

<a id="canonical-3003322021220311-1030333322302102-1121023002121322-2120310213020101-3300001010121321-3302333200002033-1222021203210123-2033330312023331"></a>

## Next pages — dhcp_server / 121320012122 / 6

- [azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-004.md#canonical-2233313322033011-0211202300302003-2201100111320103-0111123031303310-2330300033120012-1002331320131300-3021022321331321-1331233100121111)
- [azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-004.md#canonical-2103312320321332-3123311030031210-1222033012111322-3331020310201100-2231222131113301-2032000010323230-1321301330000221-2312030220312131)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-004.md#canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301)
- [azure.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-005.md#canonical-1220020203222101-3221122312100013-2111102322111013-0132220221333302-0010000311221302-0303001000030231-3332010302330313-2100010101223221)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2233313322033011-0211202300302003-2201100111320103-0111123031303310-2330300033120012-1002331320131300-3021022321331321-1331233100121111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321113031002100-0102111323323220-2230130201023032-0312031001213022-2330302302000222-1032000023003220-2022211330001200-2021330212132312"></a>

## Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 223000100103 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3001133331122323-2002221223201201-0032222121323010-1233333002303123-1020201113032312-2112110120012213-3230210131122012-2221013010311122"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

<a id="canonical-0223121230213122-3323122032302333-3222220121121303-3332232022210001-0010301110000300-0323130300030032-3332003002321023-2112230333102100"></a>

## Direct properties — automatic_from_end / 223000100103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303001011113001-3130303300333311-2302113001322022-2032033311231203-3021233320112022-2322222320121122-3330122113313311-2210111031010300"></a>

## Next pages — automatic_from_end / 223000100103 / 4

- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2103312320321332-3123311030031210-1222033012111322-3331020310201100-2231222131113301-2032000010323230-1321301330000221-2312030220312131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202021000212222-1223300230010010-0032220231202213-0201103232210030-3030121122210120-0222311202010212-0101011232211230-1020313223213012"></a>

## Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 113311201110 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3010331130300002-0220130213203222-0011212120110132-3012313123211213-0310012202001000-3112211011131300-1223020333200133-3203201221000010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

<a id="canonical-1000223021011313-3021221101233201-2300023000231010-2003020112330331-3212210122031021-2333000221031201-1332330122332320-1010112213300032"></a>

## Direct properties — automatic_from_start / 113311201110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013202202112231-1030133112101220-2001333333323312-0112322210003100-2030213122312102-3231323010011101-0121021023130302-0110123121333220"></a>

## Next pages — automatic_from_start / 113311201110 / 4

- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1132323030301002-3223321123022021-1313022301301022-0222001102020103-2100310122010121-1310320203233120-1121301021032010-1020120303101301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200320233112232-0123320100202130-1103112221322221-2200002313220202-3211030102232102-3230110121220210-1021301200300232-3120131333022113"></a>

## Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 000022332303 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Azure](resources--securemesh_site_v2--reference--group-004.md#canonical-3112231321103113-3213320211011211-3322111123112332-1212020331202310-3010203031021330-2313022310003121-0000110202133003-0001212003121320)
- [azure.not_managed](resources--securemesh_site_v2--reference--group-004.md#canonical-0011132122020132-1002110233123122-3220321320120021-1023033233101232-3113010232021131-3000322003312101-2302210310101212-0223021132300121)
- [azure.not_managed.node_list](resources--securemesh_site_v2--reference--group-004.md#canonical-0333033310232323-1123132331021120-1230311301301220-3203300331111022-1230332212310020-2210002121103002-3320121222331110-1201312320311300)
- [azure.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-004.md#canonical-2000210110320113-2102113113123131-3121010303022222-2120130023111120-1001332303133201-0302313231012233-3202333332211123-1132030313311132)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- Azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-3010332202033300-2213203233231012-2331321000202000-0012033031213120-1311221210200003-3202300212023113-3001123010120032-3133132332120333"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("dgw_address",
    "first_address"),
  validators.ConflictingListObjectAttributes("dgw_address",
    "last_address"),
  validators.ConflictingListObjectAttributes("dns_address",
    "same_as_dgw"),
  validators.ConflictingListObjectAttributes("first_address",
    "last_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
dhcp_networks {
  # Configure direct properties listed below.
}
```

<a id="canonical-1300222202023103-0320023231000323-1032333122112223-0003113210301232-2331303301212223-3101213133010021-3010213012123202-1200001202230331"></a>

## Direct properties — dhcp_networks / 000022332303 / 3

<a id="canonical-2003201232100110-0120102221312112-1322031031233213-3213022021023230-0320100222111202-2033312210330211-1202100221123120-3200312313120003"></a>

<a id="canonical-0132103001103221-1200222200123002-3001232303322221-0021330131212310-2011132022131021-2121330200321133-2110320202322003-2222112331100231"></a>

## dgw_address property — dhcp_networks / 000022332303 / 4

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1201130332003213-2232032202303111-1023013123321103-0133020033321012-1020330000130213-2022330131010123-0233233300333301-0211133103121022"></a>

<a id="canonical-1323232333300010-0333213013012300-0301302001022103-1003121023332120-0011011023201330-3221230000313112-0110032202300202-2330203313210312"></a>

## dns_address property — dhcp_networks / 000022332303 / 5

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
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

- [first_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0013330021001122-3130130213123100-0100210112200213-1123310301110002-0202331111322230-3213313200301320-2230102202312102-0122200000112303): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-005.md#canonical-0212132301200321-1023021333033300-2220131210220133-0023312212200021-3302122202001011-2331021321310302-2300222233002131-0322201023100223): complete subsection reference.

<a id="canonical-0222321023001312-0300203200332323-2111011120033301-0321221132330220-3302311112200311-0030200011030211-2122133313300002-2122322001131331"></a>

<a id="canonical-0203310012330201-3130310111211100-3030102121323321-3303033111202310-0202002312113101-3312111202211220-0231301300100330-1032123001011110"></a>

## network_prefix property — dhcp_networks / 000022332303 / 6

Type: `"string"`. Optional.

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

Upstream description:

Exclusive with \[\] Set the network prefix for the site. Ex: 192.0.2.0/24.

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

<a id="canonical-2200010010230103-1330021001030003-2130232302303221-2020032123121023-3210102112003321-1000322210123321-0332010002121023-0203011011300120"></a>

<a id="canonical-2320101202332110-2331031021032203-1101020312132121-3033231030012100-1131201321220320-2020013033232121-0332123021302011-0013100333011003"></a>

## pool_settings property — dhcp_networks / 000022332303 / 7

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Upstream description:

Identifies the how to pick the network for Interface.

Address ranges in DHCP pool list are used for IP Address allocation Address ranges in DHCP pool list
are excluded from IP Address allocation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
  "enum": [
    "INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS",
    "EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [pools](resources--securemesh_site_v2--reference--group-005.md#canonical-0210122333231303-0102210303200330-3221211103031100-2221122100301320-0320231133200130-0032221223133120-1321322321033330-3302012313202031): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-005.md#canonical-3100133112120033-1033112212330322-3212302013103101-1011210331123200-2210313030112322-2112002203311231-2221212222031322-2333131221033331): complete subsection reference.

<a id="canonical-1221230101102222-3022312131300131-2121330233023331-0121313123120330-2333233332333212-3223021321103001-0322022303222221-1330013110321010"></a>

## Next pages — dhcp_networks / 000022332303 / 8

- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-004.md#canonical-0013330021001122-3130130213123100-0100210112200213-1123310301110002-0202331111322230-3213313200301320-2230102202312102-0122200000112303)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-005.md#canonical-0212132301200321-1023021333033300-2220131210220133-0023312212200021-3302122202001011-2331021321310302-2300222233002131-0322201023100223)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-005.md#canonical-0210122333231303-0102210303200330-3221211103031100-2221122100301320-0320231133200130-0032221223133120-1321322321033330-3302012313202031)
- [azure.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-005.md#canonical-3100133112120033-1033112212330322-3212302013103101-1011210331123200-2210313030112322-2112002203311231-2221212222031322-2333131221033331)
- [azure.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-004.md#canonical-3333302011010332-2320201321011330-1301011022011133-3001200012101020-0220130102021202-2221021223212000-3111303300201232-1132200030312320)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0013330021001122-3130130213123100-0100210112200213-1123310301110002-0202331111322230-3213313200301320-2230102202312102-0122200000112303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
