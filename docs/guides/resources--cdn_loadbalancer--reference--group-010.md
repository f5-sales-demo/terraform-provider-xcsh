---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0002332021222210-0330133122030022-1013200312101012-1301331321211111-3311201023130032-1130221130013020-1030133311122101-0232320022010221"></a>

## Next pages — cdn_cache_rules / 022030112212 / 7

- [custom_cache_rule](resources--cdn_loadbalancer--reference--group-009.md#canonical-2210332032110200-0210201322313110-1131031023311132-0213211002011103-0233222021320322-2330011210220101-1033313321031003-0323003302010310)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030110033302200-1232222033000023-1010103021312233-0301321022101212-2311100321121333-1211212321213233-0321320032122310-2330321120100131"></a>

## data_guard_rules — data_guard_rules / 002330200320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- data_guard_rules

<a id="canonical-0311211012213113-1202232303221203-2102301000322113-2201211000303031-1032231132111332-1211321322020320-2232020132312002-3203312002003300"></a>

Type: `"object"`. list nested block, Optional.

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*).

Upstream description:

Data Guard prevents responses from exposing sensitive information by masking the data. The system
masks credit card numbers and social security numbers leaked from the application from within the
HTTP response with a string of asterisks (\*). Note: App Firewall should be enabled, to use Data
Guard feature.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("apply_data_guard",
    "skip_data_guard"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value")}
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

Terraform syntax:

```terraform
data_guard_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0300021220322030-0023202310223330-0301130100223320-0313113112021322-2123223103322013-1030201303312013-3112232033012003-2330110013133132"></a>

## Direct properties — data_guard_rules / 002330200320 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-010.md#canonical-0121023211312211-1300022310122122-1033302033212113-1312233223320123-3222213221010323-1032211020303312-1002232303122322-0210323123301122): complete subsection reference.

- [apply_data_guard](resources--cdn_loadbalancer--reference--group-010.md#canonical-1201222133012102-1010031021230212-0331130301020203-3120210132312331-1302200210320233-1103203312313021-2130200213320333-0012112022032021): complete subsection reference.

<a id="canonical-1102201131121112-2112203230120121-0032201300030101-1300022313022312-0331103113213123-1002222002030101-3221320020321231-1311110320203220"></a>

<a id="canonical-3011220220130100-1210102100211331-2211130233031330-0332113030222210-2130112302023013-2123000011301002-3223100100120020-2023231211232231"></a>

## exact_value property — data_guard_rules / 002330200320 / 4

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-1102133033121312-0202123102002120-2322211331213013-1111202032010222-0022322103123001-1001032000303003-3301303000033331-0333033002211030): complete subsection reference.

- [path](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221321201233220-1100220020003312-3000210031200012-2102301330322031-2230030102330302-3313203123032010-2033213021121323-3102030311233212): complete subsection reference.

- [skip_data_guard](resources--cdn_loadbalancer--reference--group-010.md#canonical-1231322011330232-0003030221220313-0113031012123311-2210213321321022-1303302022030212-0202333113330132-3023111031133332-2100202200103322): complete subsection reference.

<a id="canonical-3223101032303022-3323210122000113-1302300112130100-1201231200101222-1122122222012330-3130112311020023-1330023123010120-3211302002202203"></a>

<a id="canonical-0001221301121201-2233321032103230-3323110010032003-3023010131303111-1300123302313333-1121122210000313-3300313202022233-1001321002010320"></a>

## suffix_value property — data_guard_rules / 002330200320 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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

<a id="canonical-3000213313113313-2102210113302202-3022330130222003-3013220320010033-2123222312032332-0023023013312222-1023213231313100-1133112013331230"></a>

## Next pages — data_guard_rules / 002330200320 / 6

- [data_guard_rules.any_domain](resources--cdn_loadbalancer--reference--group-010.md#canonical-0121023211312211-1300022310122122-1033302033212113-1312233223320123-3222213221010323-1032211020303312-1002232303122322-0210323123301122)
- [data_guard_rules.apply_data_guard](resources--cdn_loadbalancer--reference--group-010.md#canonical-1201222133012102-1010031021230212-0331130301020203-3120210132312331-1302200210320233-1103203312313021-2130200213320333-0012112022032021)
- [data_guard_rules.metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-1102133033121312-0202123102002120-2322211331213013-1111202032010222-0022322103123001-1001032000303003-3301303000033331-0333033002211030)
- [data_guard_rules.path](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221321201233220-1100220020003312-3000210031200012-2102301330322031-2230030102330302-3313203123032010-2033213021121323-3102030311233212)
- [data_guard_rules.skip_data_guard](resources--cdn_loadbalancer--reference--group-010.md#canonical-1231322011330232-0003030221220313-0113031012123311-2210213321321022-1303302022030212-0202333113330132-3023111031133332-2100202200103322)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0121023211312211-1300022310122122-1033302033212113-1312233223320123-3222213221010323-1032211020303312-1002232303122322-0210323123301122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3010113032032330-3313322221012210-1301100332313030-3201221121112120-3211130131320233-0103031323303222-0122300301123001-1033320103221312"></a>

## data_guard_rules.any_domain — any_domain / 113033030201 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- data_guard_rules.any_domain

<a id="canonical-3200010310023110-1201123313121311-2231310213012111-1210211101032120-2323121100013123-3021230212130020-1212032112122231-3120201013312232"></a>

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
any_domain = {}
```

<a id="canonical-3220002132122110-0221333312013222-3112231231203032-3302200100301313-3322220101310313-2112022223311113-1113300211311322-1301022312112121"></a>

## Direct properties — any_domain / 113033030201 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1111230332323233-3210101323010120-0312000121000030-0230122031313300-1313031233133113-3313322131113202-1033102323332122-3303313111101132"></a>

## Next pages — any_domain / 113033030201 / 4

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1201222133012102-1010031021230212-0331130301020203-3120210132312331-1302200210320233-1103203312313021-2130200213320333-0012112022032021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011000210332101-0032311201301002-0013312130031202-0030232101130002-2302321002212110-3222111022312122-0102232331303223-3023130031232100"></a>

## data_guard_rules.apply_data_guard — apply_data_guard / 223130330320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- data_guard_rules.apply_data_guard

<a id="canonical-2311303101123302-1312130303200222-2002300223300012-2032323003202113-3010133131230200-0331013113330210-0123022032313203-1123112203332110"></a>

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
apply_data_guard = {}
```

<a id="canonical-0001322312133032-3203310322202203-3031212313131012-3110321233223100-1331221221332321-1333211113100103-0230103020232223-3122232231020331"></a>

## Direct properties — apply_data_guard / 223130330320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1023013031022202-1113122210131133-3212013320010132-0021313203112211-2123231132211131-2033230322021020-0231020221303213-2110122003323122"></a>

## Next pages — apply_data_guard / 223130330320 / 4

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1102133033121312-0202123102002120-2322211331213013-1111202032010222-0022322103123001-1001032000303003-3301303000033331-0333033002211030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2200333210121011-3010030100222302-3112102132122023-0213300101320231-0100132023200201-2312330312303230-1033212200033120-2310300033103030"></a>

## data_guard_rules.metadata — metadata / 301032120000 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- data_guard_rules.metadata

<a id="canonical-2002001200323000-1130312331002203-2111003121002031-1122201231000012-0300100033232302-0323022133032203-2013120303130011-3221310010132011"></a>

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

<a id="canonical-0003003211012323-2321232033222112-1103232232222220-1102030202332121-1012312203323123-2321013022201321-2233112122233300-2112031323122301"></a>

## Direct properties — metadata / 301032120000 / 3

<a id="canonical-0200011022222320-2312033231033323-1111133222011200-1331032033203320-2011300020013202-2213130221111103-1301103130030313-3301021023202130"></a>

<a id="canonical-2303311001100203-2100131020031233-1003121213311010-3200110032301302-0200020032133201-3210022112001230-1120232032200133-2211202012013112"></a>

## description_spec property — metadata / 301032120000 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1322003130033033-3132120222322113-2122001112102320-2202311111311102-3230112323222230-2030323232030220-2231201330230203-0232301100320131"></a>

<a id="canonical-3123113021012012-1221131323032123-0033013311320231-3313121312310212-0323111221201100-2233021033330012-1201300012303330-0313330302100213"></a>

## name property — metadata / 301032120000 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0213313003123312-2010013321310120-1013132301331202-0020010232320230-1223101001003312-0231020011122301-0223000332231212-3202133313102313"></a>

## Next pages — metadata / 301032120000 / 6

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2221321201233220-1100220020003312-3000210031200012-2102301330322031-2230030102330302-3313203123032010-2033213021121323-3102030311233212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3333012010133220-1302032313320321-2232211030133320-0312302302223122-3300131013102013-2200102313300032-0013230132313230-0222322001030032"></a>

## data_guard_rules.path — path / 112001002033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- data_guard_rules.path

<a id="canonical-3112130103012102-2000131210020020-0131022201313101-0123033111233100-2312000301222300-3033000230012012-1333231001313112-3113202211303311"></a>

Type: `"object"`. single nested block, Optional.

Path match of the URI can be either be, Prefix match or exact match or regular expression match.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("path",
    "prefix"),
  validators.ConflictingObjectAttributes("path",
    "regex"),
  validators.ConflictingObjectAttributes("prefix",
    "regex")}
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
  "x-ves-oneof-field-path_match": "[\"path\",\"prefix\",\"regex\"]"
}
```

Terraform syntax:

```terraform
path {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110013112301202-3013220311012020-0310033122310111-3111001333332012-0201231022211210-0031332331210303-1302223103231221-2312312200113020"></a>

## Direct properties — path / 112001002033 / 3

<a id="canonical-2323103103003330-0133300001121212-3312210310220313-0232330103323100-2300211310122203-0123233210112032-3103303103233210-1200023031200100"></a>

<a id="canonical-1120113231321001-0301103213201230-2102232223033020-0232012121332133-3122030331111121-0012213312311303-0101300121311003-3231213232332002"></a>

## path property — path / 112001002033 / 4

Type: `"string"`. Optional.

Exclusive with \[prefix regular expression\] Exact path value to match.

Upstream description:

Exclusive with \[prefix regular expression\] Exact path value to match.

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
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1013020130121132-3331301130031022-0023221102033331-2201310333000002-3221330113322220-3211300301300231-0002121232001000-1121221223310031"></a>

<a id="canonical-0103332130311111-0310102113130131-3031121233232102-1212202203023311-2111010333303002-2321303103200121-2031012210130230-1113120111330112"></a>

## prefix property — path / 112001002033 / 5

Type: `"string"`. Optional.

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

Upstream description:

Exclusive with \[path regular expression\] Path prefix to match (e.g. The value / will match on all paths)

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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="canonical-1320222021200311-2123021320333100-0300311230120122-0101221010231203-3202203013012213-1312011223132301-3021002311203033-0022022233322311"></a>

<a id="canonical-1012320120130221-0201313001101011-1023010221233121-0031221101003202-0230313200120311-3320221031230231-3221220233113323-1321211202301301"></a>

## regular expression property — path / 112001002033 / 6

Type: `"string"`. Optional.

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths).

Upstream description:

Exclusive with \[path prefix\] Regular expression of path match (e.g. The value .\* will match on
all paths)

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
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2331031221002032-1323131121013302-1102331000133303-0201332213313232-1100120211321213-2210033111232033-1133211010300320-0022210021312001"></a>

## Next pages — path / 112001002033 / 7

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1231322011330232-0003030221220313-0113031012123311-2210213321321022-1303302022030212-0202333113330132-3023111031133332-2100202200103322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3112010032030203-2032102202120210-1001013103110313-3211012231121212-1100102032000013-2100330331120121-3123313212303013-3200103101023311"></a>

## data_guard_rules.skip_data_guard — skip_data_guard / 200202021001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- data_guard_rules.skip_data_guard

<a id="canonical-3033302210321312-0022021130220231-1331220211200112-3121311102031320-3022121321133233-1223023322030223-0133132020110123-0012313333021302"></a>

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
skip_data_guard = {}
```

<a id="canonical-2122333011111202-1131321310000030-2113000102122121-3113102331203201-2313022333111101-0011021021130233-2210311331013301-2131021121200210"></a>

## Direct properties — skip_data_guard / 200202021001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1303013331032233-0100121121313021-2132212223013211-1013223011311002-3230100022211313-1331202303000331-1312103033011031-3102111311121332"></a>

## Next pages — skip_data_guard / 200202021001 / 4

- [data_guard_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-0131303033130011-1300120013211032-0221103230201231-3131000301020103-3202102131011030-0111033231131332-0212333123202100-2233232322111033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1031021101032230-2013033021300230-1201020302100221-2102231013033033-0323221002221213-3201112100323310-3122201110132000-2301300001001322"></a>

## ddos_mitigation_rules — ddos_mitigation_rules / 103203211132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- ddos_mitigation_rules

<a id="canonical-1022301231131212-1201332100310002-2100302112312333-0010012123031021-2231110212231120-2201003110213100-2022101221310203-3211322231000202"></a>

Type: `"object"`. list nested block, Optional.

Define manual mitigation rules to block L7 DDoS attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("ddos_client_source",
    "ip_prefix_list")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 256,
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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
ddos_mitigation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2223202232002233-1111321222123000-0021103213313013-1303320320033222-1231112231111000-3112103320002030-1231310033020123-3133302332333323"></a>

## Direct properties — ddos_mitigation_rules / 103203211132 / 3

- [block](resources--cdn_loadbalancer--reference--group-010.md#canonical-1101321101131010-2323112213130013-2002111122301123-3313023030113020-0112333100022223-3023013110021222-3230223301003222-2120211010231133): complete subsection reference.

- [ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320): complete subsection reference.

<a id="canonical-3203231123320310-1212330200213233-0133201022131323-3321220120311330-0223303232331202-0321122311232312-1301232333320013-0221222121330110"></a>

<a id="canonical-0110011030312030-2011222000002321-1211010131301132-0010332021122033-1223033301202112-3033122203233320-1003003210223333-3223013032310000"></a>

## expiration_timestamp property — ddos_mitigation_rules / 103203211132 / 4

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
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.timestamp.within.seconds": "31536000"
  }
}
```

- [ip_prefix_list](resources--cdn_loadbalancer--reference--group-010.md#canonical-1032202033010100-3123021202033002-2031020103231302-3003123132103022-3022131311221333-2033223333113210-1001110012203101-2113133220312302): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-0021200313320233-2232231332010313-3112121020221101-0220223211022102-1131321221322321-0330021023003303-2230313223012220-2031330101232012): complete subsection reference.

<a id="canonical-3010111130101221-0223133102301200-3313011232122112-0113301311022003-3220302120231210-0023123231031101-1012303003103220-1010223121003311"></a>

## Next pages — ddos_mitigation_rules / 103203211132 / 5

- [ddos_mitigation_rules.block](resources--cdn_loadbalancer--reference--group-010.md#canonical-1101321101131010-2323112213130013-2002111122301123-3313023030113020-0112333100022223-3023013110021222-3230223301003222-2120211010231133)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- [ddos_mitigation_rules.ip_prefix_list](resources--cdn_loadbalancer--reference--group-010.md#canonical-1032202033010100-3123021202033002-2031020103231302-3003123132103022-3022131311221333-2033223333113210-1001110012203101-2113133220312302)
- [ddos_mitigation_rules.metadata](resources--cdn_loadbalancer--reference--group-010.md#canonical-0021200313320233-2232231332010313-3112121020221101-0220223211022102-1131321221322321-0330021023003303-2230313223012220-2031330101232012)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1101321101131010-2323112213130013-2002111122301123-3313023030113020-0112333100022223-3023013110021222-3230223301003222-2120211010231133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1102320330013032-3231113212132323-2311323030120002-3131133021131211-2032232333330331-1221100222323131-2323222302211131-3000233033320001"></a>

## ddos_mitigation_rules.block — block / 023231111112 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- ddos_mitigation_rules.block

<a id="canonical-2311301120010032-2003320212201322-1312001220130003-2200302200002211-1113020230222123-3223033320021232-2313331022013211-3030333132003300"></a>

Type: `"object"`. single nested block, Optional.

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
block {}
```

<a id="canonical-0333001322021131-2022011100033311-0311200023312103-0300123322031113-1123313120130300-3002300023301013-1222011332012001-3133321010120000"></a>

## Direct properties — block / 023231111112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3230133120210111-2013102103302212-1231102131210321-1001020211301333-2332331230203323-0331112001003103-1232101030132301-1020020133212321"></a>

## Next pages — block / 023231111112 / 4

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131100311330001-2311210020220323-3100311222012113-0301312111020321-0012121012020312-3010232213221121-3012112213203130-0210212213123332"></a>

## ddos_mitigation_rules.ddos_client_source — ddos_client_source / 023121100011 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- ddos_mitigation_rules.ddos_client_source

<a id="canonical-2011301031310121-0003110120032233-1202302000023120-0320131230112211-0003231222312101-2322302310030310-3322111130200023-1011230322200330"></a>

Type: `"object"`. single nested block, Optional.

DDoS Client Source Choice. DDoS Mitigation sources to be blocked.

Upstream description:

DDoS Mitigation sources to be blocked.

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
ddos_client_source {
  # Configure direct properties listed below.
}
```

<a id="canonical-3021312211331211-2311123332203203-0303101002100130-3013221202020003-0123333303001233-1222132213131033-0100201211322120-1010030133031111"></a>

## Direct properties — ddos_client_source / 023121100011 / 3

- [asn_list](resources--cdn_loadbalancer--reference--group-010.md#canonical-3130032003320203-0303103201031010-0123021130321021-1320113320022211-2103132333323320-3203123223020321-2323020212100123-1112112321213203): complete subsection reference.

<a id="canonical-0101230302230111-0202311101032102-2313110312032301-3321311310232300-1030200331011010-2233200110100202-0030312021100312-0100332303300030"></a>

<a id="canonical-3211330103311123-3230310010303233-1102120022110200-3110210013132123-2323321210103100-2102120220333333-2233223020330001-1003022101211330"></a>

## country_list property — ddos_client_source / 023121100011 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
COUNTRY\_NONE|COUNTRY\_AD|COUNTRY\_AE|COUNTRY\_AF|COUNTRY\_AG|COUNTRY\_AI|COUNTRY\_AL|COUNTRY\_AM|COUNTRY\_AN|COUNTRY\_AO|COUNTRY\_AQ|COUNTRY\_AR|COUNTRY\_AS|COUNTRY\_AT|COUNTRY\_AU|COUNTRY\_AW|COUNTRY\_AX|COUNTRY\_AZ|COUNTRY\_BA|COUNTRY\_BB|COUNTRY\_BD|COUNTRY\_BE|COUNTRY\_BF|COUNTRY\_BG|COUNTRY\_BH|COUNTRY\_BI|COUNTRY\_BJ|COUNTRY\_BL|COUNTRY\_BM|COUNTRY\_BN|COUNTRY\_BO|COUNTRY\_BQ|COUNTRY\_BR|COUNTRY\_BS|COUNTRY\_BT|COUNTRY\_BV|COUNTRY\_BW|COUNTRY\_BY|COUNTRY\_BZ|COUNTRY\_CA|COUNTRY\_CC|COUNTRY\_CD|COUNTRY\_CF|COUNTRY\_CG|COUNTRY\_CH|COUNTRY\_CI|COUNTRY\_CK|COUNTRY\_CL|COUNTRY\_CM|COUNTRY\_CN|COUNTRY\_CO|COUNTRY\_CR|COUNTRY\_CS|COUNTRY\_CU|COUNTRY\_CV|COUNTRY\_CW|COUNTRY\_CX|COUNTRY\_CY|COUNTRY\_CZ|COUNTRY\_DE|COUNTRY\_DJ|COUNTRY\_DK|COUNTRY\_DM|COUNTRY\_DO|COUNTRY\_DZ|COUNTRY\_EC|COUNTRY\_EE|COUNTRY\_EG|COUNTRY\_EH|COUNTRY\_ER|COUNTRY\_ES|COUNTRY\_ET|COUNTRY\_FI|COUNTRY\_FJ|COUNTRY\_FK|COUNTRY\_FM|COUNTRY\_FO|COUNTRY\_FR|COUNTRY\_GA|COUNTRY\_GB|COUNTRY\_GD|COUNTRY\_GE|COUNTRY\_GF|COUNTRY\_GG|COUNTRY\_GH|COUNTRY\_GI|COUNTRY\_GL|COUNTRY\_GM|COUNTRY\_GN|COUNTRY\_GP|COUNTRY\_GQ|COUNTRY\_GR|COUNTRY\_GS|COUNTRY\_GT|COUNTRY\_GU|COUNTRY\_GW|COUNTRY\_GY|COUNTRY\_HK|COUNTRY\_HM|COUNTRY\_HN|COUNTRY\_HR|COUNTRY\_HT|COUNTRY\_HU|COUNTRY\_ID|COUNTRY\_IE|COUNTRY\_IL|COUNTRY\_IM|COUNTRY\_IN|COUNTRY\_IO|COUNTRY\_IQ|COUNTRY\_IR|COUNTRY\_IS|COUNTRY\_IT|COUNTRY\_JE|COUNTRY\_JM|COUNTRY\_JO|COUNTRY\_JP|COUNTRY\_KE|COUNTRY\_KG|COUNTRY\_KH|COUNTRY\_KI|COUNTRY\_KM|COUNTRY\_KN|COUNTRY\_KP|COUNTRY\_KR|COUNTRY\_KW|COUNTRY\_KY|COUNTRY\_KZ|COUNTRY\_LA|COUNTRY\_LB|COUNTRY\_LC|COUNTRY\_LI|COUNTRY\_LK|COUNTRY\_LR|COUNTRY\_LS|COUNTRY\_LT|COUNTRY\_LU|COUNTRY\_LV|COUNTRY\_LY|COUNTRY\_MA|COUNTRY\_MC|COUNTRY\_MD|COUNTRY\_ME|COUNTRY\_MF|COUNTRY\_MG|COUNTRY\_MH|COUNTRY\_MK|COUNTRY\_ML|COUNTRY\_MM|COUNTRY\_MN|COUNTRY\_MO|COUNTRY\_MP|COUNTRY\_MQ|COUNTRY\_MR|COUNTRY\_MS|COUNTRY\_MT|COUNTRY\_MU|COUNTRY\_MV|COUNTRY\_MW|COUNTRY\_MX|COUNTRY\_MY|COUNTRY\_MZ|COUNTRY\_NA|COUNTRY\_NC|COUNTRY\_NE|COUNTRY\_NF|COUNTRY\_NG|COUNTRY\_NI|COUNTRY\_NL|COUNTRY\_NO|COUNTRY\_NP|COUNTRY\_NR|COUNTRY\_NU|COUNTRY\_NZ|COUNTRY\_OM|COUNTRY\_PA|COUNTRY\_PE|COUNTRY\_PF|COUNTRY\_PG|COUNTRY\_PH|COUNTRY\_PK|COUNTRY\_PL|COUNTRY\_PM|COUNTRY\_PN|COUNTRY\_PR|COUNTRY\_PS|COUNTRY\_PT|COUNTRY\_PW|COUNTRY\_PY|COUNTRY\_QA|COUNTRY\_RE|COUNTRY\_RO|COUNTRY\_RS|COUNTRY\_RU|COUNTRY\_RW|COUNTRY\_SA|COUNTRY\_SB|COUNTRY\_SC|COUNTRY\_SD|COUNTRY\_SE|COUNTRY\_SG|COUNTRY\_SH|COUNTRY\_SI|COUNTRY\_SJ|COUNTRY\_SK|COUNTRY\_SL|COUNTRY\_SM|COUNTRY\_SN|COUNTRY\_SO|COUNTRY\_SR|COUNTRY\_SS|COUNTRY\_ST|COUNTRY\_SV|COUNTRY\_SX|COUNTRY\_SY|COUNTRY\_SZ|COUNTRY\_TC|COUNTRY\_TD|COUNTRY\_TF|COUNTRY\_TG|COUNTRY\_TH|COUNTRY\_TJ|COUNTRY\_TK|COUNTRY\_TL|COUNTRY\_TM|COUNTRY\_TN|COUNTRY\_TO|COUNTRY\_TR|COUNTRY\_TT|COUNTRY\_TV|COUNTRY\_TW|COUNTRY\_TZ|COUNTRY\_UA|COUNTRY\_UG|COUNTRY\_UM|COUNTRY\_US|COUNTRY\_UY|COUNTRY\_UZ|COUNTRY\_VA|COUNTRY\_VC|COUNTRY\_VE|COUNTRY\_VG|COUNTRY\_VI|COUNTRY\_VN|COUNTRY\_VU|COUNTRY\_WF|COUNTRY\_WS|COUNTRY\_XK|COUNTRY\_XT|COUNTRY\_YE|COUNTRY\_YT|COUNTRY\_ZA|COUNTRY\_ZM|COUNTRY\_ZW\]
Sources that are located in one of the countries in the given list. Possible values are
\`COUNTRY\_NONE\`, \`COUNTRY\_AD\`, \`COUNTRY\_AE\`, \`COUNTRY\_AF\`, \`COUNTRY\_AG\`,
\`COUNTRY\_AI\`, \`COUNTRY\_AL\`, \`COUNTRY\_AM\`, \`COUNTRY\_AN\`, \`COUNTRY\_AO\`,
\`COUNTRY\_AQ\`, \`COUNTRY\_AR\`, \`COUNTRY\_AS\`, \`COUNTRY\_AT\`, \`COUNTRY\_AU\`,
\`COUNTRY\_AW\`, \`COUNTRY\_AX\`, \`COUNTRY\_AZ\`, \`COUNTRY\_BA\`, \`COUNTRY\_BB\`,
\`COUNTRY\_BD\`, \`COUNTRY\_BE\`, \`COUNTRY\_BF\`, \`COUNTRY\_BG\`, \`COUNTRY\_BH\`,
\`COUNTRY\_BI\`, \`COUNTRY\_BJ\`, \`COUNTRY\_BL\`, \`COUNTRY\_BM\`, \`COUNTRY\_BN\`,
\`COUNTRY\_BO\`, \`COUNTRY\_BQ\`, \`COUNTRY\_BR\`, \`COUNTRY\_BS\`, \`COUNTRY\_BT\`,
\`COUNTRY\_BV\`, \`COUNTRY\_BW\`, \`COUNTRY\_BY\`, \`COUNTRY\_BZ\`, \`COUNTRY\_CA\`,
\`COUNTRY\_CC\`, \`COUNTRY\_CD\`, \`COUNTRY\_CF\`, \`COUNTRY\_CG\`, \`COUNTRY\_CH\`,
\`COUNTRY\_CI\`, \`COUNTRY\_CK\`, \`COUNTRY\_CL\`, \`COUNTRY\_CM\`, \`COUNTRY\_CN\`,
\`COUNTRY\_CO\`, \`COUNTRY\_CR\`, \`COUNTRY\_CS\`, \`COUNTRY\_CU\`, \`COUNTRY\_CV\`,
\`COUNTRY\_CW\`, \`COUNTRY\_CX\`, \`COUNTRY\_CY\`, \`COUNTRY\_CZ\`, \`COUNTRY\_DE\`,
\`COUNTRY\_DJ\`, \`COUNTRY\_DK\`, \`COUNTRY\_DM\`, \`COUNTRY\_DO\`, \`COUNTRY\_DZ\`,
\`COUNTRY\_EC\`, \`COUNTRY\_EE\`, \`COUNTRY\_EG\`, \`COUNTRY\_EH\`, \`COUNTRY\_ER\`,
\`COUNTRY\_ES\`, \`COUNTRY\_ET\`, \`COUNTRY\_FI\`, \`COUNTRY\_FJ\`, \`COUNTRY\_FK\`,
\`COUNTRY\_FM\`, \`COUNTRY\_FO\`, \`COUNTRY\_FR\`, \`COUNTRY\_GA\`, \`COUNTRY\_GB\`,
\`COUNTRY\_GD\`, \`COUNTRY\_GE\`, \`COUNTRY\_GF\`, \`COUNTRY\_GG\`, \`COUNTRY\_GH\`,
\`COUNTRY\_GI\`, \`COUNTRY\_GL\`, \`COUNTRY\_GM\`, \`COUNTRY\_GN\`, \`COUNTRY\_GP\`,
\`COUNTRY\_GQ\`, \`COUNTRY\_GR\`, \`COUNTRY\_GS\`, \`COUNTRY\_GT\`, \`COUNTRY\_GU\`,
\`COUNTRY\_GW\`, \`COUNTRY\_GY\`, \`COUNTRY\_HK\`, \`COUNTRY\_HM\`, \`COUNTRY\_HN\`,
\`COUNTRY\_HR\`, \`COUNTRY\_HT\`, \`COUNTRY\_HU\`, \`COUNTRY\_ID\`, \`COUNTRY\_IE\`,
\`COUNTRY\_IL\`, \`COUNTRY\_IM\`, \`COUNTRY\_IN\`, \`COUNTRY\_IO\`, \`COUNTRY\_IQ\`,
\`COUNTRY\_IR\`, \`COUNTRY\_IS\`, \`COUNTRY\_IT\`, \`COUNTRY\_JE\`, \`COUNTRY\_JM\`,
\`COUNTRY\_JO\`, \`COUNTRY\_JP\`, \`COUNTRY\_KE\`, \`COUNTRY\_KG\`, \`COUNTRY\_KH\`,
\`COUNTRY\_KI\`, \`COUNTRY\_KM\`, \`COUNTRY\_KN\`, \`COUNTRY\_KP\`, \`COUNTRY\_KR\`,
\`COUNTRY\_KW\`, \`COUNTRY\_KY\`, \`COUNTRY\_KZ\`, \`COUNTRY\_LA\`, \`COUNTRY\_LB\`,
\`COUNTRY\_LC\`, \`COUNTRY\_LI\`, \`COUNTRY\_LK\`, \`COUNTRY\_LR\`, \`COUNTRY\_LS\`,
\`COUNTRY\_LT\`, \`COUNTRY\_LU\`, \`COUNTRY\_LV\`, \`COUNTRY\_LY\`, \`COUNTRY\_MA\`,
\`COUNTRY\_MC\`, \`COUNTRY\_MD\`, \`COUNTRY\_ME\`, \`COUNTRY\_MF\`, \`COUNTRY\_MG\`,
\`COUNTRY\_MH\`, \`COUNTRY\_MK\`, \`COUNTRY\_ML\`, \`COUNTRY\_MM\`, \`COUNTRY\_MN\`,
\`COUNTRY\_MO\`, \`COUNTRY\_MP\`, \`COUNTRY\_MQ\`, \`COUNTRY\_MR\`, \`COUNTRY\_MS\`,
\`COUNTRY\_MT\`, \`COUNTRY\_MU\`, \`COUNTRY\_MV\`, \`COUNTRY\_MW\`, \`COUNTRY\_MX\`,
\`COUNTRY\_MY\`, \`COUNTRY\_MZ\`, \`COUNTRY\_NA\`, \`COUNTRY\_NC\`, \`COUNTRY\_NE\`,
\`COUNTRY\_NF\`, \`COUNTRY\_NG\`, \`COUNTRY\_NI\`, \`COUNTRY\_NL\`, \`COUNTRY\_NO\`,
\`COUNTRY\_NP\`, \`COUNTRY\_NR\`, \`COUNTRY\_NU\`, \`COUNTRY\_NZ\`, \`COUNTRY\_OM\`,
\`COUNTRY\_PA\`, \`COUNTRY\_PE\`, \`COUNTRY\_PF\`, \`COUNTRY\_PG\`, \`COUNTRY\_PH\`,
\`COUNTRY\_PK\`, \`COUNTRY\_PL\`, \`COUNTRY\_PM\`, \`COUNTRY\_PN\`, \`COUNTRY\_PR\`,
\`COUNTRY\_PS\`, \`COUNTRY\_PT\`, \`COUNTRY\_PW\`, \`COUNTRY\_PY\`, \`COUNTRY\_QA\`,
\`COUNTRY\_RE\`, \`COUNTRY\_RO\`, \`COUNTRY\_RS\`, \`COUNTRY\_RU\`, \`COUNTRY\_RW\`,
\`COUNTRY\_SA\`, \`COUNTRY\_SB\`, \`COUNTRY\_SC\`, \`COUNTRY\_SD\`, \`COUNTRY\_SE\`,
\`COUNTRY\_SG\`, \`COUNTRY\_SH\`, \`COUNTRY\_SI\`, \`COUNTRY\_SJ\`, \`COUNTRY\_SK\`,
\`COUNTRY\_SL\`, \`COUNTRY\_SM\`, \`COUNTRY\_SN\`, \`COUNTRY\_SO\`, \`COUNTRY\_SR\`,
\`COUNTRY\_SS\`, \`COUNTRY\_ST\`, \`COUNTRY\_SV\`, \`COUNTRY\_SX\`, \`COUNTRY\_SY\`,
\`COUNTRY\_SZ\`, \`COUNTRY\_TC\`, \`COUNTRY\_TD\`, \`COUNTRY\_TF\`, \`COUNTRY\_TG\`,
\`COUNTRY\_TH\`, \`COUNTRY\_TJ\`, \`COUNTRY\_TK\`, \`COUNTRY\_TL\`, \`COUNTRY\_TM\`,
\`COUNTRY\_TN\`, \`COUNTRY\_TO\`, \`COUNTRY\_TR\`, \`COUNTRY\_TT\`, \`COUNTRY\_TV\`,
\`COUNTRY\_TW\`, \`COUNTRY\_TZ\`, \`COUNTRY\_UA\`, \`COUNTRY\_UG\`, \`COUNTRY\_UM\`,
\`COUNTRY\_US\`, \`COUNTRY\_UY\`, \`COUNTRY\_UZ\`, \`COUNTRY\_VA\`, \`COUNTRY\_VC\`,
\`COUNTRY\_VE\`, \`COUNTRY\_VG\`, \`COUNTRY\_VI\`, \`COUNTRY\_VN\`, \`COUNTRY\_VU\`,
\`COUNTRY\_WF\`, \`COUNTRY\_WS\`, \`COUNTRY\_XK\`, \`COUNTRY\_XT\`, \`COUNTRY\_YE\`,
\`COUNTRY\_YT\`, \`COUNTRY\_ZA\`, \`COUNTRY\_ZM\`, \`COUNTRY\_ZW\`. Defaults to \`COUNTRY\_NONE\`.

Upstream description:

Sources that are located in one of the countries in the given list.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [ja4_tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-010.md#canonical-1221200003223120-1123132330223022-0123031300233000-3032123301003023-2232023110333223-0020303131102133-0132010200130211-1321000101133023): complete subsection reference.

- [tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-010.md#canonical-0232131223030222-2002023001132031-0113202320302321-2230333310201010-2203102312111021-1320300323110233-1111222133230010-0013202313112200): complete subsection reference.

<a id="canonical-1131133303300301-0111021301012120-0020203011303132-2033233020212022-1132121320333003-0010233131300332-3033330212111323-1301020003131013"></a>

## Next pages — ddos_client_source / 023121100011 / 5

- [ddos_mitigation_rules.ddos_client_source.asn_list](resources--cdn_loadbalancer--reference--group-010.md#canonical-3130032003320203-0303103201031010-0123021130321021-1320113320022211-2103132333323320-3203123223020321-2323020212100123-1112112321213203)
- [ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-010.md#canonical-1221200003223120-1123132330223022-0123031300233000-3032123301003023-2232023110333223-0020303131102133-0132010200130211-1321000101133023)
- [ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher](resources--cdn_loadbalancer--reference--group-010.md#canonical-0232131223030222-2002023001132031-0113202320302321-2230333310201010-2203102312111021-1320300323110233-1111222133230010-0013202313112200)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3130032003320203-0303103201031010-0123021130321021-1320113320022211-2103132333323320-3203123223020321-2323020212100123-1112112321213203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020023331233203-2331031010030220-2030101132200111-1033210110121221-2211330221032022-2113312133332231-3223021322232102-3113310002212112"></a>

## ddos_mitigation_rules.ddos_client_source.asn_list — asn_list / 011222130103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- ddos_mitigation_rules.ddos_client_source.asn_list

<a id="canonical-2131001030302230-3033011120331123-2302011200302313-0020322221023020-0013321313132120-3233101001313200-1022300022100300-2023303233112101"></a>

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

<a id="canonical-2031011032112032-1011220011021120-0121120231212123-0030322213221203-1121301321330033-1001333200330012-2112011033013221-0330002322010233"></a>

## Direct properties — asn_list / 011222130103 / 3

<a id="canonical-1001312112023123-2321102120013033-2013322322012320-1211222121103033-2122122120121103-1113131012130231-2021003000302002-2023231130230211"></a>

<a id="canonical-2302321033101332-3113202330003010-2333321110312001-3333100300330133-1003203230213200-1232112303103322-2121301231031101-2302012233212020"></a>

## as_numbers property — asn_list / 011222130103 / 4

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

<a id="canonical-0320020101133223-2213220303213103-0011230312330211-3203233112201313-3323022111220102-3103002302210113-2023122200223110-1033030012121031"></a>

## Next pages — asn_list / 011222130103 / 5

- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1221200003223120-1123132330223022-0123031300233000-3032123301003023-2232023110333223-0020303131102133-0132010200130211-1321000101133023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132132300131323-0102112031323201-3112101130201031-2300020113220303-0300323323302002-0012200113321021-3222013213203123-3003033032022323"></a>

## ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher — ja4_tls_fingerprint_matcher / 201100223013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- ddos_mitigation_rules.ddos_client_source.ja4_tls_fingerprint_matcher

<a id="canonical-2330112201213102-2332311011222131-2220222303022201-3003332312233311-2020202022132032-0223330120232232-3201001203012133-1200001130130131"></a>

Type: `"object"`. single nested block, Optional.

Extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

Upstream description:

An extended version of JA3 that includes additional fields for more comprehensive fingerprinting of
SSL/TLS clients and potentially has a different structure and length.

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
ja4_tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-0311322313221021-3203032313321221-1212123102323223-2130333331111112-3301210232130303-1322103231112313-1103001312101103-3321321002320232"></a>

## Direct properties — ja4_tls_fingerprint_matcher / 201100223013 / 3

<a id="canonical-2030213330321333-1220012102102013-0202312332323302-1021322230100030-2100010211013122-0212122123330333-3331313121331102-0023320101311330"></a>

<a id="canonical-3123102111122332-0123321310132023-3311013301302303-2310332131231323-1002131032220321-3323133010220313-3123301310331013-0311102213300122"></a>

## exact_values property — ja4_tls_fingerprint_matcher / 201100223013 / 4

Type: `["list", "string"]`. Optional.

List of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Upstream description:

A list of exact JA4 TLS fingerprint to match the input JA4 TLS fingerprint against.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "36",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2220312000131031-1102123032313220-0202211123012223-1212013312130002-1001312313023313-0121220112001310-2321200313203300-0312210022121320"></a>

## Next pages — ja4_tls_fingerprint_matcher / 201100223013 / 5

- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0232131223030222-2002023001132031-0113202320302321-2230333310201010-2203102312111021-1320300323110233-1111222133230010-0013202313112200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2121020020013033-2321200221221111-0211003313202302-1212300313100030-3003330210121011-2103211102011310-1323310202121312-0033221111321100"></a>

## ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher — tls_fingerprint_matcher / 211303102020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- ddos_mitigation_rules.ddos_client_source.tls_fingerprint_matcher

<a id="canonical-1102221121013011-3220002311310113-2310023111232010-1023221121322301-1222121112210132-3010110220310101-2022303013232110-3111220010113000"></a>

Type: `"object"`. single nested block, Optional.

TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are
satisfied..

Upstream description:

A TLS fingerprint matcher specifies multiple criteria for matching a TLS fingerprint. The set of
supported positive match criteria includes a list of known classes of TLS fingerprints and a list of
exact values. The match is considered successful if either of these positive criteria are satisfied
and the input fingerprint is not one of the excluded values.

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
tls_fingerprint_matcher {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200333023012102-1112330003001022-1330122013113322-3033302232031232-3020131321312123-0013313120320033-3310123332333110-1032122201303113"></a>

## Direct properties — tls_fingerprint_matcher / 211303102020 / 3

<a id="canonical-0010010101301122-3312122120133102-1333210011002112-2133310322133130-1300201103030132-2323013033012310-0113013023330010-1012023121011130"></a>

<a id="canonical-3113031220100012-3112213320121332-0300310300310112-1213323231003222-0012330232201212-2333333011111220-1023012333323101-2002013031002101"></a>

## classes property — tls_fingerprint_matcher / 211303102020 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
TLS\_FINGERPRINT\_NONE|ANY\_MALICIOUS\_FINGERPRINT|ADWARE|ADWIND|DRIDEX|GOOTKIT|GOZI|JBIFROST|QUAKBOT|RANSOMWARE|TROLDESH|TOFSEE|TORRENTLOCKER|TRICKBOT\]
List of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against. Possible
values are \`TLS\_FINGERPRINT\_NONE\`, \`ANY\_MALICIOUS\_FINGERPRINT\`, \`ADWARE\`, \`ADWIND\`,
\`DRIDEX\`, \`GOOTKIT\`, \`GOZI\`, \`JBIFROST\`, \`QUAKBOT\`, \`RANSOMWARE\`, \`TROLDESH\`,
\`TOFSEE\`, \`TORRENTLOCKER\`, \`TRICKBOT\`. Defaults to \`TLS\_FINGERPRINT\_NONE\`.

Upstream description:

A list of known classes of TLS fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2112021223101233-3033130211323103-2321332103123203-2102111211301002-1202202211210311-1123131131201101-1113121122123003-0231112331120033"></a>

<a id="canonical-3202110220111001-1031213132332330-3230033201202211-1220111020310321-1221021210202021-2020301221230323-2221030111300232-3002031310133123"></a>

## exact_values property — tls_fingerprint_matcher / 211303102020 / 5

Type: `["list", "string"]`. Optional.

List of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Upstream description:

A list of exact TLS JA3 fingerprints to match the input TLS JA3 fingerprint against.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0212333233221133-0013123000030001-0100321001012303-2130030231223232-0303312122221122-2131300021303132-0200202302212012-0333333131300101"></a>

<a id="canonical-1032213222121232-3123323233030012-3000112032211311-2212301232102321-0200212111033013-0122133333111030-3201032201011123-1002033000322222"></a>

## excluded_values property — tls_fingerprint_matcher / 211303102020 / 6

Type: `["list", "string"]`. Optional.

List of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can be
used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Upstream description:

A list of TLS JA3 fingerprints to be excluded when matching the input TLS JA3 fingerprint. This can
be used to skip known false positives when using one or more known TLS fingerprint classes in the
enclosing matcher.

Provider validators and defaults (from schema source):

```go
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
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.len": "32",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3222233031332020-3112110221332101-1202002021201001-1212030022200013-1113113123033130-2231112101202310-0310232023112112-3232323323303031"></a>

## Next pages — tls_fingerprint_matcher / 211303102020 / 7

- [ddos_mitigation_rules.ddos_client_source](resources--cdn_loadbalancer--reference--group-010.md#canonical-2010221100102311-2301100021021321-1132123322110312-2121020233221223-0112321033312122-1333222020130300-0333210023203323-0010022110222320)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1032202033010100-3123021202033002-2031020103231302-3003123132103022-3022131311221333-2033223333113210-1001110012203101-2113133220312302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1100202021300302-3032322313031332-1101123203202120-1110201103110011-2233310303301313-3000332122002002-1200301031110003-1233113311100020"></a>

## ddos_mitigation_rules.ip_prefix_list — ip_prefix_list / 321302011113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- ddos_mitigation_rules.ip_prefix_list

<a id="canonical-1033120112120123-3030222311123121-3110333321012100-1311212100201231-1302312303232100-1300222123331023-0103100311321223-2232111020312231"></a>

Type: `"object"`. single nested block, Optional.

List of IP Prefix strings to match against.

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
ip_prefix_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323210203102320-2223203100211202-0323203330313002-3000200211030330-3323123121322233-2221332300033220-3211132102333333-0113120333012200"></a>

## Direct properties — ip_prefix_list / 321302011113 / 3

<a id="canonical-0223213210203110-0203003323321110-1323210201300032-1330101122031233-0002033211302023-3112102232301300-1232130220131123-2112220112101322"></a>

<a id="canonical-3132201312010221-3112002111101121-2000203310121322-2100310201333333-3110223312231131-2212302032302032-2330212120102211-0033021001213103"></a>

## invert_match property — ip_prefix_list / 321302011113 / 4

Type: `"bool"`. Optional.

Invert Match Result. Invert the match result.

Upstream description:

Invert the match result.

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

<a id="canonical-1103211303322330-3012331002310301-1013320011031010-2102123213122321-0301001031332201-3333212202111330-0121101112112322-2311123101230333"></a>

<a id="canonical-2031100331102311-2013320002312330-0020003201021322-3203133011031010-1302232112201222-3031111211300202-1202230312102012-0033321121021322"></a>

## ip_prefixes property — ip_prefix_list / 321302011113 / 5

Type: `["list", "string"]`. Optional.

IPv4 Prefix List. List of IPv4 prefix strings.

Upstream description:

List of IPv4 prefix strings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(128),
}
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.items.string.not_empty": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1031230020332332-2030120330200212-2032013221233133-0211232202330232-1220120220223132-1201032012131201-0021021033222101-2212301002030001"></a>

## Next pages — ip_prefix_list / 321302011113 / 6

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0021200313320233-2232231332010313-3112121020221101-0220223211022102-1131321221322321-0330021023003303-2230313223012220-2031330101232012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302231012230103-3122210330013310-0333000330023133-0100101312333103-0212212133223201-0002122110003003-0111123323201223-3220201130103321"></a>

## ddos_mitigation_rules.metadata — metadata / 312033301133 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- ddos_mitigation_rules.metadata

<a id="canonical-3130323213113131-2301230133203230-3303200020112113-2320123223233023-0201312302132310-3323120331212110-3021322330303313-0012132123012301"></a>

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

<a id="canonical-0002223133331020-1111101030312102-3001302233203130-1112320321113333-3233211232102121-0332103132211300-0103223212211220-1331112220012130"></a>

## Direct properties — metadata / 312033301133 / 3

<a id="canonical-3011222013131130-1231031301131211-3211200033031200-2003112302112023-0021332101100121-0222010233212201-2002130121313210-2002211230122320"></a>

<a id="canonical-3203102221311313-1203220222023122-1002232222323013-0313121100223022-0220131102032102-3031201321112230-1002303013200000-0313230223310122"></a>

## description_spec property — metadata / 312033301133 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2121101120131131-2230021113221212-2202022012112302-0011102023302320-1000113331332330-3001221003220301-3313212331020102-3320323211033321"></a>

<a id="canonical-2301132211120131-0110332312123213-3310003310230233-3311030200331312-1332200211212130-2221223133303001-3030202212030101-1113330010123233"></a>

## name property — metadata / 312033301133 / 5

Type: `"string"`. Optional.

Name of the message. The value of name has to follow DNS-1035 format.

Upstream description:

This is the name of the message. The value of name has to follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2122322300033133-2230200220121120-2113332302310030-0313220332121312-3230003111321311-1320321013231020-0301232001232030-0212130230300113"></a>

## Next pages — metadata / 312033301133 / 6

- [ddos_mitigation_rules](resources--cdn_loadbalancer--reference--group-010.md#canonical-1113101222111100-0231333020320020-2111302302220332-3302220301331200-2213011011211300-1222011002133311-2332301203001000-3323001221330013)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2030222233111300-2331213302201233-3321013002010200-3033210131133233-0110330202020130-0312030130132303-0010221020232332-1100111013012103"></a>

## default_cache_action — default_cache_action / 201310222302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- default_cache_action

<a id="canonical-2033212333122032-2312110112111001-1233312213130130-1103312010323210-1023100121203202-0231332123122302-1303001230110013-2012231230210121"></a>

Type: `"object"`. single nested block, Optional.

Default Cache Behaviour. This defines a Default Cache Action.

Upstream description:

This defines a Default Cache Action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_default"),
  validators.ConflictingObjectAttributes("cache_disabled",
    "cache_ttl_override"),
  validators.ConflictingObjectAttributes("cache_ttl_default",
    "cache_ttl_override")}
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
  "x-ves-oneof-field-cache_actions": "[\"cache_disabled\",\"cache_ttl_default\",\"cache_ttl_override\"]"
}
```

Terraform syntax:

```terraform
default_cache_action {
  # Configure direct properties listed below.
}
```

<a id="canonical-3100001001201001-2332020232003031-1012332221123323-2032021332233001-2322021202002300-0112020110121001-0112020001301300-0221010000123303"></a>

## Direct properties — default_cache_action / 201310222302 / 3

- [cache_disabled](resources--cdn_loadbalancer--reference--group-010.md#canonical-1112300232302300-1311300203022231-0122331013122023-3331202321321013-0223203021300303-2330123312133031-2330013031221033-2100200031312010): complete subsection reference.

<a id="canonical-0113101211112232-3111112303311302-2311330203121223-3133121201110322-1100320032022313-3002032121031100-3201330222303223-3211133132013331"></a>

<a id="canonical-1300103310231322-3021321313000110-2203300220301301-2200030112122100-2322212010001102-1133113001130012-1312103021311013-1321311302330312"></a>

## cache_ttl_default property — default_cache_action / 201310222302 / 4

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_override\] Use Cache TTL Provided by Origin, and set a
contigency TTL value in case one is not provided.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-3201122112033133-2012200101300303-1232133123113222-0011121213202322-1320123013120223-3201003001233002-2333312131003320-1012130032010230"></a>

<a id="canonical-2232331322101231-0303132203231030-0232102123113330-0302023132023013-0320012312012310-2320010001220310-1311022210000332-3300113202122120"></a>

## cache_ttl_override property — default_cache_action / 201310222302 / 5

Type: `"string"`. Optional.

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

Upstream description:

Exclusive with \[cache\_disabled cache\_ttl\_default\] Always override the Cache TTL provided by
Origin.

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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.time_interval": "true"
  }
}
```

