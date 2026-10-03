---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1223201021301030-3012012001232033-3122223321312103-0100213000130020-2203212130102212-2233221133313030-0113130111022001-1233033311312233"></a>

## Next pages — ethernet_interface / 022021303232 / 6

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012022213030222-0133113131331232-0321230102211322-3012311203003021-0233311312210013-0330022112033230-2233302201023013-2201020330122303"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config — ipv6_auto_config / 021222132332 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3301330003033013-0130203002212230-0323332112111110-2312230222113123-2001310200201103-2021222321121133-3012332220210230-2031002210020020"></a>

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

<a id="canonical-0300122031013232-1130300031011030-3220210311323033-2311022033210031-2120100330122330-2200032301003023-1223230302320123-1331133133321311"></a>

## Direct properties — ipv6_auto_config / 021222132332 / 3

- [host](resources--securemesh_site_v2--reference--group-009.md#canonical-0321003321123303-2020001011330203-1302122331203121-3000330300033100-0112110331233102-2321012212311103-0101231121331303-0131012012212223): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211): complete subsection reference.

<a id="canonical-1212202311001200-1321131231120112-1202210130211312-3313111230230112-3131112300100100-1333220312021023-0021222101301301-2012020230201233"></a>

## Next pages — ipv6_auto_config / 021222132332 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.host](resources--securemesh_site_v2--reference--group-009.md#canonical-0321003321123303-2020001011330203-1302122331203121-3000330300033100-0112110331233102-2321012212311103-0101231121331303-0131012012212223)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0321003321123303-2020001011330203-1302122331203121-3000330300033100-0112110331233102-2321012212311103-0101231121331303-0131012012212223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2112121203201021-0203000303113010-0301323122132222-3313300312331321-1020313201131212-3222303103032303-1332310113131231-2022121213231022"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.host — host / 231332101000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-2201033211321231-3113311032033220-3021111121020211-1131032001130212-1110120122022212-1020113112222333-2203120320113102-2120031003211213"></a>

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

<a id="canonical-2200232203100321-0211112113033200-1100133322332232-0111132132121222-1221331010302133-0311031132101120-1301111321012031-2020023311130032"></a>

## Direct properties — host / 231332101000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310312301313321-0311132231230100-0003301020100013-3230203122110112-1120311132202303-3123212321333111-0000001320021231-3030213322020321"></a>

## Next pages — host / 231332101000 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303102220111121-3201131023220131-3013010102030020-3030223201111132-2102321223312112-3103221002103212-1022013120301131-0103022323313131"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router — router / 022301330331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-0201232012002311-2130231221000113-2120010213312010-0103220303022022-3002232132202113-0132122111122023-3013122132121313-0002131111121220"></a>

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

<a id="canonical-1230121330021332-2322203222130230-1320030010122022-2203301203122031-1001023212131313-2222220320202132-2313223010122121-3311302000102330"></a>

## Direct properties — router / 022301330331 / 3

- [dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222): complete subsection reference.

<a id="canonical-3112310302200012-1312013030103133-0032102200320121-2233220312333233-3133220331032333-1231332200321100-3021113212111320-2222311111233002"></a>

<a id="canonical-1300200131110011-3111010302232012-0112213110031211-1032002200103010-3210131133231013-2323323103031312-0200023330120102-3302101321122211"></a>

## network_prefix property — router / 022301330331 / 4

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

- [stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001): complete subsection reference.

<a id="canonical-3112322010112032-0310111313112122-3000312201230002-1201203231203030-1211102113131330-3321212112111130-0002320201303223-3111022013113030"></a>

## Next pages — router / 022301330331 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3311301003203333-1223030123003101-1232011300102312-3030211031033210-2001022303232222-3302300212103123-3322302031010330-2122230033001330"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config — dns_config / 130311023202 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3102333320320033-3233023212222000-1200212212201233-0022123103331332-0222013210030320-1320323333231313-2011133232033312-1200323132202132"></a>

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

<a id="canonical-1212202120002130-3302000120102013-3011201033112302-2231031201013233-2210000200221131-1110332223333123-0001300020030203-1330321011203223"></a>

## Direct properties — dns_config / 130311023202 / 3

- [configured_list](resources--securemesh_site_v2--reference--group-009.md#canonical-0200200312311010-3033110122302033-0220002223230211-1031113203202022-2221221320102111-2333111123222013-0122020231010233-2011333233000231): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113): complete subsection reference.

<a id="canonical-1022220111223011-3112013131133001-0121100323011100-1020022022012220-2213333002103233-2123303002120310-3121102112032133-1222333101120311"></a>

## Next pages — dns_config / 130311023202 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list](resources--securemesh_site_v2--reference--group-009.md#canonical-0200200312311010-3033110122302033-0220002223230211-1031113203202022-2221221320102111-2333111123222013-0122020231010233-2011333233000231)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0200200312311010-3033110122302033-0220002223230211-1031113203202022-2221221320102111-2333111123222013-0122020231010233-2011333233000231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033133031300312-1122310003221213-3231012033310122-3123231003222330-3300113312013321-1221323000210022-0020332321203303-0131123202212232"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list — configured_list / 221110133010 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1210003001011221-0211013220321313-2133001130012232-2320322321231320-3322033200113231-1333200121122123-3121230011012300-3232103300232210"></a>

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

<a id="canonical-2203110132132201-2331130330220210-2201232331221332-3223011323023101-2302223210211113-3033001202132233-1001212023223133-0231300120110102"></a>

## Direct properties — configured_list / 221110133010 / 3

<a id="canonical-1131103100303010-2123120210311022-3312311330213233-0102320322203012-1211133230102122-0000020300002102-3220302102211101-3021201333130123"></a>

<a id="canonical-0131000332020212-1322010023332230-3100312002233001-3032332011213330-0003212213220131-3232103020222022-0003000101313221-3113200001313231"></a>

## dns_list property — configured_list / 221110133010 / 4

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

<a id="canonical-0020000310031121-3010011122130221-2011322222132222-2002002210031310-0021210203211221-1332201120013321-0010121002032023-2301130313131003"></a>

## Next pages — configured_list / 221110133010 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1332020003323130-1203002003131101-1301230330301310-2210110330203221-1311233011303320-1023232202233021-1112122102132211-1301220020133121"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns — local_dns / 132003203233 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1312000020330021-2322331231123223-2202023112101212-2312311323222020-2203233321110300-3102110200102200-1110000131200202-0201112002210213"></a>

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

<a id="canonical-3302003113032100-0032133101231213-0131002021010232-2231003233330300-1302231133210010-0313233322312123-2032220111232210-3003201212213010"></a>

## Direct properties — local_dns / 132003203233 / 3

<a id="canonical-0211231033202011-3010121111210010-2123132103100000-0002310012033300-2001313010212111-0111033111111131-3300002320030112-1320203212221133"></a>

<a id="canonical-3103000000130300-1202112301321120-3113201233202220-2021110021123213-3132213230031131-1223301310321320-0130211301103201-3010131012000301"></a>

## configured_address property — local_dns / 132003203233 / 4

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

- [first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2210003231223113-2313130032223123-0222123312121323-1201003123223231-2321312213211330-1123033233322320-1121202010313121-3331110132303303): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-0012003220002220-3012023001030333-1111310233023102-1123201232020221-2221132302022301-2232112211012210-2032022330033131-1130212223320013): complete subsection reference.

<a id="canonical-1332232003002001-0331213100020013-1213022231310333-3122201323102210-2102030322103221-1310210030322010-2131123103313023-3113112210101233"></a>

## Next pages — local_dns / 132003203233 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2210003231223113-2313130032223123-0222123312121323-1201003123223231-2321312213211330-1123033233322320-1121202010313121-3331110132303303)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--securemesh_site_v2--reference--group-009.md#canonical-0012003220002220-3012023001030333-1111310233023102-1123201232020221-2221132302022301-2232112211012210-2032022330033131-1130212223320013)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2210003231223113-2313130032223123-0222123312121323-1201003123223231-2321312213211330-1123033233322320-1121202010313121-3331110132303303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333113310110033-3202303132220220-3103320003301233-0010310103021032-0212310201010231-1030000002131030-2203023213313320-1020121023011010"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 023002202112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0231120231000021-2102203223220111-1222312311311102-3113003333321123-1022323020310322-2203131111110113-1331221200221111-0021133321302013"></a>

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

<a id="canonical-2203203020021311-0111120222133220-1023030230221313-2330222311023302-3233132120120233-1021220003220103-1212031230313332-3201102310201211"></a>

## Direct properties — first_address / 023002202112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213210230032230-2113303003212022-0130011132200202-2130230032000133-1012211313301232-3221231130013311-2210330331122131-3202232012112331"></a>

## Next pages — first_address / 023002202112 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0012003220002220-3012023001030333-1111310233023102-1123201232020221-2221132302022301-2232112211012210-2032022330033131-1130212223320013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021023001313303-1111210202120230-1012101330032122-1321101322210232-2310221122003033-3302033213021301-1123213230312202-2002202101132322"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 203201011120 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-009.md#canonical-1210021003221321-2122011211331000-3031212002021122-0002101233223202-3232301233301202-2303020223333110-2132111112303102-1212333310202222)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0102223222030322-0300001002030332-3101330202113322-0313332232320111-3102310000221132-1111213310031321-1300002120130031-3331021203023121"></a>

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

<a id="canonical-2222032203023201-1120132202021330-0221131122100132-1222022031331200-2320331331223313-2323211123122331-3121110021030000-1220331300100302"></a>

## Direct properties — last_address / 203201011120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2220210220213132-2232332012330033-0312102233333132-3011201320103233-3221131212223021-2220222303321202-2232330231222032-0300033102100122"></a>

## Next pages — last_address / 203201011120 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-009.md#canonical-0132121220003022-1031211233220320-3220131033313103-3111302302112123-0223200301300300-0230213031023000-3121113113123020-2003200102330113)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212322331212201-0031121302230302-3233032110331302-1102113031031003-0333103033112223-1312030311230303-0120301233022112-1103322002203300"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful — stateful / 133103012112 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0200111033123022-3303032120200111-3020101001121211-2102132012300321-3331221233220032-2331212133120030-1022321022010020-2133302100320311"></a>

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

<a id="canonical-0332012001302220-0200032301302313-1121323231213221-2220131302020103-0122132323320001-3303132213133303-2213122333000012-1032300330123331"></a>

## Direct properties — stateful / 133103012112 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-1331130023330022-0313020312223222-3220011123011230-1113111033123320-0330112312102221-0231223120301001-0100333102303022-1113232203002301): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-3303220011212021-3323220323103102-3301223321023200-2223121132222110-3230113221111001-2100232133100230-3131313330103002-1133103011111202): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-1121102023100211-3113023031103011-3031300003201213-1210110101121001-1330102113011003-2130010333013233-0321311033211001-1021320321100032): complete subsection reference.

<a id="canonical-2020323002330013-3000113132033331-3230322113231110-1033102113132220-0333213012321030-3033032032131321-1032221230231313-3003231231013133"></a>

<a id="canonical-3230303300113231-3103202122302002-0122302303300013-2230110122310200-0303302123030112-3332121220303020-2133202131100223-3213023212300133"></a>

## fixed_ip_map property — stateful / 133103012112 / 4

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-2300002323002333-1110013212111103-1332232213131020-1333031321101003-1020031100202111-3011032233000231-1223011203223230-3120021302103213): complete subsection reference.

