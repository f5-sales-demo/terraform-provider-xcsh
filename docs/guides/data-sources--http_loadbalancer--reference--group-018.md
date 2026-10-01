---
page_title: "xcsh_http_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer reference."
---

# xcsh_http_loadbalancer reference

<a id="canonical-2101013022030112-3223013000202330-2011220020010012-3133201021211303-1233011023203202-3302330201313320-3222123100202130-3310103003231002"></a>

## namespace property — code_base_integration / 030311331232 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2322032111233332-3211010121100030-1012110202231331-0313012013311313-2222020221330121-3232023232101011-3300132103230221-1331323201133113"></a>

<a id="canonical-3213223312232233-1333300000001022-2210231330132320-3010013131023121-0221313011331100-2022101213012100-3301223203021011-1131020220201322"></a>

## tenant property — code_base_integration / 030311331232 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2002221210013012-3021102001312230-1110312302200332-0333311330001231-0111011130300001-3200332013113130-2302123011211030-3111122331203322"></a>

## Next pages — code_base_integration / 030311331232 / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2321021311110321-2231002202230230-3200232012320302-2111311121123201-2233322300121121-0110123311132030-1011312220023000-1303333200312121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1132323323033033-3033003200130312-0100322011000201-0202030102303311-3211113023231333-2011130023030332-2300300322023023-0030330021022331"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — selected_repos / 111030032223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.api_discovery_from_code_scan](data-sources--http_loadbalancer--reference--group-017.md#canonical-2203231110333303-0020330220102233-1212031302200323-2010132310300200-2323303123313032-3023312023010120-1132113021000231-1003103022033213)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-2320032312022010-3331132210103323-0232120313033333-0113312230212321-2032121003100020-2123003310303010-1133202023032210-2102112302001211"></a>

Type: `"single"`. Computed.

Select which API repositories represent the LB applications.

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

<a id="canonical-2102122023323011-2013023021331222-0021112331002003-2232301032123202-1032032002031311-0112203033132303-0233201333311131-0311232201013232"></a>

## Direct properties — selected_repos / 111030032223 / 3

<a id="canonical-0011323012203331-2032311000310331-3202122012223300-2333110330021011-2322033322231203-0233333030011203-1211303131133132-0033323031012010"></a>

<a id="canonical-3110001121202233-2012131333333301-1120311001202033-0033321023213300-0301320310223030-1133012310300331-2111020030311321-0021322311011132"></a>

## api_code_repo property — selected_repos / 111030032223 / 4

Type: `["list", "string"]`. Computed.

Code repository which contain API endpoints.

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
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1312311111102010-2121112213312321-1321131221301030-3112301123133211-2031011303022203-1203310101321323-3002123131133223-3303020310232333"></a>

## Next pages — selected_repos / 111030032223 / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](data-sources--http_loadbalancer--reference--group-017.md#canonical-1030211010301012-1013000322300110-2302132223232131-1222201210323202-0102230030303233-2031113111010121-1002332233020312-1302102010200200)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3001200131212000-3321100200121300-1113221233313122-2013110303220031-1120101002332231-3020123311233011-0203030310301322-1223331233311201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232333031213130-1331102210303103-1303011031202011-2011332013030203-1233020331111022-3131123103311302-1330020320211321-0320122111313203"></a>

## enable_api_discovery.custom_api_auth_discovery — custom_api_auth_discovery / 301332022112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-3030011312103210-3330320131011033-0030222112001221-2133302103331323-2113202033123220-0003321312101102-2331200210311003-3000133212012032"></a>

Type: `"single"`. Computed.

API Discovery Advanced Settings. API Discovery Advanced settings.

Upstream description:

API Discovery Advanced settings.

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

<a id="canonical-0112133210211132-2212300022001313-0330100030222010-2011130312013032-1011021100330213-2132010210203331-0002010132021132-0221313230032021"></a>

## Direct properties — custom_api_auth_discovery / 301332022112 / 3

