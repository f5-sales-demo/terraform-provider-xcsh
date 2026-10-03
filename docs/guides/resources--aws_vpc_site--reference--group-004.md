---
page_title: "xcsh_aws_vpc_site reference"
subcategory: "Infrastructure"
description: "Complete grouped canonical reference for xcsh_aws_vpc_site reference."
---

# xcsh_aws_vpc_site reference

<a id="canonical-3330022030111020-0211303110331232-1132332320020211-1331120113231213-0330111332320221-1223013030211210-0132301332103012-3123130311013123"></a>

## Next pages — ingress_gw / 132021220030 / 5

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210030221100330-2222030211222220-1333231330321012-3310211030101003-3012221013001202-2002220221103113-3002211331133002-2332213113103110"></a>

## ingress_gw.allowed_vip_port — allowed_vip_port / 300203100323 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- ingress_gw.allowed_vip_port

<a id="canonical-3021003322112210-0202103121023320-0020300302212320-0232323021231221-2220311223120001-1332303313000102-2133132301011102-1001023220112003"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133002330300012-2212000312131322-1311020011032111-2203221003022330-1102101013022201-2022031111221021-1102320310111200-2321200202102330"></a>

## Direct properties — allowed_vip_port / 300203100323 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-3030130100031031-3110330103102133-2211221032232332-3121213010313222-1122011110303122-3201121130233003-1310312200330102-0332130013110010): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0021321203021022-2333122100321023-3320321202012002-0031002332222121-2121020013311133-0001010303131012-0013111213213202-3120330313321223): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-3112312013123230-3030102313111030-1231010003212333-1100203123301311-2022021101122220-2000100021230001-2023002113310102-2200010110213320): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-3303103032213212-1313022202301222-1122121120120233-0012332330321030-3110313202202033-3000102032111131-3220030131010200-2320033302200210): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-0322111322012103-3111303311310212-3031112000333012-3302030220211012-0310020300300120-0101133130232301-1220020213331111-2301321002132331): complete subsection reference.

<a id="canonical-1121030023201111-0200112130032300-3233301031101132-1113222331121331-0132000102223100-3321312102322200-3211331332300100-0210303103123330"></a>

## Next pages — allowed_vip_port / 300203100323 / 4

- [ingress_gw.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-3030130100031031-3110330103102133-2211221032232332-3121213010313222-1122011110303122-3201121130233003-1310312200330102-0332130013110010)
- [ingress_gw.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0021321203021022-2333122100321023-3320321202012002-0031002332222121-2121020013311133-0001010303131012-0013111213213202-3120330313321223)
- [ingress_gw.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-3112312013123230-3030102313111030-1231010003212333-1100203123301311-2022021101122220-2000100021230001-2023002113310102-2200010110213320)
- [ingress_gw.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-3303103032213212-1313022202301222-1122121120120233-0012332330321030-3110313202202033-3000102032111131-3220030131010200-2320033302200210)
- [ingress_gw.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-0322111322012103-3111303311310212-3031112000333012-3302030220211012-0310020300300120-0101133130232301-1220020213331111-2301321002132331)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3030130100031031-3110330103102133-2211221032232332-3121213010313222-1122011110303122-3201121130233003-1310312200330102-0332130013110010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020022110131122-0220312103320131-3011003232230331-1111301001301101-0012310010012131-0213111132113211-1222033301310231-2010311203213320"></a>

## ingress_gw.allowed_vip_port.custom_ports — custom_ports / 311110212231 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- ingress_gw.allowed_vip_port.custom_ports

<a id="canonical-0231222022313320-3001322101013221-2133021133011331-2331233313120130-0133122312220310-1220030023223311-0203230303321221-1231232033121210"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2310030333231310-3222220012322130-1300202131311220-0021212203313302-3100322020020103-0102121000002222-2003012213211321-0002112230133300"></a>

## Direct properties — custom_ports / 311110212231 / 3

<a id="canonical-0030221222012223-2323322033223131-2210020330210313-3000011333202022-1012031102131021-0333320033313210-3020002303013113-2001001131012002"></a>

<a id="canonical-1301111112313333-3121201032123313-0023111022022011-0111313123312232-3333220013030130-3122211020322110-1121313222322201-2113230232201110"></a>

## port_ranges property — custom_ports / 311110212231 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-1001230121011110-3331103012122100-1002233100103121-2303110300131301-1100230310200010-2022001010303212-1003221203333223-1233023013123132"></a>

## Next pages — custom_ports / 311110212231 / 5

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0021321203021022-2333122100321023-3320321202012002-0031002332222121-2121020013311133-0001010303131012-0013111213213202-3120330313321223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123223232310101-1010010012001021-0100302303332031-0313112311021011-1220131311111303-3131320120322331-3312203032221202-1301233233003233"></a>

## ingress_gw.allowed_vip_port.disable_allowed_vip_port — disable_allowed_vip_port / 022101002101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- ingress_gw.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-0033221013102302-2323202333303221-2123211322133132-1021321220012310-3312211002323322-3201332212130010-1013333313100220-0100022332002012"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-0311113313220302-3020211220033322-0321133020331131-2210232331210331-3113201310233213-1032102101230003-0223210012201123-2222132000312121"></a>

## Direct properties — disable_allowed_vip_port / 022101002101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313011120033301-0201210203002010-2111122102010103-1303312121110023-3101230321321221-1313312002103013-0112212100223221-2300300303102000"></a>

## Next pages — disable_allowed_vip_port / 022101002101 / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3112312013123230-3030102313111030-1231010003212333-1100203123301311-2022021101122220-2000100021230001-2023002113310102-2200010110213320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102030200110301-3101303001000200-1131002130130310-1221213220330101-1232200000321022-0101313030201201-1321101031030213-1122233103130313"></a>

## ingress_gw.allowed_vip_port.use_http_https_port — use_http_https_port / 102000202130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- ingress_gw.allowed_vip_port.use_http_https_port

<a id="canonical-3313132220133323-2312312013333203-1211323133120300-0201330320103130-0310220220012032-2202112322121311-0120233333023133-1021110031032233"></a>

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
use_http_https_port = {}
```

<a id="canonical-3032113002032101-2122303101332132-2130310003120313-0011020023222220-2212203011200333-1020201012322113-0321103300323131-0020333332120212"></a>

## Direct properties — use_http_https_port / 102000202130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203311113011012-2113232310332123-3312300123133020-1230231313332202-1003301032213102-3330322203311111-1213223003222303-1323333213303003"></a>

## Next pages — use_http_https_port / 102000202130 / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3303103032213212-1313022202301222-1122121120120233-0012332330321030-3110313202202033-3000102032111131-3220030131010200-2320033302200210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212303201233113-2000130321311103-2131203303112001-0122303221022211-0011003010233030-1113112320310332-0322212320120011-3132111312330222"></a>

## ingress_gw.allowed_vip_port.use_http_port — use_http_port / 000320321100 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- ingress_gw.allowed_vip_port.use_http_port

<a id="canonical-2023330210030203-2121310211230221-1223202132221021-2032232032121323-0301301022022101-1313003211013310-2332123023022033-2033312020103200"></a>

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
use_http_port = {}
```

<a id="canonical-1331030203111030-0021112211011330-2033301112333032-0212202232112023-1311031212113321-0131200112211021-2332221213120310-0120112211320032"></a>

## Direct properties — use_http_port / 000320321100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2321320200221303-3002110320303012-3210110201113112-1022101031022132-2303031202130221-0313230303001220-1120010231110002-1301132111013213"></a>

## Next pages — use_http_port / 000320321100 / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0322111322012103-3111303311310212-3031112000333012-3302030220211012-0310020300300120-0101133130232301-1220020213331111-2301321002132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103120221101002-0311111013102031-0023330123301302-3020301313211000-2102010032311330-2102211213022020-1220302102221333-0022023332211320"></a>

## ingress_gw.allowed_vip_port.use_https_port — use_https_port / 102000222122 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- ingress_gw.allowed_vip_port.use_https_port

<a id="canonical-2310010113002121-2200012112303100-1232232230311301-3330031301102200-0301101322111230-1200232202203032-3222130031030202-0213233130230130"></a>

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
use_https_port = {}
```

<a id="canonical-2201111310100230-1213212103322323-0230232312203202-3021021202310110-1311122033031120-0132111303113310-0322330323222113-2030203223231112"></a>

## Direct properties — use_https_port / 102000222122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331330013111023-2102011031300232-3321032011323232-0033202130131002-3323301303012013-2102022131323321-2320013012300200-2010233330123032"></a>

## Next pages — use_https_port / 102000222122 / 4

- [ingress_gw.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2202312332212100-1031132303122122-1232023011302122-2222323223002320-2301112220203003-1311230220303301-1223300032132002-0201022012133120)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0111213230132322-1123100002200232-1003030323321130-1232012122220302-0113010030221312-0212000103221311-0000212213323202-2231223130300123"></a>

## ingress_gw.az_nodes — az_nodes / 002003332032 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- ingress_gw.az_nodes

<a id="canonical-2020303030020202-0123132102230322-2002003002120201-2301021023332233-3220322333301210-2100012332111133-2333120022222301-1321010112101302"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name")}
```

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
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023322333111001-2320322211133022-0121133330123022-3032132112303033-2213201010123232-1111012122011202-2130012123001323-2012023221101023"></a>

## Direct properties — az_nodes / 002003332032 / 3

<a id="canonical-3312202212100103-0100332331231020-1320111311222200-0002201111203133-1023121123200311-1331232032010230-0302323131312020-0333033232000302"></a>

<a id="canonical-0303120023332133-3333022021002203-0110333022311130-3000313112211320-0021102122210111-2332121301122113-1300221333202333-3101100002023101"></a>

## aws_az_name property — az_nodes / 002003332032 / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

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

- [local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1302233213031211-1212000333232000-2203231000222323-1231321023123321-2302301001101212-0213111221332212-2200032023323003-3321203120300110): complete subsection reference.

<a id="canonical-1012210223030102-0331302313022323-3323231300301112-0003321332103122-1111230311031023-3132210330122012-3100013033311113-1021331021303223"></a>

## Next pages — az_nodes / 002003332032 / 5

- [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1302233213031211-1212000333232000-2203231000222323-1231321023123321-2302301001101212-0213111221332212-2200032023323003-3321203120300110)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1302233213031211-1212000333232000-2203231000222323-1231321023123321-2302301001101212-0213111221332212-2200032023323003-3321203120300110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233332202311223-1032001223312313-0301013332202211-1103121111022100-1231320010110100-1000031121102312-1012323322213012-2330331301131223"></a>

## ingress_gw.az_nodes.local_subnet — local_subnet / 032102221011 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210)
- ingress_gw.az_nodes.local_subnet

<a id="canonical-2303321002201331-1000001003002002-3011310312121331-3002310000030201-3101320321220023-1211232222302311-1110231102332323-0012210002300200"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3311010133033300-3000023022031021-0322001233000132-0322113110202032-0011103100120100-3223223022132123-2203032113112133-0211002201231322"></a>

## Direct properties — local_subnet / 032102221011 / 3

<a id="canonical-3031210300322212-1222203230101121-0220033321332010-0101110033022113-3312230112101102-1323102012120201-0132221232232302-0303010132123311"></a>

<a id="canonical-2322321320130012-0021200100201310-0233331311210112-2233301002110230-1312300030222111-3322102102120123-2201113233123111-1220101310221200"></a>

## existing_subnet_id property — local_subnet / 032102221011 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-2201002020030333-2230212300313123-2023300013112031-2123312003230023-2232131012011312-1301301233211130-3033223021013122-2002132303133121): complete subsection reference.

<a id="canonical-0332030333332311-3232031202110031-1003210312323121-2231021232012230-3333321032301011-1033333200200303-2233133110023322-0130220300012133"></a>

## Next pages — local_subnet / 032102221011 / 5

- [ingress_gw.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-2201002020030333-2230212300313123-2023300013112031-2123312003230023-2232131012011312-1301301233211130-3033223021013122-2002132303133121)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2201002020030333-2230212300313123-2023300013112031-2123312003230023-2232131012011312-1301301233211130-3033223021013122-2002132303133121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020320201200133-3120033212200101-2012133332133222-0203220210133311-2222023230030302-1023222333103310-1203320113121301-3331211100110223"></a>

## ingress_gw.az_nodes.local_subnet.subnet_param — subnet_param / 221021021101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-2330123030021300-0333222331002031-2122221322133023-1323322003320020-1210003321202232-0102331330021201-0022103012213020-0103331133211210)
- [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1302233213031211-1212000333232000-2203231000222323-1231321023123321-2302301001101212-0213111221332212-2200032023323003-3321203120300110)
- ingress_gw.az_nodes.local_subnet.subnet_param