<a id="canonical-1213031010332132-2301232013111330-3010220001301101-2100200321101203-1330021333111130-2221221332211222-1223103210212011-0130030020032301"></a>

## Next pages — default_cache_action / 201310222302 / 6

- [default_cache_action.cache_disabled](resources--cdn_loadbalancer--reference--group-010.md#canonical-1112300232302300-1311300203022231-0122331013122023-3331202321321013-0223203021300303-2330123312133031-2330013031221033-2100200031312010)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1112300232302300-1311300203022231-0122331013122023-3331202321321013-0223203021300303-2330123312133031-2330013031221033-2100200031312010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220210132313210-1232021331031112-0323223010112023-2300023302030020-2311021020113221-2010122023220123-1202021200300221-1233202012311311"></a>

## default_cache_action.cache_disabled — cache_disabled / 002121031232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [default_cache_action](resources--cdn_loadbalancer--reference--group-010.md#canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313)
- default_cache_action.cache_disabled

<a id="canonical-3211233323103013-3220301332022022-2330331301333122-2012220120322331-3231022332000232-2123011210000013-2122232222332200-3021201022101133"></a>

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
cache_disabled = {}
```

<a id="canonical-0213322133113302-3220322122222112-0123110231120011-1002322232132203-0133002221202222-2230012311132111-2333123200311033-1001131210130032"></a>

## Direct properties — cache_disabled / 002121031232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0332320320132131-0221331221220201-0021111211312303-0011321303122211-2021230101313213-1020213302313320-1233200120331031-2122111221301300"></a>

## Next pages — cache_disabled / 002121031232 / 4

- [default_cache_action](resources--cdn_loadbalancer--reference--group-010.md#canonical-2300021122222120-3103201200301030-2101110210210123-3021222103220303-1221111131320223-2202110300233100-2330133310101313-0202223013312313)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2310301221200230-0201112131201231-1131101200303021-2130030310111230-3203022302002332-2030103200001230-3233113321131031-0233232233132333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1133213101032010-0032130323010132-2111230312321332-2322131132011312-0031122203212303-0103000223003302-1133310303231003-1201111130131311"></a>

## default_sensitive_data_policy — default_sensitive_data_policy / 033320333313 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- default_sensitive_data_policy

<a id="canonical-0232201112201301-3232122211222210-1201200012003200-2231121011122220-1210020311103232-1111010111310203-0231021103002332-1301120220213033"></a>

Type: `["object", {}]`. Optional.

\[OneOf: default\_sensitive\_data\_policy, sensitive\_data\_policy; Default:
default\_sensitive\_data\_policy\] Policy configuration for this feature.

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

OneOf alternatives in this subsection:

- [default_sensitive_data_policy](resources--cdn_loadbalancer--reference--group-010.md#canonical-0232201112201301-3232122211222210-1201200012003200-2231121011122220-1210020311103232-1111010111310203-0231021103002332-1301120220213033)
- [sensitive_data_policy](resources--cdn_loadbalancer--reference--group-014.md#canonical-1300020121333330-1303221303112300-0023322213301102-1002303223322122-2130031202131001-1332232121231102-1300111112133313-2022220010223230)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
default_sensitive_data_policy = {}
```

<a id="canonical-3303022101301113-1022213010002121-0033321100002113-3012200202000020-3031103333211212-0322233031022123-0103311203311010-1213131220103013"></a>

## Direct properties — default_sensitive_data_policy / 033320333313 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200333303131232-0233013203003230-3201221202111231-2211311222211231-0232312223310031-3010313021102333-3003101021000032-3103223320310331"></a>

## Next pages — default_sensitive_data_policy / 033320333313 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3312231102230133-0002321102020331-1031013220300023-0101221003122222-2100330313000330-2331331200001221-0133311233301310-3133210321022010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312221312110112-3321102203300313-2300121323303331-3102322013113332-1020321010313111-3311220233302000-1301211300231000-2130010213033223"></a>

## disable_api_definition — disable_api_definition / 320301321112 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_api_definition

<a id="canonical-1013231301331321-3201321033031021-0303223333210301-2112110233330302-0302100120123301-0132331033332233-0312121102001302-1222030220330231"></a>

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
disable_api_definition = {}
```

<a id="canonical-3000133121221020-2322102303123031-2210121121303112-1212131013012230-2102131320020132-3111100330202011-1132233210031302-2211102013132200"></a>

## Direct properties — disable_api_definition / 320301321112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2331330032003320-1121123200220232-0322303020302333-2112320100202222-2233130223103023-1012220331120112-1102221230311103-2211322001120023"></a>

## Next pages — disable_api_definition / 320301321112 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2130210132113322-0123113132000212-3230220223101021-1200003113211031-0113012202110012-1311330331231000-1333100002101020-1131001031331003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210311330300112-3213011121232310-3013101223130223-0310013021200211-2111330110002322-1202120311203233-1300303303010301-3123020311301030"></a>

## disable_api_discovery — disable_api_discovery / 113232131120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_api_discovery

<a id="canonical-1322032121300031-1321200030022101-0230311311232112-1223332321131322-3320303002032002-2300111110032023-0301130102003102-3121221211001211"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_api\_discovery, enable\_api\_discovery; Default: disable\_api\_discovery\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-1322032121300031-1321200030022101-0230311311232112-1223332321131322-3320303002032002-2300111110032023-0301130102003102-3121221211001211)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-0001201201333322-1133021323220003-2333030001222131-0223013300101113-2133221013222213-2220202223030213-2303012213101101-0010230023302110)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_api_discovery = {}
```

<a id="canonical-0020220002301211-0011131100013203-3332123011030213-1110223010333121-3233002023210303-1131110100122120-3330000100332220-2223230303132020"></a>

## Direct properties — disable_api_discovery / 113232131120 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032132213210232-1112332331013311-3100200221201112-1023232202303232-2233232312103130-0032031132221032-1203211213302322-0031002002113003"></a>

## Next pages — disable_api_discovery / 113232131120 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0322312112202232-3232301311330110-0300133101202313-0011202213000301-3113120123102301-0132223030100120-2202020201132330-2132133022103120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2201133112102323-2320223232223312-3312013030002213-0221202323302231-0003121012110331-1303003012130003-2033321312312031-2212020103300233"></a>

## disable_client_side_defense — disable_client_side_defense / 003220132301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_client_side_defense

<a id="canonical-2323022203312322-3012211311100012-3031300132031113-3310223011233211-0120103103100233-1101032030002322-1221023301030013-2002330013332333"></a>

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
disable_client_side_defense = {}
```

<a id="canonical-0131010013122221-3332211130222031-2310131210002121-1330303022201131-2033320111220030-0222222032323302-3101212123133102-3122010311000011"></a>

## Direct properties — disable_client_side_defense / 003220132301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300033113003331-3312321033122012-0013303321221101-2113322132303113-1310131100002100-2213230101022021-0310201002011222-0132333001012131"></a>

## Next pages — disable_client_side_defense / 003220132301 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1031001202212121-2320213202323013-2310203020210132-3222031002011331-1231231310110332-2313123223113001-3022212102330213-3000321230000321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301320111201231-1333033212113331-0200100302021123-0333302031013022-3010300002322102-2231321033331012-0320212223302232-2022202312200210"></a>

## disable_ip_reputation — disable_ip_reputation / 333013112223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_ip_reputation

<a id="canonical-0232112311102100-0012331211221212-1001112022113233-1022321313211112-0223113010131232-2221332211301003-1223003001103130-3211023200221123"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_ip\_reputation, enable\_ip\_reputation; Default: disable\_ip\_reputation\] Enable
this option

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

OneOf alternatives in this subsection:

- [disable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-0232112311102100-0012331211221212-1001112022113233-1022321313211112-0223113010131232-2221332211301003-1223003001103130-3211023200221123)
- [enable_ip_reputation](resources--cdn_loadbalancer--reference--group-010.md#canonical-1123223010010320-2010330211102303-3003132133301111-3013321110233132-3220003300022030-2113002030033003-3221112313032202-3320311130101320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_ip_reputation = {}
```

<a id="canonical-0223033322313301-1022312022323020-0113130302113111-1231102113032002-1122200132122232-2021113301223133-1003303321133313-2032121202211100"></a>

## Direct properties — disable_ip_reputation / 333013112223 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2302333103202023-2231301231032210-3121213311220122-0013221320012112-3300202020100022-3102123032012303-2221121003203020-0221302032133330"></a>

## Next pages — disable_ip_reputation / 333013112223 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3323303332010223-3130303213321301-1233002121331100-0110203021132310-0033302313301323-2313122210131232-3213230202131103-1322110213131001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2123113123220311-3203322331112213-1333323103200133-1021120001213011-0013003233220200-0120321123003203-0210311121201021-3313221200030031"></a>

## disable_malicious_user_detection — disable_malicious_user_detection / 000323222003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_malicious_user_detection

<a id="canonical-2333103103022200-0210221203102002-1120233210310133-0232012330110113-1331222330323203-0033030202200212-1220213202003111-1013012231033112"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_malicious\_user\_detection, enable\_malicious\_user\_detection; Default:
disable\_malicious\_user\_detection\] Configuration parameter for disable malicious user detection.

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

