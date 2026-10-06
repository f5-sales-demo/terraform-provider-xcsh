---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-3023032101103211-1222102132121230-0311102331000023-3123210102330121-2232212002002033-0223213203212313-1100003110201222-2220221203033003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.no_network_policy` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [tgw_security](data-sources--aws_tgw_site--reference--group-002.md#canonical-0232013222110121-3223210311032231-0332221012122031-2132020312222132-2303023010020032-2101212313223322-1203003033031332-1111320102110212)
- tgw_security.no_network_policy

<a id="canonical-1033230301303120-1031212130330003-1132331221223312-0331311231330122-0322302032233132-1111331031123131-2122130311310110-0232313101113211"></a>

Type: `["object", {}]`. Computed.

Policy configuration for this feature.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- vn_config

<a id="canonical-1132130212133133-1300231313303102-3012311220221030-3002121133320010-3113121310300211-1221101102322223-2121132222023331-2213031113012122"></a>

Type: `"single"`. Computed.

Virtual Network Configuration. Virtual Network Configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group_inside_vn\",\"dc_cluster_group_outside_vn\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-global_network_choice": "[\"global_network_list\",\"no_global_network\"]",
  "x-ves-oneof-field-inside_static_route_choice": "[\"inside_static_routes\",\"no_inside_static_routes\"]",
  "x-ves-oneof-field-outside_static_route_choice": "[\"no_outside_static_routes\",\"outside_static_routes\"]",
  "x-ves-oneof-field-site_mesh_group_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

<a id="canonical-3312303002330011-1022032301303131-3211201322120031-2020233303230133-1202333120000310-0203120123113113-0112222233032112-0331231123321020"></a>

### Direct properties for `vn_config`

- [allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110): complete subsection reference.

- [allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113): complete subsection reference.

- [dc_cluster_group_inside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-0123201232133002-3320121122011110-3332021222121331-3133012323330001-3313232012210123-1003323303123012-0333212223201112-3111312232100020): complete subsection reference.

- [dc_cluster_group_outside_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-3012010120110213-0221112133332320-0033033112122032-2103012332320020-1011323100121300-1113003111311133-3123003331123032-3020321300210202): complete subsection reference.

- [global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203): complete subsection reference.

- [inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320): complete subsection reference.

- [no_dc_cluster_group](data-sources--aws_tgw_site--reference--group-003.md#canonical-3322033323111233-2333300011113211-2303310012333330-1032323330033233-1133313111333130-1303303300213210-1222223033330233-0322021301022133): complete subsection reference.

- [no_global_network](data-sources--aws_tgw_site--reference--group-003.md#canonical-0211231303000231-1332322333303110-2231330333122021-3022013113310302-0331223221023310-0321000310301123-1112130032320132-0012101113120333): complete subsection reference.

- [no_inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-2133103000303110-2322132103102100-0111231000032223-1012202032122131-2132030320322112-2311233020023323-1010100332321223-0212021022322130): complete subsection reference.

- [no_outside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-2220121110310322-3310320322232321-2200001213032233-3220332331310131-1101133323003022-1221011103131111-2032233010211203-1023100311031333): complete subsection reference.

- [outside_static_routes](data-sources--aws_tgw_site--reference--group-004.md#canonical-1121012233031120-1121331231332113-2030000031313021-0331310223301120-2030100001010123-2120013302323321-2003102022200221-0301130120023301): complete subsection reference.

- [sm_connection_public_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-3310333131011322-0330100020123222-2030322120221110-2111221333212323-0130210132002022-1102312332302113-2031133120302121-0210233110231121): complete subsection reference.

- [sm_connection_pvt_ip](data-sources--aws_tgw_site--reference--group-004.md#canonical-0330213031123210-3033120010312021-0023130021033220-2313022211120033-1111033320230122-0232222032233123-2302100330110023-3000000033133323): complete subsection reference.

<a id="canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.allowed_vip_port

<a id="canonical-2021313021303233-3323323013313223-0333210033112010-1232123310002001-2222321001133230-1020020221222202-0000100100203030-1002102111231323"></a>

Type: `"single"`. Computed.

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

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

<a id="canonical-1012230021220002-3002000320101023-3102133231311303-1111332322112112-2122121123102313-0013010003012031-1201302221200331-1310100233302330"></a>

### Direct properties for `vn_config.allowed_vip_port`

- [custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-0220303100122022-0032201233010233-0103123103211001-0232333132031203-1330212330310113-3232111022231211-1001310011333311-0031313030102300): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-1110313231010221-0013213233122020-2313331002013113-3120130021222331-3112213231031000-0033201322312023-2101012211312030-0111113211022133): complete subsection reference.

- [use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-0312111012100000-1113120233111022-0202203022301031-0112203310003002-2211000110303222-0202210022130020-1302033100303232-3223313012011220): complete subsection reference.

- [use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2033110213302332-1232023311221133-2231201003203103-0321313311302302-2011200120212030-0232100022213213-2211121000231220-3200020330233232): complete subsection reference.

- [use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2203222333103031-0120331032112321-2003102123013120-2302103100112120-3033333322030030-1022333123201100-2220310203033331-2011122121211120): complete subsection reference.

<a id="canonical-0220303100122022-0032201233010233-0103123103211001-0232333132031203-1330212330310113-3232111022231211-1001310011333311-0031313030102300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.custom_ports` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110)
- vn_config.allowed_vip_port.custom_ports

