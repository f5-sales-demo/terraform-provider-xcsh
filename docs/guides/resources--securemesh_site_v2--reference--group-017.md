---
page_title: "xcsh_securemesh_site_v2 reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 reference."
---

# xcsh_securemesh_site_v2 reference

<a id="canonical-3120103210223322-2103113313030320-3023220333323022-3321031033233303-2100013013130210-2323010002121033-2233102323323210-3013223201131212"></a>

### Direct properties for `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list`

- [interface](resources--securemesh_site_v2--reference--group-017.md#canonical-0203013210220203-3011210203020001-2303302000233320-1320323230302020-3330201113321322-2033213313131303-1011320032330120-2302322033112212): complete subsection reference.

<a id="canonical-3121100230233331-2311122013110013-3122103200332133-3012011121310232-0213320313231130-3120313131202200-1013200332000131-1102211221033101"></a>

<a id="canonical-1233331113120313-3131013232320310-3300330313003110-2131202233311300-0021033220002030-2032122121002033-0322112323320100-0311200313213123"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.node` property

Type: `"string"`. Optional.

Node. Node name on this site.

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

<a id="canonical-0203013210220203-3011210203020001-2303302000233320-1320323230302020-3330201113321322-2033213313131303-1011320032330120-2302322033112212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- [segment_vrf.segment_config](resources--securemesh_site_v2--reference--group-016.md#canonical-2122023013000013-1200000021121230-0300030303303332-2331213103203310-1130023002221113-1310013101300013-0331031310323112-2113030203010021)
- [segment_vrf.segment_config.static_v6_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-3033010013033110-0021330003210110-2302332112111223-0003033003103232-3103322320123032-3120110130221120-0031101003312210-3212313313132022)
- [segment_vrf.segment_config.static_v6_routes.static_routes](resources--securemesh_site_v2--reference--group-016.md#canonical-2010113032000202-2011033201031200-0132233002323210-2233011232331332-2323103300012222-3132023113312032-2221322332032133-3022311100211312)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface](resources--securemesh_site_v2--reference--group-016.md#canonical-2211232301311022-3313233111032330-0111001103333230-2231233131203201-1313200131121321-1203110002132233-2330010020301022-0032021102132233)
- [segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list](resources--securemesh_site_v2--reference--group-016.md#canonical-3102211231023322-1331330123223021-1121220230003322-2000032102323123-3130102230302321-2030201010200002-0033321100332321-2123230332222031)
- segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface

<a id="canonical-3322023112000201-2313113111201310-0313001322311311-0120003322112233-3303130321322200-0000300211203011-0112213210021211-1310320233302032"></a>

Type: `"object"`. list nested block, Optional.

Interface. Interface reference on this node.

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
interface {
  # Configure direct properties listed below.
}
```

<a id="canonical-3321012300003202-0012330330102032-2112022120010203-3110032321011331-3201021130022323-2233201122110031-1002022333002211-3031102213330033"></a>

### Direct properties for `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface`

<a id="canonical-2100233331311301-0110333311333212-2202323121012221-2213333311112022-2332201321002210-1231330033331010-3122020010233003-0011202113011010"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.kind` property

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

<a id="canonical-1010101320213201-3331012100013031-3210303031020120-2322013231132032-3023023102301222-0233203222121100-1002230231112031-0202302213313120"></a>

<a id="canonical-1310022131121032-0203311121212020-0021020203131230-0311112132212032-0222330232210120-2110230323013112-1033023302210320-3210100120013021"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1032130221022210-2112002123323033-0210102321030320-0300213132022121-3203002112101303-1220112222121230-1331202212022330-1131230001322032"></a>

<a id="canonical-0110202200133131-2030311322112213-3030101022113221-1310032122000203-0013230200203232-1010033133223310-1220130012330111-1233132111121203"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-1103321233213300-0122232333023303-1101310320111032-0300010210103003-3213021331010120-1113012231032100-2132221012223201-2310223010020321"></a>

<a id="canonical-1303001103002331-2120131231232011-2030033231333301-2331320322022021-3322123313103101-1330021202011031-3001333233311130-1321013123101213"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.tenant` property

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

<a id="canonical-3330031110022033-0303011021331033-2122103033212232-2113000120323110-0310132130020233-2223203203211121-1220100013132123-3021000131323330"></a>

<a id="canonical-3202202222232232-3103033320031210-3330300233111012-2213133013223300-0311213210332310-2033310133102122-3100122322331103-1302100123220121"></a>

#### `segment_vrf.segment_config.static_v6_routes.static_routes.node_interface.list.interface.uid` property

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

<a id="canonical-3110213102100123-2213032002131131-2001112022032210-1133200211232232-0332033311333113-3210132330223010-0302230001102032-0313121312022203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `segment_vrf.segment_network` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [segment_vrf](resources--securemesh_site_v2--reference--group-016.md#canonical-2013033031131331-2213313223020200-2112133031330202-2303320232020100-1213211121321303-2130010121132312-0000300320202222-1331023001013310)
- segment_vrf.segment_network

<a id="canonical-3310010003302231-0220201010012300-3320220012120031-3000121322122220-2012220131121331-3103030203003100-2300222032222303-0110330120011031"></a>

Type: `"object"`. single nested block, Optional.

This type establishes a 'direct reference' from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name for public API and Uid for private API This
type of reference is called direct because the relation is explicit and concrete (as opposed to
selector reference which builds a group based on labels of selectee objects)

Terraform syntax:

```terraform
segment_network {
  # Configure direct properties listed below.
}
```

<a id="canonical-3203120113133121-1122321120021231-3230101132221210-2320332001023103-0121232330112232-3000300212200322-0211003113111001-3231221230103012"></a>

### Direct properties for `segment_vrf.segment_network`

<a id="canonical-2033202132013303-1013202332031231-1321222322111002-3300033213122211-2322230023233210-3321020032003313-2001313101302020-2103101312021221"></a>

#### `segment_vrf.segment_network.kind` property

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

<a id="canonical-0321313113030030-3010203112321201-3231223030331300-2223030312122210-1213113111102003-1300003311302200-2003201303002333-1202231302310230"></a>

<a id="canonical-1232100032123003-2200213333122112-1311111120200330-2133020100002232-3101100302121310-1102121022202302-3013221113211213-3202122021131211"></a>

#### `segment_vrf.segment_network.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3101310313100030-1203200302112323-3211001202322101-3032113120333230-3120222211133021-1132023202112201-3001212200130323-1232031003221230"></a>

<a id="canonical-0011332003033231-3332123221331110-2322301123101211-1323311310023021-2331330200223130-2320120303112233-3220111011220210-1303020222002011"></a>

#### `segment_vrf.segment_network.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-2210221032033312-1101011031200233-3303110200001122-2021231123111023-3322020012101220-2332110031213112-1001011122320003-3210323011123332"></a>

<a id="canonical-0200213011301001-0133202100103222-2330310122303010-2030021202312002-2000032100022320-1302131311031002-3323112212100300-3233213000230322"></a>

#### `segment_vrf.segment_network.tenant` property

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

<a id="canonical-2120020210210311-2232103303331021-1233100313021232-0233013122230231-2031010113111232-3121232103100331-3310122010022133-0131011020310232"></a>

<a id="canonical-0321201012330321-1013313020100332-0320121003021102-2233321031133211-1000003223120231-3010211101032121-1033032013103333-2103102023012302"></a>

#### `segment_vrf.segment_network.uid` property

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

<a id="canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- site_mesh_group_on_slo

<a id="canonical-3212222333120111-2230203121211023-0313133130013232-2310301323102000-1222112001023100-3302312232123100-3211022311121112-3323310000112231"></a>

Type: `"object"`. single nested block, Optional.

Select how the site mesh group will be connected. By default, public IPs of the control nodes of the
site will be used.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_site_mesh_group",
    "site_mesh_group"),
  validators.ConflictingObjectAttributes("sm_connection_public_ip",
    "sm_connection_pvt_ip")}
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
  "x-ves-oneof-field-site_mesh_group_choice": "[\"no_site_mesh_group\",\"site_mesh_group\"]",
  "x-ves-oneof-field-site_mesh_group_ip_choice": "[\"sm_connection_public_ip\",\"sm_connection_pvt_ip\"]"
}
```

Terraform syntax:

```terraform
site_mesh_group_on_slo {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122203213000110-1030200021232110-2212000032322002-0020221023122303-2023022130303313-0000133020331310-0321310120002102-3301030203022101"></a>

### Direct properties for `site_mesh_group_on_slo`

- [no_site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-3213012221100220-3100021321210121-2201100000112121-1010010332201323-2131222301213210-0230011302111310-2203212022110131-3230021003021110): complete subsection reference.

- [site_mesh_group](resources--securemesh_site_v2--reference--group-017.md#canonical-0221223233130001-2133301130312131-3202122030310102-1130300022002301-1220302213132021-3013131201321033-0313221122110012-2023311312012002): complete subsection reference.

- [sm_connection_public_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-2231331302323330-0213212301333112-3033310021003321-0132133122112030-3031132203021330-2213303033233103-3022122303110220-1100123022011033): complete subsection reference.

- [sm_connection_pvt_ip](resources--securemesh_site_v2--reference--group-017.md#canonical-0132123210131012-0313113211220131-2333303223213133-3101032222221100-3003000133203332-1102020013223020-3000112020112201-0320222130320130): complete subsection reference.

<a id="canonical-3213012221100220-3100021321210121-2201100000112121-1010010332201323-2131222301213210-0230011302111310-2203212022110131-3230021003021110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.no_site_mesh_group` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.no_site_mesh_group

<a id="canonical-3012320000213211-2212320133121022-2221133301312211-3332132310100323-0111300311010302-0303111301210210-0133002112233011-2310311013222031"></a>

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
no_site_mesh_group = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0221223233130001-2133301130312131-3202122030310102-1130300022002301-1220302213132021-3013131201321033-0313221122110012-2023311312012002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.site_mesh_group` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.site_mesh_group

<a id="canonical-3203031101211003-3103032221321300-2122232301113121-3133013022323300-1213213302220000-1203031332000001-1112310133312033-2231210233022121"></a>

Type: `"object"`. single nested block, Optional.

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
site_mesh_group {
  # Configure direct properties listed below.
}
```

<a id="canonical-2121130221133112-0303233222200030-0031012010022211-2320121302301333-1120230230230032-2313122120203203-3213211220101222-1323131010311313"></a>

### Direct properties for `site_mesh_group_on_slo.site_mesh_group`

<a id="canonical-2223102120030000-0131103201300122-1033300000031313-0202212301101212-1101123130320021-3022012032221011-3132121000331222-1033220010201002"></a>

#### `site_mesh_group_on_slo.site_mesh_group.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2121221023121303-0122111123030112-2122020210100030-1011013100101101-1000213213310013-0330122330213320-0123330002202211-2222221211203320"></a>

<a id="canonical-1121003010030223-3202232011201232-0222203001130223-3031033000103120-2100002031212200-1021103331201130-1020313220231101-2013130233301321"></a>

#### `site_mesh_group_on_slo.site_mesh_group.namespace` property

Type: `"string"`. Optional, Computed.

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

<a id="canonical-0211103112030103-1111230213022322-0301321301101312-2111110122120121-3032033211220021-3322031320200321-0333111131303310-1201330002211002"></a>

<a id="canonical-1211200111011301-1321030111233222-1220122002032103-2131220010023122-0220221230001330-3303223230221230-0223321220203113-3210003100332202"></a>

#### `site_mesh_group_on_slo.site_mesh_group.tenant` property

Type: `"string"`. Computed.

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

<a id="canonical-2231331302323330-0213212301333112-3033310021003321-0132133122112030-3031132203021330-2213303033233103-3022122303110220-1100123022011033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.sm_connection_public_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.sm_connection_public_ip

<a id="canonical-1320133300201220-1123311331301032-0102212203331001-2301323323000101-3211130111200033-0131303000303111-2223310312130131-1102102111032333"></a>

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
sm_connection_public_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0132123210131012-0313113211220131-2333303223213133-3101032222221100-3003000133203332-1102020013223020-3000112020112201-0320222130320130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_mesh_group_on_slo.sm_connection_pvt_ip` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [site_mesh_group_on_slo](resources--securemesh_site_v2--reference--group-017.md#canonical-1132320030201122-0222122012131332-0210322212020301-1032300301203333-3202102311232202-3000202023203231-1030101202313102-0300121202300311)
- site_mesh_group_on_slo.sm_connection_pvt_ip

<a id="canonical-3202013112113133-2210320113101033-0220010013032102-2330210232322323-1211000232032301-0031130021122002-3223313210323102-2021222332203323"></a>

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
sm_connection_pvt_ip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- software_settings

<a id="canonical-2021303230011013-3122113030331033-1131000201101130-1223003121320230-0201302021311101-2331313102113102-1032022110212020-0233323113321131"></a>

Type: `"object"`. single nested block, Optional.

Select OS and Software version for the site. All nodes in the site will run the same OS and Software
version. These settings cannot be changed after the site is created. This block is a create-only,
write-only input; changing it replaces the resource, and refresh preserves the configured value
without claiming XC observed it.

Receipt-pinned upstream constraints:

```json
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
software_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3112321100201220-3003031031223103-2112331001020123-1330221011301012-0110212011330033-1131030213121312-1320030121313200-2130331322301031"></a>

### Direct properties for `software_settings`

- [os](resources--securemesh_site_v2--reference--group-017.md#canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302): complete subsection reference.

- [sw](resources--securemesh_site_v2--reference--group-017.md#canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300): complete subsection reference.

- [waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031): complete subsection reference.

<a id="canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.os` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- software_settings.os

<a id="canonical-1310220101111131-2102110111201321-3111032230233210-2332211103222102-2023133321023231-1230021221222323-0332120102122103-2113302330301320"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_os_version",
    "operating_system_version")}
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
  "x-ves-oneof-field-operating_system_version_choice": "[\"default_os_version\",\"operating_system_version\"]"
}
```

Terraform syntax:

```terraform
os {
  # Configure direct properties listed below.
}
```

<a id="canonical-2233021331102033-2301001132012203-1011112101302013-2231002323013311-0101323201100202-1311020301021112-2000031321022233-0322023220222203"></a>

### Direct properties for `software_settings.os`

- [default_os_version](resources--securemesh_site_v2--reference--group-017.md#canonical-1030220122022302-1232023121110020-0231001010323331-1333120012030332-3301310103110120-1230101212120030-2013303130333212-2131233130011232): complete subsection reference.

<a id="canonical-3110011123301322-3023331103222310-2201322123101120-1012002231201022-1221220333012123-3010012213133323-0123020021100230-1001213202111013"></a>

<a id="canonical-0101031110311132-2200210102311223-2200311103021033-0023130201101310-3320033003223032-0020100101001312-0202212032000203-0111021212103303"></a>

#### `software_settings.os.operating_system_version` property

Type: `"string"`. Optional, Sensitive.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-1030220122022302-1232023121110020-0231001010323331-1333120012030332-3301310103110120-1230101212120030-2013303130333212-2131233130011232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.os.default_os_version` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.os](resources--securemesh_site_v2--reference--group-017.md#canonical-0322333322002322-3032123013133021-1301213130312232-1022000103110120-2323021100313310-3330102222002010-2330313032203300-3203333132213302)
- software_settings.os.default_os_version

<a id="canonical-3111110002212210-0102322331210330-1001201020332230-2031031302203002-1103221120211323-0313103112310301-2333301331323322-0133130222221312"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
default_os_version = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.sw` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- software_settings.sw

<a id="canonical-1110113213023313-1102311132110301-3323021130312011-2100333211122331-2120033223230133-1031233211311033-1120232111230220-2113330201012231"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_sw_version",
    "volterra_software_version")}
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
  "x-ves-oneof-field-volterra_sw_version_choice": "[\"default_sw_version\",\"volterra_software_version\"]"
}
```

Terraform syntax:

```terraform
sw {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130123111221101-1330322023212120-1123113330221033-0000120011000333-2110013030001031-2300203001200231-3111221131233102-2321320332113132"></a>

### Direct properties for `software_settings.sw`

- [default_sw_version](resources--securemesh_site_v2--reference--group-017.md#canonical-2321301203002322-0331112232333333-3333211033120312-2001312211012223-1020013013113030-0021212302302320-0122320312112021-0031302110201003): complete subsection reference.

<a id="canonical-2210323000122022-3231320233210302-1023302211020032-2131233123030233-1220000202132122-2031103030233111-0033223301101020-1103222333232202"></a>

<a id="canonical-2223000331030221-1211001213031230-3120122101002202-2230033232010133-0031030010323212-2201212023230223-1301123230330102-0022123122302122"></a>

#### `software_settings.sw.volterra_software_version` property

Type: `"string"`. Optional, Sensitive.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 20,
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2321301203002322-0331112232333333-3333211033120312-2001312211012223-1020013013113030-0021212302302320-0122320312112021-0031302110201003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.sw.default_sw_version` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.sw](resources--securemesh_site_v2--reference--group-017.md#canonical-1022011000033201-3130302031003303-1313330103003011-3220313123103210-0100001310330232-2023212030132102-2102021032031301-3311233103031300)
- software_settings.sw.default_sw_version

<a id="canonical-3203220002133302-1220333322120121-2112012331302320-2200100123110230-0121033011123310-2120313011200210-2110132302330203-3032220203131333"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
default_sw_version = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.waf_signatures` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- software_settings.waf_signatures

<a id="canonical-2123333211201220-1021203012100232-0221321230021311-3110233330322110-3200332000212330-1013002101200320-2333132212230221-2013331233320202"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("automatic",
    "manual")}
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
  "x-ves-oneof-field-signatures_update_mode_choice": "[\"automatic\",\"manual\"]"
}
```

Terraform syntax:

```terraform
waf_signatures {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210233033330120-1322212123000110-2201330120211100-3222102213011203-3210002020332110-2222332222120202-1330110301103231-1323033332012123"></a>

### Direct properties for `software_settings.waf_signatures`

- [automatic](resources--securemesh_site_v2--reference--group-017.md#canonical-1321112103302132-0311330003120012-3202222023122102-2222222333321323-2313231200132032-0001101331002033-1112022212311112-3010310111203233): complete subsection reference.

- [manual](resources--securemesh_site_v2--reference--group-017.md#canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023): complete subsection reference.

<a id="canonical-1321112103302132-0311330003120012-3202222023122102-2222222333321323-2313231200132032-0001101331002033-1112022212311112-3010310111203233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.waf_signatures.automatic` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- software_settings.waf_signatures.automatic

<a id="canonical-0212033110121121-2221203031112002-2220323311203110-0030030011131000-3111101323112121-0232233032330133-3310101210113132-0223222330203000"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
automatic = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1330322203032023-2120032122001101-2211212031112120-1122031202300301-2112301323133203-0101313100113103-1321120100300011-3032303101111023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `software_settings.waf_signatures.manual` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [software_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-0022231303311111-0320032323222122-1103000112000213-1212111100333213-2111302213021020-3313130203211320-3301133033221310-2322002021330133)
- [software_settings.waf_signatures](resources--securemesh_site_v2--reference--group-017.md#canonical-0010013231031330-1121200000030022-2121223331000232-0001101321233133-3012232310002232-1123032033001322-2001130231000322-2233301113322031)
- software_settings.waf_signatures.manual

<a id="canonical-3221332221221101-0102121321103320-3211111203002301-0020222103103013-0212322131112023-2102311100231210-3223200221222230-3321213001123032"></a>

Type: `["object", {}]`. Optional, Sensitive.

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
manual = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131133220012220-0222312200312320-2332131230210133-2323001331232302-0312012222021021-0213223112001210-2110313123231232-3312321032110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- timeouts

<a id="canonical-0123232131001300-0011323332301222-1331300031113011-1313301030310003-1101033221200021-3320311312021033-0002021220231213-3010101123210223"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021321101013023-0113201331010321-1223313213301203-1102122020330113-1301201111333122-0100102230121200-1102022310221121-2123111100031331"></a>

### Direct properties for `timeouts`

<a id="canonical-1110023132023112-1100103232120003-2233121122233201-3102321030123121-2111201002203301-0010022121120230-1321131213311030-0102133312300210"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-3130133011022332-2230123123000100-2021330331212132-1122103310110301-1102211033201131-1022213332002323-3032001030113000-2101002021101032"></a>

<a id="canonical-2210201022013022-3010101101010213-0321203021022011-0303021113330030-2223120120201123-0220122021311112-1320031202301311-2100203010002121"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2213013133300101-2020313200320200-3012022211113211-2233211130000331-2213130010313222-2230123120212320-3231211022000323-1002223123231322"></a>

<a id="canonical-2121232011321202-2203023312230112-3303230112222221-1230213322110100-3310010221010320-2100323100130101-3223312013132232-1113332030013223"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-0003313021213233-2220003330111102-0010200221002312-3332230211002222-1001101301312301-2030130303103123-2203003122030000-2303003002323021"></a>

<a id="canonical-2203113031000212-0112123303133102-1223122313012112-3113313200030100-2121231321120220-1302230133321100-0120032323022320-0122330223201332"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- upgrade_settings

<a id="canonical-1012300130023020-3121103011122011-3200231111011031-1132302032012101-1022231300310311-1322000133202211-3332113131103223-3020202230200002"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for upgrade settings.

Additional upstream details:

Specify how a site will be upgraded.

Receipt-pinned upstream constraints:

```json
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
upgrade_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-2122103030011130-1130031223203121-2310123120110211-3103203033131122-3303003100310103-2001220130103132-1220332013012012-1220133013020132"></a>

### Direct properties for `upgrade_settings`

- [kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300): complete subsection reference.

<a id="canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- upgrade_settings.kubernetes_upgrade_drain

<a id="canonical-3220131333120013-2202322022131000-1000010222103213-2013222300320010-2122210011001020-3001103121202113-0103222211322232-1121002031102322"></a>

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

<a id="canonical-0101321200011100-1112313033210132-2002311330030222-2231031201011003-0200110130023201-3231111111102333-3212102122231321-3220231320113002"></a>

### Direct properties for `upgrade_settings.kubernetes_upgrade_drain`

- [disable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3222130313222010-2113200000020330-2312232021233113-1330033220323210-2200321101300231-3032233121210023-1120230112112230-0000123202210303): complete subsection reference.

- [enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331): complete subsection reference.

<a id="canonical-3222130313222010-2113200000020330-2312232021233113-1330033220323210-2200321101300231-3032233121210023-1120230112112230-0000123202210303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- upgrade_settings.kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2321012113022223-1013233322221032-2113023113030230-2001203212122013-2011023203033332-0222230313023003-1203131123300310-2120212331312303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable upgrade drain.

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
disable_upgrade_drain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-1230100330002323-0110122001332122-1021203213231021-3221200310110203-1033300111021200-0012331102201322-1221222030211313-0331001023233011"></a>

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

<a id="canonical-0133031003022130-3230200220200113-2330220221333030-2333221222232301-3203113332112223-1302303123103101-3310331200132222-0203332033230330"></a>

### Direct properties for `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain`

- [disable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-3112020123021302-3203002223103220-2301221002322203-3123211132101102-1030100001330230-2033323301101002-3032133031210311-0220323010030202): complete subsection reference.

<a id="canonical-0301130333233111-3220210110033323-1111202321111333-0031212120321321-3021121210313002-0211300001100023-3020301101123012-1130123320001132"></a>

<a id="canonical-2300030022123013-0302213232301130-2200232312332323-0011123223011200-0303021232231011-2330122120132201-3332120110100121-0230300303301200"></a>

#### `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` property

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

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

<a id="canonical-2032210002330002-1330002321132010-2012221131221013-1303010103223013-3331213222211230-3333021232103212-2020101321123123-0202030321330320"></a>

<a id="canonical-3030323100203231-2331330210332002-0211130222122210-3020030121132113-1030221120121231-0001203102300022-2322000131123013-1301222122000330"></a>

#### `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` property

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-3320032100232312-3031311210022132-0222313022022211-0331213102230030-0300103210203120-3100122111030010-0031300323002322-3101233312001331"></a>

<a id="canonical-3332212003212102-3003032232202102-0313332211300031-0130112022000133-3230012300030002-3022221233111122-0300131012230123-3233013112023101"></a>

#### `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` property

Type: `"number"`. Optional.

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

- [enable_vega_upgrade_mode](resources--securemesh_site_v2--reference--group-017.md#canonical-2130101021010320-3202311123201221-3231221223003212-1033030023012101-2021121332113111-2233131230332330-3131002210023130-0321000131233100): complete subsection reference.

<a id="canonical-3112020123021302-3203002223103220-2301221002322203-3123211132101102-1030100001330230-2033323301101002-3032133031210311-0220323010030202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-3102022031121003-1323013230331033-2222213102030110-2003103132303331-0312302233301213-1113000221111210-3201010320331001-0021110130020230"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130101021010320-3202311123201221-3231221223003212-1033030023012101-2021121332113111-2233131230332330-3131002210023130-0321000131233100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [upgrade_settings](resources--securemesh_site_v2--reference--group-017.md#canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223)
- [upgrade_settings.kubernetes_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-3003201331303031-0002000113122201-0302221010301213-1101030300000222-2123213100310030-0033201302130020-1110302002210301-3122103333230300)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](resources--securemesh_site_v2--reference--group-017.md#canonical-0213021020130133-1313221201012121-1030022210233023-0112033031122113-3032132103203320-0310101221001123-1001101132203311-2021023300213331)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-1013033301121202-1312030331131013-1222303323220300-0002030002232133-0131001133211203-0322110002310211-3132322030130112-1233320312203302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

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
enable_vega_upgrade_mode = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- vmware

<a id="canonical-0013121331333101-0110031232203120-1332123300101323-0213331300203303-3102322233201100-0101131220111322-2213230000203031-1133112232132230"></a>

Type: `"object"`. single nested block, Optional.

VMware Provider Type. VMware Provider Type.

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
vmware {
  # Configure direct properties listed below.
}
```

<a id="canonical-0201033311323312-0112303311222322-1013233210001021-3303231332110320-2122111213120300-3233202232022202-0102310321012232-3031310122312312"></a>

### Direct properties for `vmware`

- [not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133): complete subsection reference.

<a id="canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- vmware.not_managed

<a id="canonical-3212320032322213-2211221212101101-0030112233232122-3121231223331210-1111233221202320-1300010012113133-0321233132100213-3211203131021331"></a>

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

<a id="canonical-0232110020133112-1030322100231002-0323310122210123-3122010001310110-3000330202000213-3322113111123221-1321003100113211-1021321231220111"></a>

### Direct properties for `vmware.not_managed`

- [node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203): complete subsection reference.

<a id="canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- vmware.not_managed.node_list

<a id="canonical-3111310102022223-1121023200313030-3113003033112012-3021002310213101-1321121112301203-1213322213010122-2133023122331111-2323112331313313"></a>

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

<a id="canonical-0101113233202300-2112332001320212-0211310233310231-0312312331131313-0300322003322110-2232321031132013-2133130320210010-2313222330021222"></a>

### Direct properties for `vmware.not_managed.node_list`

<a id="canonical-2132130101123001-0023330122030000-2322203231332312-2033102312011221-1232332201123013-3110223221201131-3210300302102203-3200111313003101"></a>

#### `vmware.not_managed.node_list.hostname` property

Type: `"string"`. Optional.

Hostname. Hostname for this Node.

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

- [interface_list](resources--securemesh_site_v2--reference--group-017.md#canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211): complete subsection reference.

<a id="canonical-1031301103231332-1300021332002120-2211112222112200-3211301303232010-2031111331131012-1031232111232200-1222223100130023-1020112002232133"></a>

<a id="canonical-1212132121013120-3230002312223111-3322002312100233-1332200013013332-1332212303330000-1301023300220202-2312331233211122-2302100303203002"></a>

#### `vmware.not_managed.node_list.public_ip` property

Type: `"string"`. Optional.

Public IP. Public IP for this Node.

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

<a id="canonical-3320022320303031-1303032012321333-1031010231022102-2110303222211121-2311331002022102-0102123021123223-0122103132022212-1300222021000020"></a>

<a id="canonical-0331303003200201-0203010211233301-2330113022103032-3320332222202111-0021231133201200-1202111003133000-1033031303002212-2000123021231131"></a>

#### `vmware.not_managed.node_list.type` property

Type: `"string"`. Optional.

\[Enum: Control|Worker\] Type for this Node, can be Control or Worker. Possible values are
\`Control\`, \`Worker\`.

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

<a id="canonical-1000113003012201-1113113021223320-3011313130323112-0022013111211030-3132033211031230-1320323200333011-1021331022022221-2030001113122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `vmware.not_managed.node_list.interface_list` properties

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md#canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120)
- [Property reference](resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [vmware](resources--securemesh_site_v2--reference--group-017.md#canonical-3001113303001011-1123231131321110-2233021103112320-0312233101131000-3202313121020020-3301002030023001-2131231233033311-1330001031200330)
- [vmware.not_managed](resources--securemesh_site_v2--reference--group-017.md#canonical-1030312232222330-1123312113121223-1333103322223102-0231022301200013-2012101123232213-0133223101033000-1121230211110023-2022120313101133)
- [vmware.not_managed.node_list](resources--securemesh_site_v2--reference--group-017.md#canonical-0223332200331100-3312031130222203-0202010123003323-0312312130130121-1201230100232032-2220133110000020-1110021320332123-1233001233132203)
- vmware.not_managed.node_list.interface_list

<a id="canonical-0310021211332003-0313231311233220-2031032300012221-0202032331000311-1203001010000022-0103323123002001-3123133100310310-0203301322211133"></a>

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

<a id="canonical-0303123011320312-2012102032323130-3303001110212131-3010000320303012-0333001022310222-3301110123213113-2213211010002212-2103000111233031"></a>

### Direct properties for `vmware.not_managed.node_list.interface_list`

- [bond_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-3322000201023200-0222202010223021-1320132020032301-2032311203011021-2131310122033320-0023230033012322-1221023202323132-3230313121310300): complete subsection reference.

<a id="canonical-2203300133132012-2213003213211030-3101330111222220-1102120023231130-3013202030002211-3133010202033012-0002210333210312-2112202231112000"></a>

<a id="canonical-1121033222131330-3230031233111022-1130200000321220-2103033330111321-1200133323203013-0101023203031120-1133331030312211-0132021130023231"></a>

#### `vmware.not_managed.node_list.interface_list.description_spec` property

Type: `"string"`. Optional.

Interface Description. Description for this Interface.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

- [dhcp_client](resources--securemesh_site_v2--reference--group-018.md#canonical-3003212100221203-1033201211311323-1212231223212200-3110102311220312-1022022312233212-0101001322220331-2230331201202110-1033201310120223): complete subsection reference.

- [dhcp_server](resources--securemesh_site_v2--reference--group-018.md#canonical-1330223203122133-1210110320131032-3311112303231013-3120330202102132-2201232310132231-3002101332013032-2103212220000101-0203030121021223): complete subsection reference.

- [ethernet_interface](resources--securemesh_site_v2--reference--group-018.md#canonical-1302202301013122-0221130300233000-2311320310300021-1300132002102020-0021122220311233-2013120031133021-0110023222200233-3023102000131112): complete subsection reference.

- [ipv6_auto_config](resources--securemesh_site_v2--reference--group-018.md#canonical-0130220212000003-2102133012121113-0200020303103020-2211022032102321-2011032113333301-0330222231101230-1132333030211312-2313331303030333): complete subsection reference.

<a id="canonical-2112313301122330-3232030220303032-0330123023311113-2311122302120132-1222022020320303-2010111300030111-1012020010301121-1222112330310201"></a>

<a id="canonical-1313023333101311-2203032120300322-0222313101232001-1301302310321322-3120230032031221-0332012022122102-2322013222303123-3233233123100203"></a>

#### `vmware.not_managed.node_list.interface_list.is_management` property

Type: `"bool"`. Computed.

Configuration for is\_management.

<a id="canonical-3033333200311010-1312230121231033-3132123323220120-3222133233200203-1200331001333220-1021021310321102-2120220310120213-3133122121213313"></a>

<a id="canonical-2113312122003200-1320112001213303-0023010031101310-1231010302203130-3131031321023010-0130033120121212-3213333323122131-3000111000213213"></a>

#### `vmware.not_managed.node_list.interface_list.is_primary` property

Type: `"bool"`. Computed.

Configuration for is\_primary.

<a id="canonical-3221212132330031-1020002003020123-3031002313001132-1031302311203112-2123331022001002-1020330030220120-3320031212211320-3011332022012123"></a>

<a id="canonical-2120303100223223-2201101003032130-3303021212222110-3202003210010020-2310010102331221-1121103220300302-2300103013020020-3101322330213021"></a>

#### `vmware.not_managed.node_list.interface_list.labels` property

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

- [monitor](resources--securemesh_site_v2--reference--group-018.md#canonical-2123002311101023-1232110121333102-3302333100103032-2301311000103120-3001332221330132-0132113300211131-0012320300112121-1102023200322330): complete subsection reference.

- [monitor_disabled](resources--securemesh_site_v2--reference--group-018.md#canonical-2133030222001222-3313302110203112-2323133211310201-3322121333130020-1031130322211313-1101022111003320-3333220123232320-3023020311132130): complete subsection reference.

<a id="canonical-1030202202133032-3203222321301022-0212332031103021-1212231210013302-0223202313030221-0201230123011002-1113003011101212-2121322312212123"></a>

<a id="canonical-2201122212021200-3023333230103301-0200102202030211-1023330131332020-1001303321013001-0003322310231000-3100022002111130-0330001132330202"></a>

#### `vmware.not_managed.node_list.interface_list.mtu` property

Type: `"number"`. Optional.

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

<a id="canonical-2100322201123332-3133122032321203-2030200031310002-3020213030233102-3201213113010321-2201133301122010-0220021122111101-0103330131230222"></a>

<a id="canonical-3313330320213333-2031223332231123-3113201013311212-2031013103103220-2320132030002310-2323011330233011-1302311211033302-1222200031322221"></a>

#### `vmware.not_managed.node_list.interface_list.name` property

Type: `"string"`. Optional.

Interface Name. Name of this Interface.

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

- [network_option](resources--securemesh_site_v2--reference--group-018.md#canonical-1132233323220222-1102200301311301-0231331120313213-1213000103330133-1122132100111221-0120322001130012-2021222310131321-0222100300322333): complete subsection reference.

- [no_ipv4_address](resources--securemesh_site_v2--reference--group-018.md#canonical-1000220003012122-1113011003022332-1003110313213233-1320031121121131-2112221201013001-2223103211323102-0001321032010320-0211220322133112): complete subsection reference.

- [no_ipv6_address](resources--securemesh_site_v2--reference--group-018.md#canonical-0220232122303320-2333231322122222-2011001010213200-1110323111112222-0320213321123312-3220021302012030-0212321322003301-1020211033310312): complete subsection reference.

<a id="canonical-1231300102003220-3223200311023233-1013123111100101-3313333311013023-1032232202313123-0332301030311123-3000112123133100-2130311221023300"></a>
