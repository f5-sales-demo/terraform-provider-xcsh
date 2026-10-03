---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0012322333200122-3323200201110013-2212013303330113-1231312311111120-2300123232323301-2202132023320310-3111213132312212-0233110210321312"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_endpoint / 120233300132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-1031210111331133-0200303110000110-2131002230010332-2201331032221231-0311233303112220-0123031220131300-0311132013302100-1203332112130000"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-1210322212323302-3320331310200313-1112002131233332-1030130203021021-1033220302121230-3323220132123303-1010000322200101-0222011333302012"></a>

## Direct properties — api_endpoint / 120233300132 / 3

<a id="canonical-1321020230013021-3223131020322233-2333112102131220-3301113331103021-3023303211210223-1130323121310233-0131031030103011-1113303100223011"></a>

<a id="canonical-3100323033201211-0210002002002202-0203210313212111-3121110100033033-3111010333020021-1202222033023320-0322310231000022-1210222131332220"></a>

## methods property — api_endpoint / 120233300132 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0023021233030202-0211131311110221-0011311331201131-3010100302021302-3220320232020200-0112323230103023-0302113211001201-3213110322003013"></a>

<a id="canonical-2010201200001232-1133200032103331-0300203222033303-2123333123011220-3203201023323031-3222221103320220-0032310231002121-3310030132332022"></a>

## path property — api_endpoint / 120233300132 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-2133100213333232-0202032123212011-1100012130001113-0132021131030102-0033100331300131-1312213020020031-3111233110222223-1313000010111011"></a>

## Next pages — api_endpoint / 120233300132 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3122022232220301-3022311110311231-2132233022101212-1033021212201102-2302313211003333-2321332322111311-2313032223301201-3112030031020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1221211101322310-3313203320000013-3032033312233003-2123102100013202-2201111232223301-3220210123313023-0320222213200210-0213012002203121"></a>

## api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — metadata / 110322231022 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-005.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-005.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-3022100003011012-1323233332033103-2302222333303300-0333211311333202-0012113310200230-1121102012223320-0002131200203300-3033213032203100"></a>

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

<a id="canonical-3220023110120333-1133022321332232-2110013213002231-2212132010323103-2321120123333133-2000011032230231-2000121312001010-2313310332211333"></a>

## Direct properties — metadata / 110322231022 / 3

<a id="canonical-0132131323203210-2103101101230321-0230331200113120-3003303123122320-0121221031211113-0030323303030000-3300202113003321-0033300301102111"></a>

<a id="canonical-0121132011223003-0230310201023320-2311313220012112-1100212313320303-0133313321223011-0021102331101022-3110113203133030-0320103212121312"></a>

## description_spec property — metadata / 110322231022 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-1021030120332302-1322031212121033-3111203123212331-3123323003303102-1231013323131100-0303012222312323-0233023203200110-3203233022023023"></a>

<a id="canonical-0212031310211213-0233320311131300-2201332222132312-0013211100331120-3000100303310102-1312330220120103-1102113310212202-1012330031031320"></a>

## name property — metadata / 110322231022 / 5

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

<a id="canonical-0121021220100332-1133331310212013-2020100321113203-3101331210311220-2021110012203103-0330200103001030-0212323031002221-3123211002031230"></a>

## Next pages — metadata / 110322231022 / 6

- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-005.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1010200121332300-3032033021323220-1232100331223301-2113002213122320-0101100233130001-1002001013313230-2022323102001101-3323212120220030"></a>

## api_specification.validation_all_spec_endpoints.settings — settings / 003312200110 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-1320313203323313-1103210203113131-3100310201103302-0321312031333031-0302311212113213-0000202102021201-2333131330112010-1223100033232311"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
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
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-3010111000230200-3232203013221322-1122222230002103-1121313212020002-1110021120212110-0020230020123210-0132213302030131-0030100230330100"></a>

## Direct properties — settings / 003312200110 / 3

- [oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-3111022133001320-3223203100220113-2132120223213223-0121022300023213-3323310102031120-1012231201032201-2000013030213320-1033210130113012): complete subsection reference.

- [oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-2301103213103110-2203331121123311-2101333100120313-2001022332232222-0233022103033123-2203312212330222-3111001020323133-2020230101022331): complete subsection reference.

- [property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123): complete subsection reference.

- [property_validation_settings_default](resources--cdn_loadbalancer--reference--group-006.md#canonical-3132303220021203-3111022333231003-1021031000232103-2010022011210210-0330232001321100-0310002113322031-2013010013000102-0313311100120131): complete subsection reference.

<a id="canonical-0001111120122210-1232103301232322-2212110230322323-0220112320332002-3220011323031033-0122122132103133-1330120300030232-0203201303102203"></a>

## Next pages — settings / 003312200110 / 4

- [api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-3111022133001320-3223203100220113-2132120223213223-0121022300023213-3323310102031120-1012231201032201-2000013030213320-1033210130113012)
- [api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-2301103213103110-2203331121123311-2101333100120313-2001022332232222-0233022103033123-2203312212330222-3111001020323133-2020230101022331)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default](resources--cdn_loadbalancer--reference--group-006.md#canonical-3132303220021203-3111022333231003-1021031000232103-2010022011210210-0330232001321100-0310002113322031-2013010013000102-0313311100120131)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3111022133001320-3223203100220113-2132120223213223-0121022300023213-3323310102031120-1012231201032201-2000013030213320-1033210130113012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1111310323001110-2223311033333013-2303130321123222-1230320210030110-2022131021012332-1022332132132012-1112100020220312-0201030203201311"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation — oversized_body_fail_validation / 103323232322 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-2000331231122003-1302312233102022-0203133320330100-3033313110210312-2020003111100131-3202320003032231-3310313113030130-3221121010121313"></a>

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
oversized_body_fail_validation = {}
```

<a id="canonical-2120031031211021-1303020222033110-1223202313001103-0211100103013132-3112211221330302-3033320233032130-3211012132220121-2103033103203233"></a>

## Direct properties — oversized_body_fail_validation / 103323232322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1200233233311213-2231111033320112-2003230120002311-0300021133330132-2132203133021313-0102122211331020-1333313331230210-0111332130032112"></a>

## Next pages — oversized_body_fail_validation / 103323232322 / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2301103213103110-2203331121123311-2101333100120313-2001022332232222-0233022103033123-2203312212330222-3111001020323133-2020230101022331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130003200113013-1020012313102012-0030210300020210-0031201310132311-2322122322032213-0100010003202132-0220223211300121-3013032130020220"></a>

## api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation — oversized_body_skip_validation / 021213333200 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-2133103022203003-2030130112201100-2222000332023203-0022021333002332-1133103001222223-3223222011212303-1321120330023122-1310331323131300"></a>

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
oversized_body_skip_validation = {}
```

<a id="canonical-2111010111012013-3212230220122123-1123003000321313-1022130000222032-1331133101233210-3012033101032130-3310120031331031-0102131110321020"></a>

## Direct properties — oversized_body_skip_validation / 021213333200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1211011113210330-0303231032323310-1103021321323313-2211032333022233-0200303032033011-0022202020210103-3031131210003110-3201213311212100"></a>

## Next pages — oversized_body_skip_validation / 021213333200 / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2213001201201303-3323022133012111-2032233200332103-3221332222202230-2003013103332120-0210101210320131-3011003120322032-3213000331130123"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom — property_validation_settings_custom / 132233211303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-2113102202220010-3223020112323301-1132132321201102-2220220032303222-1012300012113112-3210120220033133-1222012312302101-0113113113000133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
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
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-3313303330302020-1023010210201331-2211320023113302-0110233122330332-0233213231332222-2213121221321111-1223103110002220-0202002323112121"></a>

## Direct properties — property_validation_settings_custom / 132233211303 / 3

- [query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123): complete subsection reference.

<a id="canonical-3102132112330033-3113003031300210-0233203311001301-3023232211111210-1200322113313322-2301033021233120-2012302022031201-3220211102331001"></a>

## Next pages — property_validation_settings_custom / 132233211303 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102120202003210-0130220211000302-2031010113201000-0223203201103131-2213322303231330-2223332332123211-0203223333132030-0003121310023032"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters — query_parameters / 311102013020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-2020211132311321-0200030300033210-0001311323312300-2003123201032033-1132230032300002-2211330123232321-0231110312220302-1132030323231202"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2101212321221232-1003211010212230-1202100113113100-3213200010000311-0321213213310131-3112210333233033-3121320231320332-1100110232323101"></a>

## Direct properties — query_parameters / 311102013020 / 3

- [allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1013022311302032-2111213201323130-3022332330331330-0121201102302332-2133331121100022-3030002310123122-0103211303132133-0033133312011322): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-2330332323102122-0120133033102210-2213223331202203-3132333032033302-3232123202001321-3112313200320011-3033211322013330-0223133101203013): complete subsection reference.

