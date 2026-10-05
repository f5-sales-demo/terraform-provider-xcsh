---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-3310231212322113-0220122301031302-1222131010132201-1311130201302320-1203032330312020-0011330330120233-0233303031200001-1131323113120032"></a>

## origin_pool.use_tls.use_server_verification — use_server_verification / 333301201300 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.use_server_verification

<a id="canonical-2220130033132213-0011021310023330-1102010203300320-2123111000023020-3032130011203121-0102210310033012-2333032132133221-0111100323132103"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for use server verification.

Upstream description:

Upstream TLS Validation Context.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("trusted_ca",
    "trusted_ca_url")}
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
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
use_server_verification {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302200312001023-1210122322331133-3303211132030003-1102202111132113-0313222013323100-0113231000212312-2000022321230331-3033003322010013"></a>

## Direct properties — use_server_verification / 333301201300 / 3

- [trusted_ca](resources--cdn_loadbalancer--reference--group-013.md#canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300): complete subsection reference.

<a id="canonical-3012133003230110-3122032221312322-1201031021010123-1202223303033323-1012002131300311-2233220201012103-3232230202322233-0310300302333023"></a>

<a id="canonical-1330323310033201-2333300020220211-0212103101021301-3023111010320122-2011112332320202-0020311212100223-3213223321210121-1312100031120202"></a>

## trusted_ca_url property — use_server_verification / 333301201300 / 4

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Upstream description:

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 131072),
}
```

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

<a id="canonical-1332123013112001-0331021133020110-1112201113231311-0113232313323113-0020101003001320-1002012221310033-1001001303210101-0103001132010131"></a>

## Next pages — use_server_verification / 333301201300 / 5

- [origin_pool.use_tls.use_server_verification.trusted_ca](resources--cdn_loadbalancer--reference--group-013.md#canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3002213023021322-0133011031123120-1210313002112221-2202330030331233-2100222232331003-1023130312302130-0033300311121201-3021022123321300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210212023002231-2321210002112023-2100330303113023-2331032200323100-0321120120023213-2133303233002201-3012013000312031-0323110300303013"></a>

## origin_pool.use_tls.use_server_verification.trusted_ca — trusted_ca / 111220003221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023)
- origin_pool.use_tls.use_server_verification.trusted_ca

<a id="canonical-2303233022001300-2311022331230210-3111001002332233-3330113303203101-0213231122203020-1012332002020133-0112102000002220-1223010221311102"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2123300223303230-0322230021202002-3203031100200020-3022330102333203-0323221210101120-0021100213012310-1002101121301032-1031312120333321"></a>

## Direct properties — trusted_ca / 111220003221 / 3

<a id="canonical-0031000020110300-2132200001212233-2033132003332201-1021121313123231-3112130210331312-1310310203001301-1301011101311211-0022010203001200"></a>

<a id="canonical-2200311131221330-1101121333333201-1023020330001332-3221323002101333-0121303011300003-1132213121032223-3320031222023333-2113102323112231"></a>

## name property — trusted_ca / 111220003221 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2313133021321132-2223122331210320-3311110013020013-0323221300220122-2131100131021200-1131102002201022-1231212122233001-1112001333112301"></a>

<a id="canonical-1210333231230022-2322213023002020-1012231121101223-1102020231100322-2130233200033000-2031311033332311-2020002100312130-2323200332211012"></a>

## namespace property — trusted_ca / 111220003221 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2213330020231010-2112030003303333-2313001320031132-1102331220232123-0313002213032320-2122220101103210-2131303002311300-3112303311222012"></a>

<a id="canonical-3112200211202132-3120321302233013-2023103333133032-0013012123100321-3122100220010310-1120031113333021-0003303102101131-3300102001212313"></a>

## tenant property — trusted_ca / 111220003221 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2130220003310111-2120223313213011-2112033012321023-3110013303103121-0231212020130001-1123322200230222-1002000113330021-1021232322132032"></a>

## Next pages — trusted_ca / 111220003221 / 7

- [origin_pool.use_tls.use_server_verification](resources--cdn_loadbalancer--reference--group-012.md#canonical-0100201323013321-1212322220033323-3202120302123233-0233220232202213-1002311103202112-3103313101033013-1020001330210203-1011111032102023)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3132323021311131-3110113002100113-3203001123320003-0101112112132233-2030100310011022-2033133112113021-2220301213001231-2322101002133031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203100203312313-3010000333300330-0023331213031221-2002132130010300-1132201023312123-0232231321021313-3123211303331331-3333103022033103"></a>

## origin_pool.use_tls.volterra_trusted_ca — volterra_trusted_ca / 320130332110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [origin_pool](resources--cdn_loadbalancer--reference--group-012.md#canonical-1200032212323201-2220201020121210-0002002033031232-0001321032000033-1220312320030021-1321001312321022-3101031300210033-3121031303121312)
- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- origin_pool.use_tls.volterra_trusted_ca

<a id="canonical-2111132302100021-0110132102031000-1113022233110011-1303020322030321-1112330132011311-2033021022123320-2301131231223000-3322101003221033"></a>

Type: `["object", {}]`. Optional, Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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
volterra_trusted_ca = {}
```

<a id="canonical-2012222131032133-0301121112030100-0101220001222030-0113123303322033-0103023031311300-3031200021102002-1212331003003230-1220211232321022"></a>

## Direct properties — volterra_trusted_ca / 320130332110 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2311101021000113-1121011200322011-0020212231113111-2013131202231002-0022331121223312-1320302303113120-3113132312133113-2021210120331002"></a>

## Next pages — volterra_trusted_ca / 320130332110 / 4

- [origin_pool.use_tls](resources--cdn_loadbalancer--reference--group-012.md#canonical-0201031001123330-2330120010103001-2033222301230303-0111222320012213-3232310323000321-0021201303321110-2102223233331330-2020321102133200)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213113032333122-3322300230013333-0100130030330002-2110331310121031-1313031103202100-3200322103212102-1212210102321000-1030103020003313"></a>

## other_settings — other_settings / 023322023320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- other_settings

<a id="canonical-1222331321330011-0201120000010222-1203121020212220-2233310323010120-2321122200001333-2201233222312132-1302003321013013-1130233122001302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for other settings.

Upstream description:

Other Settings.

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
other_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113322112111200-0023211312003233-0030301333221330-0102233203003021-3312012022103022-2233032111232110-0312320220320322-1001323302220221"></a>

## Direct properties — other_settings / 023322023320 / 3

<a id="canonical-3001102222002010-2330013221301331-1301220323030223-3313000213011022-0330322332223111-2320332113112003-0021202120111201-2023032332032011"></a>

<a id="canonical-1131023131013201-2012001201002020-2231323101023223-0122320212013022-2030202220321230-2012223302332331-2022132113132320-2211202022201223"></a>

## add_location property — other_settings / 023322023320 / 4

Type: `"bool"`. Optional.

Add Location. X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt;
in responses.

Upstream description:

X-example: true Appends header x-F5 Distributed Cloud-location = &lt;RE-site-name&gt; in responses.

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

- [header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123): complete subsection reference.

- [logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130): complete subsection reference.

<a id="canonical-0121323010203231-0320122100102213-1022303201301311-2202123020213110-1101221231033011-3312312323221021-0120132110032300-3300200230123230"></a>

## Next pages — other_settings / 023322023320 / 5

- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233132100212301-0112010323313011-0021013012113331-3112311323223230-1022002220033003-2311101332202020-1221231033233120-3132230200313020"></a>

## other_settings.header_options — header_options / 121102122011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- other_settings.header_options

<a id="canonical-0120012201013020-1223101101033003-2320230302032330-2210121213203320-2100013023120233-3320230320003211-3333321233103120-1232221011330033"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS related to request/response headers.

Upstream description:

This defines various OPTIONS related to request/response headers.

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
header_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-3302322230001203-2021301122103221-0021113033000021-2213230102021223-1232002303103131-1202310201121232-1021113002323103-3233130003230322"></a>

## Direct properties — header_options / 121102122011 / 3

- [request_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202): complete subsection reference.

<a id="canonical-3313230232133303-2111301013112013-2230131011103332-3001113231030203-2030303002103210-3311333320200133-0310331111013021-3111023003321302"></a>

<a id="canonical-1123113231020222-3012300210233322-3223232121002122-2202221022002320-2131211131110221-1110030112010110-0022202330112210-2301003322112131"></a>

## request_headers_to_remove property — header_options / 121102122011 / 4

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP request being sent towards upstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.min_bytes": "1",
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322): complete subsection reference.

<a id="canonical-2300220121021030-1030311131100310-3213213031130331-1220300302213232-1120310220232110-2100312103312302-0231131211302231-0220121030233031"></a>

<a id="canonical-0320321101202131-0132113311012110-3023111122011010-0032332221010202-0103322211222100-1030101302030012-1202132120332021-1222032233331222"></a>

## response_headers_to_remove property — header_options / 121102122011 / 5

Type: `["list", "string"]`. Optional.

