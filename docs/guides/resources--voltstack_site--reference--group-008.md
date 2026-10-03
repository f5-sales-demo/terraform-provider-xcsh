---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-1300132202203000-2012301020010011-3211010200133221-3110123301021131-1130200011112201-2203111103301130-2123002300123022-1030233312132022"></a>

## provider_ref property — clear_secret_info / 212132010212 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1113130301130232-3322201112130001-0322313112211020-3310202010032002-3200033003023300-3221303013021313-0121023002113013-1130300003320232"></a>

<a id="canonical-2120220103012230-0030220111020222-0101232001322030-0323110223010311-2120311021322131-2000321111301022-0322001000221231-0331022112323102"></a>

## URL property — clear_secret_info / 212132010212 / 5

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Upstream description:

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3231100030323201-3311303100201210-2211103012021210-3003133202313022-2222000102011222-2332332003210212-3330123210203011-0210121001003202"></a>

## Next pages — clear_secret_info / 212132010212 / 6

- [custom_storage_config.storage_device_list.storage_devices.pure_service_orchestrator.arrays.flash_blade.flash_blades.api_token](resources--voltstack_site--reference--group-007.md#canonical-1331201220120101-0231302101001011-1111031302133231-2032331223220320-2233102032132032-2210323011310300-1000000223320012-3000111203123111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322111312111000-3132223320230311-1121313333302001-2233222023211100-1231303201012201-1022210320311231-1220222331213321-0323002012130202"></a>

## custom_storage_config.storage_interface_list — storage_interface_list / 222032103311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- custom_storage_config.storage_interface_list

<a id="canonical-1323311102310220-3220221221000133-3221112033132101-0201202120200310-3330332012110333-2232012012300312-3103000311121310-1333010101312333"></a>

Type: `"object"`. single nested block, Optional.

Configure storage interfaces for this App Stack site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("storage_interfaces")}
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
storage_interface_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2110330110103302-0022223110231300-2030010321232303-1113222333223311-1010121221311303-2332311130330311-3020231333023313-0123000203130211"></a>

## Direct properties — storage_interface_list / 222032103311 / 3

- [storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133): complete subsection reference.

<a id="canonical-1000121113233023-2202333020223202-1120300011201111-0210201312310313-1022110021331231-1002023331133300-0303002322103310-2213113131201002"></a>

## Next pages — storage_interface_list / 222032103311 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230301203223202-2120010322213210-3230011321123323-1210203232333310-1023130122033222-2211223223312000-3222110131302132-1220133131101210"></a>

## custom_storage_config.storage_interface_list.storage_interfaces — storage_interfaces / 200313010221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- custom_storage_config.storage_interface_list.storage_interfaces

<a id="canonical-0303222022120110-2100302223001103-2333302113333032-2320313011303220-0210310213331201-3312310200112101-3222002011330301-2210233132132222"></a>

Type: `"object"`. list nested block, Optional.

Configure storage interfaces for this App Stack site.

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
storage_interfaces {
  # Configure direct properties listed below.
}
```

<a id="canonical-1322231023103012-0123130210102112-3110010231021103-0220230312002111-1123300200303202-2031223002020313-2301112220233323-0303002001123222"></a>

## Direct properties — storage_interfaces / 200313010221 / 3

<a id="canonical-1010123203021022-3233130300201110-1101320303301323-2123001331132321-3202120103222110-3113111310222202-2330000132001133-0331000132022011"></a>

<a id="canonical-2300212112110202-0123233310213013-2003110331123130-0313223032100300-1323013232300211-3223223111030200-1323302022231111-2003200332222112"></a>

## description_spec property — storage_interfaces / 200313010221 / 4

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [labels](resources--voltstack_site--reference--group-008.md#canonical-3020330233311233-1013013103121221-0213133330313113-3121202110111002-3303021110200311-0323013123230323-3113132010310301-3232032120233031): complete subsection reference.

- [storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001): complete subsection reference.

<a id="canonical-3122220223231010-2122332203101231-0011110001231211-0303213133110330-2002030021021202-1030212313001112-2233203212230331-1332031330313310"></a>

## Next pages — storage_interfaces / 200313010221 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.labels](resources--voltstack_site--reference--group-008.md#canonical-3020330233311233-1013013103121221-0213133330313113-3121202110111002-3303021110200311-0323013123230323-3113132010310301-3232032120233031)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3020330233311233-1013013103121221-0213133330313113-3121202110111002-3303021110200311-0323013123230323-3113132010310301-3232032120233031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0100313321103220-3033320123111210-1302000122031030-0113203121113130-2001211131220021-2132020130030021-0201300230200031-1132200111122333"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.labels — labels / 112210203021 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- custom_storage_config.storage_interface_list.storage_interfaces.labels

<a id="canonical-1220303000313210-2002323202013220-3231111112213221-0310311011012321-1013030332203332-0302032001121033-3030210232320331-1111021131030333"></a>

Type: `"object"`. single nested block, Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0020102203120011-1020110211213000-0131020023021132-1130220023332233-0300312003311331-2103132021300031-3212001331011012-2300123000121130"></a>

## Direct properties — labels / 112210203021 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303331322131200-3333310233310210-3312102020120021-3033010330021232-2032101111012011-2200212132123002-3331001100100321-3213230221031031"></a>

## Next pages — labels / 112210203021 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232311111231113-3210222321223322-2330330033331213-0303231011232223-1021301030001132-1132020212021303-3213022321113320-0332301113001310"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface — storage_interface / 203203133022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface

<a id="canonical-2113022231013020-0321213013322313-3331130102100301-1103003030123233-1010223201011210-0013211021011332-2113132033203320-3131022030213032"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for storage interface.

Upstream description:

Ethernet Interface Configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("device"),
  validators.ConflictingObjectAttributes("cluster",
    "node"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "dhcp_server"),
  validators.ConflictingObjectAttributes("dhcp_client",
    "static_ip"),
  validators.ConflictingObjectAttributes("dhcp_server",
    "static_ip"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "no_ipv6_address"),
  validators.ConflictingObjectAttributes("ipv6_auto_config",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("is_primary",
    "not_primary"),
  validators.ConflictingObjectAttributes("monitor",
    "monitor_disabled"),
  validators.ConflictingObjectAttributes("no_ipv6_address",
    "static_ipv6_address"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "site_local_network"),
  validators.ConflictingObjectAttributes("site_local_inside_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("site_local_network",
    "storage_network"),
  validators.ConflictingObjectAttributes("untagged",
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
  },
  "x-ves-oneof-field-address_choice": "[\"dhcp_client\",\"dhcp_server\",\"static_ip\"]",
  "x-ves-oneof-field-ipv6_address_choice": "[\"ipv6_auto_config\",\"no_ipv6_address\",\"static_ipv6_address\"]",
  "x-ves-oneof-field-monitoring_choice": "[\"monitor\",\"monitor_disabled\"]",
  "x-ves-oneof-field-network_choice": "[\"segment_network\",\"site_local_inside_network\",\"site_local_network\",\"storage_network\"]",
  "x-ves-oneof-field-node_choice": "[\"cluster\",\"node\"]",
  "x-ves-oneof-field-primary_choice": "[\"is_primary\",\"not_primary\"]",
  "x-ves-oneof-field-vlan_choice": "[\"untagged\",\"vlan_id\"]"
}
```

Terraform syntax:

```terraform
storage_interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231031131332230-1010211232123311-2122222030130220-2213130201311003-2212322130112111-1302131232222101-0011331011131220-2021300333322210"></a>

## Direct properties — storage_interface / 203203133022 / 3

- [cluster](resources--voltstack_site--reference--group-008.md#canonical-1302301313023102-3202003103121122-0322132302300020-1101101221203020-2020320201202311-0200012322123112-2202210321003200-0033111323021321): complete subsection reference.

<a id="canonical-1200330212303112-1233121002332320-1302001312030100-1311330123303122-0123123003321123-0203010122231023-3221102232302333-2231021222212002"></a>

<a id="canonical-1313001302132130-0111012332230122-3011030232210130-1101030022003201-2103131321020331-3220331003221323-1033223310321132-1330303002102311"></a>

## device property — storage_interface / 203203133022 / 4

Type: `"string"`. Optional.

Interface configuration for the ethernet device.

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

- [dhcp_client](resources--voltstack_site--reference--group-008.md#canonical-3010013003002312-0112212130313311-3230112311111331-2133132012133320-1010000102230322-1333230211111321-3112123201212301-2002301302211313): complete subsection reference.

- [dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221): complete subsection reference.

- [ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131): complete subsection reference.

- [is_primary](resources--voltstack_site--reference--group-008.md#canonical-3202030333101230-1121113030133310-2311000313322112-2002221331302311-1101320221233330-3032023102230211-3102102021002002-1121321030133331): complete subsection reference.

- [monitor](resources--voltstack_site--reference--group-008.md#canonical-1320122111233123-2330221102223112-0001121121332231-3313323333121211-0311231213011220-2232300220130003-0112332010112220-1011022021110201): complete subsection reference.

- [monitor_disabled](resources--voltstack_site--reference--group-008.md#canonical-1310211020301210-1210320103002203-0123310223331313-2220310221111122-0210321011011223-3133001110021331-3121300300032001-2133311321311032): complete subsection reference.

<a id="canonical-2220222100002212-3321322333303030-1103100032010011-0211302203202132-1303211132031121-0223103131112011-0302012233032233-2232210213021102"></a>

<a id="canonical-3331310301030132-1212213332131122-3101033002232132-3133333332230332-1213120312131132-0211133131002123-3200122100033031-0221202303111323"></a>

## mtu property — storage_interface / 203203133022 / 5

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Upstream description:

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 9000.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  validators.Int64RangeSetValidator(
    validators.Int64Range{Minimum: 0, Maximum: 0},
    validators.Int64Range{Minimum: 512, Maximum: 9000},
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
    "maximum": 9000,
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
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,512-9000"
  }
}
```