OneOf alternatives in this subsection:

- [disable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-2333103103022200-0210221203102002-1120233210310133-0232012330110113-1331222330323203-0033030202200212-1220213202003111-1013012231033112)
- [enable_malicious_user_detection](resources--cdn_loadbalancer--reference--group-010.md#canonical-1123232203103010-2221320123213112-1303032131321013-2032233321303102-1212120310102022-0332230300221210-2020311101222333-2212332330213300)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_malicious_user_detection = {}
```

<a id="canonical-0223122133031313-2213130110313302-3110000200331203-3332230120330303-1322113011130132-0220012103230333-0200202320233021-2322221313012321"></a>

## Direct properties — disable_malicious_user_detection / 000323222003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122002331210121-3110120230311333-2313011111220030-0313333122213112-1213120321321001-0300220100232023-2002313211230013-2101320310233333"></a>

## Next pages — disable_malicious_user_detection / 000323222003 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2101211221210102-0323002031331010-2210220231120012-0023021303231032-1032303023320223-0003233332001012-3103103110223321-1232203302313021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0211023331023323-1212330103311110-0231332000122100-0311300110312210-1203232020312023-3333030303330213-1131011033113230-0031312232011311"></a>

## disable_rate_limit — disable_rate_limit / 123230320032 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_rate_limit

<a id="canonical-3133211010101322-0323301002321331-2233011002232323-1123100210322211-1330323131213103-3011211112311233-3123030230331311-1230310111020221"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable rate limit.

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
disable_rate_limit = {}
```

<a id="canonical-0330220020232110-1321232311100012-1201221132100223-3013233123220212-2013100203103303-2132122031331033-3210210103123031-1020100030103333"></a>

## Direct properties — disable_rate_limit / 123230320032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0323203101101111-1123032023000230-2210002202213013-1311021320113020-0233233331120321-2203121100230021-1232200130013022-0132120302103033"></a>

## Next pages — disable_rate_limit / 123230320032 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0300001220010210-2011013010302113-1031230110020223-2000000323133201-2331022100022333-0022210002101020-3232010103322200-0011200213203231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213133112330120-1112333011310321-2332210333112212-0321021113221000-1220101223021332-3133132300122223-1232310102220101-3131320130310023"></a>

## disable_threat_mesh — disable_threat_mesh / 232230313013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_threat_mesh

<a id="canonical-3212121021213202-2312330010331332-1330021101321323-3001103211201202-0231201210232013-2112333212011003-2330002011233301-1323122311231320"></a>

Type: `["object", {}]`. Optional.

\[OneOf: disable\_threat\_mesh, enable\_threat\_mesh; Default: disable\_threat\_mesh\] Enable this
option

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

OneOf alternatives in this subsection:

- [disable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-3212121021213202-2312330010331332-1330021101321323-3001103211201202-0231201210232013-2112333212011003-2330002011233301-1323122311231320)
- [enable_threat_mesh](resources--cdn_loadbalancer--reference--group-010.md#canonical-2123213233303030-1130113101312333-0212210331021330-3103233332320113-1210122212022213-2121031100211332-1031010102222213-2210032031223013)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
disable_threat_mesh = {}
```

<a id="canonical-1222223211213011-1032112021021112-0332011201103222-2303110303223233-2133111000221130-2030023133321123-2023302032030321-0201330022210210"></a>

## Direct properties — disable_threat_mesh / 232230313013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130023110112331-0111121001202200-3133200303203213-1200122122233113-3033323321210032-0110321012013313-1123023112221231-0110211013012133"></a>

## Next pages — disable_threat_mesh / 232230313013 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1120023301221013-0010101221332312-2101301230211303-3223121323013122-2033123211312311-2201011122202012-0103112113322002-3003131323102011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0322323120321301-1030131221120223-1233311301001232-3220013011221132-3032000223332131-0321120100023111-1131023102011212-1213220001021121"></a>

## disable_waf — disable_waf / 032030010031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- disable_waf

<a id="canonical-0330321120112012-0021333212231113-2232020002320311-0200223111331021-1210113213102002-2101003211223303-0223321332020310-1010211300000211"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable waf.

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
disable_waf = {}
```

<a id="canonical-2331013030010311-1220200221121320-1220021232332311-2112131002313232-1131331130121302-2213321132132323-1120012000020221-2332102010202322"></a>

## Direct properties — disable_waf / 032030010031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333202110022300-0113133033102322-2101222232213223-3013010302133013-2112123000311213-0201113220213023-3303302122012101-3210211332203001"></a>

## Next pages — disable_waf / 032030010031 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2032100310032213-1102211120033313-0001323013010022-1302112103220302-1022012200220002-3233212132031231-3233210100332123-0030110213020220"></a>

## enable_api_discovery — enable_api_discovery / 101031021213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_api_discovery

<a id="canonical-0001201201333322-1133021323220003-2333030001222131-0223013300101113-2133221013222213-2220202223030213-2303012213101101-0010230023302110"></a>

Type: `"object"`. single nested block, Optional.

Specifies the settings used for API discovery.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_api_auth_discovery",
    "default_api_auth_discovery"),
  validators.ConflictingObjectAttributes("disable_learn_from_redirect_traffic",
    "enable_learn_from_redirect_traffic")}
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
  "x-ves-oneof-field-api_discovery_settings_choice": "[\"custom_api_auth_discovery\",\"default_api_auth_discovery\"]",
  "x-ves-oneof-field-learn_from_redirect_traffic": "[\"disable_learn_from_redirect_traffic\",\"enable_learn_from_redirect_traffic\"]"
}
```

Terraform syntax:

```terraform
enable_api_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-3023011101011330-3312102001100203-1002321210311201-0220221100021111-3012232202333210-1233123032211233-0331320130320022-2122232112220121"></a>