List of keys of Headers to be removed from the HTTP response being sent towards downstream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(32),
}
```

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
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.pattern": "^[0-9A-Za-z_\\\\-\\\\.]+$",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0333133013030222-2302221131020021-2303311311110233-0231223311331130-0312223002221032-2012033022201112-0331323203200333-0112103023300323"></a>

## Next pages — header_options / 121102122011 / 6

- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1103311032222101-1220030010122210-1332310003303101-3112020123330333-3122300223313303-0013212323212133-3002120021323321-0010333220230122"></a>

## other_settings.header_options.request_headers_to_add — request_headers_to_add / 002203221302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- other_settings.header_options.request_headers_to_add

<a id="canonical-2213300010232000-0110103232330030-0312122330132333-0122333201121313-1002000301131211-3210033201310312-1020111111310122-1303123112022020"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP request being routed towards upstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
request_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-3223020032100323-0020200120223232-1323222300022131-0322111320311332-1131023131031221-1022311102320100-3210103133123110-2321012303123013"></a>

## Direct properties — request_headers_to_add / 002203221302 / 3

<a id="canonical-3103132120221331-1123001131211212-2032221232013200-1312100110113210-0113033311300201-3131310220102110-3301031001103200-2333212021003221"></a>

<a id="canonical-1110113130130221-0230323120222211-3323211213030011-0312202131032122-1323320013200003-2123120011331010-3221111210122331-0030231122222023"></a>

## append property — request_headers_to_add / 002203221302 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-0232233301120101-2212210120111112-0133001332121232-3221033110033213-0322100013302322-1122212212203113-2322222233130310-0220123333033031"></a>

<a id="canonical-2201021323100221-1120122311221010-2321002033021302-3020203221121010-3100201301312120-2033322012230122-2323102330032300-3011032331030010"></a>

## name property — request_headers_to_add / 002203221302 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213): complete subsection reference.

<a id="canonical-3022011203323013-1301002301113101-1332330300221333-0222202010333310-1212222013112030-1000122303010103-2332321021103213-2010220131330323"></a>

<a id="canonical-0331113211021100-3323222102312021-3213233212132332-1012230113130233-1210031311021121-0322001203212122-3210133000112221-2010013023110303"></a>

## value property — request_headers_to_add / 002203221302 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1002013013111000-3032120310230103-2000101322012011-3320010322200100-1212031220033220-0223232130102321-1210223122003223-2201013320330231"></a>

## Next pages — request_headers_to_add / 002203221302 / 7

- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323210230033302-2211221211332220-3110221112030300-0112221122332112-2012213130221121-2102120230002323-1130201310312320-2000021312102120"></a>

## other_settings.header_options.request_headers_to_add.secret_value — secret_value / 002210032131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- other_settings.header_options.request_headers_to_add.secret_value

<a id="canonical-3331010111222222-2100013220011320-2122101320302311-0110213332211311-1111322101222101-0132233212312033-3111312213220120-2110312032223000"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-0100212122122131-1122113331132212-1211022132103113-2100013211230223-2120111100103023-3102311213233032-3200001303320200-0333223010111323"></a>

## Direct properties — secret_value / 002210032131 / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-2331223000110010-0323233123202112-0231333012102230-1330312202120211-1230301032201200-1331220133030230-3103220021230133-0133003221223221): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-3313023202202302-3320212100333032-1100112130203220-1121213002122120-3322233331333121-3020330101330211-2101201221231103-0100103002332111): complete subsection reference.

<a id="canonical-0220121013322022-2312230321231211-0013100132030200-2330003212013212-0201021011023222-3231120103023322-3300300321233010-0121101321002220"></a>

## Next pages — secret_value / 002210032131 / 4

- [other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-2331223000110010-0323233123202112-0231333012102230-1330312202120211-1230301032201200-1331220133030230-3103220021230133-0133003221223221)
- [other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-3313023202202302-3320212100333032-1100112130203220-1121213002122120-3322233331333121-3020330101330211-2101201221231103-0100103002332111)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2331223000110010-0323233123202112-0231333012102230-1330312202120211-1230301032201200-1331220133030230-3103220021230133-0133003221223221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3332130131221001-2213132200303202-1221102230321321-1131100133322120-2232000102000312-1301001012211330-2131320231001302-1313021323131232"></a>

## other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 231113111101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- other_settings.header_options.request_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-1113111123122232-1213133021302302-2133333120231312-0132301220321032-3013212233213313-3111322110011331-3121333002232110-3131313000300011"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2332120303110303-1002201231011222-2001003301113310-0311200222120100-3102032311223320-0011102302200032-2231202002003201-1000333112232232"></a>

## Direct properties — blindfold_secret_info / 231113111101 / 3

<a id="canonical-2110220223133230-0011202123303000-2311030122331121-0323222032222233-0010003031110333-3210102311210001-3111231233110302-3110323012100013"></a>

<a id="canonical-2030032311301100-0302321101303011-0103301113133210-0021130012121000-3332031111110200-3102002013030130-0120322012331022-0310110312233322"></a>

## decryption_provider property — blindfold_secret_info / 231113111101 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-2331232003201122-0021023020120103-0211123221313311-0102111312102301-2010101022130211-1213330333300103-3300333102221321-0020122233001112"></a>

<a id="canonical-3102013102013030-0203302331012300-0031111010323100-0302101113003230-0131232031302300-2312031230221000-1200100320001213-2303132012023302"></a>

## location property — blindfold_secret_info / 231113111101 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-2101301211123232-0313013013012323-3003032331301032-2312332231002332-3230310222101313-2223203010001113-1300232210322212-3120011102110201"></a>

<a id="canonical-0031031101133231-1023131232133213-3103101102030100-3221210133210121-0330230210300321-0103223323000120-0231022300331302-1112323030202120"></a>

## store_provider property — blindfold_secret_info / 231113111101 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-2230121000133232-2333123000322120-3121300303330113-0320211120300202-0211320032330310-1110203101003203-0123323212011231-1012230133110201"></a>

## Next pages — blindfold_secret_info / 231113111101 / 7

- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3313023202202302-3320212100333032-1100112130203220-1121213002122120-3322233331333121-3020330101330211-2101201221231103-0100103002332111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222223021320011-3231131020012232-0023133120113210-0001103232122021-1333232112030003-0330323300022132-2232012011221013-2102213120303132"></a>

## other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 103131322111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.request_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110132120230100-3011303312303032-0223020113312231-3203013020003030-1203320300330101-0200120112000023-3003011101102232-3300031033011202)
- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- other_settings.header_options.request_headers_to_add.secret_value.clear_secret_info

<a id="canonical-3011233032130321-2103211302133313-1121333010232200-0331000123300032-3032001203220201-2020031003223200-1121001222331130-0033011113220120"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200333123210303-3012213231303121-1320201031111000-0021201212123101-3103330022030331-3212312113101122-0122023211333002-3202300122210202"></a>

## Direct properties — clear_secret_info / 103131322111 / 3

<a id="canonical-1212331233322232-1300113101321020-1023102103211302-0013031300031210-0230011131323323-1313331211313013-0013030031221121-2031202113320130"></a>

<a id="canonical-3312031001032231-1200312321312221-3301010233311213-2210103323000021-2233133333222231-3212221001123101-2313011030213101-1310002200211203"></a>

## provider_ref property — clear_secret_info / 103131322111 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-3001032310102333-2323310302012001-0303032302012311-3313211303002100-2313020013233122-2011201211032211-2010332110312223-2001303103102103"></a>

<a id="canonical-0001231201111231-1310311202120112-2103133112121312-2000213213332121-3301130201300210-3122022012330203-1302302122230321-3201220220323203"></a>

## URL property — clear_secret_info / 103131322111 / 5

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
EnumExtractionComplete: true
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

<a id="canonical-0112310303301103-0330223121220222-3012232011312313-0023010331301223-3113131300203213-0121312023112331-2010000011000112-2223221012321220"></a>

## Next pages — clear_secret_info / 103131322111 / 6

- [other_settings.header_options.request_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1110121111222122-3332321003021022-2201011312323333-3200022220000103-2202033000112020-3230111000233111-0301111310220112-3331130120030213)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1203332003120211-2302131212033032-2023300133133321-3211331203121311-3011012100213112-3122212231030102-1311302011300003-3212230231202131"></a>

## other_settings.header_options.response_headers_to_add — response_headers_to_add / 202010112103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- other_settings.header_options.response_headers_to_add

<a id="canonical-2100133013302121-2012033220023302-0222002231313313-2222131312032021-0221203003023003-0113333330322123-0032133230100331-3212231220030220"></a>

Type: `"object"`. list nested block, Optional.

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Upstream description:

Headers are key-value pairs to be added to HTTP response being sent towards downstream. Headers
specified at this level are applied after headers from matched Route are applied.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("secret_value",
    "value")}
```

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
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
response_headers_to_add {
  # Configure direct properties listed below.
}
```

<a id="canonical-2131022111231021-1231130102131001-0310201321301001-3003033032112311-3123230133110331-3113201120213332-3313021003201332-1301021203213023"></a>

## Direct properties — response_headers_to_add / 202010112103 / 3

<a id="canonical-0120202310300101-0003121230201023-1221133000020113-1331000001103222-3330103103110313-1231122001223133-0123100331011333-0133133230303032"></a>

<a id="canonical-1013101221110112-2000110013022122-2332200203130300-3300130031330322-3202220233212331-1132100023213311-0101313010330033-1103133210331132"></a>

## append property — response_headers_to_add / 202010112103 / 4

Type: `"bool"`. Optional.

Should the value be appended? If true, the value is appended to existing values. not append.
Defaults to \`do\`.

Upstream description:

Should the value be appended? If true, the value is appended to existing values. Default value is do
not append.

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

<a id="canonical-3123220211030303-1333013322323212-2230220310201130-2010123323030021-3121210111222022-2303220301221113-2311110121030210-1031120001330010"></a>

<a id="canonical-3201023121130211-0132031320213223-0312330202302010-2131221121000213-0110120202001121-1001032321331301-1012103110133310-0203323303223112"></a>

## name property — response_headers_to_add / 202010112103 / 5

Type: `"string"`. Optional.

Name. Name of the HTTP header.

Upstream description:

Name of the HTTP header.

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122): complete subsection reference.

<a id="canonical-1300311232031203-2021130212202333-3201330123101231-1221002013231313-2203230031333113-2022000331120002-2233321231300131-2303130100303313"></a>

<a id="canonical-3310202322302321-3211020113031203-0030301130133213-3031323323131323-0211320122133200-3230330131012102-3333130330202032-2111003010121022"></a>

## value property — response_headers_to_add / 202010112103 / 6

Type: `"string"`. Optional.

Exclusive with \[secret\_value\] Value of the HTTP header.

Upstream description:

Exclusive with \[secret\_value\] Value of the HTTP header.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(8096),
}
```

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
    "ves.io.schema.rules.string.max_len": "8096"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "8096"
  }
}
```

<a id="canonical-1311020221022201-2023130311233323-2101313103200111-2121320020201023-1013132003331213-3223212020231123-3121313220012031-1220021313230233"></a>

## Next pages — response_headers_to_add / 202010112103 / 7

- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021212121001000-1013332021101101-3322202003011102-3221212000331323-2122301232103102-1022333120111222-0300200010303101-0101323202130323"></a>

## other_settings.header_options.response_headers_to_add.secret_value — secret_value / 333121020022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- other_settings.header_options.response_headers_to_add.secret_value

<a id="canonical-2110103003010301-3213300103203013-0023222012012030-2221033001311233-3202021212133233-0000113232213302-0033300213010103-2002032301113121"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("blindfold_secret_info",
    "clear_secret_info")}
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
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
secret_value {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323122101123130-1113030000101101-0312313120203111-1203232312012121-1002130231120301-3130230211203303-0113331102230131-0330322001332302"></a>

## Direct properties — secret_value / 333121020022 / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-0110121132222012-0300301023333333-3030111211200313-3011123032013123-0313010311023311-1223212111130233-3210022201020223-2022201103322320): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-0103203311320333-0012221221223323-2121322202011033-0032311010003313-3200323120200113-0132231003130100-3100310000113203-2231110011203232): complete subsection reference.

<a id="canonical-0202212112131012-0022120032331110-3210100322300233-1102123113302232-2313113103112120-3200201211300002-1310212321231110-3303300122023003"></a>

## Next pages — secret_value / 333121020022 / 4

- [other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-0110121132222012-0300301023333333-3030111211200313-3011123032013123-0313010311023311-1223212111130233-3210022201020223-2022201103322320)
- [other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info](resources--cdn_loadbalancer--reference--group-013.md#canonical-0103203311320333-0012221221223323-2121322202011033-0032311010003313-3200323120200113-0132231003130100-3100310000113203-2231110011203232)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0110121132222012-0300301023333333-3030111211200313-3011123032013123-0313010311023311-1223212111130233-3210022201020223-2022201103322320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1322022130320100-0201320031013001-1210232301301213-0311200300111100-1102200320102300-2331012020232302-3002101111003131-1203011022220110"></a>

## other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info — blindfold_secret_info / 100011103113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- other_settings.header_options.response_headers_to_add.secret_value.blindfold_secret_info

<a id="canonical-0112332002200010-2023311110302211-1023220011102331-1232022223321312-3000210322011320-3203102313330122-3022102113313323-1213203201112232"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("location")}
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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123222203230110-1221331230022012-3320300002220332-0202330020201303-2222302020022032-2130013133312120-2322210011212122-3233231010212003"></a>

## Direct properties — blindfold_secret_info / 100011103113 / 3

<a id="canonical-1202312122323002-1130120131102031-3230200302102232-0122300131002303-2302012202231021-2010311001333332-1222332020232123-3031200031320011"></a>

<a id="canonical-3202213333310100-3211321013310302-1133231333231320-1202013021221101-3220010032133132-0010130010123321-0132212012230230-1100111300121321"></a>

## decryption_provider property — blindfold_secret_info / 100011103113 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

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

<a id="canonical-0221323133101313-1032012023303301-2201002302111022-1112012211213120-0201111000331003-0132122013112030-3200233311202213-2103300133021203"></a>

<a id="canonical-0003222213031103-3113202330103001-1013011322232113-3212032323020322-2202221300231113-0033231201331013-0033120230102011-3300033223231022"></a>

## location property — blindfold_secret_info / 100011103113 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 131072),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
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
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1233030330020320-1111110211131311-2001220132001323-3123121001031131-0121101203133220-3112223033133333-1332321321102223-1300201112212000"></a>

<a id="canonical-3113213313101003-0312022112311033-0333031133003313-2230001111022221-1132021301002123-1301023033223103-0033323123323120-2121120311132112"></a>

## store_provider property — blindfold_secret_info / 100011103113 / 6

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Upstream description:

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

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

<a id="canonical-0011100111100210-1332120321021122-2100302312003221-0330001212110023-1202331303031231-3211030110232330-3013021131003013-2231120132113131"></a>

## Next pages — blindfold_secret_info / 100011103113 / 7

- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0103203311320333-0012221221223323-2121322202011033-0032311010003313-3200323120200113-0132231003130100-3100310000113203-2231110011203232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3210001323220021-1003003212223003-0131212100301320-3312221011020201-1220232332021232-2021330023122200-0000130023323322-2111223011223021"></a>

## other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info — clear_secret_info / 301131011110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.header_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-0231012000310112-0131333323001213-3313312301320013-0200310101320310-1130103103000011-3201312132222010-0203032113203131-1133123210303123)
- [other_settings.header_options.response_headers_to_add](resources--cdn_loadbalancer--reference--group-013.md#canonical-0313220232130302-1102330030113302-3210000033313102-0333221223223031-0230232012102201-2130231123011003-0103312230110200-2230220111322322)
- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- other_settings.header_options.response_headers_to_add.secret_value.clear_secret_info

<a id="canonical-2311133021320023-2133001231310022-2322131301232223-3102331332221301-0033322213123233-3033110132230322-2113032323032233-1113212213331132"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("url")}
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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100213001301130-1202202332231231-2033220222000333-2211011103011112-0120113301100223-2313311213301030-2202312131000101-2333211113023330"></a>

## Direct properties — clear_secret_info / 301131011110 / 3

<a id="canonical-1301331330001002-0131110022123002-2113320200033002-3303022303202332-1312333011103131-2202210111221003-0303220112223203-0031110130022231"></a>

<a id="canonical-2310131332222003-0023113012230100-2103000132132203-2110311022112112-1001002213011310-3232321130210130-1323021101323020-1221102111312222"></a>

## provider_ref property — clear_secret_info / 301131011110 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1113132002110103-2330122303201312-2323313001131121-0123032213110122-2003013312012023-1110030021203033-0133233023012120-1112301002112133"></a>

<a id="canonical-3131022210002301-3201201200103021-3332201100130333-0000000331323031-2221120320021000-2130323320303201-1010211001221111-3121130300211123"></a>

## URL property — clear_secret_info / 301131011110 / 5

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
EnumExtractionComplete: true
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

<a id="canonical-3332331122113021-3123101322110311-2122132312330333-3112300221010231-1212333133020221-2302302102033121-2001013232032022-3002101311033201"></a>

## Next pages — clear_secret_info / 301131011110 / 6

- [other_settings.header_options.response_headers_to_add.secret_value](resources--cdn_loadbalancer--reference--group-013.md#canonical-1301121233102131-0121113330131231-2022211232130111-1102323112311111-0022113321320210-2221111100300010-0001311013030011-3302133033013122)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0233311130111113-2212233223031111-3031032211223033-0222103101200222-3201101310313303-0213010203013302-2120023222101301-2102211321123202"></a>

## other_settings.logging_options — logging_options / 331133010232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- other_settings.logging_options

<a id="canonical-2033033311310033-3112233220000212-1322020022231103-2131003221022103-0311112120332210-2131300130332320-2130303332121013-0233030022123001"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS related to logging.

Upstream description:

This defines various OPTIONS related to logging.

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
logging_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232200112011100-3212022302102113-0033102001102230-0211320201313233-1130022223012102-3301013302220102-0320230333222330-1120123312110103"></a>

## Direct properties — logging_options / 331133010232 / 3

- [client_log_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-3100220303110020-0030111031001233-0010210000322313-1211310232011312-0233010320311301-2000102113132210-1320131102010223-3123110102211210): complete subsection reference.

- [origin_log_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-1101322013323301-2223110001000120-2120123233101121-3222003111233313-0333012210002223-1032223320210211-1130233103030323-1200113031130213): complete subsection reference.

<a id="canonical-0330120323202321-2213333000033031-1122332001213121-1011013031103200-3110323133221021-3212002231020322-3113233202032300-1110303321332332"></a>

## Next pages — logging_options / 331133010232 / 4

- [other_settings.logging_options.client_log_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-3100220303110020-0030111031001233-0010210000322313-1211310232011312-0233010320311301-2000102113132210-1320131102010223-3123110102211210)
- [other_settings.logging_options.origin_log_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-1101322013323301-2223110001000120-2120123233101121-3222003111233313-0333012210002223-1032223320210211-1130233103030323-1200113031130213)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3100220303110020-0030111031001233-0010210000322313-1211310232011312-0233010320311301-2000102113132210-1320131102010223-3123110102211210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300122323023333-2013222212333023-3221113110211210-1223213022200200-0231302323001212-1210010203212211-2003221101001321-0301132330030232"></a>

## other_settings.logging_options.client_log_options — client_log_options / 003300010203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- other_settings.logging_options.client_log_options

<a id="canonical-3020233031032001-3201110011132302-2201300010020331-1023202122110123-3000102123102133-3120020003022223-0120030233213111-0031033322312100"></a>

Type: `"object"`. single nested block, Optional.

Headers to Log. List of headers to Log.

Upstream description:

List of headers to Log.

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
client_log_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0023330103111003-2111023331112330-3303122002112230-1322201332121111-2301001223323010-0322301333302000-1210232212323311-2211303131110112"></a>

## Direct properties — client_log_options / 003300010203 / 3

<a id="canonical-1233012030130202-1023102102231022-0222222001110020-1030012310332100-2233223020203103-2322331320213013-3223013301010210-1012321232033213"></a>

<a id="canonical-0322011100200230-0203032332223333-1020003333011033-3113003103012313-3102103333222312-1223203200020311-3200213033212121-2100012000102122"></a>

## header_list property — client_log_options / 003300010203 / 4

Type: `["list", "string"]`. Optional.

Headers. List of headers.

Upstream description:

List of headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3332103300112020-0232230030022000-3310121231233313-1333120133001303-0233212002032311-3101320311003323-1310322333203323-0322102310122301"></a>

## Next pages — client_log_options / 003300010203 / 5

- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1101322013323301-2223110001000120-2120123233101121-3222003111233313-0333012210002223-1032223320210211-1130233103030323-1200113031130213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3000101000312332-1121010332022021-0233323223230020-1310111333220131-0100031202103213-0130320021223001-1320220130002020-2302312020130020"></a>

## other_settings.logging_options.origin_log_options — origin_log_options / 212232320233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [other_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-1202031330011101-0332311032231013-2313101330202033-1021201330201212-0201100211020002-3013210230001210-2223102333313322-3023223103202022)
- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- other_settings.logging_options.origin_log_options

<a id="canonical-3112203131202231-2220133232103133-0221031232322333-0321132211002021-1221101331011213-1210003033123122-1213333031301302-2323211021031001"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for origin log options.

Upstream description:

List of headers to Log.

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
origin_log_options {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311300123211220-1021332201320022-1200330222012210-1233120233312130-0310312100032123-2201110031003030-0332211110121222-0011302020323322"></a>

## Direct properties — origin_log_options / 212232320233 / 3

<a id="canonical-1313310301020133-1130021212132030-0203232122021213-1212013300001012-3323120222003320-3223011301032120-2201120322302300-0001003300322303"></a>

<a id="canonical-3221131003031230-1221311100303032-1121212213321122-2012020322022123-3011131222022200-0121312133010000-0021010030232312-3121001222302300"></a>

## header_list property — origin_log_options / 212232320233 / 4

Type: `["list", "string"]`. Optional.

Headers. List of headers.

Upstream description:

List of headers.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0303231330331211-3010012001100133-2300220303133102-3123132221120010-0123033333311131-0112012312213332-0310101003320232-2130333231323201"></a>

## Next pages — origin_log_options / 212232320233 / 5

- [other_settings.logging_options](resources--cdn_loadbalancer--reference--group-013.md#canonical-2133311220012202-3101322300331201-1102030033220301-0033201113232321-2330221112122033-2231012130010300-1310323230120330-1112212111233130)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133221103112111-0111132110110330-0331313021013031-0301110310300000-2023223112010321-1201233010121020-2123311002133211-2122301222131112"></a>

## policy_based_challenge — policy_based_challenge / 303001202130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- policy_based_challenge

<a id="canonical-3110233220333200-3212201333323201-2020113221230221-3322221202220102-2132220310202212-2231113002233113-3013031220202301-3203221133222020"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings for policy rule based challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "always_enable_js_challenge"),
  validators.ConflictingObjectAttributes("always_enable_captcha_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("always_enable_js_challenge",
    "no_challenge"),
  validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation"),
  validators.ConflictingObjectAttributes("default_temporary_blocking_parameters",
    "temporary_user_blocking")}
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
  "x-ves-oneof-field-captcha_challenge_parameters_choice": "[\"captcha_challenge_parameters\",\"default_captcha_challenge_parameters\"]",
  "x-ves-oneof-field-challenge_choice": "[\"always_enable_captcha_challenge\",\"always_enable_js_challenge\",\"no_challenge\"]",
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]",
  "x-ves-oneof-field-temporary_blocking_parameters_choice": "[\"default_temporary_blocking_parameters\",\"temporary_user_blocking\"]"
}
```