<a id="canonical-3313033231023322-0101121202310330-3030300332321333-2202123130101331-0222330212301312-1020022312032212-1223212310201132-3201300012133002"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

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

<a id="canonical-1023201022300222-3122010021332333-0133103232121301-3012121223302023-3133300000032323-3131113123210203-3233333133102123-2102322033232113"></a>

### Direct properties for `vn_config.allowed_vip_port.custom_ports`

<a id="canonical-0312113303023131-0302230332300302-2131130001212303-3030103022122010-3333102103131232-2001202003001203-1200012313223323-1003032020232130"></a>

#### `vn_config.allowed_vip_port.custom_ports.port_ranges` property

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1110313231010221-0013213233122020-2313331002013113-3120130021222331-3112213231031000-0033201322312023-2101012211312030-0111113211022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.disable_allowed_vip_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110)
- vn_config.allowed_vip_port.disable_allowed_vip_port

<a id="canonical-0303321021030230-3022302111323231-3333211023022112-3021032032021311-3303332103003000-1000132323112233-2302013200113123-0110103011233222"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312111012100000-1113120233111022-0202203022301031-0112203310003002-2211000110303222-0202210022130020-1302033100303232-3223313012011220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.use_http_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110)
- vn_config.allowed_vip_port.use_http_https_port

<a id="canonical-3133013112010310-3310120103132101-3121201100233122-0232322211001033-1120012011021023-0201010133233032-2222031012320302-3230033231310010"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033110213302332-1232023311221133-2231201003203103-0321313311302302-2011200120212030-0232100022213213-2211121000231220-3200020330233232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.use_http_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110)
- vn_config.allowed_vip_port.use_http_port

<a id="canonical-2333123322021212-1012023313221123-1303112311003132-2232321331112220-3200201022002112-0113013200223201-2233132203032031-0100010102023021"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203222333103031-0120331032112321-2003102123013120-2302103100112120-3033333322030030-1022333123201100-2220310203033331-2011122121211120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port.use_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2312123133020312-3232021220010232-1132201301121131-0020101111100121-2103133022200200-1221020013101212-1133122300113202-3303200333320110)
- vn_config.allowed_vip_port.use_https_port

<a id="canonical-1200321100121303-2301132301320123-2030011030302332-1202332213131301-2221101132330331-2112302322202111-0122101220301032-3132320113111332"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.allowed_vip_port_sli

<a id="canonical-2012010110313302-0031030303221221-0212130003032210-0103022103101330-1121011332203210-3322231133001220-0210003233312011-2201212132202112"></a>

Type: `"single"`. Computed.

This defines the TCP port(s) which will be opened on the cloud loadbalancer. Such that the client
can use the cloud VIP IP and port combination to reach TCP/HTTP LB configured on the F5XC Site.

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

<a id="canonical-0123021003021020-2221301311321011-2202000021132131-0220233110012300-0100222000123331-0030133023102021-3011120110230211-2001031120223313"></a>

### Direct properties for `vn_config.allowed_vip_port_sli`

- [custom_ports](data-sources--aws_tgw_site--reference--group-003.md#canonical-1002233213133130-0130221313021200-1323222223302300-2120302312130210-3221101013321203-2212220223032212-1312012000311123-3133211022103010): complete subsection reference.

- [disable_allowed_vip_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-0102213133321030-0202321222003301-2111121320003120-3200131332032312-0323132333123031-1301031303013102-3003221103103212-1022213302102323): complete subsection reference.

- [use_http_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-2003303120023120-0303223013320231-1212200101222321-0202320032120021-1112133320111301-3303002103202120-1210322132310112-3120323020302302): complete subsection reference.

- [use_http_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-3330102322220212-3101301222020222-0111001333103202-2233023121013022-1001100121323100-2223103101223021-2221311320101312-1330033122013101): complete subsection reference.

- [use_https_port](data-sources--aws_tgw_site--reference--group-003.md#canonical-1311033132201113-1213111020001232-2313010203102103-1310202311311020-1101330100202112-1232013001031003-2301223223313113-2223000000221323): complete subsection reference.

<a id="canonical-1002233213133130-0130221313021200-1323222223302300-2120302312130210-3221101013321203-2212220223032212-1312012000311123-3133211022103010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.custom_ports` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113)
- vn_config.allowed_vip_port_sli.custom_ports

<a id="canonical-2132203123002302-3022122030102110-2021323121303210-3003001133331223-2003231112010010-1130201331230213-1130102231222311-2220230213211223"></a>

Type: `"single"`. Computed.

Custom Ports. List of Custom port.

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

<a id="canonical-2011320131333101-0103232123012111-3333122111123030-3302130002121200-2023101003130211-3313133301332021-3230201223210030-2223020312003201"></a>

### Direct properties for `vn_config.allowed_vip_port_sli.custom_ports`

<a id="canonical-2200002113221332-3122230301311312-1333302323312111-2113110113231231-3110212101120322-1130131200133333-0121103210001313-2000022203321122"></a>

#### `vn_config.allowed_vip_port_sli.custom_ports.port_ranges` property

Type: `"string"`. Computed.

Port Ranges. Port Ranges.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0102213133321030-0202321222003301-2111121320003120-3200131332032312-0323132333123031-1301031303013102-3003221103103212-1022213302102323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.disable_allowed_vip_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113)
- vn_config.allowed_vip_port_sli.disable_allowed_vip_port

<a id="canonical-3133023221300322-0022231131012001-3020331130323331-2203301301301110-0013021232212311-2011222112133221-1101201001112201-3112101023131003"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003303120023120-0303223013320231-1212200101222321-0202320032120021-1112133320111301-3303002103202120-1210322132310112-3120323020302302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.use_http_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113)
- vn_config.allowed_vip_port_sli.use_http_https_port

