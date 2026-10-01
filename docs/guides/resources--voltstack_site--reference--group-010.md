---
page_title: "xcsh_voltstack_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site reference."
---

# xcsh_voltstack_site reference

<a id="canonical-1032300320201132-1023210321332232-2222310220023233-3312222310032002-1312310113211311-3320101213021131-2312023313012122-3301221110002210"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name — node_name / 333100101010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-0302110121030101-3332113333221131-1101332011331303-0023133000231111-3032313222303110-3300213220233121-0030013213102013-1133233333233201)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-2122332102113320-2100300012032123-3011321202003031-1302010311120203-1110202203230230-0122100320203221-1222121111313200-3212201112120230)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-3123221313111303-2003232311002230-2022013312100232-1012211001312310-0110201211301210-1010103223000322-2131312231330323-2232312121021202)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.node_name

<a id="canonical-3022020120132033-1021310300011002-3031311013121111-0131013302201231-1100212313010210-2100313301001312-1333031112200330-2012102130022203"></a>

Type: `"object"`. single nested block, Optional.

List of nodes on which BGP routing policy has to be applied.

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
node_name {
  # Configure direct properties listed below.
}
```

<a id="canonical-1130132202333300-2032231101111012-0012031011311332-1301113031313230-0311233100332332-1033102210210333-2132121101202103-1320130133010121"></a>

## Direct properties — node_name / 333100101010 / 3

<a id="canonical-3202023103202222-2201013112210101-2232113101302102-1123121031300021-3323103322000003-0020323131320011-0113123112031112-3333121322002001"></a>

<a id="canonical-1332221000323032-0103220233113112-3311120301012122-3332110133301221-3001322131332330-3203133112331331-1012021322332333-3232032033330210"></a>

## node property — node_name / 333100101010 / 4

Type: `["list", "string"]`. Optional.

Select BGP Session on which policy will be applied.

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

<a id="canonical-3333300032130023-2233311213120222-0311331033001223-1031133223022333-2223323012011221-1301333110013300-1133330311323300-1102003021023212"></a>

## Next pages — node_name / 333100101010 / 5

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2313211011023022-0003023122131321-0223002112133210-1321220333000332-1101301221212130-3033021002122033-2302131030323001-1123300322122000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032122130000120-3000010302232313-1123210111113102-0101033013211221-2012013202022310-2331203332111123-1320031232322310-0303233312103222"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs — object_refs / 102220101022 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-0302110121030101-3332113333221131-1101332011331303-0023133000231111-3032313222303110-3300213220233121-0030013213102013-1133233333233201)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-2122332102113320-2100300012032123-3011321202003031-1302010311120203-1110202203230230-0122100320203221-1222121111313200-3212201112120230)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-3123221313111303-2003232311002230-2022013312100232-1012211001312310-0110201211301210-1010103223000322-2131312231330323-2232312121021202)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.object_refs

<a id="canonical-2222112122332030-3222121012113022-1212022222323120-1001302003223301-0323122011020110-2013323133203002-1323112103133232-3210223033111010"></a>

Type: `"object"`. list nested block, Optional.

BGP routing policy. Select route policy to apply.

Upstream description:

Select route policy to apply.

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
    "minItems": 1
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
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
object_refs {
  # Configure direct properties listed below.
}
```

<a id="canonical-2203003100021033-1131121023020000-2233020030001302-2001232320120030-3110013232303130-3202133220111232-2232331333311212-2300121101222112"></a>

## Direct properties — object_refs / 102220101022 / 3

<a id="canonical-1101130132023203-0201002012233220-1200123220020013-0200022030331110-0133033000033020-3330233022230133-0221120233330333-3110303130001201"></a>

<a id="canonical-3313211130133130-2011022010022200-1033321311320202-2011331321231133-2012121230000121-2010021210313010-0302301302202331-2121331103000320"></a>

## kind property — object_refs / 102220101022 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3101031002232012-3110212120312221-3313222231222322-1310203113320133-0321312110003110-2322002013222331-3323101000321303-1222320010121112"></a>

<a id="canonical-0223320133303133-3131233320033101-1202103010130023-1011200030100123-0322211201131103-2012323232230222-2121330133302213-1332311213221030"></a>

## name property — object_refs / 102220101022 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3303000101002132-0123010312022231-0031102013120022-3032011122122002-2031311233000332-1130330323023120-0311202123003133-2222300212211033"></a>

<a id="canonical-2202000120312021-0233233012101110-1323223021020301-2101100010221333-1032232003200333-2111303213321221-3212203020230032-0122120232310001"></a>

## namespace property — object_refs / 102220101022 / 6

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0221003111000211-1322102112332312-2111312311233320-0220111323311101-3111223111011103-1132302230000310-0233310222331303-0220311310000030"></a>

<a id="canonical-2023203330020230-0202023231033223-3222132223110312-3033130001000030-0310311123323013-3033211120123121-2203103030003320-2120021130012132"></a>

## tenant property — object_refs / 102220101022 / 7

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1102020303101103-2222320011101300-1133031033003212-1101031123223113-2211110030100212-1220120021023333-1333100201220121-1100112000023303"></a>

<a id="canonical-3112030123001330-0033232123210001-1312123303032321-0200333302122222-1112221123213121-0003320022220332-1130103232003220-3203030203321122"></a>

## uid property — object_refs / 102220101022 / 8

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0212020013010011-2100133332221232-2120300222202302-1303311223201123-1301100300323301-0303022212033211-2223122023131332-3202300321332111"></a>

