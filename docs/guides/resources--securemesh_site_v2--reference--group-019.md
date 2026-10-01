---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1301130303021200-2300033031120013-2001301112200123-3021022202212100-2130023221220020-3330221133203231-1330130033210121-1200301102313311"></a>

## vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip — node_static_ip / 131100223312 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2001232022213220-0212320021003220-1203212311232032-3331333302023320-3013312331222210-0230301110301220-0020102332002012-0103133331321300"></a>

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

<a id="canonical-3122211310221101-1131111201020123-0133330200133222-3232221112130301-0032111101313200-3311310133032003-3300133331221220-3121210303222223"></a>

## Direct properties — node_static_ip / 131100223312 / 3

<a id="canonical-0030203311312021-0001330230101023-1202113101002013-1000213002121332-0220220310010013-2002202200132300-3330310200023330-3223132000113212"></a>

<a id="canonical-1321010312311022-2101230021231303-3030312202222330-3023000101333012-1022221130022210-3331212131033120-0321023131311303-0301122131010203"></a>

## default_gw property — node_static_ip / 131100223312 / 4

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

<a id="canonical-3313302212131223-3112000233130233-0232223220221122-0100301223301231-3021130220301232-1321021120133331-0302012100000002-2213010320012123"></a>

<a id="canonical-0211323222100331-2101211201202102-3331120002121032-3302101131032303-3032321322232320-0311331202312212-3233211010101002-3120233320320102"></a>

## dns_server property — node_static_ip / 131100223312 / 5

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0301003230301122-1002301130031231-1213023323030110-1211002301122321-0302011331031202-3112213102013002-2021023230211030-1301131212213210"></a>

<a id="canonical-2213021212120130-1200033330132213-0000221100133123-2002233122032222-2212220230231133-0330133000301122-3133132133333021-3012013331300330"></a>

## ip_address property — node_static_ip / 131100223312 / 6

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

<a id="canonical-0031230213110031-3310233330122303-3131222333231323-1233121222012100-0313301323011021-0033020322110232-1232303212012120-1313230322221323"></a>

## Next pages — node_static_ip / 131100223312 / 7

- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)

<a id="canonical-0122021003320110-2010312130313220-0211022100031121-1003222012331013-3101102003210222-0113013000230332-2123230032011112-3213312112302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3100103112101210-0331313333013021-1221332111021210-0220031300331011-1023200323321021-1120122203100003-2233012233312323-0023021130220032"></a>

## vmware.not_managed.node_list.interface_list.vlan_interface — vlan_interface / 030012111131 / 2

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0033222112000013-1313333033331201-1321221321001220-1300120210320313-1102331313112030-2330301122321300-3131213121232031-2123332010230223"></a>

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

<a id="canonical-0210300232112030-1021023013032302-0120130210220131-0312222303012231-0230223031003000-0030000202030232-2020030001231100-0212113210233300"></a>

## Direct properties — vlan_interface / 030012111131 / 3

<a id="canonical-2303302012110113-2333301121032223-2320122233000101-1303001300222130-1313323001111112-1110010330123211-2130122000102120-2232022122331221"></a>

<a id="canonical-1103002321110102-1102133001233023-3110220231320310-1023300303222330-3111000131312331-3130321000210100-2322103300021123-0231020331113222"></a>

## device property — vlan_interface / 030012111131 / 4

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

<a id="canonical-0011123311231110-1321322012202231-3103330130320202-0103102112022032-3211031113223300-2312213213001221-1222333000332331-0332221221001030"></a>

<a id="canonical-3021321101103202-3223023220203011-3303122020022030-0032003203123110-1021333021133333-2210323112312200-0031213033031023-3301122002100003"></a>

## vlan_id property — vlan_interface / 030012111131 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0332023330111322-3003213133310113-3120303301120120-0001223220023023-0110311123313201-0031230213023100-1322101203231113-3330033303230322"></a>

## Next pages — vlan_interface / 030012111131 / 6

- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
