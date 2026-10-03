---
page_title: "xcsh_fleet reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet reference."
---

# xcsh_fleet reference

<a id="canonical-2011121302312100-3121101013202013-2301233033002101-3223100013110310-1003123132023112-1310310010003010-2013333330200023-2120120213310332"></a>

## network_type property — blocked_services / 012210011300 / 4

Type: `"string"`. Optional.

\[Enum:
VIRTUAL\_NETWORK\_SITE\_LOCAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE|VIRTUAL\_NETWORK\_PER\_SITE|VIRTUAL\_NETWORK\_PUBLIC|VIRTUAL\_NETWORK\_GLOBAL|VIRTUAL\_NETWORK\_SITE\_SERVICE|VIRTUAL\_NETWORK\_VER\_INTERNAL|VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE|VIRTUAL\_NETWORK\_IP\_AUTO|VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK|VIRTUAL\_NETWORK\_SRV6\_NETWORK|VIRTUAL\_NETWORK\_IP\_FABRIC|VIRTUAL\_NETWORK\_SEGMENT|VIRTUAL\_NETWORK\_MANAGEMENT\]
Different types of virtual networks understood by the system Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside) network. This is an insecure
network and is connected to public internet via NAT Gateways/firwalls Virtual-network of this type
is local to.. Possible values are \`VIRTUAL\_NETWORK\_SITE\_LOCAL\`,
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\`, \`VIRTUAL\_NETWORK\_PER\_SITE\`,
\`VIRTUAL\_NETWORK\_PUBLIC\`, \`VIRTUAL\_NETWORK\_GLOBAL\`, \`VIRTUAL\_NETWORK\_SITE\_SERVICE\`,
\`VIRTUAL\_NETWORK\_VER\_INTERNAL\`, \`VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE\`,
\`VIRTUAL\_NETWORK\_IP\_AUTO\`, \`VIRTUAL\_NETWORK\_VOLTADN\_PRIVATE\_NETWORK\`,
\`VIRTUAL\_NETWORK\_SRV6\_NETWORK\`, \`VIRTUAL\_NETWORK\_IP\_FABRIC\`,
\`VIRTUAL\_NETWORK\_SEGMENT\`, \`VIRTUAL\_NETWORK\_MANAGEMENT\`. Defaults to
\`VIRTUAL\_NETWORK\_SITE\_LOCAL\`.

Upstream description:

Different types of virtual networks understood by the system

Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL provides connectivity to public (outside)
network. This is an insecure network and is connected to public internet via NAT Gateways/firwalls
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created automatically and present on all sites
Virtual-network of type VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE is a private network inside site. It
is a secure network and is not connected to public network. Virtual-network of this type is local to
every site. Two virtual networks of this type on different sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on CE sites. This network is created during provisioning of site User defined per-site
virtual network. Scope of this virtual network is limited to the site. This is not yet supported
Virtual-network of type VIRTUAL\_NETWORK\_PUBLIC directly connects to the public internet.
Virtual-network of this type is local to every site. Two virtual networks of this type on different
sites are neither related nor connected.

Constraints: There can be atmost one virtual network of this type in a given site. This network type
is supported on RE sites only It is an internally created by the system. They must not be created by
user Virtual Networks with global scope across different sites in F5XC domain. An example global
virtual-network called "AIN Network" is created for every tenant. For F5 Distributed Cloud fabric

Constraints: It is currently only supported as internally created by the system. VK8s service
network for a given tenant. Used to advertise a virtual host only to vk8s pods for that tenant
Constraints: It is an internally created by the system. Must not be created by user VER internal
network for the site. It can only be used for virtual hosts with SMA\_PROXY type proxy Constraints:
It is an internally created by the system. Must not be created by user Virtual-network of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE\_OUTSIDE represents both VIRTUAL\_NETWORK\_SITE\_LOCAL and
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE

Constraints: This network type is only meaningful in an advertise policy When virtual-network of
type VIRTUAL\_NETWORK\_IP\_AUTO is selected for an endpoint, VER will try to determine the network
based on the provided IP address

Constraints: This network type is only meaningful in an endpoint

VoltADN Private Network is used on F5 Distributed Cloud RE(s) to connect to customer private
networks This network is created by opening a support ticket

This network is per site srv6 network VER IP Fabric network for the site. This Virtual network type
is used for exposing virtual host on IP Fabric network on the VER site or for endpoint in IP Fabric
network Constraints: It is an internally created by the system. Must not be created by user
Virtual-network of type VIRTUAL\_NETWORK\_SEGMENT for segment interface Virtual-network of type
VIRTUAL\_NETWORK\_MANAGEMENT is used for management purposes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "VIRTUAL_NETWORK_SITE_LOCAL",
  "enum": [
    "VIRTUAL_NETWORK_SITE_LOCAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE",
    "VIRTUAL_NETWORK_PER_SITE",
    "VIRTUAL_NETWORK_PUBLIC",
    "VIRTUAL_NETWORK_GLOBAL",
    "VIRTUAL_NETWORK_SITE_SERVICE",
    "VIRTUAL_NETWORK_VER_INTERNAL",
    "VIRTUAL_NETWORK_SITE_LOCAL_INSIDE_OUTSIDE",
    "VIRTUAL_NETWORK_IP_AUTO",
    "VIRTUAL_NETWORK_VOLTADN_PRIVATE_NETWORK",
    "VIRTUAL_NETWORK_SRV6_NETWORK",
    "VIRTUAL_NETWORK_IP_FABRIC",
    "VIRTUAL_NETWORK_SEGMENT",
    "VIRTUAL_NETWORK_MANAGEMENT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [SSH](resources--fleet--reference--group-002.md#canonical-1312203301110103-3130012113230101-2222302103001111-1223002010202130-2211020022130220-0002131113022003-3323311022010330-1232222133330110): complete subsection reference.

- [web_user_interface](resources--fleet--reference--group-002.md#canonical-2330030223021022-3303111302020012-1322230100321132-3313012120100100-2301301031112131-2333010110020231-2330221010013220-1231001021220120): complete subsection reference.

<a id="canonical-1020021031330313-1322223003133201-1201311211233323-2100301323330210-2311210123313000-0020311101230321-0121312321102122-0031221023311110"></a>

## Next pages — blocked_services / 012210011300 / 5

- [blocked_services.dns](resources--fleet--reference--group-002.md#canonical-3221332113111220-3010302110013032-3213220312323020-1000113322202101-2022231333313222-1303323113231210-0322133011123230-2012313233132100)
- [blocked_services.ssh](resources--fleet--reference--group-002.md#canonical-1312203301110103-3130012113230101-2222302103001111-1223002010202130-2211020022130220-0002131113022003-3323311022010330-1232222133330110)
- [blocked_services.web_user_interface](resources--fleet--reference--group-002.md#canonical-2330030223021022-3303111302020012-1322230100321132-3313012120100100-2301301031112131-2333010110020231-2330221010013220-1231001021220120)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3221332113111220-3010302110013032-3213220312323020-1000113322202101-2022231333313222-1303323113231210-0322133011123230-2012313233132100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303320322030203-0300003321202200-1010010103030310-1313131123003213-0220102210230113-0113011332113021-3211331221112312-3033003121230202"></a>

## blocked_services.DNS — DNS / 103022332131 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- blocked_services.DNS

<a id="canonical-3001003313321332-3330212121120221-2231001312302113-0130210130101030-1223311221031011-1202322303232120-3103302032023322-0021233022113320"></a>

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
dns = {}
```

<a id="canonical-0323223300312223-0000131112132213-0021111121113221-3332123312301123-3111330003313233-2232001123230022-2322123330010203-2202322301000033"></a>

## Direct properties — DNS / 103022332131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010102031002033-0120012012301110-0302123331131121-0301200232030123-3113112302302031-2031200333102201-2212113002103102-3022322021220202"></a>

## Next pages — DNS / 103022332131 / 4

- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1312203301110103-3130012113230101-2222302103001111-1223002010202130-2211020022130220-0002131113022003-3323311022010330-1232222133330110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201310323201323-0031120103032132-1332132022011112-2301010011300013-1022011132322003-1112110313123213-1213320023003302-1030220231332020"></a>

## blocked_services.SSH — SSH / 012223330311 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- blocked_services.SSH

<a id="canonical-0302223331101002-0213123231100103-3000030300221332-3300313302033202-0212203220321130-1330103300322313-0321101230120100-0002110202021030"></a>

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
ssh = {}
```

<a id="canonical-0000322222222320-1102210303101233-0203121321322033-3231333132000102-1133022022001330-1122233221001203-0333232301112031-0232003223002212"></a>

## Direct properties — SSH / 012223330311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300313333120230-2102020101130030-3212133303121112-1321102321210231-0031021001212201-2202111030310133-0021322212232130-3322313223012331"></a>

## Next pages — SSH / 012223330311 / 4

- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2330030223021022-3303111302020012-1322230100321132-3313012120100100-2301301031112131-2333010110020231-2330221010013220-1231001021220120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300033223133130-3211102030000321-0133311102130023-1102111012323023-2112321020122302-1012220233311212-3010032001121000-1120233113331300"></a>

## blocked_services.web_user_interface — web_user_interface / 313223133230 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- blocked_services.web_user_interface

<a id="canonical-1123031312320312-1220130012132210-0133032010322123-2101323021222330-2023233200223023-2311122321002303-0030311213221031-2003303232021033"></a>

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
web_user_interface = {}
```

<a id="canonical-3323203233330300-2312220313131202-0023030032101002-1331022033220120-3313033231003211-0022221311213332-2021301132000021-3022023322132003"></a>

## Direct properties — web_user_interface / 313223133230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011132232301130-1210333121223230-0023021231001311-0110112133233113-0322123011231112-1031112111211111-0001130020201323-0100200212023302"></a>

## Next pages — web_user_interface / 313223133230 / 4

- [blocked_services](resources--fleet--reference--group-001.md#canonical-1222211313112030-3230021132100033-0021111131332231-1222310321002303-3231102120213230-0221223030331102-2233211121031201-2103100311211233)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301211031130203-1030321323220031-3210332102301103-2011022210020323-0103122021323120-0312303033130200-0000313300000233-2120112101201122"></a>

## bond_device_list — bond_device_list / 201332113011 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- bond_device_list

<a id="canonical-3112331033303132-1010323223233021-0001113321020203-0020000321010310-0312210031322111-2212121113112120-0012220002121030-2332000033322102"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: bond\_device\_list, no\_bond\_devices; Default: no\_bond\_devices\] Bond Devices List. List
of bond devices for this fleet.

Upstream description:

List of bond devices for this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("bond_devices")}
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