<a id="canonical-0023200202330330-1321200332131222-2200323300223101-1313010210313023-1230201303212103-0223211021212233-0023302212101323-2212023300232210"></a>

## Next pages — query_parameters / 311102013020 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1013022311302032-2111213201323130-3022332330331330-0121201102302332-2133331121100022-3030002310123122-0103211303132133-0033133312011322)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-2330332323102122-0120133033102210-2213223331202203-3132333032033302-3232123202001321-3112313200320011-3033211322013330-0223133101203013)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1013022311302032-2111213201323130-3022332330331330-0121201102302332-2133331121100022-3030002310123122-0103211303132133-0033133312011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232012101113303-0222203313231003-3003300303232100-2230202312230132-2133201222122002-3232121200211332-3213333311120310-1300321223231301"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters — allow_additional_parameters / 003213000212 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-3002201113213032-2103023032232321-0322133000001211-2032222320131000-0323321111303112-1311113033020012-2100202321133332-2103301331103010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

<a id="canonical-3113302223103301-2221032132002001-2031210223003111-3121033232221130-3133332103032331-1110010121132222-3123321213202020-1201333010330302"></a>

## Direct properties — allow_additional_parameters / 003213000212 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0110332133112000-1302131122212011-3331330130101321-1032020120123203-1020322021322110-2100201100121322-2033023202323212-2121013021220211"></a>

## Next pages — allow_additional_parameters / 003213000212 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2330332323102122-0120133033102210-2213223331202203-3132333032033302-3232123202001321-3112313200320011-3033211322013330-0223133101203013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210300200211201-0312111012222021-0020121202213020-3101013121123010-3133012203133000-2023233002130120-2313010200232310-2112210132010202"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters — disallow_additional_parameters / 222110113130 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-2111113301120013-3210223002030322-1220220121102020-3230202233230221-1011011321132110-0312222333212323-0131313213313111-0202201030223210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

<a id="canonical-1021022031310011-1301220320202013-2123201113012222-3110011221310120-0112212310333110-0020211200112122-1021312121100211-1210000233111031"></a>

## Direct properties — disallow_additional_parameters / 222110113130 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3121220032031000-2301203130113003-0133000211101230-3102321112203033-3302130332323302-1203131220233313-0313012132111213-1202311122110102"></a>

## Next pages — disallow_additional_parameters / 222110113130 / 4

- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3132303220021203-3111022333231003-1021031000232103-2010022011210210-0330232001321100-0310002113322031-2013010013000102-0313311100120131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102003012312101-0002323233233211-1123232211230121-3023211010120323-1012332023320133-2101211313011212-2020003311021013-2211002122102222"></a>

## api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default — property_validation_settings_default / 313123113103 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-0121001130112121-1123310201321002-3010313233011000-2122223310301223-0232210002130131-1220220311033322-2011202133021101-3233311200210321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for property validation settings default.

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
property_validation_settings_default = {}
```

<a id="canonical-3003230203001102-1132202003123210-2033202032230120-0032301311001313-3032320312100301-0203001333100011-1133132023101103-2322121312301023"></a>

## Direct properties — property_validation_settings_default / 313123113103 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3010003322323213-2020213232231132-1223133031312112-2320033130202310-3322102120203221-3311133112022120-0312003210012232-0311002133230012"></a>

## Next pages — property_validation_settings_default / 313123113103 / 4

- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102130231110201-0213113322003210-1213102012202001-2002012113120232-0100030001320333-2032020122222213-0010002101103120-3022102220231321"></a>

## api_specification.validation_all_spec_endpoints.validation_mode — validation_mode / 122231332102 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-0212123003212312-3111322333000111-0100031323100101-0333202101202133-0022030312202011-3302013232312220-2322322130220221-0232220330303011"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200311130003102-1003213000113121-1122122122210221-3321312112333122-3121233032031230-1302222312000230-1303313101221230-0233233011310033"></a>

## Direct properties — validation_mode / 122231332102 / 3

- [response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-1302321003220022-2132011332213003-1322331023203221-0200233011032321-3231101312030110-2301032211311103-3333123113230022-1231203132302101): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-0112002210031102-2022111230011230-3120110121131000-2103131301230230-1222213332311232-1212013223033133-2120312023301111-0022301120310203): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212): complete subsection reference.

<a id="canonical-3102031021030333-1030113222133233-1030301312102213-2110020223212020-3103323220211313-2221113301331011-2331131023032132-1113320011311202"></a>

## Next pages — validation_mode / 122231332102 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-1302321003220022-2132011332213003-1322331023203221-0200233011032321-3231101312030110-2301032211311103-3333123113230022-1231203132302101)
- [api_specification.validation_all_spec_endpoints.validation_mode.skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-0112002210031102-2022111230011230-3120110121131000-2103131301230230-1222213332311232-1212013223033133-2120312023301111-0022301120310203)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2311312110011032-0010010103333223-2132111333302112-3323331121012103-0103000212313131-2201121302230121-3111322010130313-0101221012310202"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active — response_validation_mode_active / 211123231301 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-0313110013102010-1333023120312020-3103330001021021-0123310030120020-3113223221031302-3220232102133210-3202220120021203-2103313303102033"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-2311021210212213-1223333122212023-0132232320020313-2020333121302313-3300023323303313-1311121101132013-2210330201123323-1023210100030120"></a>

## Direct properties — response_validation_mode_active / 211123231301 / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-0210302333310310-1231311212222012-3133320013101221-1121211230120001-1130311210323112-0033122313133313-2323310223321013-3033132303130320): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-3203121032033322-0331210220132112-1331101131310332-1013110330220131-0231221000023222-3310010330112231-3223210331020010-0221120001021321): complete subsection reference.

<a id="canonical-1223112111200313-2202020311313002-0303212121112201-0122122203030200-2120131201330023-3213323012112013-0120003301010210-3000203331101011"></a>

<a id="canonical-2000111222012003-1112302121323331-3101121010310113-2103030122002112-3012010002203000-0221321302311221-2001330311001230-3333203330302223"></a>

## response_validation_properties property — response_validation_mode_active / 211123231301 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1002132111113132-1200220023032331-2002223130112311-1312330310012132-0020110103211213-2003112100210323-3013011133322310-3313311133112230"></a>

## Next pages — response_validation_mode_active / 211123231301 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-0210302333310310-1231311212222012-3133320013101221-1121211230120001-1130311210323112-0033122313133313-2323310223321013-3033132303130320)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-3203121032033322-0331210220132112-1331101131310332-1013110330220131-0231221000023222-3310010330112231-3223210331020010-0221120001021321)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0210302333310310-1231311212222012-3133320013101221-1121211230120001-1130311210323112-0033122313133313-2323310223321013-3033132303130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212232020111021-2230210021312203-1310002021303011-2023300231133323-3301203103200010-0010032312133001-3321002230332123-2303020021133012"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block — enforcement_block / 310330202122 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-1033220230221311-2210211030233302-0101333011020233-1232130100223033-1303031230333200-0030300103303030-0212032321222010-2303303213122120"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-3223230202321233-2033223021331113-1031222310303232-1110133113331132-2213230210331302-1023330212132321-0202331212033301-0220130233000020"></a>

## Direct properties — enforcement_block / 310330202122 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0300311200201133-1112230303120021-0013123032121033-0223312212002031-1111223220322233-1212131331233121-1030000000302211-0320213230112133"></a>

## Next pages — enforcement_block / 310330202122 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3203121032033322-0331210220132112-1331101131310332-1013110330220131-0231221000023222-3310010330112231-3223210331020010-0221120001021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1113221232210230-0331031112331121-3332020031101020-3010031030002013-3001300310222300-3300233233111322-0002002131231302-3213013013232131"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report — enforcement_report / 333100112303 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-2310032200131332-0231120122120330-2131233230330101-0322320122100221-2030323333030021-0223133313300103-3322100121322012-0012033130112320"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-1301010231110022-3213031102332130-1023020323323021-3313020230333323-2310223020312321-1332121111123200-2020033302123112-2322202010202013"></a>

## Direct properties — enforcement_report / 333100112303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1021222311312110-3012230230122312-0311010123221232-1003132210131112-1031300302211120-3031112021111012-2222002002112122-1313211131030230"></a>

## Next pages — enforcement_report / 333100112303 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1302321003220022-2132011332213003-1322331023203221-0200233011032321-3231101312030110-2301032211311103-3333123113230022-1231203132302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121121031333201-2310302223123301-2322001013001213-1300310202101321-1032110322212313-2302222012132322-3322333230202201-0330230232303331"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation — skip_response_validation / 220121003031 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-2020330211323302-1212220132302321-1000202232332102-1032002133210301-3132202111222001-3011121002233312-1010302233213232-2332022102012213"></a>

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
skip_response_validation = {}
```