Terraform syntax:

```terraform
policy_based_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2201111112020013-1220231321122112-0133221332020132-3300002013232020-3133031203300210-1103132302302033-3103130323220120-2231203120001222"></a>

## Direct properties — policy_based_challenge / 303001202130 / 3

- [always_enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3231100132202302-1113312322120131-1133222211000302-2032310321101021-3323031100331200-0021213230032033-0021220113101101-0303301220221031): complete subsection reference.

- [always_enable_js_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3220313211233110-3203030311210203-1321222221301101-1311012231310212-1002121131030213-1332100110321123-1313111220310010-3112031113010121): complete subsection reference.

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-1323000010110021-0032203011323300-0232330032321021-3212103130130000-0310233331320031-1320030320020332-2123111310102012-2130301311021333): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-3121331210122303-0333010020032030-2311300133222010-3103013023220220-0032113202212121-1033021203220133-1020013323120333-3130220030122211): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-1221023230023032-0022213230332233-1213130311032013-3101122202311233-3211011101333210-3031100032112110-0303320212301221-2011302321120101): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-0120231001300001-3322103130120122-3301130030331000-1001202220203231-2211200202233330-0302211021201312-0303302302033012-1012010021111100): complete subsection reference.

- [default_temporary_blocking_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-2022330021132232-2102221101220110-1010002220220113-3323023011020013-0103130102102110-1122300013010123-0103121301130020-2113313130200320): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-2013120300230010-1203120210000023-2003001321022311-1201003231201221-1322132022103111-3103302022311213-3030323332012121-0321033313031002): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-013.md#canonical-2213030000002202-0023323001313300-0100232113302220-2320331033032120-2132100010013330-2111101121100221-1132222032102331-1203323123332212): complete subsection reference.

- [no_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213200121002130-3323213132112330-2113023033200123-1131211303231121-3200332321302331-1110000333012132-2102221010223301-2021200003100300): complete subsection reference.

- [rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301): complete subsection reference.

- [temporary_user_blocking](resources--cdn_loadbalancer--reference--group-014.md#canonical-2020113203001122-0131003133101010-3112033112103122-1230132210230022-1110101121230003-3011322320122320-2001301222311323-1030230313310012): complete subsection reference.

<a id="canonical-3030121311111132-2133302123221133-2021120020221221-1301210102110023-0123231103010121-3013102202301101-3000332013120333-3033223120020313"></a>

## Next pages — policy_based_challenge / 303001202130 / 4

- [policy_based_challenge.always_enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3231100132202302-1113312322120131-1133222211000302-2032310321101021-3323031100331200-0021213230032033-0021220113101101-0303301220221031)
- [policy_based_challenge.always_enable_js_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3220313211233110-3203030311210203-1321222221301101-1311012231310212-1002121131030213-1332100110321123-1313111220310010-3112031113010121)
- [policy_based_challenge.captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-1323000010110021-0032203011323300-0232330032321021-3212103130130000-0310233331320031-1320030320020332-2123111310102012-2130301311021333)
- [policy_based_challenge.default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-3121331210122303-0333010020032030-2311300133222010-3103013023220220-0032113202212121-1033021203220133-1020013323120333-3130220030122211)
- [policy_based_challenge.default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-1221023230023032-0022213230332233-1213130311032013-3101122202311233-3211011101333210-3031100032112110-0303320212301221-2011302321120101)
- [policy_based_challenge.default_mitigation_settings](resources--cdn_loadbalancer--reference--group-013.md#canonical-0120231001300001-3322103130120122-3301130030331000-1001202220203231-2211200202233330-0302211021201312-0303302302033012-1012010021111100)
- [policy_based_challenge.default_temporary_blocking_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-2022330021132232-2102221101220110-1010002220220113-3323023011020013-0103130102102110-1122300013010123-0103121301130020-2113313130200320)
- [policy_based_challenge.js_challenge_parameters](resources--cdn_loadbalancer--reference--group-013.md#canonical-2013120300230010-1203120210000023-2003001321022311-1201003231201221-1322132022103111-3103302022311213-3030323332012121-0321033313031002)
- [policy_based_challenge.malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-013.md#canonical-2213030000002202-0023323001313300-0100232113302220-2320331033032120-2132100010013330-2111101121100221-1132222032102331-1203323123332212)
- [policy_based_challenge.no_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213200121002130-3323213132112330-2113023033200123-1131211303231121-3200332321302331-1110000333012132-2102221010223301-2021200003100300)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.temporary_user_blocking](resources--cdn_loadbalancer--reference--group-014.md#canonical-2020113203001122-0131003133101010-3112033112103122-1230132210230022-1110101121230003-3011322320122320-2001301222311323-1030230313310012)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3231100132202302-1113312322120131-1133222211000302-2032310321101021-3323031100331200-0021213230032033-0021220113101101-0303301220221031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0112233021333012-1232133301133223-2112333203110032-0123030332122210-3320333130033120-2122033120311220-1331023000020033-0032233301302013"></a>

## policy_based_challenge.always_enable_captcha_challenge — always_enable_captcha_challenge / 223330013100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.always_enable_captcha_challenge

<a id="canonical-1131120110030330-2121022011120212-3121010110212203-0301010000032110-0032320221300310-0030301302231031-3012202210112333-3030302201303101"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable captcha challenge.

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
always_enable_captcha_challenge = {}
```