<a id="canonical-0120321211232223-2202313121132120-1023101232323213-3320013301010030-2103312322221212-1221302203300223-3332102232303210-0201032200100121"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330102322220212-3101301222020222-0111001333103202-2233023121013022-1001100121323100-2223103101223021-2221311320101312-1330033122013101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.use_http_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113)
- vn_config.allowed_vip_port_sli.use_http_port

<a id="canonical-0211200231120212-3310113121000230-0020031120211333-1221212332022333-2132221131331120-3103111210110300-2223120331003003-2100103320332000"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1311033132201113-1213111020001232-2313010203102103-1310202311311020-1101330100202112-1232013001031003-2301223223313113-2223000000221323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.allowed_vip_port_sli.use_https_port` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.allowed_vip_port_sli](data-sources--aws_tgw_site--reference--group-003.md#canonical-1320122013130233-1233313130322210-1130202020333203-2131102020311033-1200013112110122-0302000210020203-3032022010322233-0302222222310113)
- vn_config.allowed_vip_port_sli.use_https_port

<a id="canonical-1301331312330112-1113302222101133-3311030300002003-0022112200130103-3321202103020121-2031023010110231-2323021230321222-1202322110102131"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123201232133002-3320121122011110-3332021222121331-3133012323330001-3313232012210123-1003323303123012-0333212223201112-3111312232100020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.dc_cluster_group_inside_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.dc_cluster_group_inside_vn

<a id="canonical-0013221021112231-1220211022001000-2002131201333221-2213322013121012-2002133003222322-2100333301322012-1030233103201023-0201213121203021"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0313310033202212-1132201000020211-3323313032112232-1032232320323211-2312131033222020-3031131212323213-1322133021121201-2321300000322022"></a>

### Direct properties for `vn_config.dc_cluster_group_inside_vn`

<a id="canonical-0130130012013132-2300233232112103-0230223031010323-0113212232330233-1313311102210112-3101133202322032-2220330303131202-3121132030312313"></a>

#### `vn_config.dc_cluster_group_inside_vn.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2303032022333220-3320331302213202-0333111003001020-0032211013322201-0011023003331111-1233323233133021-2002123231203333-3220330233231301"></a>

<a id="canonical-0330013000123202-1202001333203323-0101221203223320-2201211220033102-2100300330130203-2012023022203313-2300213012020033-1302031321002320"></a>

#### `vn_config.dc_cluster_group_inside_vn.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2033211012311130-0003032022110310-1132022330031332-0311321010332032-1303230331212230-0200011030201210-1331113001300102-3023301011333100"></a>

<a id="canonical-3030222200211223-3121013231100322-0121223030201202-1121101323012002-3230221210322322-2221012211120221-2132032110101100-1303331332111202"></a>

#### `vn_config.dc_cluster_group_inside_vn.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3012010120110213-0221112133332320-0033033112122032-2103012332320020-1011323100121300-1113003111311133-3123003331123032-3020321300210202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.dc_cluster_group_outside_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.dc_cluster_group_outside_vn

<a id="canonical-0233133011211132-2131203201020322-3230113132223030-2002012011220003-0123130322020233-2120111202223113-1120010301320030-2020010100231012"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0200231112112311-2203323210111132-0223323001320003-3302020030331303-0112133112231120-3012312330000322-1302203102131230-1232013032303302"></a>

### Direct properties for `vn_config.dc_cluster_group_outside_vn`

<a id="canonical-2133100021103131-1330030113013000-3320033222110232-0232311320301031-1313222010322120-3131212202333203-1121030122001022-3302012233011232"></a>

#### `vn_config.dc_cluster_group_outside_vn.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2031201202102112-0122223001113232-0332332211213133-1120311111131232-1022331102113132-0201112313322223-2120133230223100-1303120301020123"></a>

<a id="canonical-3202300023230223-3031222312102033-1022033201133033-3220233232021003-2223321131310012-1223111311003232-1032222111130311-0313313101301110"></a>

#### `vn_config.dc_cluster_group_outside_vn.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2030233232330121-3021000323212233-3130203033023301-2002123331133020-3231012331322223-2232220231123030-0003012110232021-3313320111230212"></a>

<a id="canonical-1303013322333013-1113111032010103-2000300120113000-0031133110101133-3213103113132200-0103211330220110-3210021321300111-3002021220211132"></a>

#### `vn_config.dc_cluster_group_outside_vn.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.global_network_list

<a id="canonical-0121312111011012-3323133333133203-1010313033013332-2332300333213313-2110102210211122-3310120233201320-1122212200301130-3312101302212101"></a>

