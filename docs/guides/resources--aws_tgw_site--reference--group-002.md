---
page_title: "xcsh_aws_tgw_site reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_aws_tgw_site reference."
---

# xcsh_aws_tgw_site reference

<a id="canonical-0300022120212210-3312311111132100-2030111122121212-3323031300112111-1131201020023012-1112013201302223-0312310103022301-3030112201131321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_disabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- direct_connect_disabled

<a id="canonical-3122301210323203-0212333020332103-3222231132103033-0120023200212032-1133103322211222-1022132010311032-1113200013022001-2103001020300111"></a>

Type: `["object", {}]`. Optional.

\[OneOf: direct\_connect\_disabled, direct\_connect\_enabled, private\_connectivity\] Enable this
option

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

OneOf alternatives in this subsection:

- [direct_connect_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-3122301210323203-0212333020332103-3222231132103033-0120023200212032-1133103322211222-1022132010311032-1113200013022001-2103001020300111)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-0101210101031120-0031320100323200-0003123032213330-2111000120012310-3013021323212201-2113331121220033-1320011211103123-3030003300300023)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0032101031203303-1213220331310310-0100112031003321-2233001013333112-0313130131201302-3200303332021210-2121212012311110-1202302003101320)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
direct_connect_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- direct_connect_enabled

<a id="canonical-0101210101031120-0031320100323200-0003123032213330-2111000120012310-3013021323212201-2113331121220033-1320011211103123-3030003300300023"></a>

Type: `"object"`. single nested block, Optional.