<a id="canonical-0333311033210310-1130213121313203-2232030211101013-2221121122121232-3223313033330113-3121003003000121-0333030102320020-2300120223022131"></a>

## Direct properties — always_enable_captcha_challenge / 223330013100 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122202111321020-3211233300303202-1001000131010231-3322211323112323-2223202102032202-3210221130201131-1313213301131213-0000003120221311"></a>

## Next pages — always_enable_captcha_challenge / 223330013100 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3220313211233110-3203030311210203-1321222221301101-1311012231310212-1002121131030213-1332100110321123-1313111220310010-3112031113010121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0320223003303010-0101320111322011-2123210133232313-2101300321022033-2023300130103003-3300031322331301-1120332303133211-2231113102312010"></a>

## policy_based_challenge.always_enable_js_challenge — always_enable_js_challenge / 021022210223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.always_enable_js_challenge

<a id="canonical-0121203231131213-3120133332233320-3013211300022321-2330013301232131-1302320221010310-2003330202213123-3300031122023110-0210020201320300"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for always enable js challenge.

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
always_enable_js_challenge = {}
```

<a id="canonical-3022231312122301-2102003122313102-1030033221101213-3003001013211230-2132301130330002-1023333001003200-2113202132212001-1112333111100222"></a>

## Direct properties — always_enable_js_challenge / 021022210223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021223201110132-3203333112130321-1103210203232203-2033000130302210-2321322003020030-0213102003110111-2131031321220212-2320111030001002"></a>

## Next pages — always_enable_js_challenge / 021022210223 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1323000010110021-0032203011323300-0232330032321021-3212103130130000-0310233331320031-1320030320020332-2123111310102012-2130301311021333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3133320031321200-1112110310332011-2200320321113313-1312310021331303-1002311033313203-0320222213313331-2100102030233013-2233032233221022"></a>

## policy_based_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 332011332132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.captcha_challenge_parameters

<a id="canonical-3210133222033210-0303032110030301-1220310011130010-2101212131112203-1031002120002012-2323301311022000-3330012232302323-1331233323321011"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry")}
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
captcha_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0010020001201113-1002130323111223-1313133032201233-3033113213230012-3000133301332133-0323331320121232-2013212130123321-1101300100312130"></a>

## Direct properties — captcha_challenge_parameters / 332011332132 / 3

<a id="canonical-2231213221002121-2010200023222031-3031012201121023-1120320321200333-0012331332330012-3033120230023331-0223222231210323-0011313222003122"></a>

<a id="canonical-1003133331010031-3230030103112123-1000220112231322-2133003210331210-2321000133231100-2033321100111120-0222013200003013-0033333102211301"></a>

## cookie_expiry property — captcha_challenge_parameters / 332011332132 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0013000233133012-0031121131213211-1110121003130030-0333123233233201-1203330222213111-1100100222211032-3021003330003200-2123102220031302"></a>

<a id="canonical-3302232311011201-3212121333122331-2210213102003202-3303330132201030-2010011120130122-3131120303213311-2202320000303333-0312332120333212"></a>

## custom_page property — captcha_challenge_parameters / 332011332132 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-1012013033311223-3223231123330203-0111122003032001-2333111110101322-0212021011221120-0311201102230230-0033311203313330-2100103110000112"></a>

## Next pages — captcha_challenge_parameters / 332011332132 / 6

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3121331210122303-0333010020032030-2311300133222010-3103013023220220-0032113202212121-1033021203220133-1020013323120333-3130220030122211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2133023320200133-3232123313231033-2320302113222020-3211032013102103-2232212011300223-1113330330223322-1220321020132330-1030203312032211"></a>

## policy_based_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 021123230332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="canonical-2333021013130001-3303131231231130-2201323131321121-0200202130213031-0310013030130321-2021201021100100-0133222302331131-0313322100303302"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_captcha_challenge_parameters = {}
```

<a id="canonical-1201201122210202-3011302202202121-0103121230022102-3212021201122033-0111000033233300-0130033332121302-1301102220222203-3211321132331230"></a>

## Direct properties — default_captcha_challenge_parameters / 021123230332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303331010312132-2301133220201102-0011120021123331-2131001103221200-0122322311001311-3213233031311003-1001332022101101-3001132231012000"></a>

## Next pages — default_captcha_challenge_parameters / 021123230332 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1221023230023032-0022213230332233-1213130311032013-3101122202311233-3211011101333210-3031100032112110-0303320212301221-2011302321120101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301221233320112-2012132102210333-1100322323000211-2010030300122023-2001220322011113-0302132301123312-0312310100330320-1300121010023133"></a>

## policy_based_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 032020120210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_js_challenge_parameters

<a id="canonical-0131110301113233-0221003002212302-2220223023001221-1231013031200030-1232013322113113-2321033312311323-0210232111000321-3202101101221312"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
default_js_challenge_parameters = {}
```

<a id="canonical-1012231002201212-3302001001232112-1002131132322331-3323112031120110-0331322033133210-1313201101023331-0030112212212302-1031322221100130"></a>

## Direct properties — default_js_challenge_parameters / 032020120210 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2021003123102222-0033120113033100-0120302233213312-1213122231110320-0203313301030303-2230310123232022-0223032031203212-0110333112130221"></a>

## Next pages — default_js_challenge_parameters / 032020120210 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0120231001300001-3322103130120122-3301130030331000-1001202220203231-2211200202233330-0302211021201312-0303302302033012-1012010021111100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333013233223222-0210221311131202-1133230233120103-0113112102201130-0330111313320302-2111312222031112-2102130111232110-0311233111032102"></a>

## policy_based_challenge.default_mitigation_settings — default_mitigation_settings / 000210210212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_mitigation_settings

<a id="canonical-0032223011111213-1231211321302322-3022212312221033-2233032331033023-3230330232133203-3012230330131300-2231300122130213-0111223010331232"></a>

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
default_mitigation_settings = {}
```

<a id="canonical-1012030102322323-3221330013221110-3310023111110303-3133310203332232-1013000101130310-3013301100033011-2110301013003013-1100112022003003"></a>

## Direct properties — default_mitigation_settings / 000210210212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2103123333312320-2120300012031202-1032021213100223-1212322233200203-2301031130312110-1332231201220332-1031033330320321-3312002302212133"></a>

## Next pages — default_mitigation_settings / 000210210212 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2022330021132232-2102221101220110-1010002220220113-3323023011020013-0103130102102110-1122300013010123-0103121301130020-2113313130200320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301001332033232-0312031013013300-0300202222223200-3321303201110301-1120312221033321-3103130013332112-3122211103300020-2021031212101203"></a>

## policy_based_challenge.default_temporary_blocking_parameters — default_temporary_blocking_parameters / 213331120112 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.default_temporary_blocking_parameters