## Direct properties — enable_api_discovery / 101031021213 / 3

- [api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302): complete subsection reference.

- [api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100): complete subsection reference.

- [custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212): complete subsection reference.

- [default_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-1013221211221310-1031120002332131-0232001133000302-3000130332131232-1303121022332312-3012002312133111-1112121330233321-2233231122001212): complete subsection reference.

- [disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331120230102020-0312102032323311-1020022132220113-1312101223301033-0301110002231123-3321223213133301-2033313330221330-1323102320001301): complete subsection reference.

- [discovered_api_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-2112030012103111-1212023301022210-3131303031130201-1232223310233020-1021210103230301-2110223201003332-1112033210222011-3100331330221113): complete subsection reference.

- [enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-0021233123301122-1013002210302002-1231011123110332-0101131021133112-1321303230003032-3032101231010222-0000133020332203-2201003101132010): complete subsection reference.

<a id="canonical-3133323321131032-3030201203321002-0310102333310023-3122320230032120-0033032033111333-3123000030212032-3302011323133013-1113301222001310"></a>

## Next pages — enable_api_discovery / 101031021213 / 4

- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212)
- [enable_api_discovery.default_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-1013221211221310-1031120002332131-0232001133000302-3000130332131232-1303121022332312-3012002312133111-1112121330233321-2233231122001212)
- [enable_api_discovery.disable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331120230102020-0312102032323311-1020022132220113-1312101223301033-0301110002231123-3321223213133301-2033313330221330-1323102320001301)
- [enable_api_discovery.discovered_api_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-2112030012103111-1212023301022210-3131303031130201-1232223310233020-1021210103230301-2110223201003332-1112033210222011-3100331330221113)
- [enable_api_discovery.enable_learn_from_redirect_traffic](resources--cdn_loadbalancer--reference--group-010.md#canonical-0021233123301122-1013002210302002-1231011123110332-0101131021133112-1321303230003032-3032101231010222-0000133020332203-2201003101132010)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031132220122031-0022213112110331-1300210103302331-0333231332031201-1031012011002123-0221113103232310-2001210031322300-2330331200001333"></a>

## enable_api_discovery.api_crawler — api_crawler / 111312022202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.api_crawler

<a id="canonical-3003102010320103-2001002303320310-3031102300112312-0303330213100202-3201321130230300-0131131021000230-1122313002322013-1302121212220232"></a>

Type: `"object"`. single nested block, Optional.

API Crawling. API Crawler message.

Upstream description:

API Crawler message.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("api_crawler_config",
    "disable_api_crawler")}
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
  "x-ves-oneof-field-api_crawler": "[\"api_crawler_config\",\"disable_api_crawler\"]"
}
```

Terraform syntax:

```terraform
api_crawler {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021301011320011-0030300331133311-1222310000323221-2022020022231231-2123100330122232-1111112313332231-2010313313331113-0303112102031322"></a>

## Direct properties — api_crawler / 111312022202 / 3

- [api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021): complete subsection reference.

- [disable_api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-0223301322310022-2112210323323023-1010031200031021-3002003330300110-1311032202200232-0120121012010001-2331132132312022-1323130100110212): complete subsection reference.

<a id="canonical-3013030301221222-1312203133333210-1120310300201333-1302222101212331-2300130323310031-2232110030033122-0032112210210213-2320033102303302"></a>

## Next pages — api_crawler / 111312022202 / 4

- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.disable_api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-0223301322310022-2112210323323023-1010031200031021-3002003330300110-1311032202200232-0120121012010001-2331132132312022-1323130100110212)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0331013313010023-2103302322110301-2003313023111100-0323132332132331-1302331233222002-0103111131310033-0022010100101113-3033023012303232"></a>

