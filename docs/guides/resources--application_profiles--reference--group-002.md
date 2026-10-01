---
page_title: "xcsh_application_profiles reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_application_profiles reference."
---

# xcsh_application_profiles reference

<a id="canonical-3010211220233111-1223213230310132-3221311032021320-1300002121232223-3213203331331111-0210320031003301-3100023013112313-1130313020233231"></a>

## Next pages — address_translation_enable / 321203012100 / 4

- [virtual_server.address_translation](resources--application_profiles--reference--group-001.md#canonical-3021112011013001-3231012202002003-2211233300032133-1313022100111002-1312313102121203-1032312311312031-2201100330021010-0102103013113132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311203311231310-1321211021211321-3001032300020100-0113301102203200-3213021320330301-0313023103022320-3132323001201313-1320201000002020"></a>

## virtual_server.auto_last_hop — auto_last_hop / 120120023323 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.auto_last_hop

<a id="canonical-2102003333002100-3120130233030201-2230023011033122-1122220031100031-0002012133223230-1220000021212230-2313302313111212-3212210233202022"></a>

Type: `"object"`. single nested block, Optional.

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if
the..

Upstream description:

When enabled, allows the system to send return traffic to the MAC address that transmitted the
request, even if the routing table points to a different network or interface. As a result, the
system can send return traffic to clients even when there is no matching route. For example, if the
system does not have a default route configured and the client is located on a remote network. This
setting is also useful when the system is load balancing transparent devices that do not modify the
source IP address of the packet. Without the last hop option enabled, the system could return
connections to a different transparent node, resulting in asymmetric routing. You can configure this
setting globally and on an object level. You set the global Auto Last Hop value on the System ::
Configuration :: Local Traffic :: General screen. To configure this setting globally, retain the
Default setting. When you configure Auto Last Hop with a value other than Default at the object
level, its setting takes precedence over the global setting. This enables you to configure auto last
hop on a per-virtual server basis. The default is Default, meaning that the system uses the global
auto-lasthop setting to send back the request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_last_hop_default",
    "auto_last_hop_disable"),
  validators.ConflictingObjectAttributes("auto_last_hop_default",
    "auto_last_hop_enable"),
  validators.ConflictingObjectAttributes("auto_last_hop_disable",
    "auto_last_hop_enable")}
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
  "x-ves-oneof-field-auto_last_hop_choice": "[\"auto_last_hop_default\",\"auto_last_hop_disable\",\"auto_last_hop_enable\"]"
}
```

Terraform syntax:

```terraform
auto_last_hop {
  # Configure direct properties listed below.
}
```

<a id="canonical-1220321000011213-1002323333231223-2031001331331132-0011113220121032-3003102331133220-2222333212302321-1103211100103231-0030301030223302"></a>

## Direct properties — auto_last_hop / 120120023323 / 3

- [auto_last_hop_default](resources--application_profiles--reference--group-002.md#canonical-3300231123130002-0223100030003331-3103032001323132-0123301233320131-2132201322230132-1322110221110023-3231131230133312-0213310310001103): complete subsection reference.

- [auto_last_hop_disable](resources--application_profiles--reference--group-002.md#canonical-1211230133221202-1200300331133130-1322013030101131-2100203101000011-0210311323233110-2031310113033322-0121010300020232-1210322123313203): complete subsection reference.

- [auto_last_hop_enable](resources--application_profiles--reference--group-002.md#canonical-0031223321131311-3301230011013221-3321012303232010-2123203031313020-1201001323030230-3020232230122012-1100313210121023-3011132112200321): complete subsection reference.

<a id="canonical-2300112333113211-2302323023133212-0020203222313100-1102220001321000-1111323020030320-2000033133302311-0001021310020331-3333230112132321"></a>

## Next pages — auto_last_hop / 120120023323 / 4

- [virtual_server.auto_last_hop.auto_last_hop_default](resources--application_profiles--reference--group-002.md#canonical-3300231123130002-0223100030003331-3103032001323132-0123301233320131-2132201322230132-1322110221110023-3231131230133312-0213310310001103)
- [virtual_server.auto_last_hop.auto_last_hop_disable](resources--application_profiles--reference--group-002.md#canonical-1211230133221202-1200300331133130-1322013030101131-2100203101000011-0210311323233110-2031310113033322-0121010300020232-1210322123313203)
- [virtual_server.auto_last_hop.auto_last_hop_enable](resources--application_profiles--reference--group-002.md#canonical-0031223321131311-3301230011013221-3321012303232010-2123203031313020-1201001323030230-3020232230122012-1100313210121023-3011132112200321)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3300231123130002-0223100030003331-3103032001323132-0123301233320131-2132201322230132-1322110221110023-3231131230133312-0213310310001103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031011302233000-0202122213233101-0020123011022221-2332311031310322-1031113001313003-2202301003300002-3211323330021110-1211320332230000"></a>

## virtual_server.auto_last_hop.auto_last_hop_default — auto_last_hop_default / 112011201130 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- virtual_server.auto_last_hop.auto_last_hop_default

<a id="canonical-0200312103221313-1103111200001032-1221002220011310-0112113223110233-3233322010331203-3230013203122122-1130012233102331-2313231022313131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for auto last hop default.

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
auto_last_hop_default = {}
```

<a id="canonical-0201123022123012-1122023033133112-2122232011300030-0232311022222232-1220002013101233-1000130233030201-0301313323003223-3223000201012000"></a>

## Direct properties — auto_last_hop_default / 112011201130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131030322020101-0333022230122313-0120213001033320-1112323022000202-1122001203203003-0322120303132121-1233333102232023-0112111310020032"></a>

## Next pages — auto_last_hop_default / 112011201130 / 4

- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1211230133221202-1200300331133130-1322013030101131-2100203101000011-0210311323233110-2031310113033322-0121010300020232-1210322123313203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213100133313220-2230130023003022-2110213302031222-1301003203003200-2120033012000332-2332222312122123-3021332231003101-0000333233002301"></a>

## virtual_server.auto_last_hop.auto_last_hop_disable — auto_last_hop_disable / 312130221203 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- virtual_server.auto_last_hop.auto_last_hop_disable

<a id="canonical-3333323311300221-1032210323101131-3110013321303220-0201100332302123-0222310211320100-1300320120332311-1223000303202000-2331032230221011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for auto last hop disable.

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
auto_last_hop_disable = {}
```

<a id="canonical-3020020331112110-3312131021102130-1230131003101201-3220231322021022-0010112022121213-1231213232322200-3230312323112101-1323111121200212"></a>

## Direct properties — auto_last_hop_disable / 312130221203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3322202033201122-2021332021222210-1312002112110232-1321110211023321-2230303021131301-3111310232023232-0101231320122322-2121321000133110"></a>

## Next pages — auto_last_hop_disable / 312130221203 / 4

- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0031223321131311-3301230011013221-3321012303232010-2123203031313020-1201001323030230-3020232230122012-1100313210121023-3011132112200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023302230231200-1203310300330203-1303032310303322-2230202201013331-0231213103001103-0013332031330130-2312103011111123-1011221330122222"></a>

## virtual_server.auto_last_hop.auto_last_hop_enable — auto_last_hop_enable / 032111002101 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- virtual_server.auto_last_hop.auto_last_hop_enable

<a id="canonical-2123101121211332-1123033021312131-0220230032033320-1312120221300022-0313300221331012-3230112211311023-0332021203320110-3123130113132311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for auto last hop enable.

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
auto_last_hop_enable = {}
```

<a id="canonical-0131301330301121-0022310312320023-2023021210111301-2132331211103110-0223313020121312-3120310211230302-2333303033003001-0321332310310031"></a>

## Direct properties — auto_last_hop_enable / 032111002101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300120030132121-1133331020133232-3121213203323202-1201130310101002-1003133003103220-1113002032110331-2210003211132220-1000120013211222"></a>

## Next pages — auto_last_hop_enable / 032111002101 / 4