<a id="canonical-1130130213332013-3320302112123021-2011313032023121-0000111311323210-3210030230031231-1111301211323211-3112333012201220-3113312211320203"></a>

## Direct properties — skip_response_validation / 220121003031 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112112230010332-3111011331002320-3310023132313032-3022131020323300-0202001023302021-2212002221003002-0221032030231200-0011222321233110"></a>

## Next pages — skip_response_validation / 220121003031 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0112002210031102-2022111230011230-3120110121131000-2103131301230230-1222213332311232-1212013223033133-2120312023301111-0022301120310203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101103111221111-3103210020031212-2320132230013220-2231101012023101-1101020010021323-2210033003302222-0010001120100012-2133122021112030"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.skip_validation — skip_validation / 323030112121 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-3010221033031110-2201311110233213-0331303222132220-1110132223230031-3020133222300203-1032203032310203-1333010221031121-2312030300320313"></a>

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
skip_validation = {}
```

<a id="canonical-3120200110102111-3200031020332320-0132130211300013-3311322011122311-0233213212030323-3020020232011033-2103323213300313-2311222012231113"></a>

## Direct properties — skip_validation / 323030112121 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0333033200112012-3200122230001110-2321223302200301-3323301313002232-0321003130120211-2111013313123200-2220230302032013-3111021100221320"></a>

## Next pages — skip_validation / 323030112121 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2111003003330133-1203001201130323-0112131331203003-0321003331011210-2120332102220302-2012011020230310-1120200023032321-1202022003023230"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active — validation_mode_active / 230201020331 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-0102002131231210-2201312221102313-0031300300001230-3310233320331011-0323112032311103-3211111001030300-1123011232302300-3212303203210022"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-3221213010110123-1332332110210332-0123310313210120-0300100332130220-1211123012021112-0321211131022232-1123013110330021-2102020022003102"></a>

## Direct properties — validation_mode_active / 230201020331 / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-2002101302303310-2002022212330203-0200313120330031-1003021002320023-3320122113202302-3330200310101101-3202032302201110-3101301002221000): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-1032133120312121-0112201100312123-3320232011311112-2200313330321023-1112021031102013-2021010202102212-1311223020003020-1122003103112201): complete subsection reference.

<a id="canonical-0032223130303122-0131203333001131-3132320012313320-2300033233233203-3001313130221201-3123132302133120-0320110132013000-2100303222321022"></a>

<a id="canonical-3030200200111320-3320020100332311-1203122021230211-2312221233013303-2022210102220312-1001003233111021-2120203320110320-1301223221201011"></a>

## request_validation_properties property — validation_mode_active / 230201020331 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3133123210001322-2132303331121113-3003132300003112-1231022103223202-1223212103021302-3010020213201012-3302223300300133-3222230231030030"></a>

## Next pages — validation_mode_active / 230201020331 / 5

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-2002101302303310-2002022212330203-0200313120330031-1003021002320023-3320122113202302-3330200310101101-3202032302201110-3101301002221000)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-1032133120312121-0112201100312123-3320232011311112-2200313330321023-1112021031102013-2021010202102212-1311223020003020-1122003103112201)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2002101302303310-2002022212330203-0200313120330031-1003021002320023-3320122113202302-3330200310101101-3202032302201110-3101301002221000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022013232010330-0111313322322233-3332212113120012-1220223111033332-2102121211002330-1201331320110321-3121201110121313-1100103021111221"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block — enforcement_block / 130013001013 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-3201321012011100-1313032113032021-3023020003122312-3132213010233300-1011330310010301-3003331110131000-0330212232303000-1320122223302212"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-1310322201330110-3123102212211103-0101221003120213-1310220030013313-1003130000313232-3111033130332111-3313233202231330-3002002031100213"></a>

## Direct properties — enforcement_block / 130013001013 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0213010113202011-1031312032033330-3013001220301102-2213010211031320-2030311300201120-3312220213101131-2203102200303033-0013200323232020"></a>

## Next pages — enforcement_block / 130013001013 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1032133120312121-0112201100312123-3320232011311112-2200313330321023-1112021031102013-2021010202102212-1311223020003020-1122003103112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1313300330202321-0001202230221331-2333312030110311-3311133233323032-3230211203122312-1022222321333233-3013212220120200-0023323213123131"></a>

## api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report — enforcement_report / 302331322023 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-005.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-2131222230313330-3001032320011013-0023303110213322-3010321022022111-2300201113232312-0023213002122130-0230201203331120-0210211320303200"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-0033212112313223-0013101213221323-2300321122302111-2331030113000120-0221301302131031-3322310212023020-2123323301112132-3300030212301221"></a>

## Direct properties — enforcement_report / 302331322023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112003301311230-3000301323201230-1112121321200030-1122321222212113-1222210221201301-3000122210102133-1222012310030032-0111310321112121"></a>

## Next pages — enforcement_report / 302331322023 / 4

- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1333331023331133-3201331211121323-2220203221033121-2320322131302321-3303311212313001-0201330203333113-0302020122023332-0211321311132023"></a>

## api_specification.validation_custom_list — validation_custom_list / 200300232223 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.validation_custom_list

<a id="canonical-3022302121112320-1202230322313102-0000211002111233-3010030103323121-3200132321223200-3333011321230303-3123023212032332-0112203121012131"></a>

Type: `"object"`. single nested block, Optional.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Upstream description:

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to "Fall Through Mode".

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
  "x-ves-oneof-field-oversized_body_choice": "[]"
}
```

Terraform syntax:

```terraform
validation_custom_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-2013330300123312-2222231002031101-1220213030000203-3130010103231030-1333332032310020-3120113101011021-3201221120020312-3212112112302200"></a>

## Direct properties — validation_custom_list / 200300232223 / 3

- [fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230): complete subsection reference.

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031): complete subsection reference.

- [settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233): complete subsection reference.

<a id="canonical-3210003203003222-3301013032003223-2030300013133302-3201313211313111-0321230213200233-1003021030003013-2111120000001231-3312000213233030"></a>

## Next pages — validation_custom_list / 200300232223 / 4

- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3131021200133122-2200101300013003-1323001330302011-3201102012301033-3113321102113031-3132303032112311-1111020300132030-3213233231023121"></a>

## api_specification.validation_custom_list.fall_through_mode — fall_through_mode / 311002121010 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-3220320111130013-1131320303110230-1232233210002021-3331031020101301-0111333122230023-3012212031133303-0210031101011131-1102132011231312"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

Upstream description:

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("fall_through_mode_allow",
    "fall_through_mode_custom")}
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
  "x-ves-oneof-field-fall_through_mode_choice": "[\"fall_through_mode_allow\",\"fall_through_mode_custom\"]"
}
```

Terraform syntax:

```terraform
fall_through_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-2103222300000123-2121333213322320-0100013213212120-1033101311231301-1222023001020312-1112131331322021-3203130121112333-2111330301133020"></a>

## Direct properties — fall_through_mode / 311002121010 / 3

- [fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-006.md#canonical-0111330000221003-2203010113323011-2331020202021302-1211021301330100-1102213021323001-0012203131213231-2322320100332223-3003130310321233): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211): complete subsection reference.

<a id="canonical-2302313211123232-1322232232303302-2303331113311202-3202321232013301-0032132312232023-2303332111100211-1101302312322110-1233322222332322"></a>

## Next pages — fall_through_mode / 311002121010 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-006.md#canonical-0111330000221003-2203010113323011-2331020202021302-1211021301330100-1102213021323001-0012203131213231-2322320100332223-3003130310321233)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0111330000221003-2203010113323011-2331020202021302-1211021301330100-1102213021323001-0012203131213231-2322320100332223-3003130310321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2100310030001102-3332023122003220-1320301121213310-1330033312123120-3130303231222200-2321103011123203-3220323103011022-3031132122222222"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow — fall_through_mode_allow / 031110311012 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-3201203110201322-2202233220223123-2133020111010311-2103202003323233-2202302003020200-1110212001122213-0221100332133212-2232202110100223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fall through mode allow.

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
fall_through_mode_allow = {}
```

<a id="canonical-3032222000010131-3021000223321221-2022122130122100-2232223023111212-3120233132201003-3332303012032012-1203120312101100-1232231001222212"></a>

## Direct properties — fall_through_mode_allow / 031110311012 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112203221102132-2000302012022321-1313032000221222-1321201033211233-2022300032322010-3000002120132110-0310233132333332-3203223003021323"></a>

## Next pages — fall_through_mode_allow / 031110311012 / 4

- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1320100310230212-0222131021103302-0032230223303203-0231201221323320-2302210313221302-1331131032133331-2330212221122320-1200123302223331"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom — fall_through_mode_custom / 102003302233 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-0102303111101020-2021131333202322-0123200220001111-3210112213322232-1200201333022202-2120322311333000-0031311101330022-1321210203303100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Upstream description:

Define the fall through settings.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("open_api_validation_rules")}
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
fall_through_mode_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000020202202320-2233220301011021-3320330311222110-0213002120023300-1311322012330011-2311220301022003-0330101310003312-3003130211333320"></a>

## Direct properties — fall_through_mode_custom / 102003302233 / 3

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033): complete subsection reference.

<a id="canonical-0110233212002323-2322003322130110-1123010012123220-2102010130222302-2112012032021213-0330220313301031-2212320321012300-2301013302313032"></a>