Type: `"single"`. Computed.

Global Network Connection List. List of global network connections.

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

<a id="canonical-1201011121202212-3223311321133203-3133313100033100-2013001200123302-2303002230301300-3130312213111233-2133231110222231-3311213202020230"></a>

### Direct properties for `vn_config.global_network_list`

- [global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102): complete subsection reference.

<a id="canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203)
- vn_config.global_network_list.global_network_connections

<a id="canonical-2322021111233212-0103333122321321-2110103100030123-2202020033001310-1112003010320323-1132131322232010-3020310032021301-1121132212032210"></a>

Type: `"list"`. Computed.

Global Network Connections. Global network connections.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2121021233221123-0232012110331333-0223003233023013-0310010130222233-2202123312221102-0320032120032012-1013302121011120-1223223121331300"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections`

- [sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-3301213310203212-0012032002200221-1021223221011303-1323023311033121-3132302231211122-0332110111033013-3311211032311330-2330321123103220): complete subsection reference.

- [slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2302130110021013-0111211003023113-2133230322321021-3032233303011311-1313032003302302-3333221023332013-2233332203000023-3331203010110131): complete subsection reference.

<a id="canonical-3301213310203212-0012032002200221-1021223221011303-1323023311033121-3132302231211122-0332110111033013-3311211032311330-2330321123103220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.sli_to_global_dr` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr

<a id="canonical-1101003112123200-3223021222301310-1022211100330111-2301311020332022-1203331121111333-1100122031031310-0001311321331112-1200222131020302"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-2031210132023130-2101322233132203-0030221032303013-1221020112031131-1321331232323003-3203223222331012-0002222310332031-0232010032122220"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.sli_to_global_dr`

- [global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-3131230122210212-0111123013313322-2110022011002321-2101313121211013-0110021312112301-1111231232230101-1103011331110100-1230301023111130): complete subsection reference.

<a id="canonical-3131230122210212-0111123013313322-2110022011002321-2101313121211013-0110021312112301-1111231232230101-1103011331110100-1230301023111130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102)
- [vn_config.global_network_list.global_network_connections.sli_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-3301213310203212-0012032002200221-1021223221011303-1323023311033121-3132302231211122-0332110111033013-3311211032311330-2330321123103220)
- vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn

<a id="canonical-0201011223203322-1333212230332110-1100022033123123-2023310321030311-3111020120022132-1102313123011130-3203022033021010-3110201233122122"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1002112200222311-0312013130301211-0132213232322022-1312012213102330-2212010022300103-0101201121323200-2122333000231321-0010030201222110"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn`

<a id="canonical-0221321122023031-3213202301300223-0013212131313212-0101222133212203-2103110312222000-2031230132212200-1013130222322031-2101230020220000"></a>

#### `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0113020102130003-3212212111323213-3301233112102122-2020113020112031-3202202023231213-3121031223112301-0311303311133332-0021131310033310"></a>

<a id="canonical-3201102323311020-1201112123222110-2020100300131303-3313113122322020-2300113200232023-2320030300122203-1312230011020232-1221003021333001"></a>

#### `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3222221221220023-0103312111013002-3301213202312222-0330322202103012-3312311122310023-0100031212110232-2103333211123030-0233121100021111"></a>

<a id="canonical-1230003303023312-1222001123103021-2002312020130331-1113030233031100-2021133122202102-2103111332120132-3211110112120232-0321131300122210"></a>

#### `vn_config.global_network_list.global_network_connections.sli_to_global_dr.global_vn.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2302130110021013-0111211003023113-2133230322321021-3032233303011311-1313032003302302-3333221023332013-2233332203000023-3331203010110131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.slo_to_global_dr` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr

<a id="canonical-2230110122313132-3311303310031123-2320112113023110-2320023100231120-1000323021131230-3033212213033033-0103021030103032-3120330230222321"></a>

Type: `"single"`. Computed.

Global network reference for direct connection.

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

<a id="canonical-2121322120333111-3111221320110333-3022330131030323-0100212113011203-3013030011111300-0330201232131200-2013212003030230-2211030022011200"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.slo_to_global_dr`

- [global_vn](data-sources--aws_tgw_site--reference--group-003.md#canonical-3312311333112200-3212101233230212-1210203120332003-0033230120132233-1311201310021232-2032002031321231-1300112211000230-0211323311110221): complete subsection reference.

<a id="canonical-3312311333112200-3212101233230212-1210203120332003-0033230120132233-1311201310021232-2032002031321231-1300112211000230-0211323311110221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.global_network_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-1102112012302300-3110320022302020-2101103231133230-2233301202100003-1121103303002200-2133212012023003-1310000112221333-1332021033111203)
- [vn_config.global_network_list.global_network_connections](data-sources--aws_tgw_site--reference--group-003.md#canonical-3232120200123113-1320230001313313-3102221313212003-1210311123333310-0112111101013103-0020001231321012-2110103023031113-2112210130220102)
- [vn_config.global_network_list.global_network_connections.slo_to_global_dr](data-sources--aws_tgw_site--reference--group-003.md#canonical-2302130110021013-0111211003023113-2133230322321021-3032233303011311-1313032003302302-3333221023332013-2233332203000023-3331203010110131)
- vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn

<a id="canonical-0113203233312220-0100030312221122-3001023301121130-1033200222312220-3011111021030330-2301210302323211-0330012003201011-2033230222202331"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0212313120231123-1101202300301121-3101101230302230-2320231333033033-0102021221112301-1001100121313120-2113101301113130-1311132032010202"></a>

### Direct properties for `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn`

<a id="canonical-0113021023301300-1030223201011201-1232100320312021-2023121230122031-0300220002113113-2211223012110232-2102100013323221-2210211322003111"></a>

#### `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2112032332220023-1133031330230200-2131131332031222-3321133233033313-1233312303200001-2020302202123300-2331030020312021-0220311203031123"></a>