- [bond_device_list](resources--fleet--reference--group-002.md#canonical-3112331033303132-1010323223233021-0001113321020203-0020000321010310-0312210031322111-2212121113112120-0012220002121030-2332000033322102)
- [no_bond_devices](resources--fleet--reference--group-002.md#canonical-2233202132112121-1300211023330221-2212133210020112-2021322122322312-0321223001202132-3220130012222211-3210110102122303-1332222011023230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
bond_device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333320023320022-2313322103001230-2213001221222112-0113312321322311-1311030202310332-0101330013000313-0221202110220000-0223223332030010"></a>

## Direct properties — bond_device_list / 201332113011 / 3

- [bond_devices](resources--fleet--reference--group-002.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113): complete subsection reference.

<a id="canonical-1313123231031103-3300310223133300-0213002200222200-1003232212000323-3100123313323313-2112220213232300-2030113021031230-0222033221131200"></a>

## Next pages — bond_device_list / 201332113011 / 4

- [bond_device_list.bond_devices](resources--fleet--reference--group-002.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3012131200112130-2220133030113232-0310222113211313-0113312033030201-2130120020300030-3313003312002330-2333131232200130-2223210123211223"></a>

## bond_device_list.bond_devices — bond_devices / 031030221032 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [bond_device_list](resources--fleet--reference--group-002.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- bond_device_list.bond_devices

<a id="canonical-3233230000001102-0311301000203233-2111011133110120-3233103113333020-3021302233123110-2002020302302210-3331023230132110-2230111323213220"></a>

Type: `"object"`. list nested block, Optional.

Bond Devices. List of bond devices.

Upstream description:

List of bond devices.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("devices",
    "link_polling_interval",
    "link_up_delay",
    "name"),
  validators.ConflictingListObjectAttributes("active_backup",
    "lacp")}
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
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
bond_devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-0000212010022101-2231302230031331-3223213113221011-2323113103320010-1302010130333332-0210010231203312-2203120213211022-1032333223101212"></a>

## Direct properties — bond_devices / 031030221032 / 3

- [active_backup](resources--fleet--reference--group-002.md#canonical-2320220031020231-2330123331333300-0121030331313323-1311323222233031-3201020003201011-3111210101112123-2301231002321223-1311303002311002): complete subsection reference.

<a id="canonical-1102101300033333-1222302001120331-3001102112133100-1113221021010223-1011033102031323-2001201332011123-3021100212311201-3021223202312132"></a>

<a id="canonical-3013220013022313-0231332320202000-3130210133221031-2312131013122312-0013103323223023-2120220123323132-2300010011023033-2301210010131221"></a>

## devices property — bond_devices / 031030221032 / 4

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

- [lacp](resources--fleet--reference--group-002.md#canonical-1232131310312013-1102232002320220-3200203002001023-0003223112323331-1230230223333130-2100022101102310-1212322333003213-1131321102112223): complete subsection reference.

<a id="canonical-3220201301033322-3120000332322333-0213203203003312-3232200020010033-2303012020310131-2302311233300131-0000013010231312-3013130100312323"></a>

<a id="canonical-3102112213200122-2211130202020203-2320001112102110-2121222011331313-0303232302020300-2320213001211321-2020220110202001-1120111023303303"></a>

## link_polling_interval property — bond_devices / 031030221032 / 5

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

<a id="canonical-2112131103330123-0303001101002111-1102112103200311-1113220212110022-2303210120033112-0031200132110330-1222020030102232-3102030331313323"></a>

<a id="canonical-1003301030310212-0332212031232101-0111013023221100-0021201003320202-2320020300102121-2123113120021030-0120211213330232-0202311133312321"></a>

## link_up_delay property — bond_devices / 031030221032 / 6

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

<a id="canonical-2133233103031220-1330321001113020-1103312233020121-0100122200001233-1103032120012000-3130012302010322-3313003331032211-2100312203130113"></a>

<a id="canonical-1013212130121001-0110201120000301-3111100233123013-1300212333101301-1231130230231233-2030003102300313-3123031001020233-3233302101032211"></a>

## name property — bond_devices / 031030221032 / 7

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

<a id="canonical-0121013220231210-1201221013321012-1332031333021300-0321312313033001-0213130210311333-2023022130101313-0033303030223213-1033333200110031"></a>

## Next pages — bond_devices / 031030221032 / 8

- [bond_device_list.bond_devices.active_backup](resources--fleet--reference--group-002.md#canonical-2320220031020231-2330123331333300-0121030331313323-1311323222233031-3201020003201011-3111210101112123-2301231002321223-1311303002311002)
- [bond_device_list.bond_devices.lacp](resources--fleet--reference--group-002.md#canonical-1232131310312013-1102232002320220-3200203002001023-0003223112323331-1230230223333130-2100022101102310-1212322333003213-1131321102112223)
- [bond_device_list](resources--fleet--reference--group-002.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2320220031020231-2330123331333300-0121030331313323-1311323222233031-3201020003201011-3111210101112123-2301231002321223-1311303002311002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332103302100302-1220311013001323-0330312323113013-2110320101023131-3231032330332132-3001112302123011-0200123302001220-2023233103030022"></a>

## bond_device_list.bond_devices.active_backup — active_backup / 300231321300 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [bond_device_list](resources--fleet--reference--group-002.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- [bond_device_list.bond_devices](resources--fleet--reference--group-002.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- bond_device_list.bond_devices.active_backup

<a id="canonical-2132113223210212-3220131323133320-0112021122002113-1331132211002101-2212032312213130-1230303323303103-3132122000332110-3231200233220012"></a>

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

<a id="canonical-0030220231100302-3012303320330303-3031202333033313-0330301011131310-2001111300333203-0021123212203212-3232002210020112-3222233030011331"></a>

## Direct properties — active_backup / 300231321300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1212313202000002-3113230033023330-2232333121003031-2122313123211011-3112011303202103-1011303113312313-3203030210021032-0113023020312210"></a>

## Next pages — active_backup / 300231321300 / 4

- [bond_device_list.bond_devices](resources--fleet--reference--group-002.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1232131310312013-1102232002320220-3200203002001023-0003223112323331-1230230223333130-2100022101102310-1212322333003213-1131321102112223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100331322111130-0331301312212221-3310121312013202-2303123123113033-0213110303130233-2122012220022013-1301032123221130-1100202301232102"></a>

## bond_device_list.bond_devices.lacp — lacp / 023032002233 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [bond_device_list](resources--fleet--reference--group-002.md#canonical-1101302100313120-2212133323013032-0201223201313230-1122203333102023-2010323123100003-2332211203232322-2313210102331313-0120210021230331)
- [bond_device_list.bond_devices](resources--fleet--reference--group-002.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- bond_device_list.bond_devices.lacp

<a id="canonical-3313323230011200-3112222003013333-1332013011020223-2101133212133213-0133030103102010-3231322320010231-0110131020223102-3201102113332030"></a>

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

<a id="canonical-3333012031032231-2020303302222323-3300003130313001-0331011232201220-3313211321231123-2222232021010300-3010331213320102-2231301131101120"></a>

## Direct properties — lacp / 023032002233 / 3

<a id="canonical-0121011021031200-1203113213112302-2311030033311221-0001200330000302-2302020222001220-1313333320120131-2133220231103311-3311213310231231"></a>

<a id="canonical-0101203212122313-1213001002010123-3033331333031301-0123232210301230-2311301022110021-3203323012103321-2122323312022213-3301010300101203"></a>

## rate property — lacp / 023032002233 / 4

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

<a id="canonical-1333312021201310-0311302103020322-1111322223311101-1212312211303033-2221123210033023-1023120012232200-0003330111002303-1303233031001330"></a>

## Next pages — lacp / 023032002233 / 5

- [bond_device_list.bond_devices](resources--fleet--reference--group-002.md#canonical-3020331033311102-0020213233201230-3031313012333131-1001102032213002-1323132210300221-2003023230101330-3011302033002012-1221333213130113)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2220011112301110-2101322302212032-2331222020113323-3333113013301010-2200210103313131-3003210110031021-1011111332222330-3232220020111203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133111233323012-3333032003030323-0311313333210011-3032030100220223-1003200200020322-3232320103321012-1313000333011022-2201222013201100"></a>

## dc_cluster_group — dc_cluster_group / 331132033221 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- dc_cluster_group

<a id="canonical-3222132011302110-3223233110303022-0010111003310331-1232121213200230-0212010203312230-2220311223301131-1013331123331132-0011233232032100"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: dc\_cluster\_group, dc\_cluster\_group\_inside, no\_dc\_cluster\_group; Default:
no\_dc\_cluster\_group\] Type establishes a direct reference from one object(the referrer) to
another(the referred). Such a reference is in form of tenant/namespace/name.

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

- [dc_cluster_group](resources--fleet--reference--group-002.md#canonical-3222132011302110-3223233110303022-0010111003310331-1232121213200230-0212010203312230-2220311223301131-1013331123331132-0011233232032100)
- [dc_cluster_group_inside](resources--fleet--reference--group-002.md#canonical-3113032020311002-1032323211212323-0330330230100021-2112321321021100-2313331222110031-1213131220011002-2013013013100021-3210322011302131)
- [no_dc_cluster_group](resources--fleet--reference--group-002.md#canonical-0300111311003312-2320102111320110-2333221333222221-2321223000332000-0132220211220113-0221111301300211-3031311131013010-3001123212213133)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
dc_cluster_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-1233131231231232-0001011222032003-0110201011100221-1003113330311103-2231112203120011-2333332212323332-1212113300031102-0320300231011302"></a>

## Direct properties — dc_cluster_group / 331132033221 / 3

<a id="canonical-0101131011112223-1230130100133332-1220213300302321-2300002103000222-3121030302332303-2313202311120022-3112022200100121-2300011020133120"></a>

<a id="canonical-3123330313233211-3220333301132301-3233121132332300-2023113110103122-1320021112203031-2301022003103130-1120213203001133-2231122023130011"></a>

## name property — dc_cluster_group / 331132033221 / 4

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

<a id="canonical-2133010021202033-1133110231202332-3311230000223312-0103230123323001-2003333320312222-0331320310323013-0303101222201001-1323102123303313"></a>

<a id="canonical-2321022123120321-0303011333202121-2022120130201313-0211132223111202-1300031223332012-2300033203102212-2220132300311320-0100231330210301"></a>

## namespace property — dc_cluster_group / 331132033221 / 5

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

<a id="canonical-2023322113123002-3100320322200311-3001321002020130-3332211000200200-3223000030112330-3303213001022330-1221003223012102-0221221123233202"></a>

<a id="canonical-0220220210022322-1002313203211120-2202302303310000-2303322311223330-2201011202301022-2031313032120022-0002130231311011-1031002233000102"></a>

## tenant property — dc_cluster_group / 331132033221 / 6

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

<a id="canonical-2210200020312212-1022110301110032-0303221332213313-0100203211132212-3232002011212131-0303200101002310-1330212021033302-0231001012001123"></a>

## Next pages — dc_cluster_group / 331132033221 / 7

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1022322123131331-3133331231203323-3213011213101211-0310133103101302-3003011111020032-0112210323033012-3222000332033020-0103003311002033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332322033101033-3221132031322203-0133330230130330-1323330021203101-1022011311023100-1103130011010100-1031100313031323-3121333210211211"></a>

## dc_cluster_group_inside — dc_cluster_group_inside / 232212000021 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- dc_cluster_group_inside

<a id="canonical-3113032020311002-1032323211212323-0330330230100021-2112321321021100-2313331222110031-1213131220011002-2013013013100021-3210322011302131"></a>

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
dc_cluster_group_inside {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301110023202212-2313232033310030-0111120022133203-0131011331000123-3321213133100210-2021312323213103-2102332100222021-0313211223120003"></a>

## Direct properties — dc_cluster_group_inside / 232212000021 / 3

<a id="canonical-2011323303112120-1312121233303233-1311311230113121-0221023112223103-2320311333333331-0203021123121111-3322023001102331-1102313203031032"></a>

<a id="canonical-1313010022201301-2131322122122332-2030013022320310-1303012331200301-2012031222123202-3232011301302023-0300120222021203-0233213112000231"></a>

## name property — dc_cluster_group_inside / 232212000021 / 4

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

<a id="canonical-2100220030012020-0321102223000001-0102112221223102-1222121232320330-1020103111112312-1221001121120222-3010133001003222-0130001013213001"></a>

<a id="canonical-0113011323112302-0102131320110102-1131130011202031-2130101301331210-1102003123023200-3100302010230210-1231200121011300-0322113332312221"></a>

## namespace property — dc_cluster_group_inside / 232212000021 / 5

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

<a id="canonical-2012030333200310-0333322331231021-3020333102301233-1322113333332233-1121311121112010-2121323022000230-3023201203213011-2113201032312211"></a>

<a id="canonical-1033100022300003-2301111123001231-0021201101331232-1111200322312023-3302313032233001-0002201303030030-0000233231112123-2000011011112231"></a>

## tenant property — dc_cluster_group_inside / 232212000021 / 6

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

<a id="canonical-3213030322130220-3111000311003120-0003201222233013-1302130000333233-2332011130003123-1211333003212000-2031201320202311-2330110113302321"></a>

## Next pages — dc_cluster_group_inside / 232212000021 / 7

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3310002022222312-0033100302223102-0112021232310302-1233201223101302-0221021103321211-0101113223023222-3000011131123213-1121030200103313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3103011333103222-0012213123121031-0021113202030020-2101030231002221-2301132110003200-1311211232130100-2011202300001013-3123200030322110"></a>

## default_config — default_config / 203213203301 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- default_config

<a id="canonical-0222331031323120-0233230301313130-2222300312211303-2131330122312100-3030030103000312-3321300032012020-1221213113311200-3302320011013321"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_config, device\_list, interface\_list; Default: default\_config\] Enable this
option

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

- [default_config](resources--fleet--reference--group-002.md#canonical-0222331031323120-0233230301313130-2222300312211303-2131330122312100-3030030103000312-3321300032012020-1221213113311200-3302320011013321)
- [device_list](resources--fleet--reference--group-002.md#canonical-3200322202323330-0323133003013331-1122130122320213-3221210232101013-1200323220311002-0232323111303111-2110131211011210-0020213222312313)
- [interface_list](resources--fleet--reference--group-002.md#canonical-1022102312212002-1321333300332120-1002202302210311-1003202220302323-2023331213223200-0220022300211302-2130200100001300-3021321322102123)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_config = {}
```

<a id="canonical-2001202331002230-3230100011320301-3201200132200103-3310222013320323-2202022332212212-1010332101103221-2321211003003322-2330021010133110"></a>

## Direct properties — default_config / 203213203301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232132213111110-1100212003213121-3313221013300121-3320020112103130-2012211332313311-1020132321111022-1122132203012322-3011330330030230"></a>

## Next pages — default_config / 203213203301 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0321320013020113-2010101112101310-3001121223201313-1212300320113031-2233220102201110-2003223123110010-3213112211322133-3220222221031301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0122001121213310-0021023031032002-1330332313330122-1033120101303301-0020233111120023-0311200303110112-3230130113311203-3221211021012333"></a>

## default_sriov_interface — default_sriov_interface / 223333102032 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- default_sriov_interface

<a id="canonical-3120231221133033-3133113300203332-3300212000321212-2230302302231033-3110323323131002-3333002303210003-1111030102211010-3220300203133231"></a>

Type: `["object", {}]`. Optional.

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

- [default_sriov_interface](resources--fleet--reference--group-002.md#canonical-3120231221133033-3133113300203332-3300212000321212-2230302302231033-3110323323131002-3333002303210003-1111030102211010-3220300203133231)
- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0122221131232302-1301033322113301-0110223212011211-3233102320110033-0121020202101031-3330230010133021-2322320310031011-0300013133213110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sriov_interface = {}
```

<a id="canonical-0133233323012112-3301011022100033-1122131113013111-1223203332223210-0212331021021320-1322311103313213-2003030211130111-2023213220133301"></a>

## Direct properties — default_sriov_interface / 223333102032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333003120200210-1030130102030133-3120223210311020-2013223301022122-1331332003332321-2333103032210332-3123122130122222-0332222003210211"></a>

## Next pages — default_sriov_interface / 223333102032 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2311012123131011-1011230030130023-3102322012232302-1222231101331323-1232103111123230-0012022301103233-1003210001001231-0212032003000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010200230130031-0232332021001010-0123330112030231-2001321132222322-3102200001331232-1302022030110231-1023221331223003-3330313010022101"></a>

## default_storage_class — default_storage_class / 120112111000 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- default_storage_class

<a id="canonical-3232310021320301-3232001110331123-2122322312002112-3020302312022233-3221320130031123-1100300210011002-2302313031231330-2332332332120130"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_storage\_class, storage\_class\_list; Default: default\_storage\_class\]
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

OneOf alternatives in this subsection:

- [default_storage_class](resources--fleet--reference--group-002.md#canonical-3232310021320301-3232001110331123-2122322312002112-3020302312022233-3221320130031123-1100300210011002-2302313031231330-2332332332120130)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-2222300023023313-2122302313220111-2233300300012313-3303221100131111-0112033102321132-1321110322120221-0103222322331100-1100211021322120)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_storage_class = {}
```

<a id="canonical-0011021022212202-2003333003000033-3110112112031021-1013333010130211-0102302202112323-0212231221010303-0210111302320100-1111220131322032"></a>

## Direct properties — default_storage_class / 120112111000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2133312220210110-3011220012310300-0301202331301022-3312302032303232-0132220101302210-1002303000332213-1302022111023023-2233021110032012"></a>

## Next pages — default_storage_class / 120112111000 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2012100320202111-1003330223212301-3032022132110231-1312231300111330-1332100132012122-1310200213231333-3313103222231312-0121120321130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103313302221210-3132220002212130-0100330213212030-0121232012031101-3221303321313021-2131011122031002-3001311103013313-1201211132320330"></a>

## deny_all_usb — deny_all_usb / 202121333033 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- deny_all_usb

<a id="canonical-3031222031130102-0220323323303102-0001201301130320-0310333220032103-2020133202021100-3213120313202223-2013331220210232-1002130131230330"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
deny_all_usb = {}
```

<a id="canonical-2113033102033113-2322012031122133-3002223020030210-2301311201111000-1020133301333223-0303201013130323-1113222233101102-3110133303321030"></a>

## Direct properties — deny_all_usb / 202121333033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3103320123023011-2122010222220002-3012031321333322-3130032102203110-2221000202113030-2010201002013013-0230320101223110-0000111020323200"></a>

## Next pages — deny_all_usb / 202121333033 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3323323033213233-1123011032113102-0103212231301130-3022330121323232-0101223220230303-1321310030220033-3312013112331001-3131300220100131"></a>

## device_list — device_list / 330003232322 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- device_list

<a id="canonical-3200322202323330-0323133003013331-1122130122320213-3221210232101013-1200323220311002-0232323111303111-2110131211011210-0020213222312313"></a>

Type: `"object"`. single nested block, Optional.

Add device for all interfaces belonging to this fleet.

Receipt-pinned upstream constraints:

```json
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
device_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022313303310230-1110202332332022-3132311222223030-3013020030120332-2013200321312000-3020230330111102-0331201023131211-2321330001221131"></a>

## Direct properties — device_list / 330003232322 / 3

- [devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101): complete subsection reference.

<a id="canonical-1321120213031132-0002022010313323-1011030210033212-0111300030021001-3223133021323030-1221220130323222-0121330122122012-3031320320211033"></a>

## Next pages — device_list / 330003232322 / 4

- [device_list.devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3023103201023012-3113120120203300-3031203133221101-3302211211012210-0201221321331030-2111002010131231-3021313232312211-2230310221032011"></a>

## device_list.devices — devices / 300210220223 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- device_list.devices

<a id="canonical-0231031102013112-0023111223302101-2021310021322003-1200312221122032-1300223112001013-2331211030032020-3300200103100133-2033202112010213"></a>

Type: `"object"`. list nested block, Optional.

Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras,
scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet
and has an corresponding interface/device.

Upstream description:

Configuration for all devices in the fleet. Examples of devices are - network interfaces, cameras,
scanners etc. Configuration a device is applied on VER node if the VER node is member of this fleet
and has an corresponding interface/device. The mapping from device configured in fleet with
interface/device in VER node depends on the type of device and is documented in device instance
specific sections.

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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
devices {
  # Configure direct properties listed below.
}
```

<a id="canonical-2333133021230121-3321233033101122-2303313231112331-0032202022100100-2212110131120122-2102003122321101-3103120101210121-3112220111310303"></a>

## Direct properties — devices / 300210220223 / 3

<a id="canonical-1100231202123210-2011110011112001-2113330223130011-2212213222220313-3201110322220330-3000120231022133-1111321100131301-0003123202132031"></a>

<a id="canonical-2023111012332010-1213033212330212-1313133212022322-2102033130030011-1333213011303130-2201113103013122-0321112000012013-2001333313031133"></a>

## name property — devices / 300210220223 / 4

Type: `"string"`. Optional.

Name of the device including the unit number (e.g. Eth0 or disk1). The name must match name of
device in host-OS of node.

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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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

- [network_device](resources--fleet--reference--group-002.md#canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310): complete subsection reference.

<a id="canonical-2211213033230302-1031233322012201-3022301211020123-0221130310223300-2010203301032032-3020303323311103-2233211130100012-1112111100021233"></a>

<a id="canonical-2112303131331001-3321202312032321-1301131332100013-2230021010201133-2201021033303012-3102133031102112-2213222302201023-0323111031111021"></a>

## owner property — devices / 300210220223 / 5

Type: `"string"`. Optional.

\[Enum:
DEVICE\_OWNER\_INVALID|DEVICE\_OWNER\_VER|DEVICE\_OWNER\_VK8S\_WORK\_LOAD|DEVICE\_OWNER\_HOST\]
Defines ownership for a device. Device owner is invalid Device is owned by VER pod. Usually it will
be network interface device or accelerator like crypto engine. Possible values are
\`DEVICE\_OWNER\_INVALID\`, \`DEVICE\_OWNER\_VER\`, \`DEVICE\_OWNER\_VK8S\_WORK\_LOAD\`,
\`DEVICE\_OWNER\_HOST\`. Defaults to \`DEVICE\_OWNER\_INVALID\`.

Upstream description:

Defines ownership for a device.

Device owner is invalid Device is owned by VER pod. Usually it will be network interface device or
accelerator like crypto engine. Device is available to be owned by vK8s workload on the site, like
camera GPU etc. Device is not available to be owned by vK8s or VER. Can be exposed via some other
service. Like TPM.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "DEVICE_OWNER_INVALID",
  "enum": [
    "DEVICE_OWNER_INVALID",
    "DEVICE_OWNER_VER",
    "DEVICE_OWNER_VK8S_WORK_LOAD",
    "DEVICE_OWNER_HOST"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1033322232301001-3233210121312110-3310122311112333-3113132212223313-1330232031031200-1331302320023302-3122023303333303-2132001013211313"></a>

## Next pages — devices / 300210220223 / 6

- [device_list.devices.network_device](resources--fleet--reference--group-002.md#canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111320102223212-3230313012322030-3133110023103100-3302322320100211-1100113031311022-0020320302320323-3001030103313230-3302130212321111"></a>

## device_list.devices.network_device — network_device / 333332033003 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- [device_list.devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101)
- device_list.devices.network_device

<a id="canonical-1233122212201230-3130021332033113-3121011121310113-3221100021110321-3012311120233122-0100013233333031-0122123300302103-3222001210333012"></a>

Type: `"object"`. single nested block, Optional.

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Upstream description:

Represents physical network interface. The 'interface' reference points to a Network Interface
object. Attributes such as Labels, MTU from Network Interface must be applied to the device.

Device mapping to nodes

A fleet can have many devices and nodes in VER customer edge site can have many interfaces. An
interface in node inherits configuration from a device by matching, &#8203;- device\_name in Network
Interface for the device &#8203;- device name for physical-interface in the node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interface")}
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
network_device {
  # Configure direct properties listed below.
}
```

<a id="canonical-1023323032122222-3021300302001121-3312030222111101-1321210103230110-0332112232102223-2302110130210322-0022322033031122-1330122110300211"></a>

## Direct properties — network_device / 333332033003 / 3

- [interface](resources--fleet--reference--group-002.md#canonical-2121333001221000-2320111310332123-3212003300122303-3001303200202330-3121002213211000-0230301230121010-1120321110201001-0223031210312232): complete subsection reference.

<a id="canonical-2231211212012001-1230322122322023-2101323312122103-0023012020120213-1103321033103103-2130100023103020-3130312223002122-3120003011331310"></a>

<a id="canonical-3031223331212010-2200110100130102-0321313100122303-2130210011232120-0032031112303001-3233100101232120-3313011110332000-3120030230020211"></a>

## use property — network_device / 333332033003 / 4

Type: `"string"`. Optional.

\[Enum:
NETWORK\_INTERFACE\_USE\_REGULAR|NETWORK\_INTERFACE\_USE\_OUTSIDE|NETWORK\_INTERFACE\_USE\_INSIDE\]
Defines how the device is used If networking device is owned by VER, it is available for users to
configure as required If networking device is owned by VER, it is included in bootstrap config and
member of outside network. If networking device is owned by VER, it is included in bootstrap
config.. Possible values are \`NETWORK\_INTERFACE\_USE\_REGULAR\`,
\`NETWORK\_INTERFACE\_USE\_OUTSIDE\`, \`NETWORK\_INTERFACE\_USE\_INSIDE\`. Defaults to
\`NETWORK\_INTERFACE\_USE\_REGULAR\`.

Upstream description:

Defines how the device is used

If networking device is owned by VER, it is available for users to configure as required If
networking device is owned by VER, it is included in bootstrap config and member of outside network.
If networking device is owned by VER, it is included in bootstrap config and member of inside
network.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "NETWORK_INTERFACE_USE_REGULAR",
  "enum": [
    "NETWORK_INTERFACE_USE_REGULAR",
    "NETWORK_INTERFACE_USE_OUTSIDE",
    "NETWORK_INTERFACE_USE_INSIDE"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0311221331103221-3212323110323222-1222323002222232-1311113322320301-1110231301231131-1311000202332101-3323131032203021-0112010310221232"></a>

## Next pages — network_device / 333332033003 / 5

- [device_list.devices.network_device.interface](resources--fleet--reference--group-002.md#canonical-2121333001221000-2320111310332123-3212003300122303-3001303200202330-3121002213211000-0230301230121010-1120321110201001-0223031210312232)
- [device_list.devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2121333001221000-2320111310332123-3212003300122303-3001303200202330-3121002213211000-0230301230121010-1120321110201001-0223031210312232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3312322310002120-0223120102011010-2212021333000122-1200002033002313-0113031321121110-3200332112032112-0330221301230210-3322200022210011"></a>

## device_list.devices.network_device.interface — interface / 212033001101 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [device_list](resources--fleet--reference--group-002.md#canonical-3120302001020302-1323301100302021-0133230333003322-3232201013100103-3132100130110221-0202113211021302-2010010333112200-0132100222000303)
- [device_list.devices](resources--fleet--reference--group-002.md#canonical-2332302030300023-0030012332113010-1332021002323222-3111212333201000-0011332200312220-0212130331121122-1202022210320310-2202333330020101)
- [device_list.devices.network_device](resources--fleet--reference--group-002.md#canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310)
- device_list.devices.network_device.interface

<a id="canonical-3100120003312303-0003110011302032-3320031212231121-2220210003310112-2103210012021210-3102323323031331-0133312302300322-1002303103203112"></a>

Type: `"object"`. list nested block, Optional.

Network Interface attributes for the device. User network interface configuration for this network
device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in
VER node If network interface refers to a virtual-network, the virtual-netowrk type must be..

Upstream description:

Network Interface attributes for the device. User network interface configuration for this network
device. Attributes like labels, MTU from the 'interface' are applied to corresponding interface in
VER node If network interface refers to a virtual-network, the virtual-netowrk type must be
consistent with use attribute given below If use is NETWORK\_INTERFACE\_USE\_REGULAR, the
virtual-network must be of type VIRTUAL\_NETWORK\_SITE\_LOCAL or
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE if use is NETWORK\_INTERFACE\_USE\_OUTSIDE, the
virtual-network must of type VIRTUAL\_NETWORK\_SITE\_LOCAL if use is
NETWORK\_INTERFACE\_USE\_INSIDE, the virtual-network must of type
VIRTUAL\_NETWORK\_SITE\_LOCAL\_INSIDE.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
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

<a id="canonical-1020231010033013-0120012312231232-2022220221120033-0322333310223333-1012123032301222-2122321323223223-3011111101330322-0100201132000202"></a>

## Direct properties — interface / 212033001101 / 3

<a id="canonical-1012010012103110-3030010120023321-1000231213302133-3332203032230213-0132012311012310-3320232210202323-3103000102220023-1013202021123323"></a>

<a id="canonical-0123020112201132-2133302302302331-3100111322011232-2231200302223222-0303220100221333-2230312231031000-2322023010330001-2220120300132111"></a>

## kind property — interface / 212033001101 / 4

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

<a id="canonical-3233333311000113-0020001313212012-1200231301321133-1222110310101332-1022220111201131-1012201330230012-1311132232333331-2033201310000320"></a>

<a id="canonical-1121211202120122-1213301220322032-0022123321020032-2332233303011320-3101333103312203-3203121202201021-2002010102202022-3330011132211020"></a>

## name property — interface / 212033001101 / 5

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

<a id="canonical-2213021333121212-1021213211210102-0122021231233013-3310220120110300-3122232133132230-2013200320003131-1021210103032312-0202312102310021"></a>

<a id="canonical-2120211022032213-1301212331200223-3200201121223121-3233102330332230-0122132320303011-0003030333222031-0333310121012321-2312003020230003"></a>

## namespace property — interface / 212033001101 / 6

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

<a id="canonical-0230321303203311-1001303000233233-3313032111020212-3230212302231332-0223323100130231-3000333322323200-3001030112332013-0301123220030223"></a>

<a id="canonical-1121202010210100-3303222301311302-3332230121213323-2323110320330133-3122002232003231-2121333023023233-1030333302320032-1113232321223133"></a>

## tenant property — interface / 212033001101 / 7

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

<a id="canonical-0233221033313310-3010220003120213-2212031120312101-0220202203310302-0331302111111313-0012001102223032-3333320201013310-1123322110301202"></a>

<a id="canonical-3312323312320222-3113202313131213-3303311233220301-2210110100101131-2230212331302210-2013132021131310-0002102031021113-3333210122302112"></a>

## uid property — interface / 212033001101 / 8

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

<a id="canonical-3320310022202301-3322212013213031-3030313130112301-3013031223232111-1310112013210101-2030332011332220-1311113031311310-0201030132133121"></a>

## Next pages — interface / 212033001101 / 9

- [device_list.devices.network_device](resources--fleet--reference--group-002.md#canonical-0223202111131211-1231231021213021-1232010223323321-3231331003132222-1121233123332331-1203102320221121-0113210302011221-1003231132020310)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0123032211233331-0131332231333212-3200202011200112-1001000300113120-0122200030312231-1201011230021221-0213033021011323-3032133220031003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103333330002013-1212210112131101-2123002032222333-0131123313332103-2221202323320202-2212221331020220-2120201031222310-3202311322113332"></a>

## disable_gpu — disable_gpu / 321320210030 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- disable_gpu

<a id="canonical-0213033021223211-2011031122232000-0031331130102311-0223100313001222-0032002333002023-3222030023220001-1220103333102311-3323302230201330"></a>

Type: `["object", {}]`. Optional.

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

- [disable_gpu](resources--fleet--reference--group-002.md#canonical-0213033021223211-2011031122232000-0031331130102311-0223100313001222-0032002333002023-3222030023220001-1220103333102311-3323302230201330)
- [enable_gpu](resources--fleet--reference--group-002.md#canonical-2200303220000331-1001222020122121-1313202222301130-3330333133022123-3302320023303020-0201020112330320-2301302111030103-3220132132312202)
- [enable_vgpu](resources--fleet--reference--group-002.md#canonical-2022102100100133-3233222033211321-3200233210103333-1131101321130311-2120202203330101-1302300011010321-3003031331100131-3122303110120030)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_gpu = {}
```

<a id="canonical-3032000020123213-3011123132031321-3100121200133211-3120211110132322-0112301030310030-3021031223112333-1000213232310200-0120103130110331"></a>

## Direct properties — disable_gpu / 321320210030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013000203031320-0000032210321122-0123210210300023-1303011201210301-2222031112320131-0320011323200233-2001011010122022-3220233222301300"></a>

## Next pages — disable_gpu / 321320210030 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0030313313212311-2001301220310203-0020001123321003-2121032212103232-1131303332211222-3020011200213223-1213020113230210-0100321121202200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122123322100222-1013030311130230-0222000310332131-0310131011221020-3312122230301000-2313112221312221-0100200001112012-1010100203112100"></a>

## disable_log_anonymization — disable_log_anonymization / 300302110130 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- disable_log_anonymization

<a id="canonical-2121001302333101-3211331033211213-3233213331231013-2301210300232123-0033110031101222-0312002133222130-0120100323223301-0331300320023001"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_log\_anonymization, enable\_log\_anonymization; Default:
disable\_log\_anonymization\] Configuration parameter for disable log anonymization.

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

- [disable_log_anonymization](resources--fleet--reference--group-002.md#canonical-2121001302333101-3211331033211213-3233213331231013-2301210300232123-0033110031101222-0312002133222130-0120100323223301-0331300320023001)
- [enable_log_anonymization](resources--fleet--reference--group-002.md#canonical-3200320200030321-3033013222300332-2010230132310221-3032313333233322-0332313113200121-1201222331033303-3120202212101301-3020103131101003)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_log_anonymization = {}
```

<a id="canonical-3021100333332311-1221320103102012-0311113232033030-2112000302230211-3133220302110123-2211223012233112-2220131223132200-1101202103321100"></a>

## Direct properties — disable_log_anonymization / 300302110130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230301320113030-1303310003112331-0323330210121001-3020111231121232-0232022302332100-0300013323131133-0202211230100223-2003223101302122"></a>

## Next pages — disable_log_anonymization / 300302110130 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0321001020131001-1211220132331302-1322103320333300-0031202311220102-1230303330102113-1311000031300313-1122113011301220-2122300130120201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2031301333122302-3200231323022210-3200112321211000-0310201201033231-2223222113131221-2300333220011203-1023012130303102-1210302122331302"></a>

## disable_vm — disable_vm / 333112313313 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- disable_vm

<a id="canonical-3230331213222102-3010132211202323-0302300222121300-1322023031220311-3323322213002302-1122103131213313-1321302120133121-2211100120320133"></a>

Type: `["object", {}]`. Optional.

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

- [disable_vm](resources--fleet--reference--group-002.md#canonical-3230331213222102-3010132211202323-0302300222121300-1322023031220311-3323322213002302-1122103131213313-1321302120133121-2211100120320133)
- [enable_vm](resources--fleet--reference--group-002.md#canonical-0300310111033333-3323012300233121-0223320333212221-2033211212300122-0320122222321001-2130221000220223-1131213121203130-2103020131022232)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_vm = {}
```

<a id="canonical-1132113001203110-3202313000120032-3300330232333301-0111102011220130-1102220203032032-0012011112021333-2110111300012333-3030011211302302"></a>

## Direct properties — disable_vm / 333112313313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3131312023311012-1003201210033231-0110022121301203-1030212032331203-2320321211312112-2302100102313200-3221120210230101-0013311230133211"></a>

## Next pages — disable_vm / 333112313313 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0002001130122322-1321022131203221-1122200222031023-1020213032321102-0222000230321311-3133232323322212-0110120030211330-3132231013210310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1122323131313203-1030120032121312-0223230012323232-1032011210033112-0321010102131031-0102201221210210-0103320302111221-1013210300022322"></a>

## enable_gpu — enable_gpu / 010110212123 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_gpu

<a id="canonical-2200303220000331-1001222020122121-1313202222301130-3330333133022123-3302320023303020-0201020112330320-2301302111030103-3220132132312202"></a>

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
enable_gpu = {}
```

<a id="canonical-1101121113301113-1311303120001013-3033331233123303-3113212130330013-1223221022122321-3312303121132212-0220313032333103-1121213312123230"></a>

## Direct properties — enable_gpu / 010110212123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023201222310330-3031212101212313-0000120123301102-1111123131120323-3121323121311123-3212311111011223-0232311112233232-1001201021310132"></a>

## Next pages — enable_gpu / 010110212123 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1300130001320220-1113103202210133-1323332200222222-1011232231123110-0001121122301303-2210010201302013-1133102312333222-1032303221030230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0330121213131103-0003213331300221-1231011301113303-2201231131313210-0302230131332233-3002202301022031-0322010333112323-3320332210221120"></a>

## enable_log_anonymization — enable_log_anonymization / 211131132020 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_log_anonymization

<a id="canonical-3200320200030321-3033013222300332-2010230132310221-3032313333233322-0332313113200121-1201222331033303-3120202212101301-3020103131101003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable log anonymization.

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
enable_log_anonymization = {}
```

<a id="canonical-3111113230121110-3201001210310230-3202301030320213-1303233133230010-2123300221011202-1232012223321311-3000012101300210-2330332221131132"></a>

## Direct properties — enable_log_anonymization / 211131132020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102123003003133-2212030100032123-1133301031221303-1300100022001333-3323000001232121-3011022321131112-1312303103330312-1110122201132223"></a>

## Next pages — enable_log_anonymization / 211131132020 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0023220101031001-1331310302022102-0213011031321233-3112103310211311-1212211111023103-3221232112323313-3332022200013311-1333313210022010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0103121101121101-0133302310120003-2303111201223123-3130032122301030-2120002102232112-3320110121200112-3213221100012332-0301333310123311"></a>

## enable_vgpu — enable_vgpu / 121102100223 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_vgpu

<a id="canonical-2022102100100133-3233222033211321-3200233210103333-1131101321130311-2120202203330101-1302300011010321-3003031331100131-3122303110120030"></a>

Type: `"object"`. single nested block, Optional.

Licensing configuration for NVIDIA vGPU.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("server_port")}
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
enable_vgpu {
  # Configure direct properties listed below.
}
```

<a id="canonical-1020112021030032-1331012201011012-1110201331000103-0210032313230113-1322132302330230-2022222231332111-0120221010102302-1132203011033133"></a>

## Direct properties — enable_vgpu / 121102100223 / 3

<a id="canonical-2023210113311333-0320211122000010-1023321012002001-3210010223100003-1321020212111320-0122033120111313-0332033011031222-0102020333322310"></a>

<a id="canonical-1103021010111000-2302203200103330-2331212233100333-2322100113302210-1203002022310110-3010211113002120-0112020020123101-1310311030333011"></a>

## feature_type property — enable_vgpu / 121102100223 / 4

Type: `"string"`. Optional.

\[Enum: UNLICENSED|VGPU|VWS|VCS\] Set feature to be enabled Operate with a degraded vGPU performance
Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation Enable NVIDIA Virtual Compute Server.
Possible values are \`UNLICENSED\`, \`VGPU\`, \`VWS\`, \`VCS\`. Defaults to \`UNLICENSED\`.

Upstream description:

Set feature to be enabled

Operate with a degraded vGPU performance Enable NVIDIA vGPU Enable NVIDIA RTX Virtual Workstation
Enable NVIDIA Virtual Compute Server.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("UNLICENSED",
    "VGPU",
    "VWS",
    "VCS"),
}
```

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

<a id="canonical-1023131222202213-0331220323130333-0131111200110203-2033123101001111-3121102110021223-0101032120223231-0121001010101202-3330220300231212"></a>

<a id="canonical-2101232012110123-3000313022331302-1123132010202103-0333030312320100-3331213213223312-1323321320300102-0310021100311000-1300321012232130"></a>

## server_address property — enable_vgpu / 121102100223 / 5

Type: `"string"`. Optional.

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
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname_or_ip": "true"
  }
}
```

<a id="canonical-0110302323212200-1232132013132232-0011133101132231-1011111220312201-2020020031121032-0320011033100232-0210111310221232-0201022112313332"></a>

<a id="canonical-1312031010123022-3321331221001010-0210303023323113-1203101032322130-2010100211102130-1231110303313003-0211000302202111-2122020101113223"></a>

## server_port property — enable_vgpu / 121102100223 / 6

Type: `"number"`. Optional.

License Server Port Number. Set License Server port number.

Upstream description:

Set License Server port number.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3002212212003013-1133332220020000-2232131203120031-1302001010100221-0320303323020001-1012103121213202-1221132133303123-2311012001201203"></a>

## Next pages — enable_vgpu / 121102100223 / 7

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0322133113031201-3211102313221333-0210203311332212-0110112033311332-0120013321023303-1301023021020300-0230103122131203-2030002233213303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0311332110013011-2213331223312312-2110331233213322-2121101200330012-0120302103033110-3233313032312102-1311200223301023-1231222301132003"></a>

## enable_vm — enable_vm / 311000231301 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- enable_vm

<a id="canonical-0300310111033333-3323012300233121-0223320333212221-2033211212300122-0320122222321001-2130221000220223-1131213121203130-2103020131022232"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_vm = {}
```

<a id="canonical-1130121032311302-0231202003310332-3302231231111020-0113210032220130-1322022121113320-3222201230102310-2102021202300102-3233231203222323"></a>

## Direct properties — enable_vm / 311000231301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211123221313303-2231112333223223-3010210002001132-1212303232032220-2110001330201312-0130221123212232-2020033020003101-3032010103233301"></a>

## Next pages — enable_vm / 311000231301 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3333200312232013-0113002003132011-3330301312033033-2120022120220222-1123100201032211-1221013003232322-1100332121322133-3130123311000222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230331331311310-2231100122013223-3110221110101000-1223202131313133-2131331030222312-2112030102323302-1012001020120320-0021231120011011"></a>

## inside_virtual_network — inside_virtual_network / 020211000203 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- inside_virtual_network

<a id="canonical-2122322131012313-3130301003123021-1212130323302111-0203212231233110-0003213231120130-0211202223132033-2230322211113002-3013330103212130"></a>

Type: `"object"`. list nested block, Optional.

Default inside (site local) virtual network for the fleet.

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

Terraform syntax:

```terraform
inside_virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323032112201003-1313202210302212-1101332211210201-2232031230101000-0321220330131001-2331100023113303-1321102212202203-0112233223100212"></a>

## Direct properties — inside_virtual_network / 020211000203 / 3

<a id="canonical-2230030320120311-1301002313331101-1033031031133302-3130011021332011-1203003022220020-1011231331303022-0010212112311232-2230211222220020"></a>

<a id="canonical-0132020303100313-1320003122113102-2120113320000311-3202033032031002-0021100322220122-1312101212011030-1103220002112202-1001330130230000"></a>

## kind property — inside_virtual_network / 020211000203 / 4

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

<a id="canonical-1011023102110022-2333212123222312-1202220112210103-3232201131211031-0133131312230321-2111002112331103-1120030101022102-3121000312223003"></a>

<a id="canonical-0311132202011001-3120001213033022-0303000003200202-0221300111130323-0031030201213001-3013232303103113-0310011001203333-3220030221000111"></a>

## name property — inside_virtual_network / 020211000203 / 5

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

<a id="canonical-2111022200220232-1310232023101322-2300133301202011-1320312010121333-2030313313300120-0303122210120001-0022301332210111-3123113233212033"></a>

<a id="canonical-1333233101030112-2321000320023300-2023102202121232-2203011123232102-0313231223110230-2020231112113210-0120202001322000-3333210311220301"></a>

## namespace property — inside_virtual_network / 020211000203 / 6

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

<a id="canonical-2022230213020013-0333210331102312-2310211033222213-3330313003233032-3102003110321220-3011030031032203-1210003232121031-2130033312131303"></a>

<a id="canonical-0103300132330132-3122233011231031-0011011321103103-3032101211111222-2230101131000002-2222100210303300-1013301033223211-0212102313200223"></a>

## tenant property — inside_virtual_network / 020211000203 / 7

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

<a id="canonical-3221222313112131-3131320132212203-3111322330230132-0023230111333013-3312113210322012-3021321332013130-1210013322202333-1213001110000322"></a>

<a id="canonical-2223301222312321-1133133211130121-3331231133213331-0103303001302132-0311022122032113-1202301120323113-1001323002123201-1002131321000221"></a>

## uid property — inside_virtual_network / 020211000203 / 8

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

<a id="canonical-0130022033301032-3022220102200013-0033100320322200-0012101123112121-1101032032221322-1002233131330120-0331122003303212-3212013012131201"></a>

## Next pages — inside_virtual_network / 020211000203 / 9

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212220211013133-3032000223210033-1301021011303210-2133102200130221-0103201030033133-2010021312201223-2332223332002311-0313210133112012"></a>

## interface_list — interface_list / 323011123223 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- interface_list

<a id="canonical-1022102312212002-1321333300332120-1002202302210311-1003202220302323-2023331213223200-0220022300211302-2130200100001300-3021321322102123"></a>

Type: `"object"`. single nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("interfaces")}
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
interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1201012201332130-0300322201123301-3313030032032232-0202002321023130-2101312103222003-2211032302322233-0213233011202300-1021113312023020"></a>

## Direct properties — interface_list / 323011123223 / 3

- [interfaces](resources--fleet--reference--group-002.md#canonical-0212032112031011-3111023103022010-0033202112313203-0310232033202222-0112031123021210-0030212302331233-0133212012233013-0111201133002333): complete subsection reference.

<a id="canonical-1121002021210021-3113323100003233-3201213110233013-1011333212220330-1200023113221112-2311303121111332-1131201100103110-3000110033110323"></a>

## Next pages — interface_list / 323011123223 / 4

- [interface_list.interfaces](resources--fleet--reference--group-002.md#canonical-0212032112031011-3111023103022010-0033202112313203-0310232033202222-0112031123021210-0030212302331233-0133212012233013-0111201133002333)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0212032112031011-3111023103022010-0033202112313203-0310232033202222-0112031123021210-0030212302331233-0133212012233013-0111201133002333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303111332000200-3133222303230002-2111211101223132-3300221103011300-1233010233030300-1031331033201110-3030331331203130-0123222122320101"></a>

## interface_list.interfaces — interfaces / 130332101313 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [interface_list](resources--fleet--reference--group-002.md#canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302)
- interface_list.interfaces

<a id="canonical-2231333023223222-2300022113210131-0103201221101230-1302323202012100-0130121122330333-1001010010030123-0302102232132101-1103210332202212"></a>

Type: `"object"`. list nested block, Optional.

Add all interfaces belonging to this fleet.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-3222131211033203-3313212021102112-3322131111003103-3003301222032130-1303223222301230-3012020100313110-2020301212103311-1010332103022100"></a>

## Direct properties — interfaces / 130332101313 / 3

<a id="canonical-0012002302301001-1302031001230033-0101120022101302-3220031301122001-3000330122031021-2112003213300001-0121301102032112-1110011131000000"></a>

<a id="canonical-2022010330212323-2210213300322201-2311311221131100-3121002003102020-2201233312023203-2313200011010323-1333302012100011-2211231131320213"></a>

## name property — interfaces / 130332101313 / 4

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

<a id="canonical-1000221230233223-2330200323332330-2233300233233020-0223011131213102-1113102002130023-0000211223313220-2312231210132112-0330130220230133"></a>

<a id="canonical-1000033212211200-1111130110122332-0222112122301101-2310131133110113-1023032003233012-3310123033103303-3032201130213213-3300000330302312"></a>

## namespace property — interfaces / 130332101313 / 5

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

<a id="canonical-0131003000100302-2302222210203211-1121332333131202-2003223101221013-1100033212023303-1333133232102301-0003312333312100-1131023003200221"></a>

<a id="canonical-0332032232322022-3033021320003321-0331201023133231-2213200322103222-2233010300230232-2313330021121220-0203220120302023-1233312133010111"></a>

## tenant property — interfaces / 130332101313 / 6

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

<a id="canonical-3202101013112021-0122033201302203-1333103332232133-2002223130012100-1123123212032301-0021231231002132-2102313332311311-3230021112210013"></a>

## Next pages — interfaces / 130332101313 / 7

- [interface_list](resources--fleet--reference--group-002.md#canonical-2033313231120123-0013102010311223-1030231302132002-1113111101332321-3323223200011101-2001323010002331-0223230320312012-1023032001121302)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131323200313301-2321000122202300-0030001233321212-1102201030113110-3210031003212033-1220121000103200-1112013222012022-1321311132222213"></a>

## kubernetes_upgrade_drain — kubernetes_upgrade_drain / 232010202031 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- kubernetes_upgrade_drain

<a id="canonical-2021130120130131-0131201122032111-0100011213212223-3331321110111101-3311000033113323-1332022231022230-0220132320100133-0221011220232023"></a>

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

<a id="canonical-0323123333233103-1033010233310333-3021212021133123-0310300012221010-2223012011233121-3232211330010110-2013102231321223-2213232032023112"></a>

## Direct properties — kubernetes_upgrade_drain / 232010202031 / 3

- [disable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-3220102311201123-3330002212131021-0130223202031202-1031023220131321-2103003012121133-2103201033221311-3023311122332132-3033021132202010): complete subsection reference.

- [enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122): complete subsection reference.

<a id="canonical-0310030100133223-0333210110301311-2222120312202030-2230300003321001-3113123322222131-0011000301133200-0332131230331133-3331322102130333"></a>

## Next pages — kubernetes_upgrade_drain / 232010202031 / 4

- [kubernetes_upgrade_drain.disable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-3220102311201123-3330002212131021-0130223202031202-1031023220131321-2103003012121133-2103201033221311-3023311122332132-3033021132202010)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3220102311201123-3330002212131021-0130223202031202-1031023220131321-2103003012121133-2103201033221311-3023311122332132-3033021132202010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2110323230333310-0223302221302030-0213003121312032-0311132233323112-2232212211302311-1031001201230202-2301011130103110-2203211301333021"></a>

## kubernetes_upgrade_drain.disable_upgrade_drain — disable_upgrade_drain / 221130200032 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-1230301023033301-2231012323222023-1022233031331311-2121301231200022-2032231011102332-1103112103023220-3100110302212212-3223221220113201"></a>

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

<a id="canonical-1332000122001202-0222110332021020-1022333010233012-2210003320120010-0111330330212101-0132311330321333-0030132013332023-0223021130123013"></a>

## Direct properties — disable_upgrade_drain / 221130200032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3100203032121300-0130103301011331-1110300333132333-0202310210111003-3113101231300213-0322002032102212-2302102010033113-2222031212201022"></a>

## Next pages — disable_upgrade_drain / 221130200032 / 4

- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331001200330323-2301032203020100-1210111203201310-2102313233332322-0200031032221003-3021302321111230-2002220112320312-1201001311300112"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain — enable_upgrade_drain / 333321000210 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-2200003321111012-2102312032003130-0331322023323100-1033223013033010-0132110322312003-2121320100203112-2220001120330032-0321201312231310"></a>

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

<a id="canonical-0100002013200220-3330331022002102-3220010230013023-3101303231200332-1001102101102032-1012010102302003-2020223331110322-0302110133110020"></a>

## Direct properties — enable_upgrade_drain / 333321000210 / 3

- [disable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-1133220013133121-3331112130002231-2321003331133202-0201202120332211-3033030202301313-3102023213011322-1120023112300210-0012231231231000): complete subsection reference.

<a id="canonical-0312303200220300-2110120203113210-0133330021211203-3022113320111311-2211100210022330-1322023011232212-0220103030102123-2031210133020101"></a>

<a id="canonical-0221211120220210-0220112013313323-3020310320211210-1331233020022133-3000333110033311-0031211023333211-2132013302333322-2220222230203313"></a>

## drain_max_unavailable_node_count property — enable_upgrade_drain / 333321000210 / 4

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

<a id="canonical-1330022231112010-3133223023113013-1310120132011001-0331200003013230-2032223111103331-3000102103002220-2302031111131021-1330001033131320"></a>

<a id="canonical-3010233222102103-1031212211313310-0203231221210133-2332111201201021-1313032030033200-0312332120311100-1300103120221133-2323131303330301"></a>

## drain_max_unavailable_node_percentage property — enable_upgrade_drain / 333321000210 / 5

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-3323030303221233-2100022310303231-0221302131302232-3011012213232112-0001210020220121-0302223122202021-3332000223033102-2023203132313122"></a>

<a id="canonical-1213223310021203-3131030200230101-3011330103030301-2021203202310230-2202220331321003-0230210311003010-1013033103320330-0012131121231002"></a>

## drain_node_timeout property — enable_upgrade_drain / 333321000210 / 6

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

- [enable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-3200231012222202-2330200310232322-0020012321110323-1102302030333121-2003110201110220-1131310233010112-0020202112123212-0302331102210222): complete subsection reference.

<a id="canonical-1013103230132131-1210103210202320-2131330330102011-0201013133003232-3330300120122212-1023002131033213-0002223313133012-1201012113220122"></a>

## Next pages — enable_upgrade_drain / 333321000210 / 7

- [kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-1133220013133121-3331112130002231-2321003331133202-0201202120332211-3033030202301313-3102023213011322-1120023112300210-0012231231231000)
- [kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode](resources--fleet--reference--group-002.md#canonical-3200231012222202-2330200310232322-0020012321110323-1102302030333121-2003110201110220-1131310233010112-0020202112123212-0302331102210222)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1133220013133121-3331112130002231-2321003331133202-0201202120332211-3033030202301313-3102023213011322-1120023112300210-0012231231231000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1230123213300330-3013132232332032-0200003101212112-3131303210223111-1221310330313303-3011201300020320-0313011031320221-0212211301200310"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode — disable_vega_upgrade_mode / 130320002030 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-3303103200232103-1120111120302102-2101210203313300-1121032310221130-1213002132012230-2233011102031033-0001200020220030-2212021010031011"></a>

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

<a id="canonical-2222133222211021-2302332231000320-0202120123231031-3131211132113331-1200200101100102-2201311001200231-1033103130321030-0100333022221323"></a>

## Direct properties — disable_vega_upgrade_mode / 130320002030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1011030213230122-2202111032130022-3302000122031100-2000131121102300-0032201022331122-1101003133133021-0313332313202222-2222022131132310"></a>

## Next pages — disable_vega_upgrade_mode / 130320002030 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3200231012222202-2330200310232322-0020012321110323-1102302030333121-2003110201110220-1131310233010112-0020202112123212-0302331102210222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010120001231010-0223030230310303-2310210002312101-0132233311110110-0121210313220320-3101023101313000-1021120001032003-0103131230320100"></a>

## kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode — enable_vega_upgrade_mode / 010223122221 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [kubernetes_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2303333022100230-0222322311231300-1201333201210221-1031333010332203-1132322322300001-2303102030023121-3011331010030130-2103210022002300)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-0001132232313221-1301212330111231-0201303232201201-0020323302213312-0210301323111031-3303033113100022-2131211311330102-2310322302233300"></a>

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

<a id="canonical-2232030123200320-2020202020201130-2232101213210133-0301320332320110-2211230203122230-1202333221200200-1132002000322132-1030133320213220"></a>

## Direct properties — enable_vega_upgrade_mode / 010223122221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013023312023000-0031220312302201-2300213113000302-1001310030212311-2302230131123010-1010100100113330-3210023013300130-0210201311130011"></a>

## Next pages — enable_vega_upgrade_mode / 010223122221 / 4

- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--fleet--reference--group-002.md#canonical-2220301303303231-1033000221330230-0032132020212001-2021103202111111-0120023223222303-3230100002133002-1001203311202332-3031100122113122)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0331330130330132-3110332130130011-0101221033032011-1321323001303111-1310212222202120-1000000110001011-0200233010121032-2221330200010033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123211333133112-3102212002213103-3031303310111001-1212222121303200-0320202112212203-3132312113302201-1331001101022211-0020320212220313"></a>

## log_receiver — log_receiver / 033131213012 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- log_receiver

<a id="canonical-3031313201013102-3012032012303320-0312101313130103-3322003301113303-1220323132023022-0301030301323133-2222033021303010-2022122130110203"></a>

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

- [log_receiver](resources--fleet--reference--group-002.md#canonical-3031313201013102-3012032012303320-0312101313130103-3322003301113303-1220323132023022-0301030301323133-2222033021303010-2022122130110203)
- [logs_streaming_disabled](resources--fleet--reference--group-002.md#canonical-2131023030210200-2311112100032332-0203003232113123-3101323130301101-1231112112221012-1033211311300001-2131011123301122-0232232301033323)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120230133321123-3230201130021300-3200203121333013-3202113221332000-2011100000023331-0102220101331033-2210013213101121-2221033000310120"></a>

## Direct properties — log_receiver / 033131213012 / 3

<a id="canonical-1130020113021020-3320232131202113-0302012111310003-0010301032313331-2123331221203013-2310211022130210-2303212231301010-3303012113202202"></a>

<a id="canonical-0112121030313303-1023002321332133-0133110321012312-1310313210010232-0312032022202121-3010113033332323-1021002203100003-1202200220013013"></a>

## name property — log_receiver / 033131213012 / 4

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

<a id="canonical-3203000220210021-2012022203121223-3231300120111030-2311331332031003-0123213133002111-1122201032101230-3010032103021231-3122320031310201"></a>

<a id="canonical-2310101330220213-0203202131230013-2111111003103222-1232031230133220-2000321030130330-3221331230301212-3123210233330322-2322201310200123"></a>

## namespace property — log_receiver / 033131213012 / 5

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

<a id="canonical-1112232211002303-3213210020311113-3230102202232032-3131313112232133-1330010131302300-3313312032212033-0302132231000212-0332111013211222"></a>

<a id="canonical-3003303022033221-2233113301310320-2000213010212111-3212133333312032-0331111133021201-1031213002300133-0121323220233212-2331033331212200"></a>

## tenant property — log_receiver / 033131213012 / 6

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

<a id="canonical-1322011101303301-2221023211313301-3101300220223112-2302110130203010-0231322003310032-2233320302302002-2221302023021301-3323022103103002"></a>

## Next pages — log_receiver / 033131213012 / 7

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3212213123002023-0202231202121031-0011310312010030-0202200201110132-0022133203021121-1110002111002320-0011010213122003-0312100022301211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320032020013101-1100231211301023-2020203211123120-0012320232112212-3330031131201121-1010303330200110-1033103320011231-0000012031320331"></a>

## logs_streaming_disabled — logs_streaming_disabled / 301021202031 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- logs_streaming_disabled

<a id="canonical-2131023030210200-2311112100032332-0203003232113123-3101323130301101-1231112112221012-1033211311300001-2131011123301122-0232232301033323"></a>

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

<a id="canonical-3110203010132221-2302300233013112-2010220030110222-2032010303310030-1121030210223313-0113231311211321-0133313000321212-1312131033202021"></a>

## Direct properties — logs_streaming_disabled / 301021202031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223030000013203-1222213011200233-3313113330233311-2002222231111231-1131301120310201-2032022133110212-3023321210030101-2222103130030001"></a>

## Next pages — logs_streaming_disabled / 301021202031 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3323111111030122-1303021120311223-1110003221103022-2020332013330231-3132231113011220-1121100020211332-0213211322111100-3303110210031121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302330323030330-3210110300133132-1321230300222320-3023231111122212-1331211200030201-2323121012322313-1311010331123231-3220123120210230"></a>

## network_connectors — network_connectors / 213231330031 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- network_connectors

<a id="canonical-2233111321103202-3320001311001223-1302200122123131-1120102313203322-1013031230100002-1200223111020300-0110311031222211-0030300120222110"></a>

Type: `"object"`. list nested block, Optional.

Network Connector defines connection between two virtual networks in a given site. Fleet defines one
or more such network connectors. The network connectors configuration is applied on all sites that
are member of the fleet.

Upstream description:

Network Connector defines connection between two virtual networks in a given site. Fleet defines one
or more such network connectors. The network connectors configuration is applied on all sites that
are member of the fleet.

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

Terraform syntax:

```terraform
network_connectors {
  # Configure direct properties listed below.
}
```

<a id="canonical-1313023130122031-3312301030031312-3332222321130113-2200130313012003-1200031231122213-3102321303022200-1220103211233331-0201120320031031"></a>

## Direct properties — network_connectors / 213231330031 / 3

<a id="canonical-2220233203213002-1010103303011202-0020302231231130-0332332120200001-3030001203330110-2201333303220222-2331012120200023-1333221330230123"></a>

<a id="canonical-2013332233223303-1133322033122230-0233230332312121-0302211330131131-3102033211021031-3120100100230321-2112031102113133-1321013203231022"></a>

## kind property — network_connectors / 213231330031 / 4

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

<a id="canonical-1222331120102022-2213033310302010-1313011233223210-1123013312203001-1003010313200211-1110200101102231-3130101130101030-3300202123100332"></a>

<a id="canonical-1021001333331311-0112000001012131-1022203303331303-3033000110221121-3233231023301021-2301130022211120-1023020302133223-2321321221001102"></a>

## name property — network_connectors / 213231330031 / 5

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

<a id="canonical-0033221133101031-1201020303111012-2030122312223333-2000210102121201-1111111123130132-2301313130220002-2321331110312333-3112011021021220"></a>

<a id="canonical-1033230311323103-0012202210211221-0232230120111031-1232330313123132-2310000322133231-3311000111003213-1001031233221330-0002330322101321"></a>

## namespace property — network_connectors / 213231330031 / 6

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

<a id="canonical-3232100332020232-0223211231000021-0132322330000322-0101123321202332-1302031103303010-0101032032132210-2323231122121121-3021000030311100"></a>

<a id="canonical-3033120331320031-3022001300001312-0132211213203320-3202031103133232-0110132312210112-0222121332130313-3123300122102310-3002113203201131"></a>

## tenant property — network_connectors / 213231330031 / 7

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

<a id="canonical-3102022112233220-1000322113230013-0332312313232121-0103231013001200-1211023212000313-2323232300103020-2111202333330322-0220212330102320"></a>

<a id="canonical-1111101312322110-3211101113023303-2212012103130003-0332212010003332-1231301231303022-1331102012300322-1101212201221211-1100321330323023"></a>

## uid property — network_connectors / 213231330031 / 8

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

<a id="canonical-2322333102201111-3013123122112102-3000203022321002-3100220032131010-3111220332033302-0031330332012212-1003303102323130-3022011303212323"></a>

## Next pages — network_connectors / 213231330031 / 9

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2101130020120112-0133200133323012-0232110022111200-1113311010303213-1313213332310323-1200301201112022-1100302111320120-3222312321231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020301002032222-0023101121102003-2333021321010113-3123132211321132-1231221023100232-0331333201102111-3212111023011220-1233000003023322"></a>

## network_firewall — network_firewall / 013020231131 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- network_firewall

<a id="canonical-2320210312033302-0003021301023012-0231131302212333-2022302003133123-1222032202030023-3301111031131101-1210122331032011-2320112313121321"></a>

Type: `"object"`. list nested block, Optional.

Network Firewall defines firewall to be applied for the virtual networks in the fleet. The network
firewall configuration is applied on all sites that are member of the fleet. Constraints The Network
Firewall is applied on Virtual Networks of type site local network and site local inside network.

Upstream description:

Network Firewall defines firewall to be applied for the virtual networks in the fleet. The network
firewall configuration is applied on all sites that are member of the fleet.

Constraints The Network Firewall is applied on Virtual Networks of type site local network and site
local inside network.

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

Terraform syntax:

```terraform
network_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-1132131030202222-3123303132300033-2003312031113031-2032023003313101-1121001203122202-1213122122002201-3323032030213003-1323122323110031"></a>

## Direct properties — network_firewall / 013020231131 / 3

<a id="canonical-0223030201022003-3120030120002220-2201022022220103-0223220033120311-0121213212010021-3001100203102230-1030200003033131-2321320300330320"></a>

<a id="canonical-3100030313012322-3300020201332230-3323232313101223-0203010100032210-2330031332220020-2303020310232310-3000202330102302-3022222131120031"></a>

## kind property — network_firewall / 013020231131 / 4

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

<a id="canonical-1021011111130311-2302113223123123-3120133233300333-1311320221333012-1312233201330110-3213202312310301-3323320320113112-3230332100202130"></a>

<a id="canonical-3023230320020103-0311321122101310-1321113022013011-0101212223200212-3321213003203013-2123320020130301-0032133230022210-1201323000220202"></a>

## name property — network_firewall / 013020231131 / 5

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

<a id="canonical-3012301122232132-3310212310201133-0023123300211213-0200022233111102-1201113132100133-0122113022013111-3023032100001110-1133030001311331"></a>

<a id="canonical-2202132233231302-1221310030110233-2332022203230232-3313131300202102-0022222202113322-1220211030030110-1120131301233201-3310300111322222"></a>

## namespace property — network_firewall / 013020231131 / 6

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

<a id="canonical-2312010111200312-1013332103323313-1030102302222122-3323230000221131-1203030230001121-0100033232311223-0100320331210120-3111131221222000"></a>

<a id="canonical-1121223113102212-0302202301301200-2111021223210200-0222101130321313-1231120233121031-1010303201100333-0221113102301003-3320312102001002"></a>

## tenant property — network_firewall / 013020231131 / 7

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

<a id="canonical-0000100210323231-3311001011020032-3322011200100021-3320011102311200-3213011023220103-0232113231202020-2230302313332210-1111320230113121"></a>

<a id="canonical-1011212000312211-1002132310021022-0113212323230323-2023001211123222-3211003133021011-0320313332031200-1122113302023300-2133110222320203"></a>

## uid property — network_firewall / 013020231131 / 8

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

<a id="canonical-2032310033110010-0111201331232213-0320020213132322-3131033313002011-1121013131222112-0112203311120320-3203323100130131-1211312230131031"></a>

## Next pages — network_firewall / 013020231131 / 9

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3203320110130321-1121100220230211-2020333200023102-0312320112101333-3102032302000203-1112313033231111-3102003021133322-0232123301133201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330323121311221-0100230000301010-0212230313130113-0020223021133032-2320112302232331-1032203001220312-0020032320122030-3132331110021122"></a>

## no_bond_devices — no_bond_devices / 100210021331 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_bond_devices

<a id="canonical-2233202132112121-1300211023330221-2212133210020112-2021322122322312-0321223001202132-3220130012222211-3210110102122303-1332222011023230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no bond devices.

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
no_bond_devices = {}
```

<a id="canonical-1123021120201301-1031032103012113-0020212223302302-2312001231303321-2003130200222102-3123000302230221-3223213013212111-0131230311212102"></a>

## Direct properties — no_bond_devices / 100210021331 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3330133230311120-2122123221311013-3230302223020300-3323022200102220-1300331233131002-1020211112332133-2002231322010230-2332220202020210"></a>

## Next pages — no_bond_devices / 100210021331 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3002230220210021-3122231112230020-3301210303100111-2000220001202332-3133200133012031-0222312213013001-0021203302213202-2210201320011200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233120113123001-0301010230311303-1333130300313132-2132313032032101-2211203130111033-2331301001102202-3023312132233230-1113133110322020"></a>

## no_dc_cluster_group — no_dc_cluster_group / 230002101112 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_dc_cluster_group

<a id="canonical-0300111311003312-2320102111320110-2333221333222221-2321223000332000-0132220211220113-0221111301300211-3031311131013010-3001123212213133"></a>

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
no_dc_cluster_group = {}
```

<a id="canonical-1331320320120201-1011112210101231-3001113210103332-1122023313032130-0120001123332100-3213120132213010-1203222202202323-1121303313113323"></a>

## Direct properties — no_dc_cluster_group / 230002101112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300312133303231-1021312102311123-1010031300112031-1011013031100102-1320313221100300-0310203313030110-0220013020222301-2123101031312002"></a>

## Next pages — no_dc_cluster_group / 230002101112 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1323333231331210-1331121020203033-1001232023202221-0323221132032011-2122012311233313-0333321001230223-1230201132012120-3320022013311213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332131220111013-1002303131322003-2030333030221120-0210001301030122-1221100013132123-2111231110213310-2101322311000131-3113112313302112"></a>

## no_storage_device — no_storage_device / 023102300221 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_storage_device

<a id="canonical-0220230132023221-0012122010222212-2110320123231220-3000120023300020-3313010120130211-3322003111321031-1011003102230130-0132323201010321"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_device, storage\_device\_list; Default: no\_storage\_device\] Configuration
parameter for no storage device.

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

- [no_storage_device](resources--fleet--reference--group-002.md#canonical-0220230132023221-0012122010222212-2110320123231220-3000120023300020-3313010120130211-3322003111321031-1011003102230130-0132323201010321)
- [storage_device_list](resources--fleet--reference--group-003.md#canonical-1220331113303212-2010311203330203-1123010023313010-2333311011310323-2013213130203030-0232022330203123-2033201030211020-3221200223133131)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_device = {}
```

<a id="canonical-2203203312332300-0211323221001300-2330312210032211-1110320130330303-3320031111330320-0003102222023103-3323132020130100-2301202230101222"></a>

## Direct properties — no_storage_device / 023102300221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112330020210203-3113221313313212-2222233230320023-0133323323221120-2120031013300203-0113023013222110-3103123112120231-3230132213221110"></a>

## Next pages — no_storage_device / 023102300221 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3021101310113300-2330120320311311-0103112023302012-0111200030121332-2122032220110113-0122101022233212-3300032121002220-0102132112203230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213212101032210-1231201032110003-0031212030021321-0023302231312333-1022030111333203-1111101122030003-1301123122013322-0032330130102211"></a>

## no_storage_interfaces — no_storage_interfaces / 203101321110 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_storage_interfaces

<a id="canonical-0132211203012133-3201031322322320-1332302222000203-1001303002202201-0323112222133113-3313003302132131-0010231220301220-2322322220001322"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_interfaces, storage\_interface\_list; Default: no\_storage\_interfaces\]
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

OneOf alternatives in this subsection:

- [no_storage_interfaces](resources--fleet--reference--group-002.md#canonical-0132211203012133-3201031322322320-1332302222000203-1001303002202201-0323112222133113-3313003302132131-0010231220301220-2322322220001322)
- [storage_interface_list](resources--fleet--reference--group-004.md#canonical-2223313213113003-2010202220123320-3322302112013032-3001322211322310-2322330013330331-1010013302213030-2221310300112203-1322231330200313)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_interfaces = {}
```

<a id="canonical-2302301322110033-0111212002113123-0203313010233220-2323023330021320-1331012222033310-0322020131211033-2212012123110202-0030103003110200"></a>

## Direct properties — no_storage_interfaces / 203101321110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311332023211101-2323331122030320-1203211021030110-3121000221102132-1000021133030323-1012332003220223-2321102231100101-1102023211110322"></a>

## Next pages — no_storage_interfaces / 203101321110 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2231011202030311-3303111022101200-2001103313313330-0212102002001302-2030113131220333-1311323330233023-2120212100232203-0031032331001130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312303230303022-3301001323122211-1131323112013011-3022321301120102-0212122021123332-0022102310003032-3111301213003111-2221210201031100"></a>

## no_storage_static_routes — no_storage_static_routes / 013021300011 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- no_storage_static_routes

<a id="canonical-3333003330002323-2020312322133222-3003221302331313-3221312221202102-1100230230310222-2202101300132202-0113222211333133-1021333131121331"></a>

Type: `["object", {}]`. Optional.

\[OneOf: no\_storage\_static\_routes, storage\_static\_routes; Default:
no\_storage\_static\_routes\] Configuration parameter for no storage static routes.

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

- [no_storage_static_routes](resources--fleet--reference--group-002.md#canonical-3333003330002323-2020312322133222-3003221302331313-3221312221202102-1100230230310222-2202101300132202-0113222211333133-1021333131121331)
- [storage_static_routes](resources--fleet--reference--group-004.md#canonical-3300220221311210-2223123102101211-3220330103110130-2013011230213113-0302003112110032-0211100233331131-0000213033313103-2010020102112211)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
no_storage_static_routes = {}
```

<a id="canonical-0112320233130210-2022222301112101-0131031023011302-1323100231201232-1022212023032331-3220101200313313-2311232020123310-2231210003210332"></a>

## Direct properties — no_storage_static_routes / 013021300011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1030333303122023-1003111213312311-1323111301113022-1020000133112130-1111311201012300-0330222031300210-1130100131012212-3221131013003210"></a>

## Next pages — no_storage_static_routes / 013021300011 / 4

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3223331122001220-3011311231202031-2020330101213300-3133211103110113-2113210013003020-2001322230121133-3212301331022221-2100121222312301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202030110010311-0312311300332203-2223030132103022-0133100222022123-0112300030012112-0100110122232011-1103200021112000-1300111021020223"></a>

## outside_virtual_network — outside_virtual_network / 003132200322 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- outside_virtual_network

<a id="canonical-3120120032101112-1320012311032120-3123110212013301-0222000011303300-0020203112032230-3201230233112233-0103200303010013-1233332133233032"></a>

Type: `"object"`. list nested block, Optional.

Default outside (site local) virtual network for the fleet.

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

Terraform syntax:

```terraform
outside_virtual_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131132201013110-1212223001323000-3010122222323113-0321323332332012-2031332330011220-0212301212121333-1012313111200213-3122231213310102"></a>

## Direct properties — outside_virtual_network / 003132200322 / 3

<a id="canonical-0103321323202010-1112100230213231-2013222330100131-3100300130311313-2123110110201013-0013230323113200-1000111231213133-2131021032013111"></a>

<a id="canonical-0000313322023332-1122102213300011-0222110111310222-3202233133221210-3120023233111310-1221012213333323-0131033223303110-1222233131000321"></a>

## kind property — outside_virtual_network / 003132200322 / 4

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

<a id="canonical-3111030303003033-1030111031000020-0202023121210221-0222210232033030-1330123013202110-2023112113130310-0321010213100132-1200310213311033"></a>

<a id="canonical-2231213131201221-0001323100103000-1021212201321012-0313103003312233-2123300032330200-1101000033021021-3213320212211121-1121001003220331"></a>

## name property — outside_virtual_network / 003132200322 / 5

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

<a id="canonical-1122113103211203-2003130010301203-1212211233311010-3002233032130233-0030102201132022-2023022231111300-3001233132033331-3131211233030110"></a>

<a id="canonical-0012103332010223-1031031033331203-2122320313102032-0010110101021002-3203312030110301-3022030331003322-0303202003322202-3332003023220211"></a>

## namespace property — outside_virtual_network / 003132200322 / 6

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

<a id="canonical-0212012002130323-1011132302233203-2020131211303310-2022201103312101-0020133112000330-2300131020232101-0020022303021000-2212310200333010"></a>

<a id="canonical-2201330302001333-3320200303120031-0113033303001212-0120033233320020-3031221132220200-0102230300021231-0310332000311113-1112033323113233"></a>

## tenant property — outside_virtual_network / 003132200322 / 7

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

<a id="canonical-1012003302101233-1223031302033032-3132122001101032-2312012212303001-1310122212210233-1303120200210121-3033311322021301-0330003121323333"></a>

<a id="canonical-1100221001313100-1100121332322303-0301332201130323-0201131223020102-1311222211131323-0323221121322222-1110223203011202-3013010101302113"></a>

## uid property — outside_virtual_network / 003132200322 / 8

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

<a id="canonical-0233011110322021-1233233300030212-0223333202000023-1202121210223020-2200132020131222-3020233031202332-1123012123130310-1321302200101133"></a>

## Next pages — outside_virtual_network / 003132200322 / 9

- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211303102202002-1111131203121220-1200231221230303-3213201011123203-1013120212212212-3300201000011203-2121133312130330-2323112121233313"></a>

## performance_enhancement_mode — performance_enhancement_mode / 102301331112 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- performance_enhancement_mode

<a id="canonical-3030113113100330-0012133311113321-0021320010121312-3022300023333220-0130331330032012-2302300003220202-3021321212323322-3003213102113110"></a>

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

<a id="canonical-0011000222212123-1112301212210222-2302123121103310-1023201033121131-2222020321231210-0033020113130223-0333220331103001-2302201013221012"></a>

## Direct properties — performance_enhancement_mode / 102301331112 / 3

- [perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103): complete subsection reference.

- [perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003): complete subsection reference.

<a id="canonical-0231233113311212-1323102321112001-0212331103321023-0012123002303303-1003002310330110-2320030113003301-1311312222202112-1231230033113222"></a>

## Next pages — performance_enhancement_mode / 102301331112 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3233131011312211-3213022000110301-3103000121302113-0001210032122303-0210123321122121-2302020311313211-1302003100233202-0232301222213301"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced — perf_mode_l3_enhanced / 200122001022 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3203210222120212-1213232111221030-1330301330231323-3320223321203311-1210023122323120-1110002201030313-1323331031332233-0222312122231221"></a>

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

<a id="canonical-0030003203113202-1212330013310312-1300200022321131-2000233210121202-0103323222112212-0212002120231222-3110100023212210-0022130201011330"></a>

## Direct properties — perf_mode_l3_enhanced / 200122001022 / 3

- [jumbo](resources--fleet--reference--group-002.md#canonical-1001231011333031-2123112303131321-1213331333230000-2112333001033020-1012323323213332-3020133020200113-1020202000330121-0023131311333213): complete subsection reference.

- [no_jumbo](resources--fleet--reference--group-002.md#canonical-3230201301023222-1303110022122011-2300010221103110-1200010220323030-0320033011332130-1002301323003230-0230131213100330-2213101312101300): complete subsection reference.

<a id="canonical-0312333023221002-3001211123023120-0303230210222023-2212201322120120-0003232330333022-1003031101331122-2013321322332130-3021310022323333"></a>

## Next pages — perf_mode_l3_enhanced / 200122001022 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--fleet--reference--group-002.md#canonical-1001231011333031-2123112303131321-1213331333230000-2112333001033020-1012323323213332-3020133020200113-1020202000330121-0023131311333213)
- [performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--fleet--reference--group-002.md#canonical-3230201301023222-1303110022122011-2300010221103110-1200010220323030-0320033011332130-1002301323003230-0230131213100330-2213101312101300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1001231011333031-2123112303131321-1213331333230000-2112333001033020-1012323323213332-3020133020200113-1020202000330121-0023131311333213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333313123021312-0230230030202031-1030301222131230-1313111121000233-0202330230000200-3001113010222103-1031110221032230-3122001000022220"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.jumbo — jumbo / 022112122103 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-2321122131222032-1311331312232023-1220132003213010-1313331233223101-2301010022103032-3321120112233322-2030223230311232-2111120030010120"></a>

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

<a id="canonical-2300220203021200-0011301211210000-3021212311121322-3230003111121210-1113022031133122-0132003301230003-0110031132131221-1021110230130113"></a>

## Direct properties — jumbo / 022112122103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121322303311110-2100121313300000-3133200021203103-0301003001010110-1021212213213302-0323312000011221-1103211301313203-2101313101331001"></a>

## Next pages — jumbo / 022112122103 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3230201301023222-1303110022122011-2300010221103110-1200010220323030-0320033011332130-1002301323003230-0230131213100330-2213101312101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302210212102232-0120300121030113-3021313223011220-1322313103311232-2302001310000331-1010000212013101-3023333221223230-0123232202133220"></a>

## performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo — no_jumbo / 123320301311 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-0221131011213303-3333302211113030-0332103013132331-0313211313323002-1312111223120222-1330302333230123-1221022132130011-2131020022022031"></a>

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

<a id="canonical-3023102313310030-2112312330022300-0112232110020230-0102232203111322-1201100330030223-0332232001210230-1132221133131000-1132020121010021"></a>

## Direct properties — no_jumbo / 123320301311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211232311120310-3111132032301320-1131330030031112-3231013303120113-3300212130333300-3303220222023301-3203011113103202-2330010013303312"></a>

## Next pages — no_jumbo / 123320301311 / 4

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--reference--group-002.md#canonical-1220221002013022-3311301112220323-2002223022201102-1023212120230320-0010133301300301-0011232300010130-3202001012220311-1330000011233103)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3301122030312022-2012012100121211-0030220031032200-3011110301210330-3030113330000213-3121202030323222-3021031023101211-2102012312131203"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced — perf_mode_l7_enhanced / 130120313121 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-0102133110100122-1213312131112131-3212313202300102-3010213002211013-1112333132023102-2303102120221133-1312032032313222-1303320333220101"></a>

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

<a id="canonical-1332222102021130-0001010333320031-3221130100321320-2212300323303021-3121200020303301-3123202011030130-1122133023322133-3113122233320320"></a>

## Direct properties — perf_mode_l7_enhanced / 130120313121 / 3

- [jumbo_disabled](resources--fleet--reference--group-002.md#canonical-2013030130220303-1200331321133120-2113203120312213-3222231223320221-2101303323100111-2330021213300330-1111101213210312-0213131131202213): complete subsection reference.

- [jumbo_enabled](resources--fleet--reference--group-002.md#canonical-1302330021030223-0330123321303210-1322121120211302-2100021230020332-0311232102333332-2132001121311132-1133203232213212-0103021132211220): complete subsection reference.

<a id="canonical-3313330210333133-2223031300322131-2213013030121211-3120033123132203-3031003303133201-3302123231320122-2031023320312210-3332323223221203"></a>

## Next pages — perf_mode_l7_enhanced / 130120313121 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--fleet--reference--group-002.md#canonical-2013030130220303-1200331321133120-2113203120312213-3222231223320221-2101303323100111-2330021213300330-1111101213210312-0213131131202213)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--fleet--reference--group-002.md#canonical-1302330021030223-0330123321303210-1322121120211302-2100021230020332-0311232102333332-2132001121311132-1133203232213212-0103021132211220)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2013030130220303-1200331321133120-2113203120312213-3222231223320221-2101303323100111-2330021213300330-1111101213210312-0213131131202213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012022231101110-2201100112203300-3023322002330103-0110012020003211-0033300002221112-1110223013023112-0300010330213230-0213330020221230"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled — jumbo_disabled / 210001230330 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-2323033031003202-0123222000220312-1310020102232010-1321133223313132-3002221002330302-2223310003123311-2211110100030312-3233313331120110"></a>

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

<a id="canonical-2013133012230132-3030131022311110-0030020010012030-0321213300101031-0133003201110301-0013131212320022-3123310233112000-2301330303303020"></a>

## Direct properties — jumbo_disabled / 210001230330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1312233030312032-3303220132330132-2101210200212101-3103023030223032-1220303331303330-2123100212020321-2000211011132333-1100313003002330"></a>

## Next pages — jumbo_disabled / 210001230330 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1302330021030223-0330123321303210-1322121120211302-2100021230020332-0311232102333332-2132001121311132-1133203232213212-0103021132211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203311311302233-1030023232113201-0031032223123133-2320022210221120-2222220031013310-1310121201232032-0312230012103122-2301013220122110"></a>

## performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled — jumbo_enabled / 301010121310 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [performance_enhancement_mode](resources--fleet--reference--group-002.md#canonical-0320032130113210-0100320100303323-1300031032102211-1232123000113130-3303213133311021-0312232322133203-3202110330133301-1330010033303333)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-3100110103313003-2220220000200110-3121111023101023-1311130111113200-1120212313003100-0323210120321110-2000120303323000-2303102232221120"></a>

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

<a id="canonical-3103302120332210-1011112223302131-1332223333312223-3110002111103220-1102202333131110-0112102001320002-3022333002101130-3103033130200232"></a>

## Direct properties — jumbo_enabled / 301010121310 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310220331203300-0311212032302331-0100330123330121-2103011030022211-0233232032211013-1210033103313201-1323130230020210-1213331001220330"></a>

## Next pages — jumbo_enabled / 301010121310 / 4

- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--reference--group-002.md#canonical-1012131121111231-1231301233030320-1112003213120200-1322321203310001-2102200321103013-2323321001322312-2302030110001110-2313133030231003)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100332211202032-1312200001220202-0013001013130002-2332323320112321-2132003130010303-2212110000132201-0201020130232213-3300233221210302"></a>

## sriov_interfaces — sriov_interfaces / 001033113233 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- sriov_interfaces

<a id="canonical-0122221131232302-1301033322113301-0110223212011211-3233102320110033-0121020202101031-3330230010133021-2322320310031011-0300013133213110"></a>

Type: `"object"`. single nested block, Optional.

List of all custom SR-IOV interfaces configuration.

Receipt-pinned upstream constraints:

```json
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
sriov_interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-3022232330103113-2220231210332011-1122322202013222-1011222200011033-1203301030001301-1111223303333233-2200022131333001-1002233222311033"></a>

## Direct properties — sriov_interfaces / 001033113233 / 3

- [sriov_interface](resources--fleet--reference--group-002.md#canonical-2200113032031301-2311202302023112-3021303002001132-0201031103201130-3203131323221301-0011022313130111-2212302233122122-2122203210130320): complete subsection reference.

<a id="canonical-0022211200213222-1001301002300222-1112203201220211-1211020011120122-1331320211231031-3213001031122022-1331231221202121-2123200230330332"></a>

## Next pages — sriov_interfaces / 001033113233 / 4

- [sriov_interfaces.sriov_interface](resources--fleet--reference--group-002.md#canonical-2200113032031301-2311202302023112-3021303002001132-0201031103201130-3203131323221301-0011022313130111-2212302233122122-2122203210130320)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2200113032031301-2311202302023112-3021303002001132-0201031103201130-3203131323221301-0011022313130111-2212302233122122-2122203210130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123311220102321-3203312033123313-1000222011120111-0333113010211202-0220113322022220-3230010000121211-3133111210302230-0031130232120202"></a>

## sriov_interfaces.sriov_interface — sriov_interface / 303310333200 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021)
- sriov_interfaces.sriov_interface

<a id="canonical-0111313130111223-0101332233300300-2203102011300122-3303001100021221-1201313331221303-2320100023133330-2300300123213223-0013031000311111"></a>

Type: `"object"`. list nested block, Optional.

Use custom SR-IOV interfaces Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interface_name",
    "number_of_vfs")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
sriov_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-2231001013230310-3001221131011102-3333011132013100-3322103103232332-0003233312311301-0202131211300003-1212130012322303-1330232213122121"></a>

## Direct properties — sriov_interface / 303310333200 / 3

<a id="canonical-2123331310213120-2210230032010112-2320220102201000-1013122020123333-2133301212021301-3302300131120120-2333320113133233-2011003233203113"></a>

<a id="canonical-0132102321323323-3021211212012131-3323111300300203-1213111311222330-2203223320022001-3031303023200120-0133230121212000-0023131032110203"></a>

## interface_name property — sriov_interface / 303310333200 / 4

Type: `"string"`. Optional.

Name of physical interface. Name of SR-IOV physical interface.

Upstream description:

Name of SR-IOV physical interface.

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

<a id="canonical-2023133122231310-0113222003213021-0313132300113010-3222133301301012-1123330223030202-1202220130201302-2123131332333001-3213231000110110"></a>

<a id="canonical-0003020011002131-0301232123213020-1312101332033310-3220033311221212-2202331312033112-2101023301100320-3201323111231210-0122011212101020"></a>

## number_of_vfio_vfs property — sriov_interface / 303310333200 / 5

Type: `"number"`. Optional.

Number of virtual functions reserved for VNFs and DPDK-based CNFs.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1233020321322311-2312010020020301-2120132130223210-2312002232001100-2331221210303221-3130003303230112-3330131212031312-1023031200111211"></a>

<a id="canonical-3320322213232330-3302331332232210-1230312021233221-2110033031333321-1101202203100203-1030110112032301-0201003231221001-1320312210301230"></a>

## number_of_vfs property — sriov_interface / 303310333200 / 6

Type: `"number"`. Optional.

Total number of virtual functions. Total number of virtual functions.

Upstream description:

Total number of virtual functions.

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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-2333032120220112-1330030222333312-0131033123102332-2323120332000021-1300213013203121-3300122330033123-2200333212100221-0021232030013330"></a>

## Next pages — sriov_interface / 303310333200 / 7

- [sriov_interfaces](resources--fleet--reference--group-002.md#canonical-0010012300131213-3030202121233221-2000103301112222-1032223132021032-0012001001222201-2110010232302022-1100022133202103-0222030012212021)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022010331230331-2123100211020202-1211010210211212-2223331202110222-0222030203321133-0300322012230030-1321121003030012-2330003002203323"></a>

## storage_class_list — storage_class_list / 113321001301 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- storage_class_list

<a id="canonical-2222300023023313-2122302313220111-2233300300012313-3303221100131111-0112033102321132-1321110322120221-0103222322331100-1100211021322120"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
storage_class_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1312201223013303-3331121111301023-1102330300121221-0001011013220033-1121322203300310-2021032101321020-3331001330031220-2313011102001231"></a>

## Direct properties — storage_class_list / 113321001301 / 3

- [storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012): complete subsection reference.

<a id="canonical-2210330031002330-3211333023331301-0223131230211310-2103302323133110-1120133213033311-1220230032221230-0121203123002022-0002230313233122"></a>

## Next pages — storage_class_list / 113321001301 / 4

- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103323201130130-0211022002333311-1310131321300112-2332323300232102-1221311131102211-2002310331322231-0133011122121022-2000031320201131"></a>

## storage_class_list.storage_classes — storage_classes / 210033210020 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- storage_class_list.storage_classes

<a id="canonical-1330213321222123-1121113101212033-0130331233330120-2323211313312210-3212020022203233-2102131212001131-0330220113011110-2023323211312113"></a>

Type: `"object"`. list nested block, Optional.

List of Storage Classes. List of custom storage classes.

Upstream description:

List of custom storage classes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("storage_class_name",
    "storage_device"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "hpe_storage"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("custom_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "netapp_trident"),
  validators.ConflictingListObjectAttributes("hpe_storage",
    "pure_service_orchestrator"),
  validators.ConflictingListObjectAttributes("netapp_trident",
    "pure_service_orchestrator")}
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

Terraform syntax:

```terraform
storage_classes {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213312020220000-0230313032031120-2330301102213112-1201320311313213-0333222222300221-0113002101330231-2120011120032312-3010201320103100"></a>

## Direct properties — storage_classes / 210033210020 / 3

<a id="canonical-2221030123300233-0110230120011332-1303202011130000-2020112230013002-3221220130102100-0102221332011012-2113210020330311-1113012001022321"></a>

<a id="canonical-0132112130003130-3121030010313012-1131221012321310-3001311111012203-0030332320121221-3110332331012110-0021112320310132-1313020031010330"></a>

## advanced_storage_parameters property — storage_classes / 210033210020 / 4

Type: `["map", "string"]`. Optional.

Advanced Parameters. Map of parameter name and string value.

Upstream description:

Map of parameter name and string value.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.max_len\":\"128\",\"ves.io.schema.rules.map.values.string.min_len\":\"1\"},\"values\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"}}")}
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
      "ves.io.schema.rules.map.values.string.max_len": "128",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 128,
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
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "128",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.max_pairs": "64",
    "ves.io.schema.rules.map.values.string.max_len": "128",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-2202133111232311-1231331012012123-1121023221003231-2102332300032331-1132210010111111-3333320013013012-1330210010031012-0023121003112131"></a>

<a id="canonical-3101132223223131-2103031001020311-2330200123133302-1222012213003313-0131313011302000-3221230220323233-3013000022033111-3220323130112312"></a>

## allow_volume_expansion property — storage_classes / 210033210020 / 5

Type: `"bool"`. Optional.

Allow Volume Expansion. Allow volume expansion.

Upstream description:

Allow volume expansion.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [custom_storage](resources--fleet--reference--group-002.md#canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221): complete subsection reference.

<a id="canonical-0310131120333103-3131013023012023-3332321000031320-0230220103031110-1122331331303233-2320211131001321-3302202012303030-3111003212101112"></a>

<a id="canonical-3202122012332233-0302012020023031-3033301011233102-0300211333021300-3100022201310230-3122233312302302-0100132223300000-2313112211233101"></a>

## default_storage_class property — storage_classes / 210033210020 / 6

Type: `"bool"`. Optional.

Make this storage class default storage class for the K8s cluster.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-1120201231001110-3312300231013311-1103123310101213-3023112130023312-1033000201010220-2131110100311320-3233213220021030-1320012310312203"></a>

<a id="canonical-2320133203030313-0320202222333122-3133120130030310-1331313133113020-3333210210331121-1232230130212200-0132312222231203-2013221302200013"></a>

## description_spec property — storage_classes / 210033210020 / 7

Type: `"string"`. Optional.

Storage Class Description. Description for this storage class.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [hpe_storage](resources--fleet--reference--group-002.md#canonical-2032302011222201-1310233322022302-0123100203223233-1210011333013002-2322333333201313-1212222011311012-2021103000021222-3002302213031013): complete subsection reference.

- [netapp_trident](resources--fleet--reference--group-002.md#canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102): complete subsection reference.

- [pure_service_orchestrator](resources--fleet--reference--group-002.md#canonical-1110101010012033-3333011200130233-2120103112121132-1312112202313113-2102222332321202-2231022203130123-1000222321310221-2201120133301120): complete subsection reference.

<a id="canonical-1221330002110232-1330210010213303-3122131210210101-2323203131323133-2210223112321131-3302203033113122-0303333011323223-1330111211100203"></a>

<a id="canonical-3001030232220130-0000002311123020-1010332111222300-3103321121133202-2332013233010313-2103011120233122-1222102113213021-2132033103302223"></a>

## reclaim_policy property — storage_classes / 210033210020 / 8

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

Reclaim Policy.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 16,
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
    "ves.io.schema.rules.string.max_len": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "16"
  }
}
```

<a id="canonical-1122312030110112-2123300032021100-1302131230102013-1010202122131123-3030221101211202-2200120003120302-2321320302132103-1322300000230231"></a>

<a id="canonical-1230300320332032-3321101011201112-1022032113022213-1312330013232002-1133031001000320-2102110331221301-1130031112002211-0330310303101323"></a>

## storage_class_name property — storage_classes / 210033210020 / 9

Type: `"string"`. Optional.

Name of the storage class as it will appear in K8s.

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
    "format": "dns-label",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-1011230231022323-1012133323300310-3301103320223322-0031203100122123-2023112220210130-0010032230233322-1221322301321213-0131000113332123"></a>

<a id="canonical-3021221231020203-2011001233110332-2210003010110220-1210200222301301-3230013313102232-1002113210023203-3232322232033122-0131002100122311"></a>

## storage_device property — storage_classes / 210033210020 / 10

Type: `"string"`. Optional.

Storage device that this class will use. The Device name defined at previous step.

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
    "byteLength": {
      "max": 64,
      "min": 1
    },
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
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0221101210322000-0133132203003300-1110020323313232-1111211101123303-0012130333133120-2320002110110312-2213230231213211-3201201332201132"></a>

## Next pages — storage_classes / 210033210020 / 11

- [storage_class_list.storage_classes.custom_storage](resources--fleet--reference--group-002.md#canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221)
- [storage_class_list.storage_classes.hpe_storage](resources--fleet--reference--group-002.md#canonical-2032302011222201-1310233322022302-0123100203223233-1210011333013002-2322333333201313-1212222011311012-2021103000021222-3002302213031013)
- [storage_class_list.storage_classes.netapp_trident](resources--fleet--reference--group-002.md#canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102)
- [storage_class_list.storage_classes.pure_service_orchestrator](resources--fleet--reference--group-002.md#canonical-1110101010012033-3333011200130233-2120103112121132-1312112202313113-2102222332321202-2231022203130123-1000222321310221-2201120133301120)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-3003121013331300-1300220222103221-2021323122133331-0011011023003132-0213101122200221-0222333220013210-1333213000113113-3202103202223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001120233323203-3032011232032213-3003303022121333-2132311023033001-2320132123213322-3100130321301322-0011000220231310-3300001330333002"></a>

## storage_class_list.storage_classes.custom_storage — custom_storage / 030211212002 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.custom_storage

<a id="canonical-3033311331231221-1311001010132332-0211110001021002-2302112231023110-2122322120033213-0223103231220001-0223003330203112-3310232133231010"></a>

Type: `"object"`. single nested block, Optional.

Custom Storage Class allows to insert Kubernetes storageclass definition which will be applied into
given site.

Receipt-pinned upstream constraints:

```json
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
custom_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-2210223121032031-1200332330330023-0232303331111213-1213010121131130-2010011310300303-0202012033102311-2003032000220003-1313011031011223"></a>

## Direct properties — custom_storage / 030211212002 / 3

<a id="canonical-0012223213023021-2022231023300113-2210222233333302-2300202213333101-1102133000202220-3303213120332332-3021112012210223-1220301023303211"></a>

<a id="canonical-0333201232302321-1212301101112003-3111322130120113-3130201001131330-2320320200232220-1110323333210131-0022101231310333-1331332013321112"></a>

## YAML property — custom_storage / 030211212002 / 4

Type: `"string"`. Optional.

Storage Class YAML. K8s YAML for StorageClass.

Upstream description:

K8s YAML for StorageClass.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 4096),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 4096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "Valid parseable YAML",
    "maxLength": 4096,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "validation": {
      "customRule": "Must be valid YAML"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "4096",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0033122331123130-1121200301123011-2222320030303300-1321111200113320-3231330133130332-0113301002102003-2100103010100212-3033203122331130"></a>

## Next pages — custom_storage / 030211212002 / 5

- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2032302011222201-1310233322022302-0123100203223233-1210011333013002-2322333333201313-1212222011311012-2021103000021222-3002302213031013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0220131123002020-3021001220222310-2003021100000000-0023131201303131-2213233013331223-1020303231020002-2323213100301232-1030220211332220"></a>

## storage_class_list.storage_classes.hpe_storage — hpe_storage / 211103123112 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.hpe_storage

<a id="canonical-0030320200321221-1122032232023202-1010033100131301-0303232232220311-0120120111310233-1311333110233221-1032103003203013-3313223100132023"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for HPE Storage.

Receipt-pinned upstream constraints:

```json
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
hpe_storage {
  # Configure direct properties listed below.
}
```

<a id="canonical-0001300203223231-3320201331203212-3032100013123222-3231031321320321-0220323330313320-3013001312133111-0132303110100331-0201322132310210"></a>

## Direct properties — hpe_storage / 211103123112 / 3

<a id="canonical-3021222303033311-0032023011212013-0131011032311111-1012211033111220-0121110212320220-3331320321311220-2111013133113221-2233032231230101"></a>

<a id="canonical-3233332101021312-2322300320103102-1123302313100123-0221020300333313-0101232220121002-3222100022330303-3110331103101223-1023232113101003"></a>

## allow_mutations property — hpe_storage / 211103123112 / 4

Type: `"string"`. Optional.

Mutation can override specified parameters.

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

<a id="canonical-1103003000332122-1130002230033202-0200130110002010-2133123131130111-3102010010002133-1233332023123312-3010132233003312-0323210003333100"></a>

<a id="canonical-3302311333221120-0122001220301111-2032122130311230-3301011321120130-2321331202202102-1302123030312103-2122031222112320-0220201232013120"></a>

## allow_overrides property — hpe_storage / 211103123112 / 5

Type: `"string"`. Optional.

AllowOverrides. PVC can override specified parameters.

Upstream description:

PVC can override specified parameters.

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

<a id="canonical-0330213203012130-2131032223020120-0232111132100010-1013333302020001-0200300013211313-0200300133302132-1013002211223202-1032222331131232"></a>

<a id="canonical-0010120233111333-2202113131012312-2222012200022321-3223102313332301-1231022212121330-1211100120321221-0231230102113211-2010122121131232"></a>

## dedupe_enabled property — hpe_storage / 211103123112 / 6

Type: `"bool"`. Optional.

Indicates that the volume should enable deduplication.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0212323122033303-1030320320313302-0131212130300330-3223123002300120-3123001121303220-0123010213302023-3013222132213321-3020203200131001"></a>

<a id="canonical-2330030220323311-1022302030000312-2303302130010221-2100331103013103-1010311222222230-0231201131030302-1220010020303132-3010123223002112"></a>

## description_spec property — hpe_storage / 211103123112 / 7

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

<a id="canonical-2032311201210020-2101301322313320-0330103020121311-1311021021023000-1011123330312332-0302022323023311-2033032323003223-1020310020313203"></a>

<a id="canonical-0001213011102332-2123232133020333-1103221300301323-0330010023221213-3121001001003320-0122202000103302-2310310102232112-2231001123310213"></a>

## destroy_on_delete property — hpe_storage / 211103123112 / 8

Type: `"bool"`. Optional.

Indicates the backing Nimble volume (including snapshots) should be destroyed when the PVC is
deleted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3213233032032202-3222120021103130-1020000100120213-0203032012220101-0023033333301311-0110102123120122-1133210300102221-3132231231121312"></a>

<a id="canonical-1102210303220320-3010131221203323-0312123222333120-0330311312031101-0032330212213123-0320222221233331-1112313211013332-3101131202103101"></a>

## encrypted property — hpe_storage / 211103123112 / 9

Type: `"bool"`. Optional.

Indicates that the volume should be encrypted.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0100223302023313-1230212121013111-3020013301330222-3000211003311213-3233111132103221-2310101332022212-1000112022203000-1012321220111313"></a>

<a id="canonical-0110312203312222-1323100302233202-1131220230103121-0313113333231022-3332220222100203-1321330220210230-0103331202112003-1121311201022223"></a>

## folder property — hpe_storage / 211103123112 / 10

Type: `"string"`. Optional.

The name of the folder in which to place the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0331312333303230-3213032021203001-3303331222003121-1211221103231122-0003210111232303-1203323130103202-2331131220233332-2020122331222211"></a>

<a id="canonical-2312230302023311-0200232311202331-2202120230031311-2103210333131200-2311220123010233-1202320101323123-3011302130303001-1012111212332130"></a>

## limit_iops property — hpe_storage / 211103123112 / 11

Type: `"string"`. Optional.

LimitIops. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-0000202201101113-2000113021101311-1312103120020030-1221112332321330-1020013122032001-2223202120311011-0101113002111033-3331022012310011"></a>

<a id="canonical-2030100003013320-2222100223010332-0211013010133313-2233020200221112-3330203232201100-2312210033022211-1312021132110312-0123023032111233"></a>

## limit_mbps property — hpe_storage / 211103123112 / 12

Type: `"string"`. Optional.

LimitMbps. The IOPS limit of the volume.

Upstream description:

The IOPS limit of the volume.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "int64",
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

<a id="canonical-2312032311313213-2103301311122211-3003002223320220-2113321111222123-0110023132211331-2123022212003132-0232101001233332-2113202313030022"></a>

<a id="canonical-3202330030331300-3023202212230323-2203231112311330-2233313132323212-1113112223302211-0201002122112112-1320222202101212-0130330032023220"></a>

## performance_policy property — hpe_storage / 211103123112 / 13

Type: `"string"`. Optional.

Policy configuration for this feature.

Upstream description:

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1123120313000032-3010213213201310-0201310123223002-2130203001013203-2113200321303231-1001020332010112-0232220012000023-3323000002030032"></a>

<a id="canonical-2002312313111113-1021312200211210-1032030021230022-2313322030310030-2200312303102031-1110100013030032-0231322222200011-3030333311311232"></a>

## pool property — hpe_storage / 211103123112 / 14

Type: `"string"`. Optional.

The name of the pool in which to place the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1130323311201032-1120010203311003-1323331001033313-2201120330021223-0312121332110113-0001010113003223-3102232331010231-3123121123013300"></a>

<a id="canonical-3111132023103211-2322310103031012-1333030211120013-2333221313201312-2022321312031231-3000330331331022-3132012102210031-3111010233202330"></a>

## protection_template property — hpe_storage / 211103123112 / 15

Type: `"string"`. Optional.

The name of the performance policy to assign to the volume.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-0303123221003200-3100113303333222-3112311000012323-0222213103100310-1201100100133021-1001010210221121-1002012323101012-0310110101333333"></a>

<a id="canonical-2000230323101323-0110201301312111-0112210300020001-0111333203300331-2102022031000111-1220022100101212-2012233113233332-3331132210121322"></a>

## secret_name property — hpe_storage / 211103123112 / 16

Type: `"string"`. Optional.

The SecretName parameter is used to identify name of secret to identify backend storage's auth
information.

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

<a id="canonical-2021201030233111-0002201221112103-2101222223120030-1220313023023210-0200030331033230-0203023022310311-2112201211212102-0020331232210010"></a>

<a id="canonical-2312303133330121-1021333000302003-3312111022331002-2300022313023300-3023021230102123-1122010130333311-2000001320311233-3010323031312213"></a>

## secret_namespace property — hpe_storage / 211103123112 / 17

Type: `"string"`. Optional.

The SecretNamespace parameter is used to identify name of namespace where secret resides.

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

<a id="canonical-1221100221101202-0002132220221330-0232110121112122-3212132202120000-2321312103122110-3210013200122132-0323112100110030-2020031100302300"></a>

<a id="canonical-3021200130330320-2001102200030210-0000213120310013-1333211022121233-0020222313030013-3131330312111123-1320110332033211-0120123323133033"></a>

## sync_on_detach property — hpe_storage / 211103123112 / 18

Type: `"bool"`. Optional.

Indicates that a snapshot of the volume should be synced to the replication partner each time it is
detached from a node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2103130213311223-0331203131333031-2012120002033221-1020313133213322-1033323230020123-1001123233122311-0010201322001133-1111122310101123"></a>

<a id="canonical-2032101213202301-3011003312010222-1213111122232202-3012003231233200-3121123012023102-2201321020221233-0122331322220110-0201133120020132"></a>

## thick property — hpe_storage / 211103123112 / 19

Type: `"bool"`. Optional.

Indicates that the volume should be thick provisioned.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3122321220312211-3012202332333112-1012203310123301-3023110313111203-3301202012121302-1231330331230322-0130220302321221-0333312310002100"></a>

## Next pages — hpe_storage / 211103123112 / 20

- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232023211320023-2032232311032212-3011200021232330-3111210133333103-0222131330330230-0022331102321131-3123012221113220-1330003112221130"></a>

## storage_class_list.storage_classes.netapp_trident — netapp_trident / 332101100011 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- storage_class_list.storage_classes.netapp_trident

<a id="canonical-3211323111122200-2022231133310222-2111133212330300-1010113231101330-0021303132000002-3100101322300112-2032232202110332-3231102032113310"></a>

Type: `"object"`. single nested block, Optional.

Storage class Device configuration for NetApp Trident.

Receipt-pinned upstream constraints:

```json
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
netapp_trident {
  # Configure direct properties listed below.
}
```

<a id="canonical-2323020330323323-3311302123331223-2030203011131121-3322223311012102-3303130300221231-2012232310130222-0213013300101020-0001323131212321"></a>

## Direct properties — netapp_trident / 332101100011 / 3

- [selector](resources--fleet--reference--group-002.md#canonical-1231211312010313-3221012302202022-1123112203201012-0332131023302331-2312131323222331-3131200232222010-0210211001112201-2201230101112003): complete subsection reference.

<a id="canonical-1323200222023300-3202123201121232-3200302013333131-1202122303320021-0313200002321110-1313123332203113-1103122333132232-1200231333320033"></a>

<a id="canonical-1011123313313013-2213302322133121-1022121012020222-2310001302003322-1121331203213003-3100323020010111-2020213300320211-1022032030122112"></a>

## storage_pools property — netapp_trident / 332101100011 / 4

Type: `"string"`. Optional.

The storagePools parameter is used to further restrict the set of pools that match any specified
attributes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
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
    "ves.io.schema.rules.string.max_len": "512"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "512"
  }
}
```

<a id="canonical-1133022303322030-3213201210220220-1211000221313032-1202233111121221-0332223321022103-2232301212021013-1213110113311110-1122100223321030"></a>

## Next pages — netapp_trident / 332101100011 / 5

- [storage_class_list.storage_classes.netapp_trident.selector](resources--fleet--reference--group-002.md#canonical-1231211312010313-3221012302202022-1123112203201012-0332131023302331-2312131323222331-3131200232222010-0210211001112201-2201230101112003)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1231211312010313-3221012302202022-1123112203201012-0332131023302331-2312131323222331-3131200232222010-0210211001112201-2201230101112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301201330301220-2203133213032230-0001312121310113-1223223332132101-3020330112132022-2332012322110011-3210201203330021-3300122020121011"></a>

## storage_class_list.storage_classes.netapp_trident.selector — selector / 310123131101 / 2

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)
- [Property reference](resources--fleet--reference--group-001.md#canonical-3312201012220032-2032322221221203-2303110033300201-0023202031100302-0322303132112120-1000212311220020-2212322120022212-1100233331203300)
- [storage_class_list](resources--fleet--reference--group-002.md#canonical-1011213213131132-0031210023033030-3201313200022303-1022122313230202-0210113023121010-0230320300003311-2023233020031113-2133031023233023)
- [storage_class_list.storage_classes](resources--fleet--reference--group-002.md#canonical-3020330102001101-3202000123022113-2313200323220022-0020122022311010-3210010202222310-3202201033010302-3033012303001033-2322132100001012)
- [storage_class_list.storage_classes.netapp_trident](resources--fleet--reference--group-002.md#canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102)
- storage_class_list.storage_classes.netapp_trident.selector

<a id="canonical-1230001331232023-1121200110230223-1320210312010322-2222111303221202-1030032033322312-1213022300230231-0123203120230022-3133132202222111"></a>

Type: `"object"`. single nested block, Optional.

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Upstream description:

Using the Selector field, each StorageClass calls out which virtual pool(s) may be used to host a
volume. The volume will have the aspects defined in the chosen virtual pool.

Receipt-pinned upstream constraints:

```json
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
selector {}
```

<a id="canonical-0113012333001123-0002212223031312-0131102330032121-0220213020212223-3230301133132013-3323222023013223-2120130320031011-1112020123300123"></a>

## Direct properties — selector / 310123131101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3031202321133033-1300131232002220-1231033201110001-3013020010032023-0212310011120001-0100103000320220-2111113102103312-3212203003133000"></a>

## Next pages — selector / 310123131101 / 4

- [storage_class_list.storage_classes.netapp_trident](resources--fleet--reference--group-002.md#canonical-2010330031210131-3132001030023300-0203201120221132-2132020001220232-3210102120202312-1202112101231101-1130010221231010-3121031322310102)
- [xcsh_fleet](../resources/fleet.md#canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131)

<a id="canonical-1110101010012033-3333011200130233-2120103112121132-1312112202313113-2102222332321202-2231022203130123-1000222321310221-2201120133301120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