<a id="canonical-2133230210023102-1112012233010110-3213013231223000-3200310333122213-1331233001231020-1033023021223202-0010232000111213-1000220312312212"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-2022230212333222-2222300210200123-2313331120022221-2220133033322113-1232031230112313-1022232211123213-3312002132002023-3323110112021230"></a>

## Direct properties — subnet_param / 221021021101 / 3

<a id="canonical-2122120030230310-0113330112121332-3201303101022232-3332012103123200-2021113130310131-0110332202312132-0313120202233123-3222130001030320"></a>

<a id="canonical-1232310330010123-2002231201321022-3312000210300331-0311211000323000-1021322032330133-2010232021311001-0223100333330113-3231332321101213"></a>

## IPv4 property — subnet_param / 221021021101 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-0202311323202303-1331302023230220-2102222001300211-1023120121132300-3113023203102213-0030222123223103-2300203131130100-2232210001312210"></a>

## Next pages — subnet_param / 221021021101 / 5

- [ingress_gw.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-1302233213031211-1212000333232000-2203231000222323-1231321023123321-2302301001101212-0213111221332212-2200032023323003-3321203120300110)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132302122330230-2002020221232103-2033201303213133-3233132310002013-3020133001113110-2030122102231000-2230112322113312-2201332330032021"></a>

## ingress_gw.performance_enhancement_mode — performance_enhancement_mode / 213121210102 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- ingress_gw.performance_enhancement_mode

<a id="canonical-0002123232300200-1312212131002311-2100031001332320-3201100111203201-3332021031333101-0022011013231210-2102200030301322-2131033232322200"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213110223222301-0211223003000032-0313123320120022-3302211102011002-3030230321320130-1030002102333113-2230332100233230-1103110022210103"></a>

## Direct properties — performance_enhancement_mode / 213121210102 / 3

- [perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010): complete subsection reference.

<a id="canonical-0013332302010032-1222121311013113-2321322311101322-0032030221020212-2203312212211103-2010001013320301-2131131023110221-1230121112023132"></a>

## Next pages — performance_enhancement_mode / 213121210102 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233110321301123-1011100031102120-3313022113103130-2333001123302133-2003213020331001-0000232130232311-1301031310210302-1200100231130302"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 301133221133 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-1220023032023012-3030212213313303-0111201100320021-3123220010311013-1111332302211331-3132120321121133-0202020221310131-3200101313002010"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020123300010023-3300003120110301-0030331233322110-1232333302013210-2110131200002022-3230102332112321-2031012130123030-3132231221202322"></a>

## Direct properties — perf_mode_l3_enhanced / 301133221133 / 3

- [jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-0301203122200213-1303331222102120-2131113231032201-2013033110103132-1323231301320123-1233303322211022-2320220223322323-2030130020011100): complete subsection reference.

- [no_jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-2233312202133211-2112021232203201-2221001011232221-0113021112220020-2222323232002202-2320312032321231-1220231201131221-1132312232312020): complete subsection reference.

<a id="canonical-1101003313012313-0031021200331021-3001123002210213-3012100201222133-3201003123230110-1103320121033211-0300231301223103-2320001313032221"></a>

## Next pages — perf_mode_l3_enhanced / 301133221133 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-0301203122200213-1303331222102120-2131113231032201-2013033110103132-1323231301320123-1233303322211022-2320220223322323-2030130020011100)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--aws_vpc_site--reference--group-004.md#canonical-2233312202133211-2112021232203201-2221001011232221-0113021112220020-2222323232002202-2320312032321231-1220231201131221-1132312232312020)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0301203122200213-1303331222102120-2131113231032201-2013033110103132-1323231301320123-1233303322211022-2320220223322323-2030130020011100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021133311311111-3102011112033003-3122110133311303-3000232222331122-1022302200113003-0123213000220232-3231020021003000-0032301201321332"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 121100130121 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-0310233322110300-3002101000102021-2202011322322313-3303112201131311-2330123132020333-3211311010102003-0123222323012000-0320121300122012"></a>

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
jumbo = {}
```

<a id="canonical-1112132203013100-0310211221330220-0000312020022332-3332313113001122-0131201131101201-1202202023201212-1312101123112213-3210031302213102"></a>

## Direct properties — jumbo / 121100130121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301123221103121-3023200230011303-3301113122311122-3033013223113103-0221021301210321-0130100330233210-3203012320121132-3330230111001213"></a>

## Next pages — jumbo / 121100130121 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2233312202133211-2112021232203201-2221001011232221-0113021112220020-2222323232002202-2320312032321231-1220231201131221-1132312232312020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333303303221301-3123332013302021-0013321123200231-0312000311213233-3113003000001320-0200130313122000-1320032001332213-2120222133211233"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 213132330020 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2331200023332332-0211110030122211-0102033102122122-1112003203131222-1230120111321001-0232121113213102-2033332021221031-2322031320032203"></a>

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
no_jumbo = {}
```

<a id="canonical-2230120030010023-1332202100231303-1303323203032013-1333100312200302-2301002300203303-2020303012111023-0013020202132112-2223002100100223"></a>

## Direct properties — no_jumbo / 213132330020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100033211121312-0111233331102022-3233123123232021-2031212022200033-0030032231330030-1031121010120121-1303321223110121-0233113300311023"></a>

## Next pages — no_jumbo / 213132330020 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-1102213232230002-2032330323032331-2201210133131120-3302030210312330-0131113122332110-0023210311133002-3113113202203000-2330003200103333)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113333323212322-1022032120330030-3203032311300131-2032322023320303-2232000122012003-1133011233320122-3323123033203001-2133123220303221"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 230102033110 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-2132031013320230-1213330322323312-3010101333120300-1312212312013310-3322000203122012-3023123331210213-2022302111113220-2333212000113011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322213011321203-0131300000103330-3323222101012213-3330321011020111-2032023001303101-0033200001022130-3202121003311321-2102223322210032"></a>

## Direct properties — perf_mode_l7_enhanced / 230102033110 / 3

- [jumbo_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-0221103230320302-1323200101032331-3030230320011133-2000111213123130-2322123102221023-1200220111212300-2200231101111110-3301133223023121): complete subsection reference.

- [jumbo_enabled](resources--aws_vpc_site--reference--group-004.md#canonical-1100332022312120-2212100023201002-2321112222132031-1110302323222103-2032332232030120-1003120320112233-1112213123312223-2123210020323012): complete subsection reference.

<a id="canonical-0332321232032100-3113012211221211-2203232112103123-0211302012001303-2310233303132321-3123320113002130-2031123111112100-1310212010020112"></a>

## Next pages — perf_mode_l7_enhanced / 230102033110 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-0221103230320302-1323200101032331-3030230320011133-2000111213123130-2322123102221023-1200220111212300-2200231101111110-3301133223023121)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--aws_vpc_site--reference--group-004.md#canonical-1100332022312120-2212100023201002-2321112222132031-1110302323222103-2032332232030120-1003120320112233-1112213123312223-2123210020323012)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0221103230320302-1323200101032331-3030230320011133-2000111213123130-2322123102221023-1200220111212300-2200231101111110-3301133223023121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020023133031101-2332312302311030-0102233020003113-2221003212221030-1101022022002023-3021213131232211-2200112133310132-0023202132001312"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 332030202033 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3131023312330111-1100000010201220-3123012203021221-1220010122231130-3112120313201331-3210330303220101-2001013021230302-0221300201221303"></a>

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
jumbo_disabled = {}
```

<a id="canonical-1030013002313120-3223013212121130-0101301213300023-3211300223300011-0111332201121002-1111120320321030-1021130302133111-2223221322101200"></a>

## Direct properties — jumbo_disabled / 332030202033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133332230231311-2311031030212303-3232001003311101-3320310312102310-0132203013330023-1320012013313311-2002213003000122-2123301032321132"></a>

## Next pages — jumbo_disabled / 332030202033 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1100332022312120-2212100023201002-2321112222132031-1110302323222103-2032332232030120-1003120320112233-1112213123312223-2123210020323012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102111131121121-1311102221132331-3331221333231331-3200211130323102-3200310200221112-1202130311232012-0211231011303101-0002310212233202"></a>

## ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 212232220111 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [ingress_gw](resources--aws_vpc_site--reference--group-003.md#canonical-0333332022003321-0130333001220102-3300331112302201-3222032002202120-2012011102331101-3213012110100103-2022011031100003-2201121222303323)
- [ingress_gw.performance_enhancement_mode](resources--aws_vpc_site--reference--group-004.md#canonical-3113111133121313-2001031200031030-3223100231113023-1120200011200033-1112201030331232-0320220033031131-0201132212302222-3102332030132220)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-3202130023332213-3203201333233122-3300300100020322-3302012303122230-2102202312000122-2322231211202123-1130332102003131-2232201100211300"></a>

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
jumbo_enabled = {}
```

<a id="canonical-0211131232101222-1220212300212210-1210203001113102-0303020310322321-1131133212231002-2202313202303311-3030020223010013-3330013331311332"></a>

## Direct properties — jumbo_enabled / 212232220111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003202213122202-1332133012333030-0033022320210230-0131203010332013-0003311221000020-3300222031211000-3332311130122213-0101021300321331"></a>

## Next pages — jumbo_enabled / 212232220111 / 4

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_vpc_site--reference--group-004.md#canonical-2103000323002023-0233333210111200-1133120013122221-2323301000331303-1320313231131002-3213211121221221-2303120012202303-0202001302013010)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233302112012003-1332013301100201-1333302101012212-2220003013121022-1330302311132223-2033220112213322-1213033002001033-2121330112223120"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 010111212002 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- kubernetes_upgrade_drain

<a id="canonical-2323233201000331-1113320001333220-2031013132230301-3132223121102100-0330233010300233-0333111313130310-3202020010020123-3302123301012121"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302333000010121-1022331223230300-0310233133300222-1101022112201221-1021301133232322-0322223010110211-0012021202102032-0123121021200212"></a>

## Direct properties — kubernetes_upgrade_drain / 010111212002 / 3

- [disable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1133031301132003-3010203023030122-2121300323102110-2032303101023013-1311111201000120-3211020111321120-0121333020122223-1220113202103210): complete subsection reference.

- [enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213): complete subsection reference.

<a id="canonical-2123213231000121-2110033320310223-2323313002232231-2323221312220132-0201023200330322-2313211201300232-3333103133022310-3132332103321112"></a>

## Next pages — kubernetes_upgrade_drain / 010111212002 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1133031301132003-3010203023030122-2121300323102110-2032303101023013-1311111201000120-3211020111321120-0121333020122223-1220113202103210)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1133031301132003-3010203023030122-2121300323102110-2032303101023013-1311111201000120-3211020111321120-0121333020122223-1220113202103210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222133031121003-3121101200012303-0103003313120002-3030003322113003-1010310132323010-3030312311330313-0031201111133112-2001133022110332"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 331320203032 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2323122121323231-2302211031220333-2203121013110321-3002223321313120-0320320213331121-3222213203221030-0102233230303331-3003101000131312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

<a id="canonical-3303213110201121-3111012132113203-1221301031031310-3302101322233123-1113121221230321-1011020222113332-2031321310120213-1320132003031232"></a>

## Direct properties — disable_upgrade_drain / 331320203032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022001222111212-2230332132330022-0303022203210130-3030232011232302-1230301223100101-1220031131130213-1303311213020323-0011202332312201"></a>

## Next pages — disable_upgrade_drain / 331320203032 / 4

- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023233102323213-3222300231210122-1323310121111102-1133001310331302-2010222311032233-1303310100030213-0332011212221000-0031321221030300"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 123000313310 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-3333013332001101-3220020012112110-0011012231011020-1001001222003321-3222203110031232-0202201331112110-1203002021121311-0222003101212030"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("drain_node_timeout"),
  validators.ConflictingObjectAttributes("disable_vega_upgrade_mode",
    "enable_vega_upgrade_mode"),
  validators.ConflictingObjectAttributes("drain_max_unavailable_node_count",
    "drain_max_unavailable_node_percentage")}
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
  "x-ves-oneof-field-drain_max_unavailable_choice": "[\"drain_max_unavailable_node_count\", \"drain_max_unavailable_node_percentage\"]",
  "x-ves-oneof-field-vega_upgrade_mode_toggle_choice": "[\"disable_vega_upgrade_mode\",\"enable_vega_upgrade_mode\"]"
}
```

Terraform syntax:

```terraform
enable_upgrade_drain {
  # Configure direct properties listed below.
}
```

<a id="canonical-3200003122223103-3223202212013132-1332213311023132-1212110322231223-1111032033313322-2101300331312030-1321121010102023-2200012022323333"></a>

## Direct properties — enable_upgrade_drain / 123000313310 / 3

- [disable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-1232313003030222-3313212222002232-0323021231201122-3211133210012022-0303023121023002-3322120332023001-1301032121312210-3001332201321220): complete subsection reference.

<a id="canonical-0303121222001030-3033100000120212-2223323213200300-2221032230120332-0213102013233311-2003313101232212-2003033331031023-0302211120313302"></a>

<a id="canonical-2113301032312310-1303220020202101-3003312121132331-3330120011301311-1302120103202130-2233101130302223-3300322102323310-1122213123020031"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 123000313310 / 4

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Upstream description:

Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 5000),
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
    "minimum": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "5000"
  }
}
```