- [api_discovery_ref](data-sources--http_loadbalancer--reference--group-018.md#canonical-2012222213213103-0323031013101033-0220233220333033-3221313001311021-1221033213132101-0031210111310323-1021232100131311-3002130331232323): complete subsection reference.

<a id="canonical-1130111003231210-3213332023033020-0311323031032221-2133223123021310-3302330322230003-0102130331023113-0211023101021321-2100331322133002"></a>

## Next pages — custom_api_auth_discovery / 301332022112 / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](data-sources--http_loadbalancer--reference--group-018.md#canonical-2012222213213103-0323031013101033-0220233220333033-3221313001311021-1221033213132101-0031210111310323-1021232100131311-3002130331232323)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2012222213213103-0323031013101033-0220233220333033-3221313001311021-1221033213132101-0031210111310323-1021232100131311-3002130331232323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001111213213021-3300010000330223-2233301312101002-3230011022233331-0220331112220201-1322300220222200-1112313112021112-0112013033311310"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — api_discovery_ref / 103030222002 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [enable_api_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-018.md#canonical-3001200131212000-3321100200121300-1113221233313122-2013110303220031-1120101002332231-3020123311233011-0203030310301322-1223331233311201)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-2013001112102300-3331322223312130-1023303111203033-3112013233000302-3100221022232203-2121010000323103-3123310210332030-3300330232223300"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

<a id="canonical-3323300031311022-2223022130233023-1323331010330023-3223322223303031-1103233001232233-3023230311112101-0020012032320000-1010020211333301"></a>

## Direct properties — api_discovery_ref / 103030222002 / 3

<a id="canonical-0231222100201010-0213003203021221-2001233231133320-0203221011121100-1323030133132013-0332121222001113-2332313210211132-2033133000303303"></a>

<a id="canonical-3230311221030200-0232313302332303-3101120101200031-1103333213200323-1110011132200031-2120210101322003-0322033103331311-1211231232133231"></a>

## name property — api_discovery_ref / 103030222002 / 4

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

<a id="canonical-0011211300330002-1230010223113333-3332310202330221-2000131122213210-2110100003231312-3132310230022022-3012301010321221-0022010232022000"></a>

<a id="canonical-3201223321032130-0021333020021223-0202320302333013-1103110030322032-0201023110223323-2011101103021013-0313232013332301-2021120131113301"></a>

## namespace property — api_discovery_ref / 103030222002 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2211113201132212-2003220302132120-1131012330200220-3302011211010230-3021210033200003-3022212313212230-0311100013013213-1321111122201323"></a>

<a id="canonical-2201100110320313-2102232002131031-3300020130012000-2102001323000020-1132001221211331-2321210320002233-2002211120020011-3011001323020222"></a>

## tenant property — api_discovery_ref / 103030222002 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-1012232333322020-1101222232322121-1223310021331001-3012221300023102-3000102112021231-0011212122310231-1232133333103310-3310130213302230"></a>

## Next pages — api_discovery_ref / 103030222002 / 7

- [enable_api_discovery.custom_api_auth_discovery](data-sources--http_loadbalancer--reference--group-018.md#canonical-3001200131212000-3321100200121300-1113221233313122-2013110303220031-1120101002332231-3020123311233011-0203030310301322-1223331233311201)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2332023220311231-2022101012211013-2010123201123233-1232120333012130-3013001110130232-3300111322230100-2212222320203322-0333122112300110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202210012032300-2320122313330312-2223020223102101-3212123331220312-3223221002302333-0313123001321033-0012130312330313-2212120202323320"></a>

## enable_api_discovery.default_api_auth_discovery — default_api_auth_discovery / 213000120131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-1103332303031223-2033311233200011-3323121101020100-3220211333231103-3020032221203201-1112333320211211-2313301311222310-2322120213231000"></a>

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

<a id="canonical-1231110331022323-3003033001322003-3300203001200232-2302023320103231-3001231202010023-3000100112102201-0322103022101211-0112301212220001"></a>

## Direct properties — default_api_auth_discovery / 213000120131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323221102222223-2012101101312320-3310313302212130-1102103000322310-1310000011111112-3032332000102120-3320013010201030-1112112220033220"></a>

## Next pages — default_api_auth_discovery / 213000120131 / 4

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3012103000100232-1333022032021133-3322020200231210-3020203332310212-3211011202000020-2113003311123103-2223221301000003-3321311331100011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0013233311321203-2301122231320221-2103221212133113-2313113032132030-3012013013010001-0001201003211223-1211301210300120-1230112110301322"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — disable_learn_from_redirect_traffic / 303123021212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-0031233031223010-1132233020221232-3011202122123101-0000002212013121-1323312010013102-2122223132103200-1230302003000001-0301311221020033"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable learn from redirect traffic.

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

<a id="canonical-1130220222120002-0130330133123311-2123300201223111-0331200133201110-3033313103330200-3113323222001013-3212210222302003-2003332120233333"></a>

## Direct properties — disable_learn_from_redirect_traffic / 303123021212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321333213220023-2230223232332121-0032330211312032-3223132221030122-2313010221210323-1311320321001133-1210333332120223-2131232013120300"></a>

## Next pages — disable_learn_from_redirect_traffic / 303123021212 / 4

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2123111312022000-2333133110203330-0220223302200110-2202032222101322-1202332001000200-3311323320022011-1211213100322320-1233320302231323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010101003230312-1221301233333132-2131012002102233-1120301301033311-2012131303331111-3113022120132101-1203012130113223-0030233333001111"></a>

## enable_api_discovery.discovered_api_settings — discovered_api_settings / 312021202303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.discovered_api_settings

<a id="canonical-1311201022323300-3033102231303223-3011232200001302-1323133312200011-0000123220321013-3311223331201212-0321310113111123-0331220230230223"></a>

Type: `"single"`. Computed.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

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

<a id="canonical-0130002221211021-2220023321100322-3211123213203323-3122023013101122-1130112023100110-3330123212311300-2313310021231203-3013203120112113"></a>

## Direct properties — discovered_api_settings / 312021202303 / 3

<a id="canonical-3010130013211332-3330120210313101-2202233022303113-0110031322102213-1110231112020111-2211201000200323-1332213222130312-3123111232031310"></a>

<a id="canonical-2311131300311012-2221133122313113-0130020203233320-3320213330232021-2022302023203330-0323021003323220-2203323330311232-2332210012231033"></a>

## purge_duration_for_inactive_discovered_apis property — discovered_api_settings / 312021202303 / 4

Type: `"number"`. Computed.

Inactive discovered API will be deleted after configured duration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 7,
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
    "ves.io.schema.rules.uint32.lte": "7"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "7"
  }
}
```

<a id="canonical-1212031120302111-0003011220213200-2100002302311020-2202220302212121-1203202133203012-2321102231213030-0201202100132120-2113332002000313"></a>

## Next pages — discovered_api_settings / 312021202303 / 5

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3132323100013201-3210103000323202-0120303213201211-3200221003122313-0313121001120132-1231133311121230-1021201002210111-3031132030331113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2203230031113331-3113101310110030-1220021202022221-3002300133332001-0330202210023102-0222211023011011-2303020311212101-1101111321302231"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_learn_from_redirect_traffic / 322130012333 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-2021020330130301-0230203110102213-1230333213203223-0223022222212020-0011323332110112-0221002323330220-2033033023322232-2331100200002130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable learn from redirect traffic.

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

<a id="canonical-1021211332022003-0111312111321301-1031313003123212-2202202121013032-3102012311033101-1201121332330010-2332010013202332-2102310013131011"></a>

## Direct properties — enable_learn_from_redirect_traffic / 322130012333 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331220311232020-1123212022010202-3332220223302320-0230011303303133-3021100310203212-1122000203001321-3231111123001101-3213230303012222"></a>

## Next pages — enable_learn_from_redirect_traffic / 322130012333 / 4

- [enable_api_discovery](data-sources--http_loadbalancer--reference--group-017.md#canonical-3102231103303332-0011102023001333-1001101113313230-0310003203030002-0230021033330011-0020301130222310-2002330010130011-0220313000312223)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130120103002232-3322011322003101-3220121132133012-1012203203032013-3123121122021201-0301211032322001-2023130233101330-3301230132003313"></a>

## enable_challenge — enable_challenge / 101202121130 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_challenge

<a id="canonical-2230132011120203-0011333021033130-1010233313001013-1100313220213320-3231313122231210-2033232202103112-3100002200122100-0000223330000031"></a>

Type: `"single"`. Computed.

Configure auto mitigation i.e risk based challenges for malicious users.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

<a id="canonical-2133322003233303-2023322131330011-2033311212103112-3021323031302211-0023110120120000-0310021101122223-1111022323112102-1010103331321230"></a>

## Direct properties — enable_challenge / 101202121130 / 3

- [captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-0121030213210321-1322302210313013-3112332123323020-1321123131333121-2131113211232020-3223111112222001-0011102330021020-3101111331133011): complete subsection reference.

- [default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-3320320211033302-3033120013010222-0121120301332033-2233122323233320-1001302333013233-0310332112021003-3332022112202301-1223121323133021): complete subsection reference.

- [default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-3102030212122002-0233012113210031-3000311301112132-3002110312202303-0121122210100000-1221122020103022-0302100222232211-3112000303123131): complete subsection reference.

- [default_mitigation_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-3033312103102312-3220031122103201-3003103002210011-1102312201122202-1320110311222313-3313112222220202-2223313011001322-2111030302232012): complete subsection reference.

- [js_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2102131322013321-3103000102211223-2203223000013002-0030321001003323-3123330210130223-1111302200033320-2030022100133123-0333011120000212): complete subsection reference.

- [malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1320102211302230-1331211303022203-3011223102123110-3012311303213230-3233312022320330-0130203012003103-0020012330231303-3212030013231120): complete subsection reference.

<a id="canonical-3201120222312331-0130130022031133-1031003332033013-1322121133321300-0102321303013121-1320000121323331-1212301001323330-1201131233302201"></a>

## Next pages — enable_challenge / 101202121130 / 4

- [enable_challenge.captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-0121030213210321-1322302210313013-3112332123323020-1321123131333121-2131113211232020-3223111112222001-0011102330021020-3101111331133011)
- [enable_challenge.default_captcha_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-3320320211033302-3033120013010222-0121120301332033-2233122323233320-1001302333013233-0310332112021003-3332022112202301-1223121323133021)
- [enable_challenge.default_js_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-3102030212122002-0233012113210031-3000311301112132-3002110312202303-0121122210100000-1221122020103022-0302100222232211-3112000303123131)
- [enable_challenge.default_mitigation_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-3033312103102312-3220031122103201-3003103002210011-1102312201122202-1320110311222313-3313112222220202-2223313011001322-2111030302232012)
- [enable_challenge.js_challenge_parameters](data-sources--http_loadbalancer--reference--group-018.md#canonical-2102131322013321-3103000102211223-2203223000013002-0030321001003323-3123330210130223-1111302200033320-2030022100133123-0333011120000212)
- [enable_challenge.malicious_user_mitigation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1320102211302230-1331211303022203-3011223102123110-3012311303213230-3233312022320330-0130203012003103-0020012330231303-3212030013231120)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0121030213210321-1322302210313013-3112332123323020-1321123131333121-2131113211232020-3223111112222001-0011102330021020-3101111331133011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121233233103301-3112003011110122-2000210023303320-3313211302120301-0222321101003001-0013003113332301-3330022100013113-2213303233122303"></a>

## enable_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 022331113200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-3032100223311200-0221203213210113-1011331133000232-1130230313120232-2303131100131332-1120112133323303-0010201123012031-2301221131103322"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform captcha challenge Captcha challenge will be based on Google
Recaptcha. With this feature enabled, only clients that pass the captcha challenge will be allowed
to complete the HTTP request. When loadbalancer is configured to do Captcha Challenge, it will
redirect..

Upstream description:

Enables loadbalancer to perform captcha challenge

Captcha challenge will be based on Google Recaptcha.

With this feature enabled, only clients that pass the captcha challenge will be allowed to complete
the HTTP request.

When loadbalancer is configured to do Captcha Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have captcha challenge embedded in it. Client
will be allowed to make the request only if the captcha challenge is successful. Loadbalancer will
tag response header with a cookie to avoid Captcha challenge for subsequent requests.

CAPTCHA is mainly used as a security check to ensure only human users can pass through. Generally,
computers or bots are not capable of solving a captcha.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-0333302031110233-0022020320031302-2032231010112231-2232322213212110-2223112100132323-0002110201133030-2200223231102322-3232320001000011"></a>

## Direct properties — captcha_challenge_parameters / 022331113200 / 3

<a id="canonical-3202123230132221-1330023031201010-0030232332310133-3202202111022123-3100221022010032-1302013311221100-1212330320123023-1320221020213102"></a>

<a id="canonical-0311030213320001-3103003022232221-0123100101320033-2132200101030033-2013130031102003-1100230121302110-0301132330230120-0300133020312223"></a>

## cookie_expiry property — captcha_challenge_parameters / 022331113200 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-3132121102220233-2030221330210210-1301232021013322-0323023132121032-0103303103210133-1301131312012313-0201001222103232-1110023022100310"></a>

<a id="canonical-2300201132013203-3332101213110021-2131223320020323-1033032133223112-1020313001022123-0110332112101011-3320101333120111-3112331221021131"></a>

## custom_page property — captcha_challenge_parameters / 022331113200 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3202130300003111-2013012200111331-1200123011232203-3030310223031312-0321301120211312-2200101013301331-0023020331210220-1130310013300020"></a>

## Next pages — captcha_challenge_parameters / 022331113200 / 6

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3320320211033302-3033120013010222-0121120301332033-2233122323233320-1001302333013233-0310332112021003-3332022112202301-1223121323133021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232111003010130-0023111230001012-3111302300011333-1230131133002201-1330301113133313-3032123123211301-1213033312320131-2023233103011021"></a>

## enable_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 021331323033 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-1311212210201333-3310002323313231-3032331230311333-1033221302322211-3312323302012013-3222132213221130-0110000133330100-1122111220210130"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

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

<a id="canonical-3221322021201001-1011010312312223-1100212201301301-0020221130303132-0312122230133130-2131003200103112-0232132333310312-2213013322203032"></a>

## Direct properties — default_captcha_challenge_parameters / 021331323033 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131002003032223-0023112211321313-3332022221320302-3031220012121032-1110323203321023-2313221020332103-0213000032231001-2201032312223001"></a>

## Next pages — default_captcha_challenge_parameters / 021331323033 / 4

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3102030212122002-0233012113210031-3000311301112132-3002110312202303-0121122210100000-1221122020103022-0302100222232211-3112000303123131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013202300231022-1302030231120231-1002232302022301-0313201123300200-1102023001101131-0111301030221212-0121130032203332-1021201113221322"></a>

## enable_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 132223023323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-3020201203201200-2231312222121222-0210332232000300-1221220323011002-2223000220201322-2030132011122120-2133113113213100-0320333100201231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default js challenge parameters.

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

<a id="canonical-2110303302233011-2130212030233032-3033012223013131-1130130320012003-0231200212030202-3333310112000121-3322002320320011-3133103110211021"></a>

## Direct properties — default_js_challenge_parameters / 132223023323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220110322131322-1232023232231311-3221123130213001-3120331132313303-0123202220011302-0302211111001303-3211100230333222-3212221212122112"></a>

## Next pages — default_js_challenge_parameters / 132223023323 / 4

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3033312103102312-3220031122103201-3003103002210011-1102312201122202-1320110311222313-3313112222220202-2223313011001322-2111030302232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1003100220320122-0210223012201020-0321023230311330-2330301302230123-2133302210000210-0322031200131332-1112311010231201-3000021011000001"></a>

## enable_challenge.default_mitigation_settings — default_mitigation_settings / 133200111200 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.default_mitigation_settings

<a id="canonical-1011011120021113-1220233020132201-0311312110031312-2330003233133123-3103102211211231-0100313033133322-1033132133130301-2213222232020200"></a>

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

<a id="canonical-3200012122132132-2220210222200200-1220333223002310-1130330023000303-0201222333113003-1103022213103313-0212113131331331-1231213213133122"></a>

## Direct properties — default_mitigation_settings / 133200111200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1000233201230333-3120322210323112-1131031023002133-2002023333321033-0120100033122213-3131012331220121-0313323131332311-2333021332332012"></a>

## Next pages — default_mitigation_settings / 133200111200 / 4

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2102131322013321-3103000102211223-2203223000013002-0030321001003323-3123330210130223-1111302200033320-2030022100133123-0333011120000212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232322301223133-0232002033003112-3231122222123013-0120312211032220-0232300130110033-1101132002122121-1121121231331221-2231132201030130"></a>

## enable_challenge.js_challenge_parameters — js_challenge_parameters / 022020100210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.js_challenge_parameters

<a id="canonical-0112022001032023-3012003032213000-3102330113232230-3110012131320122-0033013133210200-1023323100101031-3313321131332302-0002132233000130"></a>

Type: `"single"`. Computed.

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript. With this feature enabled, only clients that are capable of executing JavaScript(mostly
browsers) will be allowed to complete the HTTP request. When loadbalancer is configured to do..

Upstream description:

Enables loadbalancer to perform client browser compatibility test by redirecting to a page with
JavaScript.

With this feature enabled, only clients that are capable of executing JavaScript(mostly browsers)
will be allowed to complete the HTTP request.

When loadbalancer is configured to do JavaScript Challenge, it will redirect the browser to an HTML
page on every new HTTP request. This HTML page will have JavaScript embedded in it. Loadbalancer
chooses a set of random numbers for every new client and sends these numbers along with an encrypted
answer with the request such that it embed these numbers as input in the JavaScript. JavaScript will
run on the requester browser and perform a complex Math operation. Script will submit the answer to
loadbalancer. Loadbalancer will validate the answer by comparing the calculated answer with the
decrypted answer (which was encrypted when it was sent back as reply) and allow the request to the
upstream server only if the answer is correct. Loadbalancer will tag response header with a cookie
to avoid JavaScript challenge for subsequent requests.

JavaScript challenge serves following purposes \* Validate that the request is coming via a browser
that is capable for running JavaScript \* Force the browser to run a complex operation, f(X), that
requires it to spend a large number of CPU cycles. This is to slow down a potential DoS attacker by
making it difficult to launch a large request flood without having to spend even larger CPU cost at
their end.

You can enable either JavaScript challenge or Captcha challenge on a virtual host.

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

<a id="canonical-3101310112232122-3111303223132333-2321221202011211-0113021013100321-3202203311111023-0200220103322123-2333301211130322-1002331310232311"></a>

## Direct properties — js_challenge_parameters / 022020100210 / 3

<a id="canonical-3200331223320103-3201222111012220-0322120130303032-0223222312202200-0332332303301333-1110220311013032-2302313131222131-2120210333001000"></a>

<a id="canonical-0020032011111031-2221303330123100-2123031211121021-2113310211213100-0202301320203132-0102010112013003-2301011113301300-0221131301113101"></a>

## cookie_expiry property — js_challenge_parameters / 022020100210 / 4

Type: `"number"`. Computed.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 86400,
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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0302033301101101-0133210200223121-0312331030323012-2230320331321203-2010320312210001-1222232201203133-1201021013010110-1300220033200113"></a>

<a id="canonical-3011300123302323-0002331013001310-0301011221110113-0023303122223223-3310210102220112-0302021102211220-0232001211230321-1223033333033333"></a>

## custom_page property — js_challenge_parameters / 022020100210 / 5

Type: `"string"`. Computed.

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format.

Upstream description:

Custom message is of type uri\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in Base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". Base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 65536,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 65536,
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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1130132210022130-3323002222032031-0322210310110213-3200212000101011-3202001131221122-1321222100003313-2311310101121101-3003012230230031"></a>

<a id="canonical-0000333322111110-0333111131000302-2313020311302113-2321123221001222-3121020111002213-3000002333302100-3012331312020211-3020103203202210"></a>

## js_script_delay property — js_challenge_parameters / 022020100210 / 6

Type: `"number"`. Computed.

Delay introduced by JavaScript, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1000
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1000",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2300231113001321-3003200222013000-1203120221331022-1002211023001122-1211030023331322-3123112111333123-2200322301132303-2113022230100221"></a>

## Next pages — js_challenge_parameters / 022020100210 / 7

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1320102211302230-1331211303022203-3011223102123110-3012311303213230-3233312022320330-0130203012003103-0020012330231303-3212030013231120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131121100220011-1033322123013020-1101131201021110-1110323122303203-0120320010102232-3213202323211133-1330102203122022-0101010100101200"></a>

## enable_challenge.malicious_user_mitigation — malicious_user_mitigation / 003323321232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- enable_challenge.malicious_user_mitigation

<a id="canonical-1223221130111102-0330010312211300-2232203330222213-3131013211320030-1302220223230013-1320333032212221-0313111033023212-3131213312101012"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

<a id="canonical-2301020112313021-3011212131100221-0220232203203030-3123211103220220-0310331323000221-1313333123301200-0332232233232331-1000133223212302"></a>

## Direct properties — malicious_user_mitigation / 003323321232 / 3

<a id="canonical-3012221002001300-3003303023032311-2300000032010213-3033101213210332-2213303321321112-1301301203110200-2103102332120001-0232303113212013"></a>

<a id="canonical-1121000311023310-3221022023321103-3322310203323112-1102130122321231-3110000333022211-0313210301310223-2311301203321303-1301022101222222"></a>

## name property — malicious_user_mitigation / 003323321232 / 4

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

<a id="canonical-3331120333300000-0210220022321232-2212031122321230-1100303321002000-3321132232333330-3201230310100210-3103030000100311-2023002102020023"></a>

<a id="canonical-1032311222001311-0110033111102132-2233210131010130-0201022113111122-0121123102003302-0221202102323320-1201131112013212-2220001020200122"></a>

## namespace property — malicious_user_mitigation / 003323321232 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-2223021212333312-1332213102123031-3000023132103203-2330311123203103-0203122232233131-3013303003110120-2301210032203030-1001333110330023"></a>

<a id="canonical-2123213313131221-0312211100220200-0013233001200333-3220033021300000-0021331300203003-1110023220333110-3013312023210333-0332013120100030"></a>

## tenant property — malicious_user_mitigation / 003323321232 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3101321010232030-0222312223020300-0322020212112010-1210113330312310-0001103301230213-0010213322212101-3132113131231100-2120230330000131"></a>

## Next pages — malicious_user_mitigation / 003323321232 / 7

- [enable_challenge](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011200023030232-1103032103102003-2020010023111201-0113013102133101-1133013020111332-3130102202020210-0213300311210300-1123323322102222)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2222220232120233-1130313100210223-1110323001300223-2031010002311213-3032120333310102-3022103310121023-0103023331030110-0220211011320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123313103013301-1313033233311031-1101322033333312-1010030003331020-3223212110310321-3330011012201233-2100330202221313-0120330001333230"></a>

## enable_ip_reputation — enable_ip_reputation / 121023210103 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_ip_reputation

<a id="canonical-1303123101122023-3103220103322001-2320130011231101-2030202013121330-2011333132210133-2222302130220303-2222103112211003-3201100223033010"></a>

Type: `"single"`. Computed.

IP Threat Category List. List of IP threat categories.

Upstream description:

List of IP threat categories.

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

<a id="canonical-0331203121311120-2331303232331010-1303303230301322-3302231130310021-3233032122003101-1302112001131233-0323220132202003-0110112011330311"></a>

## Direct properties — enable_ip_reputation / 121023210103 / 3

<a id="canonical-1012112323213010-0121011310311223-0133232321330210-1032112020223022-0211133113012222-0330031303122203-1103303321001021-3120000302000011"></a>

<a id="canonical-0122130002323033-2110303221210312-0022201210131202-0133211003102223-0201111203230102-1313221122033101-1301120301000021-0103033133110000"></a>

## ip_threat_categories property — enable_ip_reputation / 121023210103 / 4

Type: `["list", "string"]`. Computed.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1102313230322131-0120301230231200-3010321300031223-3021211100112001-2311011001332212-1301022333320022-3123010130121112-1012102202131323"></a>

## Next pages — enable_ip_reputation / 121023210103 / 5

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3112321013223011-0232120223301113-2123300000310001-2303013212022230-1020121313200102-3113021131130312-3132303103233220-0130203210202122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112003120300230-0301321200230231-2331131133332102-2120003031200131-2030201200033103-2301203133113100-1233300301312023-1012232010200101"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / 122002230223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_malicious_user_detection

<a id="canonical-2110332331011132-0112220203021031-2310200300312113-0000113031331323-0022111100132113-1010312111223221-0011212101030212-2221110322223112"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable malicious user detection.

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

<a id="canonical-2233132133013102-2133001101010323-3223011310102301-0333300300221122-0221102332003031-2120113113001230-2130133032313130-3120012210230222"></a>

## Direct properties — enable_malicious_user_detection / 122002230223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211300213133220-1220311002231321-3010133122020103-1332310211313030-0333133112222112-1221030313023300-3212130022111323-1010031012302222"></a>

## Next pages — enable_malicious_user_detection / 122002230223 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1111231022131133-1331313200113310-3310302011012003-0030332010313023-0230102231302212-2013001002111223-2120221113320103-0133233131213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123311123112230-1001013313110202-0110121112300022-0230110310113312-3313010211213222-3010032011301323-0330201120323123-2221203001212301"></a>

## enable_threat_mesh — enable_threat_mesh / 330332233113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_threat_mesh

<a id="canonical-3202002102313212-2121133022323302-2332320333313223-1312303022323030-3030133001022000-0320122330221310-0003032232310211-2110113111030212"></a>

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

<a id="canonical-0133120301313101-3313000200233111-3301201031023021-0131103231031221-2113332020202300-3110000102231120-3111303033032220-3302300311212233"></a>

## Direct properties — enable_threat_mesh / 330332233113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3003321302212130-3213130020322203-1230220333202131-2333223011220131-1010321211231331-1132101113230212-2211332212120132-0001331320011312"></a>

## Next pages — enable_threat_mesh / 330332233113 / 4

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0223302231131311-1212022331012012-1201212100221020-3222300100201333-2302211010002001-3333320223120100-2001323010222333-0112102133220100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303100231010101-2310222120103021-0123300130020201-1320223220321320-0231201003303011-2332003222332330-3011302020122000-3201330123021023"></a>

## enable_trust_client_ip_headers — enable_trust_client_ip_headers / 311331322230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- enable_trust_client_ip_headers

<a id="canonical-2200022210121310-3030222023303302-1130102212110230-2323302120331110-2331233111002022-1020110033203002-0231333331120233-3013121212303110"></a>

Type: `"single"`. Computed.

Trust Client IP Headers List. List of Client IP Headers.

Upstream description:

List of Client IP Headers.

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

<a id="canonical-0132100220000310-0010210332112321-3030121023000230-2323233032110032-0012302213311012-1311133020333013-1011300130022301-0000102230000332"></a>

## Direct properties — enable_trust_client_ip_headers / 311331322230 / 3

<a id="canonical-2311033000302102-1230312213303000-0310232120202323-2122022102222021-1231201030320330-1333331122333200-2222021213313203-1231222012123122"></a>

<a id="canonical-0132223112220133-3022121230101103-0112112213331321-1333010212003112-1103301113012310-1201111003202333-3221032303333130-1313310133231033"></a>

## client_ip_headers property — enable_trust_client_ip_headers / 311331322230 / 4

Type: `["list", "string"]`. Computed.

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined..

Upstream description:

Define the list of one or more Client IP Headers. Headers will be used in order from top to bottom,
meaning if the first header is not present in the request, the system will proceed to check for the
second header, and so on, until one of the listed headers is found. If none of the defined headers
exist, or the value is not an IP address, then the system will use the source IP of the packet. If
multiple defined headers with different names are present in the request, the value of the first
header name in the configuration will be used. If multiple defined headers with the same name are
present in the request, values of all those headers will be combined. The system will read the
right-most IP address from header, if there are multiple IP addresses in the header value. For
X-Forwarded-For header, the system will read the IP address(rightmost - 1), as the client IP.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5,
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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3001233232222222-0000233312221130-2131130221301203-1100111031111330-1332320131110032-2220011223211233-0102032222310312-0102123030123200"></a>

## Next pages — enable_trust_client_ip_headers / 311331322230 / 5

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203231330111332-1223101331302103-2332023222001122-0332123231101030-1233103133331020-2222021021032021-2223211020230100-1222311310321311"></a>

## graphql_rules — graphql_rules / 302332210121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- graphql_rules

<a id="canonical-3003300301222323-0112102102301002-0113203300032213-3032201212021111-3200003212221223-0030112023020101-3221203102131101-2313232223103031"></a>

Type: `"list"`. Computed.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

<a id="canonical-1303030001312201-2120022002332303-2011033001331120-3131033313313312-2303300102330132-3120323033010033-2220101112022332-1320202112000303"></a>

## Direct properties — graphql_rules / 302332210121 / 3

- [any_domain](data-sources--http_loadbalancer--reference--group-018.md#canonical-1103202203211103-3211330223021312-1123113112120011-1233031121112012-1021100231020021-2020333210313301-0021123232303112-0300000110232202): complete subsection reference.

<a id="canonical-1211112232001330-0011213223031200-3123002002221022-1120003133100000-0220321303321130-0300321333023110-2010123312322011-2002310220001132"></a>

<a id="canonical-1202133120303200-2012000333212320-3310032113231200-2032201001323303-1222330323322220-3210000003100223-2001212022102201-0320023223121203"></a>

## exact_path property — graphql_rules / 302332210121 / 4

Type: `"string"`. Computed.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /graphql.

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-2311132103301320-3130110111011310-3121301033220113-2200220302312202-0331213331001221-2330133200320132-0332010201110230-1012211232011022"></a>

<a id="canonical-1203213322031220-0031131333011223-3010320101111021-3212230122130001-2132203013011010-1022333001000233-1132001012100300-1001113230322330"></a>

## exact_value property — graphql_rules / 302332210121 / 5

Type: `"string"`. Computed.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010): complete subsection reference.

- [metadata](data-sources--http_loadbalancer--reference--group-018.md#canonical-0331131331213103-3103312023301300-0303321120101113-1323132201332122-3010203210003332-0132112222323313-0300230120102230-0200300120030102): complete subsection reference.

- [method_get](data-sources--http_loadbalancer--reference--group-018.md#canonical-2213131120013323-3301221131310223-0021331301012202-2012020231122113-2320030003221211-2011022111211030-2033112121032000-0121013330103132): complete subsection reference.

- [method_post](data-sources--http_loadbalancer--reference--group-018.md#canonical-3222113113333222-1110230032332100-3230132102303033-3021033313031223-1323103130002112-1230113120111312-3022313200321012-2021113111302130): complete subsection reference.

<a id="canonical-2322012123010321-3223301302302110-0330011221003331-1333220231001233-0222122201303130-3321031333001133-0301203311100033-3112103221311203"></a>

<a id="canonical-0202133131201012-0100102120320302-3220212301122003-0312333132321133-0302031031303310-1020033231332320-1013231302130102-1333112121222113"></a>

## suffix_value property — graphql_rules / 302332210121 / 6

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-0201312010103113-3020332110321232-3120301212032220-0121000011303332-0032322112322301-2032000023011131-2333210110321210-3121311203312030"></a>

## Next pages — graphql_rules / 302332210121 / 7

- [graphql_rules.any_domain](data-sources--http_loadbalancer--reference--group-018.md#canonical-1103202203211103-3211330223021312-1123113112120011-1233031121112012-1021100231020021-2020333210313301-0021123232303112-0300000110232202)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- [graphql_rules.metadata](data-sources--http_loadbalancer--reference--group-018.md#canonical-0331131331213103-3103312023301300-0303321120101113-1323132201332122-3010203210003332-0132112222323313-0300230120102230-0200300120030102)
- [graphql_rules.method_get](data-sources--http_loadbalancer--reference--group-018.md#canonical-2213131120013323-3301221131310223-0021331301012202-2012020231122113-2320030003221211-2011022111211030-2033112121032000-0121013330103132)
- [graphql_rules.method_post](data-sources--http_loadbalancer--reference--group-018.md#canonical-3222113113333222-1110230032332100-3230132102303033-3021033313031223-1323103130002112-1230113120111312-3022313200321012-2021113111302130)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1103202203211103-3211330223021312-1123113112120011-1233031121112012-1021100231020021-2020333210313301-0021123232303112-0300000110232202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121022102231322-1123330120023232-1023023220130301-1212100312303122-2100122223001333-2023310123303313-0133021103221312-1102123001123031"></a>

## graphql_rules.any_domain — any_domain / 003000113222 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.any_domain

<a id="canonical-2312213232002023-1020230030200310-2130312302100223-0030203120121230-3311030320001021-2311211322310010-1303212222201231-2323103012123302"></a>

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

<a id="canonical-0112132010200330-0120320110231131-3323002112201001-1131202023220011-0010331012231000-0321311222232100-3021330013110103-1320303212232112"></a>

## Direct properties — any_domain / 003000113222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3012131113111233-2231221110231110-0113130131000033-1220331201131323-0100112112320023-1221101022131100-2200101210232200-2233210020113200"></a>

## Next pages — any_domain / 003000113222 / 4

- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133010011333231-2121223202122113-2030021111232103-1022230011213233-0110232101021003-3003233102001133-2102330230300122-0102202223222021"></a>

## graphql_rules.graphql_settings — graphql_settings / 312321033032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.graphql_settings

<a id="canonical-0303023310303203-1012022212101111-1312302301131313-2222022312011131-3130103102331113-2213013223222133-2201133320312023-3220233121132301"></a>

Type: `"single"`. Computed.

Configuration parameter for graphql settings.

Upstream description:

GraphQL configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-allow_introspection_queries_choice": "[\"disable_introspection\",\"enable_introspection\"]"
}
```