<a id="canonical-2232212001323331-2313221002113330-0220332110222331-3003233220033023-0000033010210030-3222032032032100-0011101313202030-3201010310103320"></a>

#### `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3233131200300320-1112232031133023-0123022322000131-1101133331000112-2201120100330211-0110301223021213-1012033302032312-2133031021122133"></a>

<a id="canonical-3021313130031330-3113300010110002-2030112200010212-1201233023010102-1120211213323011-3112220320120202-2002333232022001-2232133311231010"></a>

#### `vn_config.global_network_list.global_network_connections.slo_to_global_dr.global_vn.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.inside_static_routes

<a id="canonical-1313032111231222-1231012302001313-3202220000221030-1033322000331202-0030110100210113-1122200020232300-3332323330233113-1122230121113222"></a>

Type: `"single"`. Computed.

Configuration parameter for inside static routes.

Additional upstream details:

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

<a id="canonical-3011310300223021-0331233220020301-0200221000321110-3200002021300122-1032133300110203-1103022300221131-3332210221102113-0211030111032110"></a>

### Direct properties for `vn_config.inside_static_routes`

- [static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213): complete subsection reference.

<a id="canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- vn_config.inside_static_routes.static_route_list

<a id="canonical-0320002011302122-2122010312320313-0213313220322210-3123112210233132-3032302102133103-3002001323232200-0203023131333333-2302210102122012"></a>

Type: `"list"`. Computed.

List of Static Routes. List of Static routes.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0032211111111031-3332231110220122-1313110022110210-0032331011212030-2332102200101223-3311032102232323-2320231033011310-2021021010002323"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list`

- [custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323): complete subsection reference.

<a id="canonical-0000022032200223-2111110233201011-2121311132230210-1220210010033012-0202102223121130-3002010010313330-3021123333002101-2020120022310312"></a>

<a id="canonical-1121211223221032-2221311002001210-0120320013021111-3013332131023302-1010302003000112-3320121300230300-0001321321232031-1213100201021001"></a>

#### `vn_config.inside_static_routes.static_route_list.simple_static_route` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- vn_config.inside_static_routes.static_route_list.custom_static_route

<a id="canonical-2310021111233213-1012310201322132-3112230020320023-0311220010023201-3133100023302210-2012120011101133-3213300102332120-3022221302111111"></a>

Type: `"single"`. Computed.

Defines a static route, configuring a list of prefixes and a next-hop to be used for them.

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

<a id="canonical-1012113302303302-2232122321332222-0231130223210221-1100313033121232-2311132222010231-2022212211223131-0133133301100101-1201122121012211"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route`

<a id="canonical-1332000033103023-0230330202313023-0022230300202122-1233322022202330-2300211220020133-0310223320100130-1231013301031101-1112213221232021"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.attrs` property

Type: `["list", "string"]`. Computed.

\[Enum:
ROUTE\_ATTR\_NO\_OP|ROUTE\_ATTR\_ADVERTISE|ROUTE\_ATTR\_INSTALL\_HOST|ROUTE\_ATTR\_INSTALL\_FORWARDING|ROUTE\_ATTR\_MERGE\_ONLY\]
List of route attributes associated with the static route. Possible values are
\`ROUTE\_ATTR\_NO\_OP\`, \`ROUTE\_ATTR\_ADVERTISE\`, \`ROUTE\_ATTR\_INSTALL\_HOST\`,
\`ROUTE\_ATTR\_INSTALL\_FORWARDING\`, \`ROUTE\_ATTR\_MERGE\_ONLY\`. Defaults to
\`ROUTE\_ATTR\_NO\_OP\`.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [labels](data-sources--aws_tgw_site--reference--group-003.md#canonical-3221031013333310-1320331230020202-2021221310120222-0012130200202211-3103120232312010-3131123110330311-3223101233331031-1310032023000300): complete subsection reference.

- [nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310): complete subsection reference.

- [subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-3213201111321223-2310320130001123-3031233120100133-3302010200132020-0303203000003031-0202033321130321-1100223323112001-3013313000022331): complete subsection reference.

<a id="canonical-3221031013333310-1320331230020202-2021221310120222-0012130200202211-3103120232312010-3131123110330311-3223101233331031-1310032023000300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.labels` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- vn_config.inside_static_routes.static_route_list.custom_static_route.labels

<a id="canonical-2100122130313122-2012212301333201-1301201011012123-2010221001030233-0110020031123212-1101002212232030-1013231123033032-0022332311321213"></a>

Type: `"single"`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop

<a id="canonical-3120101033331320-1312230322103201-1232131023010323-1033111101322031-2001111302222330-2022311122330222-2300120120323232-1332003321333321"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1032022121300200-0221312333331202-1101200311200230-3023211002211202-1023030120031112-2322213100222323-2320323111123100-2312121323313303"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop`

- [interface](data-sources--aws_tgw_site--reference--group-003.md#canonical-3323010300231333-3023232100130130-3131100100023122-3313203210001331-1223011212133311-3312331100302002-1120133112212113-1121032311322212): complete subsection reference.

- [nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303): complete subsection reference.

<a id="canonical-0310302220123031-0202020003223230-0310021011311023-1330300102112121-0313232113220331-3230222321121310-0100231023002132-2033010033203310"></a>

<a id="canonical-2322202212132133-0311110321030021-2000002330100212-2310123130310010-3331001003002331-3321221131321220-1012111210022011-1013101022002121"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.type` property

Type: `"string"`. Computed.

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

<a id="canonical-3323010300231333-3023232100130130-3131100100023122-3313203210001331-1223011212133311-3312331100302002-1120133112212113-1121032311322212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface

<a id="canonical-2330222223012030-2332010113310222-3223122232213332-3112200212330220-0022123030011321-0130012103230330-2221121132002033-0302200211100330"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0320321220130103-3203010102133332-2321132200330100-2031113231321303-3130133002111023-0002113111132210-3102313003011230-0202321302012202"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface`

<a id="canonical-1222102310100101-2032210110121111-2331331223332223-2033131032112231-2110300231132011-2031032103213222-3230010322020003-0313200000331111"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3012113222122100-1002121132322213-2100123120121210-2331121121030101-0133121303023011-1202100122230032-1110112132100221-1120121133332113"></a>

<a id="canonical-2132322123200313-3020232301003320-0012023303202101-1110303013121021-3311120201232033-2333120322333211-2130322302202200-2312111232220030"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3101301103033322-0300201303200202-0303221210100322-0210121012123133-0111233130230332-3010102203312130-2112102002130103-2102032023022211"></a>

<a id="canonical-3132223010333212-3210221021032101-2021230212130230-2032332122213010-0333132133313322-3313100003312002-2122101300211311-1002002200023102"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.namespace` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1210303100333230-0132230002233103-2031322232320222-1111231113332012-1322121320302202-2331233301123011-1031113310012032-2010300230232110"></a>

<a id="canonical-3303103030221222-1020102120202033-2313220233330133-0223033303010110-1203030201000311-3120320321232310-3310012020001200-0133321132333101"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1310331122002021-0210331320003020-0120203013130322-2103223301012232-0320103021200320-0111123321301112-3011220003210203-3111032033311322"></a>

<a id="canonical-0210200133131110-1023031220103302-1312032210113021-1003313130220002-3101331301113103-3300330121030231-2210022201300201-2303211310002112"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.interface.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address

<a id="canonical-1013222313003032-0211213232202032-2133210222311210-0101222320223321-0200103130311203-1321203221111023-1102002221031300-2303302121200111"></a>

Type: `"single"`. Computed.

IP Address used to specify an IPv4 or IPv6 address.

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

<a id="canonical-3220113212323302-0020221323311003-0300302012101320-3232120331000113-2220201223011112-3232212231213130-0113320123220033-3010032113302222"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address`

- [dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-2123033331011132-3303300213332030-0211312133213013-3332120030330232-3212110002033012-2201323032321212-0020323221111230-3020332121120220): complete subsection reference.

- [IPv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-2131330103323012-0021133201303023-3121301010102013-3113133011110330-1321301213323003-2312033111213121-2331023032131001-1113130122132221): complete subsection reference.

- [IPv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-0000000112231222-3330033310023331-2223333210000322-0313220030103123-1333220221111231-0003030312231132-0030011332230010-0132122311333111): complete subsection reference.

<a id="canonical-2123033331011132-3303300213332030-0211312133213013-3332120030330232-3212110002033012-2201323032321212-0020323221111230-3020332121120220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack

<a id="canonical-3212110100222132-0321030100231331-2032121020130133-3103233323312021-2100102030201222-2000020223303032-2132322110221231-1013303111333221"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0301311122232031-1132131233112002-1221313030202313-1030111111201010-2200013310100010-2311120123102201-1310023002201312-1313101000310211"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack`

- [IPv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-3101303301332311-2323323032131002-1121302300002120-1001012121020130-0001100323032101-2120031232022130-0033231113210210-2202103301002333): complete subsection reference.

- [IPv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-2112333112100032-3220321332322032-1333230131232132-2320033303112321-3200032003202133-3230300223021203-3001300203300233-2312333310303210): complete subsection reference.

<a id="canonical-3101303301332311-2323323032131002-1121302300002120-1001012121020130-0001100323032101-2120031232022130-0033231113210210-2202103301002333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-2123033331011132-3303300213332030-0211312133213013-3332120030330232-3212110002033012-2201323032321212-0020323221111230-3020332121120220)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv4

<a id="canonical-3201323101310213-0201300303321103-1103100222221302-3121010213301012-0211121123001003-1223121103103112-2313331223011110-3022010212223203"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1101301212012033-3230310323131300-3330000021123302-2100210120330321-0332121332332120-0222011133300130-0221320220132120-3123222111322321"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4`

<a id="canonical-1120211020320310-2110311303323101-0110231123310302-0220213231001233-1002233322103222-1303100313231021-0302013122101030-0001111012222300"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2112333112100032-3220321332322032-1333230131232132-2320033303112321-3200032003202133-3230300223021203-3001300203300233-2312333310303210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack](data-sources--aws_tgw_site--reference--group-003.md#canonical-2123033331011132-3303300213332030-0211312133213013-3332120030330232-3212110002033012-2201323032321212-0020323221111230-3020332121120220)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.IPv6

<a id="canonical-2222101131023131-0103023101032330-0212010301002232-2332300101232131-0202111210120122-1222103020122111-0110210023222103-2231222300111011"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1113233011101010-0023022022213202-0101303333333021-1322203323111313-1020133121303211-0300022002002131-2202312013030312-3123311000223323"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6`

<a id="canonical-2011221310330123-1210030312201102-3312101301013302-1002232100201311-2011103131003023-2302031303021301-0310002100131221-3120121210111032"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.dual_stack.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2131330103323012-0021133201303023-3121301010102013-3113133011110330-1321301213323003-2312033111213121-2331023032131001-1113130122132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv4

<a id="canonical-3130333220113003-2202210202210003-0222123203223311-2033131211200003-1303203131220303-1213220002200232-3121101323110033-3031002120012333"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3232302203310113-0230122230131302-3032332331221330-3322220103332310-3320032312022302-0021102013132303-3331211023022311-2313110130131330"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4`

<a id="canonical-2021120302302210-3132013102020002-1301011202020203-1102231331101200-0110233300113000-3123132300300310-1101202221101023-3011031103120211"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv4.addr` property

Type: `"string"`. Computed.

IPv4 Address in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0000000112231222-3330033310023331-2223333210000322-0313220030103123-1333220221111231-0003030312231132-0030011332230010-0132122311333111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop](data-sources--aws_tgw_site--reference--group-003.md#canonical-0202202003001022-3312031221223310-3033102221010130-3310201100130010-0301013231120001-1210112310030312-3323302000000201-1203033001130310)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address](data-sources--aws_tgw_site--reference--group-003.md#canonical-2111303123300103-1213332213033313-2313320233031331-0301133130202000-2021020133001333-2302110301231111-0331221120033023-1300303313001303)
- vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.IPv6

<a id="canonical-0021203312233331-1033331332223313-2021310311210032-1110232020110231-0032020110222021-1200212320001022-2301020323230133-2330323103312211"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1211332300023323-1300323032130111-1032002210123032-1231323321211313-2211331123001211-1022113123332232-1103320310011221-3210210220220301"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6`

<a id="canonical-3132333302213220-2330322322111130-3132133333101111-2330320330130021-2001100331221231-1112210302020233-3331221313230032-3103020231013300"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.nexthop.nexthop_address.ipv6.addr` property

Type: `"string"`. Computed.

IPv6 Address in form of string. IPv6 address must be specified as hexadecimal numbers separated by
':' The address can be compacted by suppressing zeros e.g. '2001:db8:0:0:0:0:2:1' becomes
'2001:db8::2:1' or '2001:db8:0:0:0:2:0:0' becomes '2001:db8::2::'.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3213201111321223-2310320130001123-3031233120100133-3302010200132020-0303203000003031-0202033321130321-1100223323112001-3013313000022331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets

<a id="canonical-3200013132220230-3021210131110211-1022321232003030-2220221123232320-1032113230311233-3213220101322311-2101111013111131-0322202022033222"></a>

Type: `"list"`. Computed.

Subnets. List of route prefixes.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2210203030001102-0032011321103230-2030203231311102-2033231330023133-3132011033122222-1032032011131300-3330331310220000-2232311021102000"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets`

- [IPv4](data-sources--aws_tgw_site--reference--group-003.md#canonical-2222030313330311-1123131300302233-0330310203020111-0130331212321111-2213111222320332-1010323001000221-0211222220201002-1210201312303020): complete subsection reference.

- [IPv6](data-sources--aws_tgw_site--reference--group-003.md#canonical-2023222032010130-0203300030301303-0120132223213102-3020330032000002-0101201132030123-2333032133020303-2011100031103211-3100133231231032): complete subsection reference.

<a id="canonical-2222030313330311-1123131300302233-0330310203020111-0130331212321111-2213111222320332-1010323001000221-0211222220201002-1210201312303020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-3213201111321223-2310320130001123-3031233120100133-3302010200132020-0303203000003031-0202033321130321-1100223323112001-3013313000022331)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.IPv4

<a id="canonical-1302220333000102-1013212303331210-3312322222110110-1121313100321112-2120200232022013-0012202023103103-2302332302111001-0011212121321110"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0310030223223331-1131201223002001-0030203231011302-2313022130110203-1303022102231001-3220012233112033-1123210201122030-1103130231001212"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4`

<a id="canonical-1023220021310133-3233310121023203-1123313132021232-0102132310332101-1332220000231332-1221011221112103-2101210120012331-2013003002211301"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.plen` property

Type: `"number"`. Computed.

Prefix-length of the IPv4 subnet. Must be &lt;= 32.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1230311132220223-0302201033103122-1022300231232002-0221332131302202-3211030130211002-2121000322323221-0033020022032313-2010121110230322"></a>

<a id="canonical-1232211320020020-3330210000022312-3212113202332322-1013330012323220-2212011332121321-2200020202121213-2331230300133112-1111122031112102"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv4.prefix` property

Type: `"string"`. Computed.

Prefix part of the IPv4 subnet in string form with dot-decimal notation.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2023222032010130-0203300030301303-0120132223213102-3020330032000002-0101201132030123-2333032133020303-2011100031103211-3100133231231032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- [vn_config.inside_static_routes](data-sources--aws_tgw_site--reference--group-003.md#canonical-3201323003223211-1221213331033030-2111103323302310-0000112320110200-1211131221332221-0002101031233122-3232221000111311-3013213202022320)
- [vn_config.inside_static_routes.static_route_list](data-sources--aws_tgw_site--reference--group-003.md#canonical-0322013313230030-1212110010101111-1033022002232102-2131211223322020-0322132113023322-0223032021123023-0231000211302120-3212001213102213)
- [vn_config.inside_static_routes.static_route_list.custom_static_route](data-sources--aws_tgw_site--reference--group-003.md#canonical-3030030123312221-3202130102233013-0013331133212230-1011020112030333-3111000021331131-2103001121221103-2223332330113122-0200000201133323)
- [vn_config.inside_static_routes.static_route_list.custom_static_route.subnets](data-sources--aws_tgw_site--reference--group-003.md#canonical-3213201111321223-2310320130001123-3031233120100133-3302010200132020-0303203000003031-0202033321130321-1100223323112001-3013313000022331)
- vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.IPv6

<a id="canonical-1320103021133030-2201231130312121-0032202113321330-0202330322202001-3210310213031302-2012222312130101-2321101320330211-0313220100010030"></a>

Type: `"single"`. Computed.

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

<a id="canonical-0033330030132201-2020110200103121-0201130103032103-0121010122020010-1003312221121221-3330221121111233-3312013022130101-1002021121332023"></a>

### Direct properties for `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6`

<a id="canonical-2133213011200203-1011120033311032-3122211031300032-1212002331300023-3101102133233322-0331302000212311-0212230102122102-2220022003300131"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.plen` property

Type: `"number"`. Computed.

Prefix length of the IPv6 subnet. Must be &lt;= 128.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-0123002033213010-3232030101212011-0202001333301122-2300300202321230-0310222030121302-2313213230130133-3101222312320032-3103000220121200"></a>

<a id="canonical-0231331030302112-1300302303332320-3220302300103113-3033210112123103-1003221032022112-3331113330232001-1332113112223103-1113111320232031"></a>

#### `vn_config.inside_static_routes.static_route_list.custom_static_route.subnets.ipv6.prefix` property

Type: `"string"`. Computed.

Prefix part of the IPv6 subnet given in form of string. IPv6 address must be specified as
hexadecimal numbers separated by ':' e.g. '2001:db8:0:0:0:2:0:0' The address can be compacted by
suppressing zeros e.g. '2001:db8::2::'.

Additional upstream details:

IPv6 address must be specified as hexadecimal numbers separated by ':' e.g. "2001:db8:0:0:0:2:0:0"
The address can be compacted by suppressing zeros e.g. "2001:db8::2::"

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3322033323111233-2333300011113211-2303310012333330-1032323330033233-1133313111333130-1303303300213210-1222223033330233-0322021301022133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_dc_cluster_group` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.no_dc_cluster_group

<a id="canonical-1221132013211032-2133130332031123-0020022331000110-1020112200130302-2221330030123310-2200000332132201-3323133213320200-2002033322023101"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211231303000231-1332322333303110-2231330333122021-3022013113310302-0331223221023310-0321000310301123-1112130032320132-0012101113120333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_global_network` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.no_global_network

<a id="canonical-0232130030022032-1310310331201023-3330012201032223-2211221012303032-1302021301000203-0031201201302110-2112021232201230-0112130210311333"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133103000303110-2322132103102100-0111231000032223-1012202032122131-2132030320322112-2311233020023323-1010100332321223-0212021022322130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_inside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.no_inside_static_routes

<a id="canonical-3211213202330233-3012313121131230-2013211112102123-2031021122221202-1011013022110021-0322312213020332-1313233123312013-0320110211001120"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220121110310322-3310320322232321-2200001213032233-3220332331310131-1101133323003022-1221011103131111-2032233010211203-1023100311031333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vn_config.no_outside_static_routes` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md#canonical-2023320331133312-1100202000221021-1000031302123320-0023120213231300-3030001110003033-1212200103313203-1311302300223321-3312113211320012)
- [Property reference](data-sources--aws_tgw_site--reference--group-001.md#canonical-1212321311102123-0132200012312320-2312231211033230-2200200312230233-3132320011023221-0111320212132020-3002000002020330-2223302203200011)
- [vn_config](data-sources--aws_tgw_site--reference--group-003.md#canonical-0302230100311030-1210212010000200-3030120211220013-1202212101231321-1322201012322230-0300232300110210-1103133222223202-1011002112321330)
- vn_config.no_outside_static_routes

<a id="canonical-0103130211231220-3101133330300313-1003120032333320-0221210331301022-0033011320330133-0201022133331023-3012023002131111-1333013320013122"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.