- [no_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-1323110312222000-1332122013222211-1331223002101001-1320212001201110-2312030130101030-1102233103312311-1103123101232222-1323103011123230): complete subsection reference.

<a id="canonical-1003001310232023-0101112100230131-0100100001213032-3000112220121111-2303233110333030-2210132133033303-0030310133022130-1331313003332000"></a>

<a id="canonical-1123023300102320-3131000223013113-0323103313122110-3131303202001023-0330020113130203-3333001311003020-1111003203021033-3011131223213313"></a>

## node property — storage_interface / 203203133022 / 6

Type: `"string"`. Optional.

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

Upstream description:

Exclusive with \[cluster\] Configuration will apply to a device on the given node.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [not_primary](resources--voltstack_site--reference--group-008.md#canonical-2030001302031213-3301323223031021-0010112321210112-0032333302131232-0312133110201110-0211130013223333-2302313001012231-2000112312301001): complete subsection reference.

<a id="canonical-2010012230133132-1210323012022232-2032100110033103-3332120101110001-2102100201122320-1203223321331022-3133233133023220-0130003122022232"></a>

<a id="canonical-2213032002100230-2302011021231020-2300030210221102-1300312222203310-2012102032313212-0200121101102310-0000020223002203-0322230233230103"></a>

## priority property — storage_interface / 203203133022 / 7

Type: `"number"`. Optional.

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

Upstream description:

Priority of the network interface when multiple network interfaces are present in outside network
Greater the value, higher the priority.

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

- [site_local_inside_network](resources--voltstack_site--reference--group-008.md#canonical-2021132332003322-2303131302223130-0132230213011030-3311302120210010-3023311123330121-3303223001232121-2131133100013301-0011212303203201): complete subsection reference.

- [site_local_network](resources--voltstack_site--reference--group-008.md#canonical-2320002321301231-0320321323322121-2033113021100333-2201201302223131-1222101001333311-0230301320301111-2023232222010002-1100123212233000): complete subsection reference.

- [static_ip](resources--voltstack_site--reference--group-008.md#canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311): complete subsection reference.

- [static_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-3200300220203231-3023330333022333-3222222223230021-2200301322031220-2002330223023112-3221031232322100-2223212103201111-3333020321210313): complete subsection reference.

- [storage_network](resources--voltstack_site--reference--group-009.md#canonical-0032030332120031-2202202201100231-2322202223131331-2132011101013030-3020033220000320-2031113022110130-3321121032110233-0101132123101030): complete subsection reference.

- [untagged](resources--voltstack_site--reference--group-009.md#canonical-0233020120313000-3301111332332201-3110320120321031-2100012123320130-3220103132132023-2211113313111133-0200200131121303-2122122320303120): complete subsection reference.

<a id="canonical-0300100111231023-1212333010223023-1303230011330122-3311300233111130-3010120211210003-0102011301000032-2101310333133332-2303210233300221"></a>

<a id="canonical-3021113232121310-3112113323331022-3021301130132312-2031020212302122-0122302010201212-3220002311232331-0010302212120233-0233021202021130"></a>

## vlan_id property — storage_interface / 203203133022 / 8

Type: `"number"`. Optional.

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

Upstream description:

Exclusive with \[untagged\] Configure a VLAN tagged ethernet interface.

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
    "create": false,
    "minimum_config": false,
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

<a id="canonical-0212223223121031-1312121302332221-0103030000232121-2002023222013301-0201022311100002-2212323220332300-3111333313323120-1232010222333122"></a>

## Next pages — storage_interface / 203203133022 / 9

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster](resources--voltstack_site--reference--group-008.md#canonical-1302301313023102-3202003103121122-0322132302300020-1101101221203020-2020320201202311-0200012322123112-2202210321003200-0033111323021321)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client](resources--voltstack_site--reference--group-008.md#canonical-3010013003002312-0112212130313311-3230112311111331-2133132012133320-1010000102230322-1333230211111321-3112123201212301-2002301302211313)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary](resources--voltstack_site--reference--group-008.md#canonical-3202030333101230-1121113030133310-2311000313322112-2002221331302311-1101320221233330-3032023102230211-3102102021002002-1121321030133331)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor](resources--voltstack_site--reference--group-008.md#canonical-1320122111233123-2330221102223112-0001121121332231-3313323333121211-0311231213011220-2232300220130003-0112332010112220-1011022021110201)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled](resources--voltstack_site--reference--group-008.md#canonical-1310211020301210-1210320103002203-0123310223331313-2220310221111122-0210321011011223-3133001110021331-3121300300032001-2133311321311032)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-1323110312222000-1332122013222211-1331223002101001-1320212001201110-2312030130101030-1102233103312311-1103123101232222-1323103011123230)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary](resources--voltstack_site--reference--group-008.md#canonical-2030001302031213-3301323223031021-0010112321210112-0032333302131232-0312133110201110-0211130013223333-2302313001012231-2000112312301001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network](resources--voltstack_site--reference--group-008.md#canonical-2021132332003322-2303131302223130-0132230213011030-3311302120210010-3023311123330121-3303223001232121-2131133100013301-0011212303203201)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network](resources--voltstack_site--reference--group-008.md#canonical-2320002321301231-0320321323322121-2033113021100333-2201201302223131-1222101001333311-0230301320301111-2023232222010002-1100123212233000)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](resources--voltstack_site--reference--group-008.md#canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address](resources--voltstack_site--reference--group-008.md#canonical-3200300220203231-3023330333022333-3222222223230021-2200301322031220-2002330223023112-3221031232322100-2223212103201111-3333020321210313)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.storage_network](resources--voltstack_site--reference--group-009.md#canonical-0032030332120031-2202202201100231-2322202223131331-2132011101013030-3020033220000320-2031113022110130-3321121032110233-0101132123101030)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.untagged](resources--voltstack_site--reference--group-009.md#canonical-0233020120313000-3301111332332201-3110320120321031-2100012123320130-3220103132132023-2211113313111133-0200200131121303-2122122320303120)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1302301313023102-3202003103121122-0322132302300020-1101101221203020-2020320201202311-0200012322123112-2202210321003200-0033111323021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221211302231101-0232232113133231-1212011232020121-2010132100130232-3231132200022323-2001330000102332-3323112230322120-3202300331113001"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster — cluster / 203110203123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.cluster

<a id="canonical-3131313200301122-3303332011032020-2223220013323131-0221131322100303-0300333100301231-1230111002302122-1001220110023200-3101220211032330"></a>

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
cluster = {}
```

<a id="canonical-1013011112202210-0233332123031312-2020000131112122-0320023202033311-2231022222302323-1213300211023212-3013313210310030-1311212112021111"></a>

## Direct properties — cluster / 203110203123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2210003213001203-3103131333303333-2331033021022221-1232210103232233-3220003102302313-3123032302233233-2123220223110002-0103201111012220"></a>

## Next pages — cluster / 203110203123 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3010013003002312-0112212130313311-3230112311111331-2133132012133320-1010000102230322-1333230211111321-3112123201212301-2002301302211313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203233001011302-1111110312202221-2211233212333101-3030323133001120-2213003223103132-0001322020310311-3310011310203001-0331011231311200"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client — dhcp_client / 133321022221 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_client

<a id="canonical-0223030210021322-3211332201032012-2133031222231113-3331010132312301-3101011220312321-2132211322032203-2131101032123102-0021103223322311"></a>

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

<a id="canonical-3320311222301100-0323003222022023-0222030301210133-0222232231333023-1232122002122322-2011013013121121-0331220213102110-3112321321203121"></a>

## Direct properties — dhcp_client / 133321022221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3211010100233133-2210033312123331-2232102110000313-2122323110310030-2313103201230000-2302113300121121-2102203010032020-1010303113021211"></a>

## Next pages — dhcp_client / 133321022221 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3113131221301303-2112233320032020-3113313010123310-1100000333211002-0121232000013131-0022120120010331-3330132200033321-0202310320003003"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server — dhcp_server / 322013122113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server

<a id="canonical-3012111102221111-1331322120011030-2233022122021320-3021323000112003-1213313220203333-2111013000230321-0221332301131311-3022330120223033"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for dhcp server.

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
dhcp_server {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121220101203020-3230332012201101-1310010033013002-1300123233003320-1122020011003301-0230103220303110-3201333133213113-1331203120203010"></a>

## Direct properties — dhcp_server / 322013122113 / 3

- [automatic_from_end](resources--voltstack_site--reference--group-008.md#canonical-2330301002310023-1332201233202123-2332331131333302-0212323101310331-1123102102231201-0203132313312130-1133213332203002-2010331133000031): complete subsection reference.

- [automatic_from_start](resources--voltstack_site--reference--group-008.md#canonical-0031003022312200-0101320333211111-1330302101022022-3323232222112011-3200033331200030-3211102102000130-1210132210023113-1123331312321213): complete subsection reference.

- [dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311): complete subsection reference.

<a id="canonical-2100302330103232-1132332320212301-2212231312231322-2332012120023332-3003101022020101-2320321110322033-3212320132010013-2122330310320111"></a>

<a id="canonical-1122312330231110-1321023130200100-2201232002011003-2310203033003220-0000032211122220-0102001001302032-3322311322332130-1100301203123011"></a>

## dhcp_option82_tag property — dhcp_server / 322013122113 / 4

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-0310021330310213-3130013101212021-2321022031131111-2310231331000221-1113103103310212-3230021003112000-3220011201130003-3023211222032121"></a>

<a id="canonical-0000131110110031-3112201232330121-2311211331103200-3112102103113302-3322120301201321-2003101303020130-1130203333030011-0212101310100300"></a>

## fixed_ip_map property — dhcp_server / 322013122113 / 5

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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

- [interface_ip_map](resources--voltstack_site--reference--group-008.md#canonical-1110000210131331-0222211003000203-2323232230323312-0332322023200323-0300013122103330-0031301111323001-0120123303121031-3131001303100232): complete subsection reference.

<a id="canonical-3102032333023113-3203312321133303-3313213022333330-1230023120021112-0233130212330102-0302223103031002-3112132002203310-2330003132201012"></a>

## Next pages — dhcp_server / 322013122113 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end](resources--voltstack_site--reference--group-008.md#canonical-2330301002310023-1332201233202123-2332331131333302-0212323101310331-1123102102231201-0203132313312130-1133213332203002-2010331133000031)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start](resources--voltstack_site--reference--group-008.md#canonical-0031003022312200-0101320333211111-1330302101022022-3323232222112011-3200033331200030-3211102102000130-1210132210023113-1123331312321213)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map](resources--voltstack_site--reference--group-008.md#canonical-1110000210131331-0222211003000203-2323232230323312-0332322023200323-0300013122103330-0031301111323001-0120123303121031-3131001303100232)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2330301002310023-1332201233202123-2332331131333302-0212323101310331-1123102102231201-0203132313312130-1133213332203002-2010331133000031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113323003202112-1220112221102300-2331002033022300-3203213303203320-1112210211233211-0111111001320231-2020230320023120-0201333030100131"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end — automatic_from_end / 020232111311 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_end

<a id="canonical-3101310031011322-0132222221102122-0002233101333001-0213000323333233-3220033133302000-3122010213010331-1121003302012200-1010110031222023"></a>

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

<a id="canonical-1302310300030213-0313202020211101-3203123323301222-0213101201033100-3132032113302311-3301302102020231-2310303122213200-0300310231001030"></a>

## Direct properties — automatic_from_end / 020232111311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3321123011010320-3200222020112311-3033220011030322-2332123212300001-1033320001011002-0220103202300221-3232312033302303-3212200202213331"></a>

## Next pages — automatic_from_end / 020232111311 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0031003022312200-0101320333211111-1330302101022022-3323232222112011-3200033331200030-3211102102000130-1210132210023113-1123331312321213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313213120110120-1230302023113320-3011032231110030-1021301111012032-1103111003130033-1323221301201320-1111011310121203-0012201011223012"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start — automatic_from_start / 001313001201 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.automatic_from_start

<a id="canonical-2201213222110112-0031130020312310-0001331112110120-1121223031032011-0010331300131103-3302332303010003-0233013021300013-1300102000200032"></a>

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

<a id="canonical-2010120300302121-0002232220133321-3210012120121003-2102000213330202-3111001211321100-2231023022330132-0212221131230210-2301200203120002"></a>

## Direct properties — automatic_from_start / 001313001201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130012310113320-3222100222122023-0230023221230200-3223301310002223-0132123130122331-0223233303100112-0330321113022310-0102021202000020"></a>

## Next pages — automatic_from_start / 001313001201 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1022123003030331-3201313021021123-2101330300031322-1130012233110300-3222232033101221-1011033302031103-0122030222313210-2302032121022312"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks — dhcp_networks / 203200121333 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks

<a id="canonical-1223112220313101-2111233031101113-2211133322201300-2232101020020200-2323000233102023-2003222001032103-0321003300100302-2331102033332003"></a>

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

<a id="canonical-0012303213032301-0121131020333131-0220303320332000-2022111220003330-1322003202201322-1332113313332202-1322131332131310-3312310003003233"></a>

## Direct properties — dhcp_networks / 203200121333 / 3

<a id="canonical-0103101123330203-0302223011123122-1303130201313032-1301212310100302-2200120101221130-1320000232333311-1312020200221223-3323003201110322"></a>

<a id="canonical-3302021203203212-0310100312121212-1101330121302131-0033132321131120-2120000203200111-2003202123122220-3211003010010313-3300203132303200"></a>

## dgw_address property — dhcp_networks / 203200121333 / 4

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

<a id="canonical-1022203333203032-3301313323032322-1021032313112230-1313201111111110-3110201030003002-3333203120232310-1301121011010213-2102221011132301"></a>

<a id="canonical-2333130000312120-3321023223031221-0320210300331222-1203232302323122-2120233100230103-2232330003100123-2302220203021122-0303213302310032"></a>

## dns_address property — dhcp_networks / 203200121333 / 5

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

- [first_address](resources--voltstack_site--reference--group-008.md#canonical-3112203201013001-2113301033312321-3101032111312323-1010330222100313-1022030031322223-2302111023211111-0030112123010203-3230113332013300): complete subsection reference.

- [last_address](resources--voltstack_site--reference--group-008.md#canonical-2233223031020131-2230112131313220-2213023222122123-3012032302333322-2113310113013100-0213012000312112-2002120032232002-1130121200303002): complete subsection reference.

<a id="canonical-3003222002202013-0131211221102300-1331021313032003-3221100303122021-3210330323113213-3033301102221310-3111323022333321-0203103202322033"></a>

<a id="canonical-1202131310101323-0211200111213203-1130113312103030-1001033100313213-3203011022131023-1320120233200301-2031203302122101-2300023323331203"></a>

## network_prefix property — dhcp_networks / 203200121333 / 6

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

<a id="canonical-1123323101121310-0300320231203133-2211002230322121-2310222103200033-2103122130310201-3032001311021110-3131013303031321-2021000021103233"></a>

<a id="canonical-2301131100033213-0302022311330002-0310101310223321-2002113202230101-2210100000111211-3212101120221020-0100003333123213-0301101200112001"></a>

## pool_settings property — dhcp_networks / 203200121333 / 7

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

- [pools](resources--voltstack_site--reference--group-008.md#canonical-2131311123020011-1111112000122123-0102123221013013-1312123301111322-1131010111311202-1330002322222220-1231132332031212-0110230111320312): complete subsection reference.

- [same_as_dgw](resources--voltstack_site--reference--group-008.md#canonical-3203121112113112-0031103102031022-0221303211212331-0010311113011102-1210112000201323-3200203112212112-1102233031022010-2300220311130130): complete subsection reference.

<a id="canonical-0031010300021303-0122000121003030-2211330122103200-2321010123211103-1132121130203102-0312022202132003-2201230001121231-1102312201131300"></a>

## Next pages — dhcp_networks / 203200121333 / 8

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address](resources--voltstack_site--reference--group-008.md#canonical-3112203201013001-2113301033312321-3101032111312323-1010330222100313-1022030031322223-2302111023211111-0030112123010203-3230113332013300)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address](resources--voltstack_site--reference--group-008.md#canonical-2233223031020131-2230112131313220-2213023222122123-3012032302333322-2113310113013100-0213012000312112-2002120032232002-1130121200303002)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools](resources--voltstack_site--reference--group-008.md#canonical-2131311123020011-1111112000122123-0102123221013013-1312123301111322-1131010111311202-1330002322222220-1231132332031212-0110230111320312)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw](resources--voltstack_site--reference--group-008.md#canonical-3203121112113112-0031103102031022-0221303211212331-0010311113011102-1210112000201323-3200203112212112-1102233031022010-2300220311130130)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3112203201013001-2113301033312321-3101032111312323-1010330222100313-1022030031322223-2302111023211111-0030112123010203-3230113332013300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1012101311231010-1123000230113001-3001120212010101-0122302101300021-1012121032231132-1023302202233300-0122200120330222-0001222303112313"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address — first_address / 030232330323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.first_address

<a id="canonical-2312200120203032-0310220111223310-0200201013113103-0213332010120220-2201032332003103-2100120233103332-0001100011102012-1320323311330303"></a>

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

<a id="canonical-2212000000000332-2323310332302131-2320011323203102-1111231021210302-2101201021020003-0130100011001112-0302212001303231-3122202302300110"></a>

## Direct properties — first_address / 030232330323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112331312120301-1222111122122100-3121003123332232-3121012300203300-3013302312201011-3323112333220122-2110303300222222-3231323132202021"></a>

## Next pages — first_address / 030232330323 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2233223031020131-2230112131313220-2213023222122123-3012032302333322-2113310113013100-0213012000312112-2002120032232002-1130121200303002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000003130310013-0322021210233010-1331231101333131-0113332330013102-3031023322103320-2330023013323302-3101132132323122-1213120010321010"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address — last_address / 101101130223 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.last_address

<a id="canonical-0013320223230210-0323221111221033-0213211313133332-3212000132012231-3320000002022132-1233202310213030-0130112210022023-1032312111022130"></a>

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

<a id="canonical-0011200101001003-1120122031232013-0201230330121002-3033111220031120-2323011023110013-1333203101121131-1332211300323310-1230210312221100"></a>

## Direct properties — last_address / 101101130223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2030222003332331-1112220121012102-2321230202011201-1211121222203313-0032010110310021-0312113133233320-3021000123001332-2300032220121130"></a>

## Next pages — last_address / 101101130223 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2131311123020011-1111112000122123-0102123221013013-1312123301111322-1131010111311202-1330002322222220-1231132332031212-0110230111320312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0113323113232320-3123013230012003-0211022100001131-3120213111110233-2023222111113003-0113211020100220-1230232113110013-0132201232311313"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools — pools / 022223320000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.pools

<a id="canonical-1211301310101221-0000022122212000-3220333102102313-3003302110133002-2331002120312233-3210212211013323-1210000133322202-3212323111202011"></a>

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

<a id="canonical-2102200332031312-1021320301011022-0021203032132121-3201122110202322-1100311031302203-2200123230212301-0203033020110230-2113011001120033"></a>

## Direct properties — pools / 022223320000 / 3

<a id="canonical-0201111113330033-3110130111032302-0330020221003122-3002331321333210-0203230103232121-1232031001010110-1200330130223120-3211230203021033"></a>

<a id="canonical-1111233313231301-3333032002311321-2232330110112303-3221223120232203-3233002232202313-2321003331132321-2023002320131013-2131212210331121"></a>

## end_ip property — pools / 022223320000 / 4

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Upstream description:

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

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

<a id="canonical-3311132311220012-0231102020223333-1102331313033312-1322003230302221-3013233011000303-0031210001003203-1202001231303011-1322302012220321"></a>

<a id="canonical-2200313033200110-3033133222322323-2233120020022310-2300010030233330-3033020111233021-0111311321011230-1202111311131232-2302012202312112"></a>

## exclude property — pools / 022223320000 / 5

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-1113210222010010-0333222112221332-2130002301121302-1123303312330222-1212322203101113-3020203123033312-1223023020222102-3102321022010303"></a>

<a id="canonical-0231010022232311-1122311100301113-3110013313101110-3112313121031200-2132022020333233-1121001021303201-2023221321223332-2132011022030232"></a>

## start_ip property — pools / 022223320000 / 6

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Upstream description:

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

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

<a id="canonical-1033203212110123-3123223120121121-1321001311200022-2020121213213203-0323000113332003-0033232103312311-2023301102132203-0200023100110323"></a>

## Next pages — pools / 022223320000 / 7

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3203121112113112-0031103102031022-0221303211212331-0010311113011102-1210112000201323-3200203112212112-1102233031022010-2300220311130130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030310232123213-3310321303101033-2231010311323012-1020030133113031-1331000203012301-1200002001010231-3000331213112210-2211110233212333"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw — same_as_dgw / 233131122132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1033201023211112-1312120232310003-2303313330213311-1021111322032220-3323031322112320-0322322012120012-3110330310231200-3002110003212301"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
same_as_dgw = {}
```

<a id="canonical-1012220303223313-1013133111001212-2231222313233030-2030301103311212-1001000103321021-2103012102203120-1012222110202013-3120233100022102"></a>

## Direct properties — same_as_dgw / 233131122132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223221232131010-3313113211200032-0000123323001300-2021223312101000-0203031002202002-2200120321212332-1121120223200203-3133231120321112"></a>

## Next pages — same_as_dgw / 233131122132 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-2033012133202233-3112013110213100-2122213231330010-3133020103311121-3121332210011201-2010003222112031-1120020020013021-2220231332111311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1110000210131331-0222211003000203-2323232230323312-0332322023200323-0300013122103330-0031301111323001-0120123303121031-3131001303100232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123211110221210-3023212112010232-3330132212031010-2301300222100231-3310210012013222-2202220203212110-3131312221203130-2130012020221020"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map — interface_ip_map / 000322223113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server.interface_ip_map

<a id="canonical-0303123213233320-2100213331303022-3211320201031323-0002220100302130-3221301132230332-1301101011022020-0210111112002101-2303221101121102"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
interface_ip_map {
  # Configure direct properties listed below.
}
```

<a id="canonical-1202211231133321-0221232210221101-1010212013311113-1110200201100021-1223233310201330-3311102002020223-1313013332222103-3100303321100311"></a>

## Direct properties — interface_ip_map / 000322223113 / 3

<a id="canonical-0133213022203202-0313002313101100-3100210321103100-0121310311010303-2111000122302211-0123013302233303-3300032311212020-3302011111103212"></a>

<a id="canonical-3030112110232113-3301320032011003-3003232013003122-1223012011032003-1111130103332310-2002020112133102-3220031101101000-0332202030322200"></a>

## interface_ip_map property — interface_ip_map / 000322223113 / 4

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 64,
    "metadata": {
      "confidence": 0.75,
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

<a id="canonical-0333332122330311-3030212322030133-2333211033123130-2300032322311222-0032030003110033-2110213133011100-2020000331101331-1202331113311231"></a>

## Next pages — interface_ip_map / 000322223113 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.dhcp_server](resources--voltstack_site--reference--group-008.md#canonical-1002031101222310-1031120003212032-0212122012010133-2223101011301020-0302130312231203-1220231230332100-2332202222033222-2303311021132221)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0133102302120033-0222020110201012-2230023311030010-0021302213311321-1123322000013112-0001111113021302-2320332300201012-1310102222321213"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config — ipv6_auto_config / 232133030022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config

<a id="canonical-1103120100111133-2302031201012021-2331302110320212-2101020332221203-2212311021102321-2321111213323210-0003211110231012-3222101122011233"></a>

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

<a id="canonical-0100101330221131-3112220123000333-2332000212223333-0322311311013232-0210321122133100-0333021222020102-1103230200112001-1310232030332301"></a>

## Direct properties — ipv6_auto_config / 232133030022 / 3

- [host](resources--voltstack_site--reference--group-008.md#canonical-1313301312230213-2202223313111001-0012332202023231-0322031132010012-1310133030000112-2030231030000313-1203211001201330-2221012231123330): complete subsection reference.

- [router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311): complete subsection reference.

<a id="canonical-3223233100120112-1121122331112101-3113110232011303-1313111210123203-1333031113221323-0330211132121230-2301220233112012-0123011203222010"></a>

## Next pages — ipv6_auto_config / 232133030022 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.host](resources--voltstack_site--reference--group-008.md#canonical-1313301312230213-2202223313111001-0012332202023231-0322031132010012-1310133030000112-2030231030000313-1203211001201330-2221012231123330)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1313301312230213-2202223313111001-0012332202023231-0322031132010012-1310133030000112-2030231030000313-1203211001201330-2221012231123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000021130213003-0010023321003103-0301222101001022-2002122201302113-1202110113200133-3030232012330031-0310033003231001-2213120012103013"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.host — host / 110132111130 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.host

<a id="canonical-0211122210302310-2333203201121022-2321033203020310-1122330021023032-2103212121033100-3200013312230211-2030213322022212-0120023332000002"></a>

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

<a id="canonical-1202131301330302-0311111233331132-2303033121022303-1312222213021210-2333211003212030-0133320112231032-3100203023321003-0313221203011130"></a>

## Direct properties — host / 110132111130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011120020212220-1231011121022011-3231012213001301-1303221330200002-3130320111102210-3113130200333232-3120203211130330-0102122300121231"></a>

## Next pages — host / 110132111130 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130123133011100-3100033021131321-1323112322113303-0300033233222103-0200200102213003-1023101232123313-1032310010130123-2330030313113032"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router — router / 211230322300 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router

<a id="canonical-0310022320020111-0111131210203333-3233221130231021-2121102113133032-2001200211300032-3020313300001012-0110300130330113-1230330022121223"></a>

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

<a id="canonical-0000023332133123-2232231001331131-2231131313231321-0201111021332121-2200230000131302-1332013320222123-1333102111333000-0010123322232020"></a>

## Direct properties — router / 211230322300 / 3

- [dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321): complete subsection reference.

<a id="canonical-0000312120102303-1121133302321123-3313002323032131-1223021010001233-1120112233100032-0102201003031311-2022133011122311-1213020013011312"></a>

<a id="canonical-2133101222113013-3322211332203210-1231133031322211-3312233032201202-1313023031100213-1221302030322112-1313312202022023-2102130203002212"></a>

## network_prefix property — router / 211230322300 / 4

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

- [stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112): complete subsection reference.

<a id="canonical-2211211322011202-0000201331320222-3212233123132130-0023032113300111-0001102131022102-3310302201320020-2312220331203022-0213112033223102"></a>

## Next pages — router / 211230322300 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210012302001321-0112102122330322-3100000321110130-2202031131303022-0301111122331230-2011333023303003-1012033232102111-1233201223032233"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config — dns_config / 332121302230 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config

<a id="canonical-1003123113203331-2112311302323232-1021221013000131-2101321100200201-1110032312002003-2012131032133232-1202320220020111-3311031132100210"></a>

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

<a id="canonical-2123212310310131-0331232321033021-2033031110231330-2231130001322201-3132032002133220-0212013313212333-3301322303103231-3230203111232103"></a>

## Direct properties — dns_config / 332121302230 / 3

- [configured_list](resources--voltstack_site--reference--group-008.md#canonical-0122011300022330-0323223131302312-3313101011332221-1100210001002113-3013310330120132-1321313220230233-2130301020300130-2300222030301110): complete subsection reference.

- [local_dns](resources--voltstack_site--reference--group-008.md#canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330): complete subsection reference.

<a id="canonical-0130330031233113-0233232101232300-3023110000321000-1331200203131122-1032011200212222-2330212312232032-3120020210122012-0332021301112231"></a>

## Next pages — dns_config / 332121302230 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list](resources--voltstack_site--reference--group-008.md#canonical-0122011300022330-0323223131302312-3313101011332221-1100210001002113-3013310330120132-1321313220230233-2130301020300130-2300222030301110)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-008.md#canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0122011300022330-0323223131302312-3313101011332221-1100210001002113-3013310330120132-1321313220230233-2130301020300130-2300222030301110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330301130031230-2231032002100202-0221021230210303-0000131302000112-1201333330321111-2221021320223311-3230331320312320-0022132023212201"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list — configured_list / 212220021211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-1312302223133020-3023330323121222-2131210330303103-0101011321121211-1133003323013002-0002003012102122-0100332133232002-3003103100331022"></a>

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

<a id="canonical-3320111320113032-2320301310321301-0301000123200312-1103030132322110-0211011210321321-0312230303321232-3010111023300120-3123203213332323"></a>

## Direct properties — configured_list / 212220021211 / 3

<a id="canonical-3112303330030100-3122020312303122-2312021110030223-3131303122022230-0121332130033131-3011102011311012-2320112011013131-3100320101113001"></a>

<a id="canonical-1310330313212313-0123000112212311-1011310210113221-3003013033300210-0223333303333213-0123113233121111-3021121331122132-0213303102001021"></a>

## dns_list property — configured_list / 212220021211 / 4

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

<a id="canonical-0211121121001232-1200311003201103-0023221120333133-0003212211100120-0300102003032033-3301112103200331-3223002230013012-3301333000211202"></a>

## Next pages — configured_list / 212220021211 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213231303233031-0200122300103103-3001033111023310-1013102002311310-0000033220233331-1212022101130203-2323213222311012-2103213301122223"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns — local_dns / 211220210000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-1202120112131112-2312310312130331-1323201132232030-3303203000120333-2232232202133030-1031323100122021-0332203100213232-2121333102111323"></a>

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

<a id="canonical-3203322001223120-0111323311130111-2201311121001102-3112323023201312-0012123312001131-3030121231031330-3221231020213012-1021312210030122"></a>

## Direct properties — local_dns / 211220210000 / 3

<a id="canonical-2120330113222223-1003212130312232-3230301131232303-1313233121213030-0200120001013300-3023031002103301-0021130003332303-1331232312213222"></a>

<a id="canonical-3210230200303330-1311110202113230-3111012033012300-1120103233003220-1301103321322321-3021312132013002-2031230322213010-1233023110111033"></a>

## configured_address property — local_dns / 211220210000 / 4

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

- [first_address](resources--voltstack_site--reference--group-008.md#canonical-1233013020020332-3213023322323130-1322222021321323-0320320323110133-2303121122201331-2331322221210203-2220213102000023-0022112333313323): complete subsection reference.

- [last_address](resources--voltstack_site--reference--group-008.md#canonical-3211011033023021-1111333031301233-1230112103022132-1213021121101222-2220111202012022-3310321121003101-3312212313333333-3033020320202020): complete subsection reference.

<a id="canonical-2022313120030221-0012230030222223-0303331111221311-0012313003132211-2233113122231302-1031331310020212-0012001032122120-0011031210223010"></a>

## Next pages — local_dns / 211220210000 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.first_address](resources--voltstack_site--reference--group-008.md#canonical-1233013020020332-3213023322323130-1322222021321323-0320320323110133-2303121122201331-2331322221210203-2220213102000023-0022112333313323)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.last_address](resources--voltstack_site--reference--group-008.md#canonical-3211011033023021-1111333031301233-1230112103022132-1213021121101222-2220111202012022-3310321121003101-3312212313333333-3033020320202020)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1233013020020332-3213023322323130-1322222021321323-0320320323110133-2303121122201331-2331322221210203-2220213102000023-0022112333313323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202000121102201-1131011333211013-0033213222232021-1013322013332031-2123002022221130-1321323211303223-2220201000121203-2300031132023311"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.first_address — first_address / 231013013200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-008.md#canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-3220303100303233-2210202102313210-1100221001133330-2301322031230302-0020022313320313-1213320332223221-1012233030210102-0231331311013100"></a>

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

<a id="canonical-2332232210032032-0023131210000001-2033310320300311-1113020203221213-1311202002322222-0010102201130000-0100111132220000-1210001330000312"></a>

## Direct properties — first_address / 231013013200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323033200020130-1001230211231222-0001120212131203-2111223012132223-1332033310031332-1133130310310221-0021002000222022-3000302221020023"></a>

## Next pages — first_address / 231013013200 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-008.md#canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3211011033023021-1111333031301233-1230112103022132-1213021121101222-2220111202012022-3310321121003101-3312212313333333-3033020320202020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322332030112102-2312110023320100-0133013103313100-0031133301031323-1121321323010321-3122011113132030-2121211333211022-3303223111122122"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.last_address — last_address / 301122311211 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config](resources--voltstack_site--reference--group-008.md#canonical-3020131103330110-3222120310131012-0030132031023220-0122020121123013-1122232330233131-1313322303110123-1312302122311003-1132111302111321)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-008.md#canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-3100133120002021-2233012202003331-3100320230300323-3213020023203013-0312023231311032-0203211021122132-1331021003200210-0123100022130011"></a>

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

<a id="canonical-1001012333113220-2130020212111023-1332113213231111-3213333230200313-1200123110220332-1020033313102122-0300012233002323-3132222201211110"></a>

## Direct properties — last_address / 301122311211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323322332301021-3102213323323133-2021311130310301-2210323003321012-0201001001230212-3331022212211320-3002302101033230-3001101313211333"></a>

## Next pages — last_address / 301122311211 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.dns_config.local_dns](resources--voltstack_site--reference--group-008.md#canonical-2220201202031312-2000332220302030-3133331033133331-2030031011333122-0002003112221330-0233233233113020-1002033101010210-0323222131203330)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003031301122011-3113001132110310-3033122132322312-1100001332123002-2120312033323302-0310223020131110-1321212012203003-0300102202330201"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful — stateful / 221011201032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful

<a id="canonical-2333122022202013-2210103111020032-2100010002230022-3322101012303323-0012202003123031-2113311211121202-0033233012221231-2232333221002122"></a>

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

<a id="canonical-3233232332033201-1311120022020211-3003211232002130-1231332130032031-1001011222220220-1321002020310101-2300020231003133-3132201023212302"></a>

## Direct properties — stateful / 221011201032 / 3

- [automatic_from_end](resources--voltstack_site--reference--group-008.md#canonical-3313232031123012-1132201003330311-2110232120212010-0121010000232213-0333300030033123-2232010021213322-2030210233002101-1320231221223013): complete subsection reference.

- [automatic_from_start](resources--voltstack_site--reference--group-008.md#canonical-2130303102203012-3221331022002230-1313303012031133-1220100220012231-3032011331131021-2020230333130113-3001333320031130-3133320023222231): complete subsection reference.

- [dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-1123313111200233-1012212101211132-3223211330300011-0003200232322100-1332313123003213-3031310113131133-2021133120313331-1301010300233021): complete subsection reference.

<a id="canonical-1122200132211011-1321103322100101-1030222123120213-0032202323102230-2203102220303101-0220332010233001-3203113210210003-3330213132321220"></a>

<a id="canonical-1213202022013221-3101322231330013-0011212323013103-1200312321320331-3303211121222300-0202222000311223-3332223211201023-3111223223313011"></a>

## fixed_ip_map property — stateful / 221011201032 / 4

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Upstream description:

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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

- [interface_ip_map](resources--voltstack_site--reference--group-008.md#canonical-0300131331222301-0023230020222313-1211033000132120-2133323030132213-3013031230201102-3200211120020023-2001220032032113-1013031020213201): complete subsection reference.

<a id="canonical-0011220013321022-2100111221013330-3301233303103122-0132210300303320-1321111031121123-0210013320020132-3310222102232002-0222222333130010"></a>

## Next pages — stateful / 221011201032 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end](resources--voltstack_site--reference--group-008.md#canonical-3313232031123012-1132201003330311-2110232120212010-0121010000232213-0333300030033123-2232010021213322-2030210233002101-1320231221223013)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start](resources--voltstack_site--reference--group-008.md#canonical-2130303102203012-3221331022002230-1313303012031133-1220100220012231-3032011331131021-2020230333130113-3001333320031130-3133320023222231)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-1123313111200233-1012212101211132-3223211330300011-0003200232322100-1332313123003213-3031310113131133-2021133120313331-1301010300233021)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map](resources--voltstack_site--reference--group-008.md#canonical-0300131331222301-0023230020222313-1211033000132120-2133323030132213-3013031230201102-3200211120020023-2001220032032113-1013031020213201)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3313232031123012-1132201003330311-2110232120212010-0121010000232213-0333300030033123-2232010021213322-2030210233002101-1320231221223013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213312123323220-1311130323021320-2003311211310201-0331310120222320-0133002102331121-2220323121023121-2130013302121020-1021301121311133"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end — automatic_from_end / 013202002312 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-0113232130321102-0101102323022103-2120231110231002-2030130322310323-2203003313203320-2121203331023001-1233102112130213-1000120222200022"></a>

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

<a id="canonical-1003020100200010-0302012101113012-1320311303230330-1013122012123100-0313112200203002-3333210133232330-3231031230231122-2022303332032021"></a>

## Direct properties — automatic_from_end / 013202002312 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320103321121220-3330332201230313-2313012312120303-3301021220030302-2333102132202131-0003132102100021-2010233210110023-3331222122320101"></a>

## Next pages — automatic_from_end / 013202002312 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2130303102203012-3221331022002230-1313303012031133-1220100220012231-3032011331131021-2020230333130113-3001333320031130-3133320023222231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232012133221013-2202303010320100-1010000222100220-2221023003300010-1202111122112331-1120310220031003-3220232131121113-2213032221302132"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start — automatic_from_start / 233133112121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-1023032211002101-2233300212230000-2220313312232313-2312322010002332-3303123102232010-1322313121120320-1332312230201201-0312031312302220"></a>

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

<a id="canonical-2110022320000001-2200112230103222-3002330123020013-1033202202030003-1322000303003220-3322312022111313-2123231301331122-1300132222031000"></a>

## Direct properties — automatic_from_start / 233133112121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0301221311012112-3001312022301321-2220332000023032-3223202331313032-0323022120200120-2002001230130111-3000130232222212-1301022012102202"></a>

## Next pages — automatic_from_start / 233133112121 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1123313111200233-1012212101211132-3223211330300011-0003200232322100-1332313123003213-3031310113131133-2021133120313331-1301010300233021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222021222120321-2103020330320313-1202321310123331-2000122331023221-2110331013020022-3201220231312021-2201100032113103-3223000122311003"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks — dhcp_networks / 002121332200 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0002331332220321-2233313311001031-2033213133313011-2001022202303323-0210112212033330-3302302113212311-2312002321323112-2121202212011030"></a>

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

<a id="canonical-3002320011123102-0331200301120133-2103022300131221-1123211110201003-0201320120313131-3232211201002330-1001230100302020-2010110310103133"></a>

## Direct properties — dhcp_networks / 002121332200 / 3

<a id="canonical-1102002203212130-2211033110211101-1230220220130311-0113312322222231-1102210002121110-0012220033310233-1311012231312220-0200032010213213"></a>

<a id="canonical-0320302000013322-1332132110313322-2202111313013312-1202022223101112-3112200200120013-3202311303000311-3102131123102010-1120301220001012"></a>

## network_prefix property — dhcp_networks / 002121332200 / 4

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

<a id="canonical-1101330020122230-3303310311032303-1011123232232003-3212120121132203-0111012001012121-2232033311112310-1120331032231023-1300000221113321"></a>

<a id="canonical-3120123300302230-2301020213321111-2212031103330010-3331200111102132-2230201301010101-2003021121231230-3001031003231012-3200131331033103"></a>

## pool_settings property — dhcp_networks / 002121332200 / 5

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

- [pools](resources--voltstack_site--reference--group-008.md#canonical-0121013132130031-0133030313110121-3101203120231011-0310300302201333-0322312030322112-0221010011223212-1230331200022131-2213121302331213): complete subsection reference.

<a id="canonical-0111132312231322-0221303110311103-2122310100300121-2032032222121003-1201332303230022-1313101101303012-2212211320003033-2011100020112120"></a>

## Next pages — dhcp_networks / 002121332200 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools](resources--voltstack_site--reference--group-008.md#canonical-0121013132130031-0133030313110121-3101203120231011-0310300302201333-0322312030322112-0221010011223212-1230331200022131-2213121302331213)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0121013132130031-0133030313110121-3101203120231011-0310300302201333-0322312030322112-0221010011223212-1230331200022131-2213121302331213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2012103013320013-1020121312332212-2131333201123010-0200100300133130-0220313202312313-1220000101232321-0121321302103103-2301011003021002"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools — pools / 203032032121 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-1123313111200233-1012212101211132-3223211330300011-0003200232322100-1332313123003213-3031310113131133-2021133120313331-1301010300233021)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-3311221221330313-3031200320023122-0023203100220010-3221231020222230-2332221301321032-2331113203032220-0313210023300030-2033332020102031"></a>

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

<a id="canonical-2233021311023313-2230200202000202-3220231013021221-1130230312032003-3132131331300030-3022322321013302-3211221131200212-3303202111103120"></a>

## Direct properties — pools / 203032032121 / 3

<a id="canonical-1313032003233332-3101222223122230-2022020001202031-1332332223032003-2020200233020022-3023220320100210-3220101221230032-2123311330033113"></a>

<a id="canonical-1001331213111133-0000122232121102-1001231322330302-1103200220222203-2221202023333303-2333313031213122-0233021201110120-1233311012310112"></a>

## end_ip property — pools / 203032032121 / 4

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

<a id="canonical-0121132020020312-3321203103131123-2120103331220220-0020302220123101-3201323131321011-3132130123010331-1212223201031323-0130300022001321"></a>

<a id="canonical-0100323210303022-2120312130101031-3011031203330210-0322311022011112-1231112333031320-0000313203103131-1012321222313310-1131001021211220"></a>

## start_ip property — pools / 203032032121 / 5

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

<a id="canonical-0201213232322300-0102320312101020-3203023133200131-1002211210212323-0001323303301030-2012323302110021-1211213023311123-3001223100201011"></a>

## Next pages — pools / 203032032121 / 6

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.dhcp_networks](resources--voltstack_site--reference--group-008.md#canonical-1123313111200233-1012212101211132-3223211330300011-0003200232322100-1332313123003213-3031310113131133-2021133120313331-1301010300233021)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0300131331222301-0023230020222313-1211033000132120-2133323030132213-3013031230201102-3200211120020023-2001220032032113-1013031020213201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013120132200120-1121002203232030-0221023103212131-3021202031033321-3233002013213103-2030011113312033-1320202100022232-1213030330213320"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map — interface_ip_map / 000311033110 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config](resources--voltstack_site--reference--group-008.md#canonical-3231001030200322-0202323001311330-3033121122313122-3210022310223110-0011100033121213-1213302133103000-1310313210312031-3310233122222131)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router](resources--voltstack_site--reference--group-008.md#canonical-3103211122123221-2233021131001030-2232022301131223-2111322233203122-0210121212111031-3222310200211120-3112103001202222-2232010113200311)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-1213100031112222-1303120102023123-1223331120001310-3201103031122322-1100011033020110-3311100332202331-2000112303121021-1021233013031013"></a>

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

<a id="canonical-3113122201131123-1302331000032103-2111100330130200-0001121201313132-1220012113210002-2310102121233130-3121212213202021-1212033110211203"></a>

## Direct properties — interface_ip_map / 000311033110 / 3

<a id="canonical-2122020100030331-0330102122121202-3323103331303133-1332201222233223-3120212302113223-1211002010213233-0203132203002311-0123003003120021"></a>

<a id="canonical-0212302111220012-2000332031101122-1132322212132020-0311212300311323-3220312212203310-0231003130023233-3133133301322002-3131302210230020"></a>

## interface_ip_map property — interface_ip_map / 000311033110 / 4

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Upstream description:

Map of Site:Node to IPv6 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 64,
    "metadata": {
      "confidence": 0.75,
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

<a id="canonical-2223303102332113-3220200013010113-2112202310021232-3030213333011020-0000020331101220-0211300222012202-1110321110030213-1103232310320332"></a>

## Next pages — interface_ip_map / 000311033110 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.ipv6_auto_config.router.stateful](resources--voltstack_site--reference--group-008.md#canonical-3300303013230331-2021231023101211-2011313323303332-2312312211221013-3223221103322122-0310201001202021-2320300200112230-2132311223023112)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3202030333101230-1121113030133310-2311000313322112-2002221331302311-1101320221233330-3032023102230211-3102102021002002-1121321030133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133021030201033-3012120301000213-3310200311212013-0012002320132020-3113121312030333-1302233202132210-1332302122110200-0022220200123032"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary — is_primary / 113231321131 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.is_primary

<a id="canonical-0000003333211021-3222022102110210-1310121222113302-2001112102202212-3032032333102311-2133332330232313-1002300030332321-1102303122312230"></a>

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
is_primary = {}
```

<a id="canonical-0210203321210113-3010112210200010-0331132133331020-2133321101220121-0203132222110303-1303333202002322-2012230323102202-0302031122331130"></a>

## Direct properties — is_primary / 113231321131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312111302231011-0021221332110102-1202122001200021-2311223100121132-2232013033113031-1300323320032030-1231231130230002-3233212230330210"></a>

## Next pages — is_primary / 113231321131 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1320122111233123-2330221102223112-0001121121332231-3313323333121211-0311231213011220-2232300220130003-0112332010112220-1011022021110201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210321030203233-3131330301032322-3120003200331323-3301310230030211-3122213303302320-2103212330210303-2331300101113110-3010231001023332"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor — monitor / 032321133212 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor

<a id="canonical-0132131012230323-3321201222110300-2212102001131213-3213313112033323-2300032323303302-1123123312322120-3030331021301320-1100233010022032"></a>

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

<a id="canonical-2000010320103001-2312121303230002-2230123100330132-0222311333123131-1131230033230211-0211130033010333-0132213223103022-3302111202003002"></a>

## Direct properties — monitor / 032321133212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320213110211000-3221110332003312-0000103201331320-3200113023023310-3013220120300111-2012223300002032-3333122011310011-2012113220302102"></a>

## Next pages — monitor / 032321133212 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1310211020301210-1210320103002203-0123310223331313-2220310221111122-0210321011011223-3133001110021331-3121300300032001-2133311321311032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1211323302021100-3122012130231303-0122123232203323-0013000321331000-1031102302311020-0030031120012322-3122113123103013-2303010112201302"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled — monitor_disabled / 102110332033 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.monitor_disabled

<a id="canonical-0303302101330222-0102320130103132-0012313320130033-1322312131333232-2322231232202202-0313232300222232-2302322112310110-0020223122132231"></a>

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

<a id="canonical-2103231113131323-2031231132232112-0011132201001000-2030120000323000-0212010130233012-2331103331002330-0003113302033300-1020212311111223"></a>

## Direct properties — monitor_disabled / 102110332033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3231112201200100-3302212311223231-0201333133223022-2321321232021203-1022321210320333-3331212303031311-0300313132320323-1233030002331121"></a>

## Next pages — monitor_disabled / 102110332033 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1323110312222000-1332122013222211-1331223002101001-1320212001201110-2312030130101030-1102233103312311-1103123101232222-1323103011123230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1000331300021310-3332321223311133-1321132303333233-3120132213221132-2223200330100201-3313030230002023-1133222020023312-3332302201203301"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address — no_ipv6_address / 112311203130 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.no_ipv6_address

<a id="canonical-1113031330322203-3032013203312331-0133203120300103-0113122312213000-2023123232201300-2203332023221210-1000230001032332-3031223322222031"></a>

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

<a id="canonical-2011313001311000-1231231102212222-2101031303202311-2032301231300311-1212312210123011-3212222030032303-0023031032122321-3003222031132210"></a>

## Direct properties — no_ipv6_address / 112311203130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132332122211232-2002032003020223-3122311230320300-2010223300322210-3001212202332310-2032110132133021-0332000000230232-2210121230103013"></a>

## Next pages — no_ipv6_address / 112311203130 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2030001302031213-3301323223031021-0010112321210112-0032333302131232-0312133110201110-0211130013223333-2302313001012231-2000112312301001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1110003302322113-2211303102333013-2101331231033022-3002232013303312-3012310031000111-3021210003200001-1303100233013211-0320331010133032"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary — not_primary / 002120300123 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.not_primary

<a id="canonical-0122011313023202-3110230112123131-3232330131103310-0312323030021013-0120023221110210-3220231111100032-1102311333322310-2231023113311302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
not_primary = {}
```

<a id="canonical-3023133211322331-0333032302211213-2223301013330013-0002223313223103-2322202303310130-1002230011313311-3221222033131030-2002123010311103"></a>

## Direct properties — not_primary / 002120300123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331020223312300-2233121111200103-1232130210113023-0130222020230022-2303113022101000-1212301021020022-1102002000000033-3002122322200111"></a>

## Next pages — not_primary / 002120300123 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2021132332003322-2303131302223130-0132230213011030-3311302120210010-3023311123330121-3303223001232121-2131133100013301-0011212303203201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0020110133132110-1131032010332112-1020213310010310-3313030032333223-1133133133131023-0121212100203201-0010201222330013-1221102012112030"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network — site_local_inside_network / 013230312332 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_inside_network

<a id="canonical-3012232020013303-3101120020100211-0233222022333311-2120200303330222-2300321031222113-3030233112301201-1230002110033013-3111232032021102"></a>

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

<a id="canonical-3320013212322210-0221010202230031-1122033322330213-2323103123201333-3120101000031002-3232100332313133-0333023123302000-3331133213212000"></a>

## Direct properties — site_local_inside_network / 013230312332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001031301002223-2201201330222013-1101320103312220-3222223030031310-0332131201133302-3113122031023132-2200002303033313-1332213331013321"></a>

## Next pages — site_local_inside_network / 013230312332 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2320002321301231-0320321323322121-2033113021100333-2201201302223131-1222101001333311-0230301320301111-2023232222010002-1100123212233000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122000223132032-2121323222000020-1213221003023313-2033223212312102-0321213332232031-3133001221330000-3020003203333121-0230011130031210"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network — site_local_network / 130201022122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.site_local_network

<a id="canonical-1212000300123200-2123011312130321-1001002220113313-2212113121202333-1313332023222123-1020201233322033-0303133233200201-1313121111100201"></a>

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

<a id="canonical-0220031111220020-1300121012013022-0220203210001221-1000301321132123-3121131133232012-1322110322202023-2102201231223221-0321331321231122"></a>

## Direct properties — site_local_network / 130201022122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0312102222331322-0232201213301001-3023012320112120-3031333302032102-1310131123313332-0222322203202030-0213101110220111-2003213232322022"></a>

## Next pages — site_local_network / 130201022122 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303320203321130-1222111022110212-2303303031122113-1030032120131103-1320003201200031-2200311023201222-1212012131132100-2032020102123233"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip — static_ip / 033330203100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip

<a id="canonical-1221000003112102-2100201212020011-1302202133000002-0333121201131121-1331322031210012-3303000001021331-2002313230210212-0222010030103123"></a>

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
static_ip {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010033311003200-3212310322120303-1322222313003002-1330123133122210-2323031323012102-1200323203312232-1032120133112220-1210213310102232"></a>

## Direct properties — static_ip / 033330203100 / 3

- [cluster_static_ip](resources--voltstack_site--reference--group-008.md#canonical-3310011330130321-1101330002000001-2002102232030320-1112201101230231-2011313322002022-1212322011112133-3131333102011230-1012332122233321): complete subsection reference.

- [node_static_ip](resources--voltstack_site--reference--group-008.md#canonical-2221013113030331-2130211221132001-1221323232301232-3020131121230033-1133112133220233-1030131112000232-0202211003112322-2032030203221122): complete subsection reference.

<a id="canonical-3102111233333201-0220020320231233-3100012113201210-1232131030202030-2112122000200323-1221232120132331-0210213303021220-1113121103323202"></a>

## Next pages — static_ip / 033330203100 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip](resources--voltstack_site--reference--group-008.md#canonical-3310011330130321-1101330002000001-2002102232030320-1112201101230231-2011313322002022-1212322011112133-3131333102011230-1012332122233321)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip](resources--voltstack_site--reference--group-008.md#canonical-2221013113030331-2130211221132001-1221323232301232-3020131121230033-1133112133220233-1030131112000232-0202211003112322-2032030203221122)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3310011330130321-1101330002000001-2002102232030320-1112201101230231-2011313322002022-1212322011112133-3131333102011230-1012332122233321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200001333202203-0131302020010303-1120112012330330-1121210212221321-1123030323300303-0330023133003101-2331021002232323-1311120312131212"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip — cluster_static_ip / 030310121120 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](resources--voltstack_site--reference--group-008.md#canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.cluster_static_ip

<a id="canonical-0123003020223303-3123232312212033-3312023112303300-0313131300010222-3133332013002323-1101112121210132-2101310330011211-2201023212123102"></a>

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

<a id="canonical-2031113322321033-0333021130323231-1002332310013322-0202033301013310-3120313222231302-3213233031230300-0030311231331120-1110210323110003"></a>

## Direct properties — cluster_static_ip / 030310121120 / 3

<a id="canonical-1120101211021330-1332210110030012-1321112133013013-2031112100133123-2100112220312231-1122210203300201-3110101120221102-3023230321210312"></a>

<a id="canonical-1102111310212230-1120200311322300-2030132202230313-3231012202033130-0023002233331331-1203113310131001-3203333130033223-2000300332010320"></a>

## interface_ip_map property — cluster_static_ip / 030310121120 / 4

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "object",
    "maxProperties": 128,
    "metadata": {
      "confidence": 0.75,
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

<a id="canonical-1331021313101100-2301111223211210-3202033302022023-1020110321003322-0133222302231223-1031023223022232-3332213012133322-2333222021230020"></a>

## Next pages — cluster_static_ip / 030310121120 / 5

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](resources--voltstack_site--reference--group-008.md#canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2221013113030331-2130211221132001-1221323232301232-3020131121230033-1133112133220233-1030131112000232-0202211003112322-2032030203221122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220332100023233-0123210013231133-1323021212000132-0111320200330311-1311021111313311-3011032330323111-0002110312200333-3103023332332223"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip — node_static_ip / 313301302323 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](resources--voltstack_site--reference--group-008.md#canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip.node_static_ip

<a id="canonical-2210131301032210-0010033312202232-2220120023331310-0310322330010130-0321100230231133-0032332011032121-3002303101013233-3132321331030330"></a>

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

<a id="canonical-3002300000101201-2330233132332120-3333021321223122-2302021220001132-0222132130231222-1301323213303311-1210213323102333-1113112221313212"></a>

## Direct properties — node_static_ip / 313301302323 / 3

<a id="canonical-3001310100030003-2302322123310230-0302033133220021-2213222201021123-1031323322010020-2302012211221103-1310000212133022-1000231101233312"></a>

<a id="canonical-0121023010211300-2003130032230121-2303030132132221-0320113000230233-1103102213001210-1122100313101100-3212231131121120-2210210300033030"></a>

## default_gw property — node_static_ip / 313301302323 / 4

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

<a id="canonical-2333313022022200-2312033011112210-1121103333300022-0313311023030120-0233031130321122-1231112031303010-1210311101121110-1301232032130201"></a>

<a id="canonical-1031001103030231-1333111132020232-1003133121233013-3102102331001120-0020010232222113-2230203103231101-2100012221112011-2220323331002222"></a>

## dns_server property — node_static_ip / 313301302323 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3311200320013003-1312030013032132-0210221113323131-1233330010110122-1132321311130320-2131331230023231-0120321010312222-0100203102200200"></a>

<a id="canonical-0122111011022310-1113323230321110-1021222131101232-1021013012200010-1003023033232222-3003300113011003-0001000011132033-3011233100232013"></a>

## ip_address property — node_static_ip / 313301302323 / 6

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

<a id="canonical-1202033123333311-1213103330222322-0331203132020010-2300331310231333-2222111023321202-0122332301230332-0313211211321213-2022232201200232"></a>

## Next pages — node_static_ip / 313301302323 / 7

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ip](resources--voltstack_site--reference--group-008.md#canonical-2133033101001023-2323203002122330-3133202203232132-3231000121131303-1133122131312020-0023322223023302-0103101130321200-2302221223221311)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3200300220203231-3023330333022333-3222222223230021-2200301322031220-2002330223023112-3221031232322100-2223212103201111-3333020321210313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031112033331221-2301020103202221-0310110311000220-3101010202231303-3301301032102221-2220123311101032-2132032311232320-1130131122211033"></a>

## custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address — static_ipv6_address / 111131320122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [custom_storage_config](resources--voltstack_site--reference--group-005.md#canonical-1130033323232323-2312320032102120-2230002101333233-0122122333223330-0120123302300302-0320200333230230-3313321101003303-0102021002223303)
- [custom_storage_config.storage_interface_list](resources--voltstack_site--reference--group-008.md#canonical-2300201311112121-1233300031310103-1300331123012212-1130323002022023-2200132010222313-3000120110302110-3111023121122131-2203013232210333)
- [custom_storage_config.storage_interface_list.storage_interfaces](resources--voltstack_site--reference--group-008.md#canonical-2333022013310011-2321020123302031-0001020312122322-0021113210112322-0210331200202312-0021003313201113-0330231110200303-0332123100223133)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address

<a id="canonical-3012033303102132-1123332021013333-0300131333020131-2103101001312230-3103230120202021-3213032010221023-2122001130312332-1221200033130330"></a>

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

<a id="canonical-1133120101223003-0302301213132200-3103003120000332-2120121231330122-1321001030203211-2202210300130210-2233023003223302-0022101112302033"></a>

## Direct properties — static_ipv6_address / 111131320122 / 3

- [cluster_static_ip](resources--voltstack_site--reference--group-008.md#canonical-0013212312003003-1102020323021121-3221022332032133-3021322221230222-3032203033112333-1000311203201203-3030013000113211-2103131120112233): complete subsection reference.

- [node_static_ip](resources--voltstack_site--reference--group-009.md#canonical-0312001133021311-0320010220010200-1322020010323302-2013203211021333-0303100302013312-2133321231003233-1011013013113312-3200122310012002): complete subsection reference.

<a id="canonical-2330213132013000-2202011310133122-1132310100231121-1010001130220300-1020200231001211-0101032130010332-0201123113121222-3331112230010020"></a>

## Next pages — static_ipv6_address / 111131320122 / 4

- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.cluster_static_ip](resources--voltstack_site--reference--group-008.md#canonical-0013212312003003-1102020323021121-3221022332032133-3021322221230222-3032203033112333-1000311203201203-3030013000113211-2103131120112233)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface.static_ipv6_address.node_static_ip](resources--voltstack_site--reference--group-009.md#canonical-0312001133021311-0320010220010200-1322020010323302-2013203211021333-0303100302013312-2133321231003233-1011013013113312-3200122310012002)
- [custom_storage_config.storage_interface_list.storage_interfaces.storage_interface](resources--voltstack_site--reference--group-008.md#canonical-1023113311011113-2300230312003233-1302020010010123-1111012013103323-1333021132331021-3033122102112331-1021000112213101-0230301111133001)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0013212312003003-1102020323021121-3221022332032133-3021322221230222-3032203033112333-1000311203201203-3030013000113211-2103131120112233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