<a id="canonical-1032233322233100-1030102211223010-1312103112311232-2001203222002111-0201322222303021-2122130121033222-2031102213121123-2302020303230332"></a>

## Direct properties — graphql_settings / 312321033032 / 3

- [disable_introspection](data-sources--http_loadbalancer--reference--group-018.md#canonical-0012210321003021-0131310211201301-2231210000010011-3012020310011010-1321113220200330-3323113000331123-0011121210023111-0123133333310232): complete subsection reference.

- [enable_introspection](data-sources--http_loadbalancer--reference--group-018.md#canonical-2102110122303030-2330303013000003-2103130320031002-0103012222203332-2122133010233322-1101313001311033-2131321203213310-2103121010001121): complete subsection reference.

<a id="canonical-1113131222001120-2000213310301321-3003013322112102-2220000302321312-1323330131313231-2230212332121120-0010321332110221-0022121033331231"></a>

<a id="canonical-3322200100203302-0332013210331113-2300112023101113-0110311021202030-2221221023113001-0000232030132231-3223330023213121-3122102331330103"></a>

## max_batched_queries property — graphql_settings / 312321033032 / 4

Type: `"number"`. Computed.

Specify maximum number of queries in a single batched request.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-3121321230030223-1303132121033033-3130112323313232-0010001003010012-1131001212222102-0110133223200310-0231011003122303-1000310032012000"></a>

<a id="canonical-3201322313312332-0001312020310213-2120210003000331-2201302201102030-1222210031210332-0033103202303023-2132000323203111-1313201221031223"></a>

## max_depth property — graphql_settings / 312321033032 / 5

Type: `"number"`. Computed.

Specify maximum depth for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "20"
  }
}
```

<a id="canonical-0200000033130222-3332001233201213-1033222121130203-2111020203223012-2233202210313020-3202221030101003-1130223212210323-1130030123001101"></a>

<a id="canonical-1001211212230301-1213232331022010-0213212322321223-2201113033002302-0012002233331232-0331112000101112-0322320002110312-3210011033233310"></a>

## max_total_length property — graphql_settings / 312321033032 / 6

Type: `"number"`. Computed.

Specify maximum length in bytes for the GraphQL query.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 16386,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.uint32.lte": "16386"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "16386"
  }
}
```