Direct Connect Configuration. Direct Connect Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("auto_asn",
    "custom_asn"),
  validators.ConflictingObjectAttributes("hosted_vifs",
    "standard_vifs")}
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
  "x-ves-oneof-field-asn_choice": "[\"auto_asn\",\"custom_asn\"]",
  "x-ves-oneof-field-vif_choice": "[\"hosted_vifs\",\"standard_vifs\"]"
}
```

Terraform syntax:

```terraform
direct_connect_enabled {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000303002131002-0021323122023010-1130131003133212-0230033103022113-2311121211210013-3313031230200001-3011122211233030-3303212222103101"></a>

### Direct properties for `direct_connect_enabled`

- [auto_asn](resources--aws_tgw_site--reference--group-002.md#canonical-1212212231300100-3002120022310313-2111300031220131-3012121123321321-0022032223021300-0000311222223301-0131030232112302-2220110231321023): complete subsection reference.

<a id="canonical-2313213201302220-3211201102221122-3230330010001211-2002302330213032-1210133120210103-2313330332120033-2311333011210102-3211303130330200"></a>

<a id="canonical-2310320131202230-2003201010332320-0003032302202002-3323313323221300-0303120331122220-1201203032313113-0111111123031220-3330010320230103"></a>

#### `direct_connect_enabled.custom_asn` property

Type: `"number"`. Optional.

Exclusive with \[auto\_asn\] Custom Autonomous System Number.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Int64{
  int64validator.AtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.uint32.gte": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1"
  }
}
```

- [hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300): complete subsection reference.

- [standard_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-0203322303111331-2203130323333113-1223320210331012-3020332313111013-3323101013301101-3033212221200222-2013111020230323-2103231132202000): complete subsection reference.

<a id="canonical-1212212231300100-3002120022310313-2111300031220131-3012121123321321-0022032223021300-0000311222223301-0131030232112302-2220110231321023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.auto_asn` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- direct_connect_enabled.auto_asn

<a id="canonical-3121232211122213-2201321102313323-3313232122120001-0203232112130021-3332110312311003-1313001210220231-2133312220123231-3212233223132130"></a>

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
auto_asn = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- direct_connect_enabled.hosted_vifs

<a id="canonical-0101213200030233-1002131032330333-3322132312031301-2222130310133211-3021123001303033-0223032300213333-0031103331211302-0033111320332320"></a>

Type: `"object"`. single nested block, Optional.

AWS Direct Connect Hosted VIF Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site_registration_over_direct_connect",
    "site_registration_over_internet")}
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
  "x-ves-oneof-field-connectivity_options": "[\"site_registration_over_direct_connect\",\"site_registration_over_internet\"]"
}
```

Terraform syntax:

```terraform
hosted_vifs {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022203232101200-0213010111121233-1103122232033333-2302223231011001-3310310331323232-3122033102211111-3010103021120330-0230321223001121"></a>

### Direct properties for `direct_connect_enabled.hosted_vifs`

- [site_registration_over_direct_connect](resources--aws_tgw_site--reference--group-002.md#canonical-3211300021320132-0112201232011331-3011312231111300-3222102132100313-1013223121333313-0103003121102121-2220211033331200-2230300013010200): complete subsection reference.

- [site_registration_over_internet](resources--aws_tgw_site--reference--group-002.md#canonical-2233303031303323-2200031230302211-1120123022102133-0030301332213231-3313302220320011-2101210010103233-1312233312211302-3332120200130330): complete subsection reference.

- [vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303): complete subsection reference.

<a id="canonical-3211300021320132-0112201232011331-3011312231111300-3222102132100313-1013223121333313-0103003121102121-2220211033331200-2230300013010200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect

<a id="canonical-1302123302001321-3300232121100232-1102223012202033-0322220000222331-0323101112003120-3210001333212202-3310100023221212-3111023100031112"></a>

Type: `"object"`. single nested block, Optional.

CloudLink ADN Network Config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("cloudlink_network_name")}
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
site_registration_over_direct_connect {
  # Configure direct properties listed below.
}
```

<a id="canonical-0322012323001112-3213003120010212-2110112122111231-0101102130311130-1333201200210230-0323231131002101-2122030013211313-2212313332202233"></a>

### Direct properties for `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect`

<a id="canonical-2020113230220132-0320210213313323-2023031331130120-2003110020130301-3320231023031211-3003322212133100-1323132313203101-2100232310331321"></a>

#### `direct_connect_enabled.hosted_vifs.site_registration_over_direct_connect.cloudlink_network_name` property

Type: `"string"`. Optional.

Establish private connectivity with the F5 Distributed Cloud Global Network using a Private ADN
network. To provision a Private ADN network, please contact F5 Distributed Cloud support.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2233303031303323-2200031230302211-1120123022102133-0030301332213231-3313302220320011-2101210010103233-1312233312211302-3332120200130330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.site_registration_over_internet` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- direct_connect_enabled.hosted_vifs.site_registration_over_internet

<a id="canonical-1220203121310223-3032330333023223-3111211023101030-1020231331311101-3033032121300022-3002112230120120-2011301131131032-1302211102000210"></a>

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
site_registration_over_internet = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.vif_list` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="canonical-1102332011103332-3003102021133331-2023221300213012-1220130210320033-3211010122103223-1330212212021120-0010123001122201-1313202111312220"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("vif_id"),
  validators.ConflictingListObjectAttributes("other_region",
    "same_as_site_region")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
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
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
vif_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-3123100231311023-0002033112221333-0322312313132231-0320223201202300-2232012311212022-0113311211101320-2102100210003323-0101202102013013"></a>

### Direct properties for `direct_connect_enabled.hosted_vifs.vif_list`

<a id="canonical-1230333100130313-3123210233300333-0320133111300223-1210131101221130-2211310203333230-2010002100311202-0112210212223033-1322133330111013"></a>

#### `direct_connect_enabled.hosted_vifs.vif_list.other_region` property

Type: `"string"`. Optional.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["af-south-1","ap-east-1","ap-northeast-1","ap-northeast-2","ap-south-1","ap-southeast-1","ap-southeast-2","ap-southeast-3","ca-central-1","eu-central-1","eu-north-1","eu-south-1","eu-west-1","eu-west-2","eu-west-3","me-south-1","sa-east-1","us-east-1","us-east-2","us-west-1","us-west-2"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
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
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](resources--aws_tgw_site--reference--group-002.md#canonical-1103203201203221-2232200121201212-3001332002021010-3011110001033223-3033312013300232-3010110101202312-2123331310303333-2203213223132331): complete subsection reference.

<a id="canonical-3311133220111202-1201130302031100-3113022103321013-2011301233030231-1313013000002320-3302022213330223-1010331010121323-0123313021031111"></a>

<a id="canonical-1303130330330022-1213221020330333-3300032200322031-2100203222213020-1200202220020001-1233110001213231-3201222113013203-0000310132323003"></a>

#### `direct_connect_enabled.hosted_vifs.vif_list.vif_id` property

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

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
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

<a id="canonical-1103203201203221-2232200121201212-3001332002021010-3011110001033223-3033312013300232-3010110101202312-2123331310303333-2203213223132331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- [direct_connect_enabled.hosted_vifs](resources--aws_tgw_site--reference--group-002.md#canonical-2332030220330220-0011312032330110-3032020120103013-2010112332321323-3113221112030013-1012122201123002-2030023302301131-0101003301122300)
- [direct_connect_enabled.hosted_vifs.vif_list](resources--aws_tgw_site--reference--group-002.md#canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303)
- direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region

<a id="canonical-0203022322211121-1110020013101112-1111321310121311-3311113203210201-3233322013200211-2113003132031101-2123203202003201-1112231120003200"></a>

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
same_as_site_region = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0203322303111331-2203130323333113-1223320210331012-3020332313111013-3323101013301101-3033212221200222-2013111020230323-2103231132202000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `direct_connect_enabled.standard_vifs` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [direct_connect_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2230200213030222-3210301232310133-3003333211130302-2032210021222010-3232300331010023-0232100221031011-3110102001002130-0222113323033123)
- direct_connect_enabled.standard_vifs

<a id="canonical-3202231200200123-0213221330301310-1113000011111010-3012320231321102-2333300202232111-2133002311220111-3031312110232311-2110112112331202"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for standard vifs.

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
standard_vifs = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- kubernetes_upgrade_drain

<a id="canonical-1033200023221322-2013021020023300-1131032000223030-2302223233233033-1002030300321221-2333213323101132-1303230010021302-1311111233311103"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1302110301031112-3020002223203021-2333213300110101-3111130102200200-1312202112212110-2131333020300320-2121122113302223-3030301020220333"></a>

### Direct properties for `kubernetes_upgrade_drain`

- [disable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-0320100301132103-3223313010113130-2300120333231013-2311323302211321-1332220222003320-1222121122233230-2311200133012211-0323233000232013): complete subsection reference.

- [enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311): complete subsection reference.

<a id="canonical-0320100301132103-3223313010113130-2300120333231013-2311323302211321-1332220222003320-1222121122233230-2311200133012211-0323233000232013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.disable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- kubernetes_upgrade_drain.disable_upgrade_drain

<a id="canonical-2321311030030323-3303130110223011-1313111023330230-2312101022101201-0311101131101120-2033221321112030-0131032131210332-0000200303113003"></a>

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

<a id="canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- kubernetes_upgrade_drain.enable_upgrade_drain

<a id="canonical-1200232111122311-3110100123033033-2022302120121002-3131321030211202-1010212032220013-0312220211101223-0221011123231311-0033131231302222"></a>

Type: `"object"`. single nested block, Optional.

Specify batch upgrade settings for worker nodes within a site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0102323121311001-0011312030313303-1032213322211131-3100122230010113-1020013313110031-3021023031230213-1003220031003200-2001213022331220"></a>

### Direct properties for `kubernetes_upgrade_drain.enable_upgrade_drain`

- [disable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-0001000310120231-1020023311312103-1101002301112233-2100202200212013-2113201332131122-3303111011112332-3310111133032032-0133230100232302): complete subsection reference.

<a id="canonical-0100032102010030-0211110231131030-3310031320313032-3232302113201301-3311300030323300-2133012231131203-0212231330100231-0033131133102122"></a>

<a id="canonical-3321033101011102-0123011133303122-3202012112133312-2012000023020301-1230311201313320-2322021201020123-0233112001111132-2100130130232102"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_count` property

Type: `"number"`. Optional.

Node Batch Size Count. Exclusive with \[\]

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

<a id="canonical-2132020032200131-0122312123002001-2133220323330323-3210113101210332-1030131322022210-2100303003010030-3201032310123020-3310211112132233"></a>

<a id="canonical-1233232333011123-1113313011223202-2133313121133113-0103033230030023-2133313030302222-3012121303322233-3202013022313320-3123002103012311"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_max_unavailable_node_percentage` property

Type: `"number"`. Optional.

Maximum percentage of nodes unavailable during upgrade draining.

<a id="canonical-2332333120231310-3130220320112211-3230010203111202-3233031313002012-2332311100321131-0130312002320023-0313313012301021-2322120333330213"></a>

<a id="canonical-0012211333033211-3211002030222000-1010231331001013-0103332332220120-0020223000132011-2111120301113211-3330322313032113-1031002231032202"></a>

#### `kubernetes_upgrade_drain.enable_upgrade_drain.drain_node_timeout` property

Type: `"number"`. Optional.

Seconds to wait before initiating upgrade on the next set of nodes. Setting it to 0 will wait
indefinitely for all services on nodes to be upgraded gracefully before proceeding to the next set
of nodes. (Warning: It may block upgrade if services on a node cannot be gracefully upgraded. It is
recommended to use the default value).

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
    "ves.io.schema.rules.uint32.lte": "900"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "900"
  }
}
```

- [enable_vega_upgrade_mode](resources--aws_tgw_site--reference--group-002.md#canonical-2211032113230223-1131203232203132-0223033202031211-3332110322033201-3001122313123202-3101330102011322-1012022110220013-0002201210131200): complete subsection reference.

<a id="canonical-0001000310120231-1020023311312103-1101002301112233-2100202200212013-2113201332131122-3303111011112332-3310111133032032-0133230100232302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="canonical-1123113300100201-1232102030333230-0001112321103330-2332133220230131-3103221113110133-3233222303032031-1322310131322323-1210200331301203"></a>

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

<a id="canonical-2211032113230223-1131203232203132-0223033202031211-3332110322033201-3001122313123202-3101330102011322-1012022110220013-0002201210131200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [kubernetes_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-1133003220233223-0103011013313311-0210231331121001-1312002230012130-1301002123001123-3100032321303113-3312011031132303-0221202103220332)
- [kubernetes_upgrade_drain.enable_upgrade_drain](resources--aws_tgw_site--reference--group-002.md#canonical-2103130212033323-1230331311211111-0200302310110032-2313212322210011-1232112123132331-1122103302221313-0000323122222102-3131303021131311)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="canonical-2301300120132201-2230011223011021-0113233331330310-2131211201312221-1310021202212130-2031000122012101-0030323213113223-0020120003031212"></a>

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

<a id="canonical-1103022023311031-0103303110232123-3012233320211332-2012113211332023-0001320020120012-1023313201031223-2223210013123113-1123000001332222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `log_receiver` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- log_receiver

<a id="canonical-3200330213313000-2030032223302232-2201302110100212-2112331310122100-3222022123013021-3010123110322120-3312001030023232-2103001000200103"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: log\_receiver, logs\_streaming\_disabled\] Type establishes a direct reference from one
object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

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

OneOf alternatives in this subsection:

- [log_receiver](resources--aws_tgw_site--reference--group-002.md#canonical-3200330213313000-2030032223302232-2201302110100212-2112331310122100-3222022123013021-3010123110322120-3312001030023232-2103001000200103)
- [logs_streaming_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-2002101032021020-0020310230203133-2120133233201122-2330201230310123-1230132103102232-3201030321232221-0102033001221130-1101311022321020)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
log_receiver {
  # Configure direct properties listed below.
}
```

<a id="canonical-0332120313021231-3211122212301310-2223331133133333-3023101100312313-2031020123102022-2021213203003110-0130010303013001-1301130211120333"></a>

### Direct properties for `log_receiver`

<a id="canonical-0000202320001303-3323032101300221-3112021102003223-3023110313013100-2122130223322233-1132331232300323-2211131201212232-2301322230002022"></a>

#### `log_receiver.name` property

Type: `"string"`. Optional.

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

<a id="canonical-3110200222111133-0222012113123120-0320320200220330-1212221000020202-2121301321333221-1323200011020300-1303023130312233-3203312231230000"></a>

<a id="canonical-0132211222200220-2201022331032012-2203033133233232-1330203233230302-1233021211101211-3130102203321303-1103123210300101-1210221222222210"></a>

#### `log_receiver.namespace` property

Type: `"string"`. Optional, Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1332012211222331-1202230002030213-1212133223320202-1310121022001131-0330213202212001-0313133221203212-3133100120003233-2032030000302123"></a>

<a id="canonical-3203003310011130-1320010122312323-1220300103223212-3010102303211210-0120231101322323-1210323312222202-3100000021201313-0132322330013231"></a>

#### `log_receiver.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3331230112202213-2130303012010322-2331122233103201-0333302220133100-3033003202003300-0301310322330232-2200123022312220-2023201330131100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `logs_streaming_disabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- logs_streaming_disabled

<a id="canonical-2002101032021020-0020310230203133-2120133233201122-2330201230310123-1230132103102232-3201030321232221-0102033001221130-1101311022321020"></a>

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
logs_streaming_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- offline_survivability_mode

<a id="canonical-3300322230023203-3003210213133023-2000020032102101-2302030102313121-0313002003300123-1030030132102223-0233311001201132-0322030200021100"></a>

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

<a id="canonical-1001203002310022-3211330213233100-2020312201021020-1331112233113032-0033232202013102-2110110232002201-1231030131312030-0122223202102211"></a>

### Direct properties for `offline_survivability_mode`

- [enable_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1330200230003111-1130322033202321-0233022322200303-2120031100031312-2033330231120100-2333300203112312-1022330022220021-3300322001131333): complete subsection reference.

- [no_offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1211231003331122-3230110231022113-0022001200110322-0310103021323011-0000133310023113-3132000331002203-0332123011312211-1123102122322213): complete subsection reference.

<a id="canonical-1330200230003111-1130322033202321-0233022322200303-2120031100031312-2033330231120100-2333300203112312-1022330022220021-3300322001131333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode.enable_offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323)
- offline_survivability_mode.enable_offline_survivability_mode

<a id="canonical-3102311111032021-2102131022103122-0323013130233001-2202323223033230-3033023120201323-3001221330312231-2120012003231112-1102323011130112"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable offline survivability mode.

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
enable_offline_survivability_mode = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211231003331122-3230110231022113-0022001200110322-0310103021323011-0000133310023113-3132000331002203-0332123011312211-1123102122322213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `offline_survivability_mode.no_offline_survivability_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [offline_survivability_mode](resources--aws_tgw_site--reference--group-002.md#canonical-3001311020203333-3221233013320302-3311320030212213-1023113012213111-0220020321311331-1032303112200020-0010110000121231-3320101333113323)
- offline_survivability_mode.no_offline_survivability_mode

<a id="canonical-1210333111300321-0312011223003131-1000332133031110-2021202311030122-1012131300311001-0121202101233210-1200301212300000-0222110110212313"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for no offline survivability mode.

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
no_offline_survivability_mode = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303123002201131-0222010221210120-2033022000213311-1111020303311112-2012121320122023-3033220200130232-0231103221011210-1220321132322030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `os` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- os

<a id="canonical-2022233012203233-1131120110103021-2231122321223033-1110132121110001-0110233111012332-1133131122312010-0111113200211020-0321130302300020"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Operating System Version for the site. By default, latest available OS Version will
be used. Refer to release notes to find required released OS versions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-0032213210021031-3300301112330212-1302232313011012-1233001031032202-2110323200033311-1210222121332221-2123133213133103-0000022230220101"></a>

### Direct properties for `os`

- [default_os_version](resources--aws_tgw_site--reference--group-002.md#canonical-2331203310211212-3302313221001221-2013123300210331-0303032311322023-3102210230221323-3000303132210212-3102230122221230-1320323110323330): complete subsection reference.

<a id="canonical-0230031123322331-1223213200113230-2012100100201131-0122120131121010-2010030203111133-1302212222112301-1131003022103310-1203310311011212"></a>

<a id="canonical-3100222133222021-1021202111330231-0321131011200021-2230010211312011-1230211310230310-2122032123231211-2122031320100030-0210111332223110"></a>

#### `os.operating_system_version` property

Type: `"string"`. Optional.

Exclusive with \[default\_os\_version\] Specify a OS version to be used e.g. 9.2024.6.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2331203310211212-3302313221001221-2013123300210331-0303032311322023-3102210230221323-3000303132210212-3102230122221230-1320323110323330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `os.default_os_version` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [os](resources--aws_tgw_site--reference--group-002.md#canonical-2303123002201131-0222010221210120-2033022000213311-1111020303311112-2012121320122023-3033220200130232-0231103221011210-1220321132322030)
- os.default_os_version

<a id="canonical-1110320223313311-0102213232221200-0121123211033220-0220300200110212-3233132233230130-1133230212030031-1211301022233233-0020232011212100"></a>

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
default_os_version = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- performance_enhancement_mode

<a id="canonical-0033131101311310-3332202301032222-1231111021311112-2103103213010223-3221322322311112-2231101023030002-0030233130121020-0013122120101122"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-0112321222010131-1223101011113313-3131010302320213-0022332101121222-0102032231220130-2330331120323330-1121121130233012-2310021333000221"></a>

### Direct properties for `performance_enhancement_mode`

- [perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123): complete subsection reference.

- [perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330): complete subsection reference.

<a id="canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- performance_enhancement_mode.perf_mode_l3_enhanced

<a id="canonical-3031131233122022-3021100000201010-0100312021101123-0302201103003120-2222022303001011-0101132111330313-2102020030122230-3201322123332330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Additional upstream details:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-3030323100132130-3321321322300310-3113322122202232-0322213023220131-2220332000321223-3103112000010303-2121212201113002-0031122303330102"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l3_enhanced`

- [jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-3321002213031313-3332121033133332-2001332233010203-1100302131120132-1013002212133101-1223323301111133-1123222310033001-2331331320223313): complete subsection reference.

- [no_jumbo](resources--aws_tgw_site--reference--group-002.md#canonical-1010332012232311-3200012011213203-0100120310321002-0302223012302120-0330213001013212-3001211012002031-2230202000313123-0230321103313311): complete subsection reference.

<a id="canonical-3321002213031313-3332121033133332-2001332233010203-1100302131120132-1013002212133101-1223323301111133-1123222310033001-2331331320223313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.jumbo` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- performance_enhancement_mode.perf_mode_l3_enhanced.jumbo

<a id="canonical-1133202232121332-2001123221223200-0200201001122002-1123021220313321-3102021220212320-2233032232213203-0311313033221012-0221122210331111"></a>

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
jumbo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1010332012232311-3200012011213203-0100120310321002-0302223012302120-0330213001013212-3001211012002031-2230202000313123-0230321103313311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3003132121022311-2223212333230003-0233232222111121-3131033210122022-1121301121302013-3301022130321232-1313000102211112-3322000222130123)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="canonical-2102101130231123-3003330221202203-2100033210001220-1322113203212203-2002332203313302-1113321313213123-0232022331230110-3200112000003111"></a>

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
no_jumbo = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="canonical-3012111321231111-2101101120303013-3112002220233001-0103231321103101-1301102032000303-1121230021300303-0121131013013210-3202020210213011"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Additional upstream details:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022121330222001-2131332113232023-0330312023212220-1010122220001301-0101221020200232-2102101202201131-3313132201010300-1101220221130122"></a>

### Direct properties for `performance_enhancement_mode.perf_mode_l7_enhanced`

- [jumbo_disabled](resources--aws_tgw_site--reference--group-002.md#canonical-2312301122223030-2313002303210011-2021210333303330-0312303212220211-1013002220122022-3201002232121313-3230213311021022-0010230312132233): complete subsection reference.

- [jumbo_enabled](resources--aws_tgw_site--reference--group-002.md#canonical-2232003102101112-0000102001013332-2012231122013202-1100331100023300-1000000121223220-3303231231011221-0112331300123033-0012201133021100): complete subsection reference.

<a id="canonical-2312301122223030-2313002303210011-2021210333303330-0312303212220211-1013002220122022-3201002232121313-3230213311021022-0010230312132233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled

<a id="canonical-3232113332032210-3033020110001332-3120303332123222-1313201303211311-1323232330320102-2210133300022211-2301131313332112-2031331001121001"></a>

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
jumbo_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2232003102101112-0000102001013332-2012231122013202-1100331100023300-1000000121223220-3303231231011221-0112331300123033-0012201133021100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [performance_enhancement_mode](resources--aws_tgw_site--reference--group-002.md#canonical-1003000322233122-1203103022221013-2012312321203031-3020031313331223-2123222020211330-2332120211321231-1232113122132011-2222323003332113)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--aws_tgw_site--reference--group-002.md#canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

<a id="canonical-1302132332003232-2231010300212103-3220223101330203-1030103232103133-2320221032333322-0113023013130103-0211213101221213-3021311201102112"></a>

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
jumbo_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- private_connectivity

<a id="canonical-0032101031203303-1213220331310310-0100112031003321-2233001013333112-0313130131201302-3200303332021210-2121212012311110-1202302003101320"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for private connectivity.

Additional upstream details:

Private Connect Configuration.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("inside",
    "outside")}
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
  "x-ves-oneof-field-network_options": "[\"inside\",\"outside\"]"
}
```

Terraform syntax:

```terraform
private_connectivity {
  # Configure direct properties listed below.
}
```

<a id="canonical-2313103113033131-3223232013101131-2033302110133033-0102102100101003-2122333313310103-0111000010131012-3302232020312030-0000103000223212"></a>

### Direct properties for `private_connectivity`

- [cloud_link](resources--aws_tgw_site--reference--group-002.md#canonical-1033222232330221-0100013301112211-1303202302330111-2032023131012000-0330112130302333-2202111020220333-0331321122302132-1210221002130122): complete subsection reference.

- [inside](resources--aws_tgw_site--reference--group-002.md#canonical-2023313230113323-3212330110031301-2321233013221222-0303103323200022-3120103120231201-0000330013331130-0200202102300330-2020320010112123): complete subsection reference.

- [outside](resources--aws_tgw_site--reference--group-002.md#canonical-0120301302031320-1033020223303023-1321023321022311-1101000120300312-2200002021233212-2103030102312321-1122032310022332-2312011301133033): complete subsection reference.

<a id="canonical-1033222232330221-0100013301112211-1303202302330111-2032023131012000-0330112130302333-2202111020220333-0331321122302132-1210221002130122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity.cloud_link` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- private_connectivity.cloud_link

<a id="canonical-3111321303121122-2312011022310303-3100232113220302-2131113333123132-0011231211122320-0333233213113021-0210213211331231-1220120202231003"></a>

Type: `"object"`. single nested block, Optional.

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
cloud_link {
  # Configure direct properties listed below.
}
```

<a id="canonical-2301221000333002-0310303331210102-2012320303122200-3310233000222002-1101121222132032-0022213303012003-0223220332001300-1031320121322230"></a>

### Direct properties for `private_connectivity.cloud_link`

<a id="canonical-0232202221200200-0021331112120002-3210031312330320-2203122013111223-1320220211000013-3030111211001013-3022223310011123-2332320112133330"></a>

#### `private_connectivity.cloud_link.name` property

Type: `"string"`. Optional.

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

<a id="canonical-0013230331222100-1112002313210311-0011012202232320-0120012221100220-0303121330011123-1032322012213200-1013233013232111-3110320202202133"></a>

<a id="canonical-1020332302113000-0301033333131100-1313130312111201-2230130201020033-2333001012113022-3130302001200133-3003313213221112-0102123312223311"></a>

#### `private_connectivity.cloud_link.namespace` property

Type: `"string"`. Optional, Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2300311023332321-0130022203130210-2103210303300010-1210211122232211-1233011212030010-1211030123132122-3132220123321130-3320330233122120"></a>

<a id="canonical-3300232021130223-3131003232012320-3011210121333000-1220133210003310-2112301032021333-3313313021300211-0211311302031323-2120002003021032"></a>

#### `private_connectivity.cloud_link.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2023313230113323-3212330110031301-2321233013221222-0303103323200022-3120103120231201-0000330013331130-0200202102300330-2020320010112123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity.inside` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- private_connectivity.inside

<a id="canonical-3322222031113011-0022300130320002-0103000223333032-2002132022323331-3301233111322220-1021202102322010-3321201120003223-3300111221311300"></a>

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
inside = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120301302031320-1033020223303023-1321023321022311-1101000120300312-2200002021233212-2103030102312321-1122032310022332-2312011301133033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `private_connectivity.outside` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [private_connectivity](resources--aws_tgw_site--reference--group-002.md#canonical-0112133200322013-2300122300221313-0113120023203023-2303120023331133-1221201211013222-2113331302223210-0212130311102330-1220131330231211)
- private_connectivity.outside

<a id="canonical-3010220230021001-3111020322112312-0012310223231331-1003222101220002-2123202202201312-3230213231033031-3321013301121021-2233222320322231"></a>

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
outside = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1300131232312032-1320122110210113-3211331031330230-1322221001301333-0300033311131133-2330301300022230-1020032213210011-3220200321030231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sw` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- sw

<a id="canonical-1220203211332031-1012022012103003-2202213333310310-0003011022001030-2333301133330210-1321202102321312-3302111220323310-0220233021002132"></a>

Type: `"object"`. single nested block, Optional.

Select the F5XC Software Version for the site. By default, latest available F5XC Software Version
will be used. Refer to release notes to find required released SW versions.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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

<a id="canonical-1021103010311311-3211022220030213-2102012102113112-1103103100221322-3020200303022302-2032321132212121-0323321101220030-3311031213020030"></a>

### Direct properties for `sw`

- [default_sw_version](resources--aws_tgw_site--reference--group-002.md#canonical-2122101121313223-3322012222120200-3011303013000313-3220223332211102-3223000323012033-1321203110211113-1102203012031232-3202123010202331): complete subsection reference.

<a id="canonical-1321321203111033-1013033230200311-0123231331120201-3033201212120102-1213213211202313-0212000222010303-0123111021302231-1113023221232321"></a>

<a id="canonical-0221323102030120-1001230110011101-3201101010022102-2302213312233001-3103010002012100-3302312311211202-2030300121112112-0131331012311033"></a>

#### `sw.volterra_software_version` property

Type: `"string"`. Optional.

Exclusive with \[default\_sw\_version\] Specify a F5XC Software Version to be used e.g.
Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
    "ves.io.schema.rules.string.max_len": "20"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "20"
  }
}
```

<a id="canonical-2122101121313223-3322012222120200-3011303013000313-3220223332211102-3223000323012033-1321203110211113-1102203012031232-3202123010202331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `sw.default_sw_version` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [sw](resources--aws_tgw_site--reference--group-002.md#canonical-1300131232312032-1320122110210113-3211331031330230-1322221001301333-0300033311131133-2330301300022230-1020032213210011-3220200321030231)
- sw.default_sw_version

<a id="canonical-2331321303310332-1300321330320310-0001030011123100-0231302303203300-1210200211212113-3101132200131013-2113203202302100-0120030010221232"></a>

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
default_sw_version = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- tgw_security

<a id="canonical-1312311100131121-3233120200020121-0002032311133320-3321112003303032-1030013003323100-3023320302103001-2101023301132013-3111100330221013"></a>

Type: `"object"`. single nested block, Optional.

Security Configuration for transit gateway.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "east_west_service_policy_allow_all"),
  validators.ConflictingObjectAttributes("active_east_west_service_policies",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "active_network_policies"),
  validators.ConflictingObjectAttributes("active_enhanced_firewall_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "forward_proxy_allow_all"),
  validators.ConflictingObjectAttributes("active_forward_proxy_policies",
    "no_forward_proxy"),
  validators.ConflictingObjectAttributes("active_network_policies",
    "no_network_policy"),
  validators.ConflictingObjectAttributes("east_west_service_policy_allow_all",
    "no_east_west_policy"),
  validators.ConflictingObjectAttributes("forward_proxy_allow_all",
    "no_forward_proxy")}
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
  "x-ves-oneof-field-east_west_service_policy_choice": "[\"active_east_west_service_policies\",\"east_west_service_policy_allow_all\",\"no_east_west_policy\"]",
  "x-ves-oneof-field-forward_proxy_choice": "[\"active_forward_proxy_policies\",\"forward_proxy_allow_all\",\"no_forward_proxy\"]",
  "x-ves-oneof-field-network_policy_choice": "[\"active_enhanced_firewall_policies\",\"active_network_policies\",\"no_network_policy\"]"
}
```

Terraform syntax:

```terraform
tgw_security {
  # Configure direct properties listed below.
}
```

<a id="canonical-2230121333301331-1321322110202110-0011311220322002-3320122111103103-2101123313132213-1311331110223121-1200020221201302-1212012332133030"></a>

### Direct properties for `tgw_security`

- [active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030): complete subsection reference.

- [active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133): complete subsection reference.

- [active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311): complete subsection reference.

- [active_network_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120): complete subsection reference.

- [east_west_service_policy_allow_all](resources--aws_tgw_site--reference--group-003.md#canonical-1212231303220233-2010113303122033-3011320331223131-0321033320220030-1220333031302100-2013330003103020-3132213201030302-3013202322330332): complete subsection reference.

- [forward_proxy_allow_all](resources--aws_tgw_site--reference--group-003.md#canonical-1201120002231301-0121123222011213-0222322100210132-2331203031331310-1021212323122132-2103312223133222-1200301331133103-3000320022330201): complete subsection reference.

- [no_east_west_policy](resources--aws_tgw_site--reference--group-003.md#canonical-2023022012132131-2202121200230011-1120311221213111-2232023300131020-3021323023111122-3133100110121313-0221121132110210-2010132230010101): complete subsection reference.

- [no_forward_proxy](resources--aws_tgw_site--reference--group-003.md#canonical-2321021113032211-1112301111210302-2122022013230311-2310330302301131-1212323230212112-0001210222132001-2222022202023102-3331013003033310): complete subsection reference.

- [no_network_policy](resources--aws_tgw_site--reference--group-003.md#canonical-1203031221233023-1200122111122211-2020320201211200-1312131320023211-3002302120132300-3203030331100323-3021220313112001-0230023220113030): complete subsection reference.

<a id="canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_east_west_service_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_east_west_service_policies

<a id="canonical-1110123011202323-1301012013302030-3302320031320021-3130231311231300-3300033322110302-1120032122233231-0133230032331311-3103112210102221"></a>

Type: `"object"`. single nested block, Optional.

Active service policies for the east-west proxy.

Receipt-pinned upstream constraints:

```json
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
active_east_west_service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2033023333012323-0012223003001113-3032133300001333-1302031310012310-1131130022011013-2213032330213211-1112331011122232-0330101332033331"></a>

### Direct properties for `tgw_security.active_east_west_service_policies`

- [service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-3031210222331320-1321331113102131-0302223233021222-1112100300300103-1013233221322113-0010220002021301-1121021121101311-2233010303121102): complete subsection reference.

<a id="canonical-3031210222331320-1321331113102131-0302223233021222-1112100300300103-1013233221322113-0010220002021301-1121021121101311-2233010303121102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_east_west_service_policies.service_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_east_west_service_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2233103232021301-3211110120311313-1122222320100212-3011331000331302-3311002321132211-2222000130113213-0321012021212213-3312331031310030)
- tgw_security.active_east_west_service_policies.service_policies

<a id="canonical-0032013332330212-1223121033100231-0211013010300312-3200310103111020-3330330210001323-3130022301111011-2131112133212203-2003323231002333"></a>

Type: `"object"`. list nested block, Optional.

A list of references to service\_policy objects.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
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
service_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022201220311121-0321012200330312-0000212313132111-1113002313030122-2131102031132112-1303201121022231-0310320331023220-0112203313121033"></a>

### Direct properties for `tgw_security.active_east_west_service_policies.service_policies`

<a id="canonical-0132022012210012-0302103130320323-0112320303231331-0033212001222100-0032212010213300-0223332112130201-0121203032310123-3002310321102113"></a>

#### `tgw_security.active_east_west_service_policies.service_policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2323302332212223-3030033003231002-3111111213220311-3323122230102022-3023222010320221-2132001313322021-0111323133321322-1122130323333131"></a>

<a id="canonical-1031322230200033-1300230311232313-2100333300212112-1331332132110321-0213001313202332-1123202031200101-0321033110130302-0112220003022213"></a>

#### `tgw_security.active_east_west_service_policies.service_policies.namespace` property

Type: `"string"`. Optional, Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1203310231210323-1202132101011333-3103232112031203-1123230031333202-2302211020232013-3111332330030203-1231011322112122-1002220032102320"></a>

<a id="canonical-3230333122230032-0201101123220003-1321031133133133-3113021033201010-1223320022231100-2131112202321031-3123323102320132-2323001200222022"></a>

#### `tgw_security.active_east_west_service_policies.service_policies.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_enhanced_firewall_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_enhanced_firewall_policies

<a id="canonical-0223021110012232-2102232232321313-1300331100001221-2121131131222332-0101133022121232-3133330011203213-0101312232330033-2031021121112033"></a>

Type: `"object"`. single nested block, Optional.

List of Enhanced Firewall Policies These policies use session-based rules and provide all OPTIONS
available under firewall policies with an additional option for service insertion.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("enhanced_firewall_policies")}
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
active_enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1203221133023030-2131032102232001-0322331303121111-3201302101131321-2220021030330121-2313102222212012-2221031132013211-0111330232133132"></a>

### Direct properties for `tgw_security.active_enhanced_firewall_policies`

- [enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2331222203321201-2200122323322103-0123332223213331-2110100102323123-2220210323232113-2222232312221210-2102122211121032-3112031123130232): complete subsection reference.

<a id="canonical-2331222203321201-2200122323322103-0123332223213331-2110100102323123-2220210323232113-2222232312221210-2102122211121032-3112031123130232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_enhanced_firewall_policies](resources--aws_tgw_site--reference--group-002.md#canonical-2310103112323101-3232032222023122-3013133222200133-0011132033232220-2220103123223101-2110233232233231-3103023133133021-2202121332213133)
- tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies

<a id="canonical-0133031213232322-2011100200011223-1013312231310002-1000222322200300-0003301312122011-2310201300213302-0232033102311022-2321113013221021"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Enhanced Firewall Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
enhanced_firewall_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2100330023303200-0320020120310010-0030131300220210-2302323103131300-1210221230221313-1130120033220203-1331013001231030-2112312101131320"></a>

### Direct properties for `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies`

<a id="canonical-0002331331222220-3001220211102132-3123322023301321-0313002120232033-0131212313122033-0122112131201013-3332330000203003-1103220133003110"></a>

#### `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-1320231312112002-3233203000120332-0120113030321030-0033213100013200-2133320232132132-3302313000202220-2031133313211033-0332320321032303"></a>

<a id="canonical-3011323120303310-3021131221110002-3331111112223232-0113312231022332-2113232002313300-2332103211330221-0331031203113222-0330102121203313"></a>

#### `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.namespace` property

Type: `"string"`. Optional, Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3330302211231010-1122033211133123-2313113302203230-3113310221323120-2030332003002113-0010320011300010-1000131022333123-2131221330313230"></a>

<a id="canonical-1330300021223321-2210031001101030-0212030332313122-2202203330023333-0312123220221011-3230323031210203-3203020121313333-0022002011111113"></a>

#### `tgw_security.active_enhanced_firewall_policies.enhanced_firewall_policies.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_forward_proxy_policies

<a id="canonical-0033133103231112-1322011022002333-1101200110203123-2211112332330230-2033111032003301-1032010021000210-3120103020011233-1312333132003201"></a>

Type: `"object"`. single nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("forward_proxy_policies")}
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
active_forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-1323201320211130-0230022212033220-3000213303131110-2322302113031020-2203132003200321-1323330332103022-1000223132023023-1331003331330133"></a>

### Direct properties for `tgw_security.active_forward_proxy_policies`

- [forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1111110110312031-3311213033030211-0211323330230313-3211200330202220-0112122010221201-0123012131313101-3103301111201130-1130200312313302): complete subsection reference.

<a id="canonical-1111110110312031-3311213033030211-0211323330230313-3211200330202220-0112122010221201-0123012131313101-3103301111201130-1130200312313302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_forward_proxy_policies.forward_proxy_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- [tgw_security.active_forward_proxy_policies](resources--aws_tgw_site--reference--group-002.md#canonical-1031013310020101-2212122023230013-2013203210021030-1131300210302300-3323030102103103-2203030010012100-1133111331201201-0002132330131311)
- tgw_security.active_forward_proxy_policies.forward_proxy_policies

<a id="canonical-0120012202100031-1003213003133310-0010210302111203-2322111200130301-3101221011013332-1330111233102221-1110100230121323-0022103331111122"></a>

Type: `"object"`. list nested block, Optional.

Ordered List of Forward Proxy Policies active.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{validators.RequiredListObjectAttributes("name")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "minItems": 1,
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
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
forward_proxy_policies {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021020302310220-2033222231123303-2220200313331330-2110121103332203-2011230020101011-2301230112030003-2010000121020301-0011321330120302"></a>

### Direct properties for `tgw_security.active_forward_proxy_policies.forward_proxy_policies`

<a id="canonical-3320012022220330-2301110030300010-1010001121021323-2203031022222212-3032101222133130-1210020310230212-0222030200310211-0010301131321221"></a>

#### `tgw_security.active_forward_proxy_policies.forward_proxy_policies.name` property

Type: `"string"`. Optional.

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

<a id="canonical-2001130010322213-1121312320023202-0213213032323233-1222113322301333-2111020000232333-0313311213231333-3033220000002221-3010101312111130"></a>

<a id="canonical-0003201213010301-0212103331231123-2123130320313222-3100111112033310-1220102111132131-1022013232112313-2002312232323231-3330300132300322"></a>

#### `tgw_security.active_forward_proxy_policies.forward_proxy_policies.namespace` property

Type: `"string"`. Optional, Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2330230101002121-0102021221322121-0313131233330030-2001221111323223-2001021222112020-0323212320312022-0222120303123231-0122011012311101"></a>

<a id="canonical-2023311102221231-3002013212210312-1012010332121220-2303002312003002-1132201010311310-0312022121032103-1033111321131023-3100212320310022"></a>

#### `tgw_security.active_forward_proxy_policies.forward_proxy_policies.tenant` property

Type: `"string"`. Computed.

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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3313010131101201-0330013203223302-3310032201223232-3120232120111023-1323300221122300-1111001213032333-3033303213023132-0113333122021120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tgw_security.active_network_policies` properties

Breadcrumbs:

- [xcsh_aws_tgw_site](../resources/aws_tgw_site.md#canonical-2230112000323023-1232003130310011-1300013003021130-0212200032323001-1231220220333120-0203201112201213-0311321123312113-1121110002010112)
- [Property reference](resources--aws_tgw_site--reference--group-001.md#canonical-0010030023013322-1222303112030301-3301303233332013-0103300311223300-2202230101213131-0302322120010320-1030223332212332-3103201322012110)
- [tgw_security](resources--aws_tgw_site--reference--group-002.md#canonical-1133330010300121-0303232012001232-1332311230303120-2211233112111030-1123201320011012-1210113320203131-3032233221100133-3213002021031030)
- tgw_security.active_network_policies

<a id="canonical-1013300123232013-1133231232230201-1233133330301023-2223323200222110-3313312302312022-2032021311033220-1213111133330030-1003213120311330"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for active network policies.

Additional upstream details:

List of firewall policy views.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("network_policies")}
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
active_network_policies {
  # Configure direct properties listed below.
}
```