<a id="canonical-3320110011203230-3210303121233103-2012131122020030-2200010103002112-3101332110101332-1322132201220002-1220000323213113-1302133213110020"></a>

<a id="canonical-3120012221231133-3022103212203131-1020113132332001-3232012233232131-3310212311231010-0133113032222130-0130022022133300-3133230301131001"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 123000313310 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-3020120111210331-0332312210323203-1021213011100312-3100022330131311-3331010013110203-3200223133301201-3031211021103102-0100223300200330"></a>

<a id="canonical-0200211212330332-3223132210221011-0030003222331033-0103101301133302-3213001132001100-3103331311233312-3111102232221333-0202233021013232"></a>

## drain_node_timeout property — enable_upgrade_drain / 123000313310 / 6

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It
is..

Upstream description:

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(0, 900),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 900,
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2131020210233030-1231320302312030-3202321233013313-1322303013201303-1303100332032010-0030301020013000-2013012002033322-2002321113022020): complete subsection reference.

<a id="canonical-2220033311103110-3011022211221003-0013111022103101-3131311330100111-0000013132202001-0022200013330303-2102132332110023-3210213312012123"></a>

## Next pages — enable_upgrade_drain / 123000313310 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-1232313003030222-3313212222002232-0323021231201122-3211133210012022-0303023121023002-3322120332023001-1301032121312210-3001332201321220)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2131020210233030-1231320302312030-3202321233013313-1322303013201303-1303100332032010-0030301020013000-2013012002033322-2002321113022020)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1232313003030222-3313212222002232-0323021231201122-3211133210012022-0303023121023002-3322120332023001-1301032121312210-3001332201321220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2000230302120113-1020232301130110-1331002321231201-2110202311313123-1321130313232211-2032322112020313-1300312000123311-3213002022120203"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 123303233330 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-0032130310032231-1011131131010010-0122231201000322-3111333220233332-0131111230122103-0300021330120302-3301321211000003-1232103313332200"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

<a id="canonical-2210131221113212-0020013100000300-0112103123302001-1033330332030332-2313123101102103-0200313013131322-3122320130130200-3301122303223323"></a>

## Direct properties — disable_vega_upgrade_mode / 123303233330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1101110200202022-2132013023221312-1333011203013333-3303332022032303-0012220112020313-2013312231232210-3202132003323313-0233130323233120"></a>

## Next pages — disable_vega_upgrade_mode / 123303233330 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2131020210233030-1231320302312030-3202321233013313-1322303013201303-1303100332032010-0030301020013000-2013012002033322-2002321113022020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120010333122220-2221210031230331-3213330011221320-2321001031302200-0032320302212201-1001133031102131-1020003212330222-0032021001021020"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 023312122010 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [kubernetes_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-3013301332331333-2211212210321032-3013100011321103-0221022231110303-0303221221101223-1102301300233020-0322322133203301-3113001020112122)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-3120110303302200-2001111223332112-1111330300002331-3201302311022113-2332201023301120-0320232201313332-2013031023103033-0312132102212133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

<a id="canonical-2233311111013230-3023013001201213-3122302110230102-1202000010021103-3312303023210320-1100120023023233-1222033003222101-2123222212331030"></a>

## Direct properties — enable_vega_upgrade_mode / 023312122010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021011210331123-0203030301221313-3213323131102130-2120310112231200-2000011112221133-2300031102203011-3313113033301301-1102231112030121"></a>

## Next pages — enable_vega_upgrade_mode / 023312122010 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_vpc_site--reference--group-004.md#canonical-1100112131223303-0112212100023033-1012332302313333-3121013233310303-1121100322021130-0331000302011123-1233103311232032-0330003032123213)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3220233101332201-0221100213122133-2100201232213330-1302330031123033-1031032112102021-3331322021100121-1222111221211231-3331002202001201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031213310202220-1312122120232110-2033321121013233-3011120031202223-0102030011002132-1330122011020200-2103111000123311-1001303212011202"></a>

## log_receiver — log_receiver / 112303332032 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- log_receiver

<a id="canonical-2003213030032201-0000122320031311-3313011233322130-0100210100222012-1130101021232210-2300012013022331-3332320223010333-1031330223022230"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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

OneOf alternatives in this subsection:

- [log_receiver](resources--aws_vpc_site--reference--group-004.md#canonical-2003213030032201-0000122320031311-3313011233322130-0100210100222012-1130101021232210-2300012013022331-3332320223010333-1031330223022230)
- [logs_streaming_disabled](resources--aws_vpc_site--reference--group-004.md#canonical-2002222101023311-2002312201100012-2211123211220012-0110012100322233-0322231103202303-2230110301300200-0023321323023233-2213312131012103)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-1131303020030032-1013300110223132-3121333102323320-2113002012310302-3103311130121100-2320230302231210-3212313033002203-3212032010331322"></a>

## Direct properties — log_receiver / 112303332032 / 3

<a id="canonical-0333301110103030-1300100213013101-2131001212321332-0202033023120331-1300033220022233-1111102013120232-2120003311132120-2323112302031201"></a>

<a id="canonical-1012313233101111-2301003033302112-3021332132211233-3010120322031023-2001202011102113-2112033321332302-1233133221021130-2013201303010311"></a>

## name property — log_receiver / 112303332032 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1023231133123331-2221222212021130-0111112320030012-1033212331010200-2220333331013300-3311133302301130-3020022011231312-3223320120103121"></a>

<a id="canonical-2021001311333223-1211013330110330-2311020002003112-1200233210210333-2231220330012021-2221201322012201-2230133112333032-0113321313233030"></a>

## namespace property — log_receiver / 112303332032 / 5

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

<a id="canonical-3332300312231233-1102220300233231-3300021312001011-3023132011030000-2323330201301200-0202211101202331-1220200220123110-3132220101033111"></a>

<a id="canonical-1221012023100131-0310123202301121-1032031000210331-1020023131003310-3030032113012313-1111102032322222-0033302033230120-3322030330300011"></a>

## tenant property — log_receiver / 112303332032 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0122020312102002-2122123300200010-1321000000021103-2100033130110223-3001113301001230-1012130220131000-2120121213301211-3000023030103332"></a>

## Next pages — log_receiver / 112303332032 / 7

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2032020211212221-2302310202003310-0312213011010303-2213221031020102-0022211032123131-0011101032031322-0302123012333120-1132100112302201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122110133310012-2112332013013223-3311313223230020-2121020231002210-1221123031211302-0211220322301212-0012322133000332-1320113301100203"></a>

## logs_streaming_disabled — logs_streaming_disabled / 111123103123 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- logs_streaming_disabled

<a id="canonical-2002222101023311-2002312201100012-2211123211220012-0110012100322233-0322231103202303-2230110301300200-0023321323023233-2213312131012103"></a>

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
logs_streaming_disabled = {}
```

<a id="canonical-2120210013333333-2223203023333032-1111031211203100-2000200312131302-1220020211011121-3301023031201021-1022201000233033-3311222100103300"></a>

## Direct properties — logs_streaming_disabled / 111123103123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322113223031113-1332310130203213-1321130101310313-1312313001330013-3122033112322011-0300301033323213-0203130302120201-0032122131133033"></a>

## Next pages — logs_streaming_disabled / 111123103123 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2332333023332232-3133011010233103-0013221302323103-2213132121331123-0331101321320212-1032203101301301-2232301020321122-2200320322313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332121312323023-2220033222211313-1300032131330303-1013013223210101-2121122100211033-2003330103003120-1321321002221233-3302030323310311"></a>

## manual_routing — manual_routing / 111310203110 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- manual_routing

<a id="canonical-3020121031032231-1103013030230103-3031133330101100-2301130023123222-1233013021101120-3103232331131332-0022200132303212-1113101010212212"></a>

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
manual_routing = {}
```

<a id="canonical-2300322200031321-2212000132302002-2011203012202212-1032011231020133-2201031110331111-2121010213123021-1133231011200010-1333113223111122"></a>

## Direct properties — manual_routing / 111310203110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0220113020311031-0010232021120231-0112322030323001-2201303300120113-3202030001233011-0002232302222001-2103211030211322-1323320332011132"></a>

## Next pages — manual_routing / 111310203110 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0323111213120210-3330110331133331-1123320133233101-0112302203111100-3323311230111302-2322033010012210-0303221110112120-1110302200110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213020020000221-3322133013202010-2032200033111311-3033022120111230-1302123011212121-0200211320210203-1111233300102212-2302023222210112"></a>

## no_worker_nodes — no_worker_nodes / 223322112000 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- no_worker_nodes

<a id="canonical-0100330200230123-0211130123221033-1210312202010201-0333000212310210-2100131231233213-2113223121221203-2010213310032330-2231102112200001"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_worker\_nodes, nodes\_per\_az, total\_nodes; Default: no\_worker\_nodes\] Configuration
parameter for no worker nodes.

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

OneOf alternatives in this subsection:

- [no_worker_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-0100330200230123-0211130123221033-1210312202010201-0333000212310210-2100131231233213-2113223121221203-2010213310032330-2231102112200001)
- [nodes_per_az](resources--aws_vpc_site--reference--group-001.md#canonical-3232210103132020-3020113020222012-3122131000321021-1031203301330033-2020130333332332-2213031031300311-0132132001223132-0013031200000010)
- [total_nodes](resources--aws_vpc_site--reference--group-001.md#canonical-1332013110110103-0030221222310310-2233102023323102-0331312112001230-2223011133323200-1031133321333233-2230331022333211-0320212023101300)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_worker_nodes = {}
```

<a id="canonical-1332230311231110-2311120231200302-3321023103131101-0201103210110123-1231020211103213-1032323133233320-1303330233333131-1032102222332121"></a>

## Direct properties — no_worker_nodes / 223322112000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2020013320102332-0203023232033123-0030020120122010-0200131233003313-2101123101301002-2110232120303310-3002201322100221-1233010212313013"></a>

## Next pages — no_worker_nodes / 223322112000 / 4

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122013102131120-0202312202011033-0002110231022321-1202302232230301-1002230223300032-1121130320032011-2032202110020110-0322111320320322"></a>

## offline_survivability_mode — offline_survivability_mode / 112101011131 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- offline_survivability_mode

<a id="canonical-1303201030303311-0220223111201320-1001213033122111-3331321022302321-0002002200100133-0000223102013323-1100133322130311-1100010331230000"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("enable_offline_survivability_mode",
    "no_offline_survivability_mode")}
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
  "x-ves-oneof-field-offline_survivability_mode_choice": "[\"enable_offline_survivability_mode\",\"no_offline_survivability_mode\"]"
}
```

Terraform syntax:

```terraform
offline_survivability_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120222322132000-0202232322001123-1021232022313120-3032201012101200-3023123120002321-0310223330202002-1030320102320002-3011223303103000"></a>

## Direct properties — offline_survivability_mode / 112101011131 / 3

- [enable_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0033033030202011-1022232000302212-3111300103300011-3133222000230321-2132310111020101-1330332323210131-2131231003003211-3331233313321023): complete subsection reference.

- [no_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-1311220120232231-0323132033232101-1010203132313133-0110003121110333-3100133313213212-2322001203103021-2201013010223112-1233311033330320): complete subsection reference.

<a id="canonical-3322322002021212-1102300130322320-0321003333201202-2113123102120132-0203230201013312-2022220203033301-2122113003123333-1131201122121023"></a>

## Next pages — offline_survivability_mode / 112101011131 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-0033033030202011-1022232000302212-3111300103300011-3133222000230321-2132310111020101-1330332323210131-2131231003003211-3331233313321023)
- [offline_survivability_mode.no_offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-1311220120232231-0323132033232101-1010203132313133-0110003121110333-3100133313213212-2322001203103021-2201013010223112-1233311033330320)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0033033030202011-1022232000302212-3111300103300011-3133222000230321-2132310111020101-1330332323210131-2131231003003211-3331233313321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211231010332211-2133132323013320-3110213133032321-2202313330020232-0233110332312311-2210132123310223-0111023302203232-2233130223211300"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 300231222210 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-0023302230310002-0203332110132210-1322031001210032-2123102311212103-0101010112002223-2333120002200022-0203120331120103-0211130322313202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-1231313132030330-1211232123033333-1320111021230213-0222020101203111-0302321020003300-2210123332033330-0301002221020130-0232111202102031"></a>

## Direct properties — enable_offline_survivability_mode / 300231222210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011303112232332-2013013123002001-3010332032303301-1203100302312220-2123101302321030-2111221221023320-3003112322303012-3001111213032300"></a>

## Next pages — enable_offline_survivability_mode / 300231222210 / 4

- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1311220120232231-0323132033232101-1010203132313133-0110003121110333-3100133313213212-2322001203103021-2201013010223112-1233311033330320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130013030321002-1223021213312221-1233033001010203-2321323200211213-0211112210120110-2332122001011203-1332112033211213-3102122030313102"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 212103301301 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-0013000122232111-3213212311201302-3123020313132121-3133022302321030-0311333130031030-0231033030301012-1122331201233002-2230332103102231"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-0203001221220103-1012013332222032-2313031330310331-3223231130202012-1301112011102023-3130333320321020-3301230002330112-3303301133000223"></a>

## Direct properties — no_offline_survivability_mode / 212103301301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110311323111001-1231010202202222-1122333203023210-0330002221303222-3101023320302200-2213033110313232-1221211112333213-3000330300001210"></a>

## Next pages — no_offline_survivability_mode / 212103301301 / 4

- [offline_survivability_mode](resources--aws_vpc_site--reference--group-004.md#canonical-2230331302030233-1231312011331200-0330011122330231-1003002230012012-3322102030022203-0101330012200101-1001211032031010-0121023120133132)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3121332323311033-1213202021203320-2022333333100031-2121201303231310-0210311231023010-0313032021023311-3331022231220210-2100330111200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323012001301003-2023022321110322-1231033332010203-3230130320332332-2012121300022230-0010011320233300-1031333311011021-3330303020310123"></a>

## os — os / 102331213200 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- os

<a id="canonical-1211113230123003-2211300003302201-0200223302302020-3122123101012232-1101132332203223-0132213003031002-2332103112101303-0300002233221120"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203300030312202-2110230130232310-2030113203002030-2121232023322233-1201201211031201-3223230303312233-0201123212111111-1213013122310011"></a>

## Direct properties — os / 102331213200 / 3

- [default_os_version](resources--aws_vpc_site--reference--group-004.md#canonical-0333112203033301-0303011133121312-3313033312311031-3210130331210222-3232020333101111-1000320211132211-0230321030001110-1021221223131012): complete subsection reference.

<a id="canonical-1233102132003121-3300310020011310-1111200213223221-2310103120323133-3112211120303020-0032200101233211-3223303101331033-2212021221302123"></a>

<a id="canonical-0113132202003323-0313103301131231-1303220000013002-1002232030330022-1131202230220020-1130113020233232-3130312333220333-3120022213332301"></a>

## operating_system_version property — os / 102331213200 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2213120200111303-0310000020011203-0232003102333003-0002201313023210-2001101311203231-0302131112031312-0331201132321303-1101133000031113"></a>

## Next pages — os / 102331213200 / 5

- [os.default_os_version](resources--aws_vpc_site--reference--group-004.md#canonical-0333112203033301-0303011133121312-3313033312311031-3210130331210222-3232020333101111-1000320211132211-0230321030001110-1021221223131012)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0333112203033301-0303011133121312-3313033312311031-3210130331210222-3232020333101111-1000320211132211-0230321030001110-1021221223131012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313223303322121-0230130011320201-2233133331313230-3200320031122302-3032022332323021-0223120020300201-3213103313203132-2123002013133221"></a>

## os.default_os_version — default_os_version / 211032212233 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [os](resources--aws_vpc_site--reference--group-004.md#canonical-3121332323311033-1213202021203320-2022333333100031-2121201303231310-0210311231023010-0313032021023311-3331022231220210-2100330111200133)
- os.default_os_version

<a id="canonical-0000022220020313-0000100001111130-1200120233120300-0213312332333122-0223331311111022-3323321003301121-0310203233203323-0000333312333030"></a>

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
default_os_version = {}
```

<a id="canonical-1122030122200133-2230013223023301-0020113113323031-2032022033113012-2201011321301000-2310033020212130-0120012313010103-0213213322131121"></a>

## Direct properties — default_os_version / 211032212233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0233021022112323-2222130231321112-1221110210311030-3312012002203033-3133033020132231-0313200121130021-3011100101020030-3002100122302111"></a>

## Next pages — default_os_version / 211032212233 / 4

- [os](resources--aws_vpc_site--reference--group-004.md#canonical-3121332323311033-1213202021203320-2022333333100031-2121201303231310-0210311231023010-0313032021023311-3331022231220210-2100330111200133)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0310112221323310-2123221110321302-1012131000333231-0300332333303010-3312000320011223-0322212320303112-3202021210020031-0233123301111233"></a>

## private_connectivity — private_connectivity / 223310001211 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- private_connectivity

<a id="canonical-3302002132231203-3030202210010120-1310320012103133-1302011000212230-3111302233111200-0123000123331023-3012122210102203-0012021311213222"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for private connectivity.

Upstream description:

Private Connect Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside",
    "outside")}
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
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

Terraform syntax:

```terraform
private_connectivity {
  # Configure direct properties listed below.
}
```

<a id="canonical-0321020303020203-2122220121200023-2120301023120102-1230010333321221-0001101232211003-2131020333121013-1303031003020203-2221203102230011"></a>

## Direct properties — private_connectivity / 223310001211 / 3

- [cloud_link](resources--aws_vpc_site--reference--group-004.md#canonical-1220213301010112-2031211120200230-2310301003103333-1020121132302232-3133311010020233-1033202210213322-3102022301332202-1103323111000303): complete subsection reference.

- [inside](resources--aws_vpc_site--reference--group-004.md#canonical-0211133132022113-2233221020320212-3203320003011110-3131100111312201-3002330203131030-0223230111220303-0122123022033020-2001230331012031): complete subsection reference.

- [outside](resources--aws_vpc_site--reference--group-004.md#canonical-0021002223313121-1133033102200312-2330110313321233-2023322231000212-3200112321230313-0000020212001221-1331012331201322-1103131102232100): complete subsection reference.

<a id="canonical-3130012112323022-1033233031100323-0103323311120312-3230201232221332-1031312211011002-3010200330130132-3303331100310013-0210021301213200"></a>

## Next pages — private_connectivity / 223310001211 / 4

- [private_connectivity.cloud_link](resources--aws_vpc_site--reference--group-004.md#canonical-1220213301010112-2031211120200230-2310301003103333-1020121132302232-3133311010020233-1033202210213322-3102022301332202-1103323111000303)
- [private_connectivity.inside](resources--aws_vpc_site--reference--group-004.md#canonical-0211133132022113-2233221020320212-3203320003011110-3131100111312201-3002330203131030-0223230111220303-0122123022033020-2001230331012031)
- [private_connectivity.outside](resources--aws_vpc_site--reference--group-004.md#canonical-0021002223313121-1133033102200312-2330110313321233-2023322231000212-3200112321230313-0000020212001221-1331012331201322-1103131102232100)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1220213301010112-2031211120200230-2310301003103333-1020121132302232-3133311010020233-1033202210213322-3102022301332202-1103323111000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021301312310212-3213302211121131-1201223201012013-2102230032113113-1020303020130111-2332100021033003-0020210313031111-1111011223301311"></a>

## private_connectivity.cloud_link — cloud_link / 110311323130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- private_connectivity.cloud_link

<a id="canonical-2312120313020133-0130312103011212-0120320013133101-1330101002232300-2203112210001000-1022210013230131-2310121211111312-2101302230321303"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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
cloud_link {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213331102211012-1011000223011022-2220003113120333-0323330333120020-3002112332100030-2103102320012011-2223320001101111-1230103032120230"></a>

## Direct properties — cloud_link / 110311323130 / 3

<a id="canonical-3113030212323203-3301122231311001-2321223122330000-1322122323103202-1021010333121320-2003132103103010-0112311120231330-1231003301011012"></a>

<a id="canonical-2233321123031023-1010002300303332-0003202123203003-2201302121200221-1111122301033303-1321310112310022-2123320222303030-2320121311320010"></a>

## name property — cloud_link / 110311323130 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0111311013132000-0222013002022200-3311302100100133-3230321211102021-3011033131032221-2033123233031110-1132010031121330-1203220002120013"></a>

<a id="canonical-2211202113000212-2103001021333303-2011320122122223-2022331032312220-3101132300301232-2132231221330233-3000323202212001-1001220032132200"></a>

## namespace property — cloud_link / 110311323130 / 5

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

<a id="canonical-1221013102031232-3213313003112021-1030131121301301-2120131321001321-0222201101021032-2310233313033031-2301132113322030-1303112102202030"></a>

<a id="canonical-2210323332200213-1003110133132222-0100123030020110-3020213310220123-0012302201303233-0301003331001221-3010030011302031-0332130301320203"></a>

## tenant property — cloud_link / 110311323130 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3133202123131221-3300311101203110-1013000102013322-1303201313230113-2333110213210102-0202220202122320-3122311211310212-1200300123121223"></a>

## Next pages — cloud_link / 110311323130 / 7

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0211133132022113-2233221020320212-3203320003011110-3131100111312201-3002330203131030-0223230111220303-0122123022033020-2001230331012031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031020310013222-1123212333102112-0323033332322000-3330333303230020-2300001022033112-0221233320022131-1203213303121121-2012213003313211"></a>

## private_connectivity.inside — inside / 030021213113 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- private_connectivity.inside

<a id="canonical-2123333130123022-1110110312001211-2310133230123132-2001002022120321-1230232130202110-3113022303100132-1103301210213130-1210220301310013"></a>

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
inside = {}
```

<a id="canonical-3300120203320021-2323020013220103-1020120011013113-0031322102110032-0201001001121021-0030223202011223-3030130212131132-3023012311323312"></a>

## Direct properties — inside / 030021213113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3113112012002020-2103320202212012-1123232210303020-3200310231101200-1310103210203133-0232233221030123-3123020312100231-1223130101321203"></a>

## Next pages — inside / 030021213113 / 4

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0021002223313121-1133033102200312-2330110313321233-2023322231000212-3200112321230313-0000020212001221-1331012331201322-1103131102232100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113012131331101-1223212301111331-2222212323123112-1202333000020122-2030100022102012-3122031331331132-0113213132010023-2311212311112313"></a>

## private_connectivity.outside — outside / 310012312130 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- private_connectivity.outside

<a id="canonical-1012102111312103-0022202011202232-0330021101113010-2111023031313032-1113100320032301-3112020202000333-1233012010113022-0313132202113100"></a>

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
outside = {}
```

<a id="canonical-3001100130031222-2122323321332133-0222020310010101-2122101030032131-3303210203013023-2310230113233200-0312313331121321-1002133102312201"></a>

## Direct properties — outside / 310012312130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012332211320000-3203222211111300-2212333200010111-2211212033111331-3310320033322132-2303302200330020-3213233031022102-1331311312103200"></a>

## Next pages — outside / 310012312130 / 4

- [private_connectivity](resources--aws_vpc_site--reference--group-004.md#canonical-1033321313213300-3312102121203101-2023112202031312-2002230233131030-3320121000233132-3120322133021211-2332131021231000-2321013121233001)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1320213313311232-0331111200330103-3031121112123021-0310133212030023-0022311002322333-1033131021322001-3013223002132330-3033332120222200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131311321132320-2110323122311003-0230232200033330-1220010231010033-0311333010210333-2021322010031130-2010220300002302-1301133113320301"></a>

## sw — sw / 100303022023 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- sw

<a id="canonical-3302121131223112-0323203102333122-3221222133232322-1031230120131200-3321012133111130-1021110211210001-2110122103301100-1303100122332200"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011213112313313-1233112121102222-2000113201102030-0220023132033321-3110200230311010-0010133003001102-2123001122032010-3323232013100033"></a>

## Direct properties — sw / 100303022023 / 3

- [default_sw_version](resources--aws_vpc_site--reference--group-004.md#canonical-0022301312131020-0131020001212212-0011231321101001-0322021303033320-1130310330103121-3311001003130212-0330033103302310-3322133133331332): complete subsection reference.

<a id="canonical-0330331133313120-3323002110200203-1203132201001201-2003303011032223-1030220121330102-2223303020332310-0201332230130233-1022101131001203"></a>

<a id="canonical-0021012121131130-3011310012132120-2303112302111100-2122201311302133-2330331123330202-0102000030011102-0211323122203301-3221220131012223"></a>

## volterra_software_version property — sw / 100303022023 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-0321033330213101-1332230313330210-2310021123210133-0011212123000312-0012302031210022-0111021122302001-2330210200202030-0200313203122011"></a>

## Next pages — sw / 100303022023 / 5

- [sw.default_sw_version](resources--aws_vpc_site--reference--group-004.md#canonical-0022301312131020-0131020001212212-0011231321101001-0322021303033320-1130310330103121-3311001003130212-0330033103302310-3322133133331332)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0022301312131020-0131020001212212-0011231321101001-0322021303033320-1130310330103121-3311001003130212-0330033103302310-3322133133331332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0123101100112113-2003003031132322-0130221030010013-1122103122021213-0102012112100101-0302330002200233-3312021000113221-3013211133101223"></a>

## sw.default_sw_version — default_sw_version / 222321332103 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-1320213313311232-0331111200330103-3031121112123021-0310133212030023-0022311002322333-1033131021322001-3013223002132330-3033332120222200)
- sw.default_sw_version

<a id="canonical-2221122213230120-0302230232210121-3022010132000012-1211133002223323-3110221233001021-1203302320013021-0311302310232123-3323110223322313"></a>

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
default_sw_version = {}
```

<a id="canonical-2333112023111131-0223003320201121-2011120321321101-1200311232331133-1301303201233233-1102331113332030-2002033300203233-1211310221113033"></a>

## Direct properties — default_sw_version / 222321332103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011230030320000-1230200121323331-0331301330100210-3112122221313310-2322311230133311-2102321313313101-2220231011220330-3122313321222212"></a>

## Next pages — default_sw_version / 222321332103 / 4

- [sw](resources--aws_vpc_site--reference--group-004.md#canonical-1320213313311232-0331111200330103-3031121112123021-0310133212030023-0022311002322333-1033131021322001-3013223002132330-3033332120222200)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3220110111301131-0033121322200310-0033001321230012-3131113012020201-3201213331302322-0033003230131202-1221000010112302-2320021200023132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323312223002023-2311000230002332-1003212130213123-0112301111301011-2012130201103111-0222130300223222-2100110111200202-2130201111100032"></a>

## timeouts — timeouts / 331203333021 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- timeouts

<a id="canonical-3222102300020111-0333230030321122-2312130111020003-2000220130021003-3322312012111022-1323111133121021-0100312012230200-3012203030200003"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-3210232221100113-1303121123130113-0201031120112320-3220132220000131-1133210222200111-1112221010133110-3011303103301013-3112313032213313"></a>

## Direct properties — timeouts / 331203333021 / 3

<a id="canonical-1110033210210111-3121330113031222-0211332322112222-1013012010133300-1022230022131000-1113222211333330-2130320311021120-3022022220303203"></a>

<a id="canonical-3103100013023223-2111322003032012-0133320131210000-3123013311133022-0012103200312300-2200223311310020-1023231020133003-1122311333303121"></a>

## create property — timeouts / 331203333021 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-0021020010033011-1113002212302021-2100102111330100-1103330300211122-1033311200023132-3200012303121133-3101331333223302-0200211210120311"></a>

<a id="canonical-3013310230032312-1223302003231010-1020310233110130-0331231330233133-1012000230122232-3203302110130121-1101331003312311-0012223300031000"></a>

## delete property — timeouts / 331203333021 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2031103132012131-0320103230022321-0110303331023210-1333133312210020-1012333233110123-2310321202033022-1023232030211320-1203013201012332"></a>

<a id="canonical-0013223323023121-1332230001301321-3311220223102221-2221102102130132-2122033200213322-2221102221130203-0022000022123103-3111203232313202"></a>

## read property — timeouts / 331203333021 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0132331020113323-1203233103021011-0101221000332312-0321212000321321-1121031320300331-0311310101013321-3202012222312121-3232232221222113"></a>

<a id="canonical-1302021021100122-0323003001023302-2102033321220003-1221202313102120-1211221221313021-2212322123010101-3111330320331110-1011300201202301"></a>

## update property — timeouts / 331203333021 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3203121313223201-3220211233012103-2212130200311331-2003023211132022-0203313013302303-3303311022210002-1201230110200102-1312223212031133"></a>

## Next pages — timeouts / 331203333021 / 8

- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020012311033233-2002220012013221-2100110310313231-2212010331110023-3011223202213022-2223213132311200-1032133012100030-3232222213212220"></a>

## voltstack_cluster — voltstack_cluster / 022012013203 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- voltstack_cluster

<a id="canonical-1001100130213133-3230331301033031-0331333111030200-0210320013203113-1123312203213202-1023022223312332-2230321323121003-3130130333013112"></a>

Type: `"object"`. single nested block, Optional.

App Stack cluster of single interface AWS nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_certified_hw",
    "az_nodes"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("default_storage",
    "storage_class_list"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("global_network_list",
    "no_global_network"),
  validators.ConflictingObjectAttributes("k8s_cluster",
    "no_k8s_cluster"),
  validators.ConflictingObjectAttributes("no_outside_static_routes",
    "outside_static_routes"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-k8s_cluster_choice": "[\"k8s_cluster\",\"no_k8s_cluster\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]",
  "x-ves-oneof-field-storage_class_choice": "[\"default_storage\",\"storage_class_list\"]"
}
```

Terraform syntax:

```terraform
voltstack_cluster {
  # Configure direct properties listed below.
}
```

<a id="canonical-2330201132202223-1002230323020120-0300210203121322-1112003211321130-3220021201301012-2223112211003220-2222222013203333-2220130302012311"></a>

## Direct properties — voltstack_cluster / 022012013203 / 3

- [active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-3300032122133111-1100302023122030-3120132312222001-3133310311312001-0320310101011102-2120111032313112-3230312032013203-3210212113333211): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0220112232330111-1201221332013323-3130113332313130-3332103123231212-3301233011310123-0130201132101023-0000212222120332-3232313200012311): complete subsection reference.

- [active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2311222122231233-2010202202201021-0030022312333203-2230200320101232-2023203302121001-2321313000002221-1213202231200303-0333010033202210): complete subsection reference.

- [allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023): complete subsection reference.

<a id="canonical-0012212101111223-3213131022302031-2212321030322322-2023013222033333-3030121120031232-3013110230221033-1122102303311121-3332230211202022"></a>

<a id="canonical-1310212120332323-3222211011322102-2032001212110110-0232302102220312-2000122220201132-3122113330212303-1122220031022013-2101310313011031"></a>

## aws_certified_hw property — voltstack_cluster / 022012013203 / 4

Type: `"string"`. Optional.

\[Enum: aws-byol-voltstack-combo\] AWS Certified Hardware. Name for AWS certified hardware. The only
possible value is \`aws-byol-voltstack-combo\`.

Upstream description:

Name for AWS certified hardware.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
  stringvalidator.OneOf("aws-byol-voltstack-combo"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "aws-byol-voltstack-combo"
  ],
  "maxLength": 64,
  "x-f5xc-constraints": {
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"aws-byol-voltstack-combo\\\"]",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

- [az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-1033112221201233-2322122022112031-3200023332213310-1122122222110233-3003212100033301-0011030211022013-1012233100230102-0110132112001101): complete subsection reference.

- [dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-0123300023211213-2113303100223102-3213112230330201-3303011310221311-1113112111312231-2000110223111302-1320230001213001-3302331322001121): complete subsection reference.

- [default_storage](resources--aws_vpc_site--reference--group-004.md#canonical-2122032123302020-1011111112301023-0200231201001002-0112003210232021-2320313333312233-3120333213012003-0333010212031300-0033131033132222): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_vpc_site--reference--group-004.md#canonical-1132210111310122-1131302322203020-2012203032002333-0012012103031121-0220322320320220-2201033330021233-3112031121320000-1133122132232301): complete subsection reference.

- [global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111): complete subsection reference.

- [k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-1002223031333221-1312133120110331-2000030232301122-0320112303200023-0220312221030033-0212231012320032-1310132121121023-2031111103310323): complete subsection reference.

- [no_dc_cluster_group](resources--aws_vpc_site--reference--group-005.md#canonical-1201130103133211-0130322100310130-1030002231330321-3012320113211113-3320121023220131-3320101131300223-1033121201321000-1102000310302133): complete subsection reference.

- [no_forward_proxy](resources--aws_vpc_site--reference--group-005.md#canonical-2313231002032112-2201321312013223-0112230330312033-2003333122232032-3133311113001101-0011131103220332-0023311001021021-2311133330312311): complete subsection reference.

- [no_global_network](resources--aws_vpc_site--reference--group-005.md#canonical-2211203113330213-2300232001010302-0011302301022233-2311112202113222-0013300032330123-3003323320102322-0211311232102023-2101230102121220): complete subsection reference.

- [no_k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-2211323200131213-1211213333312303-1102030302103231-2031030230212212-0333102113012100-2211131133011220-2333203202301123-0200302011331231): complete subsection reference.

- [no_network_policy](resources--aws_vpc_site--reference--group-005.md#canonical-3003302203010233-2003212201230112-1311312031311112-3222121001020130-1000333010121301-1003303232131221-3201010031333013-2032210121203021): complete subsection reference.

- [no_outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-1310102322003113-1232131320210012-1220231032232010-1320213211120222-3133213231120020-2233132133313031-3300233101003110-2130330303133321): complete subsection reference.

- [outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211): complete subsection reference.

- [sm_connection_public_ip](resources--aws_vpc_site--reference--group-005.md#canonical-3330300220001130-0310110133023222-3100121233311020-3003033323103312-0012202211032230-0000213020231331-2311200222213300-0312320110100312): complete subsection reference.

- [sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-005.md#canonical-2232300111003333-1310011333201021-2203131132102031-1203110020323312-0113010123223312-0013231001331122-2202311013111223-2312133012222332): complete subsection reference.

- [storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-3300031200330010-1112010200123032-2113122132333302-0123210112231022-1332012112011310-0001003120130110-2310221002131212-3231202032201211): complete subsection reference.

<a id="canonical-3001113110122010-3212023112011011-0331002301003100-2230231302222332-2103333001111303-2321200012112100-3323323120013101-2201023102233312"></a>

## Next pages — voltstack_cluster / 022012013203 / 5

- [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-3300032122133111-1100302023122030-3120132312222001-3133310311312001-0320310101011102-2120111032313112-3230312032013203-3210212113333211)
- [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0220112232330111-1201221332013323-3130113332313130-3332103123231212-3301233011310123-0130201132101023-0000212222120332-3232313200012311)
- [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2311222122231233-2010202202201021-0030022312333203-2230200320101232-2023203302121001-2321313000002221-1213202231200303-0333010033202210)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-1033112221201233-2322122022112031-3200023332213310-1122122222110233-3003212100033301-0011030211022013-1012233100230102-0110132112001101)
- [voltstack_cluster.dc_cluster_group](resources--aws_vpc_site--reference--group-004.md#canonical-0123300023211213-2113303100223102-3213112230330201-3303011310221311-1113112111312231-2000110223111302-1320230001213001-3302331322001121)
- [voltstack_cluster.default_storage](resources--aws_vpc_site--reference--group-004.md#canonical-2122032123302020-1011111112301023-0200231201001002-0112003210232021-2320313333312233-3120333213012003-0333010212031300-0033131033132222)
- [voltstack_cluster.forward_proxy_allow_all](resources--aws_vpc_site--reference--group-004.md#canonical-1132210111310122-1131302322203020-2012203032002333-0012012103031121-0220322320320220-2201033330021233-3112031121320000-1133122132232301)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- [voltstack_cluster.k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-1002223031333221-1312133120110331-2000030232301122-0320112303200023-0220312221030033-0212231012320032-1310132121121023-2031111103310323)
- [voltstack_cluster.no_dc_cluster_group](resources--aws_vpc_site--reference--group-005.md#canonical-1201130103133211-0130322100310130-1030002231330321-3012320113211113-3320121023220131-3320101131300223-1033121201321000-1102000310302133)
- [voltstack_cluster.no_forward_proxy](resources--aws_vpc_site--reference--group-005.md#canonical-2313231002032112-2201321312013223-0112230330312033-2003333122232032-3133311113001101-0011131103220332-0023311001021021-2311133330312311)
- [voltstack_cluster.no_global_network](resources--aws_vpc_site--reference--group-005.md#canonical-2211203113330213-2300232001010302-0011302301022233-2311112202113222-0013300032330123-3003323320102322-0211311232102023-2101230102121220)
- [voltstack_cluster.no_k8s_cluster](resources--aws_vpc_site--reference--group-005.md#canonical-2211323200131213-1211213333312303-1102030302103231-2031030230212212-0333102113012100-2211131133011220-2333203202301123-0200302011331231)
- [voltstack_cluster.no_network_policy](resources--aws_vpc_site--reference--group-005.md#canonical-3003302203010233-2003212201230112-1311312031311112-3222121001020130-1000333010121301-1003303232131221-3201010031333013-2032210121203021)
- [voltstack_cluster.no_outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-1310102322003113-1232131320210012-1220231032232010-1320213211120222-3133213231120020-2233132133313031-3300233101003110-2130330303133321)
- [voltstack_cluster.outside_static_routes](resources--aws_vpc_site--reference--group-005.md#canonical-3200323300201300-1200122000102133-2102310323120000-0222112120023010-0203013200232201-1002322123021121-0222003002101112-1100100101331211)
- [voltstack_cluster.sm_connection_public_ip](resources--aws_vpc_site--reference--group-005.md#canonical-3330300220001130-0310110133023222-3100121233311020-3003033323103312-0012202211032230-0000213020231331-2311200222213300-0312320110100312)
- [voltstack_cluster.sm_connection_pvt_ip](resources--aws_vpc_site--reference--group-005.md#canonical-2232300111003333-1310011333201021-2203131132102031-1203110020323312-0113010123223312-0013231001331122-2202311013111223-2312133012222332)
- [voltstack_cluster.storage_class_list](resources--aws_vpc_site--reference--group-005.md#canonical-3300031200330010-1112010200123032-2113122132333302-0123210112231022-1332012112011310-0001003120130110-2310221002131212-3231202032201211)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3300032122133111-1100302023122030-3120132312222001-3133310311312001-0320310101011102-2120111032313112-3230312032013203-3210212113333211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023201230101330-3131232120330313-3123012023030211-1121100220123213-1022122120312033-3133112113010331-3101212312121101-0332103010330313"></a>

## voltstack_cluster.active_enhanced_firewall_policies — active_enhanced_firewall_policies / 120001321120 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.active_enhanced_firewall_policies

<a id="canonical-0012010130332031-1220002330320330-0202000120331303-2230122311323211-0312320312330212-1221331013231012-1020020230121321-2210100130202202"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Upstream description:

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1232101232113320-1020331110003333-2002132111111123-2022323033213012-0232120102300011-2220312111323120-1023111030012311-0301130321111233"></a>

## Direct properties — active_enhanced_firewall_policies / 120001321120 / 3

- [enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2203023131322211-3200111322311132-3331003112210210-2102201201302102-1233120203223330-2121102033321102-3033301030230220-2201001203321000): complete subsection reference.

<a id="canonical-3311330311123023-2030120013102222-3130312112203312-0202332333133333-1122002102021310-1312223020303320-3302232131133012-0002231203313113"></a>

## Next pages — active_enhanced_firewall_policies / 120001321120 / 4

- [voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2203023131322211-3200111322311132-3331003112210210-2102201201302102-1233120203223330-2121102033321102-3033301030230220-2201001203321000)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2203023131322211-3200111322311132-3331003112210210-2102201201302102-1233120203223330-2121102033321102-3033301030230220-2201001203321000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230202213200201-1300133022201300-2102102003313330-0333210333030303-0233031223303330-1122312033123002-3230101002002110-0100231221302000"></a>

## voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies — enhanced_firewall_policies / 131210103103 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-3300032122133111-1100302023122030-3120132312222001-3133310311312001-0320310101011102-2120111032313112-3230312032013203-3210212113333211)
- voltstack_cluster.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-2032020002310320-2003132002332300-2221222300002220-0203022013223110-0000230021103023-3011303012132323-0102122213021121-3110022013311101"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023313011331312-0011211033032001-1033311020210212-2121120023030120-2222130303311113-0203231001121231-3011023313200223-2310220130313032"></a>

## Direct properties — enhanced_firewall_policies / 131210103103 / 3

<a id="canonical-3331332330313032-0302200211213333-0333013231302001-1312013213100313-0111223200020023-1023110221200111-0331221213323001-0301103230332320"></a>

<a id="canonical-0303321101322323-2300132032311300-1131022322231312-1132203121301030-1033101301113213-0321031101023331-3112133021212110-2112202103122330"></a>

## name property — enhanced_firewall_policies / 131210103103 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3203103210301012-2023332331113232-2101030303210220-1233220320230121-2230101221321012-3213022202303213-2020331013211110-2030331020300223"></a>

<a id="canonical-0311300313212320-1132031312000332-2020331302021323-3300022332021213-1210123103210011-1302001331301311-2111101011010323-0202313231223310"></a>

## namespace property — enhanced_firewall_policies / 131210103103 / 5

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

<a id="canonical-3221211120213221-2131002232201031-3130103213210231-0110100102312300-2213210233223130-2223021320300300-2200301112213121-3032030322302330"></a>

<a id="canonical-0320223021101213-2211122201321121-0102102102301230-2113212323013333-2211003112321332-2233003320133301-0123233110311301-1310131223123333"></a>

## tenant property — enhanced_firewall_policies / 131210103103 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0113322211310331-0202313120322312-1131113133100223-0021203103310023-2311233121210323-1211221002201101-3201131201111200-0323311111201021"></a>

## Next pages — enhanced_firewall_policies / 131210103103 / 7

- [voltstack_cluster.active_enhanced_firewall_policies](resources--aws_vpc_site--reference--group-004.md#canonical-3300032122133111-1100302023122030-3120132312222001-3133310311312001-0320310101011102-2120111032313112-3230312032013203-3210212113333211)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0220112232330111-1201221332013323-3130113332313130-3332103123231212-3301233011310123-0130201132101023-0000212222120332-3232313200012311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3020211202023033-2331231222100311-0212321222121031-2002202013000011-2231031003033000-0012032003221011-1020311211103310-0302013031023133"></a>

## voltstack_cluster.active_forward_proxy_policies — active_forward_proxy_policies / 222220001013 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.active_forward_proxy_policies

<a id="canonical-2102312320002120-2303022231223332-2231302200021331-3303030100001013-2331323002131100-3132302320002232-1130131021113312-0311233203212010"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3220003102013022-3210210031032201-0302020130111202-1301110012323333-1131322122213202-3021032302130132-1002130031203301-2033022030120331"></a>

## Direct properties — active_forward_proxy_policies / 222220001013 / 3

- [forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-1313233332022102-3102310313030230-3101213112012002-0213223323211310-2111023233210330-2321212331032122-0200110222130301-0203132220002231): complete subsection reference.

<a id="canonical-1313102203212321-1020030121001001-0003133011323113-1232131113113120-2131320301201201-3230013221013100-3220232313032221-3322021323130030"></a>

## Next pages — active_forward_proxy_policies / 222220001013 / 4

- [voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-1313233332022102-3102310313030230-3101213112012002-0213223323211310-2111023233210330-2321212331032122-0200110222130301-0203132220002231)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1313233332022102-3102310313030230-3101213112012002-0213223323211310-2111023233210330-2321212331032122-0200110222130301-0203132220002231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121031013231233-3322220303302131-2013310030310123-3033030000101313-2230100313020132-0011031101321133-0121210220222011-3222101301320303"></a>

## voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies — forward_proxy_policies / 011030001101 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0220112232330111-1201221332013323-3130113332313130-3332103123231212-3301233011310123-0130201132101023-0000212222120332-3232313200012311)
- voltstack_cluster.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-3113211010301112-3031120113213322-3203222223102200-1023013233111110-2133123312202101-1013310210122011-2232230103203200-1002021233230320"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103301210000321-2320303300110110-2223332031303320-2232331230311011-1133211221303133-0232023231102013-2132311220022223-3010131320213321"></a>

## Direct properties — forward_proxy_policies / 011030001101 / 3

<a id="canonical-2220020322201133-2320031100023201-3202013030013022-2111110220332231-3313323222101202-1100323010332012-3230212202333012-1211031012030101"></a>

<a id="canonical-2302001112122111-3211131101002121-1000231323122331-1310213232213033-0300000123001113-1320220012000121-0121002120330122-2310222031233333"></a>

## name property — forward_proxy_policies / 011030001101 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1301200023131231-0233130202331223-1223030332321203-0132201012031322-2200113330030212-3130012231221222-2333132122203022-2132101001100213"></a>

<a id="canonical-0002101232201231-0222000300211311-2313021201112001-2202131300123301-0221231021203122-2200333110012133-2100230222301121-3131012333312310"></a>

## namespace property — forward_proxy_policies / 011030001101 / 5

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

<a id="canonical-2033301321132322-0013231112002230-1031013201202022-1310202330223123-3232121112023223-3110233300322121-3320033022121031-1123032031332120"></a>

<a id="canonical-3011312221113332-0031031013003322-2233023211131311-1110022113131211-0331102203330232-1032000323233312-3311311103220211-3201113110131233"></a>

## tenant property — forward_proxy_policies / 011030001101 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0203313021203323-2310221111131102-0111122300030011-1312031310313321-0312212133320230-0323310212032103-1310331113231300-0132013011322312"></a>

## Next pages — forward_proxy_policies / 011030001101 / 7

- [voltstack_cluster.active_forward_proxy_policies](resources--aws_vpc_site--reference--group-004.md#canonical-0220112232330111-1201221332013323-3130113332313130-3332103123231212-3301233011310123-0130201132101023-0000212222120332-3232313200012311)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2311222122231233-2010202202201021-0030022312333203-2230200320101232-2023203302121001-2321313000002221-1213202231200303-0333010033202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3322321103200113-0220310001310303-3013111312202013-2010330112311020-2113312321220213-2311033322311211-1120132002333133-0230020110101322"></a>

## voltstack_cluster.active_network_policies — active_network_policies / 103233103330 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.active_network_policies

<a id="canonical-1121320310123312-2323102223033202-0330321310303233-0030000201021020-2123111132021200-1120211203020323-0100211312102321-2000020121002011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Upstream description:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-3121333223211202-3303030111010200-3013121231212121-3312332011132211-3132221300100023-0301212222010200-2321030131233331-1011333220231000"></a>

## Direct properties — active_network_policies / 103233103330 / 3

- [network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-1323203133031322-0022310200021022-2001331203312223-2230211210030012-1323233313121002-3333013311230010-2020113203020330-3332332222012112): complete subsection reference.

<a id="canonical-3231033202331222-3210333313021020-0333033233301102-1333023002303021-0130213330112220-0202122033010111-3123323333331211-3320123001222322"></a>

## Next pages — active_network_policies / 103233103330 / 4

- [voltstack_cluster.active_network_policies.network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-1323203133031322-0022310200021022-2001331203312223-2230211210030012-1323233313121002-3333013311230010-2020113203020330-3332332222012112)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1323203133031322-0022310200021022-2001331203312223-2230211210030012-1323233313121002-3333013311230010-2020113203020330-3332332222012112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3002223113030032-2102130311201131-1333333232100213-1013322032011201-1022013102322022-2003010300220101-3121010323301223-1200122022222102"></a>

## voltstack_cluster.active_network_policies.network_policies — network_policies / 100320310021 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2311222122231233-2010202202201021-0030022312333203-2230200320101232-2023203302121001-2321313000002221-1213202231200303-0333010033202210)
- voltstack_cluster.active_network_policies.network_policies

<a id="canonical-0221220331222331-0202220101213311-3001332000200031-3212112002220013-3113220021302020-0333223013223232-2210232111010320-3200311203223201"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Firewall Policies active for this network firewall.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
network_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020101230121020-0321212302223112-1301212323022230-0300132212110012-2001213301221330-0010311111010200-0202030203113332-0330011200013002"></a>

## Direct properties — network_policies / 100320310021 / 3

<a id="canonical-3000310031301300-3133221121212313-1301021100131121-2010202211031002-2332132330200200-0211211223100330-3100110221203101-3012323103130001"></a>

<a id="canonical-0031001330333213-0323111023233321-2203231123333120-1123011113033133-3212021002100121-0222120032011022-1030032011331113-0310002111002211"></a>

## name property — network_policies / 100320310021 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0132201110223302-3202310013313031-1221111220023023-0111110011100203-3010231000222201-3222321323103332-2122321010312303-0110020102100322"></a>

<a id="canonical-0233333201000312-0332302201321232-0223311022331131-0230103103002212-1120331231130212-2331312310120012-0122022023300330-0213233230103030"></a>

## namespace property — network_policies / 100320310021 / 5

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

<a id="canonical-3233233030223032-1131330320010130-3113000332333232-0232303003101032-1322222212121221-3332210022010123-3312123011212022-1312010120230001"></a>

<a id="canonical-0201320011110200-2221210112101223-2322113200210211-3211331130202332-2200131010233301-3231213001213201-3311332303020223-0001302011311131"></a>

## tenant property — network_policies / 100320310021 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1121222131031200-3132112232023221-3323013103330210-1030231210120103-0230223322013002-1230003002002031-0312130320032022-3322013102202030"></a>

## Next pages — network_policies / 100320310021 / 7

- [voltstack_cluster.active_network_policies](resources--aws_vpc_site--reference--group-004.md#canonical-2311222122231233-2010202202201021-0030022312333203-2230200320101232-2023203302121001-2321313000002221-1213202231200303-0333010033202210)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111212133021303-1231113111102021-3201123121200122-2023113130030013-1020120110232211-3333023323121013-0201321130211322-2012323032102330"></a>

## voltstack_cluster.allowed_vip_port — allowed_vip_port / 000221101012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.allowed_vip_port

<a id="canonical-2023222222123202-2120200011223002-3221210103222203-2100302221221212-3101221132001122-1100112210212012-1202322103310021-2223033321220211"></a>

Type: `"object"`. single nested block, Optional.

Defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client can use
the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Upstream description:

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_ports",
    "disable_allowed_vip_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_http_port"),
  validators.ConflictingObjectAttributes("custom_ports",
    "use_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_https_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("disable_allowed_vip_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_http_port"),
  validators.ConflictingObjectAttributes("use_http_https_port",
    "use_https_port"),
  validators.ConflictingObjectAttributes("use_http_port",
    "use_https_port")}
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
  "x-ves-oneof-field-port_choice": "[\"custom_ports\",\"disable_allowed_vip_port\",\"use_http_https_port\",\"use_http_port\",\"use_https_port\"]"
}
```

Terraform syntax:

```terraform
allowed_vip_port {
  # Configure direct properties listed below.
}
```

<a id="canonical-1112113233020233-2001022110231113-2011322020003312-1023010011011111-2322233111020322-2303103221100321-1302310210020121-0230313100131312"></a>

## Direct properties — allowed_vip_port / 000221101012 / 3

- [custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-2323321211021222-1321302203220003-1103302303321223-0020313003332322-2321313120010022-1213213000120222-2323133012110211-2001102111303210): complete subsection reference.

- [disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2101333002120303-3320131023331303-1302302300121230-0020130103300222-3323120221333021-1002001220133202-3010001131321021-3131230210203333): complete subsection reference.

- [use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-3223310000000100-1303210232122210-3220132302003223-3132121010132213-0231220030231331-3121310103101023-0210211202102323-1133310123212303): complete subsection reference.

- [use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-1120033103312030-2321130111001232-3222230021330310-3320002013032031-1311230133312021-3213223102010210-0002330012312013-1013031031133233): complete subsection reference.

- [use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-1230131202200321-3112012312031310-1023103323013223-0112321221133021-2211101020012101-0212331200002202-3201123113301010-2001331122022223): complete subsection reference.

<a id="canonical-1313021302131211-2221330001330221-0130333313131322-3313301300011231-0110313301012203-2133303021121223-2001122310311213-1031120330231113"></a>

## Next pages — allowed_vip_port / 000221101012 / 4

- [voltstack_cluster.allowed_vip_port.custom_ports](resources--aws_vpc_site--reference--group-004.md#canonical-2323321211021222-1321302203220003-1103302303321223-0020313003332322-2321313120010022-1213213000120222-2323133012110211-2001102111303210)
- [voltstack_cluster.allowed_vip_port.disable_allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-2101333002120303-3320131023331303-1302302300121230-0020130103300222-3323120221333021-1002001220133202-3010001131321021-3131230210203333)
- [voltstack_cluster.allowed_vip_port.use_http_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-3223310000000100-1303210232122210-3220132302003223-3132121010132213-0231220030231331-3121310103101023-0210211202102323-1133310123212303)
- [voltstack_cluster.allowed_vip_port.use_http_port](resources--aws_vpc_site--reference--group-004.md#canonical-1120033103312030-2321130111001232-3222230021330310-3320002013032031-1311230133312021-3213223102010210-0002330012312013-1013031031133233)
- [voltstack_cluster.allowed_vip_port.use_https_port](resources--aws_vpc_site--reference--group-004.md#canonical-1230131202200321-3112012312031310-1023103323013223-0112321221133021-2211101020012101-0212331200002202-3201123113301010-2001331122022223)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2323321211021222-1321302203220003-1103302303321223-0020313003332322-2321313120010022-1213213000120222-2323133012110211-2001102111303210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232001203322110-1332223131012120-3122013303033222-1132001300213012-3000131120012331-1010321000113002-2111113202113021-2033323300201023"></a>

## voltstack_cluster.allowed_vip_port.custom_ports — custom_ports / 300301310112 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- voltstack_cluster.allowed_vip_port.custom_ports

<a id="canonical-2000230121300102-2211113300131022-0300121110333222-3021132323020002-2200311020103233-2320231121302300-1111100023012031-0300200030230012"></a>

Type: `"object"`. single nested block, Optional.

Custom Ports. List of Custom port.

Upstream description:

List of Custom port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("port_ranges")}
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
custom_ports {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023212023330113-2210021030110002-0132210323330132-3222202100320122-0003103012210201-3102323313232331-3030031023230311-1120003210132001"></a>

## Direct properties — custom_ports / 300301310112 / 3

<a id="canonical-3020121322131231-1222012313010101-1100330231031001-2010221332301033-0032023131201221-2213012033303021-2120101323100231-2323233100203322"></a>

<a id="canonical-1122011102101310-1120301321101002-0323210232033110-0323012011030033-2012310311231032-3121320320330100-3101300012321030-1000123211333020"></a>

## port_ranges property — custom_ports / 300301310112 / 4

Type: `"string"`. Optional.

Port Ranges. Port Ranges.

Upstream description:

Port Ranges.

Provider validators and defaults (from schema source):

```go
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.port_range_list": "true"
  }
}
```

<a id="canonical-0323123020213130-1103013202222303-3112312021211021-2202020000320233-3003313100222203-1300130103323000-0301310333122100-2331330302311203"></a>

## Next pages — custom_ports / 300301310112 / 5

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2101333002120303-3320131023331303-1302302300121230-0020130103300222-3323120221333021-1002001220133202-3010001131321021-3131230210203333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333032321003211-1111300323212112-0220210003311313-2301321121120103-2201301202011301-1102333033121122-1100013301122212-0223202312301013"></a>

## voltstack_cluster.allowed_vip_port.disable_allowed_vip_port — disable_allowed_vip_port / 111323131213 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- voltstack_cluster.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-2321200301302100-1121002122033212-1333223002310000-3333313300202103-0211332100232022-1203322113102113-1202012310202213-1323331032203210"></a>

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
disable_allowed_vip_port = {}
```

<a id="canonical-2223111222130303-3322130003002201-1321311131112002-3330132220111030-2330212200113030-1022113300133020-1132003103220321-2100330312210302"></a>

## Direct properties — disable_allowed_vip_port / 111323131213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023333233203120-1330002212012311-2202123003132311-2312102312031311-2003310133220222-3001012231211120-2312232233012012-0312310102101311"></a>

## Next pages — disable_allowed_vip_port / 111323131213 / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-3223310000000100-1303210232122210-3220132302003223-3132121010132213-0231220030231331-3121310103101023-0210211202102323-1133310123212303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022300002330123-2001103232000323-1213322323110200-3023302113101311-1330020203002330-1001130323030301-1122130221021210-2131131203022102"></a>

## voltstack_cluster.allowed_vip_port.use_http_https_port — use_http_https_port / 301021133012 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- voltstack_cluster.allowed_vip_port.use_http_https_port

<a id="canonical-0301022122123023-2122210122132022-1232101212201301-1121001032323112-0322323301232312-0021023300301113-2032111021221010-3330030032102331"></a>

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
use_http_https_port = {}
```

<a id="canonical-2201122103003323-0112123212002103-3233311302321222-0120021021333033-3030121113012012-0112310122002003-3223331311323111-3301200230230002"></a>

## Direct properties — use_http_https_port / 301021133012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3013203300111332-1030211302023132-3031120312011021-2232233021331023-2200231033210011-0311112112030121-1323132132302001-2220031233331233"></a>

## Next pages — use_http_https_port / 301021133012 / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1120033103312030-2321130111001232-3222230021330310-3320002013032031-1311230133312021-3213223102010210-0002330012312013-1013031031133233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011132013130133-1221333101212130-0313133330020022-1330302022333323-3330300100332002-1103010222302122-0210110203121122-2121210121321302"></a>

## voltstack_cluster.allowed_vip_port.use_http_port — use_http_port / 022200000123 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- voltstack_cluster.allowed_vip_port.use_http_port

<a id="canonical-2330221001303312-3220002102032100-0322001313012310-3031031033233010-3123302032233000-2232302323110323-2100001330100033-3301300331333132"></a>

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
use_http_port = {}
```

<a id="canonical-1002011233133021-0210112331312203-1121302021303212-2323231033132012-3200300011132302-0312232221100210-1031123222221013-1211111121112222"></a>

## Direct properties — use_http_port / 022200000123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0023222302230312-2001111013302122-3100003332333023-3012002112010213-1133200002311033-0100202131211122-2222322210210211-2220010002010103"></a>

## Next pages — use_http_port / 022200000123 / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1230131202200321-3112012312031310-1023103323013223-0112321221133021-2211101020012101-0212331200002202-3201123113301010-2001331122022223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2002033003013020-3012211220133311-3013332301232220-1131030111030232-3130023202333020-0110301320011231-0223030233111333-3321212332113000"></a>

## voltstack_cluster.allowed_vip_port.use_https_port — use_https_port / 102223201330 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- voltstack_cluster.allowed_vip_port.use_https_port

<a id="canonical-2333222321312300-3203000012121311-0331032230020110-2002020002332321-0232210103011301-2112320112023110-2222020101022202-2311203220020330"></a>

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
use_https_port = {}
```

<a id="canonical-0331301313010222-2133330201022112-1220102011303211-0132123203212023-0120223102311333-3121121110033303-1001213003012312-3031013211310333"></a>

## Direct properties — use_https_port / 102223201330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003032310032022-2033021333302212-0320130021121202-3210202132231231-0323132133223202-1210113100121322-1002030113300011-3302110323333201"></a>

## Next pages — use_https_port / 102223201330 / 4

- [voltstack_cluster.allowed_vip_port](resources--aws_vpc_site--reference--group-004.md#canonical-0133233213100033-2233032230021202-3120021121323111-0332211003122012-3203112200132030-0022222012132203-0021311123300133-0311203223020023)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1033112221201233-2322122022112031-3200023332213310-1122122222110233-3003212100033301-0011030211022013-1012233100230102-0110132112001101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221010301210222-0313132111203301-2032331211102122-2101110223001000-1332233333330030-0220233103202100-1302211011230333-1113333120333100"></a>

## voltstack_cluster.az_nodes — az_nodes / 120222023313 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.az_nodes

<a id="canonical-2120331132021020-0003101210232312-2300120213022310-0132130120312233-0033020102012102-2332222233300301-3312223212201033-2203332003012212"></a>

Type: `"object"`. list nested block, Optional.

Only Single AZ or Three AZ(s) nodes are supported currently.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("aws_az_name")}
```

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
    "ves.io.schema.rules.repeated.num_items": "1,3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.num_items": "1,3"
  }
}
```

Terraform syntax:

```terraform
az_nodes {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231110222022023-1023002111100023-3100131131330310-1022032002013311-1333220220023333-2311331202230020-3200013302121112-1033312322321013"></a>

## Direct properties — az_nodes / 120222023313 / 3

<a id="canonical-0333033002000230-1233201120110312-3022300323000120-2032001233003100-2230133013102023-0033000312220112-0030011031013031-3020210301133011"></a>

<a id="canonical-0122322202202011-2001111323001032-0300320220002122-3023233211330310-1102112121023230-2213213303231122-3121311100333212-3133031002000323"></a>

## aws_az_name property — az_nodes / 120222023313 / 4

Type: `"string"`. Optional.

AWS availability zone, must be consistent with the selected AWS region.

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

- [local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-0130313102230122-2223032002313201-2310201103321300-3333031201323200-0101023113112123-0033030003022222-0113133131202113-3201300013112003): complete subsection reference.

<a id="canonical-2202131220310210-0122301332322223-0323100210230030-3302130013222332-3301131311001030-2013030322210003-0222010211302101-2203220213302223"></a>

## Next pages — az_nodes / 120222023313 / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-0130313102230122-2223032002313201-2310201103321300-3333031201323200-0101023113112123-0033030003022222-0113133131202113-3201300013112003)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0130313102230122-2223032002313201-2310201103321300-3333031201323200-0101023113112123-0033030003022222-0113133131202113-3201300013112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100302021001301-2331012223332111-1330201103221110-1002001333330131-3132023330233100-1230030120013310-1002122302300301-1313323113110231"></a>

## voltstack_cluster.az_nodes.local_subnet — local_subnet / 200110311110 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-1033112221201233-2322122022112031-3200023332213310-1122122222110233-3003212100033301-0011030211022013-1012233100230102-0110132112001101)
- voltstack_cluster.az_nodes.local_subnet

<a id="canonical-1030322133020120-3322322010013332-2303323320221003-0022103031331210-3300011202232223-0133103221001303-1033320311033203-0312110103220032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for local subnet.

Upstream description:

Parameters for AWS subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("existing_subnet_id",
    "subnet_param")}
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
  "x-ves-oneof-field-choice": "[\"existing_subnet_id\",\"subnet_param\"]"
}
```

Terraform syntax:

```terraform
local_subnet {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030122130033332-0222231232113323-0112220322333203-1101030311122113-1300221012103313-2003111110331322-1121223121322211-2322310031333113"></a>

## Direct properties — local_subnet / 200110311110 / 3

<a id="canonical-1322021221023321-2101323323310022-3201332010311303-3013323033303021-3013311103302213-1220301310132113-3123000120223033-2023103301203030"></a>

<a id="canonical-2213202011100310-0322310332330030-3302123311101322-0331122103102102-2011000302323111-3120321122320023-3020313202221133-1200310122233103"></a>

## existing_subnet_id property — local_subnet / 200110311110 / 4

Type: `"string"`. Optional.

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Upstream description:

Exclusive with \[subnet\_param\] Information about existing subnet ID.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.pattern": "^(subnet-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

- [subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-2322100113121110-3000233232212333-3232322013110130-2001102212103100-3130322303110011-0221100203330201-2303113031012233-3122213000130121): complete subsection reference.

<a id="canonical-3103313323231222-3202301022232303-2310020111332233-0300313313202301-0222323320310000-1110221112230111-3010321200030030-3132002211331100"></a>

## Next pages — local_subnet / 200110311110 / 5

- [voltstack_cluster.az_nodes.local_subnet.subnet_param](resources--aws_vpc_site--reference--group-004.md#canonical-2322100113121110-3000233232212333-3232322013110130-2001102212103100-3130322303110011-0221100203330201-2303113031012233-3122213000130121)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-1033112221201233-2322122022112031-3200023332213310-1122122222110233-3003212100033301-0011030211022013-1012233100230102-0110132112001101)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2322100113121110-3000233232212333-3232322013110130-2001102212103100-3130322303110011-0221100203330201-2303113031012233-3122213000130121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321031010201302-0333302233130022-1103313233011122-0200230223321003-2202312222222311-3213321333210120-0033033100300003-0022111000332331"></a>

## voltstack_cluster.az_nodes.local_subnet.subnet_param — subnet_param / 222311033032 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.az_nodes](resources--aws_vpc_site--reference--group-004.md#canonical-1033112221201233-2322122022112031-3200023332213310-1122122222110233-3003212100033301-0011030211022013-1012233100230102-0110132112001101)
- [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-0130313102230122-2223032002313201-2310201103321300-3333031201323200-0101023113112123-0033030003022222-0113133131202113-3201300013112003)
- voltstack_cluster.az_nodes.local_subnet.subnet_param

<a id="canonical-3111030101101210-3103133023000121-3023102023012032-0021001121001312-1123333212321001-2312001033020212-2300332013100023-2200220321032330"></a>

Type: `"object"`. single nested block, Optional.

Parameters for creating a new cloud subnet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ipv4")}
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
subnet_param {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210230223111123-0201323222102123-0021113323310310-0103201132103201-0003000122220032-3203223001112301-3321130203313203-2301010222230312"></a>

## Direct properties — subnet_param / 222311033032 / 3

<a id="canonical-3212013122231123-0002231331010313-0211100302022320-3222200203231203-2332333021030313-0021002220133221-3032100222111103-3111203332111100"></a>

<a id="canonical-1113030103312330-1003123303030033-3023201011231002-2121130112320322-2000020201112112-2121231002320101-2010202200130021-3311112212033231"></a>

## IPv4 property — subnet_param / 222311033032 / 4

Type: `"string"`. Optional.

IPv4 Subnet. IPv4 subnet prefix for this subnet.

Upstream description:

IPv4 subnet prefix for this subnet.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "ipv4",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ipv4_prefix": "true",
    "ves.io.schema.rules.string.max_ip_prefix_length": "28"
  }
}
```

<a id="canonical-3103011321230333-0020310102102010-0101221123201102-1032210221312010-2322110122120002-0122100322002103-0121232233103122-0323020202110222"></a>

## Next pages — subnet_param / 222311033032 / 5

- [voltstack_cluster.az_nodes.local_subnet](resources--aws_vpc_site--reference--group-004.md#canonical-0130313102230122-2223032002313201-2310201103321300-3333031201323200-0101023113112123-0033030003022222-0113133131202113-3201300013112003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-0123300023211213-2113303100223102-3213112230330201-3303011310221311-1113112111312231-2000110223111302-1320230001213001-3302331322001121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122221103033021-1322322223100223-0231013131130102-1331111301202221-1101213323201312-2332111110030033-0203032232212022-0310301211023200"></a>

## voltstack_cluster.dc_cluster_group — dc_cluster_group / 120101122111 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.dc_cluster_group

<a id="canonical-1021210323103133-0230302223310022-2113131130310211-3120302030300022-3220110320221211-1112132120120320-2300100300220033-3103200103233001"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
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
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002112333312032-2232021301000301-2101131201120123-3113313210210230-3212322011121023-1332301312322213-1213303321320231-2131312330231223"></a>

## Direct properties — dc_cluster_group / 120101122111 / 3

<a id="canonical-1300301123101013-1012102301312003-3122110131332121-0102210201311332-3133001020222000-1203103300313030-3222023300001011-1020023031321002"></a>

<a id="canonical-1332112202233232-3100301322323331-2120313213120112-2232102211132322-0232330312201302-3130202131123321-1112312003103021-3113330213332033"></a>

## name property — dc_cluster_group / 120101122111 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0313330011021321-1321103030120220-0100000303000330-3232310103333012-2311002101013221-2011113132231220-1203001230310300-2220331013200220"></a>

<a id="canonical-2002120211310203-2011333320222121-0223320231031033-2130232333130222-3232133310130002-3200011231103033-3023003022233210-1131310331111321"></a>

## namespace property — dc_cluster_group / 120101122111 / 5

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

<a id="canonical-3001233110333101-1222033022230033-3032010010220112-3223122313202201-0110031122001001-1002302030203123-1313123323130221-1011330230303302"></a>

<a id="canonical-1002010333113200-2203301222121110-3012130020313101-0023321020002323-2203132310130013-0223113103320233-2123103230221220-2222231013120300"></a>

## tenant property — dc_cluster_group / 120101122111 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0332310301320203-3131130331010001-1103230230103012-2013111200232120-2012031301313332-2112103110200001-0012320331320013-0003310220123013"></a>

## Next pages — dc_cluster_group / 120101122111 / 7

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2122032123302020-1011111112301023-0200231201001002-0112003210232021-2320313333312233-3120333213012003-0333010212031300-0033131033132222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333203003101011-2132332021321331-1101203232021323-3331311201012003-3302302233030020-0023022313213230-2023021220331223-1220310020300322"></a>

## voltstack_cluster.default_storage — default_storage / 021122033212 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.default_storage

<a id="canonical-2032232203333220-3201331221322023-1120113200013131-3303100230201332-1033102032122333-0023221233011232-1130223222031202-0302221332221110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default storage.

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
default_storage = {}
```

<a id="canonical-1233122100330300-0221213332111000-0320022001111311-0313131230230312-2123112302113121-1003021012331222-0210022011102000-3312332230230202"></a>

## Direct properties — default_storage / 021122033212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3110012010333113-3332203202213002-2202203222012031-2211032323001030-0333311031032120-3211302331201022-1113110023102332-0313122310323032"></a>

## Next pages — default_storage / 021122033212 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-1132210111310122-1131302322203020-2012203032002333-0012012103031121-0220322320320220-2201033330021233-3112031121320000-1133122132232301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132321003323300-1222212121112322-2323030011002320-2133310021231301-0031330100320033-3220130211033011-0130203212322303-3312100132130103"></a>

## voltstack_cluster.forward_proxy_allow_all — forward_proxy_allow_all / 321022222322 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.forward_proxy_allow_all

<a id="canonical-2111332012000301-0123321313211202-3211200031003332-0030111112010031-3030211331130001-0300123332031112-0103210230230321-2031011331122230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for forward proxy allow all.

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
forward_proxy_allow_all = {}
```

<a id="canonical-3003001301102222-0133110233233111-0203330002131302-0332303203122332-2323031033201012-3012312302031030-2201023113301100-3133231100230231"></a>

## Direct properties — forward_proxy_allow_all / 321022222322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021233310010331-3203010200312113-0230301212031033-3310322200020321-2132123201103211-0212230103133303-3230010110231301-0112301003232120"></a>

## Next pages — forward_proxy_allow_all / 321022222322 / 4

- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232100221003022-1222230012031323-2023333322230211-0011310101011121-0022001111231123-0230010302031132-1312020322223320-0010021031011330"></a>

## voltstack_cluster.global_network_list — global_network_list / 320303313002 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- voltstack_cluster.global_network_list

<a id="canonical-3201311021111031-2230113102203230-1030323011111311-1001121321231032-1110113310300220-0320301231212111-0101323303311211-3201332200010110"></a>

Type: `"object"`. single nested block, Optional.

Global Network Connection List. List of global network connections.

Upstream description:

List of global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("global_network_connections")}
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
global_network_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313033310100323-2203333033033222-0310301211230123-3302310220322031-0313310231200102-1333003223101302-3200232101203330-3113201020111300"></a>

## Direct properties — global_network_list / 320303313002 / 3

- [global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311): complete subsection reference.

<a id="canonical-2312100222313032-0330320021133011-0311132311203000-2221311320123222-1212301331102132-0001323133122123-1210210010113333-2012233023201120"></a>

## Next pages — global_network_list / 320303313002 / 4

- [voltstack_cluster.global_network_list.global_network_connections](resources--aws_vpc_site--reference--group-004.md#canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)

<a id="canonical-2220100231202200-1000303320332323-0223113001021233-1013001212100200-0001230320000202-0110123212001021-2233111031310210-0023303011021311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2003111311100030-1313100313302132-0113100102201221-2100222200332022-0303211003200322-1213200230100231-1012211001330101-0303210220123131"></a>

## voltstack_cluster.global_network_list.global_network_connections — global_network_connections / 310010331231 / 2

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md#canonical-1121120120112313-2031303113302001-3322310321202321-3023211110122101-0112222301201113-1120203320333230-0020311023113132-2302133101311311)
- [Property reference](resources--aws_vpc_site--reference--group-001.md#canonical-0301233321013203-1300103011022100-2111131023322220-0102222101012002-2322331200110113-3123301012113021-1201332322031122-0103122213110330)
- [voltstack_cluster](resources--aws_vpc_site--reference--group-004.md#canonical-2321300031300012-0132210010103123-0233200321322223-1302113212020030-0212310212010100-2020221032202133-1333332221203211-0333331021201003)
- [voltstack_cluster.global_network_list](resources--aws_vpc_site--reference--group-004.md#canonical-2203021112020001-3133120221323131-3110312113100112-2200221102023130-1333113333312321-3332230222220203-3213111121221120-3032201310331111)
- voltstack_cluster.global_network_list.global_network_connections

<a id="canonical-1312113030303303-2303320113121112-0232020023132112-3023020031120303-2301123130001102-1333311231220302-3031233023023210-3000012220311032"></a>

Type: `"object"`. list nested block, Optional.

Global Network Connections. Global network connections.

Upstream description:

Global network connections.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("sli_to_global_dr",
    "slo_to_global_dr")}
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
global_network_connections {
  # Configure direct properties listed below.
}
```