<a id="canonical-3022111031011030-0313211220122230-3301102013121223-1311301021112031-2020222331330120-0001201313032101-1302230302201201-1012231113323200"></a>

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
default_temporary_blocking_parameters = {}
```

<a id="canonical-3232331122330230-1221213123303101-2330232332222122-0030303031122202-1331011322331131-3113132101101213-1210131021232030-3123013230012110"></a>

## Direct properties — default_temporary_blocking_parameters / 213331120112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0230111132012203-3102101021132102-0033231233110331-3123300211030202-3223213212310203-0313321232023101-3100121302112022-0121002003102322"></a>

## Next pages — default_temporary_blocking_parameters / 213331120112 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2013120300230010-1203120210000023-2003001321022311-1201003231201221-1322132022103111-3103302022311213-3030323332012121-0321033313031002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311310031033022-0220212220330010-1222013330303303-2301100000200003-2031233200031322-3220321331132000-1303232121323333-0033330003110230"></a>

## policy_based_challenge.js_challenge_parameters — js_challenge_parameters / 331103012302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.js_challenge_parameters

<a id="canonical-0033101233000012-0000301013201301-1230230132113303-1112320010000300-0331003323332021-2010023132211210-1322301300232330-0311000133201020"></a>

Type: `"object"`. single nested block, Optional.

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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cookie_expiry",
    "js_script_delay")}
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
js_challenge_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-0301010313202210-0300201323200232-0321110223131132-3322132032230333-0013021233320101-0122100201001032-0302223310212112-3212030221103201"></a>

## Direct properties — js_challenge_parameters / 331103012302 / 3

<a id="canonical-2010021012230323-3330222231032332-1312110330211311-3231303020302113-1311000320133020-2110013300233200-2010322012313321-1000103031320311"></a>

<a id="canonical-3131300301101111-3300002311202220-1321320021122002-1123211201321022-3003302322010002-3232313123020111-2110030302122230-3313323230000301"></a>

## cookie_expiry property — js_challenge_parameters / 331103012302 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1, 86400),
}
```

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
    "ves.io.schema.rules.uint32.lte": "86400"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "86400"
  }
}
```

<a id="canonical-0100233011111322-2232013101131321-1331010202121101-0212333300122232-2311231221002032-3321123200323300-0312022112332301-0300002103122030"></a>

<a id="canonical-0332113302033232-2121222013013102-1311313111030131-2032222211030321-2231313322312202-3111213013221333-0222232030212201-1331100333223021"></a>

## custom_page property — js_challenge_parameters / 331103012302 / 5

Type: `"string"`. Optional.

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format.

Upstream description:

Custom message is of type URI\_ref. Currently supported URL schemes is string:///. For string:///
scheme, message needs to be encoded in base64 format. You can specify this message as base64 encoded
plain text message e.g. "Please Wait.." or it can be HTML paragraph or a body string encoded as
base64 string E.g. "&lt;p&gt; Please Wait &lt;/p&gt;". base64 encoded string for this HTML is
"PHA+IFBsZWFzZSBXYWl0IDwvcD4="

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(65536),
}
```

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
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "65536",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3233313101111033-2221120123131300-2032103233310322-1032130302030201-3023203302003111-3133223210012213-3321232230330221-1212323012220011"></a>

<a id="canonical-3011023333310230-1310011100232231-1003103303223113-2220010223321311-1013021112112010-1102101303102300-3013230112302120-3201200032113211"></a>

## js_script_delay property — js_challenge_parameters / 331103012302 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.Between(1000, 60000),
}
```

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

<a id="canonical-0330133300210301-0200232221222311-2303203012003132-1012333331211121-2223100303002333-2233312202123232-2310010210100320-0300113203130012"></a>

## Next pages — js_challenge_parameters / 331103012302 / 7

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2213030000002202-0023323001313300-0100232113302220-2320331033032120-2132100010013330-2111101121100221-1132222032102331-1203323123332212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232223331020030-3033130013210010-0013123312221203-0123303330333210-3023023131232130-2201322200120221-1311112101320131-3213312000210300"></a>

## policy_based_challenge.malicious_user_mitigation — malicious_user_mitigation / 130232101123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.malicious_user_mitigation

<a id="canonical-2311303230123023-3132330022020112-1333202130031322-1022120030031201-3203003330100113-3111332102303323-0212333003200202-2032301211213002"></a>

Type: `"object"`. single nested block, Optional.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2023230330021010-3212100223131220-3021200232302130-1323200300011120-0112133103303022-0202313330130321-1102233012021120-2212100000023022"></a>

## Direct properties — malicious_user_mitigation / 130232101123 / 3

<a id="canonical-0210203220323123-0311121102223133-1100131231133212-3222010110021133-3313230010032303-2231322202230021-1010031212201220-3221121223310023"></a>

<a id="canonical-3331121222330330-0230211012133223-3013121230221102-2121310001002112-1121010323100101-1301311202030021-3101131111002332-2201231230101111"></a>

## name property — malicious_user_mitigation / 130232101123 / 4

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-3013323212221331-2210322102103101-2330322322323021-0223101323321031-2231032210220213-2200332121131113-1033202003233022-2120031321233131"></a>

<a id="canonical-2203310131222031-1012332001123201-2212101130103123-3120323232123120-0011132323230332-0211201030331003-0222001121221031-2010300023121011"></a>

## namespace property — malicious_user_mitigation / 130232101123 / 5

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2103321000331011-2131302121010202-1303220331021330-2213100301223013-0220323120123332-2101200111023312-0020232132200123-0213133303030330"></a>

<a id="canonical-1231120032103310-1002300321213333-1122113101333301-3131130210110310-0132332120010030-1010123002000310-3311212030010131-3021010030313303"></a>

## tenant property — malicious_user_mitigation / 130232101123 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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

<a id="canonical-2230022313031003-2133130110030232-0030021213132123-3313221330101103-0123112102312101-3323110023301110-1033111032120212-3311202032212311"></a>

## Next pages — malicious_user_mitigation / 130232101123 / 7

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3213200121002130-3323213132112330-2113023033200123-1131211303231121-3200332321302331-1110000333012132-2102221010223301-2021200003100300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310203222101202-2000233223013303-2102233212133121-0312232120310020-2222232021332320-1222113311030122-0230133212323133-2322312220223202"></a>

## policy_based_challenge.no_challenge — no_challenge / 132311030032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.no_challenge

<a id="canonical-0203202320131112-2302320221210211-3212131320032120-0320111122010132-0330111000211332-1210110133233323-2201113130103030-2031022102122131"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no challenge.

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
no_challenge = {}
```

<a id="canonical-3130213210112033-1031311002210033-1302221201333102-1111220210223123-1210310121332132-1121130221123123-0011200101302333-0130300220220332"></a>

## Direct properties — no_challenge / 132311030032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0101100222000230-3002031011030133-2001303133333023-0121311230000300-1101222101312220-1002031213103300-3002323010301021-1322312113320022"></a>

## Next pages — no_challenge / 132311030032 / 4

- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1030230222011133-2133003030332211-2310120020002300-2300000000213102-1300031331322321-0023032320020123-1302121331321113-0303222030001011"></a>

## policy_based_challenge.rule_list — rule_list / 323010132231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- policy_based_challenge.rule_list

<a id="canonical-2123222330010033-0203022221121212-2113002200232332-3211212012222233-0111101003133132-2031003031110011-3233303210233021-0211120333103132"></a>

Type: `"object"`. single nested block, Optional.

List of challenge rules to be used in policy based challenge.

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
rule_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1102333133110232-2002220221301331-3232220231222023-2302113021130333-1130230001001221-3201313220313232-3011101300222233-0100122200210332"></a>

## Direct properties — rule_list / 323010132231 / 3

- [rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111): complete subsection reference.

<a id="canonical-3331331021013012-2130211221222330-3231112000000211-1222112331102213-2001313023121333-0333303130312000-2031202221012332-3003320223020013"></a>

## Next pages — rule_list / 323010132231 / 4

- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2301131101132230-0110121133030203-1323313331022113-1031233130303033-2313021213331331-1120332312321121-3111121202101300-1133033313011300"></a>

## policy_based_challenge.rule_list.rules — rules / 301212100003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- policy_based_challenge.rule_list.rules

<a id="canonical-1101021211020011-0030321102302103-3003000331310110-0123033123211030-0221003232023011-2230023210001323-2303323322303303-2012020013101331"></a>

Type: `"object"`. list nested block, Optional.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

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
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
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

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033223332300302-3021100213232333-1210300122111112-0231232000111122-3213232302330203-1313003120211132-2221310030102233-0302301300223110"></a>

## Direct properties — rules / 301212100003 / 3