## Next pages — fall_through_mode_custom / 102003302233 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0002233331113020-0200103221221031-0111103130220033-3013221212202100-2023211200112112-0210012221103200-3112121211103233-1112232021322012"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules — open_api_validation_rules / 011300223210 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-2212211330233221-3332013323112111-1311123231301120-2211032130333322-0313303132303220-2322123303000212-0332201330220123-0312123102123100"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("action_block",
    "action_report"),
  validators.ConflictingListObjectAttributes("action_block",
    "action_skip"),
  validators.ConflictingListObjectAttributes("action_report",
    "action_skip"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-1311233123211123-0110130332232130-1332131003130332-0201103330232323-0121301031031120-1101120123303103-0133230201210122-0213201122113230"></a>

## Direct properties — open_api_validation_rules / 011300223210 / 3

- [action_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322013211132302-1133110233011131-2122302333323121-1300221001133133-1103232111213310-1302222031022131-2012103032331000-3130001320022020): complete subsection reference.

- [action_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-3002032202300322-1221033310333121-3000132021223031-3330231031223320-2202303121120031-3111300223101301-2122030310202300-1130001210313121): complete subsection reference.

- [action_skip](resources--cdn_loadbalancer--reference--group-006.md#canonical-1203322032320111-3011233302330222-1131332022012100-1012100013203333-2313322113323132-0020121130230113-1000111322032300-1033002032313110): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-1233011312013203-2333000201303113-1002101013330021-2032021012133213-0230121310211332-0110230321213010-2332131330121310-2031132100120311): complete subsection reference.

<a id="canonical-2011032032310232-1120121020333233-0311303212322133-1001310223222330-0021002310031322-3303021131212300-1332200032023213-2011012202101222"></a>

<a id="canonical-1021101123231321-2210101022203131-1330233121020131-2003320201203330-3210010220210113-2003221332333233-2100023331123101-2310002131310032"></a>

## api_group property — open_api_validation_rules / 011300223210 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-1222313320221101-2320332201012333-3300230230101100-0022130102213011-1333133313331130-2011313021212100-0010201223321200-0011003030002003"></a>

<a id="canonical-1203011330310110-1220321201112333-2003203122323333-1230122312112313-3201121210103103-2112112310010012-0301210130222133-2312100230322102"></a>

## base_path property — open_api_validation_rules / 011300223210 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-2213100130022110-2203122121211011-0032211321331011-2102313123013321-2322012022210132-1112211000320131-1223302321120220-3011012120122131): complete subsection reference.

<a id="canonical-2102301333211233-1000322122231223-3103022130103131-3132231010020032-2221211312123011-1210020112133123-0020002030233321-2333212102230013"></a>

## Next pages — open_api_validation_rules / 011300223210 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-0322013211132302-1133110233011131-2122302333323121-1300221001133133-1103232111213310-1302222031022131-2012103032331000-3130001320022020)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-3002032202300322-1221033310333121-3000132021223031-3330231031223320-2202303121120031-3111300223101301-2122030310202300-1130001210313121)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip](resources--cdn_loadbalancer--reference--group-006.md#canonical-1203322032320111-3011233302330222-1131332022012100-1012100013203333-2313322113323132-0020121130230113-1000111322032300-1033002032313110)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-1233011312013203-2333000201303113-1002101013330021-2032021012133213-0230121310211332-0110230321213010-2332131330121310-2031132100120311)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-2213100130022110-2203122121211011-0032211321331011-2102313123013321-2322012022210132-1112211000320131-1223302321120220-3011012120122131)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0322013211132302-1133110233011131-2122302333323121-1300221001133133-1103232111213310-1302222031022131-2012103032331000-3130001320022020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1101011303311222-2033320101230300-3230232200113110-0311021110310310-2231022021232001-0301213232122112-1112000022220112-3020133300230313"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block — action_block / 301210102231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-3002132212102200-1011131031103021-2210222203123232-3231333120230101-2320210200113311-2330033232100231-3213033122311221-0302300021312003"></a>

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
action_block = {}
```

<a id="canonical-1110023132333031-2111220133320000-2213201133312111-2132000130123231-3333110033123223-3222100203103213-2031212332322131-3033210013232003"></a>

## Direct properties — action_block / 301210102231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2023301310210122-2123311032201313-1220223222310201-3122030023320311-2333213212311123-2012320320013121-1131300223323130-0210220011031210"></a>

## Next pages — action_block / 301210102231 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3002032202300322-1221033310333121-3000132021223031-3330231031223320-2202303121120031-3111300223101301-2122030310202300-1130001210313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1123310233001021-3312310030213010-0122101223110130-2011320233202222-3200322311231223-0222030202030211-2230100303333113-2020220333103211"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report — action_report / 320220312020 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-0101031232130033-3233321220312300-0102200201111233-1012033110113110-1023131032130023-3201013203013021-2310203103132111-2321233212101000"></a>

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
action_report = {}
```

<a id="canonical-2122232201323001-0002121213131233-1233202021312310-3021211333023000-3233030312331103-1313213330022000-1233100310132003-2001023010323112"></a>

## Direct properties — action_report / 320220312020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333111001321202-1003322122211010-0001020210331132-0131220013233012-1031221313211211-2103113311313120-1110310002223201-1323301310301210"></a>

## Next pages — action_report / 320220312020 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1203322032320111-3011233302330222-1131332022012100-1012100013203333-2313322113323132-0020121130230113-1000111322032300-1033002032313110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1312233130210100-1012010003031321-1021232221220032-1102233203013200-1332130121220112-3303003213013213-3232311231001300-1101122033231031"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip — action_skip / 122223220123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-3010102213113332-2322312230102221-3013032032132230-3231030002231021-1210321332001003-0202221301330311-0210122110230321-2213012031221223"></a>

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
action_skip = {}
```

<a id="canonical-0033202212012301-1103323102333100-3333002100011101-2002033000120011-0013013120322011-2121312012020111-3310233220112112-1232320030121320"></a>

## Direct properties — action_skip / 122223220123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221332030232120-2003131113312100-2230213232010122-1211121202322232-0211303222311211-3311100313122313-3112103201322333-0300322101123311"></a>

## Next pages — action_skip / 122223220123 / 4

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1233011312013203-2333000201303113-1002101013330021-2032021012133213-0230121310211332-0110230321213010-2332131330121310-2031132100120311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1200010333131030-1230200201200130-3200311132113120-2211202233223023-2233133223301321-3013221320213201-1222333130110003-1132122011021311"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint — api_endpoint / 000103030302 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-1301220321310310-0332231021000303-1030331203302331-1232330331333330-1211332332200103-3110302230011301-0312110202123130-3112113203233301"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021110301212301-2310333132100022-3130223101310011-3113330201122311-0311030130210222-3321200030110313-2133113013032012-1132130111321321"></a>

## Direct properties — api_endpoint / 000103030302 / 3

<a id="canonical-1321232122232112-0233202323133022-1320233133211303-0101121321213213-0030013312023230-3021001201221012-3112133210132031-1320302302312200"></a>

<a id="canonical-1321133001002123-1333103210230013-2020122213102331-0322211221023230-2313122130020013-0233201101320010-1313311301103000-2323232230212202"></a>

## methods property — api_endpoint / 000103030302 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2210320231323230-3010123130321102-1130023200330320-0010123230111102-0110331303232130-1202031322310003-0110003313300001-3222232213020203"></a>

<a id="canonical-3333130333330233-0132203211202033-0210100230003113-2111323321311233-2313332310213122-1221220213223202-1010320220013023-2200212213321200"></a>

## path property — api_endpoint / 000103030302 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-1332020220103130-1312331023301212-0130220330103231-0031310113303321-2013313230232122-1112310020101100-2000223312323223-0301132010201122"></a>

## Next pages — api_endpoint / 000103030302 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2213100130022110-2203122121211011-0032211321331011-2102313123013321-2322012022210132-1112211000320131-1223302321120220-3011012120122131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2020023210200322-3111333133303331-2110002020020100-0320323012221031-3123031211033231-3210301202123220-0133002201110203-2211111133323122"></a>

## api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata — metadata / 320203200211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-2321231000031310-3311113013222023-3101321012331311-3122223323222320-3212321203102210-2202333332213131-2131111303123111-0312332310301110"></a>

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

<a id="canonical-1012132221323121-2213113003103223-2032212233131213-2311210233312020-2200031222023301-0312112001010323-2311011112220023-1303323211211012"></a>

## Direct properties — metadata / 320203200211 / 3

<a id="canonical-0011332311122120-2032120303330302-3202122102200133-3203310330330033-1323233020032202-0120111102233101-3303112221232002-2021020300232211"></a>

<a id="canonical-1221211302210132-0121030033111220-2112022203312330-0231012312210223-0100000203213113-3232102121002133-1113123333222202-1233103022002320"></a>

## description_spec property — metadata / 320203200211 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-3330300010310133-3130113333213331-2211231110303101-1223320200021320-3321113312100300-1110200301101123-1120010321200000-1103322231202002"></a>

<a id="canonical-2102202132310021-1113211301313230-2312103331202221-1302213121122121-3111233322323033-1023330210221030-3012203223020020-1023330022321312"></a>

## name property — metadata / 320203200211 / 5

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

<a id="canonical-3330323330210232-3300210220300123-1110210112220110-2333233231102013-2313223030310100-1002203231213202-0312323012301230-0331311302100123"></a>

## Next pages — metadata / 320203200211 / 6

- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213003232321320-3023013000213032-1033033200322000-1203322010013313-0310332002220121-2130322000202221-1101310203002131-0120002121022100"></a>

## api_specification.validation_custom_list.open_api_validation_rules — open_api_validation_rules / 201320003100 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-1000333221022003-3203031023202323-0213011233002102-3133310303321200-1033313231013213-0200013300133001-1020223221213011-0331220201311022"></a>

Type: `"object"`. list nested block, Optional.

Validation List. Rule or policy definition

Upstream description:

Rule or policy definition

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.ConflictingListObjectAttributes("any_domain",
    "specific_domain"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "api_group"),
  validators.ConflictingListObjectAttributes("api_endpoint",
    "base_path"),
  validators.ConflictingListObjectAttributes("api_group",
    "base_path")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 15,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 15,
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
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "15",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
open_api_validation_rules {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111012123312201-1013013300301211-3233222020122302-3020201201033033-1032331001312231-2202222312322120-1331022000101321-0200312103323113"></a>

## Direct properties — open_api_validation_rules / 201320003100 / 3

- [any_domain](resources--cdn_loadbalancer--reference--group-006.md#canonical-0001121131003212-2001030100132010-2120213033231312-2231023013230233-2021220102103203-0311221322133130-1111200202331103-2022301000301102): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321322200122101-3122230112323122-2322232020122330-3021303320220130-0311321223333111-3211220010202303-3322110010122300-2023331302112003): complete subsection reference.

<a id="canonical-1220330121202312-1201222230030123-0221112220102011-0013222300201203-0213332111000210-2213102130320120-3223010003223120-2221332301333003"></a>

<a id="canonical-1112033121200113-2313203313023000-1321131011202001-0210033231332211-0030222230022122-2313013013331211-3300002023120202-3303213230322232"></a>

## api_group property — open_api_validation_rules / 201320003100 / 4

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

<a id="canonical-2232002003002303-3302333100222333-1330002230311113-2331130130203322-3221221202020203-0033022211200120-3321211303022331-0112200101032023"></a>

<a id="canonical-1311010313231112-1012200310122013-1211001132213103-3203232332132120-1332010003311013-1111310032030333-2220013321313203-0321031031022023"></a>

## base_path property — open_api_validation_rules / 201320003100 / 5

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Upstream description:

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "128"
  }
}
```

- [metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-3031030312110130-1313002120122221-2123030002001220-3112030313120103-0200330232322123-2312322201013110-0200012303132033-1000212020012300): complete subsection reference.

<a id="canonical-3311103320132221-0231202233323320-0121123110321110-1331031311310102-0311213300233213-0223000201112030-0030313320110122-3302132102123320"></a>

<a id="canonical-1220011301330132-2223123221200101-1320320123122320-1102231320331001-3300333303020310-0101300231012311-3212233302331122-1201300303212022"></a>

## specific_domain property — open_api_validation_rules / 201320003100 / 6

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Upstream description:

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(128),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.vh_domain": "true"
  }
}
```

- [validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032): complete subsection reference.

<a id="canonical-0211331332300011-3033322200032130-2120011032212003-1221312321221230-0133031022001110-2032133233120103-1132330313221002-3021223220131023"></a>

## Next pages — open_api_validation_rules / 201320003100 / 7

- [api_specification.validation_custom_list.open_api_validation_rules.any_domain](resources--cdn_loadbalancer--reference--group-006.md#canonical-0001121131003212-2001030100132010-2120213033231312-2231023013230233-2021220102103203-0311221322133130-1111200202331103-2022301000301102)
- [api_specification.validation_custom_list.open_api_validation_rules.api_endpoint](resources--cdn_loadbalancer--reference--group-006.md#canonical-0321322200122101-3122230112323122-2322232020122330-3021303320220130-0311321223333111-3211220010202303-3322110010122300-2023331302112003)
- [api_specification.validation_custom_list.open_api_validation_rules.metadata](resources--cdn_loadbalancer--reference--group-006.md#canonical-3031030312110130-1313002120122221-2123030002001220-3112030313120103-0200330232322123-2312322201013110-0200012303132033-1000212020012300)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0001121131003212-2001030100132010-2120213033231312-2231023013230233-2021220102103203-0311221322133130-1111200202331103-2022301000301102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332210010321210-0320021133023313-0020033123300332-1022032303100102-0313321011013003-3200213022133022-2113333201300121-3212201221123031"></a>

## api_specification.validation_custom_list.open_api_validation_rules.any_domain — any_domain / 331331021112 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-2133112203211020-2333112210031123-0201301113011110-1322111310211033-3003023202132003-2001112311303201-0331100311012301-2332300123133032"></a>

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

<a id="canonical-2312110232102220-1012101101120012-0233033003321112-1202121002103020-3301011031120031-2133130110313313-3320133303103213-2331333103133131"></a>

## Direct properties — any_domain / 331331021112 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1022013113222333-0000023332213023-3223202320112322-3301112112200131-0102122100030103-1032112112231023-3312312011103023-0120311123021030"></a>

## Next pages — any_domain / 331331021112 / 4

- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0321322200122101-3122230112323122-2322232020122330-3021303320220130-0311321223333111-3211220010202303-3322110010122300-2023331302112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3303120122233032-0010331223021102-3110333303000233-0333311003130233-2313132211121222-3213123103011231-2210300011022320-2003022112012131"></a>

## api_specification.validation_custom_list.open_api_validation_rules.api_endpoint — api_endpoint / 102221102132 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-3311331013223222-0213221323310022-0332100000213113-1330220121312223-0123132022032012-2121231121103311-3100130220100120-0200310323000132"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Upstream description:

This defines API endpoint.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("path")}
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
api_endpoint {
  # Configure direct properties listed below.
}
```