<a id="canonical-0112003232330332-3203030112010321-3111033023102121-2103232001001103-0323013030320013-1322331002323012-0102130002330321-1120111213200122"></a>

## Next pages — stateful / 133103012112 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-1331130023330022-0313020312223222-3220011123011230-1113111033123320-0330112312102221-0231223120301001-0100333102303022-1113232203002301)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-3303220011212021-3323220323103102-3301223321023200-2223121132222110-3230113221111001-2100232133100230-3131313330103002-1133103011111202)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-1121102023100211-3113023031103011-3031300003201213-1210110101121001-1330102113011003-2130010333013233-0321311033211001-1021320321100032)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map](resources--securemesh_site_v2--reference--group-009.md#canonical-2300002323002333-1110013212111103-1332232213131020-1333031321101003-1020031100202111-3011032233000231-1223011203223230-3120021302103213)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1331130023330022-0313020312223222-3220011123011230-1113111033123320-0330112312102221-0231223120301001-0100333102303022-1113232203002301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103013331033101-3111101110312101-1131023110311313-0202022323210321-2200031212322010-2013112133100211-2210323131333203-1320020123232310"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 230122321101 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0120022000020210-0202233223021103-0032230232202003-2020002230123220-3322000201223301-3121303031101232-3101203021133012-1320330323101312"></a>

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

<a id="canonical-1001011103300122-2011130023102111-2110033231322303-3011301013211120-1102002003013223-1020101313101223-1332310223100201-2102113120332002"></a>

## Direct properties — automatic_from_end / 230122321101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2100202310102132-0223312120301210-1300231220222011-3333231132133012-0003031233122200-0301132012332312-1301120302203333-1221302301031130"></a>

## Next pages — automatic_from_end / 230122321101 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3303220011212021-3323220323103102-3301223321023200-2223121132222110-3230113221111001-2100232133100230-3131313330103002-1133103011111202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0321122121210000-0212101100032322-3301313233023213-1001023030022100-2002021103210230-1201113021312212-3003122022223023-0312213202011312"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 330001231333 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1302011320003320-1201032300021121-3222303010321103-0130112030200203-0110031030232310-0322233301101000-3321231200132003-1100113321112333"></a>

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

<a id="canonical-0220133031100202-3103301031220111-0321022113302112-1333211030222233-0102120230020322-2232110303100211-3322113330201021-1220122321110322"></a>

## Direct properties — automatic_from_start / 330001231333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233001133011220-0300233321330313-1020013022001333-3200100302133002-3031221120013320-0103300303202311-3232303233303312-0031002213131002"></a>

## Next pages — automatic_from_start / 330001231333 / 4

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1121102023100211-3113023031103011-3031300003201213-1210110101121001-1330102113011003-2130010333013233-0321311033211001-1021320321100032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010130123110220-1021001113000121-1121332302213012-3021031201000302-3223030313322000-2110010032101131-3112031112220213-3033333023211112"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 233303011021 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2132121012320112-1220011333023210-0301203033221131-2122301112220331-2113310000023232-1101322300311211-1310302003322103-0020230330211022"></a>

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

<a id="canonical-2300301113123213-1300200110003123-3003122312102331-2000320110231231-3020133233113002-2012120110111021-2303120301101223-1201333230023030"></a>

## Direct properties — dhcp_networks / 233303011021 / 3

<a id="canonical-3001333202110301-3120133213312111-1232032130310010-1210001323210312-3323022011000231-3330003001220002-0213102303021230-1102303122330133"></a>

<a id="canonical-1100001302311303-1032003220132312-0111121310003122-2120301022302300-2201222201102002-1032303132220123-3203110012330101-0133031000131111"></a>

## network_prefix property — dhcp_networks / 233303011021 / 4

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

<a id="canonical-3203011120133230-3312123333313020-0302233222010103-0122030131330010-3232330333032301-3032022302223230-3220312231021100-0132112230000121"></a>

<a id="canonical-1110131330301321-1220000001203220-0001213330230102-3000000100322211-1112223010320021-0030000223132132-3310103111320030-0211332331032012"></a>

## pool_settings property — dhcp_networks / 233303011021 / 5

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

- [pools](resources--securemesh_site_v2--reference--group-009.md#canonical-0132001103100020-1300221120330231-0200300213130023-1231201023123230-1201132023201102-3120012121023320-0012103323012113-2123133013311002): complete subsection reference.

<a id="canonical-0223213103322303-2110300310213230-0102131300003020-0113200121312211-3310213133230322-0302101121013303-1311300320230133-0023201030122002"></a>

## Next pages — dhcp_networks / 233303011021 / 6

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-009.md#canonical-0132001103100020-1300221120330231-0200300213130023-1231201023123230-1201132023201102-3120012121023320-0012103323012113-2123133013311002)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0132001103100020-1300221120330231-0200300213130023-1231201023123230-1201132023201102-3120012121023320-0012103323012113-2123133013311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020230133333202-1121213133313302-1100010131213101-0101302320210312-1330310321132302-2021003013333131-2333220302023131-0302120211313120"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 332323231201 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-1121102023100211-3113023031103011-3031300003201213-1210110101121001-1330102113011003-2130010333013233-0321311033211001-1021320321100032)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0100310030033301-0232312022102030-3003231000223302-2233122121123032-1022123131232122-2100221001312303-3122132223323333-1101312330202020"></a>

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

<a id="canonical-0122212100210032-0221101201013032-3112103000330233-1201211303100323-0002220220001021-0123231103322003-2202122231220033-0301332311300112"></a>

## Direct properties — pools / 332323231201 / 3

<a id="canonical-3331021123321333-2113200002021301-3101232123113202-2003313322233220-2200330213300122-0033111221110133-2211102031302110-3310230233322120"></a>

<a id="canonical-2320021132211001-3223312220202302-2000131121020002-0303112230121112-2003101310020033-0002022001300320-2101000020332022-3300103030323020"></a>

## end_ip property — pools / 332323231201 / 4

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

<a id="canonical-2231333110301212-1112031033232033-2102313023320000-0001023010011021-0110222330212330-3323320302101301-3330120012322312-0023323003013201"></a>

<a id="canonical-1001101321031330-2133011121222122-0203112201321131-1223112310131320-0233012222310132-1332301323313212-3002003013230311-1003211003002303"></a>

## start_ip property — pools / 332323231201 / 5

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

<a id="canonical-2323231202022103-1231000220211110-3112223113010011-2202211031003011-1133013012232100-3023023300132121-2203033112332213-3113221113213203"></a>

## Next pages — pools / 332323231201 / 6

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-1121102023100211-3113023031103011-3031300003201213-1210110101121001-1330102113011003-2130010333013233-0321311033211001-1021320321100032)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2300002323002333-1110013212111103-1332232213131020-1333031321101003-1020031100202111-3011032233000231-1223011203223230-3120021302103213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011103002133032-1121312101012331-0320023021123010-2300011130312112-1010121121221110-0011231010213333-1023021233013200-2121320331311333"></a>

## equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 323323231222 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-009.md#canonical-0130010230122210-0120213103010013-2111221300002310-0111131203321100-0232100021033223-3211111202030233-3132010021133331-0130033220213012)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-009.md#canonical-1011303120121001-0333333012300100-0200321132030212-0032120312032033-3132302212333321-2320121102333012-0022201311013032-3312010230313211)
- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2233110012002013-1122102333023013-0221321102001230-2112033101312120-1322011133203020-2301310333310332-3313231203021031-1303223210013210"></a>

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

<a id="canonical-1102312211002323-1021030131332331-2013130203212333-2013201332031322-2023030200031023-0330031022201213-0133311113203302-0013312312132001"></a>

## Direct properties — interface_ip_map / 323323231222 / 3

<a id="canonical-3313313321112213-2100312022321002-0020313100310322-0022232200002011-1333130212103213-1121223332130311-3021212303223332-2230330220121123"></a>

<a id="canonical-1032322110312303-2230303203301201-0202122220223200-1021001301030200-2003131103103230-3130200323220120-3112133303212231-3113231130213012"></a>

## interface_ip_map property — interface_ip_map / 323323231222 / 4

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

<a id="canonical-1113203010222110-3102211033213302-0033012330201010-3202223332130121-2221333102121112-3230303233131210-0012130320223030-3230000302332102"></a>

## Next pages — interface_ip_map / 323323231222 / 5