- [metadata](resources--cdn_loadbalancer--reference--group-013.md#canonical-3312022112120312-2330122132101001-0023311121210013-0120002103000332-2202100021232030-3133210320123312-0311302202001302-0322131111030012): complete subsection reference.

- [spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201): complete subsection reference.

<a id="canonical-0023312011301131-2000023221232131-2303210221120210-3012230031003122-2110233201213322-3010033100100202-2033211223312210-1320101333303201"></a>

## Next pages — rules / 301212100003 / 4

- [policy_based_challenge.rule_list.rules.metadata](resources--cdn_loadbalancer--reference--group-013.md#canonical-3312022112120312-2330122132101001-0023311121210013-0120002103000332-2202100021232030-3133210320123312-0311302202001302-0322131111030012)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3312022112120312-2330122132101001-0023311121210013-0120002103000332-2202100021232030-3133210320123312-0311302202001302-0322131111030012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332130100022203-0031311203102011-3220220000322220-1213013011013310-0323220112113312-2303100323003130-3130021231311121-3020223033120021"></a>

## policy_based_challenge.rule_list.rules.metadata — metadata / 203011121123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- policy_based_challenge.rule_list.rules.metadata

<a id="canonical-0101332312003220-0213301222131132-1301030213013320-1311201032302202-2113001233131031-3121021333332102-3301313213032331-3222220200332313"></a>

Type: `"object"`. single nested block, Optional.

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during
create..

Upstream description:

MessageMetaType is metadata (common attributes) of a message that only certain messages have. This
information is propagated to the metadata of a child object that gets created from the containing
message during view processing. The information in this type can be specified by user during create
and replace APIs.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013232112323301-1033310000200133-3121120301210333-3202310230113030-2323120312111022-0333101101333000-2333110003103330-1013122001012101"></a>

## Direct properties — metadata / 203011121123 / 3

<a id="canonical-1333021033320221-3120102123212323-3003122103122212-1220003233013002-1101032110013032-3103032131103002-0030303202211102-3232211113120130"></a>

<a id="canonical-0301232133103133-2300133020313212-1200310000112012-3132102321133133-3310210320130112-0301023121100230-3132132300003113-1111313312120001"></a>

## description_spec property — metadata / 203011121123 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2333300211031022-1222300310101122-3000113000323213-3212220302131122-0131030311021333-0033032230113120-0131132131022303-3313132030103111"></a>

<a id="canonical-0012220001312320-1032203121330311-3120132123023231-3331213331023231-0020332331010312-1303331302311113-1111010010030000-2022110210322321"></a>

## name property — metadata / 203011121123 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

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

<a id="canonical-0313310332031110-2310220222311312-0303002001200130-3000230231231232-0311201320032033-0312032011312013-3330122313302312-3032222011010033"></a>

## Next pages — metadata / 203011121123 / 6

- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010323020212330-3200001013012332-1320023301013123-1331300302020213-3331000301222013-2211103230302112-0023011120231032-1321100003100020"></a>

## policy_based_challenge.rule_list.rules.spec — spec / 010122213221 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- policy_based_challenge.rule_list.rules.spec

<a id="canonical-2100022312032302-2032323222000310-3320020323323030-0311311230302003-0203013303301133-1303003301003030-0233222330313000-1012020030201303"></a>

Type: `"object"`. single nested block, Optional.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_captcha_challenge"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("enable_captcha_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

<a id="canonical-0030230131311301-2020221301310303-3212333322310202-0300223002331230-2032203122001110-3102022311312130-1210021321020111-0000001000313232"></a>

## Direct properties — spec / 010122213221 / 3

- [any_asn](resources--cdn_loadbalancer--reference--group-013.md#canonical-3121222113032223-1300220113211122-3102303011211231-0001101101121000-1212013123030310-1110302210322232-2112200222032232-3210130202023000): complete subsection reference.

- [any_client](resources--cdn_loadbalancer--reference--group-013.md#canonical-1302001220201330-1331232223202101-2220300033110233-1330233330100223-1231113112012330-0032302122023001-1202003201120313-0123231113010300): complete subsection reference.

- [any_ip](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110221022133232-2300032333010202-2101013323001311-3132202333321222-0301330230222000-2122121322011301-1212020100230302-1233203323313101): complete subsection reference.

- [arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100): complete subsection reference.

- [asn_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003233223300201-0312002012103020-1303002203022023-2111233231333112-1000323100130012-1320213001201033-1202333023000302-1013322233111312): complete subsection reference.

- [asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101): complete subsection reference.

- [body_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0020310012101102-3000230032112032-0132110110013312-1020032331321312-1203133212331130-0123320022003020-2120221312120320-3221212002033010): complete subsection reference.

- [client_selector](resources--cdn_loadbalancer--reference--group-013.md#canonical-3230133003213133-3011200302131310-1110312033300333-3311222103011221-1231200100131103-1131223213322203-3302022113011000-2001203101232132): complete subsection reference.

- [cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130): complete subsection reference.

- [disable_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-0032332002211123-1202012303111223-1300130212323212-3131233302210131-2031312311130100-1303223302131323-2102102310301121-2010003203213221): complete subsection reference.

- [domain_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1031330331223002-1133112231311120-0010033010221011-3200232312203013-1321321213331110-3211120301031200-0232002103001122-1233011032121333): complete subsection reference.

- [enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-3103311303303301-3120122203132131-2113220010310230-3210210112101313-1321321213233223-2130133230221223-3223033330123030-1202010323003320): complete subsection reference.

- [enable_javascript_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101321300022020-2012121100121213-1300112332120133-3301000323233111-0133300213322003-1130203021233202-1031230021202010-0002003123331010): complete subsection reference.

<a id="canonical-2122200121320131-3201120023221012-1211131301012120-0013032102322322-2123011123130213-0111223233003021-1323002221223313-1222020032311002"></a>

<a id="canonical-3203013011223231-3322232210113212-2121002202310210-2121320113221202-1201202301022232-1323110103012033-0131102100310220-2232320011131322"></a>

## expiration_timestamp property — spec / 010122213221 / 4

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
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

- [headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020): complete subsection reference.

- [http_method](resources--cdn_loadbalancer--reference--group-014.md#canonical-2312131312221300-0230011100003021-0021131212211331-0123333322230033-1210310032303322-1320033320200030-2221332222001320-2330310311021003): complete subsection reference.

- [ip_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020): complete subsection reference.

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-1012033320010022-2222231020323103-2321000021201310-3031023311331022-0012002023131321-2230011300323100-3130312013222320-2213121223200210): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-014.md#canonical-1323031133100220-3003131300131310-1203112022000311-1230021132323020-2132121211131031-2202311300331212-3111230320020110-3011011131332111): complete subsection reference.

- [query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-3232022003302213-0032131233110110-3231221110233032-1010300113202110-1012221322120230-1322201202221221-2233333330333200-3030322000303331): complete subsection reference.

<a id="canonical-1013223111312210-3301002032002230-0301330331220022-3311223012032002-1332203122232203-3122020111332003-2312331331033100-0002100213122003"></a>

## Next pages — spec / 010122213221 / 5

- [policy_based_challenge.rule_list.rules.spec.any_asn](resources--cdn_loadbalancer--reference--group-013.md#canonical-3121222113032223-1300220113211122-3102303011211231-0001101101121000-1212013123030310-1110302210322232-2112200222032232-3210130202023000)
- [policy_based_challenge.rule_list.rules.spec.any_client](resources--cdn_loadbalancer--reference--group-013.md#canonical-1302001220201330-1331232223202101-2220300033110233-1330233330100223-1231113112012330-0032302122023001-1202003201120313-0123231113010300)
- [policy_based_challenge.rule_list.rules.spec.any_ip](resources--cdn_loadbalancer--reference--group-013.md#canonical-3110221022133232-2300032333010202-2101013323001311-3132202333321222-0301330230222000-2122121322011301-1212020100230302-1233203323313101)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- [policy_based_challenge.rule_list.rules.spec.asn_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003233223300201-0312002012103020-1303002203022023-2111233231333112-1000323100130012-1320213001201033-1202333023000302-1013322233111312)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0020310012101102-3000230032112032-0132110110013312-1020032331321312-1203133212331130-0123320022003020-2120221312120320-3221212002033010)
- [policy_based_challenge.rule_list.rules.spec.client_selector](resources--cdn_loadbalancer--reference--group-013.md#canonical-3230133003213133-3011200302131310-1110312033300333-3311222103011221-1231200100131103-1131223213322203-3302022113011000-2001203101232132)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-0032332002211123-1202012303111223-1300130212323212-3131233302210131-2031312311130100-1303223302131323-2102102310301121-2010003203213221)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1031330331223002-1133112231311120-0010033010221011-3200232312203013-1321321213331110-3211120301031200-0232002103001122-1233011032121333)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-3103311303303301-3120122203132131-2113220010310230-3210210112101313-1321321213233223-2130133230221223-3223033330123030-1202010323003320)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101321300022020-2012121100121213-1300112332120133-3301000323233111-0133300213322003-1130203021233202-1031230021202010-0002003123331010)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--cdn_loadbalancer--reference--group-014.md#canonical-3330200010012232-0013311112232203-1032330210000200-2222122211223030-3322102023003230-0100300103120211-1113231232323333-0222103303312020)
- [policy_based_challenge.rule_list.rules.spec.http_method](resources--cdn_loadbalancer--reference--group-014.md#canonical-2312131312221300-0230011100003021-0021131212211331-0123333322230033-1210310032303322-1320033320200030-2221332222001320-2330310311021003)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-1221203300113022-2132333200200013-1003100021200202-1120013111223113-1232003120333102-0311332302111030-2000323030332330-3022032101132020)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](resources--cdn_loadbalancer--reference--group-014.md#canonical-1012033320010022-2222231020323103-2321000021201310-3031023311331022-0012002023131321-2230011300323100-3130312013222320-2213121223200210)
- [policy_based_challenge.rule_list.rules.spec.path](resources--cdn_loadbalancer--reference--group-014.md#canonical-1323031133100220-3003131300131310-1203112022000311-1230021132323020-2132121211131031-2202311300331212-3111230320020110-3011011131332111)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--cdn_loadbalancer--reference--group-014.md#canonical-1101012003003233-2202222113013213-2002303112232202-1102102121311030-3310003312033111-2122121111000112-0123121121310202-3132330200111100)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-014.md#canonical-3232022003302213-0032131233110110-3231221110233032-1010300113202110-1012221322120230-1322201202221221-2233333330333200-3030322000303331)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3121222113032223-1300220113211122-3102303011211231-0001101101121000-1212013123030310-1110302210322232-2112200222032232-3210130202023000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203211311112133-0113203212030212-0103230203112133-2020231310333001-0010010332300313-0232212023013103-2320302023110323-0021321210333303"></a>

## policy_based_challenge.rule_list.rules.spec.any_asn — any_asn / 133010132030 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.any_asn

<a id="canonical-3112231331323121-3320022221133013-3131221010013300-2211320300230111-2013000321333320-1111132033001000-3201301203311200-2233300001002020"></a>

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
any_asn = {}
```

<a id="canonical-3302123320121301-0300023312001131-3100313112033301-1200231333102200-3020112233310202-3313203112011300-3023110110222233-1220003202320000"></a>

## Direct properties — any_asn / 133010132030 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2313303121303302-3113031202001302-2322303123212103-1231231321313031-0300121102223013-3100110013031302-1220002321130012-3011111020133000"></a>

## Next pages — any_asn / 133010132030 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1302001220201330-1331232223202101-2220300033110233-1330233330100223-1231113112012330-0032302122023001-1202003201120313-0123231113010300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102002331122201-0320100320301121-2010133033310232-2330312030003200-0201220010131321-2302303133103331-3103201201232221-1001220332133320"></a>

## policy_based_challenge.rule_list.rules.spec.any_client — any_client / 020131000322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.any_client

<a id="canonical-0112212111231331-2131010303013302-2030110211100123-3213013230330300-0133311300333000-3013202033201200-3212013031102103-2321213223202212"></a>

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
any_client = {}
```

<a id="canonical-2322030330221100-2223123110132000-0211330222213010-0002100120302211-0310120331123230-2012233023211322-3101000100230333-0223220310110030"></a>

## Direct properties — any_client / 020131000322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0222321331113112-1023020010312311-0022033220331321-2302233301200122-2133221010103203-0033102333202220-1103221112000131-1100232100311321"></a>

## Next pages — any_client / 020131000322 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3110221022133232-2300032333010202-2101013323001311-3132202333321222-0301330230222000-2122121322011301-1212020100230302-1233203323313101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121120202001323-0013202121233310-0113330122002300-2221220031110032-3030300223002110-3133010323003003-3020230320002201-1211010303333102"></a>

## policy_based_challenge.rule_list.rules.spec.any_ip — any_ip / 120231330203 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.any_ip

<a id="canonical-3110230313303220-3032103230312330-2231013003013003-2033332203002113-0310200302220123-0200123232211321-3120331230122210-3000232103113021"></a>

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
any_ip = {}
```

<a id="canonical-1332032100221220-1132130113101111-3130330020130202-3032232032213312-3033110132210333-1211222221312123-2013020102332111-2133200221201210"></a>

## Direct properties — any_ip / 120231330203 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2222120222323001-0033012312111331-2000100120302231-3232313031302220-3313121303022100-3000111322210212-0113030313322221-2030113021122220"></a>

## Next pages — any_ip / 120231330203 / 4

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0301202331323033-0232132122202030-2113200300002120-3301220012303200-0221212202133013-2231032101323000-0233131302233000-3011023123121300"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers — arg_matchers / 031010230013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.arg_matchers

<a id="canonical-1202331333202023-0331330213201312-2032323202013121-0330311313200330-1201333013012020-0333112213110211-1321223333001020-3212001201310000"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all POST args that need to be matched. The criteria for matching each arg are
described in individual instances of ArgMatcherType. The actual arg values are extracted from the
request API as a list of strings for each arg selector name.

Upstream description:

A list of predicates for all POST args that need to be matched. The criteria for matching each arg
are described in individual instances of ArgMatcherType. The actual arg values are extracted from
the request API as a list of strings for each arg selector name. Note that all specified arg matcher
predicates must evaluate to true. A request body greater than 64KB will not be evaluated.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
arg_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0132312122011001-0321011322330233-1101123122101000-1233203213133300-0323023231023101-1010210330033301-3231023231321013-0012111032221322"></a>

## Direct properties — arg_matchers / 031010230013 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1013323232133001-1333023212002100-1022133020223331-2113302220020330-3233023001021000-0022321120211103-1311320012110312-2021332021031131): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1210121133032101-2033110320303231-0021123031202121-1020223313123313-3013232013322003-1121031320130203-0110222312210333-1021213130201333): complete subsection reference.

<a id="canonical-1020030031110113-0012330311023011-3011133220032122-2013323112220222-2213122332113131-1010222333012211-0323300020000013-1030120232331011"></a>

<a id="canonical-3331311120311211-0000231302023223-0002113212211330-0310000203031213-2322131102222130-0232011033133121-3001233203021321-0220231033330023"></a>

## invert_matcher property — arg_matchers / 031010230013 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-013.md#canonical-0130223113313203-1133021030001223-3111100303130003-3130100310030322-1203103132102121-2101230102123033-1022122313200300-2113313231323113): complete subsection reference.

<a id="canonical-2110331012213113-3212213132303313-0201133213000113-2012122203332113-3303130311300232-3111230022130313-2323131212020332-0300222131022312"></a>

<a id="canonical-2331000210331113-1230322002003132-2221332221320211-2102121123323112-0131320013301003-1323003230311102-1303120210022212-2023033200300231"></a>

## name property — arg_matchers / 031010230013 / 5

Type: `"string"`. Optional.

Case-sensitive JSON path in the HTTP request body.

Upstream description:

A case-sensitive JSON path in the HTTP request body.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.json_path": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-0033020310022312-0122323002213003-0230102120013332-1322103333122121-1203002020202131-1012010203133231-2003333001030013-0003332012202011"></a>

## Next pages — arg_matchers / 031010230013 / 6

- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1013323232133001-1333023212002100-1022133020223331-2113302220020330-3233023001021000-0022321120211103-1311320012110312-2021332021031131)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-1210121133032101-2033110320303231-0021123031202121-1020223313123313-3013232013322003-1121031320130203-0110222312210333-1021213130201333)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers.item](resources--cdn_loadbalancer--reference--group-013.md#canonical-0130223113313203-1133021030001223-3111100303130003-3130100310030322-1203103132102121-2101230102123033-1022122313200300-2113313231323113)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1013323232133001-1333023212002100-1022133020223331-2113302220020330-3233023001021000-0022321120211103-1311320012110312-2021332021031131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0030132121032201-2230230013202213-2231210230011102-1221311322102220-0301203011113133-3133322123120220-1330133222211020-1020230032302223"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present — check_not_present / 232323102011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_not_present

<a id="canonical-2013303011030110-2120110113222231-1010332303112313-2222102230111220-0320010011011303-3103100212221020-1001002311020001-3021030333000111"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

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
check_not_present = {}
```

<a id="canonical-3020033123002131-0322113300110002-0201111332133122-0301233220033230-1011121112021002-2322321101323131-2032000010123123-3200333223102330"></a>

## Direct properties — check_not_present / 232323102011 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3210331031101020-1212221230130032-1232330330302323-3300033102201110-2001331313223103-3102000121332302-3311012332013230-3133201020212323"></a>

## Next pages — check_not_present / 232323102011 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1210121133032101-2033110320303231-0021123031202121-1020223313123313-3013232013322003-1121031320130203-0110222312210333-1021213130201333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2232332211120213-1133032201032201-0120013230200200-1113101213220033-0022323133122331-2323123220333312-1323022132201333-0210122133011130"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present — check_present / 003133222131 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.check_present

<a id="canonical-1032021110132311-1323302102033111-0022101301130301-0233312020213013-3323332212002102-1321213200321013-0100110333001202-3101303110120303"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check present.

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
check_present = {}
```

<a id="canonical-0322322200122331-1302220101133021-1021030103032323-3032213011221313-2030303201310132-1222032113200030-0300221211210113-0302221222231221"></a>

## Direct properties — check_present / 003133222131 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1301012213100212-1132221322323000-0013310101212022-2023320322203323-0021230320133020-2301001310320022-0232002311113221-0331013320220012"></a>

## Next pages — check_present / 003133222131 / 4

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0130223113313203-1133021030001223-3111100303130003-3130100310030322-1203103132102121-2101230102123033-1022122313200300-2113313231323113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102033130031202-0010123202112213-0223233303111232-0332322102333201-1133030330330233-0130100022311200-3101332220101311-1011102030230201"></a>

## policy_based_challenge.rule_list.rules.spec.arg_matchers.item — item / 230332212100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- policy_based_challenge.rule_list.rules.spec.arg_matchers.item

<a id="canonical-2321013130100133-3100211211012323-1330213213310232-3211201323310010-2132311001132132-0103313033102300-1033120102222332-0220330032231000"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
item {
  # Configure direct properties listed below.
}
```

<a id="canonical-1113010202011103-3232230013011030-0233123320003101-1200332111321100-0121232020333120-0123020100020213-2131231232103211-2113301122021321"></a>

## Direct properties — item / 230332212100 / 3

<a id="canonical-2213031201121112-1030213023202122-0333301300030100-2003230131113303-2213122021121211-0222330023110232-1000330122122111-1111213030102033"></a>

<a id="canonical-3023311201033203-2222303221010013-3300023222320323-3031020133203130-3312231113321303-2230022221210301-1212101300203221-0000331120010100"></a>

## exact_values property — item / 230332212100 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1322130112200211-2323032300330112-3033013312221023-3013300210003320-1010122112020221-0101101220022100-0102221001132020-2012321323002112"></a>

<a id="canonical-0002032203201131-3321303200211101-0320202032322302-2120213130231020-2113033113200103-0130210201201000-3003220122210103-3003322010310100"></a>

## regex_values property — item / 230332212100 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2111032001223332-1331033102323001-2022331201321132-1331333320310203-1101303103200021-3011030013201012-1023031003331031-1321231132013213"></a>

<a id="canonical-3200111020221302-3100201313013002-0221031310313000-1332223012122213-1010121331220112-0221230311202311-2333201303002303-0212312133313230"></a>

## transformers property — item / 230332212100 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2220233312222202-0133232221102220-1232001132110322-0330222203012311-0123321113122223-2113013223211323-0131333221022111-2313201323001032"></a>

## Next pages — item / 230332212100 / 7

- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--cdn_loadbalancer--reference--group-013.md#canonical-2101111013322230-3322212030112032-0100100312002133-1202202003023232-0302333011020311-3031132111200222-0021111312130021-0303012220033100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0003233223300201-0312002012103020-1303002203022023-2111233231333112-1000323100130012-1320213001201033-1202333023000302-1013322233111312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0202332121121312-3111001003333230-0322333012103222-2122032300011232-1000100313103111-2130033223002331-1322121103321333-0322312231202122"></a>

## policy_based_challenge.rule_list.rules.spec.asn_list — asn_list / 012020320330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.asn_list

<a id="canonical-3203002021210222-0203103302001100-0313033230220021-0232012022120333-3103310130200201-0032102301021332-1030222233032220-2133231210111220"></a>

Type: `"object"`. single nested block, Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("as_numbers")}
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
asn_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-0011030313322312-0000012203113020-0110302312323033-2101203031210000-0023132211030313-2121013230303320-3323020211333112-0213231113110130"></a>

## Direct properties — asn_list / 012020320330 / 3

<a id="canonical-1323232023302012-2300320120020232-2103100331230112-2332223223210323-0032031010232102-2302222333203203-1213303211001121-3120201210113223"></a>

<a id="canonical-1203130020032200-0311202211001100-3121322311010133-2313233323020211-2233011330020101-1121111203233032-3320010001013210-1110101031103123"></a>

## as_numbers property — asn_list / 012020320330 / 4

Type: `["list", "number"]`. Optional.

Unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny lists
for use in network policy or service policy. It can be used to create the allow list only for DNS
Load Balancer.

Upstream description:

An unordered set of RFC 6793 defined 4-byte AS numbers that can be used to create allow or deny
lists for use in network policy or service policy. It can be used to create the allow list only for
DNS Load Balancer.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeBetween(1, 16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "minItems": 1,
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1113321112010300-0121322010000111-2233102002120130-3301322333300333-3223010322230300-2120021230102231-3221331000230132-1331033222112333"></a>

## Next pages — asn_list / 012020320330 / 5

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001020020113120-3003023102231320-0103330023133001-3220333200223211-0310002023120120-1030023112203322-1211013331313330-2021333321121320"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher — asn_matcher / 222303020210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.asn_matcher

<a id="canonical-3231331210032231-2123312121132312-0323213001332133-0011312120300023-2211111233320302-1302103230010132-2313131301331222-0300323203312230"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2113302023123310-1220331132010331-0232311121111010-0322310130301202-3012113120321331-3303220133011322-3301302320221013-0003112212113303"></a>

## Direct properties — asn_matcher / 222303020210 / 3

- [asn_sets](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003030013130330-2331113132100113-3133013032111101-2330111202210322-2212222032222202-2223222120111202-1220132122303010-0212112303222130): complete subsection reference.

<a id="canonical-1303333213201323-3030231012001202-2231210110032123-2332301011003233-2121131102102202-1103301313003100-2221300021031012-3203032320112212"></a>

## Next pages — asn_matcher / 222303020210 / 4

- [policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003030013130330-2331113132100113-3133013032111101-2330111202210322-2212222032222202-2223222120111202-1220132122303010-0212112303222130)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0003030013130330-2331113132100113-3133013032111101-2330111202210322-2212222032222202-2223222120111202-1220132122303010-0212112303222130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310022213331313-2011122223310302-2232122301322101-0103130111023223-3100030032011002-0221301110210303-1202222203302113-2312313033213123"></a>

## policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets — asn_sets / 032303031211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101)
- policy_based_challenge.rule_list.rules.spec.asn_matcher.asn_sets

<a id="canonical-1121012211122000-0233332002121213-0002032011000131-0000200223132111-1320121001123300-1022121111103022-3122320310323001-2230201100313102"></a>

Type: `"object"`. list nested block, Optional.

List of references to bgp\_asn\_set objects.

Upstream description:

A list of references to bgp\_asn\_set objects.

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
    "ves.io.schema.rules.repeated.max_items": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4"
  }
}
```

Terraform syntax:

```terraform
asn_sets {
  # Configure direct properties listed below.
}
```

<a id="canonical-2232230123210323-1020122020203013-2321132200203112-3120031010113113-3020011213123133-0113102213031110-0120313301101011-2011321323312113"></a>

## Direct properties — asn_sets / 032303031211 / 3

<a id="canonical-0130033203003331-3111032011230332-3131330021223203-0033022001021333-3303133212212030-2210032220131022-2230000021121330-2330011121312003"></a>

<a id="canonical-1131312100111313-1113310200110331-3203023310122003-0101133031223013-2201021023000213-0203210210313023-2211331332033013-2220300310013203"></a>

## kind property — asn_sets / 032303031211 / 4

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

<a id="canonical-3111102122210211-2232221210102311-1000121120133211-0003331003211230-1112003133203131-0230320132132231-2222331123201020-0010212113301201"></a>

<a id="canonical-0111213013131020-2313030132312131-3123212232113000-0330223320111030-1230112013032311-2220100201011013-2101311313210020-2131303330111001"></a>

## name property — asn_sets / 032303031211 / 5

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

<a id="canonical-0221022230212122-0212202203221012-0313203031013022-0030232101202313-2113020010113301-3111321100023320-2122203012200201-2233232213322023"></a>

<a id="canonical-0013030203111202-3100000030133300-1013332120100122-0020320322103000-1310223332133222-2223013030002201-2203220301011123-1102321101010130"></a>

## namespace property — asn_sets / 032303031211 / 6

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-3303010112101120-1310121213122202-0112203103000330-0303221120223131-1323322103220303-1220121131123112-2110301101032213-3311311103001202"></a>

<a id="canonical-3020301131133110-1320213011012213-0120111031323123-3221022022331202-1011012223132210-1011130203022303-3122033011203202-0020100012313231"></a>

## tenant property — asn_sets / 032303031211 / 7

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

<a id="canonical-0313111031022023-3133300101321001-1023213302101123-2111230011120001-2003302112210110-2300123113330210-1021110133130230-2220102203231202"></a>

<a id="canonical-0320303233223302-2232132003033021-0303330023210330-3133102023311221-2103233203221030-1321020102031220-1230300320213023-2122101001322030"></a>

## uid property — asn_sets / 032303031211 / 8

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

<a id="canonical-3212030001020222-0302222002222221-1213322200211032-1320221332122312-0102102312202013-1031233201121002-1112302232322330-3322020022100022"></a>

## Next pages — asn_sets / 032303031211 / 9

- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--cdn_loadbalancer--reference--group-013.md#canonical-0003102132023010-3231100012022023-0021011213211001-3321012032332023-1221003122112110-0003331033100121-1032002020102300-1103330212100101)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0020310012101102-3000230032112032-0132110110013312-1020032331321312-1203133212331130-0123320022003020-2120221312120320-3221212002033010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3033012010220110-0102022132332000-1310301203021312-1210120302311233-1212202111032001-0331033203033023-1320213130030313-0022002310300233"></a>

## policy_based_challenge.rule_list.rules.spec.body_matcher — body_matcher / 303312131100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.body_matcher

<a id="canonical-1031030331011012-1102322013022301-0320330033223032-3222002223112121-1110032020123013-2002103323103102-1121113220132031-3212122032220032"></a>

Type: `"object"`. single nested block, Optional.

Matcher specifies multiple criteria for matching an input string. The match is considered successful
if any of the criteria are satisfied. The set of supported match criteria includes a list of exact
values and a list of regular expressions.

Upstream description:

A matcher specifies multiple criteria for matching an input string. The match is considered
successful if any of the criteria are satisfied. The set of supported match criteria includes a list
of exact values and a list of regular expressions.

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
body_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000103103200031-1300113131111032-3311111033012002-2103121320311201-2202303010313300-0132311102233020-0310322211323303-0000312122223212"></a>

## Direct properties — body_matcher / 303312131100 / 3

<a id="canonical-3032223032313111-0003202122300012-1032323312230310-2233032320332000-1013333012011120-2333201101032330-3311302320033122-0301230023200230"></a>

<a id="canonical-1132322002011312-2321133113033113-2320233000020000-1311232010323002-1223332200103200-2110011233111100-0231201112312033-3111010011221232"></a>

## exact_values property — body_matcher / 303312131100 / 4

Type: `["list", "string"]`. Optional.

List of exact values to match the input against.

Upstream description:

A list of exact values to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(64),
}
```

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
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0232300210032121-3313330332331310-2023232113023001-0020103210233310-0021120132021311-2323100233023300-0100001133202210-0011210121023033"></a>

<a id="canonical-2201021120313123-3033010332323332-2102011121131222-3203210032323011-1100131111311003-2230122310122210-0200320033123210-3230102120202212"></a>

## regex_values property — body_matcher / 303312131100 / 5

Type: `["list", "string"]`. Optional.

List of regular expressions to match the input against.

Upstream description:

A list of regular expressions to match the input against.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

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
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "256",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.items.string.regex": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020313011221122-3001121233212111-2123133013232020-2201033312102323-0021003330200133-1221200013132220-2131331333331120-3132322030100123"></a>

<a id="canonical-1030122032013033-1003122231230232-3000111102201132-2300020303033022-1201120332101023-0232022213121033-0002033201311313-1323133203332123"></a>

## transformers property — body_matcher / 303312131100 / 6

Type: `["list", "string"]`. Optional.

\[Enum:
LOWER\_CASE|UPPER\_CASE|BASE64\_DECODE|NORMALIZE\_PATH|REMOVE\_WHITESPACE|URL\_DECODE|TRIM\_LEFT|TRIM\_RIGHT|TRIM\]
Ordered list of transformers (starting from index 0) to be applied to the path before matching.
Possible values are \`LOWER\_CASE\`, \`UPPER\_CASE\`, \`BASE64\_DECODE\`, \`NORMALIZE\_PATH\`,
\`REMOVE\_WHITESPACE\`, \`URL\_DECODE\`, \`TRIM\_LEFT\`, \`TRIM\_RIGHT\`, \`TRIM\`.

Upstream description:

An ordered list of transformers (starting from index 0) to be applied to the path before matching.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(9),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 9,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 9,
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
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "9",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3200212221113313-1203321301332033-0131201202222110-1200313221122332-3132233312023022-2232021133201312-3330213100230100-1201213113221211"></a>

## Next pages — body_matcher / 303312131100 / 7

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3230133003213133-3011200302131310-1110312033300333-3311222103011221-1231200100131103-1131223213322203-3302022113011000-2001203101232132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2331302310233132-3233330030031010-1222331332301310-1121203221002200-1133130221321110-0020301332030311-0221203212121201-2211200130013130"></a>

## policy_based_challenge.rule_list.rules.spec.client_selector — client_selector / 133213332323 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.client_selector

<a id="canonical-2122013012232301-1120001333102030-1112313122033120-2032321322212210-1021113322221223-3311303323320002-2313032033022113-1233321132211320"></a>

Type: `"object"`. single nested block, Optional.

Type can be used to establish a 'selector reference' from one object(called selector) to a set of
other objects(called selectees) based on the value of expressions. A label selector is a label query
over a set of resources. An empty label selector matches all objects.

Upstream description:

This type can be used to establish a 'selector reference' from one object(called selector) to a set
of other objects(called selectees) based on the value of expressions. A label selector is a label
query over a set of resources. An empty label selector matches all objects. A null label selector
matches no objects. Label selector is immutable. Expressions is a list of strings of label selection
expression. Each string has "," separated values which are "AND" and all strings are logically "OR".
BNF for expression string &lt;selector-syntax&gt; ::= &lt;requirement&gt; | &lt;requirement&gt; ","
&lt;selector-syntax&gt; &lt;requirement&gt; ::= \[!\] KEY \[ &lt;set-based-restriction&gt; |
&lt;exact-match-restriction&gt; \] &lt;set-based-restriction&gt; ::= "" |
&lt;inclusion-exclusion&gt; &lt;value-set&gt; &lt;inclusion-exclusion&gt; ::= &lt;inclusion&gt; |
&lt;exclusion&gt; &lt;exclusion&gt; ::= "n&#111;tin" &lt;inclusion&gt; ::= "in" &lt;value-set&gt;
::= "(" &lt;values&gt; ")" &lt;values&gt; ::= VALUE | VALUE "," &lt;values&gt;
&lt;exact-match-restriction&gt; ::= \["="|"=="|"!="\] VALUE.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("expressions")}
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
client_selector {
  # Configure direct properties listed below.
}
```

<a id="canonical-1013222202010121-3302011133133030-0202002113323201-2001310111213002-3303100112031312-2123001220121012-2221133212230011-2021321112021111"></a>

## Direct properties — client_selector / 133213332323 / 3

<a id="canonical-0300021310102112-0233312223213331-0032223231303123-0030332203312132-1021231203021020-0133330113320333-1010033031230032-3232131303122312"></a>

<a id="canonical-1120022223121003-0333202013101023-3200003232311302-3312330212120322-2000223213113103-0022200330331023-2230313213313102-0202220221310221"></a>

## expressions property — client_selector / 133213332323 / 4

Type: `["list", "string"]`. Optional.

Expressions contains the Kubernetes style label expression for selections.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(1),
}
```

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
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.k8s_label_selector": "true",
    "ves.io.schema.rules.repeated.items.string.max_len": "4096",
    "ves.io.schema.rules.repeated.items.string.min_len": "1",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-3310212232123221-1231132010010023-1032233331003200-0213301322233132-2323231110100202-1021202323032300-1013020231222112-1023011130112221"></a>

## Next pages — client_selector / 133213332323 / 5

- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3313200200332201-1200210121121211-3220003300020012-2002110001123113-3233122133221310-0212323322101310-1333012331212122-2002300120333130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331122130032110-0000030232232002-2011201021123203-0000330331011311-2200022321103311-2012031213112333-1012012312033213-2100100310123113"></a>

## policy_based_challenge.rule_list.rules.spec.cookie_matchers — cookie_matchers / 113102222312 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [policy_based_challenge](resources--cdn_loadbalancer--reference--group-013.md#canonical-1032221321010233-1300020333320013-3312111330223002-0332221310112213-0333230120222300-1030231201022210-2100102333211122-1331111021222033)
- [policy_based_challenge.rule_list](resources--cdn_loadbalancer--reference--group-013.md#canonical-3213320011002200-0100122311300210-0321022123302121-3013011333320213-0321000121000203-2133113300312212-1110033001012210-0211320103301301)
- [policy_based_challenge.rule_list.rules](resources--cdn_loadbalancer--reference--group-013.md#canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- policy_based_challenge.rule_list.rules.spec.cookie_matchers

<a id="canonical-3122323331000232-0332310100303121-3202233023302312-1303320110303023-1131232311302033-3332313233323332-2030213320001300-2000102020032112"></a>

Type: `"object"`. list nested block, Optional.

List of predicates for all cookies that need to be matched. The criteria for matching each cookie is
described in individual instances of CookieMatcherType. The actual cookie values are extracted from
the request API as a list of strings for each cookie name.

Upstream description:

A list of predicates for all cookies that need to be matched. The criteria for matching each cookie
is described in individual instances of CookieMatcherType. The actual cookie values are extracted
from the request API as a list of strings for each cookie name. Note that all specified cookie
matcher predicates must evaluate to true.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "check_present"),
  validators.ConflictingListObjectAttributes("check_not_present",
    "item"),
  validators.ConflictingListObjectAttributes("check_present",
    "item")}
```

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
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
cookie_matchers {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230011003230322-3320220120102021-3023002200230202-3032113322230132-1313203223311011-1230200330132022-0131110122003313-0030231010302130"></a>

## Direct properties — cookie_matchers / 113102222312 / 3

- [check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-0200311131313000-1330113300332313-0201331201321232-3321212100022002-0021011312320202-3303313033012313-2131223133333032-3010121011220301): complete subsection reference.

- [check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0120122022030210-3120003232102010-1000000302001001-0102210101030320-3211221101103330-3010111213033233-0103032021303303-2112200021102030): complete subsection reference.

<a id="canonical-1323323221320221-2312213200230130-3332103333212320-2133300320000100-2121000322121322-2130133112231122-2103101021333221-1023303213123320"></a>

<a id="canonical-3020320111302323-1113312212302113-0132113321322033-1122111113003333-0021113313133133-1301131211110223-3113001223231221-3003000101202220"></a>

## invert_matcher property — cookie_matchers / 113102222312 / 4

Type: `"bool"`. Optional.

Invert Matcher. Invert Match of the expression defined.

Upstream description:

Invert Match of the expression defined.

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

- [item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0103003011323300-2200320200120210-1133213110133312-3102111212013223-0302030033013221-3011211222220211-1023030113323131-2133301300232302): complete subsection reference.

<a id="canonical-0121113003130233-0302303020330202-3311322320320211-1021332020200131-0132033110022301-0031031021301223-3010121221323300-2221203130311023"></a>

<a id="canonical-0112320221132300-2200113122330200-3103201030312223-2122020011031010-2303002033100222-0202100132300112-3300302211211001-3332202120033322"></a>

## name property — cookie_matchers / 113102222312 / 5

Type: `"string"`. Optional.

Cookie Name. A case-sensitive cookie name.

Upstream description:

A case-sensitive cookie name.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
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
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

<a id="canonical-3212311113100120-2021323313310031-0310211233303210-0002002001202201-2320110322201300-1320210003202230-1330320122310112-2313301332331010"></a>

## Next pages — cookie_matchers / 113102222312 / 6

- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_not_present](resources--cdn_loadbalancer--reference--group-013.md#canonical-0200311131313000-1330113300332313-0201331201321232-3321212100022002-0021011312320202-3303313033012313-2131223133333032-3010121011220301)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.check_present](resources--cdn_loadbalancer--reference--group-014.md#canonical-0120122022030210-3120003232102010-1000000302001001-0102210101030320-3211221101103330-3010111213033233-0103032021303303-2112200021102030)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers.item](resources--cdn_loadbalancer--reference--group-014.md#canonical-0103003011323300-2200320200120210-1133213110133312-3102111212013223-0302030033013221-3011211222220211-1023030113323131-2133301300232302)
- [policy_based_challenge.rule_list.rules.spec](resources--cdn_loadbalancer--reference--group-013.md#canonical-3001122110332033-1320123200003301-2320310012021312-1230232102320302-0121111212011320-3131020320212123-1000233013302131-2002301101132201)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0200311131313000-1330113300332313-0201331201321232-3321212100022002-0021011312320202-3303313033012313-2131223133333032-3010121011220301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