<a id="canonical-0022130132133101-1223201332102121-2122300202123233-0102301021020111-2201323301201302-2120221220120132-0030310101222103-1331131203033132"></a>

## Direct properties — api_endpoint / 102221102132 / 3

<a id="canonical-3020121201230312-0030132000102221-2033321013101113-1011003310202222-0123331103302101-2303222303111332-3021103302012002-2313013120011332"></a>

<a id="canonical-2201001011000211-1031310033133023-2330333030031102-1221112131003111-1203303133303030-2100101003030003-3302023200212312-0201102302202333"></a>

## methods property — api_endpoint / 102221102132 / 4

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

Upstream description:

Methods to be matched.

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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "16",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1033200333220300-2202111210223113-0301320222102131-3331220003300212-1030102130022123-2000212013213300-3110032100333202-1233001023202232"></a>

<a id="canonical-3313020232333032-1312130003332303-3203113123301221-3322020230013000-1133303013231211-1111000303312120-2020023201213130-0332021322132211"></a>

## path property — api_endpoint / 102221102132 / 5

Type: `"string"`. Optional.

Path. Path to be matched.

Upstream description:

Path to be matched.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
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
    "minLength": 1,
    "pattern": "^[/a-zA-Z0-9._-]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "1024",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0212123110100101-0210102311003031-0131311021023033-3101313021100112-2032300133033111-0310130202233113-0103332201112101-1321112310102012"></a>

## Next pages — api_endpoint / 102221102132 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3031030312110130-1313002120122221-2123030002001220-3112030313120103-0200330232322123-2312322201013110-0200012303132033-1000212020012300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2212023210122220-1010000231003012-2210103133223332-2210312133012302-2201333221023121-0021330133123123-1121123132022333-2130120312202300"></a>

## api_specification.validation_custom_list.open_api_validation_rules.metadata — metadata / 210332020003 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-1133321121303113-2203110032011330-3123332233220000-0201130210102011-2112320030103313-1111123132213130-3222000201101323-0123310303132000"></a>

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

<a id="canonical-1332120033231022-2321033023030323-2332330233002030-0221103013031112-0211110223321212-0230300003122011-2312201301210302-2030100131301322"></a>

## Direct properties — metadata / 210332020003 / 3

<a id="canonical-3030131301120312-2203300203303001-1203213013313120-0100222323223212-1020211333233000-0202232002111000-1213103133221010-0210200102102210"></a>

<a id="canonical-3111113200000311-1320322213032221-1313121001231322-3231311222032203-1001333033212013-0023200221221213-1013023331212122-2020222201111013"></a>

## description_spec property — metadata / 210332020003 / 4

Type: `"string"`. Optional.

Description. Human readable description.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

<a id="canonical-2330322212232122-2013030133323202-0113321230021121-2220121232003022-1223233232221033-3021021020120212-1313203233013033-0302032110333011"></a>

<a id="canonical-3210120120220012-1100022311101021-2122110313022012-2130030303231003-0013000313002022-3013222121103300-1132302120203312-2213311321101223"></a>

## name property — metadata / 210332020003 / 5

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

<a id="canonical-2302213102330233-2003213221233222-1133313202010121-1310010133211031-3211132013122130-0121111002331113-3232310103022313-2020011200113000"></a>

## Next pages — metadata / 210332020003 / 6

- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1301303232030231-3332123003111131-2002030131310032-1222201112022032-3200111001322020-1011021313301121-0030210010131132-1020020010313023"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode — validation_mode / 223130020202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-1203002200333322-3300331231320023-1222231131102113-0320011111131320-1300200001122133-1110110203133200-2102220222011021-1223113212300123"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

Upstream description:

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("response_validation_mode_active",
    "skip_response_validation"),
  validators.ConflictingObjectAttributes("skip_validation",
    "validation_mode_active")}
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
  "x-ves-oneof-field-response_validation_mode_choice": "[\"response_validation_mode_active\",\"skip_response_validation\"]",
  "x-ves-oneof-field-validation_mode_choice": "[\"skip_validation\",\"validation_mode_active\"]"
}
```

Terraform syntax:

```terraform
validation_mode {
  # Configure direct properties listed below.
}
```

<a id="canonical-1030120232331213-2203210200021332-1202020021230302-3013002200111122-3201222120221222-3123123020230211-1030210120020122-3202220221031300"></a>

## Direct properties — validation_mode / 223130020202 / 3

- [response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-2111302223301013-1313333332013013-2022001332312010-0130033332132132-0000312223312310-0101212030320322-3223001031011123-1033232001320200): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-0131000113200332-1122033201111000-0130010212313212-2013000222113202-2131202033111200-3320112131103022-0132003131013021-2021333231110030): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033): complete subsection reference.

<a id="canonical-0203333113030201-3120123302113210-0210033230022031-1110312003112330-1131132121013321-3222130231113303-3321012123311233-3122220101221113"></a>

## Next pages — validation_mode / 223130020202 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-2111302223301013-1313333332013013-2022001332312010-0130033332132132-0000312223312310-0101212030320322-3223001031011123-1033232001320200)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-0131000113200332-1122033201111000-0130010212313212-2013000222113202-2131202033111200-3320112131103022-0132003131013021-2021333231110030)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202101012213110-3122210100011002-3330020213123111-2323130000203312-1100312321012021-3021102110031321-2001032302212110-1012033012023011"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active — response_validation_mode_active / 012333213123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-1121321330230010-0121220123311032-3220121323221221-1103211321130323-2033230100312111-1112223200322230-1122032120231302-0313033030322011"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

Upstream description:

Validation mode properties of response.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("response_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
response_validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-0031330332030321-0202110032131301-0203113213121301-3320012322133223-3122300120211023-0031032213332120-2020003021230303-2112221023113013"></a>

## Direct properties — response_validation_mode_active / 012333213123 / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-0202300221020321-0102210232001200-2113123030211023-2203320100012211-2220210022000332-1300330203032303-2033220001031102-2223131023303213): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-3102120002033333-1302211001102110-1003100030330131-1021020330330000-2233013002231231-3013020212130012-3322201110230033-2030002132322002): complete subsection reference.

<a id="canonical-2200120102031100-2030122032021011-1113311320023222-3033222110212003-0030322201003222-3130111220331201-0000333323211210-0211212100132023"></a>

<a id="canonical-0011202203311332-2303021323102133-0110031131011303-1103020323100211-2321312122003002-1233020112002031-1122320323223133-1113103320032132"></a>

## response_validation_properties property — response_validation_mode_active / 012333213123 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.in": "[2,4,5,7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3231310321322111-2313302200212101-2301302000322023-2201303031331331-3100330023101231-0130112012222303-0221130330112330-1130103103032021"></a>

## Next pages — response_validation_mode_active / 012333213123 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-0202300221020321-0102210232001200-2113123030211023-2203320100012211-2220210022000332-1300330203032303-2033220001031102-2223131023303213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-3102120002033333-1302211001102110-1003100030330131-1021020330330000-2233013002231231-3013020212130012-3322201110230033-2030002132322002)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0202300221020321-0102210232001200-2113123030211023-2203320100012211-2220210022000332-1300330203032303-2033220001031102-2223131023303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1020031011213123-2002001112030023-2310002013130030-1101320113132222-0203112323122200-1103130112100101-1020122230003312-2201221002133003"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block — enforcement_block / 023212202232 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block

<a id="canonical-2302221223303032-2310032103301003-0123033232200003-3323033022230321-0113312301301013-2100202130210123-3031131333331100-3011322020003213"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-1010030312202233-1222002300211010-2002321132233033-1131333120031301-0112113322331301-0211033332330221-1022303021321100-0231312032311021"></a>

## Direct properties — enforcement_block / 023212202232 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3122210322112302-2230222000301120-1323122313213013-0110311223000212-2021131112333110-2222022020210120-3213130110101003-1330103012031303"></a>

## Next pages — enforcement_block / 023212202232 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3102120002033333-1302211001102110-1003100030330131-1021020330330000-2233013002231231-3013020212130012-3322201110230033-2030002132322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001212310120320-3210333021023221-0121123321032021-3212212202123201-3332232001111120-3132233210023230-1300001130323223-1100033210012120"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report — enforcement_report / 332020101113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report

<a id="canonical-2000212003121210-2310030110301000-0010133111133022-1312122313221300-0303210230312322-0003223011013003-1221001020033033-2323312122013120"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-3210222030222111-1311211030322000-1313300321211231-2212102110332202-1132333122213032-1100201001002332-3123033000332001-0003220331213021"></a>

## Direct properties — enforcement_report / 332020101113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3222233001203312-1210220203012213-0220212200301333-3222131003321012-0102223200003113-0122303310121031-1111130312021220-0330333002232121"></a>

## Next pages — enforcement_report / 332020101113 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2111302223301013-1313333332013013-2022001332312010-0130033332132132-0000312223312310-0101212030320322-3223001031011123-1033232001320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0120203223302003-3001100111233213-1313012303211020-2200102023211331-0022003213312010-3002023211322101-3110321200222003-3232101300002320"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation — skip_response_validation / 000132111222 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-0232023303213233-3302203033103122-0002000303033323-0110323310202000-3332000000231020-0202101010112211-3001331331231102-0201103122031310"></a>

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
skip_response_validation = {}
```

