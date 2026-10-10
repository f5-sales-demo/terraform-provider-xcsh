---
page_title: "xcsh_bigip_http_proxy reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bigip_http_proxy reference."
---

# xcsh_bigip_http_proxy reference

<a id="canonical-2203113031322030-2330022102103202-0211000330230112-2321023021232000-2132221020212300-2323003202003223-1033331032231323-1223000222120121"></a>

## Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool`

- [no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-1120300203220233-1323222220312211-0331031300103232-1032020031220021-0122233111303312-2300212310302333-1111201311313111-3122313000013000): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0032212100033122-2311302103202323-1103111310103113-0223312130120302-3012020303232232-2120331002202032-0233210232332320-1011300123211203): complete subsection reference.

<a id="canonical-1120300203220233-1323222220312211-0331031300103232-1032020031220021-0122233111303312-2300212310302333-1111201311313111-3122313000013000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-1311211321313201-3321333332130333-1111232100201310-0132003101311310-2303103123122102-2211202330330320-0120121003102110-0330330133201010)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-2123321133201330-1101022110010022-2212000130131300-3133310211302131-0030133321202301-2022230213131132-0120023221132230-3102120020000311"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032212100033122-2311302103202323-1103111310103113-0223312130120302-3012020303232232-2120331002202032-0233210232332320-1011300123211203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool](resources--bigip_http_proxy--reference--group-001.md#canonical-1311211321313201-3321333332130333-1111232100201310-0132003101311310-2303103123122102-2211202330330320-0120121003102110-0330330133201010)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-3211300102000202-3203101313113331-0331003310300202-3303003031202201-2003013322303020-2233320301203211-1122013313021221-2101220030033123"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210112103230300-3300302311102230-0200121321020310-2221213200122320-3302200013031032-0231101112023233-0210130133332012-1020000003110200"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-3302231032122000-1312303303312013-0101003333130022-2030011323001031-2022302202032303-2323120111000321-0323332302211130-1223331130300021"></a>

#### `origin_pools.pools.origin_servers.origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0011302332230132-2030011302213232-3301331313023332-1231222203031121-0301012321311130-2330300102211320-1332010123131110-3020003033110321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.k8s_service](resources--bigip_http_proxy--reference--group-001.md#canonical-0211202201221121-3001000133010112-1131111200013313-3122130332300332-1020212223330013-1311100313303320-2312112110031023-1120330330033233)
- origin_pools.pools.origin_servers.origin_servers.k8s_service.vk8s_networks

<a id="canonical-0211133031012200-2011221102322000-1112030331321300-3123012032022312-3011231122213110-2012121212030310-1210231103212022-2103221310302031"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for vk8s networks.

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
vk8s_networks = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- origin_pools.pools.origin_servers.origin_servers.private_ip

<a id="canonical-2321223011000313-2312122221131022-3113333013312112-0230003332100121-2120322130022220-0332332310100321-0133011133031000-2222010013223133"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
private_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3002302310330020-1321133200100220-0223212320311301-3203030101210112-2321202002202221-3333033033321302-3113032011001133-3321113331313001"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip`

- [inside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-3202220220231220-0201102032212232-0103210303131122-3203331020231232-1121131021103133-2030011130021023-1303022033123302-3302033220313302): complete subsection reference.

<a id="canonical-1022233231121011-3310131031220322-1311222313111030-0001211303112323-1101023111231102-0112003203110311-1130131202330210-2022203221012200"></a>

<a id="canonical-3222301001300332-2100211322112112-2312221110110333-1223031030102312-2313013223033200-1110211002322103-2303332002020232-3213301222220223"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.ip` property

Type: `"string"`. Optional.

IP. Exclusive with \[\] Private IPv4 address.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [outside_network](resources--bigip_http_proxy--reference--group-002.md#canonical-1302012031211001-1312222023022210-0320221000003321-3333101311320103-1021212313101230-2021001122100133-2100212010313301-3012021021132313): complete subsection reference.

- [segment](resources--bigip_http_proxy--reference--group-002.md#canonical-2123331231213133-1211031302232313-0220220030322111-0323303301330321-2232332221133103-3213112033233332-2113101103232110-3102311332103030): complete subsection reference.

- [site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-1230023122312032-1103223311003112-2213211131230113-1201222200312112-0222310213120110-1103130022310122-2131113221210111-2032330032013313): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0012323013103221-2113331011300123-0022012222333113-1101000223003033-2320202111220110-1232200202020301-2020120203313201-0021111002012011): complete subsection reference.

<a id="canonical-3202220220231220-0201102032212232-0103210303131122-3203331020231232-1121131021103133-2030011130021023-1303022033123302-3302033220313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- origin_pools.pools.origin_servers.origin_servers.private_ip.inside_network

<a id="canonical-0233101222100231-3302112213020231-3322003231333122-3130120112112211-2230323102202120-3213030221121330-2131211200131013-0002023100010101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for inside network.

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
inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302012031211001-1312222023022210-0320221000003321-3333101311320103-1021212313101230-2021001122100133-2100212010313301-3012021021132313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- origin_pools.pools.origin_servers.origin_servers.private_ip.outside_network

<a id="canonical-2102223102132210-2203301302312120-2212022310023003-2312013330121102-2312123012322021-3103312132133000-1211223113103113-0030101333222020"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for outside network.

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
outside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123331231213133-1211031302232313-0220220030322111-0323303301330321-2232332221133103-3213112033233332-2113101103232110-3102311332103030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.segment` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- origin_pools.pools.origin_servers.origin_servers.private_ip.segment

<a id="canonical-2331312210111131-2211020012321111-0003122010132102-3102030130113302-3301000131030110-1021133223220031-1112101220302020-0002021300220101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
segment {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021012200023301-0211331223320211-2322201223102223-2103233320303112-2330012020200123-2203112011011033-3223330323103223-3103201322023230"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.segment`

<a id="canonical-0120022000012333-1102013223103310-0233133033123210-2321201113332213-0012320113030213-0221020111122232-0112320323012311-2123032323323113"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3123003230221133-3313330330000212-3022210002302111-3101100210330222-1323031213202130-3201220102030122-1031021323123122-1303303333031130"></a>

<a id="canonical-2303320030132233-2132310313313021-1300003113010100-1332132110120222-2032011000322000-2301221301221202-1311110223210332-3100002221331113"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3000130331013221-0230003332023333-2311121333302030-0010302012200233-3102132302313023-0222303002312201-0132003203212321-0011131231312330"></a>

<a id="canonical-0303312012000303-2111020331323121-0011102013120103-0022331003001202-1023201323020130-2101032100113220-1131221110021020-2111011213121300"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.segment.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1230023122312032-1103223311003112-2213211131230113-1201222200312112-0222310213120110-1103130022310122-2131113221210111-2032330032013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator

<a id="canonical-0233221201203020-0011210022002003-3322230223201030-3133102303102002-2312110201100122-1010223331300130-1131230323330321-1012213232120130"></a>

Type: `"object"`. single nested block, Optional.

This message defines a reference to a site or virtual site object.

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
site_locator {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203203332021003-3120122321010332-3223321112221231-3100201131001023-3030032321203123-1001020202211213-1212212010001203-3200110132110110"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator`

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-2302102311011122-3312320203201103-3303223102103032-0313211321301323-0230112133301100-2330213111312321-2211031022123200-1233111301113213): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-002.md#canonical-3211220120203012-1121320222113230-0310020112232031-0012131223111031-0312113003230111-1133020320200002-3121113321223103-2012022113333123): complete subsection reference.

<a id="canonical-2302102311011122-3312320203201103-3303223102103032-0313211321301323-0230112133301100-2330213111312321-2211031022123200-1233111301113213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-1230023122312032-1103223311003112-2213211131230113-1201222200312112-0222310213120110-1103130022310122-2131113221210111-2032330032013313)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site

<a id="canonical-1232121222333030-1231220133301011-2002001330312132-3030101003132221-2310333303012230-0322233132212313-0213222210033212-3322311200231121"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110220130121203-2030033122003331-1213120020003301-2200030022202032-0020332220231300-1230011213133013-3130222323012322-3002221033130322"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site`

<a id="canonical-3122330301332313-3320001113231021-1222231131023330-0201312233003003-3010123320322232-2113303033001022-2000122302003111-0321031133331001"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3212303023220122-1312212030012230-2021301001232021-3020210300201001-0022211333101013-2220113013333212-3112320010101031-3131220330302233"></a>

<a id="canonical-0221013321323032-0130212201002110-3230232310202311-1322012321023211-3103121332110032-3223121102110003-0211233010031100-1201000200112233"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3201330000011032-2013003301013110-0202221210132320-1301203330202221-1202003030230210-2322323111002311-2003012103303103-1130032130031101"></a>

<a id="canonical-3210121300122012-2012110221011230-0310311303131221-0001303322003012-3333200102030313-1032022122132212-2033122302113231-2312022311223310"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.site.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3211220120203012-1121320222113230-0310020112232031-0012131223111031-0312113003230111-1133020320200002-3121113321223103-2012022113333123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator](resources--bigip_http_proxy--reference--group-002.md#canonical-1230023122312032-1103223311003112-2213211131230113-1201222200312112-0222310213120110-1103130022310122-2131113221210111-2032330032013313)
- origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-1330121132112212-1113012313001011-3011201133023223-1130101231133012-0200132203203022-2330010310223301-0022001133112230-2021210201331122"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
virtual_site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0230001220333010-3221200032112120-0203211123302131-0120111013123022-1133112112123023-0131303013320302-2011333020112221-2110321221123101"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site`

<a id="canonical-1230332201313020-2330100021321133-1331303310100330-3330103231201130-3322233230120032-1330102100303023-1032323031033013-1313213112023122"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1112002300210232-1101303000123323-3310110211021222-1323220103103102-0122122012122033-2213312212313011-1211213312333312-2203233301023133"></a>

<a id="canonical-2021310323002232-1303101110033103-1311130033322001-1032101302030210-2222013330122223-0213332220122311-3201100000101032-1213032013331031"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0003021032103112-0222310101301202-1231032020010131-3010002101102100-3333213303333222-2301030011132130-1232002302132211-1122032321002102"></a>

<a id="canonical-1201323331001302-2120011330300201-1302022022111001-0033033220113220-3200302133201021-2231223123201033-0300223220110011-1211003113201231"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.site_locator.virtual_site.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0012323013103221-2113331011300123-0022012222333113-1101000223003033-2320202111220110-1232200202020301-2020120203313201-0021111002012011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool

<a id="canonical-2032001111021122-2102131012223222-3102113022022213-1033013131323312-1202313102131023-1232122121101233-2033002200013133-2323203313312002"></a>

Type: `"object"`. single nested block, Optional.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

Terraform syntax:

```terraform
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0303021300201203-2101310002103331-2222230322012323-3313223322033302-2120212030110301-0330132332002213-2210323132313033-3011313102012312"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool`

- [no_snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-3010022211110123-0032313030001133-1312103100110212-1222110220010031-3233103112302231-1202302222110102-3020210111320033-0023202030330113): complete subsection reference.

- [snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-1322330113110202-0231311211012103-1023231002223200-2023023022113213-3113230201123010-3121021133310001-1333031233211021-3311002031032232): complete subsection reference.

<a id="canonical-3010022211110123-0032313030001133-1312103100110212-1222110220010031-3233103112302231-1202302222110102-3020210111320033-0023202030330113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0012323013103221-2113331011300123-0022012222333113-1101000223003033-2320202111220110-1232200202020301-2020120203313201-0021111002012011)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-2103102333211113-3103223330031332-3123121121221230-2030330200313200-3332330112002033-0231111123111111-2220010232322121-1312101023333033"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no snat pool.

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
no_snat_pool = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1322330113110202-0231311211012103-1023231002223200-2023023022113213-3113230201123010-3121021133310001-1333031233211021-3311002031032232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3232131333111312-0323012101311031-0132212133331330-0022022302222023-3301101031131030-2203330111111210-0202012220323010-3232101203203320)
- [origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool](resources--bigip_http_proxy--reference--group-002.md#canonical-0012323013103221-2113331011300123-0022012222333113-1101000223003033-2320202111220110-1232200202020301-2020120203313201-0021111002012011)
- origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-1032330311300003-3313320310030103-2123221302330333-2031303113322031-0232300010133111-2121223233003002-2132103133022013-3021331100220011"></a>

Type: `"object"`. single nested block, Optional.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
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
snat_pool {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322332001011120-3300023233132032-3131330101103030-0311113213131000-1203202211330033-3131020303313220-0221000211320311-0013103322113321"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool`

<a id="canonical-1123111031032121-3320323233310221-0221311312301131-2122302130030111-1310121313331113-2001103020313120-2322110030023002-1303330231132220"></a>

#### `origin_pools.pools.origin_servers.origin_servers.private_ip.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Optional.

List of IPv4 prefixes that represent an endpoint.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2200332312202210-0032022133011112-0321011230001311-0210031123310011-2312030333320211-1331331202212220-0001233120122303-0011232022323133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- origin_pools.pools.origin_servers.origin_servers.public_ip

<a id="canonical-2010223030113203-0233211001220121-0020103030122012-0322021030333003-0222223232011221-1032011103210000-2033333303303230-1032132220010230"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3303112322221210-3032232130030320-0121023203231011-3033020012313000-1211132030323201-3012112231233103-2120220021102321-1131113322012332"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.public_ip`

<a id="canonical-1322132030010121-3333233020031300-3012030121233123-2113012311220203-0113131030001030-0212133011301220-2303021123210322-0232133232211212"></a>

#### `origin_pools.pools.origin_servers.origin_servers.public_ip.ip` property

Type: `"string"`. Optional.

Public IPv4. Exclusive with \[\] Public IPv4 address.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2002330123011211-1100203020020233-2200222203213100-0323123022123023-2303301212202330-1220032321201102-3002033230123013-0330123233202110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_pools.pools.origin_servers.origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [origin_pools](resources--bigip_http_proxy--reference--group-001.md#canonical-3130302100213032-1133033213001002-2331132101322013-0203331122233331-3021113033330230-2313122333001332-0303103133232110-2002132020103202)
- [origin_pools.pools](resources--bigip_http_proxy--reference--group-001.md#canonical-1023031131101223-0333002000010022-2011021202323002-2121003101012320-0131313002113130-0302023223110202-2203300233110313-3320003203232022)
- [origin_pools.pools.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-1312222213103100-0321300110133131-3312031033013332-1010303313020221-3311012303023212-1010020102000132-3130213301103033-0122010331103022)
- [origin_pools.pools.origin_servers.origin_servers](resources--bigip_http_proxy--reference--group-001.md#canonical-3122233331022000-3130002202110102-2203300321323221-1321111230302130-2010113302011332-2020303212023012-2310303023303012-3120020110120320)
- origin_pools.pools.origin_servers.origin_servers.public_name

<a id="canonical-3302301122002013-1230313322333200-3202322111200011-1220333212323001-3101001320231000-3121121102003322-1303101233120102-2322002123030033"></a>

Type: `"object"`. single nested block, Optional.

Specify origin server with public DNS name.

Receipt-pinned upstream constraints:

```json
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
public_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121111332001000-3220213310121330-3323120221112003-1003231100021203-0210202203203311-3302322113033112-3302120103300213-0031211330210232"></a>

### Direct properties for `origin_pools.pools.origin_servers.origin_servers.public_name`

<a id="canonical-2311332013331031-0313230230033031-1330302213112331-1022123032312131-0230230110331312-0023020021322303-2102130222013013-2213102020213320"></a>

#### `origin_pools.pools.origin_servers.origin_servers.public_name.dns_name` property

Type: `"string"`. Optional.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0123321023020110-1321322121101332-0223330211031023-2122132021301320-2021211312200021-1032021002012120-3330000212131300-1221331120303333"></a>

<a id="canonical-2000002121020310-3100300330123011-1110223101201031-3003320012211223-3321200213112210-3022331330230122-2312300223020231-1031310131233001"></a>

#### `origin_pools.pools.origin_servers.origin_servers.public_name.refresh_interval` property

Type: `"number"`. Optional.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- proxy_advertisement

<a id="canonical-3132000202003233-2200103133232231-3123032303033200-2022013023131200-1131211311013010-1311220332231213-1130023102200031-2310220333022030"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for proxy advertisement.

Additional upstream details:

Proxy Advertisement Type.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"do_not_advertise\"]"
}
```

Terraform syntax:

```terraform
proxy_advertisement {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221021101332232-2210313100123221-3131112232130011-1212002202031210-2133312311003102-0210202103101023-3310030201111313-2201032021211233"></a>

### Direct properties for `proxy_advertisement`

- [advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330): complete subsection reference.

- [do_not_advertise](resources--bigip_http_proxy--reference--group-003.md#canonical-1332033121310100-2011203031301030-3100012202102111-0231031031001013-0100201312332130-2311031012211232-0303100202100230-3112212033310122): complete subsection reference.

<a id="canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- proxy_advertisement.advertise_custom

<a id="canonical-1311111120023312-0312200122211022-2203203010020131-3030021221200200-1321200313112302-2131221110202100-0000022232000331-3300010301300011"></a>

Type: `"object"`. single nested block, Optional.

This defines a way to advertise a VIP on specific sites.

Receipt-pinned upstream constraints:

```json
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
advertise_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3212310212202213-0310203130223010-1322231202022011-1220323011203011-3101221212332123-1232113320100120-1330210021122310-1000120030333310"></a>

### Direct properties for `proxy_advertisement.advertise_custom`

- [advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033): complete subsection reference.

<a id="canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- proxy_advertisement.advertise_custom.advertise_where

<a id="canonical-2111210223102301-3133011231011332-1331231303230113-0322321020123313-2020103311011011-3010200032113023-1323301033000100-1120000033033231"></a>

Type: `"object"`. list nested block, Optional.

Where should this load balancer be available.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0031113100120313-3230321312033003-1033321012121022-3331221323212013-2131321302003133-1111202303211121-2330323111011312-3112221331211003"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where`

- [advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-3031122032221021-2120211132203121-2301223223001230-0133010201022012-0003110211223022-0003310233032233-3313013210130230-0211033103203112): complete subsection reference.

- [advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-0111110010031303-3320222202220012-3211210031202310-0022212121230221-2312201030233321-3301023232320130-2320200012101003-3322030020222203): complete subsection reference.

- [advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-3002301022000021-0012300213120021-2233322133131231-1012301111012222-3032100322220231-0213233311031202-0220001221013121-3220201230210303): complete subsection reference.

<a id="canonical-0301112112213301-1022222002130222-3132203032012111-3323021013100330-3331003120021203-3323231223112033-1332211011201232-0120223313112022"></a>

<a id="canonical-3313220013033103-1120221030311330-0211100111000231-1032102021003213-0000101003023302-2322000220311303-3322322233003100-0203111300333122"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.port` property

Type: `"number"`. Optional.

Exclusive with \[port\_ranges use\_default\_port\] Port to Listen.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1333121130033023-0033203310302213-2312201000102001-2123003201020221-3010233323333203-0112012030122312-1132000212122113-2211001211301331"></a>

<a id="canonical-3302203202303200-3012022033322001-0301013310233333-2310102203131100-0032331002121100-0021002030211220-1023321312332123-3300030000230122"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.port_ranges` property

Type: `"string"`. Optional.

Exclusive with \[port use\_default\_port\] A string containing a comma separated list of port
ranges. Each port range consists of a single port or two ports separated by '-'.

Additional upstream details:

Each port range consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-1003101232131113-0113311323222200-0120022332103213-2021333303210112-2220131221030220-1213213031233003-3021232001130223-2313001300302330): complete subsection reference.

- [use_default_port](resources--bigip_http_proxy--reference--group-002.md#canonical-2011020322221223-2101212111130110-0220200033301330-2230203021003003-1312331100030022-0103203013330023-2130302133021003-3201302011030110): complete subsection reference.

- [virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-0303023203101011-3310022012313303-0310200232211333-2320120210030203-3001331213320020-1120321132301203-0320311000020123-1101213123321031): complete subsection reference.

- [virtual_site](resources--bigip_http_proxy--reference--group-003.md#canonical-2333212000313303-2000232320323310-1132310130123202-1021333123010110-2300203221100121-1202232211130120-3321213112310102-1011112111330230): complete subsection reference.

- [virtual_site_with_vip](resources--bigip_http_proxy--reference--group-003.md#canonical-2210222003023302-3322232233330320-1133232011210322-2302301032121031-3332022221001312-0223333301221201-1212020213132130-1210122103301303): complete subsection reference.

- [vk8s_service](resources--bigip_http_proxy--reference--group-003.md#canonical-3201012210311222-0013322302303000-0033333010112020-1110200303031321-2310032110213210-2223222313100032-1222203010020321-1021221212131032): complete subsection reference.

<a id="canonical-3031122032221021-2120211132203121-2301223223001230-0133010201022012-0003110211223022-0003310233032233-3313013210130230-0211033103203112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public

<a id="canonical-0120021301122100-2030100011101123-0131130333003222-3232101122223312-0100131303233011-3012000102312103-0103001101231010-1132100013233010"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1033222233110111-3133211032123001-2101323200212100-3212301020113310-1010020200132320-1301333233030321-3323131102202132-3101112213332323"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public`

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-3111130020221201-3200233200023203-3121133331223220-0101202202202102-0120212020312112-0010002131012321-1220001101130013-2011230322212313): complete subsection reference.

<a id="canonical-3111130020221201-3200233200023203-3121133331223220-0101202202202102-0120212020312112-0010002131012321-1220001101130013-2011230322212313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-3031122032221021-2120211132203121-2301223223001230-0133010201022012-0003110211223022-0003310233032233-3313013210130230-0211033103203112)
- proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip

<a id="canonical-1131123230220011-0333210220032210-3213321011220300-3100303231332312-2100210230201020-1021131232002003-3133131120010320-2132212312210021"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332210322313003-3031222022112213-1200131131103211-0221233310011022-3220322100323032-2123202111302032-0021312202201311-3031303131123333"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip`

<a id="canonical-2113310020200232-0032200112320230-3223332111123311-0333112202030301-2300120312101003-1130310003121322-0123102210102103-1303010303102321"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1123121121113111-0103233131303213-0000231220020230-3320123232212101-3200003211330332-3113331102121032-3310130132021333-3100100333233132"></a>

<a id="canonical-1322210101010131-1022031130202020-3002012022332022-3230332122020321-1020112322312132-3131123330130111-3001033103332333-3220033232233221"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0012221033023300-1200302003110303-0000312233211030-1123033332132232-3331200032322120-2001313020012233-0120221001132023-3330112010003200"></a>

<a id="canonical-1100122311131112-0203203100222210-1230201223013312-1012220223332203-1322012032212113-0330120331101110-3233012000202310-0310311311202013"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_dualstack_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0111110010031303-3320222202220012-3211210031202310-0022212121230221-2312201030233321-3301023232320130-2320200012101003-3322030020222203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public

<a id="canonical-0101302102033232-3222112100301323-3212211020000210-2320032230123112-0101200230332111-2230022301030320-3311101002112211-0302032001313333"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-1310003122013120-3223112111122101-2321120131012011-3013122022032002-0100311012210222-3320022330021200-0132310313013213-3222111203032032"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public`

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-2022132312221203-0323231333310101-0121221312323222-3003330300330020-2002313213223211-2211103120110103-1020103023102132-1002013323220211): complete subsection reference.

<a id="canonical-2022132312221203-0323231333310101-0121221312323222-3003330300330020-2002313213223211-2211103120110103-1020103023102132-1002013323220211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-0111110010031303-3320222202220012-3211210031202310-0022212121230221-2312201030233321-3301023232320130-2320200012101003-3322030020222203)
- proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip

<a id="canonical-3330320333320032-0301001100232020-1023122320222302-2312232013321022-0032122211131102-0222303330101003-1333020302210020-1302211101120101"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3130113203222103-0020221122103102-2002113313330332-0032003233022002-2331103001301220-3223113301331132-1333322110200111-3300302321332230"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip`

<a id="canonical-3211301212122000-2211303110202200-0122322222103032-0333332200322031-0320223103233003-2313012131222211-1320212131303313-1330110322022032"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3111103001220101-3310110132323100-3333323301301231-3220001012303132-3133010111220230-3333302010003100-0333201021120111-3100111020330311"></a>

<a id="canonical-0120001121203032-0023311312021100-3030112130211231-3130201133000101-2333220222013322-1121011223001131-1330022103200220-2332102112330130"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2300312032231221-2333010112223233-2131132221320322-1012123123121032-3332013031013331-1030313203121021-2133332012233211-2200303312312313"></a>

<a id="canonical-1113303301223130-2332331200112001-2302100030211210-0032320011231110-1303101213322001-0232112012322110-2222233123320302-1330131000110002"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3002301022000021-0012300213120021-2233322133131231-1012301111012222-3032100322220231-0213233311031202-0220001221013121-3220201230210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public

<a id="canonical-0321330023321310-2021100133302101-3031230202132103-1000320203202112-2013020131221203-3003021213321223-0331113302123033-3001212033111001"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3113223322012110-0313221301320000-3122012203123322-1201123133023212-2131330011202013-3001111131212000-2332321221233222-2112220232320202"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public`

- [public_ip](resources--bigip_http_proxy--reference--group-002.md#canonical-2312211233211301-2310300101032323-3311021311123302-0001203312232000-1003202333111102-0300130302022021-0023103022120112-0331200213221132): complete subsection reference.

<a id="canonical-2312211233211301-2310300101032323-3311021311123302-0001203312232000-1003202333111102-0300130302022021-0023103022120112-0331200213221132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public](resources--bigip_http_proxy--reference--group-002.md#canonical-3002301022000021-0012300213120021-2233322133131231-1012301111012222-3032100322220231-0213233311031202-0220001221013121-3220201230210303)
- proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip

<a id="canonical-2232313000130002-3103222300301130-3031101201230022-3210113031221220-0312121000311201-0101320112332001-2223201221230010-3330201120321002"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
public_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010300331131023-0020101113311030-1201220322231020-1103300213123220-0310302313130130-1200332113031303-0033232111001212-2211312031212323"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip`

<a id="canonical-1030122130221222-1123231212033213-2121222233130301-1011002111310332-2312030212301303-0123221330333123-2111100330100203-3123212120310012"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2122232012111211-2201002230100210-3113320002001012-0321210102121133-3112120102121300-3310313010111210-2322323002212200-3301201323301023"></a>

<a id="canonical-0033022330203012-2003102112112112-0333310103233203-2132223313112130-3320303320021120-3032213122320023-0202300322002012-0331321120222201"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1031320200012330-0213210102302123-0113023211033113-0001330201032332-0222130312113203-0120003300202233-3131201213132321-0103331312223122"></a>

<a id="canonical-0021232231220310-3312232322202010-0112110332011310-3310312013130133-0033123020000232-3300223130102222-2111120021302121-2300212022101322"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.advertise_v6_on_public.public_ip.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1003101232131113-0113311323222200-0120022332103213-2021333303210112-2220131221030220-1213213031233003-3021232001130223-2313001300302330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.site

<a id="canonical-0130122121002133-1122221012023322-1233313031202000-3000131030300231-0331002011211222-1203313031022123-2330103220032222-2110122102301030"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-2033203300202032-2311331012303003-2310121013313012-1103302030331322-1301201022100123-0303321231232210-3222030231202022-2011032332103102"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.site`

<a id="canonical-1331232011212220-1113201323130303-0311322010030300-0110303231233323-0032133123301333-2320222023331303-2031311020312121-3130303000021231"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.ip` property

Type: `"string"`. Optional.

Use given IP address as VIP on the site.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3122300101223300-0333311302020123-2030302333102330-1130302103322030-2223332201312031-3123330122102120-2210030131213021-2121132102113022"></a>

<a id="canonical-2222332201203302-1330031113302131-1001301122023313-3102310210333000-3131333211302330-1120220201121310-2203213010312131-0102330303233113"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.network` property

Type: `"string"`. Optional.

\[Enum:
SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE|SITE\_NETWORK\_INSIDE|SITE\_NETWORK\_OUTSIDE|SITE\_NETWORK\_SERVICE|SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP|SITE\_NETWORK\_IP\_FABRIC\]
Defines network types to be used on site All inside and outside networks. All inside and outside
networks with internet VIP support. All inside networks. Possible values are
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`, \`SITE\_NETWORK\_INSIDE\`, \`SITE\_NETWORK\_OUTSIDE\`,
\`SITE\_NETWORK\_SERVICE\`, \`SITE\_NETWORK\_OUTSIDE\_WITH\_INTERNET\_VIP\`,
\`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\_WITH\_INTERNET\_VIP\`, \`SITE\_NETWORK\_IP\_FABRIC\`.
Defaults to \`SITE\_NETWORK\_INSIDE\_AND\_OUTSIDE\`.

Additional upstream details:

This defines network types to be used on site

All inside and outside networks. All outside networks. All outside networks with internet VIP
support. VK8s service network. &#8203;- SITE\_NETWORK\_IP\_FABRIC: VER IP Fabric network for the
site

This Virtual network type is used for exposing virtual host on IP Fabric network on the VER site or
for endpoint in IP Fabric network.

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

- [site](resources--bigip_http_proxy--reference--group-002.md#canonical-3103110021033100-1232301311033023-3301122111303112-1012101033123321-0211021223200300-3100032003331002-1330020222231223-2012103231303222): complete subsection reference.

<a id="canonical-3103110021033100-1232301311033023-3301122111303112-1012101033123321-0211021223200300-3100032003331002-1330020222231223-2012103231303222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.site.site` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.site](resources--bigip_http_proxy--reference--group-002.md#canonical-1003101232131113-0113311323222200-0120022332103213-2021333303210112-2220131221030220-1213213031233003-3021232001130223-2313001300302330)
- proxy_advertisement.advertise_custom.advertise_where.site.site

<a id="canonical-2013032023310312-0031322113312230-3332200120231012-0001002101013123-0100020003132021-0000122003233322-2200112132303132-0221012231301203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
site {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232131210103233-1000223032220013-0222203213303201-1211132303100131-0323120101121110-2010313002010311-0032221001000002-2003100113230200"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.site.site`

<a id="canonical-0201011123020221-0330033131131013-2223231212133022-2132032312023210-3312132331312033-1333133023320321-2110133003323120-1110000111321113"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0021333213121221-1111132010232103-0003022131033120-2002330113122301-2221302213020010-1212023032023323-0031330123331331-0302232311022310"></a>

<a id="canonical-3211330321323131-2212213330013302-0223211003023030-3221200321213032-2032223001020003-0221310123210121-1332120302102333-0132202222010323"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2130110320200002-2110133002112013-2320330233030131-3123212011010022-2302320301231122-2013100300313230-0331332100120331-0301120011332120"></a>

<a id="canonical-1333321030001010-2223223231132222-3310133303123322-1121302233122201-2113232103212003-0300103222330230-1223312321123112-0232212031120201"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.site.site.tenant` property

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2011020322221223-2101212111130110-0220200033301330-2230203021003003-1312331100030022-0103203013330023-2130302133021003-3201302011030110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.use_default_port` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.use_default_port

<a id="canonical-2013220330230101-2003131300000130-1012200023330101-1333300111011111-2013012233321220-0203303100032221-2321301121110103-3000322103322310"></a>

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
use_default_port = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303023203101011-3310022012313303-0310200232211333-2320120210030203-3001331213320020-1120321132301203-0320311000020123-1101213123321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network

<a id="canonical-2000021012300133-0223211112310033-2331021312120301-2313132120011121-3113333332332113-1322020000320032-0200131310201213-2230332300033033"></a>

Type: `"object"`. single nested block, Optional.

Parameters to advertise on a given virtual network.

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

<a id="canonical-2223212311012111-3120213022131330-2231033101223022-1312233332133232-2233200300113123-2132231031323112-3323300323013133-3313200333323031"></a>

### Direct properties for `proxy_advertisement.advertise_custom.advertise_where.virtual_network`

- [default_v6_vip](resources--bigip_http_proxy--reference--group-002.md#canonical-2232132011202230-0011020032122231-2322101110210313-2213210330131212-3213102111002233-0023310113202322-0112000111002302-2120310221102111): complete subsection reference.

- [default_vip](resources--bigip_http_proxy--reference--group-003.md#canonical-3020010230021021-3130200302202230-0313320010003112-3211122021202113-3033220110320231-3023202202311100-3303230112201010-2123122113332211): complete subsection reference.

<a id="canonical-3032021001323100-1300201313010310-3102031313322320-1113032020320323-1022112212210030-0322230333211300-2233112333333130-3010123031111030"></a>

<a id="canonical-1321001132123233-2232321223031313-0022032202122012-3302312010111013-2220213012212211-0222133332032321-3332231200012132-2311103223011003"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_v6_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_v6\_vip\] Use given IPv6 address as VIP on virtual Network.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0310203013301232-0022102032113130-3300123120231130-0320021010210123-0231122000222022-2021233021330332-1120211123232010-0020111021223012"></a>

<a id="canonical-0000021310102303-2201201300103032-3133321003021010-0301210023122001-2003320011023003-0102210320130331-3020230030232112-0211223310330021"></a>

#### `proxy_advertisement.advertise_custom.advertise_where.virtual_network.specific_vip` property

Type: `"string"`. Optional.

Exclusive with \[default\_vip\] Use given IPv4 address as VIP on virtual Network.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [virtual_network](resources--bigip_http_proxy--reference--group-003.md#canonical-3222021333132003-2333132213331001-3133203331320301-0111213323031103-2331130213031223-3120230321300133-1220111302102031-3002001333122330): complete subsection reference.

<a id="canonical-2232132011202230-0011020032122231-2322101110210313-2213210330131212-3213102111002233-0023310113202322-0112000111002302-2120310221102111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip` properties

Breadcrumbs:

- [xcsh_bigip_http_proxy](../resources/bigip_http_proxy.md#canonical-1032100020020330-1220313211200112-1331211020131131-0320101130102310-2022002330330313-3011010123201310-3333021132203201-3110113102200111)
- [Property reference](resources--bigip_http_proxy--reference--group-001.md#canonical-0101201332222201-3302033021001322-1302101210311331-2210333210320003-0002033231003211-1133020003110203-0200133113131131-0101001111122000)
- [proxy_advertisement](resources--bigip_http_proxy--reference--group-002.md#canonical-0222202123223213-1021332332331003-3333013231022023-1200013112100010-3021102122231210-1210031321121132-3303131312332322-2002012022311122)
- [proxy_advertisement.advertise_custom](resources--bigip_http_proxy--reference--group-002.md#canonical-0003302100123003-0120302302033120-2033221231301030-1321310211301112-2322203120001123-1230320101110131-1230033322322311-2022333113123330)
- [proxy_advertisement.advertise_custom.advertise_where](resources--bigip_http_proxy--reference--group-002.md#canonical-0110333311321122-0102322332130222-0013022122311102-0030130031302023-0132011232311101-1023102213023221-2123002103132132-1010332102003033)
- [proxy_advertisement.advertise_custom.advertise_where.virtual_network](resources--bigip_http_proxy--reference--group-002.md#canonical-0303023203101011-3310022012313303-0310200232211333-2320120210030203-3001331213320020-1120321132301203-0320311000020123-1101213123321031)
- proxy_advertisement.advertise_custom.advertise_where.virtual_network.default_v6_vip

<a id="canonical-3031102321023301-0003233200231212-2100221311200220-1021002103112311-2121132033022033-1101123331113010-3202012121110012-3213323011331103"></a>

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
default_v6_vip = {}
```

This is an empty object or choice marker. It has no direct properties.