## enable_api_discovery.api_crawler.api_crawler_config — api_crawler_config / 303010222022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- enable_api_discovery.api_crawler.api_crawler_config

<a id="canonical-2122202003320020-0120333300302123-0330013012112013-3122023312102132-0303132002021121-1321312133200321-0013223212032013-0121323033102002"></a>

Type: `"object"`. single nested block, Optional.

Crawler Configure.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("domains")}
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
api_crawler_config {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230222101010030-3002112330100000-2332101032230121-0032010130023322-3121213201102011-3300301232220010-2322232122303323-1300232130032210"></a>

## Direct properties — api_crawler_config / 303010222022 / 3

- [domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333): complete subsection reference.

<a id="canonical-2032113100133131-0021202012111202-1313332111231303-0321330221313211-2203023013130222-0203313133201123-2301320233021002-0331002101120022"></a>

## Next pages — api_crawler_config / 303010222022 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202121031203013-3020003123123203-1311000313223033-2003200322203130-2023132232330103-1330303331223202-0330020120300001-1000030120033213"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains — domains / 013132022222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- enable_api_discovery.api_crawler.api_crawler_config.domains

<a id="canonical-1033131131213212-1333112301230031-0120220313223123-2133321012032202-0011320302323302-3120230022131033-3022110333111010-0103110231112020"></a>

Type: `"object"`. list nested block, Optional.

Enter domains and their credentials to allow authenticated API crawling. You can only include
domains you own that are associated with this Load Balancer.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("domain")}
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
    "ves.io.schema.rules.repeated.max_items": "32"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32"
  }
}
```

Terraform syntax:

```terraform
domains {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020313311101331-1201301033033122-0230101033220122-0201023322103300-1132132010022013-0322203030011230-3022321233300321-3001300003233221"></a>

## Direct properties — domains / 013132022222 / 3

<a id="canonical-0230233313133232-1022303320313013-3330131233221032-3031021233230113-2312101102332113-1221321233132122-2223212030321213-3300031122011230"></a>

<a id="canonical-0022131322332013-2132321111113320-0102021002320331-0311013300112103-2030013121020200-2202301101113111-1133303333221321-0200002320031331"></a>

## domain property — domains / 013132022222 / 4

Type: `"string"`. Optional.

Select the domain to execute API Crawling with given credentials.

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
    "format": "fqdn",
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
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111): complete subsection reference.

<a id="canonical-2310033311220321-3131031331202023-0023001321202200-1033322300222313-1300010333210233-2220202312113302-2022010313211023-2201130303002120"></a>

## Next pages — domains / 013132022222 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2333121220112103-0220323013130101-2132212203011331-1032001130222013-2212013202201201-2100222120331313-2010322032110212-2111011200320113"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login — simple_login / 233121222332 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login

<a id="canonical-0302213302123230-2232310000301231-1300002000210303-1003320103213333-2211120131213223-2231201103113013-0030131101322300-3020333102303012"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for simple login.

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
simple_login {
  # Configure direct properties listed below.
}
```

<a id="canonical-2302223212201221-2031223103123001-3030111022303031-0301031011211330-0101122020223332-3132110311033030-2310122000002233-3130002022210330"></a>

## Direct properties — simple_login / 233121222332 / 3

- [password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220): complete subsection reference.

<a id="canonical-1323013001013013-0030010300130220-3330033320220123-2021223133110032-3022011023130012-2132311110100331-3011120230032211-3220130020310033"></a>

<a id="canonical-3330321110322212-1211023201203020-0112001102011032-2311320033132331-0112013303001322-1203131220112131-2020210213000220-0221010130201012"></a>

## user property — simple_login / 233121222332 / 4

Type: `"string"`. Optional.

Enter the username to assign credentials for the selected domain to crawl.

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
      "allowed": "[a-zA-Z0-9_.-]",
      "description": "Alphanumeric with underscores, dots, hyphens"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-zA-Z0-9_.-]+$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```

<a id="canonical-2000302231213131-3121312023130322-2322203100010100-2332231210132201-1132010112021023-0323333320110010-2231211313231322-3232101323330310"></a>

## Next pages — simple_login / 233121222332 / 5

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0031310133211012-3120302002233312-2131031112020122-3213212332232120-2111032321111202-0120230322300111-1300232200311121-0230122102312031"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password — password / 331103032122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password

<a id="canonical-3332101023013122-2212130303321220-0003312321103212-2113013211200333-3330010232230023-2000233330103221-3332003232332023-2021330021010222"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Provider validators and defaults (from schema source):

```go
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
password {
  # Configure direct properties listed below.
}
```

<a id="canonical-1003310300223102-1100222221121000-2321013000030210-1033323100130211-1011101221000302-2132323303212203-2332211222102302-3333212130313313"></a>

## Direct properties — password / 331103032122 / 3

- [blindfold_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-0203021132213020-1023230023031320-0001220210233022-3131201223200301-0332210022120002-2313010213221231-3201102221031211-1223020110012220): complete subsection reference.

- [clear_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-3011333222233131-0113333113011123-3012100120011302-1112101232033213-2032213333000311-1223123111101003-2131011023013300-1103130003100103): complete subsection reference.

<a id="canonical-3312032313302312-3311010230030000-1100030033233132-0102221002030200-1200333321120300-3121212130110121-1002203132002213-1300311233123331"></a>