- [virtual_server.auto_last_hop](resources--application_profiles--reference--group-002.md#canonical-0121310003002130-3003332201223302-1200332002132003-2333323132231110-0132023021202330-1302303220222033-1033311330201233-3023023232102123)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1021302330213202-2303301231320102-2211323122010333-2021300332322031-3212031203232320-0311110303022111-1313022210003033-3300113110112122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2023110231131103-2311111011003322-2112020303133330-0033233030230320-3020331001003222-2222130232323332-0103210320223101-3013113130231100"></a>

## virtual_server.clone_pool_client — clone_pool_client / 311303202322 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.clone_pool_client

<a id="canonical-1033030230211202-1021333323012333-3323221312331011-3320122331112203-2220120321232231-0010221120323210-1310313301100022-2200300322112102"></a>

Type: `"object"`. list nested block, Optional.

Replicates client-side traffic (that is, prior to address translation) to a member of the specified
pool.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
clone_pool_client {
  # Configure direct properties listed below.
}
```

<a id="canonical-3331102130322211-0132112031003310-1312220231011012-0311132100122233-2032310013002002-2102210002221131-1311231012210010-0132213033033321"></a>

## Direct properties — clone_pool_client / 311303202322 / 3

<a id="canonical-3120313123311212-3000231030202030-1210031020310120-3333210210332233-3202133133011221-1030021010303212-2120312233210302-2023202020020100"></a>

<a id="canonical-2321323023222110-2131001222103000-2333121031131013-2323032221232122-0220123301313230-1030230221003220-0011113202321233-3112210001202211"></a>

## kind property — clone_pool_client / 311303202322 / 4

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

<a id="canonical-0103112213301230-3333211011322130-0231332103311020-1231211330030103-2010001223211303-1011311231010010-1231100100213312-0233122323023011"></a>

<a id="canonical-3301123010031231-0122100330021213-3122003000201320-1023023112121232-1223300221132321-3322211102333333-3232131302211102-0032311010323012"></a>

## name property — clone_pool_client / 311303202322 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0301313033003103-2131023200312212-1301112022013032-2211213122102220-1122022010322210-2031320113103333-3012330020320002-1103031303111023"></a>

<a id="canonical-3310200033123202-2220310331123211-1303231020023000-1201301212302003-1113212013213131-1212133132033111-3212332231220020-0220321231310123"></a>

## namespace property — clone_pool_client / 311303202322 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1131200111132033-1211033311321020-0022320230302121-2202023021000220-0103131210321232-2001310012032033-0311010322302310-3110220220220131"></a>

<a id="canonical-2211100322011231-0133011223201101-0232100023013011-1012313230111123-2231322331130010-1020111033332232-1201001313011120-2111102122330030"></a>

## tenant property — clone_pool_client / 311303202322 / 7

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

<a id="canonical-3123233321330320-0331033131032210-3232200313222301-2030302122301231-2203003022113133-2100331021232122-3022132133013223-0201330110103311"></a>

<a id="canonical-2330210203301330-3202233210331311-2211312333331233-0013111311200030-0322002303010022-1000110333013001-3002202323223311-0313001101311032"></a>

## uid property — clone_pool_client / 311303202322 / 8

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

<a id="canonical-0210320012032020-3211023110323133-3232100100101031-1133002130023122-0003113113000113-3000131112220031-2233233021323210-0232233313300002"></a>

## Next pages — clone_pool_client / 311303202322 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0110030313223221-0212311332031102-3220000320023132-0113100203203210-2300323131111132-1320212032201203-2100233233233312-1111220311120132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131330301013300-3213322103331303-0102002223303322-2021111113033020-0131330121111221-1212023332201302-2221132310310303-0321030111113122"></a>

## virtual_server.clone_pool_server — clone_pool_server / 232001000333 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.clone_pool_server

<a id="canonical-1013221131231303-0110323201210021-1301332311032223-1333023030001320-2100220021203031-0202002323323223-3032101233123232-2120000203203121"></a>

Type: `"object"`. list nested block, Optional.

Replicates server-side traffic (that is, prior to address translation) to a member of the specified
pool.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
clone_pool_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332120123230202-3012132303311200-3323131301203100-2110223321333102-0200001022311001-0312232021330321-1003331213320012-2210300331122301"></a>

## Direct properties — clone_pool_server / 232001000333 / 3

<a id="canonical-1223121310000131-0022301003011121-2221122103302113-0030111003021112-1230302210221101-0012300132033202-0002323331303012-2030330031211032"></a>

<a id="canonical-2222003100113102-2133032122121320-0130021331203013-1203012231221130-0101012020031211-0121302021020221-3333321220020233-2131322213121212"></a>

## kind property — clone_pool_server / 232001000333 / 4

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

<a id="canonical-2302010132332301-2211231221023230-1110110111231030-1302333311012223-0130311111330321-2322033223012313-0310332212302203-0320330101122133"></a>

<a id="canonical-0001133212102301-0003212011131221-1303133333301132-3030003010121213-2031310120323112-1113331320102012-1111122220233201-3012211112311323"></a>

## name property — clone_pool_server / 232001000333 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3030102120302331-0130002010320330-3332321233111002-1313330233130311-2122320233302211-3211030301202321-3021002013022323-3220202312102131"></a>

<a id="canonical-2331210113130311-2131003001223302-2012113113301012-3220002021332213-2103201233200130-1003011030023001-2201100333021211-2021120020313131"></a>

## namespace property — clone_pool_server / 232001000333 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-3113330030320212-2210233311203202-2321120111202211-0231111010022231-0210132132323001-0323002123010031-3211213020010312-0130123230233000"></a>

<a id="canonical-1100103212113112-0302012011023232-0201010010231010-3001011013002322-3320112323032003-1330202021110010-0323021222223210-0121303210233122"></a>

## tenant property — clone_pool_server / 232001000333 / 7

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

<a id="canonical-2020312020003200-3122200122222001-3131220100232322-2021232032133202-1003332112322122-2313112313203103-1332023213112111-1003312303323102"></a>

<a id="canonical-2033003222023203-0212012010333223-1010132311100230-0211110010021323-3112320013330323-1132021033130002-2033021203121032-2223021123001112"></a>

## uid property — clone_pool_server / 232001000333 / 8

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

<a id="canonical-1032211332230031-2320223003123002-2033221132112211-0102111031201010-0302302212321200-3113121022101203-2133233300312220-1232011112201032"></a>

## Next pages — clone_pool_server / 232001000333 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012033202323021-1022322102201000-2212332322100103-0201333313111121-1122003311231001-3132012311120222-3323322322032102-3200101300320330"></a>

## virtual_server.connection_rate_limit_mode — connection_rate_limit_mode / 020332210123 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.connection_rate_limit_mode

<a id="canonical-3301023333300032-2210013303130033-3332030022313210-1311132001212223-3022013121111032-0221303131332122-2021123212320013-0221332020333311"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for connection rate limit mode.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("per_destination_address",
    "per_source_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_source_destination_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_destination_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_source_destination_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_source_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_source_destination_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server",
    "per_virtual_server_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_virtual_server",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server_destination_address",
    "per_virtual_server_source_address"),
  validators.ConflictingObjectAttributes("per_virtual_server_destination_address",
    "per_virtual_server_source_destination_address"),
  validators.ConflictingObjectAttributes("per_virtual_server_source_address",
    "per_virtual_server_source_destination_address")}
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
  "x-ves-oneof-field-connection_rate_limit_mode_choice": "[\"per_destination_address\",\"per_source_address\",\"per_source_destination_address\",\"per_virtual_server\",\"per_virtual_server_destination_address\",\"per_virtual_server_source_address\",\"per_virtual_server_source_destination_address\"]"
}
```

Terraform syntax:

```terraform
connection_rate_limit_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230021133020233-1033232013323033-0010103311010220-1231301001212223-2221302330311132-2002210311113101-1332000132212233-1000321200002311"></a>

## Direct properties — connection_rate_limit_mode / 020332210123 / 3

- [per_destination_address](resources--application_profiles--reference--group-002.md#canonical-0311220212331211-3011202012101113-0232011333222301-2203212100122311-0131333201111200-3210010320312213-1012131321030100-2210101133000220): complete subsection reference.

- [per_source_address](resources--application_profiles--reference--group-002.md#canonical-3112311033110011-0231210000133330-2311033311023222-3312112003323111-0211331222300131-2100100002310130-3022112310120230-1310100121122011): complete subsection reference.

- [per_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0112133010331231-3221202200001131-0320310312102001-2223300122300312-2222302002211231-3033121312002302-0330102121121023-2331210003103123): complete subsection reference.

- [per_virtual_server](resources--application_profiles--reference--group-002.md#canonical-0002110030200031-3021230220301233-2301020300210212-2012003220332212-2101023222030132-2332012301011100-1321330200112001-1110103230220112): complete subsection reference.

- [per_virtual_server_destination_address](resources--application_profiles--reference--group-002.md#canonical-0301121332212300-0113023232303112-2333110211120123-2123112112111322-2120111311301310-3012013311220121-1203221022201203-1231323231200220): complete subsection reference.

- [per_virtual_server_source_address](resources--application_profiles--reference--group-002.md#canonical-2120210301030312-0201031221120200-1030032320011313-3212303321331232-1011211111003121-3001112002001300-0312113022222301-3110233212220201): complete subsection reference.

- [per_virtual_server_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0113212332311100-1310301131000332-1202003101033103-2322132100123101-0101000111312210-2220101121121130-3032010102130123-2011222003213122): complete subsection reference.

<a id="canonical-1000323210330311-1123300231002002-0311030330003132-2232200303322000-2221120223232230-1203013200330032-3000000210020032-2003131313301121"></a>

## Next pages — connection_rate_limit_mode / 020332210123 / 4

- [virtual_server.connection_rate_limit_mode.per_destination_address](resources--application_profiles--reference--group-002.md#canonical-0311220212331211-3011202012101113-0232011333222301-2203212100122311-0131333201111200-3210010320312213-1012131321030100-2210101133000220)
- [virtual_server.connection_rate_limit_mode.per_source_address](resources--application_profiles--reference--group-002.md#canonical-3112311033110011-0231210000133330-2311033311023222-3312112003323111-0211331222300131-2100100002310130-3022112310120230-1310100121122011)
- [virtual_server.connection_rate_limit_mode.per_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0112133010331231-3221202200001131-0320310312102001-2223300122300312-2222302002211231-3033121312002302-0330102121121023-2331210003103123)
- [virtual_server.connection_rate_limit_mode.per_virtual_server](resources--application_profiles--reference--group-002.md#canonical-0002110030200031-3021230220301233-2301020300210212-2012003220332212-2101023222030132-2332012301011100-1321330200112001-1110103230220112)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address](resources--application_profiles--reference--group-002.md#canonical-0301121332212300-0113023232303112-2333110211120123-2123112112111322-2120111311301310-3012013311220121-1203221022201203-1231323231200220)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_address](resources--application_profiles--reference--group-002.md#canonical-2120210301030312-0201031221120200-1030032320011313-3212303321331232-1011211111003121-3001112002001300-0312113022222301-3110233212220201)
- [virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address](resources--application_profiles--reference--group-002.md#canonical-0113212332311100-1310301131000332-1202003101033103-2322132100123101-0101000111312210-2220101121121130-3032010102130123-2011222003213122)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0311220212331211-3011202012101113-0232011333222301-2203212100122311-0131333201111200-3210010320312213-1012131321030100-2210101133000220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110021330133333-0211331320320212-1310113030101011-0322200330011313-0132013210230301-3122213130322132-1110101201023312-0302210123301322"></a>

## virtual_server.connection_rate_limit_mode.per_destination_address — per_destination_address / 231003221000 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_destination_address

<a id="canonical-3313232302320002-1322131102311312-3202132010023031-2321130021120013-2003133200130013-2122320003212321-3020011300103011-3302222023132003"></a>

Type: `"object"`. single nested block, Optional.

Destination Address Mask.

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
per_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1012323321132331-3103303301232130-3223021001130132-0210103230302001-0111211333030022-3330332211133123-3032020223231001-1100320122120201"></a>

## Direct properties — per_destination_address / 231003221000 / 3

<a id="canonical-1011231323321110-0220010012111331-0001312032130211-2230012002022120-2203231322201101-2003032111331110-0222310201000131-1002031220203333"></a>

<a id="canonical-2300202223300312-0022030333010201-0100320133200331-2230202100121120-3100031223322202-1303310031013330-1013201013111032-2012111021323022"></a>

## destination_mask property — per_destination_address / 231003221000 / 4

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1002102311301102-1112100202223330-0200130120300112-3112002001001310-0113221220111230-3011321100013231-2000322023322322-1123132111231201"></a>

## Next pages — per_destination_address / 231003221000 / 5

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3112311033110011-0231210000133330-2311033311023222-3312112003323111-0211331222300131-2100100002310130-3022112310120230-1310100121122011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202301133000223-0200131200332310-1300000322230233-3221333023211203-2032000312312312-0033133203330330-2213210130300003-1110012302032012"></a>

## virtual_server.connection_rate_limit_mode.per_source_address — per_source_address / 212301012101 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_source_address

<a id="canonical-1032313202203331-2213320301221002-2211210101321220-3133123230322220-1301210201332131-0321031301121220-2332221203212311-1323130120131323"></a>

Type: `"object"`. single nested block, Optional.

Source Address Mask.

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
per_source_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203111002013030-2333102210320022-2233101113110131-3230121303113201-0332122211200000-2312230213111300-1321331201233103-1200230121131300"></a>

## Direct properties — per_source_address / 212301012101 / 3

<a id="canonical-3030132331223231-2311121121311021-0023330023232220-2130321003300202-0321323130233121-2333301200322333-2132323112102003-1021023322120103"></a>

<a id="canonical-0113213300003011-1300231011100132-1000012222002012-3232211232012331-1310331211210310-1333212011030300-1111310332333031-1232210330213003"></a>

## source_mask property — per_source_address / 212301012101 / 4

Type: `"number"`. Optional.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1221223301220002-2120021023323103-1130211122320013-3132020203213211-2001301112010310-1332223120033131-3320011121100322-1133122022100202"></a>

## Next pages — per_source_address / 212301012101 / 5

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0112133010331231-3221202200001131-0320310312102001-2223300122300312-2222302002211231-3033121312002302-0330102121121023-2331210003103123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023212201103031-3222302020022033-3121203130011010-0111232300300022-1300120203220130-2311222221320333-1113301100012121-1220321301110133"></a>

## virtual_server.connection_rate_limit_mode.per_source_destination_address — per_source_destination_address / 110102201230 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_source_destination_address

<a id="canonical-1121321010302331-3210012021133201-1311000332111023-0102212333311320-1302133003221333-3332301131221312-1322321301203313-2200323113202003"></a>

Type: `"object"`. single nested block, Optional.

Destination and Source Address Mask.

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
per_source_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112333113010202-1300121321230031-0102001110303210-2121321301122010-2002133213320220-1131033003121113-0332231222230203-0121110022121303"></a>

## Direct properties — per_source_destination_address / 110102201230 / 3

<a id="canonical-3130120312101122-1301101010333203-1321211202203331-3233022013120312-1312002132212303-3103202112111212-2310220211133201-1222332310111100"></a>

<a id="canonical-2102001212330102-0022322312122322-2030121211323031-0221121201020321-1310312201130220-2232311113302133-2012111120131231-3013211021330111"></a>

## destination_mask property — per_source_destination_address / 110102201230 / 4

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3131112320301313-1101232211113011-3001121131001302-0101312032323111-1000000131131032-2202312000332200-1131023233110021-3031121111323311"></a>

<a id="canonical-2321200123200122-0230021332203021-1023112122212121-1033000020022321-0320331302032003-0110331101311030-1231022311311321-3012033230220321"></a>

## source_mask property — per_source_destination_address / 110102201230 / 5

Type: `"number"`. Optional.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-3302103012132030-2221103321111320-1013022220322000-0330223211212121-0012301323131001-3112212000110222-0112110122013102-3203120323100122"></a>

## Next pages — per_source_destination_address / 110102201230 / 6

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0002110030200031-3021230220301233-2301020300210212-2012003220332212-2101023222030132-2332012301011100-1321330200112001-1110103230220112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020310112001022-3220313121031130-0303201310200221-2323131021101022-2302131301333203-3101122012021101-0101033222102122-1213200032120203"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server — per_virtual_server / 210203112100 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server

<a id="canonical-3130113030210232-3312223213230030-2322321202130103-1302122101233331-0120310132320003-1010003012102203-1003023002131030-1022333022323012"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for per virtual server.

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
per_virtual_server = {}
```

<a id="canonical-0003013310200232-0310233200121102-1023030201122202-0201303211131021-0102220130222131-3022221010030121-3233210033112232-1030300303020013"></a>

## Direct properties — per_virtual_server / 210203112100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3011133223120012-2313030222112020-1012333301010312-3211321010120230-1222131131333323-0102323101211320-1330131330112030-3101233102220202"></a>

## Next pages — per_virtual_server / 210203112100 / 4

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0301121332212300-0113023232303112-2333110211120123-2123112112111322-2120111311301310-3012013311220121-1203221022201203-1231323231200220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310313320001012-1202300300320330-1110111012130333-1113011002010120-1112001012300300-1012221313102223-2312233331022313-1103312100101300"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address — per_virtual_server_destination_address / 302001213322 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server_destination_address

<a id="canonical-2012022113300113-2301301113011120-2020123111303122-1010131133200020-1212100222003022-2110131011230123-2303210200233110-1300223013230020"></a>

Type: `"object"`. single nested block, Optional.

Destination Address Mask.

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
per_virtual_server_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321222312300311-3022212001223120-2331110202133131-2130211013300032-2230303012233111-2210311320221220-1330130003113000-0133202131303110"></a>

## Direct properties — per_virtual_server_destination_address / 302001213322 / 3

<a id="canonical-1000011102212103-3113131323031300-1033321133312130-3011111121312303-0032332330103213-2130301222003102-0232130310131322-0012233033333233"></a>

<a id="canonical-0021010230000031-1010021121010202-3300123212231232-3333312122222302-1221011320010313-1321232321101310-1322022013103330-1013203012033132"></a>

## destination_mask property — per_virtual_server_destination_address / 302001213322 / 4

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-0301232133120123-2132323312131332-2002111033310121-0323233202301010-0112002012121213-1011023212312100-3000002323002221-0113011132332013"></a>

## Next pages — per_virtual_server_destination_address / 302001213322 / 5

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2120210301030312-0201031221120200-1030032320011313-3212303321331232-1011211111003121-3001112002001300-0312113022222301-3110233212220201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321320330123021-3321103010301133-3123131132031122-2313310232233320-3122021011110203-1021112233312311-1111302333231312-2301010200003312"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server_source_address — per_virtual_server_source_address / 310332311221 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_address

<a id="canonical-2101303130210111-1220102123110101-0231331023010311-1203131013010310-3023132221000003-3032200132201300-3222303021122200-0133330120331001"></a>

Type: `"object"`. single nested block, Optional.

Source Address Mask.

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
per_virtual_server_source_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-0020123112000133-2323101210223020-2210202313023011-0331011323113211-3003321322022000-2121312322301201-0201223113223031-1300301310022203"></a>

## Direct properties — per_virtual_server_source_address / 310332311221 / 3

<a id="canonical-0333231212030220-1021031331111010-0321211202222011-3221220033112032-2233122210303320-0132212011202312-3223331303230003-0201112110123333"></a>

<a id="canonical-2303121121322130-0120100201323331-3331133200232331-1123030233121233-2201331300020222-2303022030233022-2321231030330001-3201322002321300"></a>

## source_mask property — per_virtual_server_source_address / 310332311221 / 4

Type: `"number"`. Optional.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2030210012331232-3333011202032022-0220301202103230-3313132330033331-1011012222202120-2103010311030222-2023211231213013-3020011103002203"></a>

## Next pages — per_virtual_server_source_address / 310332311221 / 5

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0113212332311100-1310301131000332-1202003101033103-2322132100123101-0101000111312210-2220101121121130-3032010102130123-2011222003213122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212002201223103-3210210032001132-2131000202202333-0200013101122212-2233123221211101-3020222133030133-1132132101110201-2010030331303013"></a>

## virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address — per_virtual_server_source_destination_address / 030121210301 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- virtual_server.connection_rate_limit_mode.per_virtual_server_source_destination_address

<a id="canonical-0331310010022310-2302213032223221-1223232003201100-3021010332102311-2211130113110022-0223001011223003-3212310223321212-0023232130313231"></a>

Type: `"object"`. single nested block, Optional.

Destination and Source Address Mask.

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
per_virtual_server_source_destination_address {
  # Configure direct properties listed below.
}
```

<a id="canonical-1000223030213030-3332032121012112-2001110301222322-1130211321111001-2333303311323203-2032232311221332-3121123221233200-1223133332121300"></a>

## Direct properties — per_virtual_server_source_destination_address / 030121210301 / 3

<a id="canonical-1131102111332120-3131131320131203-3203120021321221-0233121222103203-2213311330001211-2100023100332312-3133011120012220-3321112230102313"></a>

<a id="canonical-3023123103021002-2302303202200313-3302023012020121-1010132221233011-2130232301203002-1103303201330121-2201203132133112-2233123131011002"></a>

## destination_mask property — per_virtual_server_source_destination_address / 030121210301 / 4

Type: `"number"`. Optional.

Configuration parameter for destination mask.

Upstream description:

Configuration parameter for destination mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-1231033213102212-0020130033122133-0330222102122121-2211323131230012-0200201101111020-1020212310313223-3322201022120002-1000312130312113"></a>

<a id="canonical-2000320110110000-0123121010313132-3113330021103223-0222002132212033-3330333312120212-2020103231223130-2110132203210310-2213122120203113"></a>

## source_mask property — per_virtual_server_source_destination_address / 030121210301 / 5

Type: `"number"`. Optional.

Configuration parameter for source mask.

Upstream description:

Configuration parameter for source mask

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 600000),
}
```

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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

<a id="canonical-2000203320122022-2203031330203002-1113131103321330-2303102211113113-3301203230003333-3101231101210100-1213313301212123-2102032201030013"></a>

## Next pages — per_virtual_server_source_destination_address / 030121210301 / 6

- [virtual_server.connection_rate_limit_mode](resources--application_profiles--reference--group-002.md#canonical-1210022010230202-1323323121231021-0100301113232321-2302233310331221-1202213321330303-3233312320100003-0332120123013031-1311100000320202)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1130200120303132-1310331020320222-2223332303331002-1001311303211331-1130223022103331-3332031011100321-0310011023011303-0330123013203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100133310032103-3002120311011123-0002020122020211-2033312300133121-0203030212122322-1131131211330303-2331021203221003-2221220332010030"></a>

## virtual_server.default_persistence_profile — default_persistence_profile / 330012120122 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.default_persistence_profile

<a id="canonical-3003330212123022-1202023302201020-2222131130302220-3103112023103232-0120013021311210-1301120010220333-0031231203030233-2001000121103120"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for default persistence profile.

Upstream description:

Configuration parameter for default persistence profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_persistence_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212330122330123-1110120203033103-2310012300231333-1121100203222233-1320122100022002-2132023020202111-3032013123000003-0121123130011323"></a>

## Direct properties — default_persistence_profile / 330012120122 / 3

<a id="canonical-1032030113301233-3012310320011110-2321311321123130-1010033223211110-3002330031100210-1033101231011032-0021331301222213-1312000100200102"></a>

<a id="canonical-2312113011003231-3223202220300332-0133130321302233-1033113322320313-3300331232032020-1113121103020011-3332311023012221-2133130010302132"></a>

## kind property — default_persistence_profile / 330012120122 / 4

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

<a id="canonical-1013302000100300-0300223331100030-2301030003001313-3211212333200301-0220220033023021-2132120331210312-3201021221330201-0231103211111301"></a>

<a id="canonical-1203110303100311-3130000133033233-2230301101113103-0330031133123013-2120122203023223-0002330023101300-1210312123203211-2333000020022111"></a>

## name property — default_persistence_profile / 330012120122 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0123110133310023-2330220232011220-2010113312321210-1213112030223011-2031102300213301-3331013321030323-0212110123021301-2022130110313013"></a>

<a id="canonical-0230232123210320-1130223002021103-0312310002210220-0202303012200233-2032302112300000-0011302000313320-2100330003003321-2213020021213121"></a>

## namespace property — default_persistence_profile / 330012120122 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-2132010111023003-1310103300310111-0023222132330330-3203010001133001-1112031101200301-3312011210302002-1201102312131100-2131122132111300"></a>

<a id="canonical-0323201113013330-3120300310020011-3300210220121011-1100213230210302-0010121331002001-0320130232202020-0032330313120323-3013023102210323"></a>

## tenant property — default_persistence_profile / 330012120122 / 7

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

<a id="canonical-2302120021022132-1013123122013002-2213233113310232-3033120213200102-1032003123003033-2103113033103000-3223303121211332-1320121123222331"></a>

<a id="canonical-0230120121030303-2020002202031020-2102213132301121-0002201230231003-3322032021211222-3310012111320301-0101201301310330-3123022201233323"></a>

## uid property — default_persistence_profile / 330012120122 / 8

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

<a id="canonical-2013103203302301-0322311020021222-3020321012233033-2120033332031033-0231332333012103-2133110300320220-0123001212211213-3203313110122200"></a>

## Next pages — default_persistence_profile / 330012120122 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0220212320310233-0020030230010211-1101200002230023-3233033023031122-2320003100323101-1300230000023000-1322202320211120-0101202210330311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2021010310231133-2133233201030321-2200003210030333-3010013202331222-2122133200332213-1203001220203020-3131100330331123-0213333330332322"></a>

## virtual_server.default_pool — default_pool / 002030110000 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.default_pool

<a id="canonical-3323300103233100-2200233303311301-1122232232202202-1300201201233111-3320200010030210-3001213020011320-2330222130011223-0013330220133002"></a>

Type: `"object"`. list nested block, Optional.

Specifies the pool name that you want the virtual server to use as the default pool. A load
balancing virtual server sends traffic to this pool automatically, unless an iRule directs the
server to send the traffic to another pool instead.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
default_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020210033012120-0021121310033331-1213003210012233-2202312313001320-0000003222023113-0000330133110100-0312031322323332-3111322112020302"></a>

## Direct properties — default_pool / 002030110000 / 3

<a id="canonical-3233011302111323-3010331222123322-2113221113201212-0300101020323000-3113313031321130-0131030001233023-0213100200220203-2230003310200201"></a>

<a id="canonical-0320021230212222-1031301031100212-1121223203110120-0231223320002231-3010230300011222-1113031011322000-2022312100212321-1320021200321203"></a>

## kind property — default_pool / 002030110000 / 4

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

<a id="canonical-1002132220121110-1130032020003033-2002033121132232-2120023033003231-0022110202310230-3332022011103131-0213321001113130-0200221001031032"></a>

<a id="canonical-2211213010201311-3023033113110101-2033230220001321-0020333031332222-1200310303121000-0313132330213022-2021311203013322-0220313121332013"></a>

## name property — default_pool / 002030110000 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0110023303023330-3310132013030213-2011022132022302-3230220213233313-1301133323333101-2033203332112300-3333200010023021-0213013001211022"></a>

<a id="canonical-3133231232203023-2122200030213011-1200120032203203-1323300313311330-1233222133010300-2333303201001312-0230331322110302-0123102300022113"></a>

## namespace property — default_pool / 002030110000 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-3302123123312220-1320322313001213-0101323301020221-0202312333130231-1110112320010232-1200110201300211-1012213220310203-0033010123332220"></a>

<a id="canonical-2023221231333233-2323323220102011-0303332311112113-2200210301133003-3310332010202202-3010012121313111-1021333100302220-1313320231320003"></a>

## tenant property — default_pool / 002030110000 / 7

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

<a id="canonical-0012023333200002-2030020322213323-3012222013333221-1221232000013020-3222212301313023-1313110223331002-1302313223301110-3332300221000103"></a>

<a id="canonical-2003320221210320-0221300010211321-1031333321312330-0301001133102001-2011003230132132-3110200112100122-1212200133203232-2022331021321321"></a>

## uid property — default_pool / 002030110000 / 8

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

<a id="canonical-1033323221033320-2330232202120221-1232111122111110-3101331230232333-0103130021332300-1213203132301332-0201210312131000-2033233230321223"></a>

## Next pages — default_pool / 002030110000 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3103321022122021-1320313333200212-3310120231330011-3010133213120203-2222103302302221-3011000212110223-0013202223013102-1230032013233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2230230211102201-2311211002302123-1103033213133123-3011203332130023-0101100233323230-0101131203213000-0211232033211132-3022020220301131"></a>

## virtual_server.fallback_persistence_profile — fallback_persistence_profile / 001111221012 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.fallback_persistence_profile

<a id="canonical-3021300023002001-2000330111223011-0331013022010230-3330313113132111-3323032232021003-3122301013122023-2212102223011130-3300130111221331"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for fallback persistence profile.

Upstream description:

Configuration parameter for fallback persistence profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
fallback_persistence_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3102310020023133-3311130303121332-2200230103200020-3010001003031131-1221222030312230-2222000201310210-1321202011223002-2203331002331010"></a>

## Direct properties — fallback_persistence_profile / 001111221012 / 3

<a id="canonical-1330112230020101-0021310312221331-0103121112122020-0210321112111223-3103001323110320-2032210302321233-0011320310231023-1030033230022203"></a>

<a id="canonical-2303133321320230-0103132221130112-0211030321321132-2200101021032000-2312302021333002-2332103013211013-0230203202233210-3032230010020311"></a>

## kind property — fallback_persistence_profile / 001111221012 / 4

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

<a id="canonical-2030221001130301-3122031011310203-0303101300031310-0000132113321210-1223121302130113-2012233200300030-2301032112033000-0112101012230311"></a>

<a id="canonical-0220300320031111-2202111030013013-2302312300330010-3211331321023210-2332213020301001-1001012302212323-1120002323033022-3103332012222322"></a>

## name property — fallback_persistence_profile / 001111221012 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1123323201121001-2013133301202012-0113102110322300-3331131322120030-3231222132013022-2323000332133031-1301002101110230-3121330230102010"></a>

<a id="canonical-0103200103231132-0113000211113103-2202131310313301-2212132131301203-3322022320303131-0301002223301100-1231323133021313-2003222312201130"></a>

## namespace property — fallback_persistence_profile / 001111221012 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1333323303123233-2212133033033032-2120232102123002-1220002122200002-1123310211023213-1031113002012221-0120113120223010-1111003031133230"></a>

<a id="canonical-3323330011321030-0101000012200220-3112231020213022-1000211031312300-3122310013133302-2103220211112231-0213320311101300-1203202201132321"></a>

## tenant property — fallback_persistence_profile / 001111221012 / 7

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

<a id="canonical-3301210113012320-2212320330000330-1011212000223311-1302222001220112-2032220300132021-0121012331200331-0131023223302302-3332022203003302"></a>

<a id="canonical-1210211123010020-0330133213030231-3021133113011331-1331213120013312-1131120021323133-1130103221313311-1022202023113112-1323122312033120"></a>

## uid property — fallback_persistence_profile / 001111221012 / 8

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

<a id="canonical-1221130112031223-2120003323302101-2313011311310122-2312330102300333-1133110123100101-0221323101123303-2332021212102000-1130320111233212"></a>

## Next pages — fallback_persistence_profile / 001111221012 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0300032301211222-3312312313011013-2022223230130323-3110133300220133-1121320101021121-2303320221132211-1112303221011231-2020012203010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113302322233233-1133231112021322-0022100030323123-2001103102111331-3213001102312102-0102110112100300-2033110023221123-3130132113201332"></a>

## virtual_server.fix_profile — fix_profile / 221100133332 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.fix_profile

<a id="canonical-2220330223223032-2003200010320322-3132202123100121-2211120320000322-1013120130311031-3101111201103333-3311130121031320-2122213302333231"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for fix profile.

Upstream description:

Configuration parameter for fix profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
fix_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313202233303020-3313320232111201-1211330213211110-0102123312133202-2001213232202300-0120021000222233-0113332021322300-1010011313332032"></a>

## Direct properties — fix_profile / 221100133332 / 3

<a id="canonical-2332022332123013-0013231020120012-1231010201023101-2132002021110223-0001312112002302-1330333131031221-1210320030303330-2122211033200233"></a>

<a id="canonical-0203111303120132-0003000130130300-3210212323222311-2031212023202321-0303302103213320-3222331121312202-2112302112020003-1001013101001032"></a>

## kind property — fix_profile / 221100133332 / 4

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

<a id="canonical-2233022010323031-2220201322202013-1002120001010322-0010201330310101-0200120100333021-1333030301222310-3033101302323022-0321333123322100"></a>

<a id="canonical-1303333202122330-1322002002013100-3030220331003323-0100211202313202-0012220222311013-3030112132312312-1310332001113110-0313321222230100"></a>

## name property — fix_profile / 221100133332 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0203032130000311-2321212301033301-2331313011220000-1003032130330230-3330310231333222-3121233002203331-2231213300212122-3213330112031323"></a>

<a id="canonical-2100212233212003-2232020313202031-0210132033130332-2322223321103130-2320333320012112-1232010112102300-1002312112213201-3210022123213223"></a>

## namespace property — fix_profile / 221100133332 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-2121121222100033-1120220212111320-2310301202130120-1222201030031022-0130123120131231-2101223003221201-0301320022220202-3323110232102322"></a>

<a id="canonical-0231111230303010-1313102133112010-3011020130132221-1323221032221013-3203213321100331-2023332313221200-1133223120112221-3203011320023330"></a>

## tenant property — fix_profile / 221100133332 / 7

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

<a id="canonical-1202223323102200-0012100002112132-0012200300132221-1313220330133232-2130023120233002-2230212102100031-3023333312023321-0022200132010232"></a>

<a id="canonical-1231101213033332-1332132100213003-3313311330203302-0033301112233023-2101231132301011-3213112030100322-0230111012000103-0330333331301102"></a>

## uid property — fix_profile / 221100133332 / 8

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

<a id="canonical-1230232300211121-0220030222331010-1302202313002312-3213332011310231-0103322233313320-2031110100330122-2212110203223212-1112032103022020"></a>

## Next pages — fix_profile / 221100133332 / 9

- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313220031113031-1110133233301202-3101102331231310-2032233212100122-3013033231132030-0102131013223123-0303201000331212-1313030130103201"></a>

## virtual_server.http — http / 313030300011 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.http

<a id="canonical-2332010000332020-0110331212332300-3330311200330200-3222000010231313-1221311303330220-0313132223022231-3301130303131001-0033103233213202"></a>

Type: `"object"`. single nested block, Optional.

HTTP profiles.

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
http {
  # Configure direct properties listed below.
}
```

<a id="canonical-0222211021313121-1333333212032233-1100313201330022-1331322321000010-0013230121030310-1220023231201021-0301313321333020-2311333002311233"></a>

## Direct properties — http / 313030300011 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3130120121301111-0232122113022331-0133313310022200-2330232310032203-1332022200120302-1321130233030333-0131333010323133-2230331203221003): complete subsection reference.

- [http2_client_profile](resources--application_profiles--reference--group-002.md#canonical-0332330302121333-0200202323301323-2011232011103020-2101231222032323-2011023122121013-2201030101010301-0301300001200132-0033123212100321): complete subsection reference.

- [http2_server_profile](resources--application_profiles--reference--group-002.md#canonical-0022331112021010-3110203310013002-0212322031113123-3313230120011121-0103101002003131-0033200323322031-0033010030130201-3323002202213230): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-002.md#canonical-2122303303313200-2013130123113121-1031102010022023-1131131323001322-2322123310020201-0021023332013321-2103212212003311-1320211331131022): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-002.md#canonical-2322322120023310-2331230320202223-2321331132010123-1120303200320012-2132131011333223-0311000303222001-0111212033133333-2113032201233303): complete subsection reference.

- [ocsp_profile](resources--application_profiles--reference--group-002.md#canonical-0212200100013300-2010330003202013-2113300023302120-3201222310221023-0313220131110103-1312321001211122-3021120111312003-3200020321033213): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3232011302112010-0200312011113331-3101111203223133-3020113313300300-0023113130010221-0133232012023323-1002302300221201-1331300202001202): complete subsection reference.

- [stream_profile](resources--application_profiles--reference--group-002.md#canonical-2003122102331011-0102223311211013-1113200013103313-3101102121131201-2101310003131021-3232021132033212-0313132302121023-2222132212203322): complete subsection reference.

- [tcp_client_profile](resources--application_profiles--reference--group-002.md#canonical-2223311101321210-2223313131302002-2310010033333222-2122013110101211-1021221123210103-1210221022002331-2101300030310020-2331210030200231): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-002.md#canonical-3313330121201111-0233011211100030-0312301221013133-0221322322010322-0221231321001010-2010112300003100-1220230110101232-1323330130100112): complete subsection reference.

- [websocket_client_profile](resources--application_profiles--reference--group-002.md#canonical-2103003320213123-2333013120322010-2312232213313300-1031221031100011-3332121220121033-0110233202012310-0011130201222110-1123332212132030): complete subsection reference.

- [websocket_server_profile](resources--application_profiles--reference--group-002.md#canonical-2133310101022000-1330312111332132-3210102331201100-1113232033330012-0200001232322231-1010212032131322-0301022332301131-3030020121130200): complete subsection reference.

<a id="canonical-1221023303130131-2321111320011000-0133322030322230-0201331112031303-3232221112232020-1233102111012121-2120103013020002-2233010310103133"></a>

## Next pages — http / 313030300011 / 4

- [virtual_server.http.client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3130120121301111-0232122113022331-0133313310022200-2330232310032203-1332022200120302-1321130233030333-0131333010323133-2230331203221003)
- [virtual_server.http.http2_client_profile](resources--application_profiles--reference--group-002.md#canonical-0332330302121333-0200202323301323-2011232011103020-2101231222032323-2011023122121013-2201030101010301-0301300001200132-0033123212100321)
- [virtual_server.http.http2_server_profile](resources--application_profiles--reference--group-002.md#canonical-0022331112021010-3110203310013002-0212322031113123-3313230120011121-0103101002003131-0033200323322031-0033010030130201-3323002202213230)
- [virtual_server.http.http_client_profile](resources--application_profiles--reference--group-002.md#canonical-2122303303313200-2013130123113121-1031102010022023-1131131323001322-2322123310020201-0021023332013321-2103212212003311-1320211331131022)
- [virtual_server.http.http_server_profile](resources--application_profiles--reference--group-002.md#canonical-2322322120023310-2331230320202223-2321331132010123-1120303200320012-2132131011333223-0311000303222001-0111212033133333-2113032201233303)
- [virtual_server.http.ocsp_profile](resources--application_profiles--reference--group-002.md#canonical-0212200100013300-2010330003202013-2113300023302120-3201222310221023-0313220131110103-1312321001211122-3021120111312003-3200020321033213)
- [virtual_server.http.server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-3232011302112010-0200312011113331-3101111203223133-3020113313300300-0023113130010221-0133232012023323-1002302300221201-1331300202001202)
- [virtual_server.http.stream_profile](resources--application_profiles--reference--group-002.md#canonical-2003122102331011-0102223311211013-1113200013103313-3101102121131201-2101310003131021-3232021132033212-0313132302121023-2222132212203322)
- [virtual_server.http.tcp_client_profile](resources--application_profiles--reference--group-002.md#canonical-2223311101321210-2223313131302002-2310010033333222-2122013110101211-1021221123210103-1210221022002331-2101300030310020-2331210030200231)
- [virtual_server.http.tcp_server_profile](resources--application_profiles--reference--group-002.md#canonical-3313330121201111-0233011211100030-0312301221013133-0221322322010322-0221231321001010-2010112300003100-1220230110101232-1323330130100112)
- [virtual_server.http.websocket_client_profile](resources--application_profiles--reference--group-002.md#canonical-2103003320213123-2333013120322010-2312232213313300-1031221031100011-3332121220121033-0110233202012310-0011130201222110-1123332212132030)
- [virtual_server.http.websocket_server_profile](resources--application_profiles--reference--group-002.md#canonical-2133310101022000-1330312111332132-3210102331201100-1113232033330012-0200001232322231-1010212032131322-0301022332301131-3030020121130200)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3130120121301111-0232122113022331-0133313310022200-2330232310032203-1332022200120302-1321130233030333-0131333010323133-2230331203221003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023213311221322-3021210021013213-0221233130220332-3100220223031030-0213230133033330-2302230030310223-0021132010030321-1221200203023231"></a>

## virtual_server.http.client_ssl_profile — client_ssl_profile / 013332110010 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.client_ssl_profile

<a id="canonical-2003221132223331-1212333232201023-3211311333310031-2221202111301221-1323210013230332-3312220301221300-0303321020102203-2122120110233333"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102100313200100-3033121113003321-2032022232201300-2213110203130002-2201031023312103-3012222001311110-2021322300333311-2031102002230101"></a>

## Direct properties — client_ssl_profile / 013332110010 / 3

<a id="canonical-3300332013002312-3210111301031111-3012202101030120-3211323013131232-1312222001320310-2211103133103111-0100030301231002-0323210022012010"></a>

<a id="canonical-2123311032221023-3130220302220101-3123322001322301-0120113202010032-0120322311121333-1123111222230113-1232030001311300-0001222202333231"></a>

## kind property — client_ssl_profile / 013332110010 / 4

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

<a id="canonical-2031111110320233-0033303223312000-0013213013110100-2001102310113302-2100311030223303-0200333220310101-3222220300022122-2131230211322113"></a>

<a id="canonical-1332223121312201-0121321100003302-0213201013200110-3201010330330113-1311123313302112-0102002102301113-0212230031022123-0031302113100021"></a>

## name property — client_ssl_profile / 013332110010 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3333032332002022-0023210022110332-3202002121330201-3123123001320133-3000221013122223-1222221202331031-3211103312010311-1101121200102033"></a>

<a id="canonical-3331201111323030-1203321010200100-0322010303212021-2322023212223333-0021210201331122-1001110302012200-0032201003023123-3021110201101121"></a>

## namespace property — client_ssl_profile / 013332110010 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0201312331013222-3021233111130210-0310113000300213-2310122230233121-1211210333210031-0320203032213200-1312031203132210-2200022313012022"></a>

<a id="canonical-1201001202121021-0020323133001200-0331110001333310-0310020101103212-3122013310331020-3000132102103202-3232021213231121-0102331211020032"></a>

## tenant property — client_ssl_profile / 013332110010 / 7

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

<a id="canonical-2032321300031201-1102321000211001-2000213130132310-3231330313312130-3012223333302123-3102320013222011-3002233220302122-1212100222200310"></a>

<a id="canonical-0221102022220101-1030220220320130-2213310233233332-0223202130300101-1232022230302313-3302212013101230-2132301333230102-3212000322113002"></a>

## uid property — client_ssl_profile / 013332110010 / 8

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

<a id="canonical-1101120133200313-0313023230013300-3301020300102212-1303131122320203-2011010020031332-0300112103223130-1011201022122112-1003333002222133"></a>

## Next pages — client_ssl_profile / 013332110010 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0332330302121333-0200202323301323-2011232011103020-2101231222032323-2011023122121013-2201030101010301-0301300001200132-0033123212100321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020130001030320-3011010122313102-3311120002302100-3032301132323033-0202213101213211-2113300223001003-3031303302030000-1203332332200130"></a>

## virtual_server.http.http2_client_profile — http2_client_profile / 312202033232 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http2_client_profile

<a id="canonical-1022312112332200-0222223130011023-1211321323203210-1222020210311023-1120011101100232-2023112312302133-0012200131203330-1232112223113013"></a>

Type: `"object"`. list nested block, Optional.

HTTP/2 Profile Client. Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0110011212021201-2030003022200302-3120003033231031-2210223110200120-0010213122211323-2330111001302103-0103011310111323-0332322310213130"></a>

## Direct properties — http2_client_profile / 312202033232 / 3

<a id="canonical-3101002331323112-2223133331130222-3111323123123223-3111021303031302-1213213010321303-3300003323313303-3101322130311202-0103130121323332"></a>

<a id="canonical-2122130123303200-2020130130200330-1310312331000131-1232122001123321-1002300202022311-2212101202033223-1221033033233102-1032120113321301"></a>

## kind property — http2_client_profile / 312202033232 / 4

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

<a id="canonical-3322131032121130-0313332202013030-1032331120313213-1011200231110213-1103222011113212-1013311131311203-1032300021000333-2033300322023010"></a>

<a id="canonical-3102313032220121-0302023123003230-2301102023202100-0302223320120121-3133112033200001-0232131320201232-1313022320000131-3310323233230310"></a>

## name property — http2_client_profile / 312202033232 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0333330322231131-1123322131032031-2311313301022110-1110211102003001-3211120023113332-2233233300331313-1302302210003320-1301121102112330"></a>

<a id="canonical-3203233132011333-0133330102332233-1322003031103320-2322222020323030-0000132023222211-3211322012022110-2213232332130320-2001311013101003"></a>

## namespace property — http2_client_profile / 312202033232 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-2103320101311310-1013130302023321-1102003220102122-0132001321212111-3022030210201221-3211331321013022-1011331003111231-0032320230233002"></a>

<a id="canonical-0210320020113021-3331313012201103-2312021221112111-3321202010022333-3300233222120131-3103301213223011-2233312331330010-3303110102230301"></a>

## tenant property — http2_client_profile / 312202033232 / 7

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

<a id="canonical-1121111232030031-2313113321020320-3211322311103002-0321310331030301-3201231100210230-3002203221331333-2001030022122330-2321022033330322"></a>

<a id="canonical-0010033103332101-1300201331031300-3333332321313100-0031112203223031-2113232221213211-3100223022033121-1322221113112012-2022132132130221"></a>

## uid property — http2_client_profile / 312202033232 / 8

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

<a id="canonical-3202030133302120-3300330013300321-1001332011213103-3323300033321331-0300111133010002-0002201011030223-2313000001301002-2312320131212213"></a>

## Next pages — http2_client_profile / 312202033232 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0022331112021010-3110203310013002-0212322031113123-3313230120011121-0103101002003131-0033200323322031-0033010030130201-3323002202213230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230331033013131-0310021313101022-1012211323102010-0003123301032203-0221123313223013-1023201110022222-0022032313011312-0130033223030302"></a>

## virtual_server.http.http2_server_profile — http2_server_profile / 132111311222 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http2_server_profile

<a id="canonical-1021113223331302-0220202131331210-2123132120213301-3310212101201220-2111233331121033-0022122331121111-3201000211233123-3120021232320232"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http2 server profile.

Upstream description:

Configuration parameter for http2 server profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http2_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1302333133030032-1030003322113212-1032121221032331-3233231313202233-3100121001321003-0032120231201121-1212300232132212-1010303201020320"></a>

## Direct properties — http2_server_profile / 132111311222 / 3

<a id="canonical-2221023020223311-0113331122232003-3123201221211132-0000330112100001-1313311120012110-2000001110113330-2112022021323332-3310022020330032"></a>

<a id="canonical-2211220033000101-2011010202233020-0232300302311021-3101233330011013-3001321122002020-1013012331330033-3013312232033230-2200021101102002"></a>

## kind property — http2_server_profile / 132111311222 / 4

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

<a id="canonical-3032011303010020-2103020203002013-0100103303220001-2203233132100002-2312302030322320-0132213311010201-2211233302013030-2302330322212133"></a>

<a id="canonical-1322203112131320-0220013310212223-1312331213222032-0313132003202230-1202113031122313-2031131310312230-0230301231121100-2200311123210022"></a>

## name property — http2_server_profile / 132111311222 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3001103022033230-1333001103131111-1022222132003100-0113211001231031-2121321120001103-3220003103310033-2101212200111231-1010221213303121"></a>

<a id="canonical-0131102101112201-0133110000332023-2133131221120030-1201331103211131-3210000122123122-1111300330111222-1311133120131022-0210211031300010"></a>

## namespace property — http2_server_profile / 132111311222 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1031220312220123-1011233100132310-1230331123101311-1233120220201033-3223312101201130-0031002211102313-1312010113110301-0111010221020213"></a>

<a id="canonical-0320322032222303-2320312212001212-2231333133132021-0111102221000313-2212313300302311-1321211233301020-3321321001100220-2230100320133021"></a>

## tenant property — http2_server_profile / 132111311222 / 7

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

<a id="canonical-0330130332102210-1110033022220012-0031132320113020-3303220320210011-2132210110312230-2003232211203320-3103023211322002-1010220333130301"></a>

<a id="canonical-0112323001312010-3320301223111312-3332030331330212-3321003313330233-1333323300232300-0330111330031032-1132023120213221-2203023021332120"></a>

## uid property — http2_server_profile / 132111311222 / 8

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

<a id="canonical-0320020132331202-0210121122011133-1130103322132301-3000130233231303-3220201212021303-2033330003020020-2122301102023113-1012132022201213"></a>

## Next pages — http2_server_profile / 132111311222 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2122303303313200-2013130123113121-1031102010022023-1131131323001322-2322123310020201-0021023332013321-2103212212003311-1320211331131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301302200300301-3202322203313131-0330030210211131-0102331223030321-3333222221310333-1231130203330212-2023031220230110-2023313110303202"></a>

## virtual_server.http.http_client_profile — http_client_profile / 333120332112 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http_client_profile

<a id="canonical-1113100222021131-2011310210010022-0120023101322100-0330202311321000-3132200232113230-2323110010233002-3321311201232300-1222113233120233"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2321131020111311-3030232223012320-2200020001203320-1032322022112220-2130331233303103-0311211222232312-2031113101323102-2200301301111230"></a>

## Direct properties — http_client_profile / 333120332112 / 3

<a id="canonical-0101021333322322-3303132201311230-2100031032312123-1333110312111133-2220121011121020-1301131002333022-1313130302222233-2223123131313002"></a>

<a id="canonical-0001223020013013-0230200233133232-2212221122102321-3322332223010130-0033020223333112-3210013223103310-3102020023000003-2333111120023131"></a>

## kind property — http_client_profile / 333120332112 / 4

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

<a id="canonical-3220303301223023-2123211100031220-0200233111011000-1111223322003323-0100301203001321-3133110102230110-2101212210012010-0230313301323011"></a>

<a id="canonical-2010001100031221-3201332111331100-3231000023022201-0133310110030002-1020232021022211-0300221312312121-1101010020203310-0112123132300221"></a>

## name property — http_client_profile / 333120332112 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1030111303102331-1302321220333001-1201010133021102-3213303223210011-3231003312330102-1332331322310330-2201211132010113-1312111020233022"></a>

<a id="canonical-2031311301013220-1121102111132031-1103103213201221-0102301013220101-2210210331301000-0103120130023133-0031023222100321-2301011302331112"></a>

## namespace property — http_client_profile / 333120332112 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-3300100233013102-1012010320002031-3302131002012222-1210310113300010-0220032232010232-1103131003232011-0321201022302000-2112030031233100"></a>

<a id="canonical-1321122003103112-2211123101320131-3323220132012201-0123112230123101-3033112120021320-1113023123022033-3133220022122312-0122311001102233"></a>

## tenant property — http_client_profile / 333120332112 / 7

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

<a id="canonical-0311132203322311-0233130001201300-1233022103211333-2133013320020233-0111320303133032-2212132112122010-0233220013131303-2002022300220012"></a>

<a id="canonical-3320232300020021-2023000132013322-0320312033100213-3302203133133031-0011022311121100-1030322110132130-0230110020231320-3311021301130032"></a>

## uid property — http_client_profile / 333120332112 / 8

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

<a id="canonical-1033132212322233-0001213132213002-2032322030032103-0333010200230033-3131010010012332-3303203200212000-0132303130020130-1020210233031232"></a>

## Next pages — http_client_profile / 333120332112 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2322322120023310-2331230320202223-2321331132010123-1120303200320012-2132131011333223-0311000303222001-0111212033133333-2113032201233303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201022131130031-1231022302023111-3001122302001221-2033111030230110-2100130012321023-0132013331203310-2220202133303232-2012203230231110"></a>

## virtual_server.http.http_server_profile — http_server_profile / 020312122121 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.http_server_profile

<a id="canonical-1213110300000223-2122321330223221-3123002011000213-0021132300332002-0032022023223023-3133330032100221-1300020310321322-0223322210021032"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3312210220220323-0000011302201031-3301313013001032-2020001210203233-1211222220231221-1002001202310302-1022122313201120-3332003110011312"></a>

## Direct properties — http_server_profile / 020312122121 / 3

<a id="canonical-0123112123120200-0102021220033201-1200130211101330-1032020001301303-0032013331130301-2130212010113221-1112202111232321-1123232123133233"></a>

<a id="canonical-2221122232203013-1331021011111011-0102020330203003-2212011310101111-0033123210132312-3101132201230113-1300300013312231-1012123220011030"></a>

## kind property — http_server_profile / 020312122121 / 4

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

<a id="canonical-1300111231210032-1020010010131011-3012310202020113-1312211021223123-3133001111312332-0201300000232011-0133320013331312-2302133101002003"></a>

<a id="canonical-2312101212200121-0233303132100131-3333032033112131-2203310233031000-0033330301331232-1003223032202133-0323013003022023-0121012333321003"></a>

## name property — http_server_profile / 020312122121 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3121232322132020-3332023201332123-0233203032212023-0220012213023322-2313011333231313-2313001112232003-3300222222031020-3100101322120233"></a>

<a id="canonical-1230231200230103-2123023222211002-0200203312023122-0300311010200003-3030333003223110-2203213221020220-0100302013123320-0132313120100133"></a>

## namespace property — http_server_profile / 020312122121 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-2323312323003313-3033312331322211-3203313301020221-2132111232212003-1333322322230012-3332113030213222-0231223211023212-0121210300123022"></a>

<a id="canonical-1103201020311122-2010230320213120-2321033230303123-0022212023122000-3212032231332003-2213211312310320-3012222003213103-3313120211021112"></a>

## tenant property — http_server_profile / 020312122121 / 7

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

<a id="canonical-0122133312001331-1022310203010302-0233012103332213-1311201022311230-0112112322122231-0231010330131111-0203100211213130-2232202123322202"></a>

<a id="canonical-0101001233313213-3000320222313332-1000133103213302-1220013213321330-1033100021112000-1111321130131310-1313031133220001-1112121301120112"></a>

## uid property — http_server_profile / 020312122121 / 8

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

<a id="canonical-1210212211302332-2032231000230303-0203103110032231-1223013213302201-1333212011113321-1300303222010022-3021000012323233-1300103110213130"></a>

## Next pages — http_server_profile / 020312122121 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0212200100013300-2010330003202013-2113300023302120-3201222310221023-0313220131110103-1312321001211122-3021120111312003-3200020321033213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0000011023130033-0131003130032033-0021233231002212-1233013223232203-2133023322111020-0132313122302023-0313103002103212-2202021312101302"></a>

## virtual_server.http.ocsp_profile — ocsp_profile / 132312333231 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.ocsp_profile

<a id="canonical-3100301002030221-1333133211330332-3023213213021130-1030100001100123-2231320331330011-2103212330310100-0332130113313302-2323130220000223"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for ocsp profile.

Upstream description:

Configuration parameter for ocsp profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
ocsp_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0313010323120223-1223013231133001-2301212023202020-1232002002213201-1101021201232200-0003011113231032-0201012021003031-0230102112231222"></a>

## Direct properties — ocsp_profile / 132312333231 / 3

<a id="canonical-0213213000023030-2031213303111300-2301011013312220-3233030330333331-2303320233112223-2302213020122111-2022131003003223-0212121233202012"></a>

<a id="canonical-1332100130111223-0232120101132220-1210102221330130-2133032232020012-2310201201333130-1123030130212023-0322110123030222-2002313222132001"></a>

## kind property — ocsp_profile / 132312333231 / 4

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

<a id="canonical-0130120301011100-3200223212121012-2202001302003132-0303010233320232-0011223121021233-0301013012310000-2003313030001033-1213100102333200"></a>

<a id="canonical-3233133320303322-0102313220220022-3013111113103123-2311113113023132-3331123001102120-1223313303103001-0010020112212003-3003233302301123"></a>

## name property — ocsp_profile / 132312333231 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0300202233010331-2211211132210022-3103231202213033-2333120002112232-2102132133211300-3331102322031120-0030033330013213-3222201230021123"></a>

<a id="canonical-2300320023212112-2321023223131101-3211332010211331-3310312332203000-1210210022211032-3322311202022202-1122013321033202-2032110332121132"></a>

## namespace property — ocsp_profile / 132312333231 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0311301120303300-3322213301301012-2130303131032301-0032202323001220-1212123331330100-0101022012112321-3020003212023300-0300201210012211"></a>

<a id="canonical-2012121023013231-1201011000122002-1113223031213113-0013311022333001-1122011013220032-2000100102213001-2200133312200331-3223000220102103"></a>

## tenant property — ocsp_profile / 132312333231 / 7

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

<a id="canonical-3001331120213330-1222231232223200-1110021010300110-0130323311033013-0013120233222121-3123011202013232-3111003112311002-2100103111312221"></a>

<a id="canonical-2131321023330001-3320321330131011-3222001230321010-2303230312021112-2220311333313330-1022101011212303-3221212022230201-3122210201021331"></a>

## uid property — ocsp_profile / 132312333231 / 8

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

<a id="canonical-2130130300221233-1321331103001300-1010011203022322-0023213100213030-0112101003301133-3220001330021323-1231221110010111-1103123012102232"></a>

## Next pages — ocsp_profile / 132312333231 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3232011302112010-0200312011113331-3101111203223133-3020113313300300-0023113130010221-0133232012023323-1002302300221201-1331300202001202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123100110320230-2111010203122031-2203110222202310-1031231312213211-1300132322132303-3202300032012102-2321202301103321-1010031002103331"></a>

## virtual_server.http.server_ssl_profile — server_ssl_profile / 102331020120 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.server_ssl_profile

<a id="canonical-3010133303300011-3222112300303320-1201223020122330-2133200201010012-2123103122120202-1121303321332202-3320312331123202-2000333302030120"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-2312333122121122-2231122003111010-2220130020031003-2331020011120121-1003122222123231-1111010033010300-1223131202131223-0010212232231031"></a>

## Direct properties — server_ssl_profile / 102331020120 / 3

<a id="canonical-1012310300203232-2210311031133003-1322030232220323-2001000301031231-3120022313323033-2022301133213233-3330000101120030-0303300012102210"></a>

<a id="canonical-3222302020330130-0222222230010203-0232011131322201-2232223013022200-3333003100100333-2022233312333113-2000300120123232-1213322321033131"></a>

## kind property — server_ssl_profile / 102331020120 / 4

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

<a id="canonical-0113130033330003-2001102213130102-1213331212222101-2132022032322212-2310003201003202-3321223130121131-1321110310331233-3103312332332022"></a>

<a id="canonical-1103212121131300-0110033331223002-3333201012203101-0110201213012202-2100000103301201-2102213131103300-2013031021101300-0030001110322322"></a>

## name property — server_ssl_profile / 102331020120 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3103311222322332-2313101313122333-3303311220230231-0012211131102130-2213100321122230-3332020232033121-1031011203311033-0122011212102310"></a>

<a id="canonical-3001333212033021-3323331012202022-3133003322333322-0000111312210132-3112130232201133-3102032302002033-3330311001122112-2310130133121111"></a>

## namespace property — server_ssl_profile / 102331020120 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0311331321301221-2032133113211303-2112213231301110-0201103110220200-3003031010213011-1133332202222133-2110211303220333-3312310130001133"></a>

<a id="canonical-1200123112310110-2110233300111232-0113130212002301-0310310123303100-3000112320230101-1021001301002301-1231103133211331-0113203103333302"></a>

## tenant property — server_ssl_profile / 102331020120 / 7

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

<a id="canonical-3000002222332213-0130233303102212-3121023212100231-2021021123110012-1303203313201212-1130232013301210-2011121003131103-0312332303203130"></a>

<a id="canonical-2023122330103123-3201230133130313-1000330230313213-3001021000022230-3133330302123011-0023112211110302-3120201133013003-0023322332202330"></a>

## uid property — server_ssl_profile / 102331020120 / 8

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

<a id="canonical-2010133330010021-1320122113031233-1010212313100301-2230230210113011-0201020031320231-3012300032010023-1322212302030033-0323103220010321"></a>

## Next pages — server_ssl_profile / 102331020120 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2003122102331011-0102223311211013-1113200013103313-3101102121131201-2101310003131021-3232021132033212-0313132302121023-2222132212203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210212331321112-1330313212003231-0320310223131111-0212110302133321-0331013130013100-3213113303302121-3202220213031032-3222020130313203"></a>

## virtual_server.http.stream_profile — stream_profile / 001120320101 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.stream_profile

<a id="canonical-3102013030110002-1100221012323030-3312013033112103-3012023300122022-0132211310102010-3301002032211022-2110002101003312-1033331011012212"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for stream profile.

Upstream description:

Configuration parameter for stream profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
stream_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311231002233221-2010311310313122-2222222322000320-1112023202330010-0213233023021023-3211310112231033-1121030121312221-1203033310222310"></a>

## Direct properties — stream_profile / 001120320101 / 3

<a id="canonical-0033323323112113-3121331332122111-1311120203312000-3203113121021202-1322103133020021-3203230201010012-0001223001120103-2232223001321020"></a>

<a id="canonical-0032312101131320-0120111002002230-2023003100221031-1300330211113111-0133033132020322-3003202221011211-2202312231331312-1003013332202232"></a>

## kind property — stream_profile / 001120320101 / 4

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

<a id="canonical-2333233112332102-0020003323010312-1121221332221311-0012213223102203-2311322033020333-2130013013203211-3313231221300103-0132200032003133"></a>

<a id="canonical-1203202022323131-0220131021212103-2010322131321020-1001333031031030-0321231320332300-0023021313002012-1332230130121220-3102031032203302"></a>

## name property — stream_profile / 001120320101 / 5

Type: `"string"`. Optional.

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

<a id="canonical-2020130232330213-2101101323212120-0010000313203013-2002103202003130-3233023012013332-2122102130203323-2321210222122320-3011210113103112"></a>

<a id="canonical-1010122120012323-2310303122021031-1311330300311202-1202330230331211-0211322100333222-0033112123331100-1321322102011032-0102220203031331"></a>

## namespace property — stream_profile / 001120320101 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-3133301103203232-3203121101200023-3133323320130221-1121000120110131-0131212201012111-0220131222022302-2233121332221033-2300300131010313"></a>

<a id="canonical-3030233202023132-0023322230223231-3201133202100210-2333321002102103-1313310233030113-1202302331033100-2130230101112122-3000013111000023"></a>

## tenant property — stream_profile / 001120320101 / 7

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

<a id="canonical-0102131022332311-1301310130321013-0101322320030322-0130333101203301-2331030033010331-2310202032210231-2301021110032231-0131310220211212"></a>

<a id="canonical-2311130322010201-0122313112012002-0222221112013321-3111000133302233-2112212003013200-0113111211121132-1121212310013220-3103102122202120"></a>

## uid property — stream_profile / 001120320101 / 8

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

<a id="canonical-1231101002223013-0032010123111303-3101310301310310-3013013003103131-3032321322212203-0102103232311002-2310031202300120-3231213203101001"></a>

## Next pages — stream_profile / 001120320101 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2223311101321210-2223313131302002-2310010033333222-2122013110101211-1021221123210103-1210221022002331-2101300030310020-2331210030200231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323130130302222-0031000121001101-3311120323113312-0212022212210001-0001011311320010-2230133221030203-1202031000233211-0100213331100222"></a>

## virtual_server.http.tcp_client_profile — tcp_client_profile / 300022313002 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.tcp_client_profile

<a id="canonical-2222103002213222-1321203331302022-1212333301101011-1232112221113131-2021123322210003-3330013232203221-1323211110021210-3203312102321133"></a>

Type: `"object"`. list nested block, Optional.

Protocol Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0033011301121000-1212320330302321-0231333232032211-1233300131032010-1331100312101321-1223110010022303-2121013330222302-0110101131322102"></a>

## Direct properties — tcp_client_profile / 300022313002 / 3

<a id="canonical-3100210000103122-1130131103112020-3303032232032300-1233220333223213-0023323313231002-1203010022010012-1001123231120003-3331313322111012"></a>

<a id="canonical-1333312231201102-2123111133300013-0122012002231213-2022212313223313-3020222122220332-0333113312110301-2033213100132110-3003320100321013"></a>

## kind property — tcp_client_profile / 300022313002 / 4

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

<a id="canonical-2101110333101233-1300123303133222-1202110230223302-1321102021300203-3002310000031003-2322223030230020-0202121331130120-1023311102201100"></a>

<a id="canonical-1030203012220130-0312130202112120-3010321113333013-1113031121233201-3333113222312313-0000232023201133-2133230012001321-2122131030231233"></a>

## name property — tcp_client_profile / 300022313002 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0013311201323312-2211312312312022-1022302111030002-1331110120012212-0023300031202200-2032313321123311-1030033232201332-0020120103230000"></a>

<a id="canonical-3232311110222012-1311211333130213-0123223033313001-3103002021203210-0302332220233011-3333020010101113-1203200321220223-0020130323000220"></a>

## namespace property — tcp_client_profile / 300022313002 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0321200223302101-0100131032311111-0222202033010030-3033233333010002-1332322212302022-2311011010122100-3020331310201322-2132231001330320"></a>

<a id="canonical-0330011101132023-3301223310033321-3132022231213020-0110302303100331-0220012310231110-0020313323223322-0022323213032100-3322322213211333"></a>

## tenant property — tcp_client_profile / 300022313002 / 7

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

<a id="canonical-2003010303310013-0031210011230133-0321133231132022-3110212212101030-2200021033221232-3213310211310021-0301001302231330-3112322321131231"></a>

<a id="canonical-3101010222332033-1312030101211220-3221001121302101-2021322301130030-2103313022031320-0201012033310233-3321001331223203-2311212223222111"></a>

## uid property — tcp_client_profile / 300022313002 / 8

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

<a id="canonical-0300013023320203-2002223020030121-2113223110100223-2302210012230232-1031032000213323-1211220333200203-3212311000023022-3323203101132111"></a>

## Next pages — tcp_client_profile / 300022313002 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-3313330121201111-0233011211100030-0312301221013133-0221322322010322-0221231321001010-2010112300003100-1220230110101232-1323330130100112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203101123131222-2301002321233231-0201233131201011-0223022003200002-1221221311012300-3221100000013211-2232133130200021-0303133131200010"></a>

## virtual_server.http.tcp_server_profile — tcp_server_profile / 003313130011 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.tcp_server_profile

<a id="canonical-0032322120000231-2113212203313122-3322322010120231-3301111113111021-2031230313322100-2220221313011021-3303121310113132-0022300100310032"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for tcp server profile.

Upstream description:

Configuration parameter for tcp server profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
tcp_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002111100133300-3323220321200231-1303021300103000-0230322232010203-0203002002233200-1132302223120111-0203333103002001-1122103122322300"></a>

## Direct properties — tcp_server_profile / 003313130011 / 3

<a id="canonical-3330231201203003-1031222202310302-0120220020202323-2322323012032011-3311300120301120-0222301302122033-2301321023001223-0011000202001000"></a>

<a id="canonical-3300222100231220-0121311321102311-2010213303122321-2113303213123130-0100210210230000-0001012313122023-2333212123120120-3223232103110211"></a>

## kind property — tcp_server_profile / 003313130011 / 4

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

<a id="canonical-3101211303223032-1231020320220223-3022020310322030-3133221330020101-1201001203121132-2101202221332231-0103020131322130-3012002001120232"></a>

<a id="canonical-0012221302203023-1113110331321323-1000020001221301-3022333332200013-0312113003010323-0123322101013323-3131210233130130-3120331021302122"></a>

## name property — tcp_server_profile / 003313130011 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3021032011003322-1332330220020233-1023122033131112-1312120123200023-0213320013313013-3132003313111030-2213012133020212-0220220010130321"></a>

<a id="canonical-0321000312101000-1231111221300120-1230221321200230-3322200102200023-2303323232300330-0011130001120200-3012100110000103-2222033110303002"></a>

## namespace property — tcp_server_profile / 003313130011 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0223011121300023-2212232213313220-2202230031232023-1332223320023032-2211330212023002-1322301003313011-1013332330122023-1313100031320223"></a>

<a id="canonical-0311202132312010-0212231300302011-0331133303300130-2323220231211102-0321213333322332-2321113131132301-3101200013120110-1100011001013023"></a>

## tenant property — tcp_server_profile / 003313130011 / 7

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

<a id="canonical-3012110322231301-2021220300300201-0311100000133220-3032032033330021-2132321233031013-3232202101330013-2012220323221322-2023123201012322"></a>

<a id="canonical-0310211311132121-0210320221021212-0023303333020030-3300112003213200-2203011101101110-2313111202222211-1310001230213030-1021000203222023"></a>

## uid property — tcp_server_profile / 003313130011 / 8

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

<a id="canonical-1110301321021022-1131302120132332-3320222303220221-3031001020233020-2311103133123021-1111230131232320-3311301031102000-2010000201101100"></a>

## Next pages — tcp_server_profile / 003313130011 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2103003320213123-2333013120322010-2312232213313300-1031221031100011-3332121220121033-0110233202012310-0011130201222110-1123332212132030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011211310013002-0020200033212310-1001212331202101-2013111311033003-0212103300131331-0130321012221222-3223132323212103-3103223001201302"></a>

## virtual_server.http.websocket_client_profile — websocket_client_profile / 033112331322 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.websocket_client_profile

<a id="canonical-2223110220120301-1213223332122020-1320112133333333-0201211010202031-1022132031331300-0332032202323012-1303202122321332-2231221120220032"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Client. Web-related configuration

Upstream description:

Web-related configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023333003330133-0300233000322203-0323311130210103-0132132331331133-1223030133030320-2132022301000103-2032131331010301-3101233033121133"></a>

## Direct properties — websocket_client_profile / 033112331322 / 3

<a id="canonical-3030032010131311-1001103223310323-0011013221123013-0013102111032322-2023230033303313-3033131033103131-0013000113201101-1313021313110222"></a>

<a id="canonical-2202221113212222-1120222002020102-0222130330113022-3113120233310300-0121023022021122-2302302233211223-0202330200021300-3222022220113021"></a>

## kind property — websocket_client_profile / 033112331322 / 4

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

<a id="canonical-1122302322322001-2102320321021121-3110330220332222-0011032330322310-0311312112011030-0003212023003011-3103113013001330-3120132110033101"></a>

<a id="canonical-1212032132010230-3032013233330210-3000121320222332-1122001312321301-2200312330012221-1131230010001003-3030201321111031-2000120130332332"></a>

## name property — websocket_client_profile / 033112331322 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3101102123202032-2113102303113130-3020312322030012-3101312332022031-1100313233303333-1131310021301322-2233032011011230-0001033110103012"></a>

<a id="canonical-0203022222201130-3011131101023012-3302023200001000-0310301231320202-1002330001002233-1232122212302032-2110331110001003-3303032020132201"></a>

## namespace property — websocket_client_profile / 033112331322 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-3200203130131213-2211031120033003-1103221001102300-0202201030013133-0023201102112131-1110003113111032-3301000331321032-3010303322220222"></a>

<a id="canonical-1112202102003312-1332303330333013-0302203303333321-3302333120313003-2223033003230332-0100023032130223-3031220000131210-0030200102332321"></a>

## tenant property — websocket_client_profile / 033112331322 / 7

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

<a id="canonical-1331111301122210-1023033323332121-3112312330002232-1223210212221120-2332030012232213-1102232323103320-2032322201002202-3221031013320323"></a>

<a id="canonical-0120000011202223-0320233011031112-1330110112033001-2112121300331011-2232303033101311-1011011311231221-3021203100301030-3021001313233002"></a>

## uid property — websocket_client_profile / 033112331322 / 8

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

<a id="canonical-0023231321201200-0211032021120013-0300332101331113-1310120112110310-3223033221233330-0223021112012101-3201320300302030-1010322201030132"></a>

## Next pages — websocket_client_profile / 033112331322 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-2133310101022000-1330312111332132-3210102331201100-1113232033330012-0200001232322231-1010212032131322-0301022332301131-3030020121130200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120331323220112-3000020020233111-0221132003311331-0323123102201100-1111000322023313-1313310230100210-2113032230333210-0100002103130303"></a>

## virtual_server.http.websocket_server_profile — websocket_server_profile / 212222222002 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- virtual_server.http.websocket_server_profile

<a id="canonical-1202232302031303-0013323212010003-2020001202002320-1103221331022212-0113213022111122-2102330111012133-0022131122303201-0103012213310211"></a>

Type: `"object"`. list nested block, Optional.

WebSocket Profile Server. Web-related configuration

Upstream description:

Web-related configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
websocket_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0312032123121303-1110333011320002-0130020303231032-2121000122023202-0132032322201311-0130132233323221-2012010132101011-0133022122310033"></a>

## Direct properties — websocket_server_profile / 212222222002 / 3

<a id="canonical-2131021222322302-1310112022132011-1112213013312312-0223320031030002-2101020321001130-1121113321202122-1212200122031320-2311103013020331"></a>

<a id="canonical-0021210202321311-0301011031002331-0012022000220310-1320222233322321-0122213111103203-0121000301203021-3300023030121032-3010023321211003"></a>

## kind property — websocket_server_profile / 212222222002 / 4

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

<a id="canonical-1231113113020320-1002213012123130-3102331221323121-2101130331232211-2001000333102302-0032102100002210-1210213211303030-3310211303022230"></a>

<a id="canonical-1112312111013103-0010001011011301-0330213032312322-1210000002232110-0013230213102010-0112021013112213-1231312110101330-3200201022023030"></a>

## name property — websocket_server_profile / 212222222002 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3333101202000331-0331320333223120-0310321333023120-1330033023011123-2300203111011200-2121130100230100-2133221300332113-2231023220210223"></a>

<a id="canonical-2322203130000032-2332003301101331-1332132331232021-0102322111222011-1002221300022131-1331201003030313-1033202133202231-0102113023302232"></a>

## namespace property — websocket_server_profile / 212222222002 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1111213312131202-0002000100232122-2301033120001112-1233023323203101-0230010112233302-0321021031112231-1123033221223311-0312211321133331"></a>

<a id="canonical-0100302000331031-1203122002333101-3200022100301023-1133132012031111-2221231103300200-0032120122202321-0233131231301310-0111030203022311"></a>

## tenant property — websocket_server_profile / 212222222002 / 7

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

<a id="canonical-0031201230102221-3321111012120001-3223303333010333-2303231200023111-3333102132002302-0000301232013021-0231233121031122-3021213231131011"></a>

<a id="canonical-0330020223111020-3320031323301130-2330030120232130-0021132113223221-0110110310303332-3310232322313310-3113101310031223-0003330232210131"></a>

## uid property — websocket_server_profile / 212222222002 / 8

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

<a id="canonical-0321133210022122-1103112323112322-2100013331311332-2103332223032123-2300220221130003-1213123223010003-0003110130213220-0211313332130133"></a>

## Next pages — websocket_server_profile / 212222222002 / 9

- [virtual_server.http](resources--application_profiles--reference--group-002.md#canonical-2132220100213310-0202210200313002-0313022312321210-1221030110101102-3203001231322331-0123200300000220-2000103122103301-0330030321001322)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022032301122103-3322001111102012-1031102102010010-1301331202012031-1111021233011312-1120201231332232-3323231221233210-3002221110212032"></a>

## virtual_server.http3 — http3 / 200203132213 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- virtual_server.http3

<a id="canonical-2002231130013121-1320001010323121-2331302320313322-2221020113232332-1113213211110201-1022230001113101-1132301132322222-3120032223001131"></a>

Type: `"object"`. single nested block, Optional.

HTTP/3 profiles.

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
http3 {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130222230313030-0201201220231233-0122020330233302-1032020010131323-0001022212321303-0320301033330212-1321103122200200-2203001020312123"></a>

## Direct properties — http3 / 200203132213 / 3

- [client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-1013312311301220-1033133010113103-2312213023210210-1011131133330212-0333233011331131-2323002012312101-3003233002300221-0011220013121002): complete subsection reference.

- [http3_profile](resources--application_profiles--reference--group-002.md#canonical-0330103031220011-2233020132233102-3122223210303012-0113012332332203-2012302030003300-2232201011112130-0001003312201001-1332113123333323): complete subsection reference.

- [http_client_profile](resources--application_profiles--reference--group-002.md#canonical-0110323101223223-0010110032230301-0033333001003323-2301130020131120-1231020331020222-0332330333033103-1020112033001230-3210003311322323): complete subsection reference.

- [http_server_profile](resources--application_profiles--reference--group-002.md#canonical-0220231220211221-1300100113032203-1202030023031212-3101012201312312-0221100301100212-0101010110231132-1003301322301023-0020200203102112): complete subsection reference.

- [quic_profile](resources--application_profiles--reference--group-002.md#canonical-0310002212103133-2323121222001013-1130122010313012-3002210333000020-2312022322312202-0112101012133013-3332302221010313-1300310211101012): complete subsection reference.

- [server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-0100301033103101-0132322221003231-0130012010003121-2100011022032313-1300221101232323-3021101001012203-3000301233212002-2110021113331313): complete subsection reference.

- [tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-1001310203000122-0311102113233113-3302310000201022-0200230311200102-1233203113320233-0132232313301222-3211300233232121-0102200101100012): complete subsection reference.

- [udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3022120130332131-3313210302001322-3231201032200020-1312101002032320-1200300302023322-3130201220303033-0122013032031230-2112001220310130): complete subsection reference.

- [udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-1001201132330211-3022323222212131-0312122213332223-0022033331023200-1322321300101231-3303303132220203-0323330033312310-0211000013331213): complete subsection reference.

<a id="canonical-3010100120102013-2002201213303311-1000011221221231-0231123113200112-3032333012201021-1302032123231112-0122020311313113-0220022301302113"></a>

## Next pages — http3 / 200203132213 / 4

- [virtual_server.http3.client_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-1013312311301220-1033133010113103-2312213023210210-1011131133330212-0333233011331131-2323002012312101-3003233002300221-0011220013121002)
- [virtual_server.http3.http3_profile](resources--application_profiles--reference--group-002.md#canonical-0330103031220011-2233020132233102-3122223210303012-0113012332332203-2012302030003300-2232201011112130-0001003312201001-1332113123333323)
- [virtual_server.http3.http_client_profile](resources--application_profiles--reference--group-002.md#canonical-0110323101223223-0010110032230301-0033333001003323-2301130020131120-1231020331020222-0332330333033103-1020112033001230-3210003311322323)
- [virtual_server.http3.http_server_profile](resources--application_profiles--reference--group-002.md#canonical-0220231220211221-1300100113032203-1202030023031212-3101012201312312-0221100301100212-0101010110231132-1003301322301023-0020200203102112)
- [virtual_server.http3.quic_profile](resources--application_profiles--reference--group-002.md#canonical-0310002212103133-2323121222001013-1130122010313012-3002210333000020-2312022322312202-0112101012133013-3332302221010313-1300310211101012)
- [virtual_server.http3.server_ssl_profile](resources--application_profiles--reference--group-002.md#canonical-0100301033103101-0132322221003231-0130012010003121-2100011022032313-1300221101232323-3021101001012203-3000301233212002-2110021113331313)
- [virtual_server.http3.tcp_server_profile](resources--application_profiles--reference--group-003.md#canonical-1001310203000122-0311102113233113-3302310000201022-0200230311200102-1233203113320233-0132232313301222-3211300233232121-0102200101100012)
- [virtual_server.http3.udp_client_profile](resources--application_profiles--reference--group-003.md#canonical-3022120130332131-3313210302001322-3231201032200020-1312101002032320-1200300302023322-3130201220303033-0122013032031230-2112001220310130)
- [virtual_server.http3.udp_server_profile](resources--application_profiles--reference--group-003.md#canonical-1001201132330211-3022323222212131-0312122213332223-0022033331023200-1322321300101231-3303303132220203-0323330033312310-0211000013331213)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-1013312311301220-1033133010113103-2312213023210210-1011131133330212-0333233011331131-2323002012312101-3003233002300221-0011220013121002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102200031333231-0120100330223112-1311321230020123-0033113201311210-1310111310311031-3312321020310001-1222030221222000-2232221213021101"></a>

## virtual_server.http3.client_ssl_profile — client_ssl_profile / 110113122002 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.client_ssl_profile

<a id="canonical-1120110303201101-1301001103033203-2001103223211123-1023320032013330-1112100022313220-3011010013203212-3110122001132223-3111011223120333"></a>

Type: `"object"`. list nested block, Optional.

Client SSL Profile. Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
client_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3011231111220130-3020002213202330-0120010322203020-0132322332101310-1120021112113101-2003023313001002-1203100213120320-0031002330121131"></a>

## Direct properties — client_ssl_profile / 110113122002 / 3

<a id="canonical-2200200203103102-3133013030202230-1021203002003221-2110333022012111-3110302013321330-1331302101000003-1311221130300200-1333203223031111"></a>

<a id="canonical-3303031121313123-2013330220323000-3011110203233122-1233200323321012-1003122211332030-0000012200120212-0133122102321223-1001031210100210"></a>

## kind property — client_ssl_profile / 110113122002 / 4

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

<a id="canonical-3221132330131232-2332303212210212-0110223210030001-1211100200023300-2132000210320310-1133332222003221-0013313032201203-0121233203313220"></a>

<a id="canonical-1322010022103123-2102121330023001-2301231311231231-3101021302330312-1201023313133003-1232303202030010-1010002001003103-3103113013233102"></a>

## name property — client_ssl_profile / 110113122002 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1031323230323132-3303012013032303-3003133221121022-2022000000202030-3203111232200222-1311120211232033-0021311303102132-3021101232132321"></a>

<a id="canonical-3323010031003132-0302213203001200-0210312223133232-1000222122233312-0131330013100123-1211002300231223-2012121012232133-3022332211233122"></a>

## namespace property — client_ssl_profile / 110113122002 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1133102131011033-0311201313100103-3322323331312302-0223000103022232-2320103302021210-0011233002222300-3202321012210232-3221020223320230"></a>

<a id="canonical-2021322313312013-2310011101132222-2332131130120331-3100033010130102-1100000221300333-2133112311331310-1232001010111222-3003123312110132"></a>

## tenant property — client_ssl_profile / 110113122002 / 7

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

<a id="canonical-2133030102220021-3331133303233312-0010033013023023-1221011200002020-0232203313301020-0021310222210131-3203120311121221-3010003211132132"></a>

<a id="canonical-3123221111011232-3322101302011201-0323221010200101-3121310203102220-1003122331211331-1300313311333011-2032030111013231-3323032303132032"></a>

## uid property — client_ssl_profile / 110113122002 / 8

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

<a id="canonical-1300302331222020-2311311300003203-1000323303302013-3131232323221301-0023331200133313-2112031133330120-3321333101300222-2000310121213312"></a>

## Next pages — client_ssl_profile / 110113122002 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0330103031220011-2233020132233102-3122223210303012-0113012332332203-2012302030003300-2232201011112130-0001003312201001-1332113123333323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212123102011211-1002201133001232-3030111313133310-2120002211331302-2022222331203100-1121232123220330-2022332102300110-0321011231033123"></a>

## virtual_server.http3.http3_profile — http3_profile / 312202013000 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.http3_profile

<a id="canonical-0322200003130310-2100000313011310-3310110111203133-0322333010302213-2113101012322222-3021021122212110-1320200333300022-1010100310322312"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http3 profile.

Upstream description:

Configuration parameter for http3 profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http3_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110230113202222-0333220012103333-0013213110202321-1131322313020010-2121130223232210-1303101220331011-1022011312330323-1211232131011003"></a>

## Direct properties — http3_profile / 312202013000 / 3

<a id="canonical-0223023331021000-2223031010030023-2311021301121001-1130222000011003-2303133002113130-0102302331300022-2202122200232233-1221133222302313"></a>

<a id="canonical-3331001321203121-1132123313303301-1113303230212210-2123320333032310-0023102130203312-3322323210321130-3222121102302302-1003203221222110"></a>

## kind property — http3_profile / 312202013000 / 4

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

<a id="canonical-1222200231311100-1101302000230030-2231211101101303-2310111330033023-2222021031303132-2212333002031030-0300302330332212-1333010032233210"></a>

<a id="canonical-0013223132123032-2202230220103131-2121003000330302-1321123231322000-3021001100213111-2023113303103133-3330311303102130-3132220032301311"></a>

## name property — http3_profile / 312202013000 / 5

Type: `"string"`. Optional.

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

<a id="canonical-0322112021132133-3021032010020012-2023301002300111-2202223300111312-0010102130331332-3312122322132123-3221222230131233-2231132201033101"></a>

<a id="canonical-1313131121123001-1313322122303010-1203011003302231-3130232030301320-3233111301112203-2212023211202012-2112002301103032-1021110021031331"></a>

## namespace property — http3_profile / 312202013000 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-1022330111013223-3102311311132110-1012010030003301-0120212110223111-3012101001002300-3101120312120021-0111103220303012-3131302201321022"></a>

<a id="canonical-0223312100220212-1233310222122030-3313121120103311-1312121023013233-1332312013203100-0321300102033311-2130212012301032-3130020000033310"></a>

## tenant property — http3_profile / 312202013000 / 7

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

<a id="canonical-1322323202021031-2012010100131333-2233031213321333-0123121021232122-3331333123011201-2133232131310333-0110012022023220-2030022300001301"></a>

<a id="canonical-1110001220112210-2210203321101020-3111232101233003-2303213302000321-2303003003000232-3310133110130323-3223121331102211-3110230132120330"></a>

## uid property — http3_profile / 312202013000 / 8

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

<a id="canonical-2113213002032322-0100322231112013-3301032112012110-0032220210233010-0100330031120020-2320332010303321-2103003201023001-3300113302233122"></a>

## Next pages — http3_profile / 312202013000 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0110323101223223-0010110032230301-0033333001003323-2301130020131120-1231020331020222-0332330333033103-1020112033001230-3210003311322323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123021302300113-2133310210123210-0300331230121202-2210213112211031-2113023220232020-1202202011121012-3311303023221330-2100313212033222"></a>

## virtual_server.http3.http_client_profile — http_client_profile / 330211312113 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.http_client_profile

<a id="canonical-2111231120012330-1233222013020212-2013003201033331-3302211331120231-1313132022313113-0110012313212012-0131120302222113-2233021322200321"></a>

Type: `"object"`. list nested block, Optional.

HTTP Profile (Client). Client-side configuration

Upstream description:

Client-side configuration

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_client_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-3131101300112321-3213102120210012-0300102110120301-2111331312302331-3111002201133322-0010220302102232-0010230021311223-3332110103200223"></a>

## Direct properties — http_client_profile / 330211312113 / 3

<a id="canonical-1232323312312320-0203200313010222-0202101012120212-3320021331120322-3212201230230303-0132321030302112-2332132213012301-1300221230320211"></a>

<a id="canonical-3123002323320231-3212121321302222-2322221222013202-0202112130332130-2303203301233321-1311013221133331-1113013120320203-0311323121002313"></a>

## kind property — http_client_profile / 330211312113 / 4

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

<a id="canonical-3133323121002120-0220100010031021-1132302332013120-0001301021111323-3222323212312302-0331213103333032-0133000103102232-1300003332222123"></a>

<a id="canonical-0233022212100313-0303301100111221-2010103003321003-2121112112010003-0312020310213012-2011310213202020-0031203020032000-3113213202012312"></a>

## name property — http_client_profile / 330211312113 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1123133003010012-1021313333201313-1003311301312113-3133120023301203-2233121221331021-2203201101011233-3120010332132013-3211111122131330"></a>

<a id="canonical-0220210332010220-3221103302300221-2302003012330110-3010200221303200-2100220023031101-3122201212012010-3010303330002002-3001200102200221"></a>

## namespace property — http_client_profile / 330211312113 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0321320130021102-3330121302313022-1001312023200033-3023322222003110-1221000110301300-3123222120321033-3000012022011133-0210012333233313"></a>

<a id="canonical-3020232000302112-2011332131231001-0331301113010201-3100021333212133-0021032021002201-3030122012110202-0233113331120300-0023330230133022"></a>

## tenant property — http_client_profile / 330211312113 / 7

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

<a id="canonical-1102013031132323-0002311212321333-0011103210021022-2100323313321233-1202311320131202-1012003310103111-2113202221132212-3030221000201110"></a>

<a id="canonical-0012300113031032-2110123210021011-3320033232201123-2312100013303230-2223030013313131-2000111013213010-0200212233022102-0303221223121132"></a>

## uid property — http_client_profile / 330211312113 / 8

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

<a id="canonical-1022002021001030-1130110200201111-3002301301301311-2130101020322232-3333212320012111-2012103031120330-3332003010001232-0311311210120203"></a>

## Next pages — http_client_profile / 330211312113 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0220231220211221-1300100113032203-1202030023031212-3101012201312312-0221100301100212-0101010110231132-1003301322301023-0020200203102112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001223031110210-3122301021313210-1033233113113000-3022000312111020-0131111131210121-0211202013202312-1132320333113100-3200111131132123"></a>

## virtual_server.http3.http_server_profile — http_server_profile / 102111300113 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.http_server_profile

<a id="canonical-2133130302111323-0311332113010102-2203021301212221-3311003113123201-2002013122221310-0203221012012023-0203001021303203-1012221323112022"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for http server profile.

Upstream description:

Configuration parameter for http server profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
http_server_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303032201211133-2211323131033100-2120211120331130-0202002023320031-1102331010330111-1120030221330202-1321123333032301-0201103210312213"></a>

## Direct properties — http_server_profile / 102111300113 / 3

<a id="canonical-2302321123330311-2011003131132333-0323321010202100-0021123221310302-3032012110320221-0331213222200032-3232003103320030-1120123022311111"></a>

<a id="canonical-0211031211110002-0333100203022102-3322322103110202-1001233303110210-3113222132211001-2311121231020023-3130013111101132-3201312010311010"></a>

## kind property — http_server_profile / 102111300113 / 4

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

<a id="canonical-2213310223202332-0121320111231201-3003031213003203-2333233003010103-0313000100201311-0013013223110221-2223330033223010-1002031312311323"></a>

<a id="canonical-1231200030311213-1130121301300201-3310330132221200-1332320303313333-3023021122130013-0223022113321232-0302302001000330-0011023300021212"></a>

## name property — http_server_profile / 102111300113 / 5

Type: `"string"`. Optional.

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

<a id="canonical-1022320011002333-0023330332201230-3113033012211223-1020201320031100-3012002112022130-3031220222302121-0333011313030101-3013021320020321"></a>

<a id="canonical-2331210201010000-0020020111122322-2032330112133331-0111302301021313-3020010101121023-1330301232333032-0102223301322133-0213300100231001"></a>

## namespace property — http_server_profile / 102111300113 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-2031231210110123-0203123100120331-3000311221212021-1120222023232111-1200323023213332-3030321303333012-0210333021313222-1031111112113330"></a>

<a id="canonical-0100311011222201-0123332331330110-0031032022003202-3020201232023323-0320013110113110-2122303200222233-1222132201222001-1233011303303310"></a>

## tenant property — http_server_profile / 102111300113 / 7

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

<a id="canonical-3330310132210002-0302201100211032-3012312202022322-0211103022100022-0310313332103213-3302222311122223-3312010012100002-0030313103221213"></a>

<a id="canonical-3130311331131303-1311012030231201-1202320120100321-2213102121003212-1222213331331011-0202222210302021-3220200222100311-2301313233011231"></a>

## uid property — http_server_profile / 102111300113 / 8

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

<a id="canonical-0013020001030220-2003333113303330-2122221331112031-1212003101212001-0020103313121111-2320131230012331-0223021002012301-3303323223021012"></a>

## Next pages — http_server_profile / 102111300113 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0310002212103133-2323121222001013-1130122010313012-3002210333000020-2312022322312202-0112101012133013-3332302221010313-1300310211101012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202002123003133-2230221221322032-0333132300112031-1332120123010222-1311002223012230-3321012231031113-2110110232221103-0213222302212021"></a>

## virtual_server.http3.quic_profile — quic_profile / 223210203013 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.quic_profile

<a id="canonical-1302232321131103-0021111103200313-2331011300003103-0001032113331103-2123300232330232-0202033332003011-0301212020002233-3203103023000203"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for quic profile.

Upstream description:

Configuration parameter for quic profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
quic_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312311021232211-0102303020122102-3233301021102103-1233120122122030-3213003213030311-2100103003300102-0331001201221102-2232021330212030"></a>

## Direct properties — quic_profile / 223210203013 / 3

<a id="canonical-2103301310301303-2213000303330332-1210031331121212-0211120322231010-2002323220303330-3322331223133220-2021000022032210-0200001200132003"></a>

<a id="canonical-0233001200122111-2100331222022032-3011323101101300-0201200113202012-1332220222323233-1013223121132013-2231103333122103-3023323330201111"></a>

## kind property — quic_profile / 223210203013 / 4

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

<a id="canonical-1011131312312103-2332110030312221-1332211112000033-2122120023202330-0030201011213303-1223101230202303-1102102121010022-1013202121303000"></a>

<a id="canonical-3231011032113113-0032333200113213-3221212331122203-3321113213100323-1102003312020231-1022033101022320-2032223201222123-2222031121331333"></a>

## name property — quic_profile / 223210203013 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3231202020111232-2023212123121022-0013223123210310-1000103102302133-1001211103302122-1233230200223323-3022002231312230-2320032132321003"></a>

<a id="canonical-0211020011332122-1013000321213322-0033223002013321-3133230130111230-3020200001203112-1023100030321331-3200232233211201-3132000030002212"></a>

## namespace property — quic_profile / 223210203013 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
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
  }
}
```

<a id="canonical-0331231031100010-3132103210320110-0022122310201313-3321332303100102-3011332133013120-0320112123112312-0310101003223203-1010123301102220"></a>

<a id="canonical-3011331200101331-2031231031303022-2102101322013102-1233031021003221-2202011032022210-0333030233112200-2332103011111211-0021232101102232"></a>

## tenant property — quic_profile / 223210203013 / 7

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

<a id="canonical-1010230110221332-1202213231322301-2211100303001133-1233022232001102-3013100232223000-1131211102001003-3020100002231201-1313212101102330"></a>

<a id="canonical-1210112302222110-0331303331132110-3021202221223321-0313331131110121-1230011232031131-0303222201113333-1302133300103202-2200103223223323"></a>

## uid property — quic_profile / 223210203013 / 8

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

<a id="canonical-0221020313112320-0130113311003120-0100132310231211-3101231322003011-0003303012033030-0330033020103201-0023032223120122-1101210103300101"></a>

## Next pages — quic_profile / 223210203013 / 9

- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)

<a id="canonical-0100301033103101-0132322221003231-0130012010003121-2100011022032313-1300221101232323-3021101001012203-3000301233212002-2110021113331313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320332011322303-2220033111111320-3303013200232311-0121000001202300-2333232012331213-1021122301231131-2000211100030332-0331312012111013"></a>

## virtual_server.http3.server_ssl_profile — server_ssl_profile / 332102001310 / 2

Breadcrumbs:

- [xcsh_application_profiles](../resources/application_profiles.md#canonical-0000323322010330-0102233213231021-3032332212013010-3121121333300110-2011320110200320-2330030330011323-3130301212021020-3133003102002130)
- [Property reference](resources--application_profiles--reference--group-001.md#canonical-1000000221200332-2332020132023330-1113003211333221-0213203120201020-3120301011022333-3130031323012121-0123000301113121-2121031301023112)
- [virtual_server](resources--application_profiles--reference--group-001.md#canonical-2022300032020022-3032121011332332-1323022132101011-3131133312011330-0010301011103100-3020120030331310-3302110001211033-0012330101021312)
- [virtual_server.http3](resources--application_profiles--reference--group-002.md#canonical-0003212332012203-3230001001122323-2300133332232210-3333200122133222-1012320031313330-3320220130122311-1230120221321232-2013131233111132)
- virtual_server.http3.server_ssl_profile

<a id="canonical-3011020032120132-2302133210113203-2231330200220020-2011300122312220-2130310201130222-1203123211131010-0021313120101311-3230333210021311"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for server ssl profile.

Upstream description:

Configuration parameter for server ssl profile

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
server_ssl_profile {
  # Configure direct properties listed below.
}
```

<a id="canonical-1002202010233121-1201300201130023-0123301300111200-1223011102212302-3131230213321112-3312213223121010-1021322311201320-2130222032331222"></a>

## Direct properties — server_ssl_profile / 332102001310 / 3

<a id="canonical-3123101132120311-1121210230331331-1233123311113011-3322112010310111-0002012322121331-3020203322032132-3113233321330202-1312213011201132"></a>

<a id="canonical-3123233012331311-0030102220332212-2210230310030320-0201300332132321-0131012112123102-3002330302100333-2121013003330010-1132132203200312"></a>

## kind property — server_ssl_profile / 332102001310 / 4

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

<a id="canonical-3303131301110201-1021303100323033-0101202321000101-3210132321223212-3200121332322323-1033002303110310-2002211221031022-2132320312200213"></a>

<a id="canonical-1130021220131201-3133113103133100-0102322221101301-1021222220203302-3133013133030111-0100231020200300-1112012233303132-3132222320322320"></a>

## name property — server_ssl_profile / 332102001310 / 5

Type: `"string"`. Optional.

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

<a id="canonical-3330222313300323-1300303111103222-1212310233133110-2310133303300312-0010323332131012-3320202010313020-1311231111223120-0230231231131333"></a>
