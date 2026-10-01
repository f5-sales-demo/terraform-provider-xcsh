---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-3310131213102320-3211111103002003-2131110030121213-3023203031103301-2211021201301323-1031220332212001-2123111310333030-2233122223033220"></a>

## Next pages — dhcp_client / 310030132122 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330330321321303-0030111331123210-1313020011111002-0200221320103033-0230102331213111-0113022231211232-1301210303021013-3213010112033312"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server — dhcp_server / 301110332011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server

<a id="canonical-0003012010003301-0303133200122231-1203133232101123-1333200323001311-0322301320012132-0033102301301023-2232033012023013-2223003122032031"></a>

Type: `"single"`. Computed.

Configuration parameter for dhcp server.

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

<a id="canonical-0121301303012132-2003111313212333-0100112132012101-0301122311312201-1001232131010200-2320300302010211-2220312220122130-0100321300003023"></a>

## Direct properties — dhcp_server / 301110332011 / 3

- [automatic_from_end](data-sources--voltstack_site--reference--group-008.md#canonical-2010302031302011-3221000013030222-2022323203233202-0022023231311220-2333202122220232-0112210031311303-3030331202003001-1322301111102130): complete subsection reference.

- [automatic_from_start](data-sources--voltstack_site--reference--group-008.md#canonical-2331223313011131-3231002011333120-1321202301020332-3022032333131010-3300020102020221-2022221001210330-1001212122221333-3221110123222301): complete subsection reference.

- [dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330): complete subsection reference.

<a id="canonical-1000012212130310-2331101233102033-3222231011321101-3003113000213333-3223322220101321-1000321312001302-0133023200101131-2001123200301130"></a>

<a id="canonical-0320311131203020-3021001133000002-0012002112332223-0201003312002210-2212202111031320-0032102310033132-3312201200201220-1031120121100113"></a>

## dhcp_option82_tag property — dhcp_server / 301110332011 / 4

Type: `"string"`. Computed.

DHCP option 82 tag.

<a id="canonical-2222211232203123-1332233002023200-0110122022332011-2223020333013300-3300303331023223-2023203022311001-1033002313120102-1113030300231200"></a>

<a id="canonical-1303202213221113-3001223032033002-3103121020212202-1013022011311302-0323232312112220-2110213032331211-1310102310010120-0212301002233320"></a>

## fixed_ip_map property — dhcp_server / 301110332011 / 5

Type: `["map", "string"]`. Computed.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
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

- [interface_ip_map](data-sources--voltstack_site--reference--group-008.md#canonical-0323101230021222-1300302121231330-3312010210013121-3303230222212032-3301202032302232-2010103102202122-3332212331011222-0210320220001311): complete subsection reference.

<a id="canonical-3122033033002010-3002202230302301-1313100100312301-0100213110331333-1130131321333332-2232201021130213-3320001102103221-0122122000312033"></a>

## Next pages — dhcp_server / 301110332011 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end](data-sources--voltstack_site--reference--group-008.md#canonical-2010302031302011-3221000013030222-2022323203233202-0022023231311220-2333202122220232-0112210031311303-3030331202003001-1322301111102130)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start](data-sources--voltstack_site--reference--group-008.md#canonical-2331223313011131-3231002011333120-1321202301020332-3022032333131010-3300020102020221-2022221001210330-1001212122221333-3221110123222301)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map](data-sources--voltstack_site--reference--group-008.md#canonical-0323101230021222-1300302121231330-3312010210013121-3303230222212032-3301202032302232-2010103102202122-3332212331011222-0210320220001311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2010302031302011-3221000013030222-2022323203233202-0022023231311220-2333202122220232-0112210031311303-3030331202003001-1322301111102130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0323033311012103-1213311003323022-1210333332123203-3121031002321310-2300231122230012-1112302303001310-1232203002113332-1231013211032101"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end — automatic_from_end / 322320301202 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end

<a id="canonical-3111313000313123-3333310011310000-2131133303322110-2133032210022202-2210233232310211-2012231311010310-3330303233201211-0022123213232112"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1113333303113200-1332320032130001-3030120303132030-0012212122103022-0030332023030232-2023300033111003-2030202103230221-3132313032101020"></a>

## Direct properties — automatic_from_end / 322320301202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202122112123302-3333303012221223-3101331332111033-0011312013011031-1320212032233132-1030311330001310-2212202000132320-3323021132012133"></a>

## Next pages — automatic_from_end / 322320301202 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2331223313011131-3231002011333120-1321202301020332-3022032333131010-3300020102020221-2022221001210330-1001212122221333-3221110123222301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131333013200132-1110331333010300-3210200331312231-0010232132203333-2112220022323331-3123303003200021-1102213210233131-3330112102212220"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start — automatic_from_start / 202100233011 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start

<a id="canonical-2113323311123120-0231110312322000-0201113012312313-1020320102333320-2030101312222023-2033200121222123-0101332220103211-0330222132012002"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3010010230212012-1123131330100220-1121011020101030-0301320100231020-2133113322233102-0303311032123131-2233321231013231-1301031020232322"></a>

## Direct properties — automatic_from_start / 202100233011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002311313301321-0332132130330012-0012230302223001-3131123200101132-3031323332031332-1000320133200031-3100220300112011-0312121323331020"></a>

## Next pages — automatic_from_start / 202100233011 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1330003023323213-1003323312201103-0031023203100032-0013321122010332-0313112032120323-2033200312331212-1020110102021102-3230331232112323"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks — dhcp_networks / 123212331022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks

<a id="canonical-3131311220203222-1013132213311000-2010300132210121-2321003223333102-2033323030110013-2131231010230300-0230210333112203-1320303101112323"></a>

Type: `"list"`. Computed.

List of networks from which DHCP Server can allocate IPv4 Addresses.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2033321112330131-0322202030332102-3023210322032310-2233133031110312-1200320232033132-3123332000011013-1303233031130000-1201021112322103"></a>

## Direct properties — dhcp_networks / 123212331022 / 3

<a id="canonical-1011212203012313-2112303101113032-3132020121330113-1021223120033322-1221100202213221-3100211100331222-2013031013021332-2130122123212033"></a>

<a id="canonical-3130231323203001-2213030210312213-1133122003100211-2233331302201210-3032230111320211-0012223031311001-3011033302111123-0231131223202033"></a>

## dgw_address property — dhcp_networks / 123212331022 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Upstream description:

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0201302312120110-1300331200221303-2000200230030212-2310213330220302-3132201200032113-0112112030333200-1232103300310210-1321321231203333"></a>

<a id="canonical-1223102030331313-2112303111222310-2001301003102032-3012002202302013-2131012201210221-2222210230122312-2112201213123133-0331301301132100"></a>

## dns_address property — dhcp_networks / 123212331022 / 5

Type: `"string"`. Computed.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Upstream description:

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [first_address](data-sources--voltstack_site--reference--group-008.md#canonical-2030102211231102-2002001033303103-2331211321303002-2203210130020030-3220210012212230-0101000210323000-0132221220330323-3212111203333231): complete subsection reference.

- [last_address](data-sources--voltstack_site--reference--group-008.md#canonical-0100030232320100-1301120133322212-0131311231130202-3310220313123123-2101210132123101-1203300121113303-2120130031020333-2333303033213023): complete subsection reference.

<a id="canonical-3021231303320301-3230200132001201-0223010110303332-0210200112232132-1223302103012130-1333233233323213-0003010100320020-2232112120123102"></a>

<a id="canonical-3000031203323221-1202020310101223-0012311313200110-0331001233021232-0223202103003300-2232233113023313-0112300031233332-0130310233302101"></a>

## network_prefix property — dhcp_networks / 123212331022 / 6

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2133322222111213-1110013200001002-2121101021033112-0122313102232231-1333133211230321-2332021130012110-0202130121102123-3232311222331203"></a>

<a id="canonical-0033322131300202-1333222303201301-0332031313000103-0030100000321113-1032001002000122-2300022130223310-3121011031122333-0201310011131112"></a>

## pool_settings property — dhcp_networks / 123212331022 / 7

Type: `"string"`. Computed.

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

- [pools](data-sources--voltstack_site--reference--group-008.md#canonical-1223311223100030-2000311203012330-2022132002030230-0111330332202200-0323011112303111-2230221012302103-2032120032230320-3211011301132333): complete subsection reference.

- [same_as_dgw](data-sources--voltstack_site--reference--group-008.md#canonical-1213210111223013-1103212332230033-3302232031133310-1110220131011112-2120230221030200-2010331331112332-3101310113121312-1320230313120001): complete subsection reference.

<a id="canonical-2211322000130211-1010330212311123-1100023022221200-1013222113112000-3220221010333323-3030213121102202-2031103311221011-0210300231100011"></a>

## Next pages — dhcp_networks / 123212331022 / 8

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address](data-sources--voltstack_site--reference--group-008.md#canonical-2030102211231102-2002001033303103-2331211321303002-2203210130020030-3220210012212230-0101000210323000-0132221220330323-3212111203333231)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address](data-sources--voltstack_site--reference--group-008.md#canonical-0100030232320100-1301120133322212-0131311231130202-3310220313123123-2101210132123101-1203300121113303-2120130031020333-2333303033213023)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools](data-sources--voltstack_site--reference--group-008.md#canonical-1223311223100030-2000311203012330-2022132002030230-0111330332202200-0323011112303111-2230221012302103-2032120032230320-3211011301132333)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw](data-sources--voltstack_site--reference--group-008.md#canonical-1213210111223013-1103212332230033-3302232031133310-1110220131011112-2120230221030200-2010331331112332-3101310113121312-1320230313120001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2030102211231102-2002001033303103-2331211321303002-2203210130020030-3220210012212230-0101000210323000-0132221220330323-3212111203333231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3211133013113130-1031233100010322-1320300201300102-3331220111030100-1233011130102110-0221213221113301-1232230313311302-0131210233222203"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address — first_address / 023010211200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-3003312112100231-0012321002331123-1132213311330322-0231323121010120-1103123202202330-3103223033200100-3000302010112132-3322122201323121"></a>

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

<a id="canonical-3232000223112232-0022123210031102-2013123301233231-2021230301203333-2010103320120022-0130111120330330-3301132331101111-1313010303321331"></a>

## Direct properties — first_address / 023010211200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023000302010023-3313000011201321-0231221021131212-0000200220013300-1032303030001102-2311100130201201-0013000330003030-2332221230133122"></a>

## Next pages — first_address / 023010211200 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0100030232320100-1301120133322212-0131311231130202-3310220313123123-2101210132123101-1203300121113303-2120130031020333-2333303033213023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212221233102231-3131023223230203-3311303032102010-2230210230111101-2320213331033202-1102001131331123-3002103311231013-2123332221031022"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address — last_address / 231230200333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-2113223323322200-1213123030310033-0211013001322021-2311302321313102-2133322121133031-3310102132123122-1231310303112021-1013331132032223"></a>

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

<a id="canonical-0031022120303030-3032233301220211-3111132001322230-2123033333332331-0103200330002300-0120220223100222-0133203120123303-0032322132313323"></a>

## Direct properties — last_address / 231230200333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3312011122302133-3110333023002111-0112010202113102-2212100311221233-0001223300110003-0323103311102330-3130223230021233-0311331113121023"></a>

## Next pages — last_address / 231230200333 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1223311223100030-2000311203012330-2022132002030230-0111330332202200-0323011112303111-2230221012302103-2032120032230320-3211011301132333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021032130101133-1333112120203021-3322311211301020-0223131023311302-0333211211110203-0312200211012100-0320321320233331-1031100212011031"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools — pools / 310010331013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-0122032102000333-2130002201221020-3110310013300232-0302012010331001-0011022222102312-0301031012211131-2021211112202213-2211103020111330"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0203221032331030-1220123233301202-2030011232231113-3031222111130110-1123130133000311-3320213032020021-3320120302133301-2033213110122001"></a>

## Direct properties — pools / 310010331013 / 3

<a id="canonical-3323221230331133-1022033003132301-1222301213011113-2001132331213001-2301022210022232-3323112210333231-1301213323301231-2033210313000031"></a>

<a id="canonical-2322201112210103-2102231112023010-0221303211011300-3232313122323313-0110322100132110-0123231222031101-2320100010013030-0100102210001031"></a>

## end_ip property — pools / 310010331013 / 4

Type: `"string"`. Computed.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2201311201022032-2021331200320100-2121313312330031-2023022222201201-0013303230212132-1320201002222110-1223032220111301-2320320110131301"></a>

<a id="canonical-2202103221021023-0333202121332323-1130212102130301-1222132111222333-1110330131301320-3320102323012103-1132332111211210-1200232031032321"></a>

## exclude property — pools / 310010331013 / 5

Type: `"bool"`. Computed.

Exclude this address range from DHCP allocation.

<a id="canonical-1230022322232322-0333002120232201-1212013132313210-0101032032032331-2010320130302321-1313231223033110-1330213203222211-2211212023322001"></a>

<a id="canonical-2233130322303210-3001212003122003-2112201022020203-3300322303220020-0301310212300232-2313012212310221-2122220001112012-1033203203101311"></a>

## start_ip property — pools / 310010331013 / 6

Type: `"string"`. Computed.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3303123021221121-3033200013303333-2300133003310021-3102010302103012-1023312210131102-2111233102310021-0012233231002021-3032023200030213"></a>

## Next pages — pools / 310010331013 / 7

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1213210111223013-1103212332230033-3302232031133310-1110220131011112-2120230221030200-2010331331112332-3101310113121312-1320230313120001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3230100220311011-1332130330313112-3030332321122323-2100333210203301-2123132333302321-0212311202200100-3023020003031231-2233322030033230"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 310030000021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-2333023130101101-3301310220023200-3110033012002111-0000301212022133-3003032333031300-0220313003332213-2321103131320030-2103211110203220"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for same as dgw.

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

<a id="canonical-2131013221120112-2102122313233033-0201232121202301-3332212102010100-1210010222020212-1122010333031230-1203002011202233-1201220030310303"></a>

## Direct properties — same_as_dgw / 310030000021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0033312023100011-2100113121032012-1130212013012032-0232022211323313-1221101003302210-3332331310120210-3010012102211003-0122313232311230"></a>

## Next pages — same_as_dgw / 310030000021 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-3301313313312323-0302132013000010-3230301013311101-1101302300101321-1302003033012020-0013200123301111-0123131123011132-3102222101102330)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0323101230021222-1300302121231330-3312010210013121-3303230222212032-3301202032302232-2010103102202122-3332212331011222-0210320220001311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102022201012102-2223131213032333-3303210002012230-0200320031203303-3220001033333110-0301102210332231-1023213302331122-3331012023011001"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map — interface_ip_map / 222110123010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map

<a id="canonical-1113300200313200-1321323013213333-2012121122222333-3213021101302213-1302200001310033-0123312122103102-1010310002210233-2212111213002013"></a>

Type: `"single"`. Computed.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Upstream description:

Specify static IPv4 addresses per node.

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

<a id="canonical-3321123321330320-3121032131220333-1130331033020312-0211003101101102-3002332230031221-3202103132300320-1120023212023330-1200103000023011"></a>

## Direct properties — interface_ip_map / 222110123010 / 3

<a id="canonical-0122110220002221-3022322323221033-0021111101132320-1022100213310022-2020323321323323-0020103313102210-2212310033302123-2110203213000321"></a>

<a id="canonical-3230320111221002-2331013132302022-3331230332121122-2300021310032311-0111221111233331-3322310321100202-3011333000210022-0122122323233320"></a>

## interface_ip_map property — interface_ip_map / 222110123010 / 4

Type: `["map", "string"]`. Computed.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
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
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.ipv4": "true"
  }
}
```

<a id="canonical-1101233201303312-1003011133121120-2030112002002103-3312312323321202-2222212232212330-1231000012032021-0323233010332121-2320020303113332"></a>

## Next pages — interface_ip_map / 222110123010 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](data-sources--voltstack_site--reference--group-008.md#canonical-1000221322122120-0220110012033020-3203330033221300-2033211110311201-2221320013110233-1233321313330101-3213301111203122-3231133202210203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131102200130122-0330203111321032-1102332233012120-2110132012022300-2111001100110211-3033210302020001-2123310130300230-1333113021133213"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config — ipv6_auto_config / 213030222023 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config

<a id="canonical-2222331232113003-2322203032013323-1223100320002122-0230311123232331-2103200010130011-1113331121210332-0213113031023231-0301002123032323"></a>

Type: `"single"`. Computed.

IPV6AutoConfigType.

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

<a id="canonical-0102033020010113-2100223013113302-3221000231312203-2012311001303010-3122320032001220-2131230200233121-0331120101203021-3300211012110120"></a>

## Direct properties — ipv6_auto_config / 213030222023 / 3

- [host](data-sources--voltstack_site--reference--group-008.md#canonical-3332200203012232-3033302031231122-0032322012030323-1332302302100220-2330313201203101-2300031131223003-2133320111112033-0202232320113300): complete subsection reference.

- [router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331): complete subsection reference.

<a id="canonical-0133120023222012-0313122021031020-2230020101213231-3233200032322132-0110120010033101-1322002201203202-1013003101122113-3122101122303132"></a>

## Next pages — ipv6_auto_config / 213030222023 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.host](data-sources--voltstack_site--reference--group-008.md#canonical-3332200203012232-3033302031231122-0032322012030323-1332302302100220-2330313201203101-2300031131223003-2133320111112033-0202232320113300)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3332200203012232-3033302031231122-0032322012030323-1332302302100220-2330313201203101-2300031131223003-2133320111112033-0202232320113300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1131231133302123-2333122231033101-1220001012120210-1303233211223123-1023313302020212-2123102323011232-0123021033302302-2020313011030021"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.host — host / 022023102002 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.host

<a id="canonical-3102021112201003-1130233121321310-3330200202320130-0323022213331033-1020331212120202-0013113213021203-0100211202033003-3110100102301033"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-0011032032300033-0320330031023203-0022231210031033-1130123020333231-3331000210322222-0221323232331200-3030310220223013-0132223312132232"></a>

## Direct properties — host / 022023102002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133321322102220-1120023013302032-0011330210300320-2313200333010112-0103023032003223-2132330323023011-0002122030202221-2213231333210303"></a>

## Next pages — host / 022023102002 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221112020122201-3010003100230202-3300322010330211-1303300012321122-1332010301031333-2110002100303031-1103030320131230-0220310221013310"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router — router / 300032203111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router

<a id="canonical-3232222201021032-0103101103311102-1013031101000030-2102333310110021-3323123200113233-1112210002330033-2330100333320312-2332200213212101"></a>

Type: `"single"`. Computed.

IPV6AutoConfigRouterType.

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

<a id="canonical-1021202023123213-0011013101110221-2122130022310023-1220202012230303-3001021100130123-0200112332013213-0330313321131020-0100221322333333"></a>

## Direct properties — router / 300032203111 / 3

- [dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011): complete subsection reference.

<a id="canonical-0221222011333322-3330011200330113-1322012302312303-1103022323310002-1011321030320023-1120121003210321-1110132132220022-1233233302103230"></a>

<a id="canonical-2322323032022122-3330032111330221-1110110033000000-1113022122021021-2232032230212010-2013333020032220-0232130330122012-1221021211331220"></a>

## network_prefix property — router / 300032203111 / 4

Type: `"string"`. Computed.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Upstream description:

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012): complete subsection reference.

<a id="canonical-1330011302322022-1123300330030111-2311210023032110-3130030301233021-2121010331033033-3121111003231030-2132221131101000-3202313220112223"></a>

## Next pages — router / 300032203111 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013022132133233-2231323301231212-3303110310203200-1201130320030122-0230112012323200-1230313221320031-3231131020112112-1332120031302311"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config — dns_config / 131021001300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config

<a id="canonical-1000122003303101-3230102112003233-0220100202131111-3122333330222111-3323121113200332-2011133213131323-1203300122112201-2301203233300122"></a>

Type: `"single"`. Computed.

IPV6DnsConfig.

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

<a id="canonical-0103203220200321-1203332320003221-0333330023212012-3321130203201020-0310011321331311-2112233321010223-3312313331131002-2112100223031102"></a>

## Direct properties — dns_config / 131021001300 / 3

- [configured_list](data-sources--voltstack_site--reference--group-008.md#canonical-2120020302001100-3232002121301313-0120202230121103-0330121203310310-3121001131231213-3223231212120001-3230020330033102-3021333201222013): complete subsection reference.

- [local_dns](data-sources--voltstack_site--reference--group-008.md#canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203): complete subsection reference.

<a id="canonical-1210231013233003-3031121120331230-1320012112320212-2223221032002311-1010222130000210-3223312230131110-3002003001103211-2030320213000032"></a>

## Next pages — dns_config / 131021001300 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list](data-sources--voltstack_site--reference--group-008.md#canonical-2120020302001100-3232002121301313-0120202230121103-0330121203310310-3121001131231213-3223231212120001-3230020330033102-3021333201222013)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-008.md#canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2120020302001100-3232002121301313-0120202230121103-0330121203310310-3121001131231213-3223231212120001-3230020330033102-3021333201222013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221332203001100-3222221002021213-1132213113300132-2121120132012313-1122113321131101-1320012331111120-0011213300223001-0311103200023333"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list — configured_list / 231131031132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3212212320110012-1313321012002231-1212322002211210-2302213103021302-3012301011301213-0133232310213030-0122002130220332-0333133321301313"></a>

Type: `"single"`. Computed.

IPV6DnsList.

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

<a id="canonical-1033001300320230-3030233020301023-1221102011132220-3211012220122211-1301232333331232-3332231303111132-2210200023220111-0002131030102110"></a>

## Direct properties — configured_list / 231131031132 / 3

<a id="canonical-0313201202322020-0232011122331010-0212032022321022-3233000131011130-0102200123003032-2122320000010321-1100133201133321-2203211222213120"></a>

<a id="canonical-1100022301321011-3010012220303220-2220112121303011-2113313122233211-1113313020321011-1112333012230220-2310302223203332-1332303110032302"></a>

## dns_list property — configured_list / 231131031132 / 4

Type: `["list", "string"]`. Computed.

List of IPv6 Addresses acting as DNS servers.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1022010000222030-0313300302312302-1303311023210001-3332313012110212-0231100321310022-2223031202202103-2310220002300301-2123333220123211"></a>

## Next pages — configured_list / 231131031132 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331221212110130-0310210203211020-3222131231130321-3232301100132221-2031111212101020-0003011221213002-3122010023322113-2231200113101010"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns — local_dns / 110030010100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1223312201231230-2011103001333013-3323300102231021-1212113113112200-2312230023020220-0121030101302120-0121100003332320-3200300233202310"></a>

Type: `"single"`. Computed.

IPV6LocalDnsAddress.

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

<a id="canonical-0300010333302022-2121123313011322-2211301223012020-1322121230121300-2222132310320031-0301003212323233-2013323211110332-2331320233102003"></a>

## Direct properties — local_dns / 110030010100 / 3

<a id="canonical-0200122000103031-3231101320021320-1021123300320012-0322021022300210-1301113022302102-2332323310022322-1110322311212333-3110000311300113"></a>

<a id="canonical-1210331133003021-3032101132033012-2222131310221033-3132200212011302-3231000233013103-0020102011111331-1211203230220110-2222303313001012"></a>

## configured_address property — local_dns / 110030010100 / 4

Type: `"string"`. Computed.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Upstream description:

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [first_address](data-sources--voltstack_site--reference--group-008.md#canonical-1033231000300333-0313000323322301-0221020110031320-2002202002330221-1311221011120312-3030001322313113-0322132122020131-1330230223310301): complete subsection reference.

- [last_address](data-sources--voltstack_site--reference--group-008.md#canonical-1233210132303012-1003032002222020-3120020001221030-1133121023203121-0221013102210303-1101323133232331-0133011333132222-0222201113113303): complete subsection reference.

<a id="canonical-2012020010301030-3111220000030001-2212122113101213-2223333021212203-1333222000122122-3111010213021110-2330012301200311-1300230013200302"></a>

## Next pages — local_dns / 110030010100 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](data-sources--voltstack_site--reference--group-008.md#canonical-1033231000300333-0313000323322301-0221020110031320-2002202002330221-1311221011120312-3030001322313113-0322132122020131-1330230223310301)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](data-sources--voltstack_site--reference--group-008.md#canonical-1233210132303012-1003032002222020-3120020001221030-1133121023203121-0221013102210303-1101323133232331-0133011333132222-0222201113113303)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1033231000300333-0313000323322301-0221020110031320-2002202002330221-1311221011120312-3030001322313113-0322132122020131-1330230223310301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120012010013202-1033210212130331-2102031311221323-3010033212221100-2332032102320112-2122301001113231-2030322322031030-2210303000000202"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 002221130301 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-008.md#canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0213100001323201-2322020001333310-1100312132000130-0103323010103220-0200212203112233-1203123012213113-1302320222103223-3213012312102122"></a>

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

<a id="canonical-1121130000211310-1320010203223211-3223011111103231-2210222100213012-0033201113013132-1023133312130212-2221312122030110-0030201301213231"></a>

## Direct properties — first_address / 002221130301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0111211322322122-3332023220231032-0111020122120320-2312133232313111-1133100032133310-0030220303201012-0320020120030123-1222011210313131"></a>

## Next pages — first_address / 002221130301 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-008.md#canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1233210132303012-1003032002222020-3120020001221030-1133121023203121-0221013102210303-1101323133232331-0133011333132222-0222201113113303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330211201331302-3332010223221223-2032323230230302-1331123313320002-2303012202201110-0302031031131120-2003121023120333-0030113333012013"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 001001032222 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](data-sources--voltstack_site--reference--group-008.md#canonical-0122220301133021-3300202330133101-2123132220303121-1331312323032230-0002103320003013-2212310010213000-3021003023113013-1200030200012011)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-008.md#canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2110332101003110-1330311110220112-3321110321213132-0202013303210133-3022023120102333-2133220130133031-2133030031203233-2201021123000112"></a>

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

<a id="canonical-3122231221313221-3011003130210100-1103222112303011-3013320112333332-2132120223301130-3022331213333201-2120223032013300-0032102232001323"></a>

## Direct properties — last_address / 001001032222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303030120022020-2102030133203011-1023330302123120-1301323130110010-0321133123010131-0021020212033030-0302302312331231-3000232020022123"></a>

## Next pages — last_address / 001001032222 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](data-sources--voltstack_site--reference--group-008.md#canonical-2021322002300332-2331201030313331-3322213003113202-1131331301333030-0331200331230003-3302122321223001-3112100023022010-0102010302221203)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211310301132232-3220233320200132-0201213020121103-0213001031220230-2313131122021120-3321003320220022-2031201332102302-0231320021332130"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful — stateful / 303303300121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful

<a id="canonical-2030020121302132-3032323013231222-0233323013320122-3212300021223211-2223312330013002-3231323322122313-3033203221230331-2210212222321303"></a>

Type: `"single"`. Computed.

DHCPIPV6 Stateful Server.

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

<a id="canonical-0011033203103002-1221223012231323-1332231002112030-2121312121200213-2023203202233022-0333300311103133-3002302301123013-2023231221112320"></a>

## Direct properties — stateful / 303303300121 / 3

- [automatic_from_end](data-sources--voltstack_site--reference--group-008.md#canonical-3321320001032322-0113133010201200-0310111023033022-1130003131310223-3001021222131323-1213032021233032-0201113231133230-1221033221132001): complete subsection reference.

- [automatic_from_start](data-sources--voltstack_site--reference--group-008.md#canonical-3332130311330200-0213203013111112-2310122303331212-2020020231000303-0303032122133003-1210003021232000-1333331030133203-0122001300133202): complete subsection reference.

- [dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-2312010301000011-1300320322331132-1120002033230333-3233333020311033-2121230120031232-1302122132302312-0231122332101103-3310102311132021): complete subsection reference.

<a id="canonical-3222032132321033-2310000233103001-0113110022113311-3020023113200000-2222120301211330-0023022033101010-2010213101332210-1213322202302003"></a>

<a id="canonical-3322112020000013-3202312030210111-3010201111300020-0322022132103212-1110102301222211-1222230112033121-2030311200000010-3132302211130332"></a>

## fixed_ip_map property — stateful / 303303300121 / 4

Type: `["map", "string"]`. Computed.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
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

- [interface_ip_map](data-sources--voltstack_site--reference--group-008.md#canonical-1323222332220011-0301213033221021-0021020110231332-0222110321332330-2120322230021111-0033300202212200-3212112120032012-1110111233232010): complete subsection reference.

<a id="canonical-1323101220021110-0222202331200312-3012301101300002-1122220233231312-3311132301121123-1002120133001231-1231122200330031-3213221030200031"></a>

## Next pages — stateful / 303303300121 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end](data-sources--voltstack_site--reference--group-008.md#canonical-3321320001032322-0113133010201200-0310111023033022-1130003131310223-3001021222131323-1213032021233032-0201113231133230-1221033221132001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start](data-sources--voltstack_site--reference--group-008.md#canonical-3332130311330200-0213203013111112-2310122303331212-2020020231000303-0303032122133003-1210003021232000-1333331030133203-0122001300133202)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-2312010301000011-1300320322331132-1120002033230333-3233333020311033-2121230120031232-1302122132302312-0231122332101103-3310102311132021)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map](data-sources--voltstack_site--reference--group-008.md#canonical-1323222332220011-0301213033221021-0021020110231332-0222110321332330-2120322230021111-0033300202212200-3212112120032012-1110111233232010)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3321320001032322-0113133010201200-0310111023033022-1130003131310223-3001021222131323-1213032021233032-0201113231133230-1221033221132001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010032132101101-0132130013032030-2100030301111320-0121031332102011-2033103223212011-0002103330221333-3110031130133202-3332103303321102"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 122223103013 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-2022011012111322-3223223313230121-1101331213110021-1030312131032220-2000301312332121-1320102011002320-1021001010012102-0030311302303312"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-3322332213203002-1311003021131023-1112321212231121-1313332223310211-1311131332032231-2223111033002300-3032003103310203-0120212012211111"></a>

## Direct properties — automatic_from_end / 122223103013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1113333233012112-1103100102303220-3130002113202031-2032221111021112-2021121320230312-1331013333203230-3022333030002123-1233110220010212"></a>

## Next pages — automatic_from_end / 122223103013 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3332130311330200-0213203013111112-2310122303331212-2020020231000303-0303032122133003-1210003021232000-1333331030133203-0122001300133202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233033003200213-1332211110303301-3101031312132201-2101010122323132-2231133202221210-2233123103303232-0303102213132000-1310011121002211"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 202010221130 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0200220123032320-0031303130221123-1012301013002202-0222020011223322-3012320201320030-0020211102023013-3122132220200310-0002312331000012"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-1011202231301232-1223302322123012-2102003300212120-0200110333302331-2032223332312113-3311011230300333-0030031103103100-0131232320230123"></a>

## Direct properties — automatic_from_start / 202010221130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111112331200331-3322000212223030-1100220110331022-3310030201112312-1331103001113313-0312010023122113-1323012032030300-2302133131100111"></a>

## Next pages — automatic_from_start / 202010221130 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2312010301000011-1300320322331132-1120002033230333-3233333020311033-2121230120031232-1302122132302312-0231122332101103-3310102311132021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312233202001002-3102021023133210-2021110303030320-2300212230111001-2211123133313213-0231321221201223-2031023113020001-1203121030222200"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 200312331103 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2110131311030122-3122130201321211-1322333322211101-1110113213222233-1233202300110023-2023133000302100-2033221202131233-3133331003322323"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3100330202030022-3102223312132313-3132020212023031-3222310130113230-2320112210033132-1313230103210133-0020231003323221-3113221130323221"></a>

## Direct properties — dhcp_networks / 200312331103 / 3

<a id="canonical-2123303023000110-2232320303223201-3330233221220021-3311110313322233-1011101332032133-2132223201123033-2033332122013012-3110021223303322"></a>

<a id="canonical-0010131213330310-3202302011320010-0233113110032033-1112122320011123-1013202312300220-2220132211202203-3111012033011320-3213301231323123"></a>

## network_prefix property — dhcp_networks / 200312331103 / 4

Type: `"string"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3232311010011112-0312231111211100-2310233333210102-3013002011000133-2121301330222221-0122132102312120-0202112131230013-3002001212133221"></a>

<a id="canonical-2332102332203002-3020200220232023-1313100303331002-2112230102301111-0102022223212223-0201223010201203-1102211130110112-1302210322032101"></a>

## pool_settings property — dhcp_networks / 200312331103 / 5

Type: `"string"`. Computed.

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

- [pools](data-sources--voltstack_site--reference--group-008.md#canonical-1111132221203222-1133311010002102-0012121001133222-1222200321213212-0133112001333230-3213200002122321-0202333100112010-3013111230220030): complete subsection reference.

<a id="canonical-3330103030013032-3231133012321012-1323112113022000-3330020003103232-2203113213101100-0300223111011121-2322211302122100-3230213011321230"></a>

## Next pages — dhcp_networks / 200312331103 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](data-sources--voltstack_site--reference--group-008.md#canonical-1111132221203222-1133311010002102-0012121001133222-1222200321213212-0133112001333230-3213200002122321-0202333100112010-3013111230220030)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1111132221203222-1133311010002102-0012121001133222-1222200321213212-0133112001333230-3213200002122321-0202333100112010-3013111230220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112301113023332-0220302222013021-2213232110210332-0102113033101203-0023113133233101-2132113310321213-0111013230233303-0200022030001212"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 111120302221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-2312010301000011-1300320322331132-1120002033230333-3233333020311033-2121230120031232-1302122132302312-0231122332101103-3310102311132021)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-2012220021101113-1233023211013203-0132112132000132-2312320331312021-0332321123232232-0312233031233300-3321021301203011-1321002013113232"></a>

Type: `"list"`. Computed.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1303020012212010-1323130122103012-2110213001123310-3330031023131323-2211313113101111-0201113222303110-2013020123020310-0200320012033001"></a>

## Direct properties — pools / 111120302221 / 3

<a id="canonical-0201231300103133-0113020302223101-0223112001021203-0233113123110302-0322111123203001-2331203310001001-1333023020031232-3131023030232111"></a>

<a id="canonical-3203232120130220-0202320221233003-1320333213331302-1213302123212023-0202133121331011-1231023100133100-1231300103220013-1110210333323211"></a>

## end_ip property — pools / 111120302221 / 4

Type: `"string"`. Computed.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Upstream description:

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2213231302330210-3330122122220211-3332220012113222-3221213122033223-0313323230103102-2131332320310130-1013302023213002-1331231112020112"></a>

<a id="canonical-3232202112201132-3120313012103030-1122013021000233-3201130202121000-3210021300233212-2321002233330301-0031000320223320-3300310332113002"></a>

## start_ip property — pools / 111120302221 / 5

Type: `"string"`. Computed.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Upstream description:

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3030120100200012-1332311133212031-1112311201313233-3201311313031030-1213123013210021-1201022322322201-2013320220103230-0100302211220330"></a>

## Next pages — pools / 111120302221 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](data-sources--voltstack_site--reference--group-008.md#canonical-2312010301000011-1300320322331132-1120002033230333-3233333020311033-2121230120031232-1302122132302312-0231122332101103-3310102311132021)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1323222332220011-0301213033221021-0021020110231332-0222110321332330-2120322230021111-0033300202212200-3212112120032012-1110111233232010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122230311313101-3030200112230213-3300012011221213-3223102110233323-0232323031303013-2230030221333013-0012023100233302-1201230031212333"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 101203020330 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](data-sources--voltstack_site--reference--group-008.md#canonical-0232201311012022-2203200103130330-1230200301133231-0223112030323231-3222101202003012-1232202321132103-2020213333221203-2002111221022122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](data-sources--voltstack_site--reference--group-008.md#canonical-3333102131123230-2112021300300103-3031223013201111-0022233131033323-2221232212013032-3013223200200122-3332132220203011-2113000021301331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2021200313322032-0223320101220330-0023313332300022-0120030103033200-3003202112201201-0113303200133022-0313030122012123-3123031122200102"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2133021303111121-1101122232330203-2020103030303102-0032310332320122-2212200321302221-3010301312103031-2120112003222033-2110310213023112"></a>

## Direct properties — interface_ip_map / 101203020330 / 3

<a id="canonical-1222220102302301-2202332232221333-1100311102203203-3132000333322200-3002233132313022-1120222020130033-3123113312302110-2201120020212310"></a>

<a id="canonical-3111000101102011-2010133333101133-0200311320230313-2022320201020213-1131103033301013-3101312012110132-0220122132300101-1000033331032132"></a>

## interface_ip_map property — interface_ip_map / 101203020330 / 4

Type: `["map", "string"]`. Computed.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-2032200322213103-2112211303323100-1332333323131313-3211233102102031-3330021210200331-2302301322002210-2011301332220100-0021020212330232"></a>

## Next pages — interface_ip_map / 101203020330 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](data-sources--voltstack_site--reference--group-008.md#canonical-2221212331330302-2131303032011303-0201113033120313-3200210001320013-3030113023223331-2033032022132300-2030032223233101-3121310332103012)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1220202132132123-2001001023131212-2022033031313001-1021013121202233-3222221323021101-2030033312333122-2031021320000120-0111303130010220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103302323103233-3112101300310123-1311212003311333-0123001103120310-0000111221212033-1113321132222313-2300131210001300-2303120032303131"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary — is_primary / 103120112220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary

<a id="canonical-0311323120210003-3123220021320130-3031103100321202-1332321210230020-2013000013311210-0330120020003133-3120220022011112-1303120102212100"></a>

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

<a id="canonical-0213220321313131-1321310222133112-3313122102003210-2033223031211210-1000101021311212-2122333231310113-0030232021222201-0121211230320103"></a>

## Direct properties — is_primary / 103120112220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011313121202110-2130013013030120-0103302023203302-3030201313022023-2211102021220122-0223011132021213-2003121031011032-2012101102020001"></a>

## Next pages — is_primary / 103120112220 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2110112312110023-0030332011121312-2133230302333032-3202002201323033-2223113200303030-2013321031230213-3313031022003300-3321210001020313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1212013133301020-1200211233131333-0033002001223200-2230201301202102-2210210302101311-3111323100220002-0310321122002202-0100131310120032"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor — monitor / 312213021131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor

<a id="canonical-1321213003111100-2102133020130112-0131012201203222-2302100101222010-1110203013230022-3032120011202000-3120123233231313-3201212120321201"></a>

Type: `["object", {}]`. Computed.

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

<a id="canonical-2010123103001113-0021001103320103-2122120133030033-1031331330220011-0121303310302121-1302221213131131-3002033122302030-2130200121122303"></a>

## Direct properties — monitor / 312213021131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0002232132310011-2022313330001320-3122320203022232-0023132103202021-3121320211210122-3220003003110320-0122300301122212-1321201133221201"></a>

## Next pages — monitor / 312213021131 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2030023310001011-2320112113220133-1112333003301203-1030201120132303-3213031312011323-3231203201022102-0122100332013323-0031211120021101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013212030110232-2103332203220232-2121330222022321-3323331012012132-0033103130200120-2300101120111202-2321212331323010-0223100010101221"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled — monitor_disabled / 202232122131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled

<a id="canonical-0201220032302122-2312202231131223-1200002103101312-2223003330300233-2333311232020103-2112132321001201-1032301310111311-3321300111030322"></a>

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

<a id="canonical-3310203031010033-1212120101322333-1203030102222223-2203233123030203-1333002330220300-1330011222212202-3111020022002323-3110211302230100"></a>

## Direct properties — monitor_disabled / 202232122131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2200220230112201-0221213021220111-3020023310120102-2320220330220302-0333312020101320-2331113121200213-0130032302003002-0111031102103312"></a>

## Next pages — monitor_disabled / 202232122131 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0320032312103201-1103000230110221-2120111022313313-3231123111311033-2101321132232012-1111322130303230-3222102031321103-0123232111033011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0101221103331321-0313023111100100-2021013301012201-3121301112113110-3013300131123220-0102030132022132-2312301033322310-3110203100102030"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address — no_ipv6_address / 000123112020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address

<a id="canonical-1003220121132010-1133330021020323-3330303222020233-0021021031031020-1322210211131102-2323100321100310-2302331300100121-1331233201030322"></a>

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

<a id="canonical-0122033301001201-0223100202032000-2123223332101202-3312103231031200-3213103103310112-3220212031130332-2303303300332021-2022323333021100"></a>

## Direct properties — no_ipv6_address / 000123112020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0031203001330212-1211001203131002-2213120230302210-2330220232123122-1233320313223301-2220320323311020-2012123011002033-2023321023130330"></a>

## Next pages — no_ipv6_address / 000123112020 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3202123223312121-2201331233210200-2003113020033322-1122203000011230-3002212211321012-2120120300202323-0000303222101123-2121113222112011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001303022330202-2320033020101100-0010102221103220-1302023211102110-3303000120312111-2300101020321311-2111320212231121-2323101111003313"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary — not_primary / 230213212110 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary

<a id="canonical-2310231013120312-2323203221300022-2322120220122312-1200211010311331-3122122320230231-1133003103133102-2020111330113203-0320230333221330"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for not primary.

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

<a id="canonical-3023211022222221-2312120333001122-3301103132112211-1101103300013201-3133021020112020-0010232031122232-2210131332312303-1122211200301333"></a>

## Direct properties — not_primary / 230213212110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330321231203231-0322130331030210-3212230231013030-1023023000120320-1230010302110112-3222020320212131-1101111212303222-0031012220322200"></a>

## Next pages — not_primary / 230213212110 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1102121322230302-0333023220300213-1230120212220211-2302312033010013-0002301002132122-3233222020101032-1023111203132021-1021122230301213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3231333301100112-2320212231031300-1002202013032322-3223232202212031-1221130122220113-2323203312002302-0222113223222122-0320003322123330"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network — site_local_inside_network / 232012302112 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network

<a id="canonical-3303200002230200-3312213121013310-1101022201120133-1023331201312201-0332113331233300-3113010011123221-3313301121130002-1303002301011202"></a>

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

<a id="canonical-0330003020210013-1233023102302323-3020222032312211-1032011132202200-0223032003231221-3032303203222213-1011200312113013-0002320011002010"></a>

## Direct properties — site_local_inside_network / 232012302112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0000020310102221-0003223132303111-3021101203322012-1013100212012331-1111223110030131-3333030012100213-0122122133220030-0223333201333200"></a>

## Next pages — site_local_inside_network / 232012302112 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3010303102201030-0212132332013133-2020120320211111-0000011003030021-1212320132231103-0021223113211322-0210131310312221-3333032022023003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320322011321102-1210312331113010-2233321033031000-3311020232221102-0013232031013000-1232220210331120-0023321320123220-1321231033021202"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network — site_local_network / 313100331001 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network

<a id="canonical-3211233112112221-1330202103220233-3122131333213310-0233312322202002-2333332312111123-1310133000131003-0300103123122201-2032310202222020"></a>

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

<a id="canonical-1020231321303231-3311030321320233-0000112013201002-2210210023100022-1131020203321210-2002011130033023-2001312022313111-1033133212311010"></a>

## Direct properties — site_local_network / 313100331001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110223021320011-3023220322021110-2033222320110302-1112002320113222-1311313000120331-1312211123100202-1303121301323012-2132301122122331"></a>

## Next pages — site_local_network / 313100331001 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113113200310221-0213112200033030-2303210331001130-3221010323031100-3113011031132001-0223323103031223-0002130030223130-1323121301102213"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip — static_ip / 300100130133 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip

<a id="canonical-2231303030121301-3302121110123030-3333310320031010-2023131011001332-0023132031301333-0313131223011022-2100120120032000-2213133201020212"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

<a id="canonical-0200121333110331-0230013103000302-1120231332203132-0023133333211033-1213132110003300-3220220311021210-1313302121002212-1131220212022301"></a>

## Direct properties — static_ip / 300100130133 / 3

- [cluster_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-3121200131332202-2033312013132100-1303221323320013-2100312001003012-3211200323102130-2101010322322032-2010100303220300-3203031123103200): complete subsection reference.

- [node_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2110000121222131-1311212233223302-3010130103232023-2333022010121230-1000132333303123-0000211030013221-0123010213033010-2102333331230310): complete subsection reference.

<a id="canonical-2202200120030003-2220102110231120-1032110203231233-1331211210323300-3111233333210232-1310001123223303-1122311223131123-2123120323330122"></a>

## Next pages — static_ip / 300100130133 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-3121200131332202-2033312013132100-1303221323320013-2100312001003012-3211200323102130-2101010322322032-2010100303220300-3203031123103200)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2110000121222131-1311212233223302-3010130103232023-2333022010121230-1000132333303123-0000211030013221-0123010213033010-2102333331230310)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3121200131332202-2033312013132100-1303221323320013-2100312001003012-3211200323102130-2101010322322032-2010100303220300-3203031123103200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303121321032123-0022033300100312-2120100330031013-3111321221222301-2312110221100210-0311301310312333-3102110022223110-1303200003012112"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip — cluster_static_ip / 002021120023 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip

<a id="canonical-3312202032032213-2130003032132023-3000313021202021-2021020223322212-3323323000012323-1320002230131011-0000023210112220-2221110012331212"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3133322220021313-3112001220323121-0110211220330023-3211322312302232-0012110232001312-0303333331222120-1031212201111132-3031312022323103"></a>

## Direct properties — cluster_static_ip / 002021120023 / 3

<a id="canonical-2123221210201223-3330021003000311-1312232110223312-0012231232101112-2110011320323031-0110323203102022-3032012121130123-2213112312230011"></a>

<a id="canonical-2303320221232100-0033300213122011-1101001103302210-0123321021332032-0100031320230323-3330110123121332-2332030030100002-3303131100311313"></a>

## interface_ip_map property — cluster_static_ip / 002021120023 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-3300133211022103-0231322132021222-2022220202300222-0130131321201220-2233100132320220-3200132133113213-3211023220010020-3100102122101310"></a>

## Next pages — cluster_static_ip / 002021120023 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2110000121222131-1311212233223302-3010130103232023-2333022010121230-1000132333303123-0000211030013221-0123010213033010-2102333331230310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022201002322231-0030332301210220-3320013222302122-0131232112313030-3331002120003212-1312303111100210-3011121122013201-3001311100312303"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip — node_static_ip / 132311220111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip

<a id="canonical-0322000021202332-2331132310210110-3303232230001011-0100202313130330-0132131130223310-0023231023201102-2120100303010021-0131223013101103"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-3023223011213000-0031300120000200-0210211100301101-1030101031003221-2311301311200330-1200130323230033-1330110221210320-1022321010103002"></a>

## Direct properties — node_static_ip / 132311220111 / 3

<a id="canonical-1333123321300000-1103122332122333-0103231222312132-3121232101011003-0331103200132001-0032213001021202-2311130222000001-2020213013301202"></a>

<a id="canonical-1331111210212120-3011222031232010-3110100321012200-1221311130113121-0311323132103022-3033011022233332-2003102011211230-2232002212213331"></a>

## default_gw property — node_static_ip / 132311220111 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3122031301232220-1122032101201120-1000133011133102-0310032220002030-2111131333101322-1110002211323030-0031222213310120-1032113012222013"></a>

<a id="canonical-0211210030001330-0203101003211102-0310231221113031-0320313223030011-3010113110231023-3323231301303302-3300212031112032-1000220032330223"></a>

## dns_server property — node_static_ip / 132311220111 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-0100023112213110-0233220311101331-2001122030011100-0323031103101030-2130313303103032-0220001231022022-1022332303311231-0311023021223200"></a>

<a id="canonical-2213300003013121-3300223323222002-1101312300003023-1212203023310122-3103003101230313-1022213223113200-1332112001111132-3023000310110321"></a>

## ip_address property — node_static_ip / 132311220111 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0303221030100111-1110023030302230-1212323110020113-2103213032101121-0113111232012300-3123211003020200-3212302202322311-2012010333011213"></a>

## Next pages — node_static_ip / 132311220111 / 7

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2001102210110011-2301000210121201-1001131031120112-2200331212101002-1111010023012221-2020300312031230-3002311223111202-2113013123122122)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200111311030001-2230020010120330-0331313321312202-2132310211121313-3011323110320220-3300131320213330-2220313303332100-2132000321120313"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address — static_ipv6_address / 000303010121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address

<a id="canonical-2110011331200323-0012003103301121-2310133312121110-1312130010303220-3130310301300013-2103111230011023-3033300032030032-0000313313102223"></a>

Type: `"single"`. Computed.

Static IP Parameters. Configure Static IP parameters.

Upstream description:

Configure Static IP parameters.

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

<a id="canonical-0203320121223301-1023102110302103-1001003321201311-3200213212020030-2030011332323100-0111120020033310-2120023000010203-2311123103221200"></a>

## Direct properties — static_ipv6_address / 000303010121 / 3

- [cluster_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-0330003113311232-3031101302333123-2323020103310230-3213203211130213-0303021320113113-0132130232332301-0221310311113021-1110202332200323): complete subsection reference.

- [node_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2321132003013022-0120331133010210-0011121313031303-3321001022121313-3320020233201201-0121122022321331-0013312333220313-0313002010213110): complete subsection reference.

<a id="canonical-3313223300122001-2311032333033332-2210032122132102-1313200212103002-2132031333130313-3000112123013100-2333331220223132-0220120121223031"></a>

## Next pages — static_ipv6_address / 000303010121 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-0330003113311232-3031101302333123-2323020103310230-3213203211130213-0303021320113113-0132130232332301-0221310311113021-1110202332200323)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.node_static_ip](data-sources--voltstack_site--reference--group-008.md#canonical-2321132003013022-0120331133010210-0011121313031303-3321001022121313-3320020233201201-0121122022321331-0013312333220313-0313002010213110)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0330003113311232-3031101302333123-2323020103310230-3213203211130213-0303021320113113-0132130232332301-0221310311113021-1110202332200323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210201212010031-0312331231100011-0013032033000131-1232020313013322-1330202320210110-1322031100113322-0011120300220230-3212332302102330"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip — cluster_static_ip / 202023033010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip

<a id="canonical-2211101233320222-2320002312031310-0202321321312111-2213012132223323-0220222113110303-3331211302101302-0213023020331321-0122201101201221"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3101302002213330-2131011232221200-1203330111232313-3331333131233130-0103202011033301-1223300202313133-3230203232213120-2320211113233102"></a>

## Direct properties — cluster_static_ip / 202023033010 / 3

<a id="canonical-3001102011000103-1031321220211111-2300333322212331-1032233021010310-2301010220230310-0233200103330310-2101010310002222-0132130031330202"></a>

<a id="canonical-2301130213332210-3121132112021121-3020320120332222-3313203303001322-3103131203320302-3330112100220032-3210332302301133-1322101202231031"></a>

## interface_ip_map property — cluster_static_ip / 202023033010 / 4

Type: `["map", "string"]`. Computed.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
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

<a id="canonical-2230311203223202-0010020321012230-3200330032222003-0132023011003020-2000220010111130-2331021023120221-1211211220123131-2213313211332121"></a>

## Next pages — cluster_static_ip / 202023033010 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2321132003013022-0120331133010210-0011121313031303-3321001022121313-3320020233201201-0121122022321331-0013312333220313-0313002010213110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122003320233231-2333123312123211-2330031232123130-1102022300131210-1232201020021113-2301320003003322-1100121200200110-3013103320122001"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.node_static_ip — node_static_ip / 003220201012 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.node_static_ip

<a id="canonical-3031201132001322-3220303013000222-3112013301230033-1032113133123102-2120320200000210-2221101121032011-3023233003133330-2202312200302210"></a>

Type: `"single"`. Computed.

Configure Static IP parameters for a node.

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

<a id="canonical-1232021102102111-1011012231020211-1331113033231020-3033320012212200-2113301232033132-3233011131023001-3031012300001132-3123323103001113"></a>

## Direct properties — node_static_ip / 003220201012 / 3

<a id="canonical-3303010201201123-2233211021320330-2323231302332102-3032020233220220-0331211122112332-2330332103320233-2101110121230202-3102330110310000"></a>

<a id="canonical-3320023230022233-0323022332203303-1101113331331212-1302330202211331-3321330203323313-1030021210111133-0112313312312003-2322223223132330"></a>

## default_gw property — node_static_ip / 003220201012 / 4

Type: `"string"`. Computed.

Default Gateway. IP address of the default gateway.

Upstream description:

IP address of the default gateway.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3303123103010313-2312121210020103-0013230303333323-2123200123212103-3312011020323202-3101323312132201-1103313021100100-2223121021012100"></a>

<a id="canonical-1332111221211020-2232002222320232-1300210103103031-2303112123313213-2021213032230322-0320033002222110-1230133113312132-1220311202131230"></a>

## dns_server property — node_static_ip / 003220201012 / 5

Type: `"string"`. Computed.

DNS server address for the static interface configuration.

<a id="canonical-0011130030031331-2010332011300122-0022301230101320-2211213132303202-2030302312030231-3111223100230302-1000330223001103-3032103213122212"></a>

<a id="canonical-2213022133131322-0103123020221111-3123302032202221-2101010110303101-1130131032313022-2101102222122322-1000022100213133-2103202330130111"></a>

## ip_address property — node_static_ip / 003220201012 / 6

Type: `"string"`. Computed.

IP address of the interface and prefix length.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1331003000002323-2200021032210301-0313200310101311-1000311321012311-1021012233232223-1030233330112212-0322112033223323-0121202212300202"></a>

## Next pages — node_static_ip / 003220201012 / 7

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](data-sources--voltstack_site--reference--group-008.md#canonical-0201303232323320-3233203102011101-0232000020301320-2330223130030202-1211111211002013-3123331333003232-1310300331030201-2231230000000131)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3201101213330023-0130032023121030-2300130300313111-1303202211122101-0230110103133123-2120112313121323-2200033023122233-0200313213123323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2222000111032213-1301013203030323-0312212123310333-2212313020203331-0222132032300233-3103022113332031-1321321230230332-0322322000011032"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network — storage_network / 222132230333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network

<a id="canonical-3130302213101330-3331131202323112-0320332221003130-0230030123201020-1102000031133212-2033330103001113-1222300230012132-3003331321021130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for storage network.

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

<a id="canonical-0120013131023013-0222003311011300-0313121303110231-3012300332200011-2220213230320120-1233310012312330-3112233011120123-0213310333030101"></a>

## Direct properties — storage_network / 222132230333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332101323012302-3013331010000002-3223123030010302-2110321232012302-1013133232322202-2132221120232203-1130011221031101-2300302221212030"></a>

## Next pages — storage_network / 222132230333 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0102000330232030-2013030113312230-1220203131223031-1112231133220313-0321203100221212-1322001103231112-0222311123022203-1300310312313132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330022000012103-3320330011113310-0201111312010230-1202212233303123-2000001120121012-0222001000211010-0121131213111120-1201131333011101"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged — untagged / 032013200200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [custom_storage_config](data-sources--voltstack_site--reference--group-005.md#canonical-1322330111022022-1222331222013311-0203320011033023-1312120310031232-0021011331210300-1111322303223100-1320332332030123-1110131212333020)
- [custom_storage_config.storage_interface_list](data-sources--voltstack_site--reference--group-007.md#canonical-0323230103311210-3030023003220002-0001300310102332-3012322001200113-3233220113213113-2233320301323200-0222321232211121-0022031220011321)
- [custom_storage_config.storage_interface_list.storage_interfaces](data-sources--voltstack_site--reference--group-007.md#canonical-3101323032130103-0121313201200303-2002022323010113-1102101211210302-2112003000320221-3030232310022110-3103100221111001-3303131231011132)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged

<a id="canonical-2233002301331130-2012130233023033-3332330031010313-2122130300220120-2233213330030331-2211220020102012-1313003122003012-0200311011012012"></a>

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

<a id="canonical-1001123210301013-0301300302303302-0212331221031102-3120321113331200-2322022032213121-1332221220222031-3003111020103333-2233002100320310"></a>

## Direct properties — untagged / 032013200200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330300021032210-1003002030202201-0311213100232120-1110333230223112-2211202013032001-1203223213020202-2003032022103310-3123121010130013"></a>

## Next pages — untagged / 032013200200 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](data-sources--voltstack_site--reference--group-007.md#canonical-0002211011223120-1231223332232103-0203300100110000-1331122231310032-2300033200220211-3002003132003021-2100311231203031-2030230331233112)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3213101201031113-3231303201013301-0231002333210112-1212331331322010-0231102213222122-1331321110330131-0100113322130111-3112310031310310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1233303301200231-2201320333321203-1103233212222302-3220323023210211-3302023131120000-1321300300131021-3031012103321000-2300313020231102"></a>

## default_blocked_services — default_blocked_services / 113300001300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- default_blocked_services

<a id="canonical-3032310112100032-1132301333302212-0223113120023211-2301310313132001-3022121102311011-3032010102321210-3210123213302323-2223020012003132"></a>

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

<a id="canonical-1222131022200312-1121222303033103-1030130231102210-2013031121112011-3213321121211102-0130133231010102-3231131323323030-0021321113333313"></a>

## Direct properties — default_blocked_services / 113300001300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3301112032210021-0101213132300321-3020213322303121-3203131113203201-0303311323101212-3303031313001020-0131213311001132-0112121301201321"></a>

## Next pages — default_blocked_services / 113300001300 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2131230020322123-1303110032220233-0323021011312130-2022321132301211-2212210322021202-3222312320032321-0313130320333323-2113222132113330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212313321100100-0333303211113132-3231121321110313-3110111311300102-1130122112202030-1012100210201303-2020203201121002-3110230320031332"></a>

## default_network_config — default_network_config / 323222102331 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- default_network_config

<a id="canonical-0031322002210113-2230011133033212-2102301022321301-2323202203222121-2001130120113310-2120100120011311-3202111111010301-2030302332220211"></a>

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

<a id="canonical-2012331203032131-3132321021002123-2112210001012322-3032010101202220-1333330321322322-0321000101121300-0000010313121003-2221310011123030"></a>

## Direct properties — default_network_config / 323222102331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3323310301220013-3102031123113220-2102232113220302-2300203320230001-1113332130010222-1020123121202330-2113313000210231-3012321321133320"></a>

## Next pages — default_network_config / 323222102331 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0030112211201131-0233032203330233-3012211220221210-1101021313232031-2202102032101020-1213210233030113-2330202220220200-1203113120131202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331031321012032-0031121302223233-2201120110002213-2013331123333111-2130301103021222-1230302120023002-1120312022322201-3313113113103322"></a>

## default_sriov_interface — default_sriov_interface / 031220211111 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- default_sriov_interface

<a id="canonical-3321100002013210-2102212231132100-1230002220002310-1110020111010213-3000003003221332-1321323333210101-2211021312122030-3112032322000332"></a>

Type: `["object", {}]`. Computed.

\[OneOf: default\_sriov\_interface, sriov\_interfaces; Default: default\_sriov\_interface\]
Configuration parameter for default sriov interface.

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

- [default_sriov_interface](data-sources--voltstack_site--reference--group-008.md#canonical-3321100002013210-2102212231132100-1230002220002310-1110020111010213-3000003003221332-1321323333210101-2211021312122030-3112032322000332)
- [sriov_interfaces](data-sources--voltstack_site--reference--group-010.md#canonical-2130201231213303-1230003033112020-0213313101120303-2003323211110112-1013013211313330-2012210030303101-3100322023220312-0232331302001023)

Select alternatives according to the provider validators above.

<a id="canonical-0013232303132123-1003122230211021-2320001111030233-3031311120031303-3102123111122313-0300000023320102-3330023000333200-3000201021101313"></a>

## Direct properties — default_sriov_interface / 031220211111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303331330323013-2202332333213003-1110321100302022-0311130121132300-1211300022013322-2100100200013203-3301120130033010-1213221203323120"></a>

## Next pages — default_sriov_interface / 031220211111 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0021033110303103-0003030222103213-0103111210211111-3220133201312030-3103021232313120-1031020022023221-3102223131032022-2133202231130231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102030312222331-1202030223301002-3121210323320021-0110003212021113-0312133202231021-2323330330310302-1213131131312221-0231231031021103"></a>

## default_storage_config — default_storage_config / 312030230002 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- default_storage_config

<a id="canonical-3112201001132100-1130202120221200-0220303013130011-0011331213301220-0331123030130031-2321023013233332-0300222320113013-0023020100230102"></a>

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

<a id="canonical-2223133323101012-3003330032022212-0233220030131232-3011002003211312-2111003223220131-2020001323103321-3021002001033313-3033031020331300"></a>

## Direct properties — default_storage_config / 312030230002 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2203212332001231-1010000223312123-3220130212132020-1222000330233101-0013110122102003-1320311112010332-0230322310001222-3010212222330203"></a>

## Next pages — default_storage_config / 312030230002 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-0102023001301202-2201311011312032-0331020330222220-1022123133210333-1111000022131312-0020010200333130-0312111120133202-0331021201333110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3122311302010211-2231320021220231-3000231210210010-3012121303020020-0133000023000123-3312000102232033-3213202002011211-1300200212132331"></a>

## deny_all_usb — deny_all_usb / 131012003000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- deny_all_usb

<a id="canonical-2201130011322112-3210131013111333-3122310133103330-2213332210333323-0322311123122320-1011320331212101-0032111203213021-2313333002122301"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for deny all usb.

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

<a id="canonical-0132013322303311-3222030112132032-1303302012120000-1201013021021133-2230203133311130-3131313020010211-2001223331033100-1122201210301321"></a>

## Direct properties — deny_all_usb / 131012003000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321313201010301-1332001222020033-1012122302320120-0132220132100200-1132011213132012-3101333211200103-0221311132022302-0220131000102322"></a>

## Next pages — deny_all_usb / 131012003000 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2013133233301023-3002101131132333-0132130032032231-1201022321001011-2201310323103301-3212311322221321-3312032120220131-0232123003003001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232132020231003-1303010132011133-1033102001222333-2023002131320313-2322011333030301-2323230300322100-1102300110202000-2233103133312212"></a>

## disable_gpu — disable_gpu / 123221123313 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- disable_gpu

<a id="canonical-1222220001201131-0321323023203101-0113120331211011-3131113010130101-0121132230301332-3100001000012213-3122031033331332-2130120303311033"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_gpu, enable\_gpu, enable\_vgpu; Default: disable\_gpu\] Configuration parameter
for disable GPU.

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

- [disable_gpu](data-sources--voltstack_site--reference--group-008.md#canonical-1222220001201131-0321323023203101-0113120331211011-3131113010130101-0121132230301332-3100001000012213-3122031033331332-2130120303311033)
- [enable_gpu](data-sources--voltstack_site--reference--group-008.md#canonical-2233012122202301-1300030020100010-2300013102033032-0302203210103110-0313321211233130-1220322312123330-2211300300312030-3130031211001302)
- [enable_vgpu](data-sources--voltstack_site--reference--group-008.md#canonical-0030020213330000-1122013130013133-2312022203212133-3023110222323100-2321220211002123-3000200310110032-2312020211113232-3122113023111101)

Select alternatives according to the provider validators above.

<a id="canonical-3011233022111203-0301232010303333-3302031032213032-1111220103301000-1132111231221323-2312123001120032-3031321310031101-2210022211113011"></a>

## Direct properties — disable_gpu / 123221123313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223022313121202-0223223033213320-3110231122233302-3323223211121022-0211331103122101-0112212320213131-0002230011031030-2112221320120111"></a>

## Next pages — disable_gpu / 123221123313 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1033000201210201-3132201233301200-1212311330301013-3030133000212021-3320323323311230-1211202021132213-3011000201300121-0321301200302120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001231233022101-3231233211332323-1132330033232300-1012120020212101-0100223321121201-2002031103231022-3323300333012010-0103310213233101"></a>

## disable_vm — disable_vm / 200010030221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- disable_vm

<a id="canonical-0303003110001212-3120322010101101-1321122030330112-2323221223101302-1000031111023032-0103101101001232-3301223220030200-2101030133322000"></a>

Type: `["object", {}]`. Computed.

\[OneOf: disable\_vm, enable\_vm; Default: disable\_vm\] Enable this option

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

- [disable_vm](data-sources--voltstack_site--reference--group-008.md#canonical-0303003110001212-3120322010101101-1321122030330112-2323221223101302-1000031111023032-0103101101001232-3301223220030200-2101030133322000)
- [enable_vm](data-sources--voltstack_site--reference--group-008.md#canonical-2020200311313011-3232123011223033-3323320322331112-0103113121133303-2213221102313321-3312230011021111-2200133132202212-2332020333122130)

Select alternatives according to the provider validators above.

<a id="canonical-3130030232032013-1332231122330131-1221212133321003-3303311001002330-1102013201202200-3022000023002123-3200211210110200-1100231133101122"></a>

## Direct properties — disable_vm / 200010030221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021313103300300-1202032321110101-0023112132200222-1121222013110023-1021200011102120-3201330201011012-2001120213003200-3000000003103033"></a>

## Next pages — disable_vm / 200010030221 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-3301023003030020-3333310331323132-3301311002332122-0011132321032203-1003232133133122-1220020221013232-2022011123101120-1230033021010322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333002102201023-1021321333330022-3132112123232122-2133001211220103-3111022332320300-1100100222110302-2110232310203030-3203100120023132"></a>

## enable_gpu — enable_gpu / 002302010331 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- enable_gpu

<a id="canonical-2233012122202301-1300030020100010-2300013102033032-0302203210103110-0313321211233130-1220322312123330-2211300300312030-3130031211001302"></a>

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

<a id="canonical-0210323032302013-2023033212033221-3230112312303030-0200300112130012-0113310001121222-0331300123310112-2122120023310232-2312220330213121"></a>

## Direct properties — enable_gpu / 002302010331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022132023002301-3310002332131322-1303331133321032-1321022003122330-3212330130311002-3022022222211222-2122000232301002-3013012233011112"></a>

## Next pages — enable_gpu / 002302010331 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1131330000323001-2000332111122333-3330122310233213-2112201320320113-3011213113023220-1211302100220022-2031013022001131-1002210211033323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302300323302020-1123320022202331-2012101231202011-2002012313111312-0130301012023100-0132321100222110-1033223320001000-2333201121311001"></a>

## enable_vgpu — enable_vgpu / 302003122012 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- enable_vgpu

<a id="canonical-0030020213330000-1122013130013133-2312022203212133-3023110222323100-2321220211002123-3000200310110032-2312020211113232-3122113023111101"></a>

Type: `"single"`. Computed.

Licensing configuration for NVIDIA vGPU.

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

<a id="canonical-1232303211312300-0201033010233221-1311122011323123-0002200013301303-1311103302101233-1130123233011232-3223020102331031-2122023132301132"></a>

## Direct properties — enable_vgpu / 302003122012 / 3

<a id="canonical-2332322213000012-2122133133010221-1031223120332333-2203220312323132-3030303033132130-1322302233010302-0230232320032023-3133011301011101"></a>

<a id="canonical-0100323132300312-3022012220221003-3031033201233311-1023013232031210-0313101231033331-0000302321013333-2012110030002312-3302021031132220"></a>

## feature_type property — enable_vgpu / 302003122012 / 4

Type: `"string"`. Computed.

\[Enum: UNLICENSED|VGPU|VWS|VCS\] Set feature to be enabled Operate with a degraded vGPU performance
Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.
Possible values are \`UNLICENSED\`, \`VGPU\`, \`VWS\`, \`VCS\`. Defaults to \`UNLICENSED\`.

Upstream description:

Set feature to be enabled

Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation
Enable NVIDIA Virtual Compute Server.

Receipt-pinned upstream constraints:

```json
{
  "default": "UNLICENSED",
  "enum": [
    "UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1212020303113111-0113022022330121-2322313002030130-3312202110100230-2232011303001312-2210311302323300-1223130201322113-2003011023201030"></a>

<a id="canonical-3101333312201231-2332222221232102-1213332100202320-2112210133022331-0132020002210010-3020201010110033-2100023010020332-1330110200003301"></a>

## server_address property — enable_vgpu / 302003122012 / 5

Type: `"string"`. Computed.

License Server Address. Set License Server Address.

Upstream description:

Set License Server Address.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

<a id="canonical-2031100103103321-0100103100221231-1213213213322202-1202002101130012-1331102013122302-2323110100232320-1020210002310202-0001122230230222"></a>

<a id="canonical-1321221120003012-0221132200223310-2023130030030030-2211202232212213-3321222003131013-1300033022230131-2331323031232111-0332023011022013"></a>

## server_port property — enable_vgpu / 302003122012 / 6

Type: `"number"`. Computed.

License Server Port Number. Set License Server port number.

Upstream description:

Set License Server port number.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-1332131133333333-2010321320202212-0311033211211032-0221313213203132-3020323300221320-2310331333212323-0221113203210120-1223110011332112"></a>

## Next pages — enable_vgpu / 302003122012 / 7

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-2120033220230201-3201323001000303-2300010003321030-1322013303121300-1311012110311000-2112222220331213-2322312333333022-2023330301130102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003201321300102-3103311022200203-1123032313100232-3010101032013211-1121311333322323-3230221112220231-0132120221221303-3012122222302032"></a>

## enable_vm — enable_vm / 212222322202 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- enable_vm

<a id="canonical-2020200311313011-3232123011223033-3323320322331112-0103113121133303-2213221102313321-3312230011021111-2200133132202212-2332020333122130"></a>

Type: `["object", {}]`. Computed.

VM Configuration. VMs support configuration.

Upstream description:

VMs support configuration.

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

<a id="canonical-2102130221220100-2130331131223233-3232323033221310-1133033221121333-1103233312203333-0201021230133230-2111332010110102-3222123322301203"></a>

## Direct properties — enable_vm / 212222322202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3331121021332202-3202102211212020-3120012102332132-3221133220202212-3230023231211301-3111003231010330-1201221210133231-2310331132033002"></a>

## Next pages — enable_vm / 212222322202 / 4

- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)

<a id="canonical-1132230233133203-3131223100302000-2120313313123100-3220310332231001-0231212333120223-3101313210021001-1002131331120000-1311031032021230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320102130021002-0231332312122000-3301312200323021-0323302222033102-2133331321032030-3223332232030133-0010233233210303-1222021121211231"></a>

## k8s_cluster — k8s_cluster / 023210323211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../data-sources/voltstack_site.md#canonical-3011021120121122-1120101211211330-2030121003211010-0212132020130223-0212111100103011-3003130023302333-3021021330130133-1133312111123000)
- [Property reference](data-sources--voltstack_site--reference--group-001.md#canonical-3210130022101001-3112232101022102-2032333031300013-3112003110021012-3222001330012233-0002330230030031-3201232233001313-3100032330211031)
- k8s_cluster

<a id="canonical-2311100022221013-1222221000021030-1011010103103301-1213133132103320-2222320100220012-2232023100002321-1102310223031122-1100212010322121"></a>

Type: `"single"`. Computed.

\[OneOf: k8s\_cluster, no\_k8s\_cluster; Default: no\_k8s\_cluster\] Type establishes a direct
reference from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

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

OneOf alternatives in this subsection:

- [k8s_cluster](data-sources--voltstack_site--reference--group-008.md#canonical-2311100022221013-1222221000021030-1011010103103301-1213133132103320-2222320100220012-2232023100002321-1102310223031122-1100212010322121)
- [no_k8s_cluster](data-sources--voltstack_site--reference--group-009.md#canonical-1013201012231322-1122033102111202-3123031303112213-3001013311103230-3000012313111132-3211203112221102-3332233221001323-0030300122300212)

Select alternatives according to the provider validators above.

<a id="canonical-3111121031313312-1300321000103230-1302133202133211-2330123201232023-3000323123010122-2001123212230320-0320220102032132-2000021101100213"></a>

## Direct properties — k8s_cluster / 023210323211 / 3

<a id="canonical-1121300331321002-2321103130333332-3320310110220000-3303120322323300-0112201012223211-3230332022210201-1312231032120121-0220011111100010"></a>

<a id="canonical-0033200022311122-1323022213232013-2322132303202330-1331000313212103-2013310101203312-1313101332011312-2312212331233030-2000332120001032"></a>

## name property — k8s_cluster / 023210323211 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1023122021113322-2302231321123203-3312331321331103-3302021331223023-0022131311313333-3121132321111100-1002201012312333-0032100120302310"></a>
