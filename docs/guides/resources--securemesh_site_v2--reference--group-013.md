---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

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

<a id="canonical-1231301312010313-2213232000112222-3013230012232221-0133121113311310-1103002210331213-3120322321210302-3322223321233130-2113230112301313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-2333313103002033-1000213110000323-2220311111212113-3220120103233203-2230010131033223-2323112111211212-2201330100330130-3300032023321303"></a>

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

<a id="canonical-3213332311302012-0322110013322032-2023302321211111-0301121010312031-3201102121211013-0110331013133311-3000110302320313-1333330332032231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-2032113313100332-1010101201113221-3033332130003213-3333320222021222-3110321211210332-0112023323113030-2131321320112303-3203011220132001"></a>

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

<a id="canonical-3000311211321033-1020312111303303-0333321101203202-0120301212302320-2133131020320003-0302021131332021-1303313201120311-0232213200222202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.static_ip

<a id="canonical-1030231320002020-3003213333032222-0022030230121200-1102010320001231-2120221223332131-1022010101032002-0122202133332123-3120203232121202"></a>

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

<a id="canonical-3101131032032312-0203132003133112-0201310103301023-2002321012121030-1200111120230313-0232323200313233-3001201022033321-2121132313113212"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ip`

<a id="canonical-2311212131100001-2322212223332100-1030000231201103-3003201100022002-0121133210100131-1021211323233330-1103233011030000-2333031211202321"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ip.default_gw` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1021011213321111-2320221000011111-0201203302131222-2131001221203010-2201023030313232-2211110322033102-3222133012310210-2220002223001032"></a>

<a id="canonical-3031221103230203-2121322012103231-2312112300100031-2303022211332020-2202223231021130-0211301321032133-2320323210332111-3222212102103332"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-0202023300230120-2120011113312111-2033123321020000-3003020331331321-1230312230230223-0223130002100131-0133011123131312-0223011201133113"></a>

<a id="canonical-1210332223223021-3012312220201120-0100132113033303-0211332031203101-3033202122223112-0232020120222323-1131033030212133-2222233331121332"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2132202020202332-1110010011212323-2210113032102032-1032011021313132-2320203033333212-1130002223312202-1300201210000223-0101311100130333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-2310201133302331-2230010200302031-3331221222103033-0230122202311213-0201123122021133-3121133121010221-1310013332101011-2010111310000303"></a>

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

<a id="canonical-1210212331100313-2231020323222103-2313322231020012-3212111102031132-1201302123003200-3323223210332032-3022222102011300-0102121011231102"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-2102032130303323-0323002203023103-0033001333320102-1001323333312302-2011123323232233-0200031001102321-3331120320120122-1032022203331022): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-3212333302232033-1000012112223023-2130000311012321-1333020130201010-1120122012322210-0013331212021303-3223233030320133-0202003023102320): complete subsection reference.

<a id="canonical-2102032130303323-0323002203023103-0033001333320102-1001323333312302-2011123323232233-0200031001102321-3331120320120122-1032022203331022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2132202020202332-1110010011212323-2210113032102032-1032011021313132-2320203033333212-1130002223312202-1300201210000223-0101311100130333)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-2012323121123312-1003103203300110-0001120023233133-1102003120203021-0123012011131110-3001031003330302-0102031003100032-3130033032213302"></a>

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

<a id="canonical-1022310113332200-2111002112020212-0211032310321301-1021133212322221-1020220102320210-3111311021321002-0002331330000201-0212233101022132"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-1321122310231112-3300302222301113-2102130310010101-3302111002220222-3133320030310031-0100131200323323-0133322100130301-0211303200311203"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-3212333302232033-1000012112223023-2130000311012321-1333020130201010-1120122012322210-0013331212021303-3223233030320133-0202003023102320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- [nutanix.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2132202020202332-1110010011212323-2210113032102032-1032011021313132-2320203033333212-1130002223312202-1300201210000223-0101311100130333)
- nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-1201330022231101-0200020031010113-1033201332330201-0200121111300033-1013313002311332-2102221023311330-0331203321031011-1332100333312323"></a>

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

<a id="canonical-3120233333331300-0211303001310003-0133222112230211-3111332120031232-3122012131120000-2013101032100203-3203321111121323-1331232322203220"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-2000022321000112-3302031000000333-1312320321121302-1031303223111202-1203212313330210-0301031203312111-1313013312310102-1232010033110133"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1123001211311011-3011123121311210-1331300332303333-1313100012130012-2310223111232323-2003200223303302-3230032121303011-1310121000323311"></a>

<a id="canonical-3332102223211322-3233213211120210-3130011100220122-2201023021033112-3120222000003023-3133211200312203-3003001330220012-1210132311302030"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-2222303002120131-1000021130313200-1301110232212333-3203000103123313-0222302032311300-1310011232023332-1222311212233331-2202120012200303"></a>

<a id="canonical-2300210032221011-1300332110123303-0111001113103231-0122301300121302-0131321213222013-3310133222133012-3002312300333023-2133113120310110"></a>