## Next pages — object_refs / 102220101022 / 9

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1212320313112202-2100113103123131-0030332132310311-0302333122010233-1131311033110112-3012203033000120-3111031100002033-2333320001013313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203113300000233-1311331102030222-2130131122320132-0233121302330122-3013133000202311-1032321331111103-3331031322120220-3001131022202100"></a>

## local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound — outbound / 023232222012 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- [local_control_plane.bgp_config](resources--voltstack_site--reference--group-009.md#canonical-0302110121030101-3332113333221131-1101332011331303-0023133000231111-3032313222303110-3300213220233121-0030013213102013-1133233333233201)
- [local_control_plane.bgp_config.peers](resources--voltstack_site--reference--group-009.md#canonical-2122332102113320-2100300012032123-3011321202003031-1302010311120203-1110202203230230-0122100320203221-1222121111313200-3212201112120230)
- [local_control_plane.bgp_config.peers.routing_policies](resources--voltstack_site--reference--group-009.md#canonical-3123221313111303-2003232311002230-2022013312100232-1012211001312310-0110201211301210-1010103223000322-2131312231330323-2232312121021202)
- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202)
- local_control_plane.bgp_config.peers.routing_policies.route_policy.outbound

<a id="canonical-3321301223033003-0302203000211211-1232323020221211-0002023322230321-3223112323210103-2202212013231223-1133330020003233-2230323233200321"></a>

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
outbound = {}
```

<a id="canonical-3101013022202110-1033333012113230-0110123201222121-3303131332213130-3321103020130021-2120013112132330-1332322003002301-1223303312013231"></a>

## Direct properties — outbound / 023232222012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0123030320331000-1221132331011332-1100023132020011-3023111130202311-0012331112300231-0133113321121230-3213212311302233-0301022130110301"></a>

## Next pages — outbound / 023232222012 / 4

- [local_control_plane.bgp_config.peers.routing_policies.route_policy](resources--voltstack_site--reference--group-009.md#canonical-0111023012031033-1110322032011310-1223223311112332-0323231100303332-1330310002132020-1311303332130313-0100010310011110-1330102031211202)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1100312102203023-2000332022310300-2230210032022231-3200310130221221-0302232113000310-1332212020331200-1232220333020022-0331131203233312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232312001021302-1023212121221112-3121120102202320-1130312320202202-0200023011211310-1211120201311023-0322202102130313-2122332010001203"></a>

## local_control_plane.inside_vn — inside_vn / 133002101003 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- local_control_plane.inside_vn

<a id="canonical-3020032100110232-1310213310031120-0220312113333102-3031213222223110-3330133133132102-0033021303110133-3121012030131001-3232322222322030"></a>

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
inside_vn = {}
```

<a id="canonical-0210233132232310-2102131102131333-0122110030310123-3322103323101332-0121313301010311-0003123120223122-2220003123230010-3302322323300301"></a>

## Direct properties — inside_vn / 133002101003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221332321202213-1232012130213313-2111031010030023-0022102100310200-3020033202223311-3220100113333110-2023030330223333-1332312122133120"></a>

## Next pages — inside_vn / 133002101003 / 4

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0231023010332031-0201100220010121-0313222311122110-3220200233120031-0120013222123210-2132212031122002-2000122013100023-3322010202200321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013211012011010-3110023123310223-3321103321221002-1320312030210031-3203001013021230-2030232201311023-3312232212310112-2211003101323310"></a>

## local_control_plane.outside_vn — outside_vn / 022333302010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- local_control_plane.outside_vn

<a id="canonical-1303123311330330-0020230232313102-3210231123232313-2122310302013301-3312023132123021-2331130013101303-3212131232221003-0310212313211113"></a>

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
outside_vn = {}
```

<a id="canonical-1001311320123122-1320103133022000-2310312032303220-1103003012313223-1210210003331102-0333322203110301-2330003100303232-3002322033133131"></a>

## Direct properties — outside_vn / 022333302010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230222010111022-1011100332132103-0023012001131110-3011312131222302-1232011330320233-1211313311203210-2231231133212231-3223002122001110"></a>

## Next pages — outside_vn / 022333302010 / 4

- [local_control_plane](resources--voltstack_site--reference--group-009.md#canonical-1132320332333230-2221113332303320-2330321010200300-2023122310102023-1022200001333011-3303122211221301-1130130210003310-1133211030011101)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3310021323320200-2131332321230123-0021023332302120-2312333000133000-0003300021330330-3031120120202103-0332000001202021-2012232322101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230133120100012-3330023002111332-1110320123223202-3320113011300201-2030332102100100-1301222031332302-0212201221211332-2330213322003111"></a>

## log_receiver — log_receiver / 302020100331 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- log_receiver

<a id="canonical-1000201122033323-0200003023021302-1003230322132133-1020210312023323-2010300111230030-1122121310212133-1312021102010221-2101102222321112"></a>

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

- [log_receiver](resources--voltstack_site--reference--group-010.md#canonical-1000201122033323-0200003023021302-1003230322132133-1020210312023323-2010300111230030-1122121310212133-1312021102010221-2101102222321112)
- [logs_streaming_disabled](resources--voltstack_site--reference--group-010.md#canonical-1111102121231212-1111332311122011-3222302221321103-3021011200103120-0330202120112230-1321332233222023-0111033201031112-3022323100023002)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230113313201302-2103000232321212-3322000100323023-3011210303031001-1021220021110030-3212100211123202-1111030001113330-0303022012310212"></a>

## Direct properties — log_receiver / 302020100331 / 3

<a id="canonical-0101203032311030-1021221131132201-2012310200122310-2333102103101230-1120230231323110-3133312210132100-1230221013102000-3113332021300311"></a>

<a id="canonical-0122302021122213-0032100213023032-0211020231200201-3132210003032323-1100201123233201-1002321231100210-1101030003011302-1233331013223222"></a>

## name property — log_receiver / 302020100331 / 4

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

<a id="canonical-1232101001111332-2200303031030232-2132233121200103-1131122003313101-0323223120231020-1111311013232003-3311221232022131-0311220001222320"></a>

<a id="canonical-3121130102001301-2322011222133331-0112312011033101-2210320203230102-3032331001001130-1212210023000322-0201100013203011-2302011030103313"></a>

## namespace property — log_receiver / 302020100331 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0302032220311013-2220322011230333-0311003022120320-3300102012201103-3221121321302312-2011121232013301-0011202013212333-2111323102201111"></a>

<a id="canonical-1122212333311013-2231221011131012-1200100313022223-2210031121331020-0000202213332312-0213103000132302-3101223303103033-1220133203321202"></a>

## tenant property — log_receiver / 302020100331 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3120300012123330-3032033022032232-0103011232001010-1223210333312123-0231331120111020-0010330111030131-3132323003321212-1132003001223302"></a>

## Next pages — log_receiver / 302020100331 / 7

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2013003210030303-0313301312303101-1102333300333132-2302212310311100-1312130101123323-1332121302130312-0122131212310021-2013303121303232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0012112210202332-3212132313112221-1310331332111320-2131210003111002-2112133210330011-0101310113011220-2132201233033323-3113020012330132"></a>

## logs_streaming_disabled — logs_streaming_disabled / 232223033100 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- logs_streaming_disabled

<a id="canonical-1111102121231212-1111332311122011-3222302221321103-3021011200103120-0330202120112230-1321332233222023-0111033201031112-3022323100023002"></a>

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

<a id="canonical-2112212003032033-1213110002022321-1313200331321221-3211010232120022-3230332002232312-3011110230220330-3031133320102003-2023311013202121"></a>

## Direct properties — logs_streaming_disabled / 232223033100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212302302033130-3112203010201321-2223201313310033-0321323010300303-3121112113111313-3330022223013010-0132202200102233-2231131023033110"></a>

## Next pages — logs_streaming_disabled / 232223033100 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2231013002202321-1310031001031301-2022032301131330-1102123002203131-0101322113031031-0102323000220120-2121200313301302-1032303230301033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001000320000322-0330023012022310-1320000113201110-2332122201013202-1302213002331311-1111100130202021-1201030311013330-0221022111032002"></a>

## master_node_configuration — master_node_configuration / 101132211032 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- master_node_configuration

<a id="canonical-3023031113112123-0122203021122131-1132120100301212-3112123320101321-0210320032213023-2111200302003102-0032230330213201-3103230030212000"></a>

Type: `"object"`. list nested block, Optional.

Master Nodes. Configuration of master nodes.

Upstream description:

Configuration of master nodes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "3"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3"
  }
}
```

Terraform syntax:

```terraform
master_node_configuration {
  # Configure direct properties listed below.
}
```

<a id="canonical-1033121333000102-3312011020210212-1133023011213323-0110021013230120-1113133202030112-1122023320211222-1322303113032201-2320013120221201"></a>

## Direct properties — master_node_configuration / 101132211032 / 3

<a id="canonical-1131320300330110-0310302020100322-1323120333231112-3200301031011100-0211033103221331-0322132011201123-2331022313130031-1313012211033031"></a>

<a id="canonical-0022231202000332-0123020330123031-2312321120301213-0313333020323113-1122033113210302-1113333331120300-1332133312202210-3300012211112211"></a>

## name property — master_node_configuration / 101132211032 / 4

Type: `"string"`. Optional.

Name. Names of master node.

Upstream description:

Names of master node.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-0232233102131131-3010330103230200-1030300031210212-0100231133103003-2133113131131103-2131313200031313-2132221130102102-1212113213020121"></a>

<a id="canonical-2111021231212202-2030033231001031-1100000230022102-1132013020121212-3013013021303212-1312221332300323-1012001122013323-0121321122323322"></a>

## public_ip property — master_node_configuration / 101132211032 / 5

Type: `"string"`. Optional.

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

Upstream description:

IP Address of the master node. This IP will be used when other sites connect via Site Mesh Group.

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

<a id="canonical-0103300031131123-2220112023110223-2213200132133010-3010123201233330-0221221001312103-0321212121310023-0121302112013021-0322130320000101"></a>

## Next pages — master_node_configuration / 101132211032 / 6

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0111122002330203-3213023011101331-1202010121032023-2310021300012322-3231300020223000-0312320110132201-3233220021110031-1221330021232331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002123011311003-3123012300222132-0000032021132131-1210201031232123-3010122000110220-0022212210230123-1200202012213330-2300320330112232"></a>

## no_bond_devices — no_bond_devices / 301200302122 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- no_bond_devices

<a id="canonical-2010003223301021-1221033133220013-0223013010132321-1320001001210102-1213030210031020-3200013132100232-0213232102032221-2331022330323010"></a>

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

<a id="canonical-3122302102201100-1112130123230222-3103203020131332-1113200021121210-1123020211132333-3113110202120303-0310332312303130-2012302013312122"></a>

## Direct properties — no_bond_devices / 301200302122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000331012220302-0221120221310210-0323101223022310-1222020001030032-1302212002323232-3332231222320222-2201122132030131-3033122102023102"></a>

## Next pages — no_bond_devices / 301200302122 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1110311231110111-0010333301003210-2303111023211302-3123121321000123-1332030003230203-2130132012333122-1121202021001232-1223130230213223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021012320310010-2000231021313100-3121120332030223-3232201213321211-0232113110012222-2313223121321331-1222201002032011-3003120110233001"></a>

## no_k8s_cluster — no_k8s_cluster / 133233323000 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- no_k8s_cluster

<a id="canonical-1232133321122033-0301112232311301-1100310013013301-2113110101031030-3331032322330120-3201010302332322-3220221132323123-0200231003302011"></a>

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
no_k8s_cluster = {}
```

<a id="canonical-3322130011200100-3330332013220002-3103112333331231-3030021031213221-1202303133031021-3231103212230101-1123131123120310-3013213031232002"></a>

## Direct properties — no_k8s_cluster / 133233323000 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0013131103110203-1023311020230012-3221223302120110-0001210213123111-2113132020333320-0130301131220313-1221212130210302-1322000213213233"></a>

## Next pages — no_k8s_cluster / 133233323000 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1212130122323321-1021203223212003-2213000310220323-3003323121130210-3312333213031112-1210302211021100-1311032332331321-1322130033211213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3203103331320300-0201222230220023-1200033032031103-2301022033113310-3131223302210113-1211123133232133-3313130033132031-0332120020222130"></a>

## no_local_control_plane — no_local_control_plane / 031021233213 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- no_local_control_plane

<a id="canonical-0231010211100300-0332023000301113-2312310330020311-2303000211221033-1232322322011121-3301013133223020-3333131011202120-1110303101033300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no local control plane.

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
no_local_control_plane = {}
```

<a id="canonical-2321001100020133-1232200232310112-3320020032332330-3033032131032300-1200113311013103-0202210003210021-2102301020321222-0122332200211303"></a>

## Direct properties — no_local_control_plane / 031021233213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033110330110032-2320010320211332-1211220031121302-0231232102101113-1220103301321003-3002000110332222-0312101113203331-2101230221030111"></a>

## Next pages — no_local_control_plane / 031021233213 / 4

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0120123201103100-3121123323333132-2323310020001010-3232000323111313-0230313201222013-3202130322013321-2133303000333212-0112033111011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210122211233301-1013202001322333-2320121013131231-0113033130012321-3310111333302230-0100233320120002-2023200012120131-2033312030332202"></a>

## offline_survivability_mode — offline_survivability_mode / 101220302233 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- offline_survivability_mode

<a id="canonical-1303312203020103-0322202122332132-1033222200032003-2001001001310111-2130232313012210-3020132113131332-1033112030232023-0201302133300232"></a>

Type: `"object"`. single nested block, Optional.

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7..

Upstream description:

Offline Survivability allows the Site to continue functioning normally without traffic loss during
periods of connectivity loss to the Regional Edge (RE) or the Global Controller (GC). When this
feature is enabled, a site can continue to function as is with existing configuration for upto 7
days, even when the site is offline. The certificates needed to keep the services running on this
site are signed using a local CA. Secrets would also be cached locally to handle the connectivity
loss. When the mode is toggled, services will restart and traffic disruption will be seen.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3213003223200312-1310332300021122-3033001031233320-1303110113021210-2222030022122223-3320031120312203-1023102110021022-2202312010221112"></a>

## Direct properties — offline_survivability_mode / 101220302233 / 3

- [enable_offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-2330122013022010-2020332003233100-1023313102321021-0232003310230200-3330130111310021-0331210300303133-0032021133122121-1013212022231002): complete subsection reference.

- [no_offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-2210201011002230-1213023312211222-3330132321221313-2223120001033232-1323123200331301-0030331212022123-3322123023301202-0012000211001233): complete subsection reference.

<a id="canonical-0223220110012320-1212320113221101-3123013331010122-0103231203001000-2010033313300230-3011000022003300-2012003203101211-1332033310132030"></a>

## Next pages — offline_survivability_mode / 101220302233 / 4

- [offline_survivability_mode.enable_offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-2330122013022010-2020332003233100-1023313102321021-0232003310230200-3330130111310021-0331210300303133-0032021133122121-1013212022231002)
- [offline_survivability_mode.no_offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-2210201011002230-1213023312211222-3330132321221313-2223120001033232-1323123200331301-0030331212022123-3322123023301202-0012000211001233)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2330122013022010-2020332003233100-1023313102321021-0232003310230200-3330130111310021-0331210300303133-0032021133122121-1013212022231002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123320012333202-3313222000120132-3302031112002130-2231303302010231-2220101101132033-0300002230100032-2212012301003232-3103330103011111"></a>

## offline_survivability_mode.enable_offline_survivability_mode — enable_offline_survivability_mode / 110111101202 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-0120123201103100-3121123323333132-2323310020001010-3232000323111313-0230313201222013-3202130322013321-2133303000333212-0112033111011301)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-2131020303113111-2231020120122200-0322310322120131-1033113012210020-0201200121031101-2033200330302130-3011213120100121-1303201003032211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

<a id="canonical-2112233213320223-3111210203233110-2232313333320130-3003212020303023-1323123310320303-3103311301313102-3301030321003203-3121030223233212"></a>

## Direct properties — enable_offline_survivability_mode / 110111101202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322023110212212-2201122120131110-2332120313123002-3210332223010323-0113002111112020-0023002320113022-3023332123321303-1230230232113233"></a>

## Next pages — enable_offline_survivability_mode / 110111101202 / 4

- [offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-0120123201103100-3121123323333132-2323310020001010-3232000323111313-0230313201222013-3202130322013321-2133303000333212-0112033111011301)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2210201011002230-1213023312211222-3330132321221313-2223120001033232-1323123200331301-0030331212022123-3322123023301202-0012000211001233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103322322020120-0223010112121221-0213233021033320-0230313201213023-0022131100221210-0233320013233301-1102222301332303-3001131133202110"></a>

## offline_survivability_mode.no_offline_survivability_mode — no_offline_survivability_mode / 101321133132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-0120123201103100-3121123323333132-2323310020001010-3232000323111313-0230313201222013-3202130322013321-2133303000333212-0112033111011301)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-3201012132030223-1011312120022000-0232331211021021-3202210333313233-1200231332302210-3030212210132212-2020331121010222-3222301313023221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

<a id="canonical-0030231020113113-0210031213110332-0021113030023312-3100013210022201-0210200312103202-1103120302230321-0322311232301031-3101011203130212"></a>

## Direct properties — no_offline_survivability_mode / 101321133132 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000101111121230-1100023010002230-2231223102322003-0031201330200031-3223302112130100-1333323020101303-2112200112103302-2221011331123122"></a>

## Next pages — no_offline_survivability_mode / 101321133132 / 4

- [offline_survivability_mode](resources--voltstack_site--reference--group-010.md#canonical-0120123201103100-3121123323333132-2323310020001010-3232000323111313-0230313201222013-3202130322013321-2133303000333212-0112033111011301)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2230223331213210-2003111102030202-0101123130331303-1321202222013223-3333311120130122-1231202131312311-1202210112301113-3211213301122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220300013202312-0210121310301123-3300033102201312-0030212032111010-3130020000113133-2201012031030300-0303023311032012-1232020311122122"></a>

## os — os / 110010322010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- os

<a id="canonical-1330030202331110-0322331232001131-3123022203032331-2320201003122321-2300313020311122-0032333133221313-1233030010200312-0132032101212001"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Upstream description:

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

<a id="canonical-1312303113330123-2203013111100232-1310101211310333-1101030103312332-2320020302113121-0120133001212002-3332300020000202-0112212012302200"></a>

## Direct properties — os / 110010322010 / 3

- [default_os_version](resources--voltstack_site--reference--group-010.md#canonical-0100032303230030-1133300001120113-0001200101220001-3111133302021130-0012321132033321-1332220212200331-1033233031111111-1123111111000120): complete subsection reference.

<a id="canonical-1302112030012102-1220311002202303-1121003330332100-1123203110003233-3203102210020303-1021001201111122-3333102112000131-1211300033313221"></a>

<a id="canonical-0033312300330011-1100221022233233-0111222012122230-0321031202130113-2033123102323212-2112120123132112-3002222031230311-1032002002001101"></a>

## operating_system_version property — os / 110010322010 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-3013213022113210-0102322212201102-3021311003033322-1323333010103000-0032210220110133-3011213231330233-2030002100110002-0201023223310101"></a>

## Next pages — os / 110010322010 / 5

- [os.default_os_version](resources--voltstack_site--reference--group-010.md#canonical-0100032303230030-1133300001120113-0001200101220001-3111133302021130-0012321132033321-1332220212200331-1033233031111111-1123111111000120)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-0100032303230030-1133300001120113-0001200101220001-3111133302021130-0012321132033321-1332220212200331-1033233031111111-1123111111000120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3030002223230111-2210012231313303-0320003033113120-3200100121331103-0231000103122312-3210301032003311-0003020200321311-0313203303000103"></a>

## os.default_os_version — default_os_version / 113223022010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [os](resources--voltstack_site--reference--group-010.md#canonical-2230223331213210-2003111102030202-0101123130331303-1321202222013223-3333311120130122-1231202131312311-1202210112301113-3211213301122003)
- os.default_os_version

<a id="canonical-0110322203201203-2210320313313322-3123112221022133-3230312331223212-3323003330001033-2121031003102121-0102332212012122-1131013332332230"></a>

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
default_os_version = {}
```

<a id="canonical-2001113311013231-1333130013002032-3112223321011321-0100321231331203-3231203302331322-0122333303231301-0120233030010312-2213310120111230"></a>

## Direct properties — default_os_version / 113223022010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1120103102130101-0233313231102333-1211132311021200-0233323333201033-2233032200110333-1331213321233301-0021321020300220-2232311031123011"></a>

## Next pages — default_os_version / 113223022010 / 4

- [os](resources--voltstack_site--reference--group-010.md#canonical-2230223331213210-2003111102030202-0101123130331303-1321202222013223-3333311120130122-1231202131312311-1202210112301113-3211213301122003)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3202202012201132-3231030210000322-0123032022001103-3031211120312222-3332121322033010-1220331122103013-1123232303130120-1200100021331233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020013032023332-0211012323330301-2331311313320201-3302331321122302-3002003121130022-0100113121321320-1310232032011223-3000121213031213"></a>

## sriov_interfaces — sriov_interfaces / 201311321331 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- sriov_interfaces

<a id="canonical-1003120011111311-2102230121010132-1231113301123200-2231002033323212-2322023220320232-2122201010303332-3130213010200123-3032322003010200"></a>

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

<a id="canonical-1001233013231023-0132330030220302-2003201000333231-2312111022122121-3330030313303022-1231331100031032-3231022301300302-3311111033201213"></a>

## Direct properties — sriov_interfaces / 201311321331 / 3

- [sriov_interface](resources--voltstack_site--reference--group-010.md#canonical-1211212101221221-1021122103330121-2310112103213310-1123213210212020-3203103101013212-0200213311212023-0223331201133333-0301333023112233): complete subsection reference.

<a id="canonical-2010210013213203-2100212213300230-1030232301210013-2311022332201331-2033002131110230-3102310033300113-2112220231233030-1212000220030330"></a>

## Next pages — sriov_interfaces / 201311321331 / 4

- [sriov_interfaces.sriov_interface](resources--voltstack_site--reference--group-010.md#canonical-1211212101221221-1021122103330121-2310112103213310-1123213210212020-3203103101013212-0200213311212023-0223331201133333-0301333023112233)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1211212101221221-1021122103330121-2310112103213310-1123213210212020-3203103101013212-0200213311212023-0223331201133333-0301333023112233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3032310100310213-3100330010001110-0022302031020230-2122232131021322-0222203110302300-3321130010232132-0213320203003113-0030300132311110"></a>

## sriov_interfaces.sriov_interface — sriov_interface / 210103100113 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [sriov_interfaces](resources--voltstack_site--reference--group-010.md#canonical-3202202012201132-3231030210000322-0123032022001103-3031211120312222-3332121322033010-1220331122103013-1123232303130120-1200100021331233)
- sriov_interfaces.sriov_interface

<a id="canonical-3022032011120231-1101020332122101-3211221013012230-0110323023122030-3100200332220023-0322211333222000-2210032212030030-0220301033102131"></a>

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3130303302030002-3310101110311211-0011201100010323-0333131231201122-2030103302101013-2201023021212112-3020112310021302-1231230311012200"></a>

## Direct properties — sriov_interface / 210103100113 / 3

<a id="canonical-3231011101331130-1211303132210300-1331132013322303-0112022120012013-2021000333131110-3122303210233013-1220032301203222-0101321003330100"></a>

<a id="canonical-1300233233322212-3133301233011131-1102030212320200-1002013201330220-3023021202303120-1120012001203032-1132320201211213-3202010122020030"></a>

## interface_name property — sriov_interface / 210103100113 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2022121032113221-3231321212132011-2021111230320123-1311210001232201-0132031312110112-0011122020130020-2313301010032001-0211223112301002"></a>

<a id="canonical-2012213010100010-0230021312132013-1103203122313223-2200220321132213-2102303330212003-3300000332311122-3000302232033313-3022100113323001"></a>

## number_of_vfio_vfs property — sriov_interface / 210103100113 / 5

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

<a id="canonical-1312002212032311-2023030320103101-0231022033220030-3211221223302211-1310210013203231-1230313212023110-1210320030332023-1133110002203103"></a>

<a id="canonical-1102313100223011-1331302331032120-3023220222123221-0230321112001300-3211231112120331-3123021100203101-1312023221003112-1203210133203132"></a>

## number_of_vfs property — sriov_interface / 210103100113 / 6

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

<a id="canonical-0130301100100103-3030202320112221-3030003123003222-1131023302122213-3202323212221113-3001112113202303-3103023223311321-0222100323120100"></a>

## Next pages — sriov_interface / 210103100113 / 7

- [sriov_interfaces](resources--voltstack_site--reference--group-010.md#canonical-3202202012201132-3231030210000322-0123032022001103-3031211120312222-3332121322033010-1220331122103013-1123232303130120-1200100021331233)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1310110011022322-0333211000331321-3000220132023131-1232213111323122-3332210311120230-2120231033120210-3213230302031001-1112101021020320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332113332133120-0132220320101302-1032031120110101-3033312322002333-3032302233002203-1102132002110102-0000111001010233-3323010210233230"></a>

## sw — sw / 311301333132 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- sw

<a id="canonical-0321312033220222-3133231030203002-2012312012102321-1120230322213320-1202022113100322-2321212131200111-2233200033100223-1100030120112303"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Upstream description:

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

<a id="canonical-2311110212222001-1113013130033002-0103120003133332-1333323322320332-2301330100203313-1212210022020002-1111130230020020-1220222222031022"></a>

## Direct properties — sw / 311301333132 / 3

- [default_sw_version](resources--voltstack_site--reference--group-010.md#canonical-3032312030201331-1021100031031203-0311213032322130-3003112203320001-1113033300300122-0121103110123033-1133302120120311-2011101201213001): complete subsection reference.

<a id="canonical-3331310222113213-1023113312011002-0311212323001233-0102233130310203-0130221232112000-1203100120003201-3133100310223110-3323323212231302"></a>

<a id="canonical-2211002313213011-1121102111113002-0011012222233301-0211111011220233-2201310330101033-2310232102210302-3021212303221122-3113333300133312"></a>

## volterra_software_version property — sw / 311301333132 / 4

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Upstream description:

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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-1322223020323200-3312320123310221-1322232121101221-0103311010011110-1022201120332013-0033122321311330-1100220312102232-3212020230131203"></a>

## Next pages — sw / 311301333132 / 5

- [sw.default_sw_version](resources--voltstack_site--reference--group-010.md#canonical-3032312030201331-1021100031031203-0311213032322130-3003112203320001-1113033300300122-0121103110123033-1133302120120311-2011101201213001)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3032312030201331-1021100031031203-0311213032322130-3003112203320001-1113033300300122-0121103110123033-1133302120120311-2011101201213001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010031331013132-2003113100300123-2322133301110301-1333223220122213-2113103013302121-1002203013213100-0032030010313130-1232031230000000"></a>

## sw.default_sw_version — default_sw_version / 012211133210 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [sw](resources--voltstack_site--reference--group-010.md#canonical-1310110011022322-0333211000331321-3000220132023131-1232213111323122-3332210311120230-2120231033120210-3213230302031001-1112101021020320)
- sw.default_sw_version

<a id="canonical-1032033221320322-1303121120030021-1230222000310232-0111330023102331-2121113011113313-1111101012200130-0112120003032221-2230121300321112"></a>

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
default_sw_version = {}
```

<a id="canonical-0223112233312010-0011323222110213-1133222130313210-2003103011222112-2202020113000322-0132210321122211-3303201022201201-1210133130231200"></a>

## Direct properties — default_sw_version / 012211133210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3311123223203133-3111322031233131-2333302001221021-0113320003012010-0313013221203031-1321232211230313-0313223023311132-3133023132020220"></a>

## Next pages — default_sw_version / 012211133210 / 4

- [sw](resources--voltstack_site--reference--group-010.md#canonical-1310110011022322-0333211000331321-3000220132023131-1232213111323122-3332210311120230-2120231033120210-3213230302031001-1112101021020320)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1130122203230021-0301012333023233-1330223323223133-3120132222110331-2110010321001222-2011322100322230-0323000302330311-0022221031222233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301230003022211-1220320320130311-0230021032032211-2102111132333230-1131203013203112-0232101122300230-3133011030301032-3331231032011333"></a>

## timeouts — timeouts / 332232023010 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- timeouts

<a id="canonical-2101121100300003-3303323320011031-3011202220032131-3312020000301220-0132232122101220-2111022020230112-0132230312311020-3203211321222113"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113313223310312-0112103221303112-0311132133302210-0312101220333012-0022333122302300-1303220322131023-0333303032131130-3031213331132033"></a>

## Direct properties — timeouts / 332232023010 / 3

<a id="canonical-2022323202321212-3121220013133022-2300230312200321-3103102113203220-2103123033333100-3230112032012012-1012100213330221-1322333023012202"></a>

<a id="canonical-2033320123031233-0110033220002312-1103302233012311-0321322002010020-0121233211103101-2321000231121233-1212200203112223-2100032031102111"></a>

## create property — timeouts / 332232023010 / 4

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1301122223002131-0212221020113303-2131231321020032-0211320323011100-0203020221322001-2133302203130231-0331122211311030-3201131101231131"></a>

<a id="canonical-1122323333322030-0102302202232230-2033013200023113-2322332311223022-2220332032113331-1300131101201013-3110023011023231-1301212100211333"></a>

## delete property — timeouts / 332232023010 / 5

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-1011313233211313-1021013323022122-2312023301230220-3332103123200223-3121332123133033-2121031020321131-1223133021121103-2101223032211033"></a>

<a id="canonical-0113222120201011-1020211001211112-0312232121011011-3033322010312002-1322230110002220-3002321200223102-0212232112023300-0203220202100133"></a>

## read property — timeouts / 332232023010 / 6

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-2010103233330111-1122220231223323-0132130203000002-1002132223130002-2202332221233320-0123010033023100-1330002002301131-3302022012032223"></a>

<a id="canonical-1212120212101321-1101232003311212-2032221310230002-2013110021021121-3330113010021103-0233131120111113-2331030132333023-1101121133202321"></a>

## update property — timeouts / 332232023010 / 7

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-1001232220301030-1220103101132320-2203113321010001-1012013301210202-1221203020313310-0100222011002131-1213203122120102-1000210320202220"></a>

## Next pages — timeouts / 332232023010 / 8

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3332312032102223-1310111330103012-1312102310030023-3213120011021310-2232310220220331-3331131211023113-3023103301303211-1013210310330303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2122230123202321-3302120221310320-0013100022333133-0311220123120300-3202032000213332-2301100033133123-3002031302102132-0023032201232213"></a>

## usb_policy — usb_policy / 233300313320 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- usb_policy

<a id="canonical-3311022202102320-3313322200122230-2333102332000211-0221203101131011-1122020311121032-2233311323303311-0313020330320222-0103122213222312"></a>

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
usb_policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-3013103210203202-0123020001220100-3201322313311120-2030203111302233-3100311223001132-1211003022332033-0021132321111333-1112230311013000"></a>

## Direct properties — usb_policy / 233300313320 / 3

<a id="canonical-1212002133301232-1032112300203130-2231223323130030-2212123303200101-3000311302221312-0302000223010122-1311112100023213-2133001002023200"></a>

<a id="canonical-1310130002300202-0302022203301113-0121322022232300-1213122013010322-3112312301023100-3031031310221001-0103332030103100-2301330210221323"></a>

## name property — usb_policy / 233300313320 / 4

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

<a id="canonical-3331033113033123-3103201313001233-2330230212112102-2302023331010021-0011211303313020-3122210201333002-0030302311001210-1013212010111310"></a>

<a id="canonical-3132113102310302-2330030302123203-3011030021120011-3003112321333003-0013220331310303-1013221330311330-3203003300330320-1300332232133012"></a>

## namespace property — usb_policy / 233300313320 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1310302133133302-1211210221333222-2013233031133102-3030201113131121-0230011302201213-1110301031320232-3130200320010212-0121310323321220"></a>

<a id="canonical-0332310310000032-0330310231132010-2311303203031323-3221113231213031-2011101200102130-3311223021321001-2220233301033213-0322021103213212"></a>

## tenant property — usb_policy / 233300313320 / 6

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3210120113010203-3300001011033303-3132023103002303-3013232110000212-2131112311302202-0101332332321322-2203200030012001-1223221001000100"></a>

## Next pages — usb_policy / 233300313320 / 7

- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3111030230112323-2102333310010310-3013310312010302-0120130130203331-3010103133133201-0301102111310021-3100333123132021-2101021003011122"></a>

## waf_signatures — waf_signatures / 013122311020 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- waf_signatures

<a id="canonical-2112323322222322-2312100120333220-1302330320203303-3023002302122002-3202221232003313-3033212100001023-0221113130323200-3233320330303321"></a>

Type: `"object"`. single nested block, Optional.

Select F5XC WAF Signatures update mode for the site. By default, new signatures will be applied
manually. Refer to release notes for details about available Signatures update modes.

Upstream description:

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

<a id="canonical-3223323031030231-0131030133103300-1020033232210113-0010013011232230-0301111233003010-1232332121022312-2321121013130301-0032112123102230"></a>

## Direct properties — waf_signatures / 013122311020 / 3

- [automatic](resources--voltstack_site--reference--group-010.md#canonical-2132112212021130-3210203033100002-2211201021311003-0120122031113011-3223100001032123-1300023232202013-2320003100132030-2200021003132333): complete subsection reference.

- [manual](resources--voltstack_site--reference--group-010.md#canonical-1010332120003233-0100110202322233-0111220220011233-1203321020031332-2310320123311232-0302110103113201-3331222212202121-0111131100301312): complete subsection reference.

<a id="canonical-2211033212031311-3233132101021312-2221230231332330-3220322333113003-2300233310200130-3232332221230302-1002110023323312-1203220322323202"></a>

## Next pages — waf_signatures / 013122311020 / 4

- [waf_signatures.automatic](resources--voltstack_site--reference--group-010.md#canonical-2132112212021130-3210203033100002-2211201021311003-0120122031113011-3223100001032123-1300023232202013-2320003100132030-2200021003132333)
- [waf_signatures.manual](resources--voltstack_site--reference--group-010.md#canonical-1010332120003233-0100110202322233-0111220220011233-1203321020031332-2310320123311232-0302110103113201-3331222212202121-0111131100301312)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-2132112212021130-3210203033100002-2211201021311003-0120122031113011-3223100001032123-1300023232202013-2320003100132030-2200021003132333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030233030320103-1213001233031121-2132211123213300-1300103131233230-0313330003012321-1001112221201002-3322321012332201-1301231233302201"></a>

## waf_signatures.automatic — automatic / 222132002220 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132)
- waf_signatures.automatic

<a id="canonical-2202112002303311-1312201332111103-2111223022320110-3323123313333220-1110023100331023-0100022132030202-0212110022213021-3113022023030122"></a>

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
automatic = {}
```

<a id="canonical-2232120312032223-2011210201320012-2103312133032010-3220020232010231-1312122030311333-1211332133230230-3013123232001210-2332212033320223"></a>

## Direct properties — automatic / 222132002220 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111131312033321-3200030110130111-3002330203031103-3032212311322220-2303303323210323-0322023331102221-1030331111020202-2302101121030333"></a>

## Next pages — automatic / 222132002220 / 4

- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)

<a id="canonical-1010332120003233-0100110202322233-0111220220011233-1203321020031332-2310320123311232-0302110103113201-3331222212202121-0111131100301312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210211321012302-1011310322301231-2210003123301130-3133321021120031-3003021230320232-2010331133213011-0113133033003112-0330110011021003"></a>

## waf_signatures.manual — manual / 112231322120 / 2

Breadcrumbs:

- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
- [Property reference](resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132)
- waf_signatures.manual

<a id="canonical-0322202312203100-3310123001123202-3132021010223121-1022330122102132-0000321010123010-0322131222202210-3000110332011130-0303110031220210"></a>

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
manual = {}
```

<a id="canonical-2122330323221033-0231002322203210-0111322310033311-0011031031030023-0113132102011303-2013031231301010-1311202021212003-0300323313011302"></a>

## Direct properties — manual / 112231322120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2120302131332112-1220030131122100-1113231303202123-2303332012203023-0221211223323303-0210112232233122-0322312022122303-1121212300111111"></a>

## Next pages — manual / 112231322120 / 4

- [waf_signatures](resources--voltstack_site--reference--group-010.md#canonical-3001210003032302-1322031333003130-3131300313220213-0112200202131003-3012101010211301-1323112023133110-3121212221011012-2031210320020132)
- [xcsh_voltstack_site](../resources/voltstack_site.md#canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113)