<a id="canonical-3030203000220022-3012330022012003-2022022311321322-0303330321111210-3231230213132202-0101313112103221-1230023121202032-2130332120130121"></a>

## Next pages — graphql_settings / 312321033032 / 7

- [graphql_rules.graphql_settings.disable_introspection](data-sources--http_loadbalancer--reference--group-018.md#canonical-0012210321003021-0131310211201301-2231210000010011-3012020310011010-1321113220200330-3323113000331123-0011121210023111-0123133333310232)
- [graphql_rules.graphql_settings.enable_introspection](data-sources--http_loadbalancer--reference--group-018.md#canonical-2102110122303030-2330303013000003-2103130320031002-0103012222203332-2122133010233322-1101313001311033-2131321203213310-2103121010001121)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0012210321003021-0131310211201301-2231210000010011-3012020310011010-1321113220200330-3323113000331123-0011121210023111-0123133333310232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232001033131313-1030110111331110-1101212213120322-2003330032302331-3303331112032011-2223220001101330-3130113111312301-0312113300312312"></a>

## graphql_rules.graphql_settings.disable_introspection — disable_introspection / 221303230211 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- graphql_rules.graphql_settings.disable_introspection

<a id="canonical-1302310210100031-1123101200113103-3220003022122001-0302011303301231-3133301230200303-3303103303201031-3020112022331230-3233323012003221"></a>

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

<a id="canonical-3201302211013000-2012001111002322-1322013101100311-1013312023132001-1001033211023210-0122230011121333-0303312031210332-0010331221302123"></a>

## Direct properties — disable_introspection / 221303230211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2000310233020130-0111212301220231-3122333222331323-0332333010131220-0012010002100202-3010310102012110-1330302012011103-3101120200323330"></a>

## Next pages — disable_introspection / 221303230211 / 4

- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2102110122303030-2330303013000003-2103130320031002-0103012222203332-2122133010233322-1101313001311033-2131321203213310-2103121010001121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233320101223313-0110130222032031-3103002003013100-2103211310133103-1112203310213220-2223101020301021-0013010210231303-2031221100012301"></a>

## graphql_rules.graphql_settings.enable_introspection — enable_introspection / 332032200131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- graphql_rules.graphql_settings.enable_introspection

<a id="canonical-0212011322132211-0112331031021003-0021313310333023-1232300201333210-2111132120023311-2212210220310133-3103211000120111-3231333013322012"></a>

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

<a id="canonical-2312023121011001-1310012120233121-0331000003303232-2101312012200233-1123333322322130-1032201103020102-0332010312300233-1212021330131331"></a>

## Direct properties — enable_introspection / 332032200131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3030233012133333-2001131302032030-3211232210022332-0210230023132221-1132003223031021-3133002103101033-3033133122013200-2011200330113111"></a>

## Next pages — enable_introspection / 332032200131 / 4

- [graphql_rules.graphql_settings](data-sources--http_loadbalancer--reference--group-018.md#canonical-0123001120213220-2233121321230101-1121121021122232-2222133312301221-0002121031222210-2213300001322210-1121112013301123-0233210112333010)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0331131331213103-3103312023301300-0303321120101113-1323132201332122-3010203210003332-0132112222323313-0300230120102230-0200300120030102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010232220113001-1013232130222113-2310230002013133-1211131011201301-3112333211103320-2111002232102213-0233021302233133-3203300131102312"></a>

## graphql_rules.metadata — metadata / 212223103313 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.metadata

<a id="canonical-1013113033301002-0322300022011321-1112033212321301-3121303300023132-2032223231230021-0030213320030110-2220330200330200-1213200132201321"></a>

Type: `"single"`. Computed.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

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

<a id="canonical-2022111113020220-3031232231023133-0311310313110220-3022200122100233-1121110013133113-3000202030133302-3012112202220331-0122031001331130"></a>

## Direct properties — metadata / 212223103313 / 3

<a id="canonical-3011103202022300-2311330331122223-0003133313301230-0102000020133230-1110320312231021-2321221211220231-3113223202032020-3123002313310303"></a>

<a id="canonical-1122312121220033-3323113130112301-0200100010113010-1002110320102030-0110130213310001-0312332021311221-2011031020302002-2100321001222233"></a>

## description_spec property — metadata / 212223103313 / 4

Type: `"string"`. Computed.

Description. Human readable description.

<a id="canonical-3131323021112101-3333133230112300-1230111301231301-3131231331233030-2001012102031303-0303010011011033-1123132323131230-3330001312032330"></a>

<a id="canonical-1322013112020320-0231321110233221-0020123302101120-1330322033111213-3013102002031333-0201300333202201-2012323133320000-3311001313010103"></a>

## name property — metadata / 212223103313 / 5

Type: `"string"`. Computed.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Receipt-pinned upstream constraints:

```json
{
  "minLength": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.min_len": "1",
    "ves.io.schema.rules.string.ves_object_name": "true"
  }
}
```

<a id="canonical-2322233322132310-2001311002001333-1122112103320302-1012001312032231-2031222101102110-3111231332200023-2332303013230100-1203311011110022"></a>

## Next pages — metadata / 212223103313 / 6

- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2213131120013323-3301221131310223-0021331301012202-2012020231122113-2320030003221211-2011022111211030-2033112121032000-0121013330103132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003321131222221-0303022101321010-1002310111222132-3301033133200222-2032131013311233-3331120013321111-0211131320221102-2023313330013032"></a>

## graphql_rules.method_get — method_get / 312010120323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.method_get

<a id="canonical-3122110333003031-3301301230323030-1033002311011021-1232332313312032-3313311202212020-1222121230101132-0112213210210022-2200202033210210"></a>

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

<a id="canonical-0130223200113110-3321210220220221-3220210300301232-2131203000323130-3212013312302232-2002200233223220-2033122101123221-2011033311302211"></a>

## Direct properties — method_get / 312010120323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131011013311220-2022103121231321-0100222131030333-2102300003130202-0031013010321110-3020002021022133-3012302110211132-2001201122020300"></a>

## Next pages — method_get / 312010120323 / 4

- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3222113113333222-1110230032332100-3230132102303033-3021033313031223-1323103130002112-1230113120111312-3022313200321012-2021113111302130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030030220131300-3310100232120121-1213121031032230-1031233211232221-2320311222203013-1022212312330121-2132331013102011-1313103130221020"></a>

## graphql_rules.method_post — method_post / 200321203210 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- graphql_rules.method_post

<a id="canonical-2102133311231221-3231102332301020-1200103133131020-3333003021112000-1020003202011201-1111202302312110-1130120110130303-1122321033203133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for method post.

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

<a id="canonical-0133310210030002-0021222000330022-2213312003032013-2221212123203122-3123201020132011-3132001121221310-0012002200221220-2022213333330032"></a>

## Direct properties — method_post / 200321203210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003003132321232-1322211111131110-0233103102121301-2103110103313200-3023202211120022-0323222012323002-2222311123221102-3102200332120113"></a>

## Next pages — method_post / 200321203210 / 4

- [graphql_rules](data-sources--http_loadbalancer--reference--group-018.md#canonical-0103130023112220-0323202202232201-0100110021021230-0011021230100313-0322001310012013-3211213332301231-0231033133330111-0023122020322232)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1221000122101103-3100213033303232-2132010133302010-2012231203222223-2220001112313332-0002130322200033-3133223330023000-2102003233113221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301013210122310-3210231010213302-2002013230010311-0200020002223023-0130200031212210-1113110120201022-3323112300223130-3312322202020210"></a>

## http — http / 002120132123 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- http

<a id="canonical-0330312012020130-1032221113102101-2323210100123332-0323300213303303-1312010310011021-1100003013333313-1230010000210213-0202111322331322"></a>

Type: `"single"`. Computed.

\[OneOf: http, https, https\_auto\_cert; Default: https\_auto\_cert\] HTTP Choice. Choice for
selecting HTTP proxy.

Upstream description:

Choice for selecting HTTP proxy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]"
}
```

OneOf alternatives in this subsection:

- [http](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330312012020130-1032221113102101-2323210100123332-0323300213303303-1312010310011021-1100003013333313-1230010000210213-0202111322331322)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-1230020122213233-2010123001220332-2010332323100211-3023001312301001-0100321202103310-1030231133101202-0011031010011003-2002013130023101)
- [https_auto_cert](data-sources--http_loadbalancer--reference--group-019.md#canonical-2110233032332000-2230100230102301-1331113121321211-2111113312132303-0303213311122331-0023321211013320-1001321133320230-1213120312211202)

Select alternatives according to the provider validators above.

<a id="canonical-1201111003131033-0223321011332003-1332330130202110-3121102330222303-2030000320212130-0110202010202110-1020221311100212-1132132110211011"></a>

## Direct properties — http / 002120132123 / 3

<a id="canonical-2333320123122231-0113023133300110-1130333332032230-2323021310122230-2331123132221110-0100301232330020-1120323203310123-1003320220301121"></a>

<a id="canonical-3220013032313220-1020031332123213-3213112313302313-3300120312033030-0222113302203302-0313102123222033-0112231203013132-2302131320303212"></a>

## dns_volterra_managed property — http / 002120132123 / 4

Type: `"bool"`. Computed.

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

Upstream description:

DNS records for domains will be managed automatically by F5 Distributed Cloud. As a prerequisite,
the domain must be delegated to F5 Distributed Cloud using Delegated domain feature or a DNS CNAME
record should be created in your DNS provider's portal.

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

<a id="canonical-2300222103020010-2213302110302023-1121302331202330-1123202110201310-2323230011231322-0200231233233322-1313131212303121-3012220023002123"></a>

<a id="canonical-2120110010101033-1323022000120231-1230032001212303-2322221203132010-1230030123111232-1132221033110121-0021201133302232-1130210320120333"></a>

## port property — http / 002120132123 / 5

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTP port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTP port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3222322022130121-3123001301122303-1010323020213032-3101302222321312-3133311303110113-2102020021203222-1102230210101010-3333220232320010"></a>

<a id="canonical-0312201020311212-1130203311313003-3203132112333222-3223330001210302-2023101010310002-3123223221033302-1330123111010232-3313020331200020"></a>

## port_ranges property — http / 002120132123 / 6

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0030130302301031-3101011232003222-1131331323321113-0223302012210330-0012320003301300-3333220001030031-3020212030132320-1322123023203220"></a>

## Next pages — http / 002120132123 / 7

- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0212320231230202-3111023303203223-2220311021302312-1131131120012200-2020320033301303-3212313011131130-1103013011232032-0011102130110012"></a>

## https — https / 023232103212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- https

<a id="canonical-1230020122213233-2010123001220332-2010332323100211-3023001312301001-0100321202103310-1030231133101202-0011031010011003-2002013130023101"></a>

Type: `"single"`. Computed.

Choice for selecting HTTP proxy with bring your own certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-default_lb_choice": "[\"default_loadbalancer\",\"non_default_loadbalancer\"]",
  "x-ves-oneof-field-path_normalize_choice": "[\"disable_path_normalize\",\"enable_path_normalize\"]",
  "x-ves-oneof-field-port_choice": "[\"port\",\"port_ranges\"]",
  "x-ves-oneof-field-server_header_choice": "[\"append_server_name\",\"default_header\",\"pass_through\",\"server_name\"]",
  "x-ves-oneof-field-tls_certificates_choice": "[\"tls_cert_params\",\"tls_parameters\"]"
}
```

<a id="canonical-0133021200020131-1010032132022303-1101312112312332-2013121002323023-0121122113220133-1013020023320001-1313013030012321-1323333030113003"></a>

## Direct properties — https / 023232103212 / 3

<a id="canonical-2330100312310220-1130123313303213-2320002233211023-2221122020330233-1322022213302320-0200010301232131-0211201330012232-0311101032031331"></a>

<a id="canonical-1010310022131223-2211333003201131-3302212322131113-3101300100301331-1232123222131301-3331031130013132-1320211322221010-3130312320322330"></a>

## add_hsts property — https / 023232103212 / 4

Type: `"bool"`. Computed.

Add HTTP Strict-Transport-Security response header.

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

<a id="canonical-1312321213302200-2311130132312113-2221003330200222-3013300111111203-2332101010302020-1332203121111130-1222012020200203-1102221010301231"></a>

<a id="canonical-1330021220022031-0012131012301331-2321330023023201-0331310023321122-1312330011103131-1231303023320031-0331023101103301-3032212221103133"></a>

## append_server_name property — https / 023232103212 / 5

Type: `"string"`. Computed.

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Upstream description:

Exclusive with \[default\_header pass\_through server\_name\] Define the header value for the header
name “server”. If header value is already present, it is not overwritten and passed as-is.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332): complete subsection reference.

<a id="canonical-1001301223000032-1020300012030222-1203311102100300-2030223230203003-2232111200202103-2122121330003311-1120203120321212-2301212021200331"></a>

<a id="canonical-0313022203230120-2223333120231230-1130202033020200-0301220033132330-3311002113333211-1032230113110131-2232213113203003-0320221232210100"></a>

## connection_idle_timeout property — https / 023232103212 / 6

Type: `"number"`. Computed.

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed.

Upstream description:

The idle timeout for downstream connections. The idle timeout is defined as the period in which
there are no active requests. When the idle timeout is reached the connection will be closed. Note
that request based timeouts mean that HTTP/2 PINGs will not keep the connection alive. This is
specified in milliseconds. The default value is 2 minutes.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 600000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.lte": "600000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "600000"
  }
}
```

- [default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-1011220220213223-0023022311131111-3321100321212103-2302012100332100-3310212111103331-2030012330210201-1100032222211300-2333103131312233): complete subsection reference.

- [default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1123003233230101-3321333100231312-1212220312212102-3022231131212132-1132223102003323-3030022222021003-1030200021021213-1030221101223101): complete subsection reference.

- [disable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-1313230312203121-0221130100202030-2130030031311110-0000203003102103-0330223030021321-3212202011001322-2323331303122122-3302211302001001): complete subsection reference.

- [enable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-3211010213301321-0310002022003033-2320121222112212-3322033031120213-3201231212313320-2233231203130122-3211232302302121-2012120022302233): complete subsection reference.

- [http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021): complete subsection reference.

<a id="canonical-3031001212231120-3111211022232210-0331103001030120-0113311233000023-2322202320221102-3120231023311013-3010302033101002-0012201010103110"></a>

<a id="canonical-1001212303003123-3323030333103313-0211021300100001-3212222030011203-0101133102012322-1222100121332310-2330232000130013-1001113302121200"></a>

## http_redirect property — https / 023232103212 / 7

Type: `"bool"`. Computed.

HTTP Redirect to HTTPS. Redirect HTTP traffic to HTTPS.

Upstream description:

Redirect HTTP traffic to HTTPS.

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

- [non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1101330312012103-0233323210311221-1033012102103323-0220330330101233-1110011203231222-2231210133200221-1022100012023223-3112010202021011): complete subsection reference.

- [pass_through](data-sources--http_loadbalancer--reference--group-018.md#canonical-3210232200212003-0002112320100202-1003013010223122-1001223222303210-2322123332210310-3322112011113221-2311013202301022-0201323130200031): complete subsection reference.

<a id="canonical-2221113200201003-0032023033301302-2020120122212023-0003333331133102-1331201130310022-0033023130203113-2332302223212000-2213322012001302"></a>

<a id="canonical-1312013231020220-2010021032112021-2232112310303333-3102311303021023-3110333232113022-1310232111022031-1003003301030313-2330302110002331"></a>

## port property — https / 023232103212 / 8

Type: `"number"`. Computed.

Exclusive with \[port\_ranges\] HTTPS port to Listen.

Upstream description:

Exclusive with \[port\_ranges\] HTTPS port to Listen.

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
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

<a id="canonical-3001111332001110-0303213033020310-3003021121310020-0100033122330120-1320310003000230-2020013321002122-0013333330232023-1201003030132111"></a>

<a id="canonical-2101331112300013-2320233202202130-3223200101013013-0320113313313313-3122102330122011-2013021332233323-2131021222001220-1331221122212121"></a>

## port_ranges property — https / 023232103212 / 9

Type: `"string"`. Computed.

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by '-'.

Upstream description:

Exclusive with \[port\] A string containing a comma separated list of port ranges. Each port range
consists of a single port or two ports separated by "-".

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3203301010030220-1312311200112222-1021030233110111-0000200320200320-3031233213222101-2213233322332111-1123301030323213-3023302310030333"></a>

<a id="canonical-0102022323321110-2311303011203022-2010121302323121-0323311203230212-3011102112030221-1030302313111021-1131032130332022-2333030030221312"></a>

## server_name property — https / 023232103212 / 10

Type: `"string"`. Computed.

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Upstream description:

Exclusive with \[append\_server\_name default\_header pass\_through\] Define the header value for
the header name “server”. This will overwrite existing values, if any, for the server header.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 8096,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 8096,
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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

- [tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302): complete subsection reference.

- [tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031): complete subsection reference.

<a id="canonical-0113321300212011-1021020000021122-1003213331100300-2013222312331232-2213301233022111-3310222202212010-3333300002313311-3203232131332100"></a>

## Next pages — https / 023232103212 / 11

- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- [https.default_header](data-sources--http_loadbalancer--reference--group-018.md#canonical-1011220220213223-0023022311131111-3321100321212103-2302012100332100-3310212111103331-2030012330210201-1100032222211300-2333103131312233)
- [https.default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1123003233230101-3321333100231312-1212220312212102-3022231131212132-1132223102003323-3030022222021003-1030200021021213-1030221101223101)
- [https.disable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-1313230312203121-0221130100202030-2130030031311110-0000203003102103-0330223030021321-3212202011001322-2323331303122122-3302211302001001)
- [https.enable_path_normalize](data-sources--http_loadbalancer--reference--group-018.md#canonical-3211010213301321-0310002022003033-2320121222112212-3322033031120213-3201231212313320-2233231203130122-3211232302302121-2012120022302233)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.non_default_loadbalancer](data-sources--http_loadbalancer--reference--group-018.md#canonical-1101330312012103-0233323210311221-1033012102103323-0220330330101233-1110011203231222-2231210133200221-1022100012023223-3112010202021011)
- [https.pass_through](data-sources--http_loadbalancer--reference--group-018.md#canonical-3210232200212003-0002112320100202-1003013010223122-1001223222303210-2322123332210310-3322112011113221-2311013202301022-0201323130200031)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_parameters](data-sources--http_loadbalancer--reference--group-019.md#canonical-2223111233023031-1102331112100211-2213002320223310-1321030210202102-0101220103231020-1022131232001003-3130213120020031-1113011302100031)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310001210132030-2231232213032313-1300312300020203-0113330103322022-0020301202322131-0032001310030333-3001010331231121-3113031023213203"></a>

## https.coalescing_options — coalescing_options / 120323012132 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.coalescing_options

<a id="canonical-0202223120121212-2311130013123131-0220023023210102-3100231213033113-1213022002201210-1222111202100133-2213311313000211-1112201321201201"></a>

Type: `"single"`. Computed.

TLS connection coalescing configuration (not compatible with mTLS).

Upstream description:

TLS connection coalescing configuration (not compatible with mTLS)

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-coalescing_choice": "[\"default_coalescing\",\"strict_coalescing\"]"
}
```