- [equinix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-009.md#canonical-3023322212003231-1310311302032003-2101133030101320-1001311021123220-3112302110030030-3113232310021021-0032022023303010-1001223000002001)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3323331120001023-2103203213002302-1100311012120223-3131032211023223-3322031201322212-2101212123020100-2001123011023213-0222232203101210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220321111212313-3202011110233132-3321202312131232-0322123121123002-0130123321133033-0020222332321032-2321011230300201-1030231210331200"></a>

## equinix.not_managed.node_list.interface_list.monitor — monitor / 311001312123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.monitor

<a id="canonical-3123220210331022-3010101113201000-3121011233323111-3023303031003031-1331001023222121-1111032300020330-2203132111013010-2102102201013302"></a>

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

<a id="canonical-0221100321012013-3113201223330222-0121302223121130-0001202122101002-2002230013122010-3333302202231033-2012113000201012-0010113323221113"></a>

## Direct properties — monitor / 311001312123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0012033333232133-0013313302020003-0003130322300232-2102233102000030-1031133232010022-0131020030030133-3320232233132001-0031132031331130"></a>

## Next pages — monitor / 311001312123 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1120332122331202-0100310320123121-0010230211321301-3033030022201102-2013202121322201-0330121012102232-3221023301222122-2031133113330102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200130230133100-2333022100203220-0102212203120210-1103033001320221-2311023310131001-3132320321230210-2200130000320023-1111130122221230"></a>

## equinix.not_managed.node_list.interface_list.monitor_disabled — monitor_disabled / 001202300211 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-1323000132000002-3210221311102101-2120102323032300-1301023322220123-1003022011301303-1231212100300032-1323331310032100-1102001313313211"></a>

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

<a id="canonical-0301030111010023-1322301311300230-3203322331221330-2001232131200301-2201332013013313-1320101031331233-1032320120122032-3202120002002131"></a>

## Direct properties — monitor_disabled / 001202300211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122220211121132-3333221302323122-0233113130202100-0133122213100333-2001033333213020-1302031211233330-2222100131210323-0200001113033200"></a>

## Next pages — monitor_disabled / 001202300211 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103332300100311-3100130002133033-0001221003202300-0302221031330331-0211322101212201-2212313232010221-3330101223213022-3033013113313233"></a>

## equinix.not_managed.node_list.interface_list.network_option — network_option / 220301101223 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.network_option

<a id="canonical-1010221002202122-2013302302210321-1302202032131031-0000132132213021-1011122013213320-1123031133210201-0132031121123020-3222001133203210"></a>

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

<a id="canonical-2202201332300222-1030000123020303-3110212220200201-3202303233011312-2032110120002310-1222320130202001-2122220210031133-2120331110100011"></a>

## Direct properties — network_option / 220301101223 / 3

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-009.md#canonical-0113332310320110-1203312101320020-0112003212303023-2222310003330312-0321022322031313-1312211031020023-3003012111001331-1010213022031233): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-009.md#canonical-1000212202030300-1330011200003200-2233102020233002-2013002222233031-2110313321122303-2311222002223301-3011220023310212-1022300231130101): complete subsection reference.

<a id="canonical-1100020222003130-2120020120322001-0002103222032323-1311203322202323-0312230021310030-0003112312303323-3012003212303000-3003122322223222"></a>

## Next pages — network_option / 220301101223 / 4

- [equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network](resources--securemesh_site_v2--reference--group-009.md#canonical-0113332310320110-1203312101320020-0112003212303023-2222310003330312-0321022322031313-1312211031020023-3003012111001331-1010213022031233)
- [equinix.not_managed.node_list.interface_list.network_option.site_local_network](resources--securemesh_site_v2--reference--group-009.md#canonical-1000212202030300-1330011200003200-2233102020233002-2013002222233031-2110313321122303-2311222002223301-3011220023310212-1022300231130101)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0113332310320110-1203312101320020-0112003212303023-2222310003330312-0321022322031313-1312211031020023-3003012111001331-1010213022031233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2313221030131211-0312310020020310-3110333220220010-2332033201030112-0320133302201103-2110201330112330-1002102323111130-0002333313102133"></a>

## equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network — site_local_inside_network / 210132233131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120)
- equinix.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-3320110011111103-3122113313333113-2112213101120330-2210120323303220-3031210120012010-2013001311010222-1032100113122022-0032320330313031"></a>

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

<a id="canonical-0012311110302132-0302312113321233-3321000330002210-2113000012211110-0030212022030301-2210021321311030-1011231121201223-2300301122013331"></a>

## Direct properties — site_local_inside_network / 210132233131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102230022211102-2000302021120230-0213211201010133-2120032321011232-2331303111211210-1022102003112320-3023130022123021-0202233312131000"></a>

## Next pages — site_local_inside_network / 210132233131 / 4

- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1000212202030300-1330011200003200-2233102020233002-2013002222233031-2110313321122303-2311222002223301-3011220023310212-1022300231130101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212123322003030-3100133131001202-3202032013320121-1201022233032310-3200221000330103-0322020213003133-0331120121331021-3301112301322213"></a>

## equinix.not_managed.node_list.interface_list.network_option.site_local_network — site_local_network / 200011031213 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120)
- equinix.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0221331210012223-0303032212310032-2310100011311020-3003331303000120-1111332120103001-1110333000001320-1122330212331302-0331110132303113"></a>

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

<a id="canonical-1110212211233010-2013200013011013-2121111022321000-0032102201311010-3200232102003333-3301113320201303-2013001102320130-1030003201331101"></a>

## Direct properties — site_local_network / 200011031213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021323210112313-3000030323312010-1011200333033332-2001223123203012-3223300201313030-3022010331310200-0323022222103211-0230222233302113"></a>

## Next pages — site_local_network / 200011031213 / 4

- [equinix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-009.md#canonical-1222323010131121-0131002310231222-2321230221223003-1031310022112032-0111302122303213-2032033311202322-2232230022220331-3232001031230120)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1113123221220321-3001202320101212-1013301013311033-1011012033233210-0322231000233322-3000000333010321-0300220022322333-3200122212120133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011010312203003-3130300110222223-0313312103201001-3122320000123112-1201030100110331-0001331200211110-1200330311210033-3011031323121310"></a>

## equinix.not_managed.node_list.interface_list.no_ipv4_address — no_ipv4_address / 321112231011 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-2110001100202103-1300033110120332-2030012110323213-3212033100011212-3111210101101310-1003010231320002-0233223000211301-1010111013012003"></a>

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

<a id="canonical-0222131102211030-3120211202322132-0223302133321210-0010113003231121-0002310312232131-2302321012122121-3302230303012033-1332123232002321"></a>

## Direct properties — no_ipv4_address / 321112231011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2013231321231231-3113110010213203-0112203223232030-0223312021002310-2130013132021031-1303131100100033-2001323000132310-2130230102033203"></a>

## Next pages — no_ipv4_address / 321112231011 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1211220320012111-1133221130312110-3030113322023323-3002020201110233-0000233011131032-0321210201001020-3002031131313100-0220310210201110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0011211200311231-3110111300220213-0330223021332232-2030210300133131-0012123310211302-0212200001103001-3030113210221001-2002130323330222"></a>

## equinix.not_managed.node_list.interface_list.no_ipv6_address — no_ipv6_address / 032120300311 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0201231231111331-0332012102033330-0310011113002302-3301231022023331-2100223102112121-1011223311322021-0223103022111033-2101113321012023"></a>

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

<a id="canonical-3233122231202112-1202013210302101-3201000022021110-3210312213012313-2300211121110300-0202321222313303-1011232132101331-0310221111020330"></a>

## Direct properties — no_ipv6_address / 032120300311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2323000010122323-1100213102113022-1222223231133120-3332033132323221-3111011233031202-2010110233011201-0133122032031102-0303211210113020"></a>

## Next pages — no_ipv6_address / 032120300311 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1000302322130123-3211101132112110-1100312012320333-2033331000011310-3203123313213200-2120322131323103-1133023202200222-1130131331211101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010001320121222-1132133221201323-2000111010133011-1300010133230312-1103032222102201-3201130112300201-1000031300303231-0232301022330130"></a>

## equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled — site_to_site_connectivity_interface_disabled / 200303231203 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3330002322023020-0123021100033113-1031311112331013-2031220213201112-2133313121210320-3210110030201222-1303333021333121-0222231310020233"></a>

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

<a id="canonical-1103213303300212-2323101310013011-1111001110112231-0031122221011013-1011110103231101-3010021323311213-3210033320002131-2012323322023013"></a>

## Direct properties — site_to_site_connectivity_interface_disabled / 200303231203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3300320231112321-2111011210112331-1322210210310122-1303321313211320-1023332320221023-2331231133200111-3221231303031132-0231121013213030"></a>

## Next pages — site_to_site_connectivity_interface_disabled / 200303231203 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2313031021311203-0012332220211021-2322000220201103-1311103011023020-3033310110123113-1223223231112210-0130200212132323-1300112200130222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202221332333323-2000200301020321-0320132132021313-2302201203321333-2220002133122313-1131121102121100-1000020322313321-2013323230012203"></a>

## equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled — site_to_site_connectivity_interface_enabled / 113103110321 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-1103203212302330-1311120102331123-3123000310121032-3022013333332112-3223232320223200-0200133220120101-3210321130122320-2111120032021312"></a>

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

<a id="canonical-2311110300132211-0112013030003011-0202031210123213-1321030222230230-1312112222202113-2301012321232021-0132330330012233-2121133010100302"></a>

## Direct properties — site_to_site_connectivity_interface_enabled / 113103110321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222002320112011-3133233221111010-0031101031033020-1133220130223210-1023313303201021-0122011030133200-2032210022320133-3320132123032132"></a>

## Next pages — site_to_site_connectivity_interface_enabled / 113103110321 / 4

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2021020013311333-2023023000123300-2103322212112131-3300003132202332-1001030123222102-3322101101033120-2333000113130302-3232232321120023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322001130012111-0232013211311122-3310211210133310-0021333112031331-3222303002120020-0331303231212333-1023331330321230-1001113032200221"></a>

## equinix.not_managed.node_list.interface_list.static_ip — static_ip / 021220333310 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.static_ip

<a id="canonical-3130203111323323-1300103101000300-3220103302000113-2233130102130032-2100331100332102-0222133230112023-1213101122233200-3323032002103313"></a>

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

<a id="canonical-0130030101201031-1023232313031030-2123120321010110-1323310130203331-0001212020120302-0011233122321120-2311012013330213-3210231320021123"></a>

## Direct properties — static_ip / 021220333310 / 3

<a id="canonical-3332131223113223-1002010111230311-2033222210012330-0131323100122331-2111232032000333-0011033020233103-0222100110120003-1131210011033301"></a>

<a id="canonical-2220013211012200-0220011121212002-0113232202222322-0301311322122120-0102000321310300-1220222233103001-3102300323033002-1130322200332122"></a>

## default_gw property — static_ip / 021220333310 / 4

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

<a id="canonical-2330111131003001-0112020112310010-1030330111032212-3312111221111030-0000312032212130-3111131110321032-0121322001000130-0212100301032232"></a>

<a id="canonical-2202200000033302-2233230300021130-0003212033122103-1122323001213332-3212230120030230-2213100102030003-1033111223130210-3211030111211323"></a>

## dns_server property — static_ip / 021220333310 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0303010201223333-3200021313323303-1330223132100030-2012010021321000-0200210131312033-2110131021123231-3102000333133033-0012022230021122"></a>

<a id="canonical-1303232220100002-3020212220020112-3301121000202111-0223013020233101-1303210010233032-2103001033323123-1201221031030331-3100321010110011"></a>

## ip_address property — static_ip / 021220333310 / 6

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

<a id="canonical-1330033023013011-0001133213310313-2333222321302132-2230100113020333-3223100211100111-0300221002222332-1231312010100322-3032011120033302"></a>

## Next pages — static_ip / 021220333310 / 7

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130111021330310-2133123030031300-0221233100223100-0211212222310003-0222333313032122-1321130011031023-2221123112003201-3303113231113103"></a>

## equinix.not_managed.node_list.interface_list.static_ipv6_address — static_ipv6_address / 221033022322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-1010212233221121-3321222002002021-1123113021133203-3022320330202222-2011031031103121-1211003223301112-0032211302010300-0023213022031031"></a>

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

<a id="canonical-3213010131100330-2123100112032213-2012211002311300-3131332233232011-3013311330001331-1031222012321202-0333030021013103-3021211112312200"></a>

## Direct properties — static_ipv6_address / 221033022322 / 3

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-0113100223111200-3203223302133220-3033103120233300-2202301213231002-3102203002031310-2101301321210023-0031101332003300-3133133032100000): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-2313213231331131-0020333010211021-3033130023322303-0023021303020333-3323210110002213-0312120210030301-3300210131021210-0120202232302202): complete subsection reference.

<a id="canonical-1322113331112123-1031032213202331-1001210130231022-1213202230031122-2321323233303300-1131213332130033-3113300320212010-0323333230203211"></a>

## Next pages — static_ipv6_address / 221033022322 / 4

- [equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-0113100223111200-3203223302133220-3033103120233300-2202301213231002-3102203002031310-2101301321210023-0031101332003300-3133133032100000)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip](resources--securemesh_site_v2--reference--group-009.md#canonical-2313213231331131-0020333010211021-3033130023322303-0023021303020333-3323210110002213-0312120210030301-3300210131021210-0120202232302202)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0113100223111200-3203223302133220-3033103120233300-2202301213231002-3102203002031310-2101301321210023-0031101332003300-3133133032100000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321112333210103-3310323121120200-0200231103023123-2211311223220232-3122310232331222-1203013233210012-0123000031113203-0312303200133011"></a>

## equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip — cluster_static_ip / 002000301312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330)
- equinix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2213320110021333-1302121323322230-1200231330203103-3331211312023113-0331311113033121-1003121222121300-0230023233203321-2231212310010203"></a>

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

<a id="canonical-3101232222331011-1301320101122102-3210233202103321-3310111021133330-3211332322113213-3022110101231030-1312120110311112-1023112023002311"></a>

## Direct properties — cluster_static_ip / 002000301312 / 3

<a id="canonical-2003003132222132-0130102031102003-2222301103011320-2210121232332123-3112323331022011-1200223230331202-1310122012121231-0230300200300322"></a>

<a id="canonical-3030212110223011-2221211310031131-3013000023110011-1313211022023321-1122002113101312-3202003032330310-1020120033132203-2103232320220003"></a>

## interface_ip_map property — cluster_static_ip / 002000301312 / 4

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

<a id="canonical-1122130120221031-2103111133032331-0010033022022232-3200230231203103-2021301223113321-0100023313031110-0010220101013100-0100012000331022"></a>

## Next pages — cluster_static_ip / 002000301312 / 5

- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2313213231331131-0020333010211021-3033130023322303-0023021303020333-3323210110002213-0312120210030301-3300210131021210-0120202232302202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210213010132030-2210310331323323-0201322231302222-2312120012010311-2100031303231320-3123100020001212-0002323212330212-2330010113011311"></a>

## equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 332312033122 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330)
- equinix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1012230203333221-2213220033312211-2230120122102221-2300232333102012-0100133010332022-2011130223032321-0030233213303220-2112312102320102"></a>

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

<a id="canonical-2120132031103200-3330311201101201-0011201202302123-1000133232132113-0300100320033110-1212203000320203-1031322121300202-0110203220303313"></a>

## Direct properties — node_static_ip / 332312033122 / 3

<a id="canonical-2320110331202132-0102111211302330-3311230312200320-0010223100310302-0021300301233231-3103033202012102-0322010222212333-1003113323232320"></a>

<a id="canonical-3012333112112300-3200010101332110-2001200203301120-3120102230030011-1221110003233120-1033010320100131-3100113210311102-3213312210333211"></a>

## default_gw property — node_static_ip / 332312033122 / 4

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

<a id="canonical-0021331023112300-1310120013113123-1222113310113203-1220310113013021-0111022230002103-2130301100211303-1233223303233222-1333313012332231"></a>

<a id="canonical-1332213022030111-3211232211112332-0332300202201033-2022101231103230-3332223310132233-1313333021032002-3333130132022112-1002301113130323"></a>

## dns_server property — node_static_ip / 332312033122 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-1021101110300120-0102200110121331-3232030002103223-0332320133120310-3021203222323210-3130320211132011-0123023123113212-0331122010222302"></a>

<a id="canonical-3202230132122210-1220102013030213-2011011123331012-2202232001223331-3130333303200121-1300210330302313-2211132001121302-1303023232330300"></a>

## ip_address property — node_static_ip / 332312033122 / 6

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

<a id="canonical-1201220322110222-2123321133133222-0322121131220322-0331210020232100-2332030203122131-2021031130130101-3130002222111111-1020003210033032"></a>

## Next pages — node_static_ip / 332312033122 / 7

- [equinix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-009.md#canonical-2100222011001113-1310320111302103-1201113123233031-0120121030312023-2301313031101303-2011123303232030-0310123322310000-2232110232312330)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1320201032311302-2010322011120020-2300202030302120-2310032210331223-3102333123002122-0011212110103133-2031023020012131-3013110312323121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310210203100011-3032010031023132-2230231231000011-2111010223031111-2003300020100100-3011221113311003-3011001332213221-0310022332213201"></a>

## equinix.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 230211313130 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [equinix](resources--securemesh_site_v2--reference--group-008.md#canonical-2332330220333003-0112232110213210-0301011120213331-0000032322213303-1310020022232313-3122100333322101-0013220013132310-3033223312203022)
- [equinix.not_managed](resources--securemesh_site_v2--reference--group-008.md#canonical-2021201112031032-1332002301131213-0033332113012332-3112130331312032-1120301300110231-0233232133211313-2000012123202221-0013133322312100)
- [equinix.not_managed.node_list](resources--securemesh_site_v2--reference--group-008.md#canonical-2002210330003020-2120230332133020-1331322202202111-0102302332202201-0231320030320101-0000132322233010-2331311331011133-1112100233200333)
- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- equinix.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-3210331030201113-1113322232221232-2111013111230011-0312100101033110-2220003233032221-0001130100212311-2302333222330123-0022330300003212"></a>

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

<a id="canonical-3202223211201033-3211212003023023-1120330313132313-3131131012203332-0303322212102213-3200321002212302-3310332001220013-2220022033011200"></a>

## Direct properties — vlan_interface / 230211313130 / 3

<a id="canonical-2111220301313103-0202332303202003-3222221303110323-2133220320012303-1103002331013012-0302211020300112-2011010333310322-1322222213223131"></a>

<a id="canonical-3332310112101333-2101203100032211-3212031212202111-2002300233213321-3221320120232020-2030300110023113-2033202132231022-0310121300021201"></a>

## device property — vlan_interface / 230211313130 / 4

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

<a id="canonical-3120013220030313-0202022133323300-2033230011020330-3320222010223212-2300213133311212-3221103113132120-2130200132133123-0232322133233001"></a>

<a id="canonical-2030101213102130-3232130213202331-3131113032213322-2121322330203301-0100120223321302-1220031331322120-0012330103023111-3121310323011320"></a>

## vlan_id property — vlan_interface / 230211313130 / 5

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

<a id="canonical-3122320110102310-2022001313200200-3333032023300032-3122221212033110-1203330110103321-1201123233221121-3012120302103130-0312013301101012"></a>

## Next pages — vlan_interface / 230211313130 / 6

- [equinix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-008.md#canonical-0220203221311322-0100000210112021-2203133202031013-3102133232230211-1301321320130221-3002132110301033-1123101223230320-0301300203220130)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3021301201330310-0230130120101310-2231211023332220-1302200202031013-3111011231130033-1300202330311330-3213031120313332-2203211111320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001200031313132-1313030312221322-3213020202223221-2300103002110000-0310200200102002-3313221203103123-0213210330330320-2320213131202112"></a>

## f5_proxy — f5_proxy / 031123113220 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- f5_proxy

<a id="canonical-0323202033202102-2203130323103033-2120333100022103-1031333331023101-0001233230123023-0302313331331233-0112221333211032-1203223313013302"></a>

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
f5_proxy = {}
```

<a id="canonical-3110102113002100-3300011310221221-3222002330320210-3212022002033110-3201030233223123-2133112322001002-3020201022212033-2013333102212213"></a>

## Direct properties — f5_proxy / 031123113220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133323110130000-1010021031301131-0322030312003110-0002002033033211-1303213030323003-3120002022320301-3121330233030123-3011211110232103"></a>

## Next pages — f5_proxy / 031123113220 / 4

- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310331202203013-2012020331201310-2001031331232211-3321001033013310-2301110330323033-0302100123312012-3030120310131211-2212021221331313"></a>

## gcp — gcp / 300011121220 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- gcp

<a id="canonical-3233111213210322-1302300132033222-1221222310302113-1332221020020031-1112303220213132-0002100303122120-0301031121331102-0321031110313331"></a>

Type: `"object"`. single nested block, Optional.

GCP Provider Type. GCP Provider Type.

Upstream description:

GCP Provider Type.

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
gcp {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302022202101301-1110322023210020-2323022000212232-2213103000230002-1322330223103313-2211220010223311-3131233001210312-0033310323323130"></a>

## Direct properties — gcp / 300011121220 / 3

- [not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101): complete subsection reference.

<a id="canonical-1003312202202130-2031021030121133-1232220302100021-2303023132201031-0003030300220022-0332000220221223-0010031022312212-0123303300322222"></a>

## Next pages — gcp / 300011121220 / 4

- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212122301302212-1022021211320112-1331022202121002-0313302232023231-3123021010133031-3020023233311011-0312203300133113-0212322203221231"></a>

## gcp.not_managed — not_managed / 023011222000 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- gcp.not_managed

<a id="canonical-1132021120011022-1120321333220211-1023331000303231-0323231023101330-2120211202102333-2020031212331013-1110313032222321-2331302113113303"></a>

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

<a id="canonical-0231000111333100-1232333313313130-2023132222311311-1300202310320323-1220002003210110-0013203103133032-3220131230311100-0110313313212203"></a>

## Direct properties — not_managed / 023011222000 / 3

- [node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330): complete subsection reference.

<a id="canonical-0210313213111332-3131301323210312-2020302331110113-3033311031023113-3022121130200311-0333202232232221-0220131310033021-0230010223303331"></a>

## Next pages — not_managed / 023011222000 / 4

- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332221212012321-0300321131323202-2132121330200312-2310202333131020-1221111303021213-2310113311020303-3110130101103103-0111210300303211"></a>

## gcp.not_managed.node_list — node_list / 300002320331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- gcp.not_managed.node_list

<a id="canonical-2012220023132312-0323332000121312-3202233012000122-2011202212031102-3030131210133111-1332033231213213-0032121012101030-1122302212232012"></a>

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

<a id="canonical-1033313332103130-1333222031210200-0223233033110113-1020222200020131-2122202221202310-3211001032301132-0302111322102111-2021011100333010"></a>

## Direct properties — node_list / 300002320331 / 3

<a id="canonical-3102111110111312-3200211333101330-1221200121330111-0012313300221030-0222211211112132-3310333332111233-2221013212021330-1211313031200202"></a>

<a id="canonical-1201021032003222-0201023122301132-1331232330300220-3232233320212323-0023330333223100-3110231333332221-3330303201111230-3001201303330330"></a>

## hostname property — node_list / 300002320331 / 4

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

- [interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331): complete subsection reference.

<a id="canonical-0302123330000010-0311331200033000-3322031101321200-1220311103333213-0110011322031032-3032233313212122-0223021120102333-3303000133330203"></a>

<a id="canonical-1031112000130311-1230202211201001-3023331110131220-1102032030032003-2112133130002030-0021233130203212-2212230121210203-1321000030223332"></a>

## public_ip property — node_list / 300002320331 / 5

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

<a id="canonical-1030021030033032-2102223202321131-0113333102333213-0201101330103102-1032120323010110-2223333202313000-0113032210332330-2313010203211322"></a>

<a id="canonical-2233300123110312-2331002221333110-0020121001220332-3002332111202213-3232231133233010-2200023101233303-2231101323033111-1021213011322110"></a>

## type property — node_list / 300002320331 / 6

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

<a id="canonical-2211033213110321-3022002302113023-2101033201010130-1012030212322211-1002023210003333-0220332001031003-1203303020213110-1332102212012112"></a>

## Next pages — node_list / 300002320331 / 7

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310120033012330-1213220200132321-0223333232310100-2102221122333310-1330332221032112-0033322233132120-2303120010301231-2222100200120103"></a>

## gcp.not_managed.node_list.interface_list — interface_list / 210023233032 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- gcp.not_managed.node_list.interface_list

<a id="canonical-1323313113202203-1202113333111123-0231133222302133-3310133323302022-0231211122200001-2310310110303000-2031331300213221-3100310032003030"></a>

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

<a id="canonical-0302001212221303-1013311313311320-0130131103002021-2031032130331002-1001301313012130-2322113001300302-2131031111112310-3131112303000021"></a>

## Direct properties — interface_list / 210023233032 / 3

- [bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002): complete subsection reference.

<a id="canonical-3312330221121222-3200302132110202-1231001301112003-1313112103113113-0001231312332312-0133012000203123-0131110302111113-1101123302131331"></a>

<a id="canonical-0021212131201201-2000231203113321-1333003013332311-1210212022112101-2122233230230033-3013102301321231-1000130001200230-3030203231311103"></a>

## description_spec property — interface_list / 210023233032 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-009.md#canonical-0033122033012121-3302331120212023-3323021011132313-1123200320201001-0033122103112313-3030132200003013-2110330032333033-1232013121003210): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-2211332222002230-0122303122321113-3201122320313333-3310131113000011-1101321120213101-1203213233003121-2033310231011313-1012230033113111): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301): complete subsection reference.

<a id="canonical-2320031232313202-2010022000323211-1331220032100130-1312012203213033-2223233122000313-0232101000232323-3120120013222012-3303310031131033"></a>

<a id="canonical-3221110130033203-0230300020223111-1222111231011223-3232111110313013-1032200321003133-1230130303010021-1322010300212230-2330123202210003"></a>

## is_management property — interface_list / 210023233032 / 5

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3323101032132122-2331001133323211-2223032300102223-2201102211302121-2212010210212023-0320100231321320-2311221000121331-0203323201230200"></a>

<a id="canonical-3130031123311113-3032121300001311-1303221011121011-1023013111021103-0300021013121113-1223023122300302-1213313333213311-0313322021012233"></a>

## is_primary property — interface_list / 210023233032 / 6

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-0133001220231023-3110021333021130-2223330321003100-0000300021222122-0203233321311321-2103232002123203-2321020232132221-1203132221120031"></a>

<a id="canonical-0333320003001103-1132311212312013-0213130103021130-2133311123100102-3320130321122320-0310132022130122-3013232221013210-1031221201100120"></a>

## labels property — interface_list / 210023233032 / 7

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

- [monitor](resources--securemesh_site_v2--reference--group-010.md#canonical-1211223311200301-2113213233010230-1203012301302120-1022300331103323-0211120332111200-1030031000001303-2222221212332011-1213312211122110): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-2110013100330201-0203313102212200-0032133203312013-0101212301101020-2313122133233302-2220111322101030-0222330312032021-2111000113010023): complete subsection reference.

<a id="canonical-3311222223203222-2022021033323201-2221200213101330-0112000022011110-2102220002322300-2310130021120003-0010301021033323-0311021022202323"></a>

<a id="canonical-0311210320203101-0020223002011212-1122331002300101-0001132130332121-1111322130123110-1101021301000033-2133223020323200-0312021222102001"></a>

## mtu property — interface_list / 210023233032 / 8

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

<a id="canonical-3113333132213021-0012022322311223-0233323303300101-1103211002321021-3232202220200301-2300031223020133-3312320212121010-0001113112312233"></a>

<a id="canonical-3132233231201203-3220013322113230-2312213301222001-3333020232001303-0101032103100003-3333110031333311-3302222030111033-1233321130311232"></a>

## name property — interface_list / 210023233032 / 9

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

- [network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1203323220122012-2201322220300230-0230002310100030-3330031323321311-0300313103011131-2300131311231110-3013203313301020-2320101230133033): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-3000103023231312-2301010323330333-0333003302000002-1322000323123232-2002220330300031-1331112001201102-2321333303303203-2001023303022112): complete subsection reference.

<a id="canonical-0121230123100333-2233032212020112-3321110230100023-0123013202322310-0330000203301310-3122020110200211-3222201023003220-0230002011111231"></a>

<a id="canonical-0020112113013220-1003313012131213-0113233213322331-2301213311112133-3233313013103232-0301130303201103-1031123212230223-2213032300323100"></a>

## priority property — interface_list / 210023233032 / 10

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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-1133201233031230-3333012130033031-0221332312123131-3003211133222313-1013122333012331-0100231220131122-1212222211103002-0131010321033230): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-010.md#canonical-0220100020332132-3221320131203333-1233013032031323-3222113210012223-1121032113121201-1001013201031031-1210302012100032-0323200312201112): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-1030120100013303-3022221311322001-3212222021021102-1131222030210221-2220333320001121-0123333131111212-3121032230202313-3133313121120333): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-2322033023330322-3002023133132222-1111333320023023-1223323000130112-0000221001032331-3310323132320101-1112120210112123-1112101101200111): complete subsection reference.

<a id="canonical-2032230020002231-1011020013212011-1022232130310111-3233321022101200-1210230213113020-3320131101200332-2301113001013100-3303232222213122"></a>

## Next pages — interface_list / 210023233032 / 11

- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- [gcp.not_managed.node_list.interface_list.dhcp_client](resources--securemesh_site_v2--reference--group-009.md#canonical-0033122033012121-3302331120212023-3323021011132313-1123200320201001-0033122103112313-3030132200003013-2110330032333033-1232013121003210)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [gcp.not_managed.node_list.interface_list.ethernet_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-2211332222002230-0122303122321113-3201122320313333-3310131113000011-1101321120213101-1203213233003121-2033310231011313-1012230033113111)
- [gcp.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-010.md#canonical-0030000233303003-0323012301000003-1232112033233111-1322002021231332-3203003130030132-3012232003101220-1320232111102200-3322001220201301)
- [gcp.not_managed.node_list.interface_list.monitor](resources--securemesh_site_v2--reference--group-010.md#canonical-1211223311200301-2113213233010230-1203012301302120-1022300331103323-0211120332111200-1030031000001303-2222221212332011-1213312211122110)
- [gcp.not_managed.node_list.interface_list.monitor_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-2110013100330201-0203313102212200-0032133203312013-0101212301101020-2313122133233302-2220111322101030-0222330312032021-2111000113010023)
- [gcp.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-010.md#canonical-0130223033333312-1121121211313103-3113202033331032-2233322113123200-0332010230120331-1010000111303301-0210131100330121-2332322223233223)
- [gcp.not_managed.node_list.interface_list.no_ipv4_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1203323220122012-2201322220300230-0230002310100030-3330031323321311-0300313103011131-2300131311231110-3013203313301020-2320101230133033)
- [gcp.not_managed.node_list.interface_list.no_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-3000103023231312-2301010323330333-0333003302000002-1322000323123232-2002220330300031-1331112001201102-2321333303303203-2001023303022112)
- [gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-010.md#canonical-1133201233031230-3333012130033031-0221332312123131-3003211133222313-1013122333012331-0100231220131122-1212222211103002-0131010321033230)
- [gcp.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-010.md#canonical-0220100020332132-3221320131203333-1233013032031323-3222113210012223-1121032113121201-1001013201031031-1210302012100032-0323200312201112)
- [gcp.not_managed.node_list.interface_list.static_ip](resources--securemesh_site_v2--reference--group-010.md#canonical-1030120100013303-3022221311322001-3212222021021102-1131222030210221-2220333320001121-0123333131111212-3121032230202313-3133313121120333)
- [gcp.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-010.md#canonical-2211000311220321-1300310212321311-1320211122001332-3133210132012012-0003101310213320-1320323100233021-0320003310222021-2200322311012310)
- [gcp.not_managed.node_list.interface_list.vlan_interface](resources--securemesh_site_v2--reference--group-010.md#canonical-2322033023330322-3002023133132222-1111333320023023-1223323000130112-0000221001032331-3310323132320101-1112120210112123-1112101101200111)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010200312032123-3221323033010132-1100001201310211-0030011200301222-2113013202230202-2003330033222030-1222003101211330-1220120310110203"></a>

## gcp.not_managed.node_list.interface_list.bond_interface — bond_interface / 113331122331 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.bond_interface

<a id="canonical-2001222321030122-3333010021131112-2330230331112230-1130012211032103-2301220012233021-2131333320132300-2132203310033322-1330033120021221"></a>

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

<a id="canonical-3200121301010101-1300321323012030-0231103001222323-3003203013013210-1230200013303112-2032002031120110-1320203120321210-0321003332133310"></a>

## Direct properties — bond_interface / 113331122331 / 3

- [active_backup](resources--securemesh_site_v2--reference--group-009.md#canonical-2233223120230031-2320103303230112-0131121110300032-1333112312010101-2233103303000103-0332322300202330-3313021111031333-0330011212131203): complete subsection reference.

<a id="canonical-3002201223003111-0011203100210201-0213333332111101-1331210331200321-1112230002122032-2011112013320310-3223211300022200-0312231013330132"></a>

<a id="canonical-0313210321131001-1322223220321222-0103312030200331-0112203030223210-2120020231033232-1320313220112311-1120311332033313-2010000230123010"></a>

## devices property — bond_interface / 113331122331 / 4

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

- [lacp](resources--securemesh_site_v2--reference--group-009.md#canonical-3210102320303311-3232301012312300-2010032031112122-2313033032010202-3121231320032320-2222111013233230-0112231103110233-1301013211022330): complete subsection reference.

<a id="canonical-3332101220301211-3201332233031312-2322332232110131-3222321130131302-3131102121131313-1223211123110333-0211311010200310-2122122003013321"></a>

<a id="canonical-0120312302100131-3310011111122033-2020000012001013-1122201331301013-1003021020321000-1323212023312312-0330112121030013-3301330100011330"></a>

## link_polling_interval property — bond_interface / 113331122331 / 5

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

<a id="canonical-1021322000122023-0221033322000222-0322330123012200-1132022121331212-0120303032103123-3130232132300301-1010100110010000-3000100011100012"></a>

<a id="canonical-3132323202200332-2123030003101303-1002310320000222-1221302321323210-3122220112310001-1103320112101013-3132301131301023-3100311022233013"></a>

## link_up_delay property — bond_interface / 113331122331 / 6

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

<a id="canonical-0020011012022332-1130333220033023-2213322122032200-1102300033223222-1010201223301222-3210003101313032-0021000000233222-1010132130013303"></a>

<a id="canonical-3001213123001302-0122201203313001-0100101032322100-1111323001103310-3112330110312221-2001012213311102-2032033131232123-3120230333223322"></a>

## name property — bond_interface / 113331122331 / 7

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

<a id="canonical-3013012100331130-2111011110002022-3201100122023200-0031221113123332-0122230020231112-3101221122313310-0201210211330122-1303121012012112"></a>

## Next pages — bond_interface / 113331122331 / 8

- [gcp.not_managed.node_list.interface_list.bond_interface.active_backup](resources--securemesh_site_v2--reference--group-009.md#canonical-2233223120230031-2320103303230112-0131121110300032-1333112312010101-2233103303000103-0332322300202330-3313021111031333-0330011212131203)
- [gcp.not_managed.node_list.interface_list.bond_interface.lacp](resources--securemesh_site_v2--reference--group-009.md#canonical-3210102320303311-3232301012312300-2010032031112122-2313033032010202-3121231320032320-2222111013233230-0112231103110233-1301013211022330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2233223120230031-2320103303230112-0131121110300032-1333112312010101-2233103303000103-0332322300202330-3313021111031333-0330011212131203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001313132030231-3322133232103201-3011132130210102-3311222332331211-3133202023221231-1001003000233232-2312320303033120-0121032130103033"></a>

## gcp.not_managed.node_list.interface_list.bond_interface.active_backup — active_backup / 013000003322 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- gcp.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-0103310101200112-3103110103223223-0232211211023323-3111333131210333-2113010031111133-3333202101022300-2211133112311230-2111033212030021"></a>

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

<a id="canonical-1001112301033311-0011113212311210-1003311023233031-1001232303031202-1022122122333202-0011131230130130-1312221203101211-2121133230012302"></a>

## Direct properties — active_backup / 013000003322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123001112211231-2122221101313312-0103122233111333-3213113323030100-2201030111311301-2111101130101213-2312033332210330-2213311313301102"></a>

## Next pages — active_backup / 013000003322 / 4

- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3210102320303311-3232301012312300-2010032031112122-2313033032010202-3121231320032320-2222111013233230-0112231103110233-1301013211022330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221113331023113-1103202013023013-2121200002013132-3311013032123113-1320022213010232-0233102120330031-3213100203002232-3102201201300210"></a>

## gcp.not_managed.node_list.interface_list.bond_interface.lacp — lacp / 230020312123 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- gcp.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-1031110200131023-2031130001331322-2113313211122031-3213222332211230-3210100211310222-1313033111331131-1000333031323131-3322122320202200"></a>

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

<a id="canonical-2233231323102030-3303100113203332-1113123121212200-0331202200331313-2232222233233130-1222323000122230-1120110213030020-2131201313113033"></a>

## Direct properties — lacp / 230020312123 / 3

<a id="canonical-3120000010132321-3322331103231003-3311300130100031-3113311202023321-2223132300203110-0221002021230313-1020200102210303-2333003123121003"></a>

<a id="canonical-1301113122130010-1231230002113302-1201211212101331-0000210233321300-0011302332301112-0002302012330030-1020131031232011-3203031233112113"></a>

## rate property — lacp / 230020312123 / 4

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

<a id="canonical-1310101023121103-2313101001330313-3310032021010123-1110221323212313-1311331013020133-2322013231202133-0203102321211001-0330332112320111"></a>

## Next pages — lacp / 230020312123 / 5

- [gcp.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-009.md#canonical-1312131101221311-2123312020201302-0323030202131210-3030102033211112-1113331330011011-1033012003223302-3330223011123101-2300121002010002)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0033122033012121-3302331120212023-3323021011132313-1123200320201001-0033122103112313-3030132200003013-2110330032333033-1232013121003210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002020222312213-0230132230221002-2313321223332003-0202133012230233-2230323330203321-3020020210102013-1130121120231213-2332233133322110"></a>

## gcp.not_managed.node_list.interface_list.dhcp_client — dhcp_client / 033210001030 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-3213131300211013-0031022002201033-1020323021211022-1002333123032201-0103102000312110-1123200113300020-1301111130120322-3001030031121230"></a>

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

<a id="canonical-3102220032000302-1330003221200332-3100111302321200-2233300221013310-2222220222213221-1001230313031100-1322302021112002-0101032321021003"></a>

## Direct properties — dhcp_client / 033210001030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1033201323113021-1113311230203000-2010010332302032-1210232113310301-1121033333320023-1202000323002322-3203212033120221-0112120213331212"></a>

## Next pages — dhcp_client / 033210001030 / 4

- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221203221012201-2210212122110011-2303012002321201-0030120230310030-1232121332311231-3013131201013102-3203230013220310-1232320001203113"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server — dhcp_server / 002023033023 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- gcp.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1202231222103132-2213223331320022-1111320230100300-0333302030203330-3230013310203011-1213333113103222-3220011132000132-1211200313212132"></a>

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

<a id="canonical-3332333210330110-3100021002123201-3122022020131133-3220333322203331-0323003313312223-1022002103133231-2011323023222123-2223112230203021"></a>

## Direct properties — dhcp_server / 002023033023 / 3

- [automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-3230201023201000-0031203031300322-0130003111203121-0221031120033200-2320100230310111-3203202120223003-2200122002201120-1130102300303323): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-3021112221211102-1120023232212220-0102020133000323-3312003300330210-2120230223212123-3221201200132033-1322200123220301-3212121021112120): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132): complete subsection reference.

<a id="canonical-0021201321301330-1022203301131303-2003200120030122-3003313111030333-2031130202021313-2032003323023021-1121200030201031-0023002131102003"></a>

<a id="canonical-3110320003120102-0302310101123330-1313321021312302-3021032030322012-1301103211230123-0131131021100100-3022000133100231-2022111322203313"></a>

## dhcp_option82_tag property — dhcp_server / 002023033023 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0021030113311002-1331300302000301-3010330030023303-0332110313233021-2222321313023121-2101201103030113-3002230231311030-3203120301123000"></a>

<a id="canonical-3233012101131323-3211311130031322-3331300312323310-2231313313003203-2210103322221313-3201123232200312-0022202211300323-1123002230210012"></a>

## fixed_ip_map property — dhcp_server / 002023033023 / 5

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-2111220011212201-2210211320212123-0002230223223032-0103310130311120-3030102221113221-3203102202021320-2333322310332131-2200113023033310): complete subsection reference.

<a id="canonical-2000110223223313-0300220230132101-1233221321230303-0033000213220010-1002122100212310-2131322202131322-1023030102132111-3101221123300031"></a>

## Next pages — dhcp_server / 002023033023 / 6

- [gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end](resources--securemesh_site_v2--reference--group-009.md#canonical-3230201023201000-0031203031300322-0130003111203121-0221031120033200-2320100230310111-3203202120223003-2200122002201120-1130102300303323)
- [gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start](resources--securemesh_site_v2--reference--group-009.md#canonical-3021112221211102-1120023232212220-0102020133000323-3312003300330210-2120230223212123-3221201200132033-1322200123220301-3212121021112120)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-009.md#canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132)
- [gcp.not_managed.node_list.interface_list.dhcp_server.interface_ip_map](resources--securemesh_site_v2--reference--group-010.md#canonical-2111220011212201-2210211320212123-0002230223223032-0103310130311120-3030102221113221-3203102202021320-2333322310332131-2200113023033310)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3230201023201000-0031203031300322-0130003111203121-0221031120033200-2320100230310111-3203202120223003-2200122002201120-1130102300303323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333033021020303-3211211022232212-2122301111301200-3200333302311101-0002323101210300-1231330330020233-3212330103213103-2102213202322332"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end — automatic_from_end / 232021103111 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0320320020011200-0310321333330223-2113000331220213-1033130131230110-3322330200013331-0001330333320133-0120211320223013-1100300002021323"></a>

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

<a id="canonical-3301300003033131-3013302303333201-2210221123322211-1033032232212223-2333212131020030-0202330111103211-2030331300310320-2021011122331212"></a>

## Direct properties — automatic_from_end / 232021103111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010311322032331-2020331332223333-2222002212113331-0020120221220012-2321212223110301-3223022210100100-2221020320002330-1230231203233101"></a>

## Next pages — automatic_from_end / 232021103111 / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-3021112221211102-1120023232212220-0102020133000323-3312003300330210-2120230223212123-3221201200132033-1322200123220301-3212121021112120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113120311010023-2203320101213303-1202010202300110-3320022222013212-2120332031003311-0303333122320030-2212323011121021-2300221133121233"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start — automatic_from_start / 223311030100 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-2222122000020320-2213023112023330-0332021300111131-0102100202221302-2133112023132003-2220201111133312-0122001310212332-0200122300031321"></a>

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

<a id="canonical-1212002222032013-3123330330330100-3210003102100113-3100201312020331-2331330312030200-3300100223202032-2103311122201132-3321302231313212"></a>

## Direct properties — automatic_from_start / 223311030100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011003303131120-2132333110020100-2222103013103121-3312202200101012-3320323103233001-2210110302322122-3121022002031001-1202202221102331"></a>

## Next pages — automatic_from_start / 223311030100 / 4

- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-2000303313200021-2122230100001312-2101331200232322-1030323003003300-3033030003021033-3101000321032210-1301211320233010-2311333111320132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211212103321322-3330023320011021-1101020101222120-3023111320311122-1022311213120331-0112111021123231-2233303012302232-2000212122030332"></a>

## gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks — dhcp_networks / 203031201131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [gcp](resources--securemesh_site_v2--reference--group-009.md#canonical-1232231101021302-2212123100200301-1321330010022103-1021211002121101-2301000320030230-0003030030322203-2001112121212002-3322301022001012)
- [gcp.not_managed](resources--securemesh_site_v2--reference--group-009.md#canonical-3103301333221321-1333033221102003-1011311100022010-1320300121122300-2001220211321132-2012111110102212-1100103021003211-1330120110000101)
- [gcp.not_managed.node_list](resources--securemesh_site_v2--reference--group-009.md#canonical-2311020113312212-3200122033003021-0311002021110300-3011011323130023-0310013010331013-2233120021113320-3202233332101122-3011321011103330)
- [gcp.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-009.md#canonical-1232103210002000-2313130012212130-3002033331000231-3011123223230203-0000121220200111-2320013101322303-1231010301030223-1131101031311331)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2103120333232032-0222310333022131-1232130203013230-3021323300303222-2012010001203230-2300010233120202-0222123312120020-2222303200012102"></a>

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

<a id="canonical-2012311123132131-1301121030102312-1033332301232133-0230120130330301-1320322120100212-1202133111213032-0300213333011100-2231010101301133"></a>

## Direct properties — dhcp_networks / 203031201131 / 3

<a id="canonical-2323103330202101-3313320003331123-1320131023002211-0302033213311332-3133201122013021-2021113110310033-0110020230310233-1300322203201130"></a>

<a id="canonical-0132220212320222-3003132020210122-0213100321220200-2303320003111202-0232320300311013-2232033011300223-2133302030202103-1310303102312331"></a>

## dgw_address property — dhcp_networks / 203031201131 / 4

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

<a id="canonical-0320232001133133-0211000120101202-3230113200010213-3031332203032303-3000222222030022-2033001100130012-2213120100020100-1233033313002112"></a>

<a id="canonical-2302011302120220-1231323320012133-2031013123100300-0213030223311212-2223021302303322-3133011223132132-0300001011001021-3312200013022323"></a>

## dns_address property — dhcp_networks / 203031201131 / 5

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

- [first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-0312003331312110-3111001100022001-2222300312023132-2122133211103120-0020032122130203-0112333013130301-2230113023320120-1221112300121112): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1311222001020112-2112133132212013-0031233202312313-2013003011202013-0320120020020030-2030111033010310-1202132011310200-1222312310232123): complete subsection reference.

<a id="canonical-1210102102322200-2111122012322032-0122020032122000-0210102023122211-2210012132213303-0201312311010302-3001233123132321-3022200321331011"></a>

<a id="canonical-2030220203113130-2122330312222332-1112321123112101-2132103113122120-2223221001122323-2020230211322131-1113123203013332-2120220230122213"></a>

## network_prefix property — dhcp_networks / 203031201131 / 6

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

<a id="canonical-0201302303301202-0013021300131023-3022300123320113-2101112122011000-1221300003212110-1313231020202032-3003200002223203-2230310100233100"></a>

<a id="canonical-3030023222301121-2002220102331031-1100210023211333-0212033300331202-0123333321103210-3010212013023330-0301312332203332-0111031213221120"></a>

## pool_settings property — dhcp_networks / 203031201131 / 7

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

- [pools](resources--securemesh_site_v2--reference--group-010.md#canonical-1230113202102212-1123120311220322-1031123313022131-0112023321111331-3301330310322223-0033132130330122-0331013200321001-1202011312012121): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-010.md#canonical-3131011201201112-3113332132333013-2033231312003100-1121111323123320-3120202232101232-1003132033323032-0332020332132112-3010301022003223): complete subsection reference.

<a id="canonical-0220010111230000-3212020312002202-3132010233330312-0303102212021322-1231110001231211-0223023001120332-0002000031003031-0221131201320133"></a>

## Next pages — dhcp_networks / 203031201131 / 8

- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address](resources--securemesh_site_v2--reference--group-009.md#canonical-0312003331312110-3111001100022001-2222300312023132-2122133211103120-0020032122130203-0112333013130301-2230113023320120-1221112300121112)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address](resources--securemesh_site_v2--reference--group-010.md#canonical-1311222001020112-2112133132212013-0031233202312313-2013003011202013-0320120020020030-2030111033010310-1202132011310200-1222312310232123)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools](resources--securemesh_site_v2--reference--group-010.md#canonical-1230113202102212-1123120311220322-1031123313022131-0112023321111331-3301330310322223-0033132130330122-0331013200321001-1202011312012121)
- [gcp.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw](resources--securemesh_site_v2--reference--group-010.md#canonical-3131011201201112-3113332132333013-2033231312003100-1121111323123320-3120202232101232-1003132033323032-0332020332132112-3010301022003223)
- [gcp.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-009.md#canonical-3313330022031132-3223201011202123-1333210120013030-0221103213121310-0113121203231303-2110113003211302-3033122200033020-1210000032110313)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0312003331312110-3111001100022001-2222300312023132-2122133211103120-0020032122130203-0112333013130301-2230113023320120-1221112300121112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