## Next pages — password / 331103032122 / 4

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-0203021132213020-1023230023031320-0001220210233022-3131201223200301-0332210022120002-2313010213221231-3201102221031211-1223020110012220)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info](resources--cdn_loadbalancer--reference--group-010.md#canonical-3011333222233131-0113333113011123-3012100120011302-1112101232033213-2032213333000311-1223123111101003-2131011023013300-1103130003100103)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0203021132213020-1023230023031320-0001220210233022-3131201223200301-0332210022120002-2313010213221231-3201102221031211-1223020110012220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331121330300001-0003123232231321-0200002313033202-1211300101031010-3220202131323310-1012101132312001-0210201023113202-3012032102103303"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info — blindfold_secret_info / 321131210100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.blindfold_secret_info

<a id="canonical-2301201131312000-1123301133012000-3331211200230211-0102332200330130-3233103232202110-0321211311331013-2213112103223031-1012101110213333"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-0120021220331211-0030311303201331-2121032301030301-2003001221113121-0203002311221222-0310123310132130-1021010201103313-2110220223033122"></a>

## Direct properties — blindfold_secret_info / 321131210100 / 3

<a id="canonical-2302132232012032-0131302000213103-3031312331231223-2000331033332023-2130212201213231-2222310131110302-2120322023222313-0012320233313013"></a>

<a id="canonical-0122120210330301-3202022230022122-2000103012112022-3101231300301033-0320323030232221-2100100001211120-3313213001213100-1012013231220300"></a>

## decryption_provider property — blindfold_secret_info / 321131210100 / 4

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

<a id="canonical-2212330103001132-2211032113010132-0202331221223110-3023001300132223-2013202312232331-0011302033030313-2323333013101312-0101223010033222"></a>

<a id="canonical-3322301001321013-3212323311132033-2213130101331203-1133013201013332-3332330100323221-2122320311233223-1001300001123020-0332322132130111"></a>

## location property — blindfold_secret_info / 321131210100 / 5

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Upstream description:

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Provider validators and defaults (from schema source):

```go
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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1301231330230201-0021111020130311-3131212203310110-1313132211020121-2332330230111303-3200033313330111-3102001301200220-3010222202110003"></a>

<a id="canonical-0203023130203203-3102010220322012-0131200303333202-1211103321213302-0022331030232201-3102021313112301-3332031020100111-1201010100322021"></a>

## store_provider property — blindfold_secret_info / 321131210100 / 6

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

<a id="canonical-3200223000032131-2000110311211200-0020102022310231-1322220323112233-2022303103101012-0232300220022221-2001100011131032-3031222003311333"></a>

## Next pages — blindfold_secret_info / 321131210100 / 7

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3011333222233131-0113333113011123-3012100120011302-1112101232033213-2032213333000311-1223123111101003-2131011023013300-1103130003100103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3200221111202312-0013223323302122-0032030231030120-1230023110130233-0303311120220330-3211233013221030-1210003301033302-3231220123200101"></a>

## enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info — clear_secret_info / 103332221101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [enable_api_discovery.api_crawler.api_crawler_config](resources--cdn_loadbalancer--reference--group-010.md#canonical-0332323223200020-3010332300032301-2210312101133212-1012222313113202-1110100001211033-2113312231222111-2201030312320313-0201201300203021)
- [enable_api_discovery.api_crawler.api_crawler_config.domains](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121220310110302-1022331020010110-1302331100202300-3020212122232233-2120310330111201-0012230022233003-2312222003111100-1123123002232333)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login](resources--cdn_loadbalancer--reference--group-010.md#canonical-2221103021022202-1300111132112132-3021211202212302-2022010021032103-1333113203201303-1232031231211112-0111210301033220-2111332311023111)
- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password.clear_secret_info

<a id="canonical-2123113313220110-0310302301200313-3212113121210233-3200210232310020-0110010012200211-1300033200023320-0202132210212032-0010330003311023"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2320312111013312-3221230123132013-2103233102121111-0012122122130300-2211310022103323-1103013122200030-2033032113211323-2120323003203110"></a>

## Direct properties — clear_secret_info / 103332221101 / 3

<a id="canonical-3322003002213231-0230310333030102-1201313103300313-2312120312212110-3320222111221010-2231130331313002-3003011231302320-2333332011012303"></a>

<a id="canonical-3200003130112213-0200032123030010-1221202210113011-3333233211222332-1203313021233001-3132003030203300-2002233113123120-2332121133123313"></a>

## provider_ref property — clear_secret_info / 103332221101 / 4

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1331210200012223-3311013231010113-3302111203111003-1313113203320212-3203012011122210-2322112310002101-3212322121102220-3330313223320323"></a>

<a id="canonical-1103012213332002-3103000130103312-1221012023122011-1001201022322030-1031000332213022-3313032212330112-0303032323203111-3100202111320130"></a>

## URL property — clear_secret_info / 103332221101 / 5

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3030333302120002-2020230013232303-1112220201011311-0323330100130212-1121113223212311-3120332013030330-1013222001022031-3031013211132200"></a>

## Next pages — clear_secret_info / 103332221101 / 6

- [enable_api_discovery.api_crawler.api_crawler_config.domains.simple_login.password](resources--cdn_loadbalancer--reference--group-010.md#canonical-1122021031323310-1230303113123121-3223331212330003-0231301202311001-3113112123023231-1121213233201100-3010330333211133-2223012221023220)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0223301322310022-2112210323323023-1010031200031021-3002003330300110-1311032202200232-0120121012010001-2331132132312022-1323130100110212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3031213231312220-1323201131231021-2300132110033300-1011103323121331-2121131001200013-2020111002100011-0102203011230333-0230110031201322"></a>

## enable_api_discovery.api_crawler.disable_api_crawler — disable_api_crawler / 313313300211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- enable_api_discovery.api_crawler.disable_api_crawler

<a id="canonical-3133131002302221-3223111311103332-0120003223003223-0020102313300121-1303332123233333-0312011031011130-0021320031010010-3112031301233302"></a>

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
disable_api_crawler = {}
```

<a id="canonical-3032322312131231-1030232323303133-0300002303123303-1202003101113313-2012130333123002-1333320320002002-1333233200122211-2222130213131020"></a>

## Direct properties — disable_api_crawler / 313313300211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2322310121132300-2301320321021222-1201323103313300-0322210001031122-2010112101123022-0011312103033112-2332232010011021-2033323103202310"></a>

## Next pages — disable_api_crawler / 313313300211 / 4

- [enable_api_discovery.api_crawler](resources--cdn_loadbalancer--reference--group-010.md#canonical-2312010323201210-3001303223201201-0123213333321222-3212233302021213-1330131213202213-0311122211302103-3133331232101210-2310312032020302)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302121320122110-0332000002303032-2001131200102122-3010221110120102-0302013002033123-0031111313103323-0330001323121113-2333111301322310"></a>

## enable_api_discovery.api_discovery_from_code_scan — api_discovery_from_code_scan / 223113301112 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.api_discovery_from_code_scan

<a id="canonical-1012300311230310-0001121003022332-0223030222213323-1030300202303311-1300332101121211-3123203113211231-2221201300122130-1311020030100120"></a>

Type: `"object"`. single nested block, Optional.

Select codebase and Repositories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("code_base_integrations")}
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
api_discovery_from_code_scan {
  # Configure direct properties listed below.
}
```

<a id="canonical-2002222023100200-1113102022312200-1333220030020031-2211332013133003-3123310022201132-2321032332221021-3323332032223220-3301332210332113"></a>

## Direct properties — api_discovery_from_code_scan / 223113301112 / 3

- [code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031): complete subsection reference.

<a id="canonical-2132331212210133-1230010310033101-0011222331121130-2310333223032321-3200302102120032-3022312003210231-0133022312030022-3230020203011003"></a>

## Next pages — api_discovery_from_code_scan / 223113301112 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333013212300302-2132321122200333-1100313110333221-3232313110102101-2330010333110330-3323101021203203-1002330132231123-3221310031313111"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations — code_base_integrations / 312202001101 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations

<a id="canonical-1133201031102003-3131001102211330-2110133203202212-2310302321232220-3030011203110303-2113121210013301-1120112103332321-3031023312012223"></a>

Type: `"object"`. list nested block, Optional.

Configuration parameter for codebase integrations.

Upstream description:

Configuration parameter for codebase integrations

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("all_repos",
    "selected_repos")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5,
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
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "5",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
code_base_integrations {
  # Configure direct properties listed below.
}
```

<a id="canonical-3213202011301332-2213302321312130-0021322120310010-3233020010223322-0110311122100330-2301202311111102-1132232312110100-1020231030032231"></a>

## Direct properties — code_base_integrations / 312202001101 / 3

- [all_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232): complete subsection reference.

- [code_base_integration](resources--cdn_loadbalancer--reference--group-010.md#canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223): complete subsection reference.

- [selected_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331): complete subsection reference.

<a id="canonical-1113113132000310-1322300020322120-2321111220333320-3233300201202011-0313223233211223-1020210130323203-3231211213113322-0120302313302200"></a>

## Next pages — code_base_integrations / 312202001101 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration](resources--cdn_loadbalancer--reference--group-010.md#canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos](resources--cdn_loadbalancer--reference--group-010.md#canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1133222312233233-3230003202230132-0331020231301303-1013330021223021-2003332220203123-1333213033231321-2101011233031021-1301003213132232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220321030023103-0131301031023113-0120010330320231-1320201031001311-1323003212103331-2021121213001101-1011303313230131-3202321031210222"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos — all_repos / 223322230213 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.all_repos

<a id="canonical-1000331013330222-3000113103113211-0001211010301300-1033013110223123-3221120202333012-2132002301113030-1001000130303322-0320300120233121"></a>

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
all_repos = {}
```

<a id="canonical-1200111122333121-2123221020103012-0201130301021010-0210001230310013-2333231231303123-2220232203220023-0331103212112023-2110203310312203"></a>

## Direct properties — all_repos / 223322230213 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022121123100310-0331133320033101-2122121132320331-3303201102231302-0322212131031112-3320101202203203-2011003031210131-0112220132013221"></a>

## Next pages — all_repos / 223322230213 / 4

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2121123213033030-0101132301023222-0210321103113032-3130313132023121-1023321313310131-1122011031001221-2122200023311221-1003322232210223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111021330021332-0020032120013313-2310231313101220-1321010210201303-0000321301311332-3102122011332303-0023301101312010-0101311323021131"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration — code_base_integration / 033130120120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.code_base_integration

<a id="canonical-3200022321013010-2212130332311312-3132123120032023-3231000103301002-2121202022103002-3030130133310002-3310023202310011-0330112210321032"></a>

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
code_base_integration {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111021122133120-1300100202300011-1123300301131212-0002110221311103-3030203111331210-0021013101232333-2123330223102132-2113311222211213"></a>

## Direct properties — code_base_integration / 033130120120 / 3

<a id="canonical-2220112313013320-1132100032313302-1000023013203022-2111020010001130-1013130133110121-3101130033123200-0020031032133123-1213112011220020"></a>

<a id="canonical-0032310122020112-2330023301133210-2121333231111111-0222231223113002-1100021003121202-0013130231301300-0233302332022113-0332212203310021"></a>

## name property — code_base_integration / 033130120120 / 4

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

<a id="canonical-2031312132122123-2121030213300113-1202213111200320-3230112321322101-2201031000211321-1020003321312133-3233303111003332-3210033322011222"></a>

<a id="canonical-2213100012011212-2023123221313100-2310201001112311-0312322231133211-0012300302322332-2232130221203321-0323031212023230-1231301123013330"></a>

## namespace property — code_base_integration / 033130120120 / 5

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

<a id="canonical-1300122230111202-2232322101032031-2012031232002231-3212101023303003-1231101131332320-2303301011231221-0311013221013232-2222110310033132"></a>

<a id="canonical-1002233220313013-0203332221200020-1200012123200031-2230231033023333-1320222130002302-1122103020012023-3322303103313232-2302020030210130"></a>

## tenant property — code_base_integration / 033130120120 / 6

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

<a id="canonical-0200232323320223-0033202200002230-1101003002010203-2202103032032121-1021310300300300-1121123132003212-1230022013203302-2031110320113213"></a>

## Next pages — code_base_integration / 033130120120 / 7

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3332003330020233-1321203101311112-1032320222112212-0303123303022320-0130233233212212-2331122210022100-2012123323130120-2300212333133331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0230001300133222-1222320301010301-3213330003233221-1111010030212233-1203312011200033-3231112333333121-0331210210121000-3210122110100203"></a>

## enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos — selected_repos / 130132332311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.api_discovery_from_code_scan](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012103001211100-1323221221232120-3012331302022232-3010203312233332-1323221131030032-2232202201310232-0020302012112200-0330211200110100)
- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- enable_api_discovery.api_discovery_from_code_scan.code_base_integrations.selected_repos

<a id="canonical-2023232133013100-2200211232213000-2201011220303123-1320133103031123-3102112121111312-2220023103310102-3100313313032210-1101013010020132"></a>

Type: `"object"`. single nested block, Optional.

Select which API repositories represent the LB applications.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("api_code_repo")}
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
selected_repos {
  # Configure direct properties listed below.
}
```

<a id="canonical-0231322021301000-2133223022230020-2022020103312230-1120221201331103-1300020223333021-3203032122231200-2032030132032233-0321210330202133"></a>

## Direct properties — selected_repos / 130132332311 / 3

<a id="canonical-0303003011020122-3333212220000322-2133133022133233-1221133210312310-0022110323330303-2231201103001302-2333311331321221-2022031032102221"></a>

<a id="canonical-0130300010220021-0221312320211332-0102130020030022-1322200000131033-0033213210301012-3011002230202113-2221030020011300-3330330003111322"></a>

## api_code_repo property — selected_repos / 130132332311 / 4

Type: `["list", "string"]`. Optional.

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

<a id="canonical-1131031302322010-1012020301331203-1232223132033233-0311020232123313-2233332021033231-0123031031102002-1210213222312103-1303331031102110"></a>

## Next pages — selected_repos / 130132332311 / 5

- [enable_api_discovery.api_discovery_from_code_scan.code_base_integrations](resources--cdn_loadbalancer--reference--group-010.md#canonical-1021021130013332-2232033332122023-0230000023030021-1220311223310132-1133102303121103-0121133300102103-2200033232313312-0321331113131031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0231221221302103-0211202002021321-2323021110133323-2303023132001020-2022202331221132-3322011233021212-1211030012303320-2102020231112210"></a>

## enable_api_discovery.custom_api_auth_discovery — custom_api_auth_discovery / 202233312311 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.custom_api_auth_discovery

<a id="canonical-2000000012223230-0211010000233313-1010013330320313-3003002303303013-3003012223220213-3332302231300133-2301220230113223-0113213212221220"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
custom_api_auth_discovery {
  # Configure direct properties listed below.
}
```

<a id="canonical-0102302330212102-1102200123123120-3000033303122110-0110131023210302-3303000013021020-3213101101002030-3330013021021031-3303302132121303"></a>

## Direct properties — custom_api_auth_discovery / 202233312311 / 3

- [api_discovery_ref](resources--cdn_loadbalancer--reference--group-010.md#canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330): complete subsection reference.

<a id="canonical-1030201032123032-1200313121100213-0213122110313213-2120003002221330-1210322003112011-3230130231033333-1023303132311011-3321002012211220"></a>

## Next pages — custom_api_auth_discovery / 202233312311 / 4

- [enable_api_discovery.custom_api_auth_discovery.api_discovery_ref](resources--cdn_loadbalancer--reference--group-010.md#canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1032323002010020-1022331203312301-0223203013030002-1003233130313211-0303313013223302-0222322103303123-0212103001230001-1121320112232330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3212101222033110-2113330102332322-3101332221113302-3020132121331313-1330201030121130-2221123003122231-0303023031102033-0213121133101020"></a>

## enable_api_discovery.custom_api_auth_discovery.api_discovery_ref — api_discovery_ref / 020320230120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212)
- enable_api_discovery.custom_api_auth_discovery.api_discovery_ref

<a id="canonical-2100223211110032-0022232233231102-3002001313210133-2120130303331110-3000320010313301-2113212302322012-0110123000000132-0013123202113003"></a>

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
api_discovery_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-1321203231133220-1010000332303122-3122033002303320-2122132212222221-0113311001011012-2022220312110100-0200132002213202-1121032202300320"></a>

## Direct properties — api_discovery_ref / 020320230120 / 3

<a id="canonical-1202013332032113-1311221103102202-0313222201100110-0220120112122133-1210103221112011-3121121213013331-3220132123302120-1002201232013312"></a>

<a id="canonical-3103110223120001-3220122002332200-3000230231202213-0311022300200313-2011231100133231-3110203233000211-3313100131331002-1313031212121213"></a>

## name property — api_discovery_ref / 020320230120 / 4

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

<a id="canonical-0211313023320013-0302112230002111-0003010130213213-1233001202103230-0323302221323222-3312132022011001-2003332110301021-2100211103311310"></a>

<a id="canonical-1033212010031100-3303122222320322-0101123032203031-1200130202210313-0300022321223011-0002113030030221-0033120002113222-1133332202123323"></a>

## namespace property — api_discovery_ref / 020320230120 / 5

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

<a id="canonical-1230233100131211-1010220312033333-2312032321233232-1303012123221001-0333021220203130-3333013233232321-1022100113322231-0102233302231111"></a>

<a id="canonical-3011212023322023-3221212321222133-3232021303031321-0112103132212013-3312223333201201-2321013120230330-0322032311021201-3210010000220123"></a>

## tenant property — api_discovery_ref / 020320230120 / 6

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

<a id="canonical-0322021200323031-3312300301303130-2122332331230031-1113232123011123-2311233110210310-1331103221100301-0310020232013003-3123220031200332"></a>

## Next pages — api_discovery_ref / 020320230120 / 7

- [enable_api_discovery.custom_api_auth_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2331112101123230-1113132102020030-3010300012020021-0030012202001320-0222213210303231-2210122300013233-3210320102310032-0332303223323212)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1013221211221310-1031120002332131-0232001133000302-3000130332131232-1303121022332312-3012002312133111-1112121330233321-2233231122001212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222133001100133-3131022220202112-0130210331103211-2222121120300230-2110100101312210-2330133120011033-1013220111200122-0003130000323033"></a>

## enable_api_discovery.default_api_auth_discovery — default_api_auth_discovery / 300022323001 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.default_api_auth_discovery

<a id="canonical-3330312033101230-0332222200233332-0101032312130121-0031230222130131-0023200000201230-1103120303010303-1320313222002032-1020121123313123"></a>

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
default_api_auth_discovery = {}
```

<a id="canonical-2322103023003133-1312313132100203-2102201112020121-1023131231331002-0311011100311320-3132312001321220-2313013321221333-3021033000031211"></a>

## Direct properties — default_api_auth_discovery / 300022323001 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021020103120222-0330132301100030-1233213230211002-1120001132201310-1330003323101030-2233313321311321-0200302032322030-2321311131011212"></a>

## Next pages — default_api_auth_discovery / 300022323001 / 4

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2331120230102020-0312102032323311-1020022132220113-1312101223301033-0301110002231123-3321223213133301-2033313330221330-1323102320001301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0210122230030103-3021223213033312-1000010110132300-0111323131011123-2012123103223032-0210013012222100-2012000311213312-1113301203311213"></a>

## enable_api_discovery.disable_learn_from_redirect_traffic — disable_learn_from_redirect_traffic / 001312102113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.disable_learn_from_redirect_traffic

<a id="canonical-1013231022323130-0133022002100312-2200002332003020-0032222010313033-3210300023103030-2312320031121232-0001223313301012-0033220122221221"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
disable_learn_from_redirect_traffic = {}
```

<a id="canonical-2332231220201331-1321210112303130-0301112333123303-2033230323121321-1123223123220330-3003103321021023-2331233011010303-0233301133332232"></a>

## Direct properties — disable_learn_from_redirect_traffic / 001312102113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132202331010333-0030022133122200-3033033312223102-2110222212223121-0310131303131322-1323312212211031-2233311302220222-3103233311133022"></a>

## Next pages — disable_learn_from_redirect_traffic / 001312102113 / 4

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2112030012103111-1212023301022210-3131303031130201-1232223310233020-1021210103230301-2110223201003332-1112033210222011-3100331330221113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3221232302112131-0300300230320213-2213311201330311-2230032033031110-1112102323131210-1222333220010223-2032022332301002-2013112131112023"></a>

## enable_api_discovery.discovered_api_settings — discovered_api_settings / 201332333012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.discovered_api_settings

<a id="canonical-3211302123213233-2213103023232312-1322230201223031-3113130221113323-3213022012202033-3103002213022020-0110212023330223-1131133203203212"></a>

Type: `"object"`. single nested block, Optional.

Discovered API Settings. Configure Discovered API Settings.

Upstream description:

Configure Discovered API Settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("purge_duration_for_inactive_discovered_apis")}
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
discovered_api_settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-1010030310321223-0330032231120113-3030233102302101-2303123302300313-0312113110101302-3323332001320110-1203020123113102-2232333221230013"></a>

## Direct properties — discovered_api_settings / 201332333012 / 3

<a id="canonical-0000013231022122-2030002133013321-3310220211312000-2200131132100102-3003021310122113-0002122013323233-1302113121220031-1110220311213120"></a>

<a id="canonical-0301020110222002-1210201203312213-3300313132321200-2021012001032322-0330111303222111-2131000000323102-2133322223333213-3101102012301311"></a>

## purge_duration_for_inactive_discovered_apis property — discovered_api_settings / 201332333012 / 4

Type: `"number"`. Optional.

Inactive discovered API will be deleted after configured duration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 7),
}
```

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

<a id="canonical-3220332233303121-0122230311221020-2303110021133002-3321330130210232-1012221101331023-2302201202331221-1122311023122331-1303130231210122"></a>

## Next pages — discovered_api_settings / 201332333012 / 5

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0021233123301122-1013002210302002-1231011123110332-0101131021133112-1321303230003032-3032101231010222-0000133020332203-2201003101132010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101123320303313-2232310223131102-3233011310013101-2011330132102321-2130301120102230-0130033120130303-3200020321000031-3313332230122321"></a>

## enable_api_discovery.enable_learn_from_redirect_traffic — enable_learn_from_redirect_traffic / 233130120212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- enable_api_discovery.enable_learn_from_redirect_traffic

<a id="canonical-1330012322210220-2231101001022001-3233031123022210-0230202013013011-2230333011002231-3220333302031202-1002320112023320-1231321311011222"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_learn_from_redirect_traffic = {}
```

<a id="canonical-2032200212213323-1221102111012200-0113313032003003-0231130220321201-2223120223010320-3213303012000113-2312130022131222-3231213010002330"></a>

## Direct properties — enable_learn_from_redirect_traffic / 233130120212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2211222122022102-3321211211023130-2003022103200110-0310231322103331-2030021302330101-2132211130203202-0121331110200211-3031310331112033"></a>

## Next pages — enable_learn_from_redirect_traffic / 233130120212 / 4

- [enable_api_discovery](resources--cdn_loadbalancer--reference--group-010.md#canonical-2303100100232313-3302230030213302-2031002323021102-1003031120212212-0002100220123231-1031131212300012-0213030103300100-2212030033012323)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2312320100311220-2121111023223010-2222123023231000-3231230201333232-2000321013003312-3122121312131333-2311101310003322-3032200103202113"></a>

## enable_challenge — enable_challenge / 022120221021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_challenge

<a id="canonical-1201312312020123-3230111211203203-0323103120212310-2002022312313232-2200320302313111-0002122232302200-2233200202311231-0212233210001213"></a>

Type: `"object"`. single nested block, Optional.

Configure auto mitigation i.e risk based challenges for malicious users.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("captcha_challenge_parameters",
    "default_captcha_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_js_challenge_parameters",
    "js_challenge_parameters"),
  validators.ConflictingObjectAttributes("default_mitigation_settings",
    "malicious_user_mitigation")}
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
  "x-ves-oneof-field-js_challenge_parameters_choice": "[\"default_js_challenge_parameters\",\"js_challenge_parameters\"]",
  "x-ves-oneof-field-malicious_user_mitigation_choice": "[\"default_mitigation_settings\",\"malicious_user_mitigation\"]"
}
```

Terraform syntax:

```terraform
enable_challenge {
  # Configure direct properties listed below.
}
```

<a id="canonical-2222020321233202-2001020131113233-1223132031032133-3012001211032030-3021002100021331-3322102200032212-2123021130333321-2122100311101231"></a>

## Direct properties — enable_challenge / 022120221021 / 3

- [captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012): complete subsection reference.

- [default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212): complete subsection reference.

- [default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110): complete subsection reference.

- [default_mitigation_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003): complete subsection reference.

- [js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122): complete subsection reference.

- [malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-010.md#canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203): complete subsection reference.

<a id="canonical-2012320132330302-1311223102200320-3310231021322221-3303312122200011-0300002133300133-2122232023332113-2101133223303011-0222010011013121"></a>

## Next pages — enable_challenge / 022120221021 / 4

- [enable_challenge.captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012)
- [enable_challenge.default_captcha_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212)
- [enable_challenge.default_js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110)
- [enable_challenge.default_mitigation_settings](resources--cdn_loadbalancer--reference--group-010.md#canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003)
- [enable_challenge.js_challenge_parameters](resources--cdn_loadbalancer--reference--group-010.md#canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122)
- [enable_challenge.malicious_user_mitigation](resources--cdn_loadbalancer--reference--group-010.md#canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1012112303301110-3222223022111121-0322323321333132-0233230321301001-3011120230103011-1321123300033331-0013312223020331-0102200200231012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3022331321200331-2101333111332103-3211120120210111-3113020131111101-1020312332331130-1213113123003213-3122021112223312-3121002300110012"></a>

## enable_challenge.captcha_challenge_parameters — captcha_challenge_parameters / 331120011033 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.captcha_challenge_parameters

<a id="canonical-0211123311231020-0003112120310203-2220013131110110-0021333120132010-1232202233030213-1131200301211130-3001200210101112-0323130031020303"></a>

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

<a id="canonical-2112000300021300-3033223223201021-1110032013133101-2223023020331232-1131322211233020-2221102003031030-3323132100031320-0232001211311013"></a>

## Direct properties — captcha_challenge_parameters / 331120011033 / 3

<a id="canonical-0113232331333331-3301201121111010-3210223121223023-3111113102333310-0313002211120202-0033130020101121-2320101211323120-0111223312232301"></a>

<a id="canonical-2230020020201002-1310313133102001-1002322003223001-2122210011222333-2102223023000203-2301302233002111-1200302031020220-1331223023300111"></a>

## cookie_expiry property — captcha_challenge_parameters / 331120011033 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-3311131020313213-2132233320310202-2310102223101331-2332231120032002-1013210123311021-1133200320131330-1213200133232303-0011012332001002"></a>

<a id="canonical-2013122013120100-0332330222020021-2123133101303302-2012321033132131-3231213020031330-0021332222201032-2013023332303303-0113323102222202"></a>

## custom_page property — captcha_challenge_parameters / 331120011033 / 5

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

<a id="canonical-2331101113322223-0003312003023131-1231312302312102-3132022122310300-0110213000323121-3202003110301003-0120022323212323-0002011220013113"></a>

## Next pages — captcha_challenge_parameters / 331120011033 / 6

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2212120120103301-3233031210133001-0010332200203330-3033001222120213-3112132200312113-1322100201132022-0200002302133322-0333302223330212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3132212211032020-3300220011020233-1122021333310003-2103130322322022-3001320231010022-0203021220213321-2331232113010330-2302230203033122"></a>

## enable_challenge.default_captcha_challenge_parameters — default_captcha_challenge_parameters / 020123030211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_captcha_challenge_parameters

<a id="canonical-3333113032103032-1001030133122233-0000132133113311-3120220111203201-0123211100233312-3311212113122220-2201021103002213-2021333032202220"></a>

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

<a id="canonical-0012300022001012-1210201012113021-1000233100123011-2222230313000321-0122131103231322-3133001121322203-2232221023210121-3230031001132203"></a>

## Direct properties — default_captcha_challenge_parameters / 020123030211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1310131312331222-2023203302123010-0020331031003112-0110021312022302-2332331213003200-1301120023203022-2003311033113212-0031321030101211"></a>

## Next pages — default_captcha_challenge_parameters / 020123030211 / 4

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1121110313000012-3310321102200200-3113100211012323-3003132311310212-1230221132012011-1002211112133301-0031330011112203-0302213312133110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001322231230212-0323233302322111-0231111312000022-3203233130221200-2003320200021233-2323020223002022-3222100021003012-1213321312030331"></a>

## enable_challenge.default_js_challenge_parameters — default_js_challenge_parameters / 203011131303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_js_challenge_parameters

<a id="canonical-1210331103002233-1230202212223322-0232231320321303-3010212202200130-3332331010221113-1011112221033100-3010101032032100-0011013032203001"></a>

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

<a id="canonical-1022033031330320-1121012110312132-1122212122332111-1300310310303022-0023131222110202-2102212113211003-1321012221213313-2210313111211332"></a>

## Direct properties — default_js_challenge_parameters / 203011131303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201231323213201-2013303311111313-2322100123221132-3013131302223012-2020302121203323-3121003322021122-3131021133202302-2023332313021211"></a>

## Next pages — default_js_challenge_parameters / 203011131303 / 4

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3003123010123300-3301100122332312-3002011132220020-2333111001110100-1212100130022031-2213210313032210-0121003122010333-3322202033313003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1021230112121301-3013131331001232-2023320033313001-1120101032312011-3030322221123130-2333102322332100-0202332201021133-2230311022121103"></a>

## enable_challenge.default_mitigation_settings — default_mitigation_settings / 101122200320 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.default_mitigation_settings

<a id="canonical-0300221213130223-0123223121003201-0322310232332121-0202121230200012-2321332300132320-0223302031303011-2223230010221233-3011330222120330"></a>

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

<a id="canonical-2213211030100331-3232020230222022-2111203321023322-2321022331223213-1022333132130222-1001013300123232-0101031212031323-1121032122121212"></a>

## Direct properties — default_mitigation_settings / 101122200320 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1332203020101201-0223322112002002-1213331220002023-3200032320020101-3321303100033233-1210322101221131-2230210203330133-3113002221310201"></a>

## Next pages — default_mitigation_settings / 101122200320 / 4

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0312101012232110-0001011211003323-1030010110111113-0310302333212012-0111321322111020-3021022301102212-0232230030320121-0333232231023122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202103312023101-2330031321110031-2331000030310023-0233233313320020-3102213330000302-2001200210100022-3001212222000230-0330031121220110"></a>

## enable_challenge.js_challenge_parameters — js_challenge_parameters / 002303301022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.js_challenge_parameters

<a id="canonical-0323203300302101-1222220231002133-1303201020320013-0122301013313202-1231200110111332-2233103301103001-2123203121202012-3010322110322023"></a>

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

<a id="canonical-2301000310332320-3313203212013322-3012332102102101-2330020211033202-2322230222330313-2303213003023010-0321323333203030-0220010133311100"></a>

## Direct properties — js_challenge_parameters / 002303301022 / 3

<a id="canonical-0033003031323131-2020101113101303-1303311130110220-2331313032022201-3310330101322221-1013302110103030-2220012103001131-3020213102210113"></a>

<a id="canonical-1310230120221300-0011231200200121-2100012232320232-3111020112312102-0312132220221310-0101122022313332-0132221102302200-2301033331022203"></a>

## cookie_expiry property — js_challenge_parameters / 002303301022 / 4

Type: `"number"`. Optional.

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Upstream description:

Cookie expiration period, in seconds. An expired cookie causes the loadbalancer to issue a new
challenge.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-2331310112023121-3121123302133023-2311110113133302-2100311211030323-0313310322203231-3220311320200112-3300131121012320-1122202120330202"></a>

<a id="canonical-3232331100323000-2301231333312332-1011210110332113-1211000122100012-0021330311100202-2333101023330231-3232031112202002-3322123212131332"></a>

## custom_page property — js_challenge_parameters / 002303301022 / 5

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

<a id="canonical-0202132112111103-3231121320123332-0013222333311010-2020101013220303-0230203203120232-3103212103230013-1213132122223213-2233212311210300"></a>

<a id="canonical-3330321013101020-2111132131112002-0131012201020023-0312230303232203-2022220220221001-3013323020023223-3133002333023232-1203223101203112"></a>

## js_script_delay property — js_challenge_parameters / 002303301022 / 6

Type: `"number"`. Optional.

Delay introduced by JavaScript, in milliseconds.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1310222000222031-1201123211303130-2013121131132103-2121233212311003-1000100130033012-3001132300230300-2133112100112233-0110210211003101"></a>

## Next pages — js_challenge_parameters / 002303301022 / 7

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0323300102211113-0133120131332130-3221200230202231-0232201333102322-1131003112030231-0213300330113222-2103122033011023-1113122230202203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3011030003301312-0323131112213330-1232013330203320-0223310012003231-2311010312330121-1213202330311231-1032001111220331-0212332130012303"></a>

## enable_challenge.malicious_user_mitigation — malicious_user_mitigation / 301332102333 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- enable_challenge.malicious_user_mitigation

<a id="canonical-3011302213322302-0112111011213203-2211103221013111-1102302231012201-2103101231221103-3211100203310131-1122133201203103-3020123030303130"></a>

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
malicious_user_mitigation {
  # Configure direct properties listed below.
}
```

<a id="canonical-3202031212211010-1000310230001220-1222320033110102-1120112023332202-2310133002110323-1332120232200302-0213033100311013-2232331031012221"></a>

## Direct properties — malicious_user_mitigation / 301332102333 / 3

<a id="canonical-3230211000302300-2312302011233301-1233333123111232-0001122110132201-0231133230330302-0112222301013123-1102121201133032-3232220323103333"></a>

<a id="canonical-3000321313331123-3233131101110002-2022100221330130-2100333101131323-3023001223213330-0332320213100300-3122010313131301-1120213111323213"></a>

## name property — malicious_user_mitigation / 301332102333 / 4

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

<a id="canonical-3122013203130123-3220023132222120-1101033220011321-0312333222120001-3201133220312033-1221213222320332-1221332331021321-0132312333121211"></a>

<a id="canonical-3011322003121032-2211131311103133-2013222233221320-2032123203223311-3110220230000132-0301123332102300-2021011230323123-3211131031211021"></a>

## namespace property — malicious_user_mitigation / 301332102333 / 5

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

<a id="canonical-2322103131322131-0122312130133303-0300133222030010-1000211333211300-2121310322330221-1113023202220030-2202123212232232-1101202220020330"></a>

<a id="canonical-3133000133223001-1232122021010013-1120130210130010-0002210020101322-1323020201310213-0131020101223221-2313122212001303-3030221221001031"></a>

## tenant property — malicious_user_mitigation / 301332102333 / 6

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

<a id="canonical-3222223132311302-0031113222300323-1031212001133233-1031211103232321-3231012232321300-1320112321331201-0133100331202032-0113330312210102"></a>

## Next pages — malicious_user_mitigation / 301332102333 / 7

- [enable_challenge](resources--cdn_loadbalancer--reference--group-010.md#canonical-2332130232332220-3122100131132301-2112222332200300-2232023130202233-3132033002233230-3320011212210312-2203130020311011-3020303300221303)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3132020110023330-0221123212220300-2303122132103021-2301032210323100-0112321332310001-2233030031121302-2100003011131031-0321200323303122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1032321122230201-3300312212102222-2111201333313311-1200211223033002-0303130300310102-3012200321100321-3303211021120200-1023322223303132"></a>

## enable_ip_reputation — enable_ip_reputation / 011333310120 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_ip_reputation

<a id="canonical-1123223010010320-2010330211102303-3003132133301111-3013321110233132-3220003300022030-2113002030033003-3221112313032202-3320311130101320"></a>

Type: `"object"`. single nested block, Optional.

IP Threat Category List. List of IP threat categories.

Upstream description:

List of IP threat categories.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("ip_threat_categories")}
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
enable_ip_reputation {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202302133130323-3302313200220232-1031233133212103-3223020332303122-1102323130030300-1210100302023323-3001323123332221-0213333203100111"></a>

## Direct properties — enable_ip_reputation / 011333310120 / 3

<a id="canonical-2222233221300223-2333220231013132-2233000201012203-2203002130020121-2110220333122330-3013100100110121-3121002111113221-3113320200211320"></a>

<a id="canonical-0212110200032301-2212022013102223-3000311012033031-1231302333023212-0013122222322333-1220320021303113-0111030002000320-0312320321231223"></a>

## ip_threat_categories property — enable_ip_reputation / 011333310120 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
SPAM\_SOURCES|WINDOWS\_EXPLOITS|WEB\_ATTACKS|BOTNETS|SCANNERS|REPUTATION|PHISHING|PROXY|MOBILE\_THREATS|TOR\_PROXY|DENIAL\_OF\_SERVICE|NETWORK\]
If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied. Possible values are \`SPAM\_SOURCES\`, \`WINDOWS\_EXPLOITS\`, \`WEB\_ATTACKS\`, \`BOTNETS\`,
\`SCANNERS\`, \`REPUTATION\`, \`PHISHING\`, \`PROXY\`, \`MOBILE\_THREATS\`, \`TOR\_PROXY\`,
\`DENIAL\_OF\_SERVICE\`, \`NETWORK\`. Defaults to \`SPAM\_SOURCES\`.

Upstream description:

If the source IP matches on atleast one of the enabled IP threat categories, the request will be
denied.

Provider validators and defaults (from schema source):

```go
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

<a id="canonical-1001022102011101-1101102322201131-1020331120211310-0013222030030130-2012023121031200-2333003133202210-0310010210321011-3331232203120301"></a>

## Next pages — enable_ip_reputation / 011333310120 / 5

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1303213031102030-3122302223322303-3002003203030013-0131022002010001-0002312200302222-3331123211021133-0230132330122210-3301210122110300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202221332020332-1033130233031213-1113202110321232-3330003332303300-1231032102023332-1233011111110332-0110300213301231-2213311320230100"></a>

## enable_malicious_user_detection — enable_malicious_user_detection / 231031112113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_malicious_user_detection

<a id="canonical-1123232203103010-2221320123213112-1303032131321013-2032233321303102-1212120310102022-0332230300221210-2020311101222333-2212332330213300"></a>

Type: `["object", {}]`. Optional.

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

Terraform syntax:

```terraform
enable_malicious_user_detection = {}
```

<a id="canonical-0111222121100130-1321020303000310-1323200120223320-0000221202033311-3000113210011000-3123102221131300-1021002021222132-3223011331302213"></a>

## Direct properties — enable_malicious_user_detection / 231031112113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1001223221223230-1320001111030022-0331222313231320-2311013320310332-0011312001222301-3333033311223020-2030010301113233-3230121103000122"></a>

## Next pages — enable_malicious_user_detection / 231031112113 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3301230133222123-3331320000010320-3300322323323202-1103213233103032-2103000031113100-2212030000211223-1213133100220200-0023212123220030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0222103031101031-2000122332311230-2130132310133322-0202333202101320-0203301312220113-2030313011110020-2133233130020313-3232310322130132"></a>

## enable_threat_mesh — enable_threat_mesh / 302101001233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- enable_threat_mesh

<a id="canonical-2123213233303030-1130113101312333-0212210331021330-3103233332320113-1210122212022213-2121031100211332-1031010102222213-2210032031223013"></a>

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
enable_threat_mesh = {}
```

<a id="canonical-2203313221022221-2122022223331002-3332011223112320-3011110020333311-3323221310010220-3012110123120303-1013130001311110-2112331131021220"></a>

## Direct properties — enable_threat_mesh / 302101001233 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122113220321013-1221320001201311-1232002200311101-3232003023231311-3120330022221030-3112223012221010-0011302202311233-2312021102332303"></a>

## Next pages — enable_threat_mesh / 302101001233 / 4

- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0131310203002311-1002202102331210-3120001102321131-2313330213020203-3112000203032220-1033322112213220-3032123210010110-1010202111302223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3222132010310333-3320233301012221-2020133023031332-0102123222032003-2220313210230110-2102203012301223-3033102130122113-0311302333213312"></a>

## graphql_rules — graphql_rules / 010023330303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- graphql_rules

<a id="canonical-3033113222332331-3132301221000332-2212230233231332-2330333210220301-0322131001130303-1131320103233222-1330301211123330-2232113032230103"></a>

Type: `"object"`. list nested block, Optional.

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy..

Upstream description:

GraphQL is a query language and server-side runtime for APIs which provides a complete and
understandable description of the data in API. GraphQL gives clients the power to ask for exactly
what they need, makes it easier to evolve APIs over time, and enables powerful developer tools.
Policy configuration to analyze GraphQL queries and prevent GraphQL tailored attacks.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("exact_path"),
  validators.ConflictingListObjectAttributes("any_domain",
    "exact_value"),
  validators.ConflictingListObjectAttributes("any_domain",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("exact_value",
    "suffix_value"),
  validators.ConflictingListObjectAttributes("method_get",
    "method_post")}
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

Terraform syntax:

```terraform
graphql_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-0032322322232311-1121101200302211-3120110320323313-3233031121122133-1102213020323321-1322333312032123-0021201210002113-3102230313032030"></a>

## Direct properties — graphql_rules / 010023330303 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-011.md#canonical-0310312333200300-1001133112201301-2231012213003033-1211231123311203-0233000022332111-0012232012120121-1100033002132011-3333220110033032): complete subsection reference.

<a id="canonical-2112013200300110-3201131023132300-0102210030000203-0033322220111131-1213310331123121-0000301132321131-1232033122011130-3020003032010123"></a>

<a id="canonical-0102021211022121-2100310030312022-2302231322120111-1220112231223012-2111110103102221-1232212332121222-2301313032223300-1102221313203320"></a>

## exact_path property — graphql_rules / 010023330303 / 4

Type: `"string"`. Optional.

Specifies the exact path to GraphQL endpoint. Defaults to \`/graphql\`.

Upstream description:

Specifies the exact path to GraphQL endpoint. Default value is /GraphQL.

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

<a id="canonical-1100232233221102-3302033233011201-2230220112220103-2301130031231210-3322031300000320-0300022123221231-0333320021011211-3210112013322302"></a>

<a id="canonical-0333313022001313-2023033012100312-3220102223121123-2001030323213122-0210330102231133-2021012012121320-2121210332123200-3330222312300211"></a>

## exact_value property — graphql_rules / 010023330303 / 5

Type: `"string"`. Optional.

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

Upstream description:

Exclusive with \[any\_domain suffix\_value\] Exact domain name.

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

- [graphql_settings](resources--cdn_loadbalancer--reference--group-011.md#canonical-0303202121321220-1213302012121122-3010013331001221-0330202211132123-3330213012212101-3031001203002213-1113001320220013-0230201221030331): complete subsection reference.

- [metadata](resources--cdn_loadbalancer--reference--group-011.md#canonical-2111103032113221-2300330331032303-3021202023221012-1300020310021031-1001113031002312-0122031211203333-2111112020000120-2203030003322211): complete subsection reference.

- [method_get](resources--cdn_loadbalancer--reference--group-011.md#canonical-3112311302031310-2013010311020001-1012213233002331-3203312232102320-3332222023031121-2113023320012123-3003232203122311-3020001112232321): complete subsection reference.

- [method_post](resources--cdn_loadbalancer--reference--group-011.md#canonical-3101201020133222-2123332031201300-1233023232013013-3010301220131123-0000011132001021-1031001333131330-3303102132212030-1201311020130133): complete subsection reference.

<a id="canonical-3332221223013123-2033322221131331-3312213031321211-1013303313313001-2210013302100121-3132133122013122-2103110032030032-0321001312300003"></a>

<a id="canonical-1130302133012222-0133103310333233-2213112333021301-1221331011111302-2033301031010011-1031221222301113-3133020130212320-0020220132222323"></a>

## suffix_value property — graphql_rules / 010023330303 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Upstream description:

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g "xyz.com" will match
"\*.xyz.com" and "xyz.com"

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
