---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.network_option

<a id="canonical-0000110010300103-1001310000023333-3230222020231030-0123030022312233-0022330013232221-2330000323011101-0101113022213011-0002013122012123"></a>

Type: `"object"`. single nested block, Optional.

Select virtual network (VRF) for this interface. There are 2 kinds of VRFs, local VRFs which are
local to the site and global VRFs which extend into multiple sites. A site can have 2 Local VRFs,
Site Local Outside (SLO), which is required for every site and Site Local Inside (SLI) which is
optional. Global VRFs are configured via Networking &gt; Segments. A site can have multiple Network
Segments (global VRFs).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1311223021320123-0103001333123132-0221212021332202-2301232012313222-2131102030330301-0113123320123230-1022312331123031-1233312133321302"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-019.md#canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-019.md#canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132): complete subsection reference.

<a id="canonical-1030303230330033-3233022020212132-1200032213112003-2230200010130010-0300030230310020-2323310013313002-2112311330030101-2230322232011120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- vmware.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-0021332210202012-3323031302101032-3322011323100302-0221032112021230-3012211230033021-0020022010122230-3230030033111202-0310332013033332"></a>

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

<a id="canonical-0011213002312300-3322013032110301-1020021333010121-3311332003121202-3232013003231101-0320231203020201-3232320310320023-2001031121210132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-019.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333)
- vmware.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-0112333032000011-2211320102031221-2222030220310002-0311211231221010-3211220320202003-2032212132102110-0321201232323023-0220213000320313"></a>

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

<a id="canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3220332231331012-2221310003003010-1323111102311003-2131113312211320-2202202332223010-3021302001212230-2031222123212332-3300001003110101"></a>

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

<a id="canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-0303221212303102-3032012032020231-2133232133100023-3003122000331202-1010132332032203-1001300003232003-0022113331213121-0123031002313111"></a>

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

<a id="canonical-2322001313121132-3312101121211223-2202213121003332-2210031110113011-1032002312002332-3213313220012333-2333302100232111-0030320302210211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-1322321123011233-3012311300311111-3102111312131031-3113313031033033-1032022230010022-2302011133010120-1222232131102311-0023321122302233"></a>

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
site_to_site_connectivity_interface_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2011222131222303-1120320332311132-2230113030002332-0022233303003103-2011112021303120-3031020201232030-3103122201032211-1323103122102202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-0201203123030221-3312023020123121-3102331031320013-2230302121033103-3121031101033101-2210003212130320-3212301020122013-0101113133033033"></a>

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
site_to_site_connectivity_interface_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302310220332003-3233011030011030-1331333103020220-2022131230232223-0321013011113202-3230130331131103-3300020020223203-3131022022232000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.static_ip

<a id="canonical-0321323012033110-0130002002320113-0023221021200032-2022300001330200-2132301012303011-2120120330223323-0333002032103200-3320111300233213"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0330130301101001-1303333012122231-1033312131300133-3212032313202301-0212131100323021-1013301203123121-1112002233122122-1101302011332211"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ip`

<a id="canonical-3032310032211133-0121302203132211-0121131332003333-2331310133020222-3003120202331010-1202330122203202-3103331132013103-2301122333122303"></a>

#### `vmware.not_managed.node_list.interface_list.static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0121013012310121-2300301231320211-0320221222301012-2123033311122032-1030000132302333-0001333311332210-1100310322011113-1000011300032223"></a>

<a id="canonical-0001212221313313-1033332320100013-1033232111130030-2021121120130312-1103101311131023-3130130332211020-3300202300130021-0030101220103022"></a>

#### `vmware.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0313222311032030-1133332203322010-0333122321002000-3313231221323320-1322101102300223-0131110201101322-1032032022031302-0232020012300003"></a>

<a id="canonical-0000313012300211-0011301102320303-0030133001333101-1323113013031220-2101211011110123-2213222003201301-3111323303231121-1211323221120112"></a>

#### `vmware.not_managed.node_list.interface_list.static_ip.ip_address` property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-1210012201220010-1122313132100032-3202000232113200-0210133321320020-0111132302331030-0100202232322131-0201213310210013-2302130010302031"></a>

Type: `"object"`. single nested block, Optional.

Static IP Parameters. Configure Static IP parameters.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0020222103322302-1030123202301033-2101211013212202-2332003221232130-0201010020003331-1003132331322001-1131110033330330-2123233302013332"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-019.md#canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031): complete subsection reference.

<a id="canonical-1113002301030132-2333323233320201-3210212223202330-3201303311222021-0033113232100120-2013222230200103-1103131100321030-1011332301211232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1121020030103121-0011012302201011-0111110121000002-1301032122120020-0011211213122232-3002130330100323-1203330320313330-2330321030310112"></a>

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

<a id="canonical-3322220031021120-2113221311332011-1201032133331333-1111333302002123-3132311111211023-2101233022102323-1131333101031201-3332331012332012"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-3133311110032300-3111020202211223-3120312131311010-3301112333301000-3310102011113310-0033311002112321-2101221321231211-0131020321033023"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Map of Node to Static IP configuration value, Key:Node, Value:IP Address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1322303221121010-1211211031021110-3013011121322031-0212113331203323-2331331220021311-1101233311011013-2131002333113132-3132231100301031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- [vmware.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-019.md#canonical-1031322112313113-0003301133110003-1202212102111211-0012030111230133-3323103200233100-2003231301311333-2203002101103223-3102230111310000)
- vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-2001232022213220-0212320021003220-1203212311232032-3331333302023320-3013312331222210-0230301110301220-0020102332002012-0103133331321300"></a>

Type: `"object"`. single nested block, Optional.

Configure Static IP parameters for a node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1301130303021200-2300033031120013-2001301112200123-3021022202212100-2130023221220020-3330221133203231-1330130033210121-1200301102313311"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-0030203311312021-0001330230101023-1202113101002013-1000213002121332-0220220310010013-2002202200132300-3330310200023330-3223132000113212"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

Type: `"string"`. Optional.

Default Gateway. IP address of the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3122211310221101-1131111201020123-0133330200133222-3232221112130301-0032111101313200-3311310133032003-3300133331221220-3121210303222223"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0301003230301122-1002301130031231-1213023323030110-1211002301122321-0302011331031202-3112213102013002-2021023230211030-1301131212213210"></a>

<a id="canonical-1321010312311022-2101230021231303-3030312202222330-3023000101333012-1022221130022210-3331212131033120-0321023131311303-0301122131010203"></a>

#### `vmware.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

Type: `"string"`. Optional.

IP address of the interface and prefix length.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0122021003320110-2010312130313220-0211022100031121-1003222012331013-3101102003210222-0113013000230332-2123230032011112-3213312112302021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- [vmware.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-018.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211)
- vmware.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0033222112000013-1313333033331201-1321221321001220-1300120210320313-1102331313112030-2330301122321300-3131213121232031-2123332010230223"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for vlan interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3100103112101210-0331313333013021-1221332111021210-0220031300331011-1023200323321021-1120122203100003-2233012233312323-0023021130220032"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-2303302012110113-2333301121032223-2320122233000101-1303001300222130-1313323001111112-1110010330123211-2130122000102120-2232022122331221"></a>

#### `vmware.not_managed.node_list.interface_list.vlan_interface.device` property

Type: `"string"`. Optional.

Select a parent interface from the dropdown.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-0210300232112030-1021023013032302-0120130210220131-0312222303012231-0230223031003000-0030000202030232-2020030001231100-0212113210233300"></a>

#### `vmware.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

Type: `"number"`. Optional.

Configure the VLAN tag for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
