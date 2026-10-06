---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-0212200312302111-3110000001302021-2232303331320213-3022320103221030-0303011021330300-2112130011021031-3012021003333110-2311212213113212"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.devices` property

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

- [lacp](resources--securemesh_site_v2--reference--group-012.md#canonical-2202133101000121-1033000230102033-2013100322323133-2013101102103103-0212203013023001-1300133001031230-1030131212302111-2130322331331120): complete subsection reference.

<a id="canonical-2310103201222112-2001212110232333-1021302003002121-2030021233210021-0121220030201123-2222130212223023-3032331001310131-2211131013113023"></a>

<a id="canonical-2330211303310001-0230133020211100-3232320031023320-0213200302221312-2110301000201010-1100330123212320-3320013131131001-0100112321033320"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

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

<a id="canonical-2232332202030010-3200133112332303-3101330132122002-2301123330033331-2031032222121131-3113132212332301-2321013300130102-1310011332002120"></a>

<a id="canonical-3312130112112120-0030010010321323-2033312323222031-2221122321332203-3002221012230122-3003331011330030-1111033102020330-0121333112123111"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

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

<a id="canonical-0210133202120311-0322202310012331-2333100201220230-0331201100333033-2223320123130322-1123323310200230-2002210000033012-2311001030030122"></a>

<a id="canonical-3122032110321323-2013123312331330-1002213212322233-2021122021021000-3202132313000332-2310202131302011-0010322311103323-0122221022131230"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

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

<a id="canonical-3111202113131003-0213120221131113-1123323300322122-3121230102333222-0110133110311013-0230210303310310-1220030112120023-3012013333112330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- nutanix.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-2123310022220210-2311213321320030-2002031012322313-3103223231320013-2030303032130211-0023110331022333-0032210232230201-0001202332012131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for active backup.

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
active_backup = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2202133101000121-1033000230102033-2013100322323133-2013101102103103-0212203013023001-1300133001031230-1030131212302111-2130322331331120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-011.md#canonical-3230002221230310-1311132021112332-1111002221011100-2300132200220100-0320301210013002-2321300302231213-1122030030112302-3012100332302312)
- nutanix.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-0030210333102012-1311331330203212-0121330121302300-3211133122303003-2200331302223133-1030320001322310-3122113002020220-1301112332321131"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

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

<a id="canonical-3332120212010133-3302311220012003-3133123231022022-3103001031303333-1232130312311031-0222233033030132-3311032312131021-2200022002300023"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-2013021322033013-0330223203333131-1331223120311122-1003323100102200-2321000110100322-0301311323223103-3213103021202231-3112013102212211"></a>

#### `nutanix.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

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

<a id="canonical-0122301031222001-3012001030032031-1220230130013231-2113030122330033-3223230311003133-2330110130310303-1103123101201022-1002300311110013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-0211122002310120-0102301312113310-2200322113210212-1211313030030231-2212203231120102-1021033202323131-0303231200123123-3121102112032130"></a>

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
dhcp_client = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-0130003112301020-0220201333202102-0103312132301330-3001103332321202-3103021231203002-3131121111110232-3031131211113201-3110322312333113"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Additional upstream details:

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

<a id="canonical-2102121221013113-2323300111122012-3110022202102222-2101020233103320-3123311313122113-1103232232233123-3312301021222112-2132311120120202"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-012.md#canonical-0121011231312233-2213310221211130-1211000322220132-3333011011201332-2332321003102103-1223130233032031-0330222100332000-0020323222230230): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-012.md#canonical-0102023312223300-0033330111301323-0113231113102132-3232212102023131-0033000231001121-0312332033113330-1300201012313011-3232020302002331): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101): complete subsection reference.

<a id="canonical-0321330132101310-1002321323222223-0020203200133211-1023113030303220-2012221022322103-0000210000321330-0310230133313021-2211202032220231"></a>

<a id="canonical-0003012001313300-3302233120132111-2302023212012322-3033000101203101-1202121332003203-1012332003231022-1210101020203320-2021101101023220"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-2220033023202202-2122330031011220-0332131033221311-0021103021332013-0233200310033321-1111100330030330-2323100021212331-2221210010201333"></a>

<a id="canonical-1011023030221301-3310300211310310-1222033303203023-3013320111332003-1111100213022332-1010201130111023-2022200122112333-3001010133122122"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-012.md#canonical-1300230220001111-2120230133130003-0311321100332200-0220200211203310-2321030013223320-1123232033133010-0133312221201200-1233100210101201): complete subsection reference.

<a id="canonical-0121011231312233-2213310221211130-1211000322220132-3333011011201332-2332321003102103-1223130233032031-0330222100332000-0020323222230230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-3312101103122013-1310021312210132-0321122201132302-0020023103123221-1212322031300223-2033002203321033-3002220220222023-2031332233203030"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0102023312223300-0033330111301323-0113231113102132-3232212102023131-0033000231001121-0312332033113330-1300201012313011-3232020302002331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- nutanix.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-1331101030212330-2021203132010211-0121321001002130-0032233202102332-0031031013133323-3310200202120332-1321302201023023-2211330321230110"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-3002321220312121-3111323100233023-0303330312322213-1203232330100131-2302002121022012-0312000131323300-0211212003013110-0302303300123130"></a>

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

<a id="canonical-1110220221323010-3010311113132100-3123223213100320-1321000220311213-2321033332210110-0312123211133331-2332232123232013-2201023301202010"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-3001131132120322-2321233303020222-3310113303231001-0300311331020323-0130032330021000-3300121203212231-0201330202332230-2020332212313012"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Optional.

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

<a id="canonical-2232220122131113-0310313320030002-3013210303330133-3023230311332031-1212121303012020-2331101032233011-3322012311231231-1330203011020220"></a>

<a id="canonical-3301100020333223-0110010230001203-1021110012211033-2312301223301132-3020202122020132-0131303221320210-3223301030300101-0000322011210203"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Optional.

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

- [first_address](resources--securemesh_site_v2--reference--group-012.md#canonical-3311130012000023-2010000010122223-3220201303203111-2011112010232332-2220010003223013-2300130103223321-1023023203021101-3120033203230331): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-012.md#canonical-2233130230221312-3120220332101221-0011031112210111-0132203312132130-3113302223133231-3011030332200232-1012030002132122-1000321001212202): complete subsection reference.

<a id="canonical-0230323310211032-2200011010210200-1302220001300022-1311021113232013-1120320030101211-1202330321300312-1202110331111301-2032032332120230"></a>

<a id="canonical-2210233000021030-3020131331130301-1233333023302131-1321102110100323-3131301113021013-0010112021002230-0121210133121222-2000203123120110"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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

<a id="canonical-0321232013030201-0230002032102312-3221022323113200-0203010030103321-3002302021320102-1010211203333001-0000112032003121-0111131220212110"></a>

<a id="canonical-3212110103013022-2101110021102323-1111120202033212-0022200102213000-2113100103300212-3022123030200122-3303200031320123-3332321010213123"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

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

- [pools](resources--securemesh_site_v2--reference--group-012.md#canonical-2033023301332103-0331201333010013-2303302313011112-0101310211001111-1033030123033311-2311313313000102-0032323121222032-2022102020300033): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-012.md#canonical-2322222112203112-0212002012332303-3130120002230202-1131210111200012-2301010300310032-2212322130223231-2333010100230211-0203313330333201): complete subsection reference.

<a id="canonical-3311130012000023-2010000010122223-3220201303203111-2011112010232332-2220010003223013-2300130103223321-1023023203021101-3120033203230331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-2003013302320113-1022312110100032-3130123033102123-2000000300113232-3122001212230302-1110010233321332-0133212011312303-2323310200200201"></a>

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
first_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2233130230221312-3120220332101221-0011031112210111-0132203312132130-3113302223133231-3011030332200232-1012030002132122-1000321001212202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-3230130312213301-1213012313202030-0233123123012103-3203202330132211-1313021213132333-3032332131010333-3220123203131012-0300213122220300"></a>

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
last_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033023301332103-0331201333010013-2303302313011112-0101310211001111-1033030123033311-2311313313000102-0032323121222032-2022102020300033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-3212012210313130-0131103112332312-2122222030320030-3110320333331333-0030111000312120-0130232331311223-1323231023032303-0213301330202110"></a>

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

<a id="canonical-0310302031301301-1323123220110221-2312103203130013-3010002223202213-3023031213232300-3213122321222311-0022023321110233-3100133102312322"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-1300030113003013-1101111030011130-0032301210212303-0023200132203010-0333310130002330-0030210220231012-0030300302322233-3233123320312121"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

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

<a id="canonical-3200023202212122-2032031321311103-3221001100213022-3310012010311012-0203030130210101-2221020020023023-0110111020222010-3130303322300123"></a>

<a id="canonical-1003323321133022-2100112102203023-1000202202023030-3032102032100033-2313033121120111-1200121310231122-0323021001123323-0011002112120210"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-3223132231302232-2210232230002033-3332112323210000-3322122200321012-0322120323033000-3211331001133133-1012021233221303-0112230033100123"></a>

<a id="canonical-2231220303302123-0123212232130301-3032312032132120-2113123112012211-2121021121121031-2331231321003222-3220200023323212-2110301030222202"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

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

<a id="canonical-2322222112203112-0212002012332303-3130120002230202-1131210111200012-2301010300310032-2212322130223231-2333010100230211-0203313330333201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- [nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-0310311333332031-1101221222033203-1111222202122001-3321011303221201-2111130230122020-0230100321321010-3221130311130223-1101310210023101)
- nutanix.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-3102302320013020-0230132210112221-3323321032113221-3033333031013130-2123200112120012-3121230212032332-1221313032001121-0002013123202213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for same as dgw.

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
same_as_dgw = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300230220001111-2120230133130003-0311321100332200-0220200211203310-2321030013223320-1123232033133010-0133312221201200-1233100210101201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-012.md#canonical-1110330222113323-3023013011102013-0322322313130312-2010131100020212-0021331323111332-1230330311310210-2110333120002103-1212001131022220)
- nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2122033310002033-0000332213122010-3100300002333330-2123113301300323-2103030123312221-0213001211002100-1330201313213233-3013111303323013"></a>

Type: `"object"`. single nested block, Optional.

Interface IPv4 Assignments. Specify static IPv4 addresses per node.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3030231120330230-3230330022320312-0233012310031332-1021020100131013-0001121222120220-0131031233011200-2012022321201131-0123120302310312"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-2200102330031021-1121032201103101-0001013120301002-1301121202212112-2033003222330223-2032033212323233-0331300330203021-0233100010021332"></a>

#### `nutanix.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Map{validators.MapConstraintsValidator("{\"cardinality\":{\"maxProperties\":64},\"category\":\"discovery\",\"constraintType\":\"map\",\"deterministic\":true,\"keys\":{\"maxLength\":128,\"minLength\":1,\"type\":\"string\"},\"originalRules\":{\"ves.io.schema.rules.map.keys.string.max_len\":\"128\",\"ves.io.schema.rules.map.keys.string.min_len\":\"1\",\"ves.io.schema.rules.map.max_pairs\":\"64\",\"ves.io.schema.rules.map.values.string.ipv4\":\"true\"},\"values\":{\"format\":\"ipv4\",\"type\":\"string\"}}")}
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

<a id="canonical-1220303103223201-0002300332200330-0210220311211213-3220113033021232-1111312303002001-1030223331130300-3322301012231302-2331321131003323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-0231232303302120-3130210001211023-3322303323000121-2020002101332313-1023120023320323-1110103120330203-3012030300232101-2001102020010233"></a>

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

<a id="canonical-1210210121020302-0230011310301113-3330222302103031-1332111323223213-0000131233221023-1120112222232100-0012211220300203-1202300123311232"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-2202211103033022-3112201211303013-2122103013002030-2321101313023212-0110300220113010-3212231012020103-2232321223123313-0233220121313001"></a>

#### `nutanix.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Optional.

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

<a id="canonical-2030312101221120-2022003221002021-3032123220030203-0322033312202101-2232211323133220-1320213332302112-3112120002320121-1313013020232121"></a>

<a id="canonical-0333123132123001-3302310111122321-3012101212122311-3130300232013121-3211332203130033-1212311230113213-0223200112031302-0103302103130301"></a>

#### `nutanix.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

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

<a id="canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-1311300010211331-1003103023132323-3131002321202031-0202122123133211-3121203311210201-0311031013002301-2221121102303221-3222321032130310"></a>

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

<a id="canonical-3022030312130032-3210300221203031-3211130203100021-2001000131033233-0312311311122233-1130331101103313-0020220330023210-3233020113111322"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-012.md#canonical-2311120322323230-2131001222122112-1011232300302301-2011102232000201-1333131313110220-3002320133120023-1023032132302130-2032330231103331): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210): complete subsection reference.

<a id="canonical-2311120322323230-2131001222122112-1011232300302301-2011102232000201-1333131313110220-3002320133120023-1023032132302130-2032330231103331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-3131321203010122-2002111102033333-3030333303312130-2313211020322302-3103303023331110-1130233311200203-2211320322311122-2231212000030213"></a>

Type: `["object", {}]`. Optional.

Hostname or IP address of the target server.

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
host = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1311211203232112-3101101102111013-2221102132010212-0123231203233323-2202230120223231-0332220232231032-3300321023213003-3212003222331230"></a>

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

<a id="canonical-3112322211311203-2231020222211013-0311321020200233-1132020133202301-3331232312310123-2032203102010131-0323213333321111-0201222220331200"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0201200002231323-0312221130311302-2321210311001103-1213033312302013-1203322310022330-1112210210131200-2203230030002122-3030312030330230): complete subsection reference.

<a id="canonical-2100310233223231-0020002201012023-0122333010211331-2023132320312112-3121133122131320-1303203201032033-2100133230133223-1000123212222220"></a>

<a id="canonical-1231031300113230-3112132131110222-1212222323001200-2232221011202311-2331133001122002-0212131003112223-1013220102031101-2132313200202223"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Optional.

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

- [stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020): complete subsection reference.

<a id="canonical-0201200002231323-0312221130311302-2321210311001103-1213033312302013-1203322310022330-1112210210131200-2203230030002122-3030312030330230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3303210033200320-0303332110102323-1113311202102122-0000000212210300-1131011331322101-2331131103030131-2113010122100113-0102212211303033"></a>

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

<a id="canonical-2000312000131310-0313232302031022-1220212323020030-0203013111031031-3130313031022220-2003010101313100-1023201230211020-2110221132101111"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-012.md#canonical-2223111121233002-3031123301031110-1212200331131123-2200020133222123-0200031122032131-2320133213232000-2202100030322122-1301313221121330): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-3131033001031213-3320102230231131-3312102302300133-1323033122201330-3200323301200211-1001220000121331-2122032223330213-0201130003123030): complete subsection reference.

<a id="canonical-2223111121233002-3031123301031110-1212200331131123-2200020133222123-0200031122032131-2320133213232000-2202100030322122-1301313221121330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0201200002231323-0312221130311302-2321210311001103-1213033312302013-1203322310022330-1112210210131200-2203230030002122-3030312030330230)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-3113102223003101-3223333230202302-0221111113202230-0210012221012100-2013301303021100-3021020332331233-3233332012022012-2032121333312113"></a>

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

<a id="canonical-3123032011123311-1220220102310200-3211022300132023-1103302131300031-2330331013132323-0011330010230301-2223222232311201-1111310221010102"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-1313123002003130-1333020102100311-3331333231302211-0033213202033200-1012120003323020-1312012231133103-1223203211133220-0111023200312013"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

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

<a id="canonical-3131033001031213-3320102230231131-3312102302300133-1323033122201330-3200323301200211-1001220000121331-2122032223330213-0201130003123030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0201200002231323-0312221130311302-2321210311001103-1213033312302013-1203322310022330-1112210210131200-2203230030002122-3030312030330230)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2233030032311220-0231330111223233-1010102312222011-3331113313333121-0000200111222312-3120021223313002-1322213322211030-0122222101011301"></a>

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

<a id="canonical-2311311003020132-1221333220333300-3013002102032222-2313302310222302-0101033230230031-3202231231001032-2130130302010203-0201100101333201"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-0113110322013000-1332000213132200-0301132012231100-2322332022312000-2002301313120113-3331103031010121-0013001213213332-2123102230310123"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Optional.

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

- [first_address](resources--securemesh_site_v2--reference--group-012.md#canonical-0031122211230331-1133230230022103-0011113002202333-2223231313131232-1130312102200312-0233322323132301-1313233331101322-0032123000213203): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-012.md#canonical-2012231210301311-2213212211331232-3011021330332201-0201122022211021-2201112202201003-1213113310211132-0301301110021001-0101133000112220): complete subsection reference.

<a id="canonical-0031122211230331-1133230230022103-0011113002202333-2223231313131232-1130312102200312-0233322323132301-1313233331101322-0032123000213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0201200002231323-0312221130311302-2321210311001103-1213033312302013-1203322310022330-1112210210131200-2203230030002122-3030312030330230)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-3131033001031213-3320102230231131-3312102302300133-1323033122201330-3200323301200211-1001220000121331-2122032223330213-0201130003123030)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-0032011102302311-0321130010010110-2123332311122001-0100101011223300-1212313120231012-1030132303331101-3333013330231120-2303213231012332"></a>

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
first_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012231210301311-2213212211331232-3011021330332201-0201122022211021-2201112202201003-1213113310211132-0301301110021001-0101133000112220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-012.md#canonical-0201200002231323-0312221130311302-2321210311001103-1213033312302013-1203322310022330-1112210210131200-2203230030002122-3030312030330230)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-012.md#canonical-3131033001031213-3320102230231131-3312102302300133-1323033122201330-3200323301200211-1001220000121331-2122032223330213-0201130003123030)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-0321211320220320-2022222033220330-0202032313021303-3303331111220333-2200003023201302-1123011330122010-2301203313303032-3100003303111303"></a>

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
last_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-0222310312312311-3323110331022223-0102131210101232-1010303220120121-2100210033030002-3203232333011131-3121211020322310-1023013330131030"></a>

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

<a id="canonical-3300332320013032-2220313323031010-3021120002033112-0300011010322113-0230101112011101-3230331131013020-2322020302110323-2013201102132021"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-012.md#canonical-1220111010300111-3232020103013112-3203310233330113-0030022121123221-3313010023110000-1332203202210212-1331212332133003-1322132333333010): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-012.md#canonical-2021133013102301-2200211130132010-3203211203102201-0001013302332122-2301303200302023-1120322133203311-2231112222011300-2002313310020323): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-3222300231032221-1123010333323113-0130301021130222-3323312033310230-0330123322133322-0211303000032201-1003200020001331-3130212133002131): complete subsection reference.

<a id="canonical-1202102202011320-3133220022122201-2030212231020030-3022330313003030-3031231010113313-0011221021233010-3111311222311020-3301020233010013"></a>

<a id="canonical-1301003200323120-1103201303212030-2312311213302211-1110223030010020-1032322301203011-0223013303131333-1233312010230010-2232302002313013"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-012.md#canonical-2223032132132023-2333110001300313-3232022020101133-1220102002001232-1000303023033321-1111222032312301-2231110132312132-1001321031102100): complete subsection reference.

<a id="canonical-1220111010300111-3232020103013112-3203310233330113-0030022121123221-3313010023110000-1332203202210212-1331212332133003-1322132333333010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-1130012120030120-0011000332110103-1222322303323021-1033313310100300-1122002323030323-0332212310013311-2110332331111132-1332322221120101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from end.

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
automatic_from_end = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021133013102301-2200211130132010-3203211203102201-0001013302332122-2301303200302023-1120322133203311-2231112222011300-2002313310020323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0010333330011233-0032203010201213-3102130022111320-0033130121233030-3130130323102223-3131321001113031-3231133231003230-3000320001200330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for automatic from start.

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
automatic_from_start = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222300231032221-1123010333323113-0130301021130222-3323312033310230-0330123322133322-0211303000032201-1003200020001331-3130212133002131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-0120212132100210-2201203220312133-3211112111321033-0313000311201111-3310031032212012-3333223323032100-0312103213131302-1322232001311023"></a>

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

<a id="canonical-2113011003302330-2333033311100203-0100212110332120-2311101100203210-0300131201301001-3312010301221302-3311022031201322-2312121231131301"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0022031231121320-3233101132023333-3320310120210130-2130231300222123-0003011113021212-2023332033312230-3312220011301213-0230120223210002"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

Type: `"string"`. Optional.

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

<a id="canonical-2130220033103013-2113200222302221-3203320031311133-1120323120112003-0323301212313301-0213203020013232-2122321000101122-1333212123111313"></a>

<a id="canonical-0031233212332103-0131102021103333-2020033233320320-3002202002231011-3213013222020130-3023223113212310-2330123130111111-3110212321002101"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

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

- [pools](resources--securemesh_site_v2--reference--group-012.md#canonical-1300001003120321-0103123321031020-1333113123030100-3200010300223031-2120111121312311-3031233223213333-3220000032213122-3021013202011313): complete subsection reference.

<a id="canonical-1300001003120321-0103123321031020-1333113123030100-3200010300223031-2120111121312311-3031233223213333-3220000032213122-3021013202011313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-012.md#canonical-3222300231032221-1123010333323113-0130301021130222-3323312033310230-0330123322133322-0211303000032201-1003200020001331-3130212133002131)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-0221112320130233-2302113222202112-2332030010112313-3020022223033222-0330320111010212-2223202030330022-1101101130011210-3332310202132212"></a>

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

<a id="canonical-1333100222302323-2212012230010131-2010030211003221-0110130222133212-3211102311003232-3302100012310211-2302303230000203-3212331233103231"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-2321133033212202-0112303100212100-1000210122012310-3022312302200211-1130322102102230-1200203003013212-2103130012013321-3300303312230000"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

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

<a id="canonical-0231001210200122-0222110122233102-0201001101233003-1331032000302232-3321211122111301-3203132021132101-0320110230030131-3102303200022201"></a>

<a id="canonical-2203030120303020-2311111021033233-2022300223033213-3221030002231213-1111012101223301-2113010022020221-0221222121201132-0102230123322122"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

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

<a id="canonical-2223032132132023-2333110001300313-3232022020101133-1220102002001232-1000303023033321-1111222032312301-2231110132312132-1001321031102100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-012.md#canonical-2222032223112113-3013221102131202-0213311332322121-1132031303023333-3012330201120013-3130200003011123-2303012131331223-0131030303223202)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-012.md#canonical-2113023031322030-2021300103323212-1212322131310322-0212203100021201-3113222233031330-2333223333022210-1230312033221132-2023313222212210)
- [nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-012.md#canonical-2012203230010333-2003330011033113-3022012022033223-1222320221023312-0233233303200320-2231202123113311-0301322200023332-1133202122320020)
- nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-0331212133000132-2122230210102013-2233120032000032-2230323220123311-2330102110121113-3231313202323222-0013232030330122-0311112101001231"></a>

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

<a id="canonical-3331212003302120-1022002012330033-0231211212323003-2232313331300030-1102332323132113-3120232223022320-1031222111102123-1013101020201213"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-3112023121011322-3203010223211332-0333100313022110-1200130120201322-1021023101033203-2222003332031102-2031320213322321-2231031300121013"></a>

#### `nutanix.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

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

<a id="canonical-3123320032100001-1031133131223231-0011130301203201-1123122201301021-0120221213231131-1110333302111203-0032120301303121-3232123031302301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.monitor

<a id="canonical-1133201213301130-1321323003233330-0001111202332101-2310001311300223-1010032223302231-3202333213331113-0330120230222320-1031211033301120"></a>

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3123322001322102-3101003102311001-3021021212103320-1211232031000131-3021021232200112-1233221230310332-3303221022112021-3131022012122013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-3223320300111321-2031303201311202-2103211211213223-1031022113130011-0310122122330113-0132222203223101-1210300221331021-0113033301301032"></a>

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
monitor_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033233301013022-2201212323332122-2311010013222211-0233122301232213-2022203031021300-2002200331012323-3201033130200012-2133230201113131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.network_option

<a id="canonical-0111022223321300-0321101000033033-1212313212231322-3222231101111112-1321333030020011-0330131220020313-1130322021031100-1121310101120321"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3111023013312103-0113311000020010-2103132003301233-1231320221011132-1011002100113013-2321320001300101-3011300023133202-3100103133031202"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-012.md#canonical-3130323333121200-0110102201002123-1320302133301102-0230000110321223-1030232032002301-1103202303233023-2033331330113332-2221101222131110): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-012.md#canonical-1223322132200200-1233233020202100-1020212132331010-1313132112003311-0103023123001201-3333130333003100-3001030311301022-3110121201020201): complete subsection reference.

<a id="canonical-3130323333121200-0110102201002123-1320302133301102-0230000110321223-1030232032002301-1103202303233023-2033331330113332-2221101222131110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-012.md#canonical-2033233301013022-2201212323332122-2311010013222211-0233122301232213-2022203031021300-2002200331012323-3201033130200012-2133230201113131)
- nutanix.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0121310200300010-3003011330233123-0331212023121212-2313223032211231-0113330101212101-2310000121110223-2212012213310300-1033003100133000"></a>

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
site_local_inside_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1223322132200200-1233233020202100-1020212132331010-1313132112003311-0103023123001201-3333130333003100-3001030311301022-3110121201020201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-012.md#canonical-2033233301013022-2201212323332122-2311010013222211-0233122301232213-2022203031021300-2002200331012323-3201033130200012-2133230201113131)
- nutanix.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-1210233222323013-2123231312210221-1032322232312211-0100222033320003-2300230300011322-0230113223322203-0321223332300020-3120120030213332"></a>

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
site_local_network = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300113120101223-1220103101320303-2033212212031203-0302232302113211-3320033201230113-3331203111100003-1222310233300013-2003010201023320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-0202313113011011-0133012123111221-3122323323110031-1033122110131133-2310231231231000-3233232032321313-2231113232210330-2103333202312302"></a>

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
no_ipv4_address = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003012130113301-0102301003301331-3223303121002030-1103313212120000-3101301301110332-1210200200131110-2221021033010310-1313302111203121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-2212332211203202-0231032103311333-1000323223200311-3011221220220303-3002111320203312-3200001023031010-3211310211030021-3000331310230312"></a>

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
no_ipv6_address = {}
```

This is an empty object or choice marker. It has no direct properties.