#### `nutanix.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0001022032320011-2010231101113231-3102012231032321-2010221100201020-2131313312332202-3123322133001323-3312023211310120-1100230201223210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `nutanix.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [nutanix](resources--securemesh_site_v2--reference--group-011.md#canonical-3313330100011000-1321012221111023-2203201131220011-0021313133031220-2300300000311310-1301020003123302-1012202333312321-3300313210310321)
- [nutanix.not_managed](resources--securemesh_site_v2--reference--group-011.md#canonical-3220112021101122-3110130321202300-3000232201030020-2030130010113300-3130023211030002-3301102302032211-2131231030302103-0302111102333000)
- [nutanix.not_managed.node_list](resources--securemesh_site_v2--reference--group-011.md#canonical-1223230220233333-3021121210321110-3201032121033100-2123303013201331-0102030230013033-3331232111223103-3201332331212232-2110010123111231)
- [nutanix.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-011.md#canonical-3132013111121322-1131131200100030-3111302000312210-0102200232202332-3223003333332223-1333213321210012-2013112001021313-2333000031110202)
- nutanix.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-3013032031232220-3130002011000011-0300000232112102-3102311110323120-2333212031210000-2200212302101113-1000202130200230-1103030133001121"></a>

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

<a id="canonical-1132200112022221-0203111320103202-2020130332333323-0122011313103033-0221330332210223-1302303133321003-0323012131123103-0313230201322320"></a>

### Direct properties for `nutanix.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-0122331023330203-3313122103202200-3110102332202121-0313011131110101-1131220211000132-1222201332111132-2231132211330100-0231113111021302"></a>

#### `nutanix.not_managed.node_list.interface_list.vlan_interface.device` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0321202311031133-1031323313012030-1301010310012010-2102032100210221-1200222103221032-3120212022311132-2213020233030202-3320321320221301"></a>

<a id="canonical-0300113111012200-2323230232102111-2100030302210322-1002230200033201-2210222211320133-2031121312120220-2033033013333202-3111110130222303"></a>

#### `nutanix.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- oci

<a id="canonical-2201132123330101-1202132101032121-3211232123013100-2201213331013122-1203112031100031-2310320333103212-3013300202203221-0202021112102331"></a>

Type: `"object"`. single nested block, Optional.

OCI Provider Type. OCI Provider Type.

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
oci {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212300001221302-0303202312131023-1333223301131110-2010111113011231-3202133233201011-2001210323230202-2110131122300010-3120332022210132"></a>

### Direct properties for `oci`

- [not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221): complete subsection reference.

<a id="canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- oci.not_managed

<a id="canonical-3200120021203121-2110112200010302-0203131230232312-1021002233222210-2131011121311132-1300120030133113-3200330121010113-0221231233120013"></a>

Type: `"object"`. single nested block, Optional.

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

<a id="canonical-3000010030010302-2110223121311121-1132202311010102-1023330210022233-0101320030333313-0313331112301322-2321312112322200-1100211321203200"></a>

### Direct properties for `oci.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221): complete subsection reference.

<a id="canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- oci.not_managed.node_list

<a id="canonical-1110022201111121-2023001310203022-1111230031023311-0323332331313330-2021013020021322-2303131223023101-3303003022020211-2330132013001102"></a>

Type: `"object"`. list nested block, Optional.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3020000012131323-3131332303310120-1003222311133001-2101003023100222-3030033011033312-1101111023133033-2022313211010311-3130313330233313"></a>

### Direct properties for `oci.not_managed.node_list`

<a id="canonical-1331200011111023-3021333313131032-0131212210212300-0221030100320032-2122332010230333-2303313233130320-2221003323320322-1022231121301011"></a>

#### `oci.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311): complete subsection reference.

<a id="canonical-1030233221022312-3022203122330112-3222110313001111-3311010310303001-3330001002300333-0323113221320201-3302332130132032-1023003102213121"></a>

<a id="canonical-2022113030211200-1130233300331021-2133112001231223-2331202122231102-3231122221210213-3222130100101133-3310000002311001-2201010003333031"></a>

#### `oci.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0300001131120210-3121200123002213-1001231031200033-3311002212331322-2103001013320110-2131133300030201-2203132221333023-3332222230013222"></a>

<a id="canonical-3210231313233001-1322322211021230-0311003132130112-1000100031313121-1002023123010003-2210132030110202-3312230102030231-1321021131211012"></a>

#### `oci.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["Control","Worker"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- oci.not_managed.node_list.interface_list

<a id="canonical-3312312311022023-3333130112032321-0133211000312302-3130333212011332-1021122021230010-2003330132112231-1102310232211110-3211003113003223"></a>

Type: `"object"`. list nested block, Optional.

Manage interfaces belonging to this node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2121110303301201-0132021003020000-0030203321100102-3100011023222021-2010000233322213-3300003033300303-2022310312112200-2232301303323010"></a>

### Direct properties for `oci.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-2323111001332120-2210132200230122-3321210130001331-3030102133111120-1032231332130300-1201112012200233-0213230001313010-2132033030023132): complete subsection reference.

<a id="canonical-0323002002300102-1112033333121012-0121203112231313-1210123003012203-2010222030321212-3231210102320020-2312102003211313-2312023130102022"></a>

<a id="canonical-2030312301103230-0032003031031320-2020122203213021-1323302123133233-3113202210032010-0330232102033130-2003232123010002-0102002301013030"></a>

#### `oci.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-013.md#canonical-3303112303003202-1100003211213131-3310213300320131-2233002220331030-1021200322022103-1022003023201330-0013122010113333-0312111100301001): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-0102100102313212-0321212212322322-1201012200002323-1233303312311112-1131103313332322-2000330320123022-2032313010223323-2133321132021131): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230): complete subsection reference.

<a id="canonical-1320130031023203-0021223231133110-3301211322132200-1022020220221013-0310233023013330-3222032223110310-2122123222210302-1003023123002230"></a>

<a id="canonical-0201132321201122-2022232310303000-0120012120231300-2001331001000212-0330110031011033-1203212112232122-2313211003321211-1323023230220301"></a>

#### `oci.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-2002112332003000-1132102302003222-1201102011112032-2331310113323331-3332123230333220-2202013030232110-1203000310200231-0231130021030320"></a>

<a id="canonical-2110003111200030-0302212311233330-2001103032100023-1233333310231232-2320201203213132-0311333021322230-1012002313321103-1233200133233122"></a>

#### `oci.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-2121102121131013-3321333002011112-0100121200303220-0301320112130222-2233123222231023-1331000021200132-0022211011201220-3220031323122213"></a>

<a id="canonical-2313302310201002-0130020311313113-1033102100230032-2131013331031320-2210200203222021-2300002010000322-2030332200232110-0300000332112200"></a>

#### `oci.not_managed.node_list.interface_list.labels` property

Type: `["map", "string"]`. Optional.

Add Labels for this Interface, these labels can be used in firewall policy.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [monitor](resources--securemesh_site_v2--reference--group-013.md#canonical-1012121000003222-0223311300012000-1321322100020300-2111313001312211-1132031013310230-2112202303300133-0233031220021033-0230103333112311): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3111133321323221-0013021133332313-3321010120011200-3303110120132132-3201001233220033-0301213202033310-2021201320202320-1031121200023322): complete subsection reference.

<a id="canonical-3303320313201130-1232000320322100-2000200113223013-0012010331233311-1203010313222220-2103300120022302-2100022213121201-3011010303320203"></a>

<a id="canonical-1303230223032223-1122110221221111-3120313020130321-0310312213233122-1122312130021130-1233321303230021-1320320200311210-0130032212113030"></a>

#### `oci.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

Maximum packet size (Maximum Transfer Unit) of the interface When configured, MTU must be between
512 and 8000.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0133212201012013-1032320213132010-1110120332001103-1310320303312013-0113211201231032-0031300203031303-3213233012331330-2211220210000021"></a>

<a id="canonical-1211311310021133-3031232200321102-2001320213023000-0111302320312011-0200003231121221-1202230201122111-1033233230031033-0203102311130020"></a>

#### `oci.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-0210321201302320-3311001021020211-2233102321311302-1121221100213323-1113211331021112-1310311322333213-1003111212130221-2300312123233230): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2130101021120001-2133213211000101-2300112023322331-2020212103120332-3211301312013310-2012020321032321-3231320133123220-0232130330333231): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-0300112132322302-3020221123000320-1210301232122310-2222333310111201-2122302132101232-0100013233220120-2121311102120331-0032012132312102): complete subsection reference.

<a id="canonical-1131020203203010-2113223301223213-3202200320310021-0313131122321311-1022203011213202-3032332002303311-2203123223311330-2022101112203012"></a>

<a id="canonical-1100211013003230-2111102001032231-2100232010202000-0011103321231001-2022332021111332-1112302221311000-2031111212320000-0002013123312231"></a>

#### `oci.not_managed.node_list.interface_list.priority` property

Type: `"number"`. Optional.

For a node, if multiple interfaces are configured in a VRF, interfaces with highest priority will be
used as active and interfaces with lower priority will be used as backup. If multiple interfaces
have the same priority, ECMP will be used. Greater the value, higher the priority.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [site_to_site_connectivity_interface_disabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3003202120102201-3213210231231301-0332020103012310-2233300011102333-2333023102013220-0232120303110010-0003302110312032-3110123330111121): complete subsection reference.

- [site_to_site_connectivity_interface_enabled](resources--securemesh_site_v2--reference--group-013.md#canonical-3033302102222012-3000301233233301-2213032233212122-3131331003110300-0212132220003022-3322232313013302-0133322301103221-2233221222132002): complete subsection reference.

- [static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-3010222121020100-3200320332020122-3001031023010002-2031200013313032-0301001232022221-1020321212302012-0123013001102011-3102003213111310): complete subsection reference.

- [static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-0022320311200201-2022123231133030-1003100132020003-2023121231320123-1130203301010300-0010101333111201-1013111001221323-3131232313332233): complete subsection reference.

- [vlan_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-3122131201212313-0011330320013131-1213002113003132-3301222030302110-1031321231302011-3201100110111221-1332320023322201-1220211211132031): complete subsection reference.

<a id="canonical-2323111001332120-2210132200230122-3321210130001331-3030102133111120-1032231332130300-1201112012200233-0213230001313010-2132033030023132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.bond_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.bond_interface

<a id="canonical-1313100101110123-2312323332111330-0220123102210331-1030332100123333-1032121133332210-0313010320323121-3002120001133211-0233003303021033"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for bond interface.

Additional upstream details:

Bond devices configuration for fleet.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0311323031213033-3132002110300011-0333202201010223-3331211130303110-2010023232333130-1223010231333010-0120202122233321-1323020300100023"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.bond_interface`

- [active_backup](resources--securemesh_site_v2--reference--group-013.md#canonical-1312220333003022-3312113001020101-2102000230110130-0212200203032332-3300123303100132-1230131230112131-3133103313331001-2212312203111321): complete subsection reference.

<a id="canonical-3310113222312313-1203131331321203-0323313301220311-3232220022231003-2210013213023310-1102102233132011-2032012220330320-2030133020233202"></a>

<a id="canonical-3312122101113333-3012100113122331-0202232333102213-2003310213213013-1033212110000222-0330302302203121-0031013123021120-2130131031120223"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.devices` property

Type: `["list", "string"]`. Optional.

Ethernet devices that will make up this bond.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [lacp](resources--securemesh_site_v2--reference--group-013.md#canonical-3022121003032122-2302013221030012-0322223010130310-1333013320010000-2002112302011111-0220101300321313-1321221301233333-1322321102131103): complete subsection reference.

<a id="canonical-3111200011011023-0210130113212230-1331322101012012-1202101300020120-3313003021033030-0003332113130132-1121120102030133-1100323030312303"></a>

<a id="canonical-1122032012032200-2312331300332313-3200332221233303-0031310130232232-3323031003222000-2103320033220231-0223202331303010-3203311022001013"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.link_polling_interval` property

Type: `"number"`. Optional.

Link Polling Interval. Link polling interval in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0031131033313330-3200220323100221-0221310211330223-2301130230323231-2213320310212022-2033032003103212-2110323213011233-3313202320211301"></a>

<a id="canonical-3002103100221202-3123112003331131-1332330121001000-0221001112023220-1320001203210021-3212213033112131-1123133233311011-3101012000000303"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.link_up_delay` property

Type: `"number"`. Optional.

Milliseconds wait before link is declared up.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3321231330323203-0222213213103122-3333023223303122-0211231222300030-2313320220033333-0001030303030202-0220302112100013-3331130130102033"></a>

<a id="canonical-2021210210123111-0200020310000002-3302303111310223-3011023023122001-3200201233201323-1221103223123300-2213120113133011-0031110312232203"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.name` property

Type: `"string"`. Optional.

Bond Device Name. Name for the Bond. Ex 'bond0'

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1312220333003022-3312113001020101-2102000230110130-0212200203032332-3300123303100132-1230131230112131-3133103313331001-2212312203111321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.bond_interface.active_backup` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-2323111001332120-2210132200230122-3321210130001331-3030102133111120-1032231332130300-1201112012200233-0213230001313010-2132033030023132)
- oci.not_managed.node_list.interface_list.bond_interface.active_backup

<a id="canonical-1301123020013113-3023022133221122-1200201331123201-1232231303210330-0012200110131322-3000002010222303-3012233130212222-1000030311220331"></a>

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

<a id="canonical-3022121003032122-2302013221030012-0322223010130310-1333013320010000-2002112302011111-0220101300321313-1321221301233333-1322321102131103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.bond_interface.lacp` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.bond_interface](resources--securemesh_site_v2--reference--group-013.md#canonical-2323111001332120-2210132200230122-3321210130001331-3030102133111120-1032231332130300-1201112012200233-0213230001313010-2132033030023132)
- oci.not_managed.node_list.interface_list.bond_interface.lacp

<a id="canonical-2221320003322110-2111332220102211-1233313001030130-1212323000300113-0230000310232233-1003301313122032-0303101202232133-3312112200113033"></a>

Type: `"object"`. single nested block, Optional.

LACP parameters. LACP parameters for the bond device.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3300132113032221-0331021231033020-2032130132322111-2232310003330313-0211111023223222-2202210121203112-1201131131311233-3023012312121123"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.bond_interface.lacp`

<a id="canonical-2233111002223021-3001210230330311-1122210223030132-1213122112312002-1233030220003331-1132123212113213-2322013033102210-3321302002120302"></a>

#### `oci.not_managed.node_list.interface_list.bond_interface.lacp.rate` property

Type: `"number"`. Optional.

Interval in seconds to transmit LACP packets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3303112303003202-1100003211213131-3310213300320131-2233002220331030-1021200322022103-1022003023201330-0013122010113333-0312111100301001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_client` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.dhcp_client

<a id="canonical-1020212313102332-1220211003123212-1302313212201023-2021332101101121-3200101131130131-0321120002033230-0130201301030122-2301020002111131"></a>

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

<a id="canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.dhcp_server

<a id="canonical-1201103222032033-2112100001231113-0322032131310331-0202232020322232-3123211033031322-0210120321300111-1310233133220001-1323202322120311"></a>

Type: `"object"`. single nested block, Optional.

DHCPServerParametersType.

Additional upstream details:

DHCP server configuration for this interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2333102120323212-0212221002013003-3113013221133210-3203130233200120-1320220032222221-3002321132203111-1030020001232301-2131010013101130"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.dhcp_server`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-013.md#canonical-2303000213222023-1003313320223021-0000001233221111-2102312301113313-0120012231212310-2302100030212233-0103313222013230-3011321001203130): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-013.md#canonical-2211203201033320-1021230321221121-3201212133321102-1300222011300200-2100020221200020-0011123132333311-0101211202100332-1131122312013011): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-0003302121233210-3121203120111122-0300112120120333-2233310221023132-0222033100230231-3303232232133331-2023002233011312-1001211003100120): complete subsection reference.

<a id="canonical-0003210233113221-1312201313113113-3013221100311302-1123002203331020-1033101312232021-1010122001123132-3101232332000203-3311012231033213"></a>

<a id="canonical-2231103313022031-3023313330302331-3230311203123231-2112212200100133-3220321231003122-2203233230323011-2111203031201333-2211001112112133"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_option82_tag` property

Type: `"string"`. Optional.

DHCP option 82 tag.

<a id="canonical-2010320323211131-0322132312201201-2130312131133033-1112113210121010-2013120101103123-3122101222220202-3123320213030110-0300031003100233"></a>

<a id="canonical-3300032033312322-1200030330133200-0222102111013023-1033322333011000-2000313131331221-1120201000112020-0232323030222001-3010132023303210"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Assign fixed IPv4 addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-013.md#canonical-2323022122021103-2102101223130012-2102022220321102-1211031332122313-2232233011233122-1110010000322221-0202321131301223-2230321120333113): complete subsection reference.

<a id="canonical-2303000213222023-1003313320223021-0000001233221111-2102312301113313-0120012231212310-2302100030212233-0103313222013230-3011321001203130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_end

<a id="canonical-0331123121111010-2213230003201122-2122220200332223-1102221201032213-0323113110023231-3102300133123213-3230020200120010-3300013123222201"></a>

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

<a id="canonical-2211203201033320-1021230321221121-3201212133321102-1300222011300200-2100020221200020-0011123132333311-0101211202100332-1131122312013011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- oci.not_managed.node_list.interface_list.dhcp_server.automatic_from_start

<a id="canonical-3023032021221131-0311133023203010-1123313133320203-0122030321321102-2330202313000020-2112322221030211-3000000002023310-0203101201030130"></a>

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

<a id="canonical-0003302121233210-3121203120111122-0300112120120333-2233310221023132-0222033100230231-3303232232133331-2023002233011312-1001211003100120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks

<a id="canonical-2221230133122312-0120123302221301-3011200122101021-3003132303313001-2011132023302202-3332000212101211-2213113023231112-2222233210300120"></a>

Type: `"object"`. list nested block, Optional.

List of networks from which DHCP Server can allocate IPv4 Addresses.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3100113323003003-2022030013121201-2112020031133123-0210210001231013-2032202312202023-1233213332031233-1130033201012320-1210232121111330"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks`

<a id="canonical-3322303011212112-2010111230320301-0110111321231120-1111120322100303-2313321110210002-2232021113203000-0330010323212102-0030112101301101"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dgw_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Enter a IPv4 address from the network prefix to be
used as the default gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3110210221311000-0302330233123233-0320232313031213-1323103011211313-1113312102102112-1010330201130132-3212300330203221-1321231302013011"></a>

<a id="canonical-0210330020022000-0010012221223233-2131010322113220-2101103030303221-3121131321312122-2020312012132331-0223321210123310-2232130222102220"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.dns_address` property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_dgw\] Enter a IPv4 address from the network prefix to be used as the DNS
server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-013.md#canonical-0200333310300330-1000320202011232-0003303120302223-2302020002123003-2111201132023101-0003112123302002-1312221323000222-0130120033313103): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-013.md#canonical-1000321313301121-1122300321222210-2220113001102311-1220033113123101-1232113033121103-2122213012010300-1130310123313322-2102003301333320): complete subsection reference.

<a id="canonical-0013021113302302-1330022300202210-1033300100231100-3032121210032113-3011231320233231-3311112122223300-2331002322212033-2121233011022100"></a>

<a id="canonical-3011001023231111-0131013311010222-0000221210001203-2323330113100213-1313220123211210-2132221130013003-3311201031123212-0201132313112013"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2330302001032000-3202302302233032-2233210021200300-1113302210002330-1112303120120030-1112333210133213-1032122330023033-2202113133131031"></a>

<a id="canonical-3112310300131203-1231332001200322-3312033123233322-3220331303330021-3321021022032333-2203302313321231-2320011230212303-1021023110311132"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [pools](resources--securemesh_site_v2--reference--group-013.md#canonical-0311101131321230-2231011202233201-2320231101233130-2201211221123300-3310012313202213-0222321102133311-2321123020013312-2021300033033200): complete subsection reference.

- [same_as_dgw](resources--securemesh_site_v2--reference--group-013.md#canonical-1212213020113332-3130220130302023-2012031331320013-2221310210300123-1303213213120212-2103320131231311-2300210130023122-2221023310110022): complete subsection reference.

<a id="canonical-0200333310300330-1000320202011232-0003303120302223-2302020002123003-2111201132023101-0003112123302002-1312221323000222-0130120033313103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-0003302121233210-3121203120111122-0300112120120333-2233310221023132-0222033100230231-3303232232133331-2023002233011312-1001211003100120)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.first_address

<a id="canonical-0012301213211113-2212223101021030-3021322231033102-3212222221330230-1101001120001120-3300000033101130-3000303300300010-3011211312302222"></a>

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

<a id="canonical-1000321313301121-1122300321222210-2220113001102311-1220033113123101-1232113033121103-2122213012010300-1130310123313322-2102003301333320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-0003302121233210-3121203120111122-0300112120120333-2233310221023132-0222033100230231-3303232232133331-2023002233011312-1001211003100120)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.last_address

<a id="canonical-0331311203120310-3111311332101302-0121230032201013-3321313031011312-3303212130332031-0203332002020033-1012332202100102-2301231321122322"></a>

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

<a id="canonical-0311101131321230-2231011202233201-2320231101233130-2201211221123300-3310012313202213-0222321102133311-2321123020013312-2021300033033200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-0003302121233210-3121203120111122-0300112120120333-2233310221023132-0222033100230231-3303232232133331-2023002233011312-1001211003100120)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools

<a id="canonical-3123122001231313-0112302221023321-1031303102311101-3223023232033020-0111012031232231-2301113133100323-0001012003112302-3131230103201132"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2013133303102011-3001122012210011-2210201230222212-2333010213123133-3300112301200221-0201213201301131-1133320120221212-2131100120310220"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools`

<a id="canonical-1001212232120312-0211130323022110-0211310021213011-0221211313121102-1212330033103332-2100333232300102-2222122333010001-1120210122022231"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.39 with prefix length of 24, end offset is 192.0.2.186.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0003023001022303-2020002031022303-1220233121233031-1133033023130003-2031213100211300-1011032333132221-0022100030003013-1301201102333000"></a>

<a id="canonical-2110311200200232-2302013002230030-2122003130323000-2001010211321211-0022333021011303-2102200330300213-0131213033002123-2212000122201133"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.exclude` property

Type: `"bool"`. Optional.

Exclude this address range from DHCP allocation.

<a id="canonical-3201313321012202-1201322313320120-1211102223013202-0021020122103111-2031103130303022-0013310131200021-1011312332122201-2301321301311232"></a>

<a id="canonical-3300233211202311-3332220000210313-2031311012112301-3210200132231010-1120302012331010-3023111133123122-3301103312101200-1021311022332300"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IP of the pool range. In case of address allocator, offset is derived based on network
prefix. 192.0.2.173 with prefix length of 24, start offset is 192.0.2.96.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1212213020113332-3130220130302023-2012031331320013-2221310210300123-1303213213120212-2103320131231311-2300210130023122-2221023310110022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- [oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-0003302121233210-3121203120111122-0300112120120333-2233310221023132-0222033100230231-3303232232133331-2023002233011312-1001211003100120)
- oci.not_managed.node_list.interface_list.dhcp_server.dhcp_networks.same_as_dgw

<a id="canonical-1023100232202222-3310222212202122-2322232133032002-3230333000223003-0122331203221303-1221232200310111-2210332220111230-2200320132202013"></a>

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

<a id="canonical-2323022122021103-2102101223130012-2102022220321102-1211031332122313-2232233011233122-1110010000322221-0202321131301223-2230321120333113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.dhcp_server](resources--securemesh_site_v2--reference--group-013.md#canonical-1221133003201201-0332321131311023-1311111300013323-0303333000031331-3012103321021231-2230020311132002-0032320311333013-1321312221001013)
- oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map

<a id="canonical-2333321111010310-2200110020211102-3221000030301111-3131202033101021-0110211222101110-1000113222220303-3330323110100203-2310231012003303"></a>

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

<a id="canonical-0031133331213113-2321123220111300-0023123213013211-3213023331313002-2001001002030003-3330200133312330-1002332232012022-1300301303203131"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map`

<a id="canonical-2121203100330323-2233000133100221-2030012322320001-2033132232321030-1032231131032232-1212300133112032-0220102032223020-0331212230231212"></a>

#### `oci.not_managed.node_list.interface_list.dhcp_server.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Specify static IPv4 addresses per site:node.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0102100102313212-0321212212322322-1201012200002323-1233303312311112-1131103313332322-2000330320123022-2032313010223323-2133321132021131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ethernet_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.ethernet_interface

<a id="canonical-1331322221020112-3112123003101310-3220202221300333-3002220323302120-0200020120111231-1130111130321312-0131111301121101-2301030223330232"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for ethernet interface.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1210303133001313-0011331333221222-3330100100331000-1302300031320121-1030211122211311-1311101213303030-2220333111231013-3131121133211321"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ethernet_interface`

<a id="canonical-2100213331103023-2301233030211200-1113232132301301-1210003103300301-3312220222000303-1131311130101111-3003231103130323-2132313121022230"></a>

#### `oci.not_managed.node_list.interface_list.ethernet_interface.device` property

Type: `"string"`. Optional.

Select an Ethernet device from the discovered interfaces to configure. Once configured, this
interface will be part of this sites dataplane and can participate in the networking services
configured on this site.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1112332013331012-3020100032112322-1003332313321320-1332103300002023-1200130020022012-0000320102202332-2222312011001103-0012031333013123"></a>

<a id="canonical-0013132300121233-3021132232121320-0133211103102323-1233321332113031-0001212110133003-2010233002330011-0302311030302032-2302303002223113"></a>

#### `oci.not_managed.node_list.interface_list.ethernet_interface.mac` property

Type: `"string"`. Optional.

MAC Address. Configuration parameter for mac

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.ipv6_auto_config

<a id="canonical-3023312133111233-2320101113200000-3100223022302010-2112332102322102-2232222310333022-3123003223223111-0010120131310112-2110123232330303"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2302210313131011-3100122313102313-1201130030103212-2310110221032320-2132211320023210-3012331332331122-3211233033333020-2111023101201111"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config`

- [host](resources--securemesh_site_v2--reference--group-013.md#canonical-0300232201233231-3311102131301013-1011000002022230-2132221122001312-0313330112323030-1121311020301131-3112312231131301-2231323122012100): complete subsection reference.

- [router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103): complete subsection reference.

<a id="canonical-0300232201233231-3311102131301013-1011000002022230-2132221122001312-0313330112323030-1121311020301131-3112312231131301-2231323122012100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.host` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.host

<a id="canonical-0200133231312130-3300013302113201-0201233133331010-0223023103031312-3211131111221323-3130310321123320-2202133231021232-1200121312203303"></a>

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

<a id="canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router

<a id="canonical-1333000100200332-2221113221213113-1010023312231323-2012321200231021-0121300301102131-1320000203232233-3211300102012223-2011101132032122"></a>

Type: `"object"`. single nested block, Optional.

IPV6AutoConfigRouterType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2000101021311220-0003230002213323-2112221202333120-3113200333201031-1221022211221102-3311300223221323-1301211332303331-0221030013112221"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router`

- [dns_config](resources--securemesh_site_v2--reference--group-013.md#canonical-3023302031122320-2310232310032103-1122322213113020-3013311300112112-0230300330013032-1232120300232311-0230121112321210-1002002010002131): complete subsection reference.

<a id="canonical-3211033011120100-3303022231313111-2133320222313220-2311000210211212-1310032323010031-0130323221312311-3122112200230121-0212010000212220"></a>

<a id="canonical-3011131012201300-1323120321000230-3210333011033013-3302021211033103-3223221232323113-0231032320332133-2300023232323321-0333230210311332"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.network_prefix` property

Type: `"string"`. Optional.

Exclusive with \[stateful\] Network prefix that is used as Prefix information Allowed only /64
prefix length as per RFC 4862.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [stateful](resources--securemesh_site_v2--reference--group-013.md#canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022): complete subsection reference.

<a id="canonical-3023302031122320-2310232310032103-1122322213113020-3013311300112112-0230300330013032-1232120300232311-0230121112321210-1002002010002131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config

<a id="canonical-3230221001022013-1020322303012130-3000023122113300-2203132232131222-1120030032313022-3030322212332211-3222210320111100-2003322320100312"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsConfig.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1021131010312013-1313113233331100-3303003131200133-3133222223232231-3100031310002220-2213033010201000-2100213231121120-3321120030303221"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config`

- [configured_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2001003130010103-2111201003203033-1020023123110331-0113313310031013-0230303101110330-3231310031112120-2002333030312320-2311011311330323): complete subsection reference.

- [local_dns](resources--securemesh_site_v2--reference--group-013.md#canonical-1210202021032131-3220111232121102-3311330222031033-1131133311032113-3102103030123233-1302102101131231-2231313101330303-1011013023330212): complete subsection reference.

<a id="canonical-2001003130010103-2111201003203033-1020023123110331-0113313310031013-0230303101110330-3231310031112120-2002333030312320-2311011311330323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-013.md#canonical-3023302031122320-2310232310032103-1122322213113020-3013311300112112-0230300330013032-1232120300232311-0230121112321210-1002002010002131)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list

<a id="canonical-0100032030012023-2102133230133130-1213122332112232-0133213310033021-3231321012322233-0332211030232212-2202120013110123-2330210300300223"></a>

Type: `"object"`. single nested block, Optional.

IPV6DnsList.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-2213221021312223-1331321203332011-0312120032303123-2322231031301231-0211332003331112-2023232302212202-1300313201231001-2301031001013312"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list`

<a id="canonical-3222300111331203-2230023331232320-0003133312301202-3323033220201222-1211300232210311-1211310011123333-1231121303033332-1021313330113322"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.configured_list.dns_list` property

Type: `["list", "string"]`. Optional.

List of IPv6 Addresses acting as DNS servers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1210202021032131-3220111232121102-3311330222031033-1131133311032113-3102103030123233-1302102101131231-2231313101330303-1011013023330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-013.md#canonical-3023302031122320-2310232310032103-1122322213113020-3013311300112112-0230300330013032-1232120300232311-0230121112321210-1002002010002131)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns

<a id="canonical-2321103023022122-3000003010022020-1100002212131331-2300011102113320-2020012333321113-2002212200213101-1121030220001323-1102330103003023"></a>

Type: `"object"`. single nested block, Optional.

IPV6LocalDnsAddress.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0031302022010213-1012213122213322-0231200303110200-3230303030303330-1312313300121312-1010031333021311-1332021301122122-0120111323103033"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns`

<a id="canonical-2203023231201131-2310332131320213-0003011110233223-2121031313001332-0300212322002211-2011120321221113-2201232330020103-2311103003220133"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.configured_address` property

Type: `"string"`. Optional.

Exclusive with \[first\_address last\_address\] Configured address from the network prefix is chosen
as DNS server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [first_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2200213312201013-2023010120222230-3211131022113022-3220203322323232-1112312110203021-0303112220031121-3312100201211321-2302122120232132): complete subsection reference.

- [last_address](resources--securemesh_site_v2--reference--group-013.md#canonical-2123322303132021-0320021302131033-1301222220010301-0110310113200200-0013303331033333-0112130333230300-0330202220222233-1030120330103332): complete subsection reference.

<a id="canonical-2200213312201013-2023010120222230-3211131022113022-3220203322323232-1112312110203021-0303112220031121-3312100201211321-2302122120232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-013.md#canonical-3023302031122320-2310232310032103-1122322213113020-3013311300112112-0230300330013032-1232120300232311-0230121112321210-1002002010002131)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-013.md#canonical-1210202021032131-3220111232121102-3311330222031033-1131133311032113-3102103030123233-1302102101131231-2231313101330303-1011013023330212)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.first_address

<a id="canonical-1210011301103111-0012232310032201-3030332201210122-1301010300132331-3110131301220001-1200100131103201-0202213001223210-0301301021031132"></a>

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

<a id="canonical-2123322303132021-0320021302131033-1301222220010301-0110310113200200-0013303331033333-0112130333230300-0330202220222233-1030120330103332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config](resources--securemesh_site_v2--reference--group-013.md#canonical-3023302031122320-2310232310032103-1122322213113020-3013311300112112-0230300330013032-1232120300232311-0230121112321210-1002002010002131)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns](resources--securemesh_site_v2--reference--group-013.md#canonical-1210202021032131-3220111232121102-3311330222031033-1131133311032113-3102103030123233-1302102101131231-2231313101330303-1011013023330212)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.dns_config.local_dns.last_address

<a id="canonical-2313302121003231-0033221101022223-3133002300203130-0310121220213010-2301313303322320-3300212110303033-2120131130332000-3213311131023330"></a>

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

<a id="canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful

<a id="canonical-3110022213200232-3303123303012020-1313313301122112-0301132133200012-2332222201213230-1133103330321321-2321220330011330-1021331111222033"></a>

Type: `"object"`. single nested block, Optional.

DHCPIPV6 Stateful Server.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3221231013030311-1232210012010310-3130113323230223-1322130310130310-3220231002120300-1202123212231113-2120303021233302-2012001200122003"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful`

- [automatic_from_end](resources--securemesh_site_v2--reference--group-013.md#canonical-0222202023001023-2233121322112330-3301022031010323-2312210333133120-2103212331123310-3210020330212333-0013321301120222-1021033103310212): complete subsection reference.

- [automatic_from_start](resources--securemesh_site_v2--reference--group-013.md#canonical-1021111200112112-0002102020202012-1201220031103322-2232313210210021-0002121320310333-3130012301020221-0111313020000310-2222301231113132): complete subsection reference.

- [dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-3110233122300010-0020131032022131-1033210223210033-0112212101210120-2222122103321103-1110221123103323-1310320300011323-2123220221033311): complete subsection reference.

<a id="canonical-0210033230103132-2003230003332100-0013223212120311-1033113300323323-0223223113220021-0312212132112011-1332121222231212-3101213010133213"></a>

<a id="canonical-1001313102033133-3313030320212323-2032003021123003-1111333030031102-2323100321201220-3210003201222122-3123231131032121-2000332231203332"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.fixed_ip_map` property

Type: `["map", "string"]`. Optional.

Fixed MAC address to IPv6 assignments, Key: MAC address, Value: IPv6 Address Assign fixed IPv6
addresses based on the MAC Address of the DHCP Client.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

- [interface_ip_map](resources--securemesh_site_v2--reference--group-013.md#canonical-3232320222301203-2321001231001012-3212122121323001-1133023211320223-0310023021030111-0133210202230221-3020221122300220-2221011210331013): complete subsection reference.

<a id="canonical-0222202023001023-2233121322112330-3301022031010323-2312210333133120-2103212331123310-3210020330212333-0013321301120222-1021033103310212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-013.md#canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_end

<a id="canonical-3000322122223012-2321033102222121-3101331100002213-3122330020033310-0300011212313231-0222300133321331-0302210110202132-3201223012001121"></a>

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

<a id="canonical-1021111200112112-0002102020202012-1201220031103322-2232313210210021-0002121320310333-3130012301020221-0111313020000310-2222301231113132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-013.md#canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.automatic_from_start

<a id="canonical-0311230222211310-0001120030000311-2230311010011132-1020100302323211-1203203011223231-0010010211212200-3210133112100122-0230113013332210"></a>

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

<a id="canonical-3110233122300010-0020131032022131-1033210223210033-0112212101210120-2222122103321103-1110221123103323-1310320300011323-2123220221033311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-013.md#canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks

<a id="canonical-2301011120123130-3322032331230230-0001120021203100-2132110313300011-1020221011122130-0032311113232012-3023310010033322-1030033102113301"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1201203210102310-2002121202203220-1003231012000011-3132113100332222-2130303031323312-1132333210030313-3101221322132223-2133200132212001"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks`

<a id="canonical-0123021302103302-2002022103221323-3130200123003000-2021213002131233-3312212030300213-3113330200213320-3122131123021122-0010100131012012"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.network_prefix` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0201201221122332-2020221130232001-3012330111101302-0002123030321211-1120010331321330-1301233202210321-0002300321113010-3302030010111133"></a>

<a id="canonical-1113120113231100-2310231310330102-1013221332023123-2031302012200032-0111202100232101-2120110021331301-2220013223020000-0032132321332302"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pool_settings` property

Type: `"string"`. Optional.

\[Enum: INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS|EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\]
Identifies the how to pick the network for Interface. Address ranges in DHCP pool list are used for
IP Address allocation Address ranges in DHCP pool list are excluded from IP Address allocation.
Possible values are \`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`,
\`EXCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`. Defaults to
\`INCLUDE\_IP\_ADDRESSES\_FROM\_DHCP\_POOLS\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["EXCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS","INCLUDE_IP_ADDRESSES_FROM_DHCP_POOLS"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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

- [pools](resources--securemesh_site_v2--reference--group-013.md#canonical-1220232320023012-1012201312111001-2230020102023110-1031011001033003-2321310131200020-1002212123123000-0201030113313300-0121000223030231): complete subsection reference.

<a id="canonical-1220232320023012-1012201312111001-2230020102023110-1031011001033003-2321310131200020-1002212123123000-0201030113313300-0121000223030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-013.md#canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks](resources--securemesh_site_v2--reference--group-013.md#canonical-3110233122300010-0020131032022131-1033210223210033-0112212101210120-2222122103321103-1110221123103323-1310320300011323-2123220221033311)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools

<a id="canonical-2301222301013203-2320133200322033-1011211123300012-2113233331203110-0110330330331102-2231000120321202-1131013000122000-1033023333230003"></a>

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3230333123112201-3020311302012022-3002310222023003-2302201112312003-1100332010300222-1211020230000230-1000002023210030-1112222202012022"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools`

<a id="canonical-2131321120233222-0223222012222213-3311132212012212-3210113320331231-1031232213311323-2310103101031321-2030031102113131-0023132231321100"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.end_ip` property

Type: `"string"`. Optional.

Ending IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2212231231002013-2120030002222311-3030202301223113-0112130303113013-1310231032020021-3033101230213320-2120100000013203-1110103330133000"></a>

<a id="canonical-2302330122332021-1332300000100120-2233033133123232-0312022001233310-2123202301001032-0320211111311202-2220301213012001-3131022101322023"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.dhcp_networks.pools.start_ip` property

Type: `"string"`. Optional.

Starting IPv6 address of the pool range. In case of address allocator, offset is derived based on
network prefix. 2001::1 with prefix length of 64, start offset is 5.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3232320222301203-2321001231001012-3212122121323001-1133023211320223-0310023021030111-0133210202230221-3020221122300220-2221011210331013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config](resources--securemesh_site_v2--reference--group-013.md#canonical-1032322310212020-1211122210021011-1032311031223112-1031023223002010-3303011031333123-2233330222322310-2222120213120300-3022011312113230)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router](resources--securemesh_site_v2--reference--group-013.md#canonical-0031201330021023-1330112233210210-0311213313103113-3123232031301111-3003133313023100-3233220100223201-3121332020011033-0133233121023103)
- [oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful](resources--securemesh_site_v2--reference--group-013.md#canonical-1312122030232112-3031100013032102-1222232223122002-0020003332312321-2332012311132231-0300311031221231-1002330010131101-3033313033300022)
- oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map

<a id="canonical-2021102302302232-3202131200102003-1113023223012113-0003222101321112-3222133011213211-0210203122213311-1122320110303202-3023000032032220"></a>

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

<a id="canonical-0112031320002221-1110221203122023-1213000031230120-0213330133021000-2221210023231332-0021332100022003-3202223112000121-1031203323310210"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map`

<a id="canonical-0312213211311322-0111233130231100-2111023100233120-0311112000331021-3222210331103221-0232020123323332-1121233310001213-1233320303300312"></a>

#### `oci.not_managed.node_list.interface_list.ipv6_auto_config.router.stateful.interface_ip_map.interface_ip_map` property

Type: `["map", "string"]`. Optional.

Site:Node to IPv6 Mapping. Map of Site:Node to IPv6 address.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1012121000003222-0223311300012000-1321322100020300-2111313001312211-1132031013310230-2112202303300133-0233031220021033-0230103333112311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.monitor` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.monitor

<a id="canonical-2203232100321300-2212310021011222-1001200323132221-1003201332220033-1001300102310001-2332202100010222-0031221313112101-0001303133031112"></a>

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

<a id="canonical-3111133321323221-0013021133332313-3321010120011200-3303110120132132-3201001233220033-0301213202033310-2021201320202320-1031121200023322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.monitor_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.monitor_disabled

<a id="canonical-0130021023032223-1303010102032000-3120210010021120-1313332022233030-3213011332232131-3101301232130303-1021230210223010-3221213033110002"></a>

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

<a id="canonical-0210321201302320-3311001021020211-2233102321311302-1121221100213323-1113211331021112-1310311322333213-1003111212130221-2300312123233230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.network_option` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.network_option

<a id="canonical-2113200200011021-1021002320210023-1220330120010033-2033213332010310-2121131012033312-2001030121321013-3233200110031023-1033213201212030"></a>

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

<a id="canonical-0220010231213233-1031222021323022-3011332331302101-2300231111023213-0230301323000200-3013301312221001-3233301220331302-3323201221003201"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.network_option`

- [site_local_inside_network](resources--securemesh_site_v2--reference--group-013.md#canonical-0123311130200102-0113303001100223-0112201332202321-0233311332020230-3330332120302223-0120331103023221-1013303233133212-1032123203301311): complete subsection reference.

- [site_local_network](resources--securemesh_site_v2--reference--group-013.md#canonical-3032031203010320-2120122320031012-0110212302200230-2300013120030321-0302231232302220-2322020223011201-2330322020132323-3133112030323100): complete subsection reference.

<a id="canonical-0123311130200102-0113303001100223-0112201332202321-0233311332020230-3330332120302223-0120331103023221-1013303233133212-1032123203301311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.network_option.site_local_inside_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-0210321201302320-3311001021020211-2233102321311302-1121221100213323-1113211331021112-1310311322333213-1003111212130221-2300312123233230)
- oci.not_managed.node_list.interface_list.network_option.site_local_inside_network

<a id="canonical-1313303132031202-1210102301312303-1213120131211030-0222200123130020-3111333111002012-3321010020112220-1111203213221330-1233033013012333"></a>

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

<a id="canonical-3032031203010320-2120122320031012-0110212302200230-2300013120030321-0302231232302220-2322020223011201-2330322020132323-3133112030323100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.network_option.site_local_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.network_option](resources--securemesh_site_v2--reference--group-013.md#canonical-0210321201302320-3311001021020211-2233102321311302-1121221100213323-1113211331021112-1310311322333213-1003111212130221-2300312123233230)
- oci.not_managed.node_list.interface_list.network_option.site_local_network

<a id="canonical-1012201012311301-0223130321200112-3131312020120022-2031023230020030-1322310130221122-3020113103130122-1100330100120103-3211113232302023"></a>

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

<a id="canonical-2130101021120001-2133213211000101-2300112023322331-2020212103120332-3211301312013310-2012020321032321-3231320133123220-0232130330333231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.no_ipv4_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.no_ipv4_address

<a id="canonical-3331212123221000-1100012320301112-3333233103101222-0002103301331332-2200211301022112-1000210012003022-1220012131010132-2110331013203213"></a>

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

<a id="canonical-0300112132322302-3020221123000320-1210301232122310-2222333310111201-2122302132101232-0100013233220120-2121311102120331-0032012132312102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.no_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.no_ipv6_address

<a id="canonical-2310123322022221-3012030123012010-1011203301010231-3130110310200331-3302012311330102-3302110210321133-3211030300312123-3111030013021021"></a>

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

<a id="canonical-3003202120102201-3213210231231301-0332020103012310-2233300011102333-2333023102013220-0232120303110010-0003302110312032-3110123330111121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_disabled

<a id="canonical-3013122310210223-0321202203301201-1000220121201123-0031010211110111-1022113131130123-2101122203001132-0222223332221201-3101221200333021"></a>

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

<a id="canonical-3033302102222012-3000301233233301-2213032233212122-3131331003110300-0212132220003022-3322232313013302-0133322301103221-2233221222132002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.site_to_site_connectivity_interface_enabled

<a id="canonical-2203030120313200-2233212012112213-1012012000231310-3201222130322223-0100333020313211-3232233102231022-0321200220311310-2131003011212003"></a>

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

<a id="canonical-3010222121020100-3200320332020122-3001031023010002-2031200013313032-0301001232022221-1020321212302012-0123013001102011-3102003213111310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.static_ip

<a id="canonical-3030102011110303-2210323113321300-0102323022000011-2003302312302333-2102023333121100-1120022020332123-3023331113023020-3311323320133212"></a>

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

<a id="canonical-2212311301212131-1021110012303001-3203033033033033-3332330131223121-3313323132001203-1000120312030222-0223012320123002-0130112230302303"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.static_ip`

<a id="canonical-2002310303320031-1233001320101111-3302101221031000-2023003231312303-3321213002231200-1120023021321102-1302201002230320-3321032200031213"></a>

#### `oci.not_managed.node_list.interface_list.static_ip.default_gw` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3330232023101300-0002101311302132-2321120112301300-2112223303021230-2310001103332321-3002213112030110-1000223332303021-2023030132303001"></a>

<a id="canonical-3032120213313010-1111220130111233-2221123303210131-2313221000021033-1100021223302012-1012111001211323-0103120313221223-1312322333233332"></a>

#### `oci.not_managed.node_list.interface_list.static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-1313000223321202-1113313203331113-1333132223301133-2332123103131310-0122321011113123-2321312303101201-0101211023331203-0302233322131311"></a>

<a id="canonical-1301321131111000-2211322302301112-2030301001013311-0331201001020102-1311011332101113-1132123023200030-3111320212320032-0333300220312322"></a>

#### `oci.not_managed.node_list.interface_list.static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-0022320311200201-2022123231133030-1003100132020003-2023121231320123-1130203301010300-0010101333111201-1013111001221323-3131232313332233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.static_ipv6_address` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.static_ipv6_address

<a id="canonical-3303121222313003-2202120211102200-0311112002133022-1133302133101232-3010332120201220-3202202333100023-3122021233023312-1202210122233332"></a>

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

<a id="canonical-0023333120200000-2022003102110332-2331300330022000-3201230002100033-3010022122001133-2020230300020011-3120213032323110-1000302110321003"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.static_ipv6_address`

- [cluster_static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-2233311221131210-2331210130100231-3212122312233112-1131211101101023-2223300303110101-0303301133010110-3102012113013020-0223201001212122): complete subsection reference.

- [node_static_ip](resources--securemesh_site_v2--reference--group-013.md#canonical-2300131112333301-2132102311301102-3123031023230113-1133122303021010-1221112330120023-2300123130311001-2212123102031203-3030333112030111): complete subsection reference.

<a id="canonical-2233311221131210-2331210130100231-3212122312233112-1131211101101023-2223300303110101-0303301133010110-3102012113013020-0223201001212122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-0022320311200201-2022123231133030-1003100132020003-2023121231320123-1130203301010300-0010101333111201-1013111001221323-3131232313332233)
- oci.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip

<a id="canonical-1301122202202033-2013031121332221-1002001213122220-3013212021113031-3022203210323002-1232021333020333-0222002003312131-3102030321211323"></a>

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

<a id="canonical-0310300202310210-2322223111213021-1010130323131312-1102333303202300-2213310002320131-3331313121201022-2020311220232302-1232123000022221"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip`

<a id="canonical-0211013013321101-0010033321230012-2032003103320001-2120013021031331-3033220323101303-2023032330032331-3231101110312032-1020331330000303"></a>

#### `oci.not_managed.node_list.interface_list.static_ipv6_address.cluster_static_ip.interface_ip_map` property

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

<a id="canonical-2300131112333301-2132102311301102-3123031023230113-1133122303021010-1221112330120023-2300123130311001-2212123102031203-3030333112030111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- [oci.not_managed.node_list.interface_list.static_ipv6_address](resources--securemesh_site_v2--reference--group-013.md#canonical-0022320311200201-2022123231133030-1003100132020003-2023121231320123-1130203301010300-0010101333111201-1013111001221323-3131232313332233)
- oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip

<a id="canonical-0022312120223002-0102220112101313-1321103213002221-1212213302310312-0332102103021013-1200311211102101-1032213101230200-1122023222020212"></a>

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

<a id="canonical-1123112122101330-3132002100232233-2131312122333121-1303211001113310-1101232102330020-2013132000020331-0002303213323310-3010230133120130"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip`

<a id="canonical-3222233232003201-3033213020020300-2112332322031321-2311311000300133-0020301321030110-2231130133122122-1032332013330313-2110131213232000"></a>

#### `oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.default_gw` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3223101203110023-3031223030133012-0332030210321133-3202301231032030-1021022111222110-3232222133313230-2131200030112323-1111103220113330"></a>

<a id="canonical-3220202220000001-2321120321220133-0020313303032123-0031112100312103-2033301200013122-0021220313231221-1230001303023210-1311312200112321"></a>

#### `oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.dns_server` property

Type: `"string"`. Optional.

DNS server address for the static interface configuration.

<a id="canonical-3222203122333101-2023230031112123-2330103302211302-3300221002123010-3330200213112322-0023321211320331-0233130213201330-0012203002122001"></a>

<a id="canonical-3000223123020211-0233301230131012-2021132003110030-1301332212032021-3102213123232212-2221321020311100-3201321120101310-1100200231121332"></a>

#### `oci.not_managed.node_list.interface_list.static_ipv6_address.node_static_ip.ip_address` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-3122131201212313-0011330320013131-1213002113003132-3301222030302110-1031321231302011-3201100110111221-1332320023322201-1220211211132031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `oci.not_managed.node_list.interface_list.vlan_interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [oci](resources--securemesh_site_v2--reference--group-013.md#canonical-1332211122203232-2032213012031311-0312131000222131-0313233003113010-1010303203211031-2310121002301201-0221000203033232-3212021221312311)
- [oci.not_managed](resources--securemesh_site_v2--reference--group-013.md#canonical-2223032130230102-1220212221102301-3023000223010021-0230203120122230-0313303332010310-0331213001120201-2311110301221120-0003012220112221)
- [oci.not_managed.node_list](resources--securemesh_site_v2--reference--group-013.md#canonical-1131312101200303-3011212302111200-1130300211000103-0333102120321213-0333101301020310-1102102330222123-3123011122002112-2130231322222221)
- [oci.not_managed.node_list.interface_list](resources--securemesh_site_v2--reference--group-013.md#canonical-2030321032330302-1313133321122210-0010211101110131-1031132303202202-2113132123303120-2231223003113312-2113012020212120-0302301211121311)
- oci.not_managed.node_list.interface_list.vlan_interface

<a id="canonical-0331202132010330-3121011022222230-3031303123112212-0021301202133033-2031212223200031-3321301101220221-0202222220000230-1110113220202320"></a>

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

<a id="canonical-1110230030012322-3312132101230013-2201123122232130-0232322123103333-0321102332220010-1231112111220023-2221001010300111-3222023200301333"></a>

### Direct properties for `oci.not_managed.node_list.interface_list.vlan_interface`

<a id="canonical-2112110323313222-2311313110310030-2212101003321121-3030102110100001-3210031112201133-0021232313133311-0230303001010301-2110113120111211"></a>

#### `oci.not_managed.node_list.interface_list.vlan_interface.device` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-1311212212113301-0330131011013030-1031221310013312-3003001101113133-3013102312212232-1211232223002131-3210302132330113-3021312311201322"></a>

<a id="canonical-0200303130303311-0222120120333003-1201120033100131-3332330013211331-2101100220300321-0122030000231221-2010122201022233-1332212210203021"></a>

#### `oci.not_managed.node_list.interface_list.vlan_interface.vlan_id` property

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2001312130230101-3132012012311122-0230112101013213-1333302011302001-3223331311230002-0320013221303203-3122311112320113-1320320233110001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- offline_survivability_mode

<a id="canonical-3203213033022033-0233302332212131-2202111210010301-2132011323032220-2130030113131311-3211113010120211-2223111003303213-2022320213320300"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