<a id="canonical-1011322301210003-0112313113003303-3103130020303132-3201322222331201-3223020013032031-2300000100111212-1110123011023102-3211300301023032"></a>

## Direct properties — coalescing_options / 120323012132 / 3

- [default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-1230031131322300-0203330331131212-1121001011020002-0130210100301302-0120110010120333-3220333103210232-2111032023332220-2300213123311220): complete subsection reference.

- [strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-2310100133023231-2213123132321011-2120323001032322-1320201323220323-2302223031032123-2022333320330330-3030020322303211-2010112320230212): complete subsection reference.

<a id="canonical-2123202111021131-0012312031223230-1322133321130020-1131023201123202-3303322133223001-3201021021112110-2311232211212110-2200003201301112"></a>

## Next pages — coalescing_options / 120323012132 / 4

- [https.coalescing_options.default_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-1230031131322300-0203330331131212-1121001011020002-0130210100301302-0120110010120333-3220333103210232-2111032023332220-2300213123311220)
- [https.coalescing_options.strict_coalescing](data-sources--http_loadbalancer--reference--group-018.md#canonical-2310100133023231-2213123132321011-2120323001032322-1320201323220323-2302223031032123-2022333320330330-3030020322303211-2010112320230212)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1230031131322300-0203330331131212-1121001011020002-0130210100301302-0120110010120333-3220333103210232-2111032023332220-2300213123311220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001110003210211-0331131301112212-0230312313030323-2320022222331010-1223302020212300-0032203023121023-2033001331020211-1111030232031202"></a>

## https.coalescing_options.default_coalescing — default_coalescing / 130003230323 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- https.coalescing_options.default_coalescing

<a id="canonical-3320213130112310-2233113110331030-2133310210020131-1013021112113203-1133323000230020-3110201222201023-2010112213201332-3313233333200303"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default coalescing.

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

<a id="canonical-1010212003133223-3333101210211100-2110130222113020-0332012030303313-0031121022002333-3321131201030303-0310102202102121-0211212122002211"></a>

## Direct properties — default_coalescing / 130003230323 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0122003031113321-3223200122123000-1323210013023011-3101012132312133-1201113311102202-3100130213020221-3322302322102031-0030113110120120"></a>

## Next pages — default_coalescing / 130003230323 / 4

- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2310100133023231-2213123132321011-2120323001032322-1320201323220323-2302223031032123-2022333320330330-3030020322303211-2010112320230212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1033210332230112-0232332330220331-2011133000022312-2300130103203003-2231112200023033-2312120133220320-0123202002133231-3132200201110322"></a>

## https.coalescing_options.strict_coalescing — strict_coalescing / 311010313013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- https.coalescing_options.strict_coalescing

<a id="canonical-1001220333201210-2133213300202022-0101311012010211-0032232201131330-3320021033333111-1002111013011331-2122302023133230-2313020200302102"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict coalescing.

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

<a id="canonical-3313212101311113-3010003212313101-2021202002310322-0322301332132022-0022100130310330-0323001303301201-1113131201101202-3313013222332311"></a>

## Direct properties — strict_coalescing / 311010313013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0113021110103102-3133232212231221-2001002103303303-3221121100323031-2020133011210331-1113303210002212-0100312200203213-0123002221312022"></a>

## Next pages — strict_coalescing / 311010313013 / 4

- [https.coalescing_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-1001220203312330-3031013021133302-3130013202123010-1232020313310122-0213332300123110-3132233211103221-0220011120321111-3220032023002332)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1011220220213223-0023022311131111-3321100321212103-2302012100332100-3310212111103331-2030012330210201-1100032222211300-2333103131312233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2013022232120202-0031330111303020-2111023211233030-1103311202212111-0301020203311202-2001003201000201-1110032311322010-2313322012312133"></a>

## https.default_header — default_header / 220011131203 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.default_header

<a id="canonical-0222103132323210-3210011212101022-2002102233121002-2122312321202133-3123212120210232-0221212103021010-1200311301000011-2122213002323313"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default header.

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

<a id="canonical-2001110201002122-3231012210012012-2221103131031101-2200002222032003-2211321322133322-1001303123332302-3330031211332302-3302201022202200"></a>

## Direct properties — default_header / 220011131203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211120012013220-3121310111112312-3003120132301331-0303131132303011-0223002031011122-1323100322222103-0113331223010213-0312200332002300"></a>

## Next pages — default_header / 220011131203 / 4

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1123003233230101-3321333100231312-1212220312212102-3022231131212132-1132223102003323-3030022222021003-1030200021021213-1030221101223101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112122213033110-0332200221213022-3311230332101210-1000102000312223-2231310030010223-2332312103210001-0301112030301010-1022120313132133"></a>

## https.default_loadbalancer — default_loadbalancer / 020001000230 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.default_loadbalancer

<a id="canonical-2331031330211323-1202303110133100-2223313211012211-1313312200332113-1131210121120303-3332101000222033-2000131211313320-3201101213313133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default loadbalancer.

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

<a id="canonical-2002221220313301-3013220021232032-0103000202121130-2100320311101030-2021333312102320-0300111033011101-0332130023220313-3022121213111302"></a>

## Direct properties — default_loadbalancer / 020001000230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123300321201112-1300023031100330-0301312101232133-1122300101301230-3100101323220033-1120101301223313-0110021300230010-2230310331232110"></a>

## Next pages — default_loadbalancer / 020001000230 / 4

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1313230312203121-0221130100202030-2130030031311110-0000203003102103-0330223030021321-3212202011001322-2323331303122122-3302211302001001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102211130002122-0122012330223233-1203131210113222-0032313020031200-3112131111323303-3101302320202202-2022010322031202-2101001333130331"></a>

## https.disable_path_normalize — disable_path_normalize / 302322113311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.disable_path_normalize

<a id="canonical-0012123310211301-2122213003032100-3212100223103311-1032223102013203-3311123320222111-3031301021130103-2303203020001233-3221333103301212"></a>

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

<a id="canonical-0033123220201001-3332133100231023-2003300222321020-0210312213330313-0303120121011110-1200302300300210-2212021011001322-3103212103013213"></a>

## Direct properties — disable_path_normalize / 302322113311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0201213030301301-1133003120211013-0013002111223232-0012301202321231-2122030212300103-2111032022332321-0313231122222211-1131122230312112"></a>

## Next pages — disable_path_normalize / 302322113311 / 4

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3211010213301321-0310002022003033-2320121222112212-3322033031120213-3201231212313320-2233231203130122-3211232302302121-2012120022302233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321111311222023-3211121220302310-1100211003330031-3231211201330300-1313201330101131-3123333111311303-3211020210322023-2020311332131003"></a>

## https.enable_path_normalize — enable_path_normalize / 003031103011 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.enable_path_normalize

<a id="canonical-1003023302220222-2302331233003301-2113022230130101-3221320232300130-3120320301100100-0133330010321321-0333030122200130-2032020213320331"></a>

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

<a id="canonical-2321023320233012-0110213330301112-3033022322001321-1220203322102301-1221023000123003-0311201103021223-2212221021023103-1133122330122223"></a>

## Direct properties — enable_path_normalize / 003031103011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323312201002010-0212202300130320-3010120203223321-2331103230203010-2312030102032011-2101131303112211-2332221020011323-2101311022011233"></a>

## Next pages — enable_path_normalize / 003031103011 / 4

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1231203331312002-1212120301222301-0030312113033121-1122303113212113-2113131033312022-3212200231333000-2301222203010320-3133023112133121"></a>

## https.http_protocol_options — http_protocol_options / 330231301013 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.http_protocol_options

<a id="canonical-0331333210213110-1000010230021220-2202131230100023-2311311131111133-0332120100223232-2232233202201301-0032303213122333-2313333310030013"></a>

Type: `"single"`. Computed.

HTTP protocol configuration OPTIONS for downstream connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-http_protocol_choice": "[\"http_protocol_enable_v1_only\",\"http_protocol_enable_v1_v2\",\"http_protocol_enable_v2_only\"]"
}
```

<a id="canonical-1333031333201022-3222100030213211-0331133021033320-3012301221033200-0023122033210331-2031231210113133-3333123210103003-0001222333021032"></a>

## Direct properties — http_protocol_options / 330231301013 / 3

- [http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322): complete subsection reference.

- [http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-018.md#canonical-3123333211101233-2011211330112003-1122020031211202-0332221113321002-0200002300022131-3003033110131232-0230101212213231-1221232312332022): complete subsection reference.

- [http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-1130003130113323-3323300300133313-1002312300021013-1023122000132201-0231101201200133-2331101003201001-3220011220132332-2211322332112133): complete subsection reference.

<a id="canonical-0101131013122322-0301131203013220-1200121322312023-2302132320203323-1010110330102223-0013323131000213-2223121201200210-0133020122010023"></a>

## Next pages — http_protocol_options / 330231301013 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_v2](data-sources--http_loadbalancer--reference--group-018.md#canonical-3123333211101233-2011211330112003-1122020031211202-0332221113321002-0200002300022131-3003033110131232-0230101212213231-1221232312332022)
- [https.http_protocol_options.http_protocol_enable_v2_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-1130003130113323-3323300300133313-1002312300021013-1023122000132201-0231101201200133-2331101003201001-3220011220132332-2211322332112133)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021123301203210-1033100131200121-2310203130002121-1012203211131101-3231212322301111-3110120111023301-2030301201021021-0333112011232020"></a>

## https.http_protocol_options.http_protocol_enable_v1_only — http_protocol_enable_v1_only / 311030313131 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- https.http_protocol_options.http_protocol_enable_v1_only

<a id="canonical-0021131313021203-1020002021110101-0112313000001311-0000233120133200-3331210103012213-1111013323130311-2301310230230122-1132211312003103"></a>

Type: `"single"`. Computed.

HTTP/1.1 Protocol OPTIONS for downstream connections.

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

<a id="canonical-0220201001312231-0230022331220003-3122310312010011-2123030131102033-0030103332322302-0221322013100110-0201310120200121-3222201322301102"></a>

## Direct properties — http_protocol_enable_v1_only / 311030313131 / 3

- [header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030): complete subsection reference.

<a id="canonical-1001002133323102-2123030320010200-3312300311212310-1131322000111322-1313201222313030-2002112121020102-2233211111133330-3122133221100302"></a>

## Next pages — http_protocol_enable_v1_only / 311030313131 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113121113100030-2031222201311332-3103000231331123-0202221013131133-0031133033030110-0113030213031313-1303123221033333-2231103102300003"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation — header_transformation / 303320131303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation

<a id="canonical-0011103212003100-3122001323220110-0223303311021233-3131203013311210-0301031231310020-0032221130311131-3101300132321221-0100332230020232"></a>

Type: `"single"`. Computed.

Header Transformation OPTIONS for HTTP/1.1 request/response headers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-header_transformation_choice": "[\"default_header_transformation\",\"preserve_case_header_transformation\",\"proper_case_header_transformation\"]"
}
```

<a id="canonical-0002313123320031-2133220220120033-0010320030102022-3222102001121212-2302230310110332-3000223123001322-2120313300113031-0221002021022032"></a>

## Direct properties — header_transformation / 303320131303 / 3

- [default_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0212231112231303-3110332202132222-2112001213020312-0022110013323210-3111223000101213-3333210101120311-3131102223033011-0110130330303311): complete subsection reference.

- [preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101031213333122-1122323003122010-1212011021122202-3132110001231312-1022210133303112-2212213133001130-1112133213030021-0120100133212023): complete subsection reference.

- [proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1131012011232031-3011220131031031-2223313233001213-0020010012103312-2021303331223301-3111133222023212-3203231323332010-0212230320313133): complete subsection reference.

<a id="canonical-3211123330310123-1032021312331332-0203223312112132-1333110213023210-2033303303201001-3330123313321331-0333021332231211-2322210303310332"></a>

## Next pages — header_transformation / 303320131303 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0212231112231303-3110332202132222-2112001213020312-0022110013323210-3111223000101213-3333210101120311-3131102223033011-0110130330303311)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101031213333122-1122323003122010-1212011021122202-3132110001231312-1022210133303112-2212213133001130-1112133213030021-0120100133212023)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-1131012011232031-3011220131031031-2223313233001213-0020010012103312-2021303331223301-3111133222023212-3203231323332010-0212230320313133)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0212231112231303-3110332202132222-2112001213020312-0022110013323210-3111223000101213-3333210101120311-3131102223033011-0110130330303311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3220210031032120-3013230012002020-1130203003102122-0221303331030003-1100113300120101-1210133001022131-1032321332022011-0103311021303001"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation — default_header_transformation / 221310221112 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.default_header_transformation

<a id="canonical-0223213122123031-1203331030133110-2033000132112012-1221333130201233-0123101221000000-1110312332122223-1220101121021311-2022331122122022"></a>

Type: `["object", {}]`. Computed.

Use the platform's current default HTTP header transformation behavior.

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

<a id="canonical-0133010311023110-1200223302320223-3323012023323101-0012000201112310-2031220230233332-0102322230202333-1333322220212000-3112112202100222"></a>

## Direct properties — default_header_transformation / 221310221112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333323212020203-2213330021303022-2230021120233333-1200322200332300-2032232011221332-2021300221003321-0321020211331333-1103233320310313"></a>

## Next pages — default_header_transformation / 221310221112 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0101031213333122-1122323003122010-1212011021122202-3132110001231312-1022210133303112-2212213133001130-1112133213030021-0120100133212023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203121111302310-0203101333201213-1221110032121301-3331023022322032-1311211122331301-1203330201311313-3013232001132002-0123032303300003"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation — preserve_case_header_transformation / 010302010113 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.preserve_case_header_transformation

<a id="canonical-0122002100132211-1000033221323131-2130201033330311-3121302110030131-0033003312013310-1033332301000031-0112110312201323-1213122321322311"></a>

Type: `["object", {}]`. Computed.

Preserve HTTP header-name case when upstream case must remain unchanged.

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

<a id="canonical-0130023222322210-0223122012312123-2211302021231333-0120301223111323-1311002232130022-0213112201333110-1113210202203212-1220123113002212"></a>

## Direct properties — preserve_case_header_transformation / 010302010113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2110203130301321-0112322203013221-2303003023320330-1220021201233210-3212102302323101-1330322121231011-3201121313213222-3230301201113210"></a>

## Next pages — preserve_case_header_transformation / 010302010113 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1131012011232031-3011220131031031-2223313233001213-0020010012103312-2021303331223301-3111133222023212-3203231323332010-0212230320313133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3013002220302011-1113022122303001-3301221013313202-0012330201003203-3221031211101220-1311131323333000-0212323112213303-1112202221203312"></a>

## https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation — proper_case_header_transformation / 023110013121 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [https.http_protocol_options.http_protocol_enable_v1_only](data-sources--http_loadbalancer--reference--group-018.md#canonical-2131231302230011-3310322232331211-2030001330101211-0203321332010113-1210021203131212-2233023230202032-0200233320020102-0220312013203322)
- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- https.http_protocol_options.http_protocol_enable_v1_only.header_transformation.proper_case_header_transformation

<a id="canonical-1132120303330223-2213302003231303-0102333312130121-3220111023230232-3123222322332333-1110033022032010-0222010002131203-1131022232012003"></a>

Type: `["object", {}]`. Computed.

Transform HTTP header names to proper case when explicit transformation is required.

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

<a id="canonical-1231001230110203-1202331002302103-2303030332113120-3013121133332231-0203331013010002-3110302220312310-0330220231022020-3300322300232110"></a>

## Direct properties — proper_case_header_transformation / 023110013121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120111022001331-1012310012333001-1111333030030031-1323202330131110-1011002113323302-0002233002300131-2221123223102131-0213110312033100"></a>

## Next pages — proper_case_header_transformation / 023110013121 / 4

- [https.http_protocol_options.http_protocol_enable_v1_only.header_transformation](data-sources--http_loadbalancer--reference--group-018.md#canonical-0310021231311132-0310332010021331-3020123330001232-2022001313200222-2220233211302110-3230100022101233-0212202212133201-0023221100111030)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3123333211101233-2011211330112003-1122020031211202-0332221113321002-0200002300022131-3003033110131232-0230101212213231-1221232312332022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333132123310213-3020131021203313-3030212301320232-0101210002112101-1020121233333332-0320230332132201-2330200100131113-0110300102032211"></a>

## https.http_protocol_options.http_protocol_enable_v1_v2 — http_protocol_enable_v1_v2 / 210021122202 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- https.http_protocol_options.http_protocol_enable_v1_v2

<a id="canonical-2022311233313011-0323000202123222-2121021202001200-2000223030012133-2223031011323030-3023302233112231-1010333221300313-0100223032130103"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v1 v2.

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

<a id="canonical-1021132231101303-2001020030200023-2131333021122321-1133200210223003-1220112301102100-0333330213010012-0200003233122322-1231031211311102"></a>

## Direct properties — http_protocol_enable_v1_v2 / 210021122202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130331010312112-3002113223200203-0102333032233331-2232301000201321-1320022000320201-3112313133103320-2231020023122320-3311200322021010"></a>

## Next pages — http_protocol_enable_v1_v2 / 210021122202 / 4

- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1130003130113323-3323300300133313-1002312300021013-1023122000132201-0231101201200133-2331101003201001-3220011220132332-2211322332112133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1303313202230231-3212303112133201-0233023322032211-1300230223330100-1132300200221221-1201211320112303-0313101031201312-2132230331210233"></a>

## https.http_protocol_options.http_protocol_enable_v2_only — http_protocol_enable_v2_only / 300010103221 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- https.http_protocol_options.http_protocol_enable_v2_only

<a id="canonical-2202101213113122-1101012000231012-2000033211033011-3201311103302232-3012220321002333-1113221002331313-3222231110332013-3231000112113332"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for http protocol enable v2 only.

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

<a id="canonical-3033300000123303-2133232320012303-1233103121313220-1333121113111021-2031300010210110-0322211131113020-3101231011323302-3111232202032323"></a>

## Direct properties — http_protocol_enable_v2_only / 300010103221 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212332302203231-2232313030001221-3303300110110220-2132230333310310-0000232220000212-3122220322001331-1213120300312020-0230203201102030"></a>

## Next pages — http_protocol_enable_v2_only / 300010103221 / 4

- [https.http_protocol_options](data-sources--http_loadbalancer--reference--group-018.md#canonical-0003020203103000-2333100023323311-1111213021211113-2303201112013303-2123231030312110-3202013021320132-3330331220213021-1110230322200021)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1101330312012103-0233323210311221-1033012102103323-0220330330101233-1110011203231222-2231210133200221-1022100012023223-3112010202021011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003113330013331-0311310220011332-0000313232223010-3022103311313301-2201332021311333-3311011103113202-1301123103230203-2221110002031132"></a>

## https.non_default_loadbalancer — non_default_loadbalancer / 223321111023 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.non_default_loadbalancer

<a id="canonical-3022332322133201-3122111332221310-2303311112031200-3012332122103121-3330302003203013-0021310103100123-2330113103011121-3022110211012320"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for non default loadbalancer.

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

<a id="canonical-0133111023211213-0230230331111210-1033012102201201-2212132333201120-1321202330101101-1331233231030212-1012302013130001-1131312102120001"></a>

## Direct properties — non_default_loadbalancer / 223321111023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303102223323030-0130123232220001-2323121103022213-2012311330030103-2121233210031323-1311201030031302-2033201001320010-0032102333321031"></a>

## Next pages — non_default_loadbalancer / 223321111023 / 4

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3210232200212003-0002112320100202-1003013010223122-1001223222303210-2322123332210310-3322112011113221-2311013202301022-0201323130200031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3110313031200120-3321113202322021-0201331201333021-3003213300022011-1003122231030212-3103020020231322-0102101100312033-0112221230211222"></a>

## https.pass_through — pass_through / 220312322311 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.pass_through

<a id="canonical-1221302222302211-0011201000033021-0320010103220203-2110031312322200-3113222122102102-1220100030232100-0213212203210223-0230323001200133"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for pass through.

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

<a id="canonical-3021301323120211-3130112222003101-3212222302203210-1131022301101220-0112123333303020-1210302232310122-3132310230020211-0332321013312113"></a>

## Direct properties — pass_through / 220312322311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2003310232103110-0222032131211322-3023203202313201-1300311113132330-3120101323222020-3332222122303001-0202212300210301-2121013010000031"></a>

## Next pages — pass_through / 220312322311 / 4

- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323012233332233-2300031101333001-2223210231012213-3101313023132323-1032000032120021-0312212021303021-3313201303330202-0311021001133012"></a>

## https.tls_cert_params — tls_cert_params / 101001110020 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- https.tls_cert_params

<a id="canonical-1131030322122230-2321302021230200-1202302123312222-1123302122220220-0020220323233321-2013203301201210-1320212200231112-2023130102322202"></a>

Type: `"single"`. Computed.

Configuration parameter for tls cert params.

Upstream description:

Select TLS Parameters and Certificates.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

<a id="canonical-2313021300012230-0033133212320331-2302331323203030-0011012111222212-3310212320123020-2012320010102132-3310230320332213-2220112220332233"></a>

## Direct properties — tls_cert_params / 101001110020 / 3

- [certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-2312203201012133-1132133332020110-0112332133100322-2212200200230021-0203033110100302-1301102101101330-3200320103310333-1123222313100333): complete subsection reference.

- [no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223133102312020-3233203020002133-2303010022213201-3210211100122220-3221130303102112-2303220003302120-2023322301131321-1100122120021020): complete subsection reference.

- [tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001): complete subsection reference.

- [use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100): complete subsection reference.

<a id="canonical-1310101120312302-3200233201123133-3001313110131030-3003303200011110-1021121230333303-1220331212132113-0113020113102200-2301130133100323"></a>

## Next pages — tls_cert_params / 101001110020 / 4

- [https.tls_cert_params.certificates](data-sources--http_loadbalancer--reference--group-018.md#canonical-2312203201012133-1132133332020110-0112332133100322-2212200200230021-0203033110100302-1301102101101330-3200320103310333-1123222313100333)
- [https.tls_cert_params.no_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0223133102312020-3233203020002133-2303010022213201-3210211100122220-3221130303102112-2303220003302120-2023322301131321-1100122120021020)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2312203201012133-1132133332020110-0112332133100322-2212200200230021-0203033110100302-1301102101101330-3200320103310333-1123222313100333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133030232220320-2202302311131022-1012221333301120-0123222322311213-3102333030331020-3332201023103300-0301033033101330-0220121301303021"></a>

## https.tls_cert_params.certificates — certificates / 030122132213 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.certificates

<a id="canonical-1120102101120123-2320123033022110-1023031020010332-3023332003031301-1303123310100213-2323131320201120-0023033312303221-2232233032203211"></a>

Type: `"list"`. Computed.

Select one or more certificates with any domain names.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3022330200232303-1331120011123000-3030333230011200-2110022113231322-2210200212023321-3203210101030100-1120313133130222-3201103230032123"></a>

## Direct properties — certificates / 030122132213 / 3

<a id="canonical-3231230013320110-0222020222103123-2310232020022100-2223022310330223-3220031103020012-0131313233223310-0200003033022220-3013303210132213"></a>

<a id="canonical-1203233112120001-1100213233231120-0332021103303002-0110031010200102-0133001123221313-0202211123233331-2300030203132310-0132100033323111"></a>

## name property — certificates / 030122132213 / 4

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

<a id="canonical-3032321302131012-0202112021130221-0001112013212331-0311221212213313-1220300121010033-0031123320210321-0100003230223213-1322333031323203"></a>

<a id="canonical-1331020121121220-1121232112203023-2200200123301020-0221000132002031-0303320300010302-3110320302003200-1202003312210333-0212202212233011"></a>

## namespace property — certificates / 030122132213 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-1003000122023113-1322131222101211-1021321013232300-1310332101213003-2102320012232210-3121123031212112-1002011002232201-1021110222013102"></a>

<a id="canonical-1310222323310300-0131212211303323-0033303203021213-2002123020333333-3010202113222103-1001300032100302-2223122203221203-3120202120312020"></a>

## tenant property — certificates / 030122132213 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2020321203102123-2221220012132111-3012021232202002-1131220312101320-1110232311331230-0131311222233231-0220310230010012-0322213333023330"></a>

## Next pages — certificates / 030122132213 / 7

- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0223133102312020-3233203020002133-2303010022213201-3210211100122220-3221130303102112-2303220003302120-2023322301131321-1100122120021020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000103312300323-1032132300022030-1203121200203103-3023303003212013-2111322122113300-2330101233200103-2212201203331212-1131223133123130"></a>

## https.tls_cert_params.no_mtls — no_mtls / 033013120032 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.no_mtls

<a id="canonical-1212130101301300-1200311122200223-1102001223031000-1211100303323103-3300210313002102-3301121131021330-2222100133023331-0231210332011113"></a>

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

<a id="canonical-1210120031101233-1132021220312330-2301112100131222-3203323031310220-0330122010000311-3132303213001032-3030000321210233-1210230232221301"></a>

## Direct properties — no_mtls / 033013120032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230202120223303-2001002123301122-0113020022311121-1122323201333032-3313130123030121-1311001011033011-1311211220333131-0300323310322300"></a>

## Next pages — no_mtls / 033013120032 / 4

- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0300222211131012-2300303012133011-1331233103012313-1300311301322020-2133313011333330-0130203133322321-3020321111302212-0020303021302103"></a>

## https.tls_cert_params.tls_config — tls_config / 131003002232 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.tls_config

<a id="canonical-3132332323332323-1330201111102230-2303222301010003-0110202022001022-3211310012032202-1022001320233031-1230021221313333-1022112121101110"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-2311111022012321-2220003130032002-3321103203333032-1323331120013323-3000103201313203-2311212323332303-0033122333033313-2102023013333102"></a>

## Direct properties — tls_config / 131003002232 / 3

- [custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-2320123123023013-2210220122302000-0310020301201310-3312122223321030-2321012331122220-1332231220120032-2133031212023320-0230312002103003): complete subsection reference.

- [default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-1312123121000103-2131020023213321-0200320131030212-1301323003312120-3201030013320331-0321113330013220-1220223102010122-0033300221233103): complete subsection reference.

- [low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-0002220110031322-2000022330002023-1210033032322201-2222102300031303-3102020301103131-2121013132310021-0212213012103132-2201330202321122): complete subsection reference.

- [medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-3203032113231011-0312131120020200-2200310022201122-0130021221021222-3312322001011302-1231211323221220-3320013323033222-1221033013130332): complete subsection reference.

<a id="canonical-2330230212003321-1201031320110132-3121032310102313-3032132101321333-2010223331300333-2330101000030232-3110102121312033-1311332323311112"></a>

## Next pages — tls_config / 131003002232 / 4

- [https.tls_cert_params.tls_config.custom_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-2320123123023013-2210220122302000-0310020301201310-3312122223321030-2321012331122220-1332231220120032-2133031212023320-0230312002103003)
- [https.tls_cert_params.tls_config.default_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-1312123121000103-2131020023213321-0200320131030212-1301323003312120-3201030013320331-0321113330013220-1220223102010122-0033300221233103)
- [https.tls_cert_params.tls_config.low_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-0002220110031322-2000022330002023-1210033032322201-2222102300031303-3102020301103131-2121013132310021-0212213012103132-2201330202321122)
- [https.tls_cert_params.tls_config.medium_security](data-sources--http_loadbalancer--reference--group-018.md#canonical-3203032113231011-0312131120020200-2200310022201122-0130021221021222-3312322001011302-1231211323221220-3320013323033222-1221033013130332)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2320123123023013-2210220122302000-0310020301201310-3312122223321030-2321012331122220-1332231220120032-2133031212023320-0230312002103003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123301210212102-3321310130111110-2320301101113030-0110221133033131-0212120300133200-0111021030103123-3030300302221200-1010322100333221"></a>

## https.tls_cert_params.tls_config.custom_security — custom_security / 000303301212 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.custom_security

<a id="canonical-2132132210332122-1303031131313203-3120222101101220-3101001213332003-1113301010310222-2023211301020210-1101303131233133-2202101220133110"></a>

Type: `"single"`. Computed.

Defines TLS protocol config including min/max versions and allowed ciphers.

Upstream description:

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-2332221013030221-2103132222220330-0002233201111210-3323021020231013-0231020023100203-2122300330112031-1031331312230000-1020313111311130"></a>

## Direct properties — custom_security / 000303301212 / 3

<a id="canonical-3111000330222101-1000023332013311-0232331332131230-2122331232111200-1020101230231002-1232230122030301-2320313100121320-1002110212003223"></a>

<a id="canonical-0130313110323103-3213220323132230-2230221011231312-2003100130333120-1330122033130302-2310030231110130-2220302022200133-1332323210111113"></a>

## cipher_suites property — custom_security / 000303301212 / 4

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1222131331031232-3010221200333230-0001302300233313-3020230203003100-1020300312120120-2123122313110223-3212230311200120-3002210310321221"></a>

<a id="canonical-3212230131301120-0211200202111312-3032200323030101-3023301223012022-2113030023221011-1202211321031001-1022020222022310-1301200233123213"></a>

## max_version property — custom_security / 000303301212 / 5

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2221113133121120-1020032000103101-2033232003331220-1320201031211001-1121000320121300-2013232120123322-1122212322310233-1201110300023032"></a>

<a id="canonical-3313112210310002-3312203103002223-3202323100123300-2023001213332231-0202210112103011-0203100201023200-2333103033233112-0002223012010103"></a>

## min_version property — custom_security / 000303301212 / 6

Type: `"string"`. Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Upstream description:

TlsProtocol is enumeration of supported TLS versions

F5 Distributed Cloud will choose the optimal TLS version.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2002131023211121-0322220303033313-1113202120203210-1201031002233203-0133032230231031-0032010213130303-1330200310201112-1300030302130222"></a>

## Next pages — custom_security / 000303301212 / 7

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1312123121000103-2131020023213321-0200320131030212-1301323003312120-3201030013320331-0321113330013220-1220223102010122-0033300221233103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0221333121201133-3111201121010333-1130102023203331-1203331211111010-3330031231000230-2231112321103221-1230212110220231-1300221222010001"></a>

## https.tls_cert_params.tls_config.default_security — default_security / 312101023321 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.default_security

<a id="canonical-3322200322120202-2200000121000232-3311322221132002-2121313103112033-0301332000200222-2103013300201000-3100000303211011-1020330032231131"></a>

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

<a id="canonical-0231213221012033-3022120202303021-0020220331020201-1133101201122312-2003013301003320-1231132221232030-2310300021212113-0110121110310033"></a>

## Direct properties — default_security / 312101023321 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0003100222302003-2023301121023231-0102122202120321-0110131013313310-1023031030031132-0110301031133332-3322130003221330-3103012232213030"></a>

## Next pages — default_security / 312101023321 / 4

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0002220110031322-2000022330002023-1210033032322201-2222102300031303-3102020301103131-2121013132310021-0212213012103132-2201330202321122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001101011122112-0221112201310322-1322203100000311-2122133111111132-1322222103021100-2123311112111321-0232302113011000-1133122023100210"></a>

## https.tls_cert_params.tls_config.low_security — low_security / 311131131302 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.low_security

<a id="canonical-0111210212313120-2012301312113313-3312310033200313-3200010002302220-3311303331131321-0013201220321203-1203103020221031-0223032022030323"></a>

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

<a id="canonical-1212110222120230-3201003123130203-3122132221200301-1303303011323021-1111210200210323-3101223031330133-2002112332332111-3002323223023320"></a>

## Direct properties — low_security / 311131131302 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1112201233301132-2010131222203220-1222302033012103-2320201321332210-3213222330132320-1000312312221121-1201233133220331-1221131303310313"></a>

## Next pages — low_security / 311131131302 / 4

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3203032113231011-0312131120020200-2200310022201122-0130021221021222-3312322001011302-1231211323221220-3320013323033222-1221033013130332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003120203233032-2202330221221210-2332032223003001-3303321312013022-0320313212300023-1131313130001323-1022132303102230-1332003230200103"></a>

## https.tls_cert_params.tls_config.medium_security — medium_security / 231133211010 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- https.tls_cert_params.tls_config.medium_security

<a id="canonical-0333033101123201-0020322121030321-3332133321200122-3200322112101032-1210322031220320-0223221201202321-1213001202023110-3212110202300231"></a>

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

<a id="canonical-3120120033220212-0120321011000010-2232222022120122-1203231022311101-2021232233210322-0003030003021003-0110301211311130-1332122201102102"></a>

## Direct properties — medium_security / 231133211010 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1213103201320311-3321210000113011-1123222233210121-3003200223333231-0032133301021331-1121213021300213-1333221302120122-3011200332212100"></a>

## Next pages — medium_security / 231133211010 / 4

- [https.tls_cert_params.tls_config](data-sources--http_loadbalancer--reference--group-018.md#canonical-0330213002100010-3210123313013210-0111102220002103-1222200012322122-1323203130221323-1301013230233031-3110101122322333-0033031012123001)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3121100212103131-0022233101312033-3110312130103032-2130332020112121-1132203130011312-1113103221330122-1113332021322103-3110230031301123"></a>

## https.tls_cert_params.use_mtls — use_mtls / 032102032320 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- https.tls_cert_params.use_mtls

<a id="canonical-1001033312230212-1211032321133220-3111300223331030-0212320310232012-0130220033221203-2211122032330020-2201201310310132-2210222321003323"></a>

Type: `"single"`. Computed.

Validation context for downstream client TLS connections.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-crl_choice": "[\"crl\",\"no_crl\"]",
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]",
  "x-ves-oneof-field-xfcc_header": "[\"xfcc_disabled\",\"xfcc_options\"]"
}
```

<a id="canonical-2220202300222111-1300013211223212-3122312221210233-1210331332022201-3233200231133333-1022300232031223-3011120120330202-2322331011231112"></a>

## Direct properties — use_mtls / 032102032320 / 3

<a id="canonical-1300211232332133-1310201331032020-0121311300011030-2220311321300030-0013120202110000-2301310102121301-0033033121332313-0011223112122111"></a>

<a id="canonical-3320312203311313-1000133220212202-0203333222033032-3123100133003332-0202300312103023-3000123312023122-2200130330210200-3333302211011010"></a>

## client_certificate_optional property — use_mtls / 032102032320 / 4

Type: `"bool"`. Computed.

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated.

Upstream description:

Client certificate is optional. If the client has provided a certificate, the load balancer will
verify it. If certification verification fails, the connection will be terminated. If the client
does not provide a certificate, the connection will be accepted.

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

- [crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-2132103301102021-1132103122101230-3123022210011220-0320030213232331-2311321203000203-0221302121132220-1123231211020323-1032331010322322): complete subsection reference.

- [no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-3213132300132122-1300233220030300-2200001200000201-3223021231022201-1031301301232330-0011113121023200-3202230312323002-0300101313321032): complete subsection reference.

- [trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-1221321023303223-0313120023012131-3100323220013221-2112001021000211-2302311022133303-2200303100331113-3200311001223200-1323113213332130): complete subsection reference.

<a id="canonical-3320211123102210-0111332210001123-0231300303110011-2121031020201210-2013110210213002-2223222113022022-1231030033333330-1131230100130301"></a>

<a id="canonical-0020110020001011-0013032122201112-0000120310212122-3020231301321112-1020102020003231-2301300310322332-1301203312310013-0021133003121102"></a>

## trusted_ca_url property — use_mtls / 032102032320 / 5

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Load Balancer.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

- [xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101113332133200-2011200000230031-0032201320100301-2333000103132121-0121221113003010-3220211212032320-1023031111012212-1113123323123003): complete subsection reference.

- [xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3013120300333313-2132201100310030-1310220122232030-3302131210133100-0133030331022131-2300210103021313-2322311123331223-0123230011202223): complete subsection reference.

<a id="canonical-3031020211103323-0311122123333211-2020200321311233-1133111122031233-1221031230032021-3321010203101121-2322102031230220-2020321312022322"></a>

## Next pages — use_mtls / 032102032320 / 6

- [https.tls_cert_params.use_mtls.crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-2132103301102021-1132103122101230-3123022210011220-0320030213232331-2311321203000203-0221302121132220-1123231211020323-1032331010322322)
- [https.tls_cert_params.use_mtls.no_crl](data-sources--http_loadbalancer--reference--group-018.md#canonical-3213132300132122-1300233220030300-2200001200000201-3223021231022201-1031301301232330-0011113121023200-3202230312323002-0300101313321032)
- [https.tls_cert_params.use_mtls.trusted_ca](data-sources--http_loadbalancer--reference--group-018.md#canonical-1221321023303223-0313120023012131-3100323220013221-2112001021000211-2302311022133303-2200303100331113-3200311001223200-1323113213332130)
- [https.tls_cert_params.use_mtls.xfcc_disabled](data-sources--http_loadbalancer--reference--group-018.md#canonical-0101113332133200-2011200000230031-0032201320100301-2333000103132121-0121221113003010-3220211212032320-1023031111012212-1113123323123003)
- [https.tls_cert_params.use_mtls.xfcc_options](data-sources--http_loadbalancer--reference--group-019.md#canonical-3013120300333313-2132201100310030-1310220122232030-3302131210133100-0133030331022131-2300210103021313-2322311123331223-0123230011202223)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-2132103301102021-1132103122101230-3123022210011220-0320030213232331-2311321203000203-0221302121132220-1123231211020323-1032331010322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212313103021310-0332233220021010-1230311300033200-0012110320022230-2201011331212122-0221131003021303-1112002203332111-2032222321213313"></a>

## https.tls_cert_params.use_mtls.crl — crl / 011231321100 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.crl

<a id="canonical-1210222203113202-0333012120220033-3030323113120011-1133111213321133-2001332000210130-3220311012203101-3100201010200121-0020101133001300"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

<a id="canonical-0210122031221012-1133001010213101-3010110301002232-2232223220211230-2021122333113000-1223300002133313-3200120121301313-3022100020310311"></a>

## Direct properties — crl / 011231321100 / 3

<a id="canonical-2122203201320103-1211231210303023-3313330010321213-0331002032300330-1101321222331320-2000210021102200-0211133003131131-0001200302001133"></a>

<a id="canonical-1101033020301022-1323030021132012-3330023021330302-2223032213010101-2032320102022011-0233020012130320-0113202121310221-1301321232102112"></a>

## name property — crl / 011231321100 / 4

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

<a id="canonical-2303130111210311-1113313301301223-0310130231201121-1120311102131120-1231313232023102-0301212102032333-0031320130110030-1302322313331313"></a>

<a id="canonical-3120231212003321-2110021202102002-2032232321110013-1313203222310001-0312121323003313-3000030311220100-3122233033322023-1102130123303111"></a>

## namespace property — crl / 011231321100 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3332333011331301-2203022320030220-2231132310013103-3023112023122010-3323111233213030-0121303200310220-3323122033000023-1131231203033012"></a>

<a id="canonical-1310231022311303-3223222201022102-0230321003120013-3330030111101213-0010130120102031-0320011112200121-0320202212233201-2313023332002000"></a>

## tenant property — crl / 011231321100 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-2201010203232331-2133102012111301-3303011232032121-1013223113303022-0021112030200000-2211020122211330-0001213233210300-2321030310220222"></a>

## Next pages — crl / 011231321100 / 7

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-3213132300132122-1300233220030300-2200001200000201-3223021231022201-1031301301232330-0011113121023200-3202230312323002-0300101313321032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3300113310333203-3201133323331302-0203020301023110-3220131110122223-3211310120012331-3113213012113132-0101333102110233-2130333220302002"></a>

## https.tls_cert_params.use_mtls.no_crl — no_crl / 212121331111 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.no_crl

<a id="canonical-1230323110030311-2121301101130301-0301130221013000-0212302022123022-1300202203001002-3002000211033201-2213131321121310-1120103211030031"></a>

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

<a id="canonical-3320211122013330-0021310313320131-3012012012123203-3001301033333130-1003213001230011-2112012033023303-0011202202213131-0101330021231032"></a>

## Direct properties — no_crl / 212121331111 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2121203213030221-3020031123333322-2023001303223012-3303003002011323-0221113332130202-3033320331332230-2032230320231023-1303321231221201"></a>

## Next pages — no_crl / 212121331111 / 4

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-1221321023303223-0313120023012131-3100323220013221-2112001021000211-2302311022133303-2200303100331113-3200311001223200-1323113213332130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0223213022203033-3123303003101000-0013021012300130-3032131123201233-0001322030201321-2033232112023003-1012222321000111-0132121221201321"></a>

## https.tls_cert_params.use_mtls.trusted_ca — trusted_ca / 231212133303 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.trusted_ca

<a id="canonical-0122021112200112-1111230123132222-0113022013033320-1232321113330213-2312210210023121-2212013113021130-0012232310021023-3111221112233201"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

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

<a id="canonical-3001203131301231-1332110222313031-3323131320202103-1200032331110010-2000121002222020-2132021303031112-1020321311210021-1231032221021020"></a>

## Direct properties — trusted_ca / 231212133303 / 3

<a id="canonical-0232301222302201-2131033302202020-2113011101002332-0101012022123321-3121220033213000-1303200130300110-3103002211223112-3320023210002312"></a>

<a id="canonical-0132231330203223-1222113032202013-3200122232211311-3010103000300322-2112000220013210-0311111013333322-2012323013121002-2120100130122230"></a>

## name property — trusted_ca / 231212133303 / 4

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

<a id="canonical-0101330022303022-2213311001331012-2322000333101230-1010120003000222-2031232111101200-2311332210102221-3123023122111133-0301233003020221"></a>

<a id="canonical-2310110010200203-0132310222222201-1203212132110010-3113013122030003-0313300222301010-2022010210000030-2111200322331201-3322021223333131"></a>

## namespace property — trusted_ca / 231212133303 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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

<a id="canonical-3223132220333130-3133220000311002-2332220013020020-1321302023020331-3223200231211223-2303222312023303-3201210331210210-2122321001122323"></a>

<a id="canonical-1232323011312201-0120012111123323-3323122332331022-0110222123003321-0213220311323111-0011130000122210-3221202221010230-2001110002322101"></a>

## tenant property — trusted_ca / 231212133303 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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

<a id="canonical-3333203203132111-2202100331223302-0011123111031001-0131312330003122-3100210020112122-2222212030323020-2130113132013111-1120113100330301"></a>

## Next pages — trusted_ca / 231212133303 / 7

- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)

<a id="canonical-0101113332133200-2011200000230031-0032201320100301-2333000103132121-0121221113003010-3220211212032320-1023031111012212-1113123323123003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101112102202113-2121230230301233-0211110213032203-1110223211201333-3020003011303030-2210332033330130-2110221123102032-0112202121030001"></a>

## https.tls_cert_params.use_mtls.xfcc_disabled — xfcc_disabled / 130130330223 / 2

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md#canonical-2321021211310331-0013021131330032-3102301033221021-2110211001101012-3233200230311321-0111111122130023-0121000122111210-2130013030322000)
- [Property reference](data-sources--http_loadbalancer--reference--group-001.md#canonical-1120010220221032-3212233022220221-3011110301103233-1232002110102021-0012103312230301-1322002232013201-2131001010211321-3011122013022320)
- [https](data-sources--http_loadbalancer--reference--group-018.md#canonical-0112223033331211-0131010222320312-0220012121111333-2323003223120200-1033012212133000-1022211023310312-1210322030333100-1313202312011301)
- [https.tls_cert_params](data-sources--http_loadbalancer--reference--group-018.md#canonical-0011121033233220-3311202112112001-0220333111322122-2213202322313002-3213103232311233-2200102030312213-0032120231111113-3201230203322302)
- [https.tls_cert_params.use_mtls](data-sources--http_loadbalancer--reference--group-018.md#canonical-0100211003133012-2022100102310023-1213133030310333-1112311330321323-1113030000201010-3012132103130231-3110303333323213-0302113200332100)
- https.tls_cert_params.use_mtls.xfcc_disabled

<a id="canonical-2322010031032131-3312013130331212-0102311232120110-2330300003210030-1020233133212212-0100102011213000-0032223113221113-0230012220113302"></a>

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

<a id="canonical-0332200020330312-1330333303222300-2133331333333013-3000311222213222-2330212332202112-2113132031101203-0222033130121012-1103322311213323"></a>

## Direct properties — xfcc_disabled / 130130330223 / 3

This is an empty object or choice marker. It has no direct properties.