<a id="canonical-2132222131300203-1002012213032023-2221230120132311-1223023210312031-1232203230303112-3221311211322123-2223322111000203-2310123133120112"></a>

## Direct properties — skip_response_validation / 000132111222 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2010222200131322-2331113121131301-3100001312000322-2331132102323301-2232022330100331-2112322233131300-0212221020101220-2033010222212333"></a>

## Next pages — skip_response_validation / 000132111222 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0131000113200332-1122033201111000-0130010212313212-2013000222113202-2131202033111200-3320112131103022-0132003131013021-2021333231110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3102212030132001-3001310222030222-0233311121223330-2111101230012111-0201132012211002-3121001203203021-0002313210331233-2203202002122301"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation — skip_validation / 320332023113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-2310223133223310-3122123020321111-2113021203320322-0200231112212320-2311130213123303-3130000211123331-1010223101211233-2300222211102321"></a>

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
skip_validation = {}
```

<a id="canonical-0131111300232023-2320202210331303-1210221200111100-1330112000000100-0111323212201221-2212311321302021-1312131210310203-1120023321103100"></a>

## Direct properties — skip_validation / 320332023113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2221113111300031-3321331100333310-3113221112023310-3122332202013033-2130032333002011-3322110111330110-2213222330123313-3311123122133031"></a>

## Next pages — skip_validation / 320332023113 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131332031200300-0110200021113002-2313333330133301-3112233332332030-0112123320103010-0231001322122201-0233220101210221-2321302232311110"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active — validation_mode_active / 311202212111 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-2112301023102212-2322120303123032-1203303022311212-2233130310300123-1211231223121211-0020121221032302-3323333123230010-0022001022000221"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("request_validation_properties"),
  validators.ConflictingObjectAttributes("enforcement_block",
    "enforcement_report")}
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
  "x-ves-oneof-field-validation_enforcement_type": "[\"enforcement_block\",\"enforcement_report\"]"
}
```

Terraform syntax:

```terraform
validation_mode_active {
  # Configure direct properties listed below.
}
```

<a id="canonical-3103110030311331-0123102012113221-0332102212231300-2122303110221120-3130332201313110-1002011223223333-2121033322330212-3213301203012110"></a>

## Direct properties — validation_mode_active / 311202212111 / 3

- [enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-2331311033131132-1030322002333001-3031323011200330-0312130102231021-0111000000213010-0233000201100002-0203212311120031-3231133133130213): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-2131232230131012-3331321221102330-3103301030130233-1113321213010313-2101313211331133-0301220131231232-2131313023300033-0131310312011121): complete subsection reference.

<a id="canonical-2020200300220130-3032113222110002-3112022011100100-1321023023011120-1330213010222010-3233033111100302-2220210231121233-2321231203311010"></a>

<a id="canonical-1111222132301203-3232033110013331-0303031132232103-3123322300230231-1013022120233332-1102330310302110-0310002132331200-1031013221232200"></a>

## request_validation_properties property — validation_mode_active / 311202212111 / 4

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

Upstream description:

List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger)

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtLeast(1),
}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[7]",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1222120202200203-3022130110013121-3000113320230020-1031321211302313-2012100231031030-1233023100310002-3022001322200101-2213030302103111"></a>

## Next pages — validation_mode_active / 311202212111 / 5

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block](resources--cdn_loadbalancer--reference--group-006.md#canonical-2331311033131132-1030322002333001-3031323011200330-0312130102231021-0111000000213010-0233000201100002-0203212311120031-3231133133130213)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report](resources--cdn_loadbalancer--reference--group-006.md#canonical-2131232230131012-3331321221102330-3103301030130233-1113321213010313-2101313211331133-0301220131231232-2131313023300033-0131310312011121)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2331311033131132-1030322002333001-3031323011200330-0312130102231021-0111000000213010-0233000201100002-0203212311120031-3231133133130213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2132121330123031-2122200103023023-3212211002132233-2023010310230020-0311110211313032-0023221200323032-0233101311010212-0022120031130322"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block — enforcement_block / 223330133113 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block

<a id="canonical-1003200311010312-1030020300300313-0300101320310012-0011230012303020-1320030122031102-3303121112312033-2130012302230110-2021103221120220"></a>

Type: `["object", {}]`. Optional.

Blocking validation: reject traffic that violates the selected OpenAPI validation properties.
Invalid requests are returned as HTTP 403.

Receipt-pinned upstream constraints:

```json
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
enforcement_block = {}
```

<a id="canonical-2113133132203122-2220223330211013-3003030230201000-0012201203012111-0030300101330003-3130003301133023-2030030012213331-3201231311132002"></a>

## Direct properties — enforcement_block / 223330133113 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0010122132133313-0233223020323010-0220202001220002-0012302200002133-2103120012300330-0302011232023210-1133022030030021-1322233032020221"></a>

## Next pages — enforcement_block / 223330133113 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2131232230131012-3331321221102330-3103301030130233-1113321213010313-2101313211331133-0301220131231232-2131313023300033-0131310312011121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3001233121123101-2311031212202233-0132220003011003-1330202300121303-2223202223113023-2202233113113120-0131110223211313-1123213001323130"></a>

## api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report — enforcement_report / 203033102211 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-006.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report

<a id="canonical-1121213023001001-1201102211221221-0021301221321320-1223333331302031-0013113130121230-0213120130001032-3210123132301013-0133302131321321"></a>

Type: `["object", {}]`. Optional.

Report-only validation: record OpenAPI violations while allowing the request or response to
continue.

Receipt-pinned upstream constraints:

```json
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
enforcement_report = {}
```

<a id="canonical-1022223232332023-3200211320123002-1212132123033220-0201213023110330-3210232302110212-1020301111201130-0133322131332012-1333332101230330"></a>

## Direct properties — enforcement_report / 203033102211 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1323101120232310-2022011310012000-3111123301000312-3300330202131002-0111032013020021-2132011120123120-1111213322223330-1312322013102023"></a>

## Next pages — enforcement_report / 203033102211 / 4

- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-006.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3310210123303103-2101231133333132-1121121211223333-2132222311000100-1220130221211121-1313011113102233-1121221133010312-0223310132022200"></a>

## api_specification.validation_custom_list.settings — settings / 210220233021 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- api_specification.validation_custom_list.settings

<a id="canonical-0323313333323012-0032031231333130-2023003121233003-2232110230112010-2113223321211211-0322120022122331-0330220001320010-1230022122001330"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Upstream description:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("oversized_body_fail_validation",
    "oversized_body_skip_validation"),
  validators.ConflictingObjectAttributes("property_validation_settings_custom",
    "property_validation_settings_default")}
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
  "x-ves-oneof-field-fail_configuration": "[]",
  "x-ves-oneof-field-oversized_body_choice": "[\"oversized_body_fail_validation\",\"oversized_body_skip_validation\"]",
  "x-ves-oneof-field-property_validation_settings_choice": "[\"property_validation_settings_custom\",\"property_validation_settings_default\"]"
}
```

Terraform syntax:

```terraform
settings {
  # Configure direct properties listed below.
}
```

<a id="canonical-1111111323002001-2320323221112223-1032212331233101-2101111131023120-1123323201223320-3230212012133231-1221032001110001-3212033333230013"></a>

## Direct properties — settings / 210220233021 / 3

- [oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320110021320333-2333110032221120-1031221221223012-0110310021303120-2200203222110032-3203111130311330-2312021011132120-1231303320203131): complete subsection reference.

- [oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-0320322023000222-1021103031102021-1333223013312132-3313310323032010-3302220210323201-0013232312131320-3310232033112132-1210312302302132): complete subsection reference.

- [property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322): complete subsection reference.

- [property_validation_settings_default](resources--cdn_loadbalancer--reference--group-007.md#canonical-0223110103210310-0130320302013122-2020111111213303-0013023302202000-2031301003230301-2200000202201320-0301110302102210-2210030303123032): complete subsection reference.

<a id="canonical-0121333002301233-0233311021111302-2130313113330130-3233200010321113-1100131302010323-2201112113201203-3201122331322301-2011303330302311"></a>

## Next pages — settings / 210220233021 / 4

- [api_specification.validation_custom_list.settings.oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320110021320333-2333110032221120-1031221221223012-0110310021303120-2200203222110032-3203111130311330-2312021011132120-1231303320203131)
- [api_specification.validation_custom_list.settings.oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-006.md#canonical-0320322023000222-1021103031102021-1333223013312132-3313310323032010-3302220210323201-0013232312131320-3310232033112132-1210312302302132)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322)
- [api_specification.validation_custom_list.settings.property_validation_settings_default](resources--cdn_loadbalancer--reference--group-007.md#canonical-0223110103210310-0130320302013122-2020111111213303-0013023302202000-2031301003230301-2200000202201320-0301110302102210-2210030303123032)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2320110021320333-2333110032221120-1031221221223012-0110310021303120-2200203222110032-3203111130311330-2312021011132120-1231303320203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213103221310221-1101201230101121-3121222033120233-0131310101202201-0210103100300202-0230122000023000-2113021000323001-0233331132110231"></a>

## api_specification.validation_custom_list.settings.oversized_body_fail_validation — oversized_body_fail_validation / 210312322123 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-1002300013312123-1133123020333013-0123331023001021-1111130301012023-0012121211203223-2321123200123303-1321323232312220-0002023013302300"></a>

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
oversized_body_fail_validation = {}
```

<a id="canonical-1133313201023100-2233332032211322-2311211223003301-3300221323210210-1230003231210023-1203032103223021-1133021321300110-3323032333332310"></a>

## Direct properties — oversized_body_fail_validation / 210312322123 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0303031001110221-3212322313020300-2210100220112001-0311033302331121-0003021113032211-2003113113211202-1002002012103103-2200221123211101"></a>

## Next pages — oversized_body_fail_validation / 210312322123 / 4

- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-0320322023000222-1021103031102021-1333223013312132-3313310323032010-3302220210323201-0013232312131320-3310232033112132-1210312302302132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313010133010212-2312213203210202-1023301222003131-1210321301101123-0120113321132212-0200003203123312-0211121230202022-1301031232231021"></a>

## api_specification.validation_custom_list.settings.oversized_body_skip_validation — oversized_body_skip_validation / 230112210330 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-0100112203001313-1013110110223121-2013330100111320-0323310113130321-1223232130130012-2011311013322023-0003012003101110-1221120302121313"></a>

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
oversized_body_skip_validation = {}
```

<a id="canonical-3130333300211121-0020202221032332-1121202013302113-2320201130301022-2031100110332310-2132022232021130-2020112210320031-0102211320232003"></a>

## Direct properties — oversized_body_skip_validation / 230112210330 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2312030232110231-2013211321100330-1002331003111033-3231331202111331-2112311011301022-0021032120210213-3033220213301210-0012132113133221"></a>

## Next pages — oversized_body_skip_validation / 230112210330 / 4

- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3302213213201213-0333110203123023-0222202311202132-1122222333031023-1112210130233212-0033220221030112-1213032011122213-1100223303221300"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom — property_validation_settings_custom / 312210113202 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-0332203131331201-1310012120132032-1003311022033020-3030122310313130-2313110221210002-3123313322310202-0121103310321000-3232030112123302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Upstream description:

Custom property validation settings.

Receipt-pinned upstream constraints:

```json
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
property_validation_settings_custom {
  # Configure direct properties listed below.
}
```

<a id="canonical-2000330032223311-0303020113022223-3033002022102022-0120302001131010-2201123330122320-3313031013202123-2323121220320023-0222100332323221"></a>

## Direct properties — property_validation_settings_custom / 312210113202 / 3

- [query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110): complete subsection reference.

<a id="canonical-1320210031322103-3101133331030010-2322123011102100-3020203113330031-0110201301003332-1321120310020322-1330110012121120-3202032022021101"></a>

## Next pages — property_validation_settings_custom / 312210113202 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3120133201220300-0101020322200200-1001221021001131-0130312022102020-1010322132031230-3333132321033110-0301013203133302-2233033201023101"></a>

## api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters — query_parameters / 302210031231 / 2

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-005.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-006.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-006.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-1032121321131310-0321311331023331-0031103010122121-1310012110113231-0303231002312012-0010220333132032-3132110230131312-0320230322310320"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("allow_additional_parameters",
    "disallow_additional_parameters")}
```

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2200320300000201-3031223322312100-2321331323122001-3030101100301122-2301203100233213-0123221010322132-2333331310113301-0012301032200033"></a>

## Direct properties — query_parameters / 302210031231 / 3

- [allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-2312013122222333-1220113330102333-3230220223330113-3220132113230323-3130322310000031-0013033221030110-2323100213021133-2221211103323213): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-0112231002011022-0230023311030233-3211330222331103-0323211212212012-2102000313122130-3032002320331300-3222020100121322-3132020120110133): complete subsection reference.

<a id="canonical-3023333330033110-0112231101210210-1100133311031310-3022301101210231-3101120010021131-3010101332213223-2202133133031320-0030130023203001"></a>

## Next pages — query_parameters / 302210031231 / 4

- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters](resources--cdn_loadbalancer--reference--group-006.md#canonical-2312013122222333-1220113330102333-3230220223330113-3220132113230323-3130322310000031-0013033221030110-2323100213021133-2221211103323213)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-0112231002011022-0230023311030233-3211330222331103-0323211212212012-2102000313122130-3032002320331300-3222020100121322-3132020120110133)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)

<a id="canonical-2312013122222333-1220113330102333-3230220223330113-3220132113230323-3130322310000031-0013033221030110-2323100213021133-2221211103323213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->
