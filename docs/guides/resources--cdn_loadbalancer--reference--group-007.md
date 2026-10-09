---
page_title: "xcsh_cdn_loadbalancer reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_cdn_loadbalancer reference."
---

# xcsh_cdn_loadbalancer reference

<a id="canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-0110010023210300-0103132130022200-1210332122023122-3231223000021230-1021130322203310-3131021021011003-2202332301223323-0301113330121123"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0212020030131210-0120332123112303-2021023012301033-3032033200130211-2310202233103133-0103031213312103-1112113021223000-0201332021230131"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022021220103010-3102230030122201-0202213022330321-0022123120302221-0332000121111211-0100310301330132-2323131231202100-2023300332012223): complete subsection reference.

- [action_report](resources--cdn_loadbalancer--reference--group-007.md#canonical-2132011323120203-2013132122023310-0133011203202102-1011022211221121-3121031311220000-0310203001323302-1130201120032311-2200103301011121): complete subsection reference.

- [action_skip](resources--cdn_loadbalancer--reference--group-007.md#canonical-0133032033320222-3111131132223303-0321131322312202-0013130131231010-2013103300303010-2321321013031322-1203321123222202-2202121220201232): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-007.md#canonical-1103111311121333-0113333123303023-1322223010120310-3112011320132332-2222310112112311-0233131102323212-0321130302221232-2011300231311323): complete subsection reference.

<a id="canonical-2311222322200332-0303101101122110-1203120303101303-0010202031031021-0213001020213310-0033003301212130-2132101012302032-3130211202100110"></a>

<a id="canonical-0101130230022032-3110100322021102-1003213213132030-1113012320121031-3003232333231031-3200002201001200-0200110102223103-1323221313100202"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0021010110031132-2301313010310103-2223223320210030-2211113213120102-2021231102032030-0322321310212130-0332202231021223-1102212330102021"></a>

<a id="canonical-3233010012121132-1011113333222031-3022320023330333-1112011213302213-0000103300011202-1322312230003303-0012320120031030-0332300130103303"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-3122022232220301-3022311110311231-2132233022101212-1033021212201102-2302313211003333-2321332322111311-2313032223301201-3112030031020321): complete subsection reference.

<a id="canonical-3022021220103010-3102230030122201-0202213022330321-0022123120302221-0332000121111211-0100310301330132-2323131231202100-2023300332012223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-1020030013131122-3011313121220303-0023112320321223-1000302330011011-2031022010133111-1002323231113231-0120131313032203-2122201020130031"></a>

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
action_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2132011323120203-2013132122023310-0133011203202102-1011022211221121-3121031311220000-0310203001323302-1130201120032311-2200103301011121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-3021302223123330-3323011220130231-1211223330310120-0330110232001030-3112210032321003-2030211132120232-2032221132002100-3132010211311111"></a>

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
action_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0133032033320222-3111131132223303-0321131322312202-0013130131231010-2013103300303010-2321321013031322-1203321123222202-2202121220201232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-0013222000310230-3002102300221210-0002121321021101-0030311301303132-0133010001231201-2101233303022002-2123122220232331-0333032103021123"></a>

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
action_skip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1103111311121333-0113333123303023-1322223010120310-3112011320132332-2222310112112311-0233131102323212-0321130302221232-2011300231311323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-1031210111331133-0200303110000110-2131002230010332-2201331032221231-0311233303112220-0123031220131300-0311132013302100-1203332112130000"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-0012322333200122-3323200201110013-2212013303330113-1231312311111120-2300123232323301-2202132023320310-3111213132312212-0233110210321312"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-1321020230013021-3223131020322233-2333112102131220-3301113331103021-3023303211210223-1130323121310233-0131031030103011-1113303100223011"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1210322212323302-3320331310200313-1112002131233332-1030130203021021-1033220302121230-3323220132123303-1010000322200101-0222011333302012"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3122022232220301-3022311110311231-2132233022101212-1033021212201102-2302313211003333-2321332322111311-2313032223301201-3112030031020321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.fall_through_mode](resources--cdn_loadbalancer--reference--group-006.md#canonical-1132202011301023-1311003221212003-0223220113211003-0022210103300313-0200312302211323-3013210300201133-2131013023312002-3210220003000332)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-006.md#canonical-1313112202023113-1100031331202211-3202200203100020-1313201021000212-3033331010021322-2232130310013013-0313201103013202-3001021303232013)
- [api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0123300033001212-0103231020012121-0122330200112313-1113202010311313-0313101101311223-0121320230132013-3110202300123132-3323213313332100)
- api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-3022100003011012-1323233332033103-2302222333303300-0333211311333202-0012113310200230-1121102012223320-0002131200203300-3033213032203100"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1221211101322310-3313203320000013-3032033312233003-2123102100013202-2201111232223301-3220210123313023-0320222213200210-0213012002203121"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-0132131323203210-2103101101230321-0230331200113120-3003303123122320-0121221031211113-0030323303030000-3300202113003321-0033300301102111"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-1021030120332302-1322031212121033-3111203123212331-3123323003303102-1231013323131100-0303012222312323-0233023203200110-3203233022023023"></a>

<a id="canonical-3220023110120333-1133022321332232-2110013213002231-2212132010323103-2321120123333133-2000011032230231-2000121312001010-2313310332211333"></a>

#### `api_specification.validation_all_spec_endpoints.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- api_specification.validation_all_spec_endpoints.settings

<a id="canonical-1320313203323313-1103210203113131-3100310201103302-0321312031333031-0302311212113213-0000202102021201-2333131330112010-1223100033232311"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

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

<a id="canonical-1010200121332300-3032033021323220-1232100331223301-2113002213122320-0101100233130001-1002001013313230-2022323102001101-3323212120220030"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings`

- [oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-3111022133001320-3223203100220113-2132120223213223-0121022300023213-3323310102031120-1012231201032201-2000013030213320-1033210130113012): complete subsection reference.

- [oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-2301103213103110-2203331121123311-2101333100120313-2001022332232222-0233022103033123-2203312212330222-3111001020323133-2020230101022331): complete subsection reference.

- [property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123): complete subsection reference.

- [property_validation_settings_default](resources--cdn_loadbalancer--reference--group-007.md#canonical-3132303220021203-3111022333231003-1021031000232103-2010022011210210-0330232001321100-0310002113322031-2013010013000102-0313311100120131): complete subsection reference.

<a id="canonical-3111022133001320-3223203100220113-2132120223213223-0121022300023213-3323310102031120-1012231201032201-2000013030213320-1033210130113012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_fail_validation

<a id="canonical-2000331231122003-1302312233102022-0203133320330100-3033313110210312-2020003111100131-3202320003032231-3310313113030130-3221121010121313"></a>

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
oversized_body_fail_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2301103213103110-2203331121123311-2101333100120313-2001022332232222-0233022103033123-2203312212330222-3111001020323133-2020230101022331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.oversized_body_skip_validation

<a id="canonical-2133103022203003-2030130112201100-2222000332023203-0022021333002332-1133103001222223-3223222011212303-1321120330023122-1310331323131300"></a>

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
oversized_body_skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom

<a id="canonical-2113102202220010-3223020112323301-1132132321201102-2220220032303222-1012300012113112-3210120220033133-1222012312302101-0113113113000133"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Additional upstream details:

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

<a id="canonical-2213001201201303-3323022133012111-2032233200332103-3221332222202230-2003013103332120-0210101210320131-3011003120322032-3213000331130123"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom`

- [query_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123): complete subsection reference.

<a id="canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters

<a id="canonical-2020211132311321-0200030300033210-0001311323312300-2003123201032033-1132230032300002-2211330123232321-0231110312220302-1132030323231202"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-2102120202003210-0130220211000302-2031010113201000-0223203201103131-2213322303231330-2223332332123211-0203223333132030-0003121310023032"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1013022311302032-2111213201323130-3022332330331330-0121201102302332-2133331121100022-3030002310123122-0103211303132133-0033133312011322): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-2330332323102122-0120133033102210-2213223331202203-3132333032033302-3232123202001321-3112313200320011-3033211322013330-0223133101203013): complete subsection reference.

<a id="canonical-1013022311302032-2111213201323130-3022332330331330-0121201102302332-2133331121100022-3030002310123122-0103211303132133-0033133312011322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-3002201113213032-2103023032232321-0322133000001211-2032222320131000-0323321111303112-1311113033020012-2100202321133332-2103301331103010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2330332323102122-0120133033102210-2213223331202203-3132333032033302-3232123202001321-3112313200320011-3033211322013330-0223133101203013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3313121022323211-0210110222302303-0130331011000322-2232031122333323-3302130331222103-3123013111013301-0320210221311201-2213222120321123)
- [api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1323030223233110-3102033230222013-2110223213031022-1302223222010000-2001011013322130-1323113010231312-3232030313312213-2221212111230123)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-2111113301120013-3210223002030322-1220220121102020-3230202233230221-1011011321132110-0312222333212323-0131313213313111-0202201030223210"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132303220021203-3111022333231003-1021031000232103-2010022011210210-0330232001321100-0310002113322031-2013010013000102-0313311100120131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212122210213200-2011131031130121-1310310331303200-0003121202031101-2212233312312021-1333022320112232-0110003003230311-3212032312203310)
- api_specification.validation_all_spec_endpoints.settings.property_validation_settings_default

<a id="canonical-0121001130112121-1123310201321002-3010313233011000-2122223310301223-0232210002130131-1220220311033322-2011202133021101-3233311200210321"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for property validation settings default.

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
property_validation_settings_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- api_specification.validation_all_spec_endpoints.validation_mode

<a id="canonical-0212123003212312-3111322333000111-0100031323100101-0333202101202133-0022030312202011-3302013232312220-2322322130220221-0232220330303011"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

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

<a id="canonical-0102130231110201-0213113322003210-1213102012202001-2002012113120232-0100030001320333-2032020122222213-0010002101103120-3022102220231321"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode`

- [response_validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-1302321003220022-2132011332213003-1322331023203221-0200233011032321-3231101312030110-2301032211311103-3333123113230022-1231203132302101): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-0112002210031102-2022111230011230-3120110121131000-2103131301230230-1222213332311232-1212013223033133-2120312023301111-0022301120310203): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212): complete subsection reference.

<a id="canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active

<a id="canonical-0313110013102010-1333023120312020-3103330001021021-0123310030120020-3113223221031302-3220232102133210-3202220120021203-2103313303102033"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

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

<a id="canonical-2311312110011032-0010010103333223-2132111333302112-3323331121012103-0103000212313131-2201121302230121-3111322010130313-0101221012310202"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active`

- [enforcement_block](resources--cdn_loadbalancer--reference--group-007.md#canonical-0210302333310310-1231311212222012-3133320013101221-1121211230120001-1130311210323112-0033122313133313-2323310223321013-3033132303130320): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-007.md#canonical-3203121032033322-0331210220132112-1331101131310332-1013110330220131-0231221000023222-3310010330112231-3223210331020010-0221120001021321): complete subsection reference.

<a id="canonical-1223112111200313-2202020311313002-0303212121112201-0122122203030200-2120131201330023-3213323012112013-0120003301010210-3000203331101011"></a>

<a id="canonical-2311021210212213-1223333122212023-0132232320020313-2020333121302313-3300023323303313-1311121101132013-2210330201123323-1023210100030120"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0210302333310310-1231311212222012-3133320013101221-1121211230120001-1130311210323112-0033122313133313-2323310223321013-3033132303130320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3203121032033322-0331210220132112-1331101131310332-1013110330220131-0231221000023222-3310010330112231-3223210331020010-0221120001021321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022122310113222-0203022111102221-3202100012001231-3133000133010221-0022023121002013-2003001222312211-0330321320012302-1310223031010233)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302321003220022-2132011332213003-1322331023203221-0200233011032321-3231101312030110-2301032211311103-3333123113230022-1231203132302101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_response_validation

<a id="canonical-2020330211323302-1212220132302321-1000202232332102-1032002133210301-3132202111222001-3011121002233312-1010302233213232-2332022102012213"></a>

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
skip_response_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112002210031102-2022111230011230-3120110121131000-2103131301230230-1222213332311232-1212013223033133-2120312023301111-0022301120310203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.skip_validation

<a id="canonical-3010221033031110-2201311110233213-0331303222132220-1110132223230031-3020133222300203-1032203032310203-1333010221031121-2312030300320313"></a>

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
skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active

<a id="canonical-0102002131231210-2201312221102313-0031300300001230-3310233320331011-0323112032311103-3211111001030300-1123011232302300-3212303203210022"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

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

<a id="canonical-2111003003330133-1203001201130323-0112131331203003-0321003331011210-2120332102220302-2012011020230310-1120200023032321-1202022003023230"></a>

### Direct properties for `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active`

- [enforcement_block](resources--cdn_loadbalancer--reference--group-007.md#canonical-2002101302303310-2002022212330203-0200313120330031-1003021002320023-3320122113202302-3330200310101101-3202032302201110-3101301002221000): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-007.md#canonical-1032133120312121-0112201100312123-3320232011311112-2200313330321023-1112021031102013-2021010202102212-1311223020003020-1122003103112201): complete subsection reference.

<a id="canonical-0032223130303122-0131203333001131-3132320012313320-2300033233233203-3001313130221201-3123132302133120-0320110132013000-2100303222321022"></a>

<a id="canonical-3221213010110123-1332332110210332-0123310313210120-0300100332130220-1211123012021112-0321211131022232-1123013110330021-2102020022003102"></a>

#### `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2002101302303310-2002022212330203-0200313120330031-1003021002320023-3320122113202302-3330200310101101-3202032302201110-3101301002221000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1032133120312121-0112201100312123-3320232011311112-2200313330321023-1112021031102013-2021010202102212-1311223020003020-1122003103112201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_all_spec_endpoints](resources--cdn_loadbalancer--reference--group-006.md#canonical-3323203003320030-1232313132310021-0103112110123230-3120321220223022-1020110020221000-0010330221033132-0130001101310222-0231010233112333)
- [api_specification.validation_all_spec_endpoints.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-0032000232202023-1021303223203232-2311231222130001-0301110330022200-1322133020231031-0122001231011131-3132131303303103-1013212230000131)
- [api_specification.validation_all_spec_endpoints.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-0322312301121222-3230112021320322-2332012201233121-2031231321001200-2030201211123211-3102012323302021-3201310033203130-1303300310322212)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.validation_custom_list

<a id="canonical-3022302121112320-1202230322313102-0000211002111233-3010030103323121-3200132321223200-3333011321230303-3123023212032332-0112203121012131"></a>

Type: `"object"`. single nested block, Optional.

Define API groups, base paths, or API endpoints and their OpenAPI validation modes. Any other
API-endpoint not listed will act according to 'Fall Through Mode'.

Additional upstream details:

Any other API-endpoint not listed will act according to "Fall Through Mode".

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

<a id="canonical-1333331023331133-3201331211121323-2220203221033121-2320322131302321-3303311212313001-0201330203333113-0302020122023332-0211321311132023"></a>

### Direct properties for `api_specification.validation_custom_list`

- [fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230): complete subsection reference.

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031): complete subsection reference.

- [settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233): complete subsection reference.

<a id="canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- api_specification.validation_custom_list.fall_through_mode

<a id="canonical-3220320111130013-1131320303110230-1232233210002021-3331031020101301-0111333122230023-3012212031133303-0210031101011131-1102132011231312"></a>

Type: `"object"`. single nested block, Optional.

Determine what to do with unprotected endpoints (not in the OpenAPI specification file (a.k.a.
Swagger) or doesn't have a specific rule in custom rules).

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

<a id="canonical-3131021200133122-2200101300013003-1323001330302011-3201102012301033-3113321102113031-3132303032112311-1111020300132030-3213233231023121"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode`

- [fall_through_mode_allow](resources--cdn_loadbalancer--reference--group-007.md#canonical-0111330000221003-2203010113323011-2331020202021302-1211021301330100-1102213021323001-0012203131213231-2322320100332223-3003130310321233): complete subsection reference.

- [fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211): complete subsection reference.

<a id="canonical-0111330000221003-2203010113323011-2331020202021302-1211021301330100-1102213021323001-0012203131213231-2322320100332223-3003130310321233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_allow

<a id="canonical-3201203110201322-2202233220223123-2133020111010311-2103202003323233-2202302003020200-1110212001122213-0221100332133212-2232202110100223"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for fall through mode allow.

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
fall_through_mode_allow = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom

<a id="canonical-0102303111101020-2021131333202322-0123200220001111-3210112213322232-1200201333022202-2120322311333000-0031311101330022-1321210203303100"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for fall through mode custom.

Additional upstream details:

Define the fall through settings.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1320100310230212-0222131021103302-0032230223303203-0231201221323320-2302210313221302-1331131032133331-2330212221122320-1200123302223331"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom`

- [open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033): complete subsection reference.

<a id="canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules

<a id="canonical-2212211330233221-3332013323112111-1311123231301120-2211032130333322-0313303132303220-2322123303000212-0332201330220123-0312123102123100"></a>

Type: `"object"`. list nested block, Optional.

Custom Fall Through Rule List. Rule or policy definition

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0002233331113020-0200103221221031-0111103130220033-3013221212202100-2023211200112112-0210012221103200-3112121211103233-1112232021322012"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules`

- [action_block](resources--cdn_loadbalancer--reference--group-007.md#canonical-0322013211132302-1133110233011131-2122302333323121-1300221001133133-1103232111213310-1302222031022131-2012103032331000-3130001320022020): complete subsection reference.

- [action_report](resources--cdn_loadbalancer--reference--group-007.md#canonical-3002032202300322-1221033310333121-3000132021223031-3330231031223320-2202303121120031-3111300223101301-2122030310202300-1130001210313121): complete subsection reference.

- [action_skip](resources--cdn_loadbalancer--reference--group-007.md#canonical-1203322032320111-3011233302330222-1131332022012100-1012100013203333-2313322113323132-0020121130230113-1000111322032300-1033002032313110): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-007.md#canonical-1233011312013203-2333000201303113-1002101013330021-2032021012133213-0230121310211332-0110230321213010-2332131330121310-2031132100120311): complete subsection reference.

<a id="canonical-2011032032310232-1120121020333233-0311303212322133-1001310223222330-0021002310031322-3303021131212300-1332200032023213-2011012202101222"></a>

<a id="canonical-1311233123211123-0110130332232130-1332131003130332-0201103330232323-0121301031031120-1101120123303103-0133230201210122-0213201122113230"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_group` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1021101123231321-2210101022203131-1330233121020131-2003320201203330-3210010220210113-2003221332333233-2100023331123101-2310002131310032"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-2213100130022110-2203122121211011-0032211321331011-2102313123013321-2322012022210132-1112211000320131-1223302321120220-3011012120122131): complete subsection reference.

<a id="canonical-0322013211132302-1133110233011131-2122302333323121-1300221001133133-1103232111213310-1302222031022131-2012103032331000-3130001320022020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_block

<a id="canonical-3002132212102200-1011131031103021-2210222203123232-3231333120230101-2320210200113311-2330033232100231-3213033122311221-0302300021312003"></a>

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
action_block = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002032202300322-1221033310333121-3000132021223031-3330231031223320-2202303121120031-3111300223101301-2122030310202300-1130001210313121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_report

<a id="canonical-0101031232130033-3233321220312300-0102200201111233-1012033110113110-1023131032130023-3201013203013021-2310203103132111-2321233212101000"></a>

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
action_report = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1203322032320111-3011233302330222-1131332022012100-1012100013203333-2313322113323132-0020121130230113-1000111322032300-1033002032313110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.action_skip

<a id="canonical-3010102213113332-2322312230102221-3013032032132230-3231030002231021-1210321332001003-0202221301330311-0210122110230321-2213012031221223"></a>

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
action_skip = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233011312013203-2333000201303113-1002101013330021-2032021012133213-0230121310211332-0110230321213010-2332131330121310-2031132100120311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint

<a id="canonical-1301220321310310-0332231021000303-1030331203302331-1232330331333330-1211332332200103-3110302230011301-0312110202123130-3112113203233301"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-1200010333131030-1230200201200130-3200311132113120-2211202233223023-2233133223301321-3013221320213201-1222333130110003-1132122011021311"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint`

<a id="canonical-1321232122232112-0233202323133022-1320233133211303-0101121321213213-0030013312023230-3021001201221012-3112133210132031-1320302302312200"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2021110301212301-2310333132100022-3130223101310011-3113330201122311-0311030130210222-3321200030110313-2133113013032012-1132130111321321"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2213100130022110-2203122121211011-0032211321331011-2102313123013321-2322012022210132-1112211000320131-1223302321120220-3011012120122131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.fall_through_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-1333000011110022-0303220331331101-1231030023003033-3310020201100322-1132232122120331-2311222200231322-1213033220210020-0021133031302230)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-2102230001220021-0303023030113230-2023122320323220-2332023320010002-1302200213211231-2032333002222300-3313032112013010-1101130003020211)
- [api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-2212313230111030-0133320322323320-3033221202001102-2020312310223131-0031311320021203-3332303121300303-1323122312203232-3212131101030033)
- api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata

<a id="canonical-2321231000031310-3311113013222023-3101321012331311-3122223323222320-3212321203102210-2202333332213131-2131111303123111-0312332310301110"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2020023210200322-3111333133303331-2110002020020100-0320323012221031-3123031211033231-3210301202123220-0133002201110203-2211111133323122"></a>

### Direct properties for `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata`

<a id="canonical-0011332311122120-2032120303330302-3202122102200133-3203310330330033-1323233020032202-0120111102233101-3303112221232002-2021020300232211"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-3330300010310133-3130113333213331-2211231110303101-1223320200021320-3321113312100300-1110200301101123-1120010321200000-1103322231202002"></a>

<a id="canonical-1012132221323121-2213113003103223-2032212233131213-2311210233312020-2200031222023301-0312112001010323-2311011112220023-1303323211211012"></a>

#### `api_specification.validation_custom_list.fall_through_mode.fall_through_mode_custom.open_api_validation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- api_specification.validation_custom_list.open_api_validation_rules

<a id="canonical-1000333221022003-3203031023202323-0213011233002102-3133310303321200-1033313231013213-0200013300133001-1020223221213011-0331220201311022"></a>

Type: `"object"`. list nested block, Optional.

Validation List. Rule or policy definition

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1213003232321320-3023013000213032-1033033200322000-1203322010013313-0310332002220121-2130322000202221-1101310203002131-0120002121022100"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules`

- [any_domain](resources--cdn_loadbalancer--reference--group-007.md#canonical-0001121131003212-2001030100132010-2120213033231312-2231023013230233-2021220102103203-0311221322133130-1111200202331103-2022301000301102): complete subsection reference.

- [api_endpoint](resources--cdn_loadbalancer--reference--group-007.md#canonical-0321322200122101-3122230112323122-2322232020122330-3021303320220130-0311321223333111-3211220010202303-3322110010122300-2023331302112003): complete subsection reference.

<a id="canonical-1220330121202312-1201222230030123-0221112220102011-0013222300201203-0213332111000210-2213102130320120-3223010003223120-2221332301333003"></a>

<a id="canonical-2111012123312201-1013013300301211-3233222020122302-3020201201033033-1032331001312231-2202222312322120-1331022000101321-0200312103323113"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_group` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint base\_path\] The API group which this validation applies to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-1112033121200113-2313203313023000-1321131011202001-0210033231332211-0030222230022122-2313013013331211-3300002023120202-3303213230322232"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.base_path` property

Type: `"string"`. Optional.

Exclusive with \[api\_endpoint api\_group\] The base path which this validation applies to.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-3031030312110130-1313002120122221-2123030002001220-3112030313120103-0200330232322123-2312322201013110-0200012303132033-1000212020012300): complete subsection reference.

<a id="canonical-3311103320132221-0231202233323320-0121123110321110-1331031311310102-0311213300233213-0223000201112030-0030313320110122-3302132102123320"></a>

<a id="canonical-1311010313231112-1012200310122013-1211001132213103-3203232332132120-1332010003311013-1111310032030333-2220013321313203-0321031031022023"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.specific_domain` property

Type: `"string"`. Optional.

Exclusive with \[any\_domain\] The rule will apply for a specific domain.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032): complete subsection reference.

<a id="canonical-0001121131003212-2001030100132010-2120213033231312-2231023013230233-2021220102103203-0311221322133130-1111200202331103-2022301000301102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.any_domain` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.any_domain

<a id="canonical-2133112203211020-2333112210031123-0201301113011110-1322111310211033-3003023202132003-2001112311303201-0331100311012301-2332300123133032"></a>

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
any_domain = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0321322200122101-3122230112323122-2322232020122330-3021303320220130-0311321223333111-3211220010202303-3322110010122300-2023331302112003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.api_endpoint

<a id="canonical-3311331013223222-0213221323310022-0332100000213113-1330220121312223-0123132022032012-2121231121103311-3100130220100120-0200310323000132"></a>

Type: `"object"`. single nested block, Optional.

API Endpoint. This defines API endpoint.

Receipt-pinned upstream constraints:

```json
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

<a id="canonical-3303120122233032-0010331223021102-3110333303000233-0333311003130233-2313132211121222-3213123103011231-2210300011022320-2003022112012131"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint`

<a id="canonical-3020121201230312-0030132000102221-2033321013101113-1011003310202222-0123331103302101-2303222303111332-3021103302012002-2313013120011332"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.methods` property

Type: `["list", "string"]`. Optional.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Methods. Methods to be
matched. Possible values are \`ANY\`, \`GET\`, \`HEAD\`, \`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`,
\`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to \`ANY\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0022130132133101-1223201332102121-2122300202123233-0102301021020111-2201323301201302-2120221220120132-0030310101222103-1331131203033132"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.api_endpoint.path` property

Type: `"string"`. Optional.

Path. Path to be matched.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3031030312110130-1313002120122221-2123030002001220-3112030313120103-0200330232322123-2312322201013110-0200012303132033-1000212020012300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.metadata

<a id="canonical-1133321121303113-2203110032011330-3123332233220000-0201130210102011-2112320030103313-1111123132213130-3222000201101323-0123310303132000"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-2212023210122220-1010000231003012-2210103133223332-2210312133012302-2201333221023121-0021330133123123-1121123132022333-2130120312202300"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.metadata`

<a id="canonical-3030131301120312-2203300203303001-1203213013313120-0100222323223212-1020211333233000-0202232002111000-1213103133221010-0210200102102210"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2330322212232122-2013030133323202-0113321230021121-2220121232003022-1223233232221033-3021021020120212-1313203233013033-0302032110333011"></a>

<a id="canonical-1332120033231022-2321033023030323-2332330233002030-0221103013031112-0211110223321212-0230300003122011-2312201301210302-2030100131301322"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode

<a id="canonical-1203002200333322-3300331231320023-1222231131102113-0320011111131320-1300200001122133-1110110203133200-2102220222011021-1223113212300123"></a>

Type: `"object"`. single nested block, Optional.

Validation mode of OpenAPI specification. When a validation mismatch occurs on a request to one of
the endpoints listed on the OpenAPI specification file (a.k.a. Swagger).

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

<a id="canonical-1301303232030231-3332123003111131-2002030131310032-1222201112022032-3200111001322020-1011021313301121-0030210010131132-1020020010313023"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode`

- [response_validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202): complete subsection reference.

- [skip_response_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-2111302223301013-1313333332013013-2022001332312010-0130033332132132-0000312223312310-0101212030320322-3223001031011123-1033232001320200): complete subsection reference.

- [skip_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-0131000113200332-1122033201111000-0130010212313212-2013000222113202-2131202033111200-3320112131103022-0132003131013021-2021333231110030): complete subsection reference.

- [validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033): complete subsection reference.

<a id="canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active

<a id="canonical-1121321330230010-0121220123311032-3220121323221221-1103211321130323-2033230100312111-1112223200322230-1122032120231302-0313033030322011"></a>

Type: `"object"`. single nested block, Optional.

Open API Validation Mode Active. Validation mode properties of response.

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

<a id="canonical-2202101012213110-3122210100011002-3330020213123111-2323130000203312-1100312321012021-3021102110031321-2001032302212110-1012033012023011"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active`

- [enforcement_block](resources--cdn_loadbalancer--reference--group-007.md#canonical-0202300221020321-0102210232001200-2113123030211023-2203320100012211-2220210022000332-1300330203032303-2033220001031102-2223131023303213): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-007.md#canonical-3102120002033333-1302211001102110-1003100030330131-1021020330330000-2233013002231231-3013020212130012-3322201110230033-2030002132322002): complete subsection reference.

<a id="canonical-2200120102031100-2030122032021011-1113311320023222-3033222110212003-0030322201003222-3130111220331201-0000333323211210-0211212100132023"></a>

<a id="canonical-0031330332030321-0202110032131301-0203113213121301-3320012322133223-3122300120211023-0031032213332120-2020003021230303-2112221023113013"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.response_validation_properties` property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the response to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0202300221020321-0102210232001200-2113123030211023-2203320100012211-2220210022000332-1300330203032303-2033220001031102-2223131023303213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3102120002033333-1302211001102110-1003100030330131-1021020330330000-2233013002231231-3013020212130012-3322201110230033-2030002132322002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.response_validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-0121022301320020-3331103131021030-0320132121230232-0010211112301101-2101010302111110-2022320002302213-3220023302111002-1301320331220202)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2111302223301013-1313333332013013-2022001332312010-0130033332132132-0000312223312310-0101212030320322-3223001031011123-1033232001320200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_response_validation

<a id="canonical-0232023303213233-3302203033103122-0002000303033323-0110323310202000-3332000000231020-0202101010112211-3001331331231102-0201103122031310"></a>

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
skip_response_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0131000113200332-1122033201111000-0130010212313212-2013000222113202-2131202033111200-3320112131103022-0132003131013021-2021333231110030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.skip_validation

<a id="canonical-2310223133223310-3122123020321111-2113021203320322-0200231112212320-2311130213123303-3130000211123331-1010223101211233-2300222211102321"></a>

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
skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active

<a id="canonical-2112301023102212-2322120303123032-1203303022311212-2233130310300123-1211231223121211-0020121221032302-3323333123230010-0022001022000221"></a>

Type: `"object"`. single nested block, Optional.

Enable OpenAPI validation and explicitly select enforcement\_report to allow and log invalid
traffic, or enforcement\_block to reject invalid requests with HTTP 403.

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

<a id="canonical-2131332031200300-0110200021113002-2313333330133301-3112233332332030-0112123320103010-0231001322122201-0233220101210221-2321302232311110"></a>

### Direct properties for `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active`

- [enforcement_block](resources--cdn_loadbalancer--reference--group-007.md#canonical-2331311033131132-1030322002333001-3031323011200330-0312130102231021-0111000000213010-0233000201100002-0203212311120031-3231133133130213): complete subsection reference.

- [enforcement_report](resources--cdn_loadbalancer--reference--group-007.md#canonical-2131232230131012-3331321221102330-3103301030130233-1113321213010313-2101313211331133-0301220131231232-2131313023300033-0131310312011121): complete subsection reference.

<a id="canonical-2020200300220130-3032113222110002-3112022011100100-1321023023011120-1330213010222010-3233033111100302-2220210231121233-2321231203311010"></a>

<a id="canonical-3103110030311331-0123102012113221-0332102212231300-2122303110221120-3130332201313110-1002011223223333-2121033322330212-3213301203012110"></a>

#### `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.request_validation_properties` property

Type: `["list", "string"]`. Optional.

\[Enum:
PROPERTY\_QUERY\_PARAMETERS|PROPERTY\_PATH\_PARAMETERS|PROPERTY\_CONTENT\_TYPE|PROPERTY\_COOKIE\_PARAMETERS|PROPERTY\_HTTP\_HEADERS|PROPERTY\_HTTP\_BODY|PROPERTY\_SECURITY\_SCHEMA|PROPERTY\_RESPONSE\_CODE\]
List of properties of the request to validate according to the OpenAPI specification file (a.k.a.
Swagger). Possible values are \`PROPERTY\_QUERY\_PARAMETERS\`, \`PROPERTY\_PATH\_PARAMETERS\`,
\`PROPERTY\_CONTENT\_TYPE\`, \`PROPERTY\_COOKIE\_PARAMETERS\`, \`PROPERTY\_HTTP\_HEADERS\`,
\`PROPERTY\_HTTP\_BODY\`, \`PROPERTY\_SECURITY\_SCHEMA\`, \`PROPERTY\_RESPONSE\_CODE\`. Defaults to
\`PROPERTY\_QUERY\_PARAMETERS\`.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2331311033131132-1030322002333001-3031323011200330-0312130102231021-0111000000213010-0233000201100002-0203212311120031-3231133133130213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_block` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2131232230131012-3331321221102330-3103301030130233-1113321213010313-2101313211331133-0301220131231232-2131313023300033-0131310312011121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active.enforcement_report` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.open_api_validation_rules](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120331321231033-2212122313301103-0223131221101301-0022122110103302-3021312322201000-3113102133203302-1130010032023130-0220131313203031)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113321233111233-2111202310220232-2220031212320301-0021310011210322-3130233332332122-3331033202021212-0222130202022321-1030002333311032)
- [api_specification.validation_custom_list.open_api_validation_rules.validation_mode.validation_mode_active](resources--cdn_loadbalancer--reference--group-007.md#canonical-1321232211103331-2200212133311102-2321333332322201-1122223310202012-0100222201113230-1020023212301313-2131301201002232-1323100231003033)
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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- api_specification.validation_custom_list.settings

<a id="canonical-0323313333323012-0032031231333130-2023003121233003-2232110230112010-2113223321211211-0322120022122331-0330220001320010-1230022122001330"></a>

Type: `"object"`. single nested block, Optional.

OpenAPI specification validation settings relevant for 'API Inventory' enforcement and for 'Custom
list' enforcement.

Additional upstream details:

OpenAPI specification validation settings relevant for "API Inventory" enforcement and for "Custom
list" enforcement.

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

<a id="canonical-3310210123303103-2101231133333132-1121121211223333-2132222311000100-1220130221211121-1313011113102233-1121221133010312-0223310132022200"></a>

### Direct properties for `api_specification.validation_custom_list.settings`

- [oversized_body_fail_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320110021320333-2333110032221120-1031221221223012-0110310021303120-2200203222110032-3203111130311330-2312021011132120-1231303320203131): complete subsection reference.

- [oversized_body_skip_validation](resources--cdn_loadbalancer--reference--group-007.md#canonical-0320322023000222-1021103031102021-1333223013312132-3313310323032010-3302220210323201-0013232312131320-3310232033112132-1210312302302132): complete subsection reference.

- [property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322): complete subsection reference.

- [property_validation_settings_default](resources--cdn_loadbalancer--reference--group-007.md#canonical-0223110103210310-0130320302013122-2020111111213303-0013023302202000-2031301003230301-2200000202201320-0301110302102210-2210030303123032): complete subsection reference.

<a id="canonical-2320110021320333-2333110032221120-1031221221223012-0110310021303120-2200203222110032-3203111130311330-2312021011132120-1231303320203131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.oversized_body_fail_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.oversized_body_fail_validation

<a id="canonical-1002300013312123-1133123020333013-0123331023001021-1111130301012023-0012121211203223-2321123200123303-1321323232312220-0002023013302300"></a>

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
oversized_body_fail_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320322023000222-1021103031102021-1333223013312132-3313310323032010-3302220210323201-0013232312131320-3310232033112132-1210312302302132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.oversized_body_skip_validation` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.oversized_body_skip_validation

<a id="canonical-0100112203001313-1013110110223121-2013330100111320-0323310113130321-1223232130130012-2011311013322023-0003012003101110-1221120302121313"></a>

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
oversized_body_skip_validation = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.property_validation_settings_custom

<a id="canonical-0332203131331201-1310012120132032-1003311022033020-3030122310313130-2313110221210002-3123313322310202-0121103310321000-3232030112123302"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for property validation settings custom.

Additional upstream details:

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

<a id="canonical-3302213213201213-0333110203123023-0222202311202132-1122222333031023-1112210130233212-0033220221030112-1213032011122213-1100223303221300"></a>

### Direct properties for `api_specification.validation_custom_list.settings.property_validation_settings_custom`

- [query_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110): complete subsection reference.

<a id="canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters

<a id="canonical-1032121321131310-0321311331023331-0031103010122121-1310012110113231-0303231002312012-0010220333132032-3132110230131312-0320230322310320"></a>

Type: `"object"`. single nested block, Optional.

Custom settings for query parameters validation.

Terraform syntax:

```terraform
query_parameters {
  # Configure direct properties listed below.
}
```

<a id="canonical-3120133201220300-0101020322200200-1001221021001131-0130312022102020-1010322132031230-3333132321033110-0301013203133302-2233033201023101"></a>

### Direct properties for `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters`

- [allow_additional_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-2312013122222333-1220113330102333-3230220223330113-3220132113230323-3130322310000031-0013033221030110-2323100213021133-2221211103323213): complete subsection reference.

- [disallow_additional_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-0112231002011022-0230023311030233-3211330222331103-0323211212212012-2102000313122130-3032002320331300-3222020100121322-3132020120110133): complete subsection reference.

<a id="canonical-2312013122222333-1220113330102333-3230220223330113-3220132113230323-3130322310000031-0013033221030110-2323100213021133-2221211103323213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.allow_additional_parameters

<a id="canonical-3013102213321312-2000123002132120-2311223110230121-2303111020232131-1033102312111022-0302202120011132-0301010021310302-2133133322113003"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for allow additional parameters.

Terraform syntax:

```terraform
allow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112231002011022-0230023311030233-3211330222331103-0323211212212012-2102000313122130-3032002320331300-3222020100121322-3132020120110133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom](resources--cdn_loadbalancer--reference--group-007.md#canonical-3022031303303303-3010313201332222-1013011001111002-2111030012011001-0312312321003123-3110200101010203-2000131311323113-3331322012132322)
- [api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters](resources--cdn_loadbalancer--reference--group-007.md#canonical-1000200113012312-1213233230310003-3331010123211031-2312011301332111-3202110021022322-3133012130322101-1133001121322212-3123212230310110)
- api_specification.validation_custom_list.settings.property_validation_settings_custom.query_parameters.disallow_additional_parameters

<a id="canonical-3120103332100113-2232111232122232-3322013010030032-2321233111302333-3031302021201220-0211332202221303-3210121133102022-2030033021302010"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disallow additional parameters.

Terraform syntax:

```terraform
disallow_additional_parameters = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0223110103210310-0130320302013122-2020111111213303-0013023302202000-2031301003230301-2200000202201320-0301110302102210-2210030303123032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_custom_list.settings.property_validation_settings_default` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- [api_specification.validation_custom_list](resources--cdn_loadbalancer--reference--group-007.md#canonical-0232212211000012-3311313113110113-2330211310230303-2211201320310023-2330131112201112-0122323101322112-0122213300313111-3222211223111133)
- [api_specification.validation_custom_list.settings](resources--cdn_loadbalancer--reference--group-007.md#canonical-2320111231222321-2013310323203122-1210010313233001-1312322222323020-2111100203003203-0200223323302132-3320102112223211-0123310010313233)
- api_specification.validation_custom_list.settings.property_validation_settings_default

<a id="canonical-1102222321321323-0330322200202300-3200222012200211-2201230003010213-2113011113201302-0021103223121312-0333130203331312-0202221010002302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for property validation settings default.

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
property_validation_settings_default = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3201130202332300-2123131032222110-0303233101102201-2330112312130010-0032100122031021-0111232221032121-0030201333002210-1203121111323232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_specification.validation_disabled` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [api_specification](resources--cdn_loadbalancer--reference--group-006.md#canonical-3010132111313002-2022320021321211-2210100222202312-1123233131321130-0220000321012023-1012110310113330-2331213333023232-1123012202210312)
- api_specification.validation_disabled

<a id="canonical-3131132030033121-2100010023021023-1230302323013333-0023230220231310-2132313112111331-2102110332213100-2313021330111100-1023122333131310"></a>

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
validation_disabled = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0320211303002011-0303200110322030-3020012231113003-3312221223200013-2030312203122030-1312030231213021-0122121213001310-1020121102323120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `app_firewall` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- app_firewall

<a id="canonical-1331220121113202-3212233013121033-2001031102032201-0220032123121012-3130200201332011-0103210132120031-3303201013313103-3032033300231030"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: app\_firewall, disable\_waf; Default: disable\_waf\] Type establishes a direct reference
from one object(the referrer) to another(the referred). Such a reference is in form of
tenant/namespace/name.

Additional upstream details:

This type establishes a direct reference from one object(the referrer) to another(the referred).

Receipt-pinned upstream constraints:

```json
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

- [app_firewall](resources--cdn_loadbalancer--reference--group-007.md#canonical-1331220121113202-3212233013121033-2001031102032201-0220032123121012-3130200201332011-0103210132120031-3303201013313103-3032033300231030)
- [disable_waf](resources--cdn_loadbalancer--reference--group-010.md#canonical-0330321120112012-0021333212231113-2232020002320311-0200223111331021-1210113213102002-2101003211223303-0223321332020310-1010211300000211)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
app_firewall {
  # Configure direct properties listed below.
}
```

<a id="canonical-3110302200021000-2331112212211200-1112131000322331-2102101001202332-3011330031301023-0202320333022121-0200331210321111-0212220031312033"></a>

### Direct properties for `app_firewall`

<a id="canonical-3223010300031210-0121022000223221-0300320222003033-1033322100012000-0320000201130312-1222103223121303-1102130313200310-3212213000333302"></a>

#### `app_firewall.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3200023100203332-0111110133131320-0120131323021231-2021103100200320-0232011023231030-0121211102303101-1220310121120311-1310301013101303"></a>

<a id="canonical-2113120222122220-2321331330213021-2111300330230200-2321333320311300-1110322201111112-3102033022232112-3130123210322020-3210302022000213"></a>

#### `app_firewall.namespace` property

Type: `"string"`. Optional, Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-3120230023012102-3233010130212311-0021322313020213-2010211322230011-3202201103230211-3003233323032311-0121320021202312-1132022332121021"></a>

<a id="canonical-0310330321331231-0003122211212132-2301101311230122-2333123203010220-2211222000123120-1222230322123032-3301320200232300-0212103113331233"></a>

#### `app_firewall.tenant` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- blocked_clients

<a id="canonical-1123011321312302-2103310132112201-1202312233031220-1221322232321300-1102203201112233-2101000110010202-2133203310113200-2113332102312211"></a>

Type: `"object"`. list nested block, Optional.

Define rules to block IP Prefixes or AS numbers.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
blocked_clients {
  # Configure direct properties listed below.
}
```

<a id="canonical-1121203012320212-3100310333213221-0031331232232133-1123301030212022-3310302332030210-3130121022213212-3131030110102320-1001121123201020"></a>

### Direct properties for `blocked_clients`

<a id="canonical-3200300301020233-0200132111001300-1212103333033312-1123230013230003-3321001302202123-0301121313132203-3223231311231212-0223010120012312"></a>

#### `blocked_clients.actions` property

Type: `["list", "string"]`. Optional.

\[Enum:
SKIP\_PROCESSING\_WAF|SKIP\_PROCESSING\_BOT|SKIP\_PROCESSING\_MUM|SKIP\_PROCESSING\_IP\_REPUTATION|SKIP\_PROCESSING\_API\_PROTECTION|SKIP\_PROCESSING\_OAS\_VALIDATION|SKIP\_PROCESSING\_DDOS\_PROTECTION|SKIP\_PROCESSING\_THREAT\_MESH|SKIP\_PROCESSING\_MALWARE\_PROTECTION\]
Actions that should be taken when client identifier matches the rule. Possible values are
\`SKIP\_PROCESSING\_WAF\`, \`SKIP\_PROCESSING\_BOT\`, \`SKIP\_PROCESSING\_MUM\`,
\`SKIP\_PROCESSING\_IP\_REPUTATION\`, \`SKIP\_PROCESSING\_API\_PROTECTION\`,
\`SKIP\_PROCESSING\_OAS\_VALIDATION\`, \`SKIP\_PROCESSING\_DDOS\_PROTECTION\`,
\`SKIP\_PROCESSING\_THREAT\_MESH\`, \`SKIP\_PROCESSING\_MALWARE\_PROTECTION\`. Defaults to
\`SKIP\_PROCESSING\_WAF\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2002323300110310-1001302220212212-0212302032032313-2300202032112230-2300121231321133-1210112322013132-1201312101323122-3332030000111122"></a>

<a id="canonical-1012320111201300-1002300103001212-3112201322321033-3211230123203213-2213221323300300-3333103302212122-2132213123321012-0213013220032200"></a>

#### `blocked_clients.as_number` property

Type: `"number"`. Optional.

Exclusive with \[http\_header ip\_prefix IPv6\_prefix user\_identifier\] RFC 6793 defined 4-byte AS
number.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 401308,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.lte": "401308"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "1",
    "ves.io.schema.rules.uint32.lte": "401308"
  }
}
```

- [bot_skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-2211002221013310-2033110000213011-1333322000223330-2200022132321222-1233130322233302-1111312201230313-3300012022030101-1021123221203211): complete subsection reference.

<a id="canonical-3302220131220122-0230031302003110-2202111231111132-1323333030010332-1201112130223103-3031010323100313-2003020013203122-2002301211223221"></a>

<a id="canonical-0010230023110301-3010010131130310-1012311330212203-3110001332300222-2133222303011223-3001310101320311-3032213121132113-2330012132301323"></a>

#### `blocked_clients.expiration_timestamp` property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Additional upstream details:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [http_header](resources--cdn_loadbalancer--reference--group-007.md#canonical-0001133200202201-0112321323013102-2223132032103103-3300233111313011-3331213132131322-0202013020301310-3320013131330120-0022033221232033): complete subsection reference.

<a id="canonical-2111332213303312-1011312010230021-0321031133021320-2111203022120023-0131112333122033-1131233311032200-3212001112121032-1000300013321033"></a>

<a id="canonical-2200103021120223-3031201302321303-2012033010103120-0200212031011023-0013313112120030-2003022102131003-1303120232021002-0330011011220102"></a>

#### `blocked_clients.ip_prefix` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header IPv6\_prefix user\_identifier\] IPv4 prefix string.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2121111133231221-1230112102232003-1102021212102331-3102330101223033-0010213202300101-3012331230332022-1320320113302200-1220303002331011"></a>

<a id="canonical-3000321120300303-1021133321211132-2022210322113122-3112230100330001-2130310312013013-2100001233210022-2003100101301020-2310023233233013"></a>

#### `blocked_clients.ipv6_prefix` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix user\_identifier\] IPv6 prefix string.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [metadata](resources--cdn_loadbalancer--reference--group-007.md#canonical-2200200000300132-1332301203032202-2100200211211313-3111120203213220-3330122020102220-1312102213133013-3200101022303223-1200011131230313): complete subsection reference.

- [skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-2202101210222131-0232132212311320-2112012121130132-2133121132210112-1132033031020303-2231231321333131-0211031330112200-2221130212000313): complete subsection reference.

<a id="canonical-1012323301131130-0112231011022021-1133020223220201-0023132133133003-1323321321112122-0231303302323331-1020123222312022-0202122030103001"></a>

<a id="canonical-3130201330230123-1103331021212023-1133121121310211-3003302201232222-1212310323100230-1222212200010100-0100213100311011-3212033032231023"></a>

#### `blocked_clients.user_identifier` property

Type: `"string"`. Optional.

Exclusive with \[as\_number http\_header ip\_prefix IPv6\_prefix\] Identify user based on user
identifier. User identifier value needs to be copied from security event.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

- [waf_skip_processing](resources--cdn_loadbalancer--reference--group-007.md#canonical-1320031132031333-0322332200020113-1301203133232332-3232323223120032-1000133112100222-2233323333133332-1003310032331111-3012013120333133): complete subsection reference.

<a id="canonical-2211002221013310-2033110000213011-1333322000223330-2200022132321222-1233130322233302-1111312201230313-3300012022030101-1021123221203211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.bot_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-007.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.bot_skip_processing

<a id="canonical-3313010011113230-1222023103023033-2133313112312302-1331023023232002-3021333313311231-0301102101133310-2233020320012333-3032331221003101"></a>

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
bot_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0001133200202201-0112321323013102-2223132032103103-3300233111313011-3331213132131322-0202013020301310-3320013131330120-0022033221232033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.http_header` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-007.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.http_header

<a id="canonical-3120231002333330-3130120231322101-0000122210201230-3111221321212223-2223010000130220-3111102131332031-0332022133120312-0020222022021301"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for http header.

Additional upstream details:

Request header name and value pairs.

Receipt-pinned upstream constraints:

```json
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
http_header {
  # Configure direct properties listed below.
}
```

<a id="canonical-0210211122303312-1031010031123333-1320213001321220-0232222223231112-2012100030101310-1200313331113300-2031303033113130-1003121023023231"></a>

### Direct properties for `blocked_clients.http_header`

- [headers](resources--cdn_loadbalancer--reference--group-007.md#canonical-0013310131033210-0230332203001000-2313311132103321-3003302333213302-3010002321103021-0303222132010013-2311310011332301-1002122313030003): complete subsection reference.

<a id="canonical-0013310131033210-0230332203001000-2313311132103321-3003302333213302-3010002321103021-0303222132010013-2311310011332301-1002122313030003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.http_header.headers` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-007.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- [blocked_clients.http_header](resources--cdn_loadbalancer--reference--group-007.md#canonical-0001133200202201-0112321323013102-2223132032103103-3300233111313011-3331213132131322-0202013020301310-3320013131330120-0022033221232033)
- blocked_clients.http_header.headers

<a id="canonical-1220232031210330-3121022100110200-2102311301031231-3330000202120012-3101230113022212-1321033032131203-3020123130322001-1022210021303110"></a>

Type: `"object"`. list nested block, Optional.

List of HTTP header name and value pairs.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

Terraform syntax:

```terraform
headers {
  # Configure direct properties listed below.
}
```

<a id="canonical-0130113121310131-0031022332030233-3110331103212203-2222222021111010-0220303221021030-3110010212211012-2003013031122221-0320122132220130"></a>

### Direct properties for `blocked_clients.http_header.headers`

<a id="canonical-2221213110320123-2122322303103320-0033101300310210-0212331111231321-0212031113011313-0133023011222200-0120012330000012-1321303132133030"></a>

#### `blocked_clients.http_header.headers.exact` property

Type: `"string"`. Optional.

Exclusive with \[presence regular expression\] Header value to match exactly.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.not_empty": "true"
  }
}
```

<a id="canonical-3033323300322331-0322300011001333-2130230110233210-0130302210323231-3010123323033321-0203210100012220-2303102200320233-1023020131121313"></a>

<a id="canonical-3122312110222112-2210211331200100-2113311310302311-2113012103030201-3102322120121120-0213222101023010-0323101120022123-2022323120331112"></a>

#### `blocked_clients.http_header.headers.invert_match` property

Type: `"bool"`. Optional.

Invert the result of the match to detect missing header or non-matching value.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2233220021320221-3221320031033311-1202332000033310-3100103020220330-1103220330033233-3101213210032132-1003032330322302-0301223331010120"></a>

<a id="canonical-3323103133203020-0211022203031000-3113213011013321-3001311101030123-2110101023012300-0133010122023102-3322200200320001-3310131210322322"></a>

#### `blocked_clients.http_header.headers.name` property

Type: `"string"`. Optional.

Name. Name of the header.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3221201103123201-1331122120000311-1110203331132312-2020112230303223-1122320322201013-0313233202112000-1333203221303113-1111302012133000"></a>

<a id="canonical-1023002320110031-1222002003233032-1300032211231302-3021110313201010-1202111111111122-1001211112320111-3231012313103033-1211313010221323"></a>

#### `blocked_clients.http_header.headers.presence` property

Type: `"bool"`. Optional.

Exclusive with \[exact regular expression\] If true, check for presence of header.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0011103231301130-3033022323022330-2030212230333001-2232212122102033-3233102211123332-1302302223332000-0311230302030321-1313320103212332"></a>

<a id="canonical-2310021211112122-3111330301033122-1113132301023002-2211100323033223-1030122321200000-1013331231012201-1130211300331032-2102302201211200"></a>

#### `blocked_clients.http_header.headers.regex` property

Type: `"string"`. Optional.

Exclusive with \[exact presence\] regular expression match of the header value in re2 format.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="canonical-2200200000300132-1332301203032202-2100200211211313-3111120203213220-3330122020102220-1312102213133013-3200101022303223-1200011131230313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.metadata` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-007.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.metadata

<a id="canonical-3233123130130120-0000022110120132-1311332230033000-3111220231133123-2122320212130223-2301111301110032-2000133100022100-1112110330133203"></a>

Type: `"object"`. single nested block, Optional.

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

Terraform syntax:

```terraform
metadata {
  # Configure direct properties listed below.
}
```

<a id="canonical-1123023332302112-2123112113301121-0222132012013123-3023011122323100-0320330313232232-2201033110323033-0032100320123223-3222203112033021"></a>

### Direct properties for `blocked_clients.metadata`

<a id="canonical-0213203010313133-1033010022303031-3123302222203032-1130013310231322-0013211121303303-2310323232230322-2023232230033013-2333321023312032"></a>

#### `blocked_clients.metadata.description_spec` property

Type: `"string"`. Optional.

Description. Human readable description.

<a id="canonical-2132320113103022-3133132302312111-3031310022033222-2023110312313212-1130321123311303-0200032132311120-3232312211003332-1233032201302132"></a>

<a id="canonical-1223320200102201-3021303133003203-2200330010001203-3111101020332003-3122223002210002-0202131321133113-2233211021130203-1102213212113200"></a>

#### `blocked_clients.metadata.name` property

Type: `"string"`. Optional.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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

<a id="canonical-2202101210222131-0232132212311320-2112012121130132-2133121132210112-1132033031020303-2231231321333131-0211031330112200-2221130212000313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-007.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.skip_processing

<a id="canonical-3302200001332131-1100010112223102-2122233303001222-2203031132213320-2113003122030312-0310110332102012-3321101323200111-1132203312310303"></a>

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
skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320031132031333-0322332200020113-1301203133232332-3232323223120032-1000133112100222-2233323333133332-1003310032331111-3012013120333133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `blocked_clients.waf_skip_processing` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [blocked_clients](resources--cdn_loadbalancer--reference--group-007.md#canonical-0101200112203333-2313100322333123-1133301001021123-0132230132312103-2222120211222232-1330312333203021-1213021020200201-2312220003310232)
- blocked_clients.waf_skip_processing

<a id="canonical-2123212213310320-0333131113101212-0012232020022321-2031011023320302-0002220200332121-3223033310222233-1320012023032010-3300323022113013"></a>

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
waf_skip_processing = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- bot_defense

<a id="canonical-2000311131220203-1013013112311122-1031332010101300-3222113321032031-2112031330230210-2033321311130203-3322121100201010-3202000031000203"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

Terraform syntax:

```terraform
bot_defense {
  # Configure direct properties listed below.
}
```

<a id="canonical-2003200012320130-0221013011001110-2232203102020200-2333303121120331-1000013320331331-2200030230011110-3301020213331021-2222010100211130"></a>

### Direct properties for `bot_defense`

- [disable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-2110323220222022-1200020013313230-2012021003132203-1202331100201102-1112222110003322-2212332311211210-2030033313013123-1300002003210011): complete subsection reference.

- [enable_cors_support](resources--cdn_loadbalancer--reference--group-007.md#canonical-1130023222310301-2001120103202032-3003321211323123-1123113313100301-1031313022330103-2301233102331112-1310020101212311-2122332220301103): complete subsection reference.

- [policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111): complete subsection reference.

<a id="canonical-2002232020020132-0323030221013223-0202010111023322-2120331311013001-2113003212023301-2122333110101121-1120200011132003-3121102032002100"></a>

<a id="canonical-2333020020130101-3121302003312230-2013301100210332-2000102123203133-0313201021200131-2130231131130002-0231120112221112-2113203121303121"></a>

#### `bot_defense.regional_endpoint` property

Type: `"string"`. Optional.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2000101121230031-1332003033312230-3323103101322101-1300100130120012-2221032220023033-0303312022233223-3302332021301211-2333211213201010"></a>

<a id="canonical-3130231102101110-2230210103023011-1330132120212203-2231201313332221-3312131132110321-0033312102300133-1300100113313323-1211100110211201"></a>

#### `bot_defense.timeout` property

Type: `"number"`. Optional.

The timeout for the inference check, in milliseconds.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
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
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

<a id="canonical-2110323220222022-1200020013313230-2012021003132203-1202331100201102-1112222110003322-2212332311211210-2030033313013123-1300002003210011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.disable_cors_support` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- bot_defense.disable_cors_support

<a id="canonical-1000113332122230-0221302010230000-1030001220302000-1212233221112202-2010101103111303-3121030310002302-0321112310023332-0230323232031231"></a>

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
disable_cors_support = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1130023222310301-2001120103202032-3003321211323123-1123113313100301-1031313022330103-2301233102331112-1310020101212311-2122332220301103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.enable_cors_support` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- bot_defense.enable_cors_support

<a id="canonical-0223001222223213-0012200330232312-2211033301003222-2333303023022323-3032330123203130-2131132212221330-3013210323131231-0033023301033232"></a>

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
enable_cors_support = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- bot_defense.policy

<a id="canonical-2101120103201033-1320310231012320-1011310220333012-1330200013111022-1112311320312332-2020223102331030-2313230033112203-0100000320001210"></a>

Type: `"object"`. single nested block, Optional.

This defines various configuration OPTIONS for Bot Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

Terraform syntax:

```terraform
policy {
  # Configure direct properties listed below.
}
```

<a id="canonical-0232200213321030-2313203123231112-1203211030330220-0303010320203312-3332210231032333-1310303300221200-1130023022022231-1322021330213132"></a>

### Direct properties for `bot_defense.policy`

- [disable_js_insert](resources--cdn_loadbalancer--reference--group-007.md#canonical-3113032133121203-2001232210033102-3203122102001200-2032012232202212-1230223221132133-2103202223101322-0013100213101123-3121201130103200): complete subsection reference.

- [disable_mobile_sdk](resources--cdn_loadbalancer--reference--group-007.md#canonical-0120322022300212-0310022310212003-3032313020212132-1021133110312313-3031102102101121-3031010112312113-0133123103012203-0231310000332213): complete subsection reference.

<a id="canonical-0012113233011302-3002222312100303-3313220210303232-3121203022010220-3030311012011012-1123023130301301-1020102011210133-0330230023120312"></a>

<a id="canonical-2122200331221320-0020332223030120-1210203131102310-0103323310210122-3003112122303021-3211100003310223-0232020212333300-3222301110031120"></a>

#### `bot_defense.policy.javascript_mode` property

Type: `"string"`. Optional.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2113232011012300-0032031330323322-1332031032313213-3332022221222123-0230232233332202-1301210333130220-0310301332023322-3023320221020312"></a>

<a id="canonical-2301120220002130-0133201211102312-0311000013000133-3010201213312221-2111112113303230-3020231002122210-0310332212300233-1021200110002330"></a>

#### `bot_defense.policy.js_download_path` property

Type: `"string"`. Optional.

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](resources--cdn_loadbalancer--reference--group-007.md#canonical-0211123212103031-2332332103330000-2013300321210120-1330332123230233-3030322322220130-2123121303221223-1211202023303210-0113321012012302): complete subsection reference.

- [js_insert_all_pages_except](resources--cdn_loadbalancer--reference--group-008.md#canonical-2033123012202321-1211101311022113-1201002111032111-3210022010212021-3321310021201313-3202021033020030-1101021222300022-1110131112213111): complete subsection reference.

- [js_insertion_rules](resources--cdn_loadbalancer--reference--group-008.md#canonical-3101332010211112-3031030103123313-2120311030130113-1132001221131033-1103310102132033-0213103111312323-3003202121131230-3323212032330321): complete subsection reference.

- [mobile_sdk_config](resources--cdn_loadbalancer--reference--group-008.md#canonical-0211110332021202-1223322002233230-1010121033102111-2333021322220011-0303302202101022-2020222312011102-1313020332132132-3230110202201322): complete subsection reference.

- [protected_app_endpoints](resources--cdn_loadbalancer--reference--group-008.md#canonical-3300001103213222-2032130220300122-0220101231033121-3323032100130212-2013030302131232-2300203120320330-3202133021002303-3000110003101013): complete subsection reference.

<a id="canonical-3113032133121203-2001232210033102-3203122102001200-2032012232202212-1230223221132133-2103202223101322-0013100213101123-3121201130103200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_js_insert` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.disable_js_insert

<a id="canonical-2031101020003311-2313303032123003-1122010212122122-2011012202211320-1310003300200021-3131333332201113-3030100233003211-2222332131210001"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable js insert.

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
disable_js_insert = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0120322022300212-0310022310212003-3032313020212132-1021133110312313-3031102102101121-3031010112312113-0133123103012203-0231310000332213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.disable_mobile_sdk` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.disable_mobile_sdk

<a id="canonical-2223012123201123-2211003312020332-0321030312010130-1030023102232000-0101331331033303-0322330020102131-0223323032110203-0220121013321321"></a>

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
disable_mobile_sdk = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0211123212103031-2332332103330000-2013300321210120-1330332123230233-3030322322220130-2123121303221223-1211202023303210-0113321012012302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `bot_defense.policy.js_insert_all_pages` properties

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md#canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323)
- [Property reference](resources--cdn_loadbalancer--reference--group-001.md#canonical-1012133122032113-1010211312001122-2210220230020233-3330322220122132-1302030221223031-1122222203331001-2113210011222311-0201322132233330)
- [bot_defense](resources--cdn_loadbalancer--reference--group-007.md#canonical-1231232212013132-2130001032222220-0320033110013221-0131313023101120-1102230212130003-0122310320031322-3013233223020302-2211010022010100)
- [bot_defense.policy](resources--cdn_loadbalancer--reference--group-007.md#canonical-3233222220320220-2012123201001112-2001111122200333-3000232203203221-3010100330312230-1021020232003031-0231321333322111-1012011130013111)
- bot_defense.policy.js_insert_all_pages

<a id="canonical-2102110023323330-2003211321210003-0132022131012132-0011100202211231-3020021101010021-0011100323310320-0031200123123323-3202301022003312"></a>

Type: `"object"`. single nested block, Optional.

Insert Bot Defense JavaScript in all pages.

Receipt-pinned upstream constraints:

```json
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
js_insert_all_pages {
  # Configure direct properties listed below.
}
```

<a id="canonical-3333212132300122-1330022012330202-0202112313312300-2131102100313232-2002230011103133-3133230111321332-3013032302011333-0313133220203322"></a>

### Direct properties for `bot_defense.policy.js_insert_all_pages`

<a id="canonical-3212021200302102-2100333220210112-1033111021311232-0013220123110301-3220012021113121-2310302331320012-1210012001123131-1220330330012113"></a>

#### `bot_defense.policy.js_insert_all_pages.javascript_location` property

Type: `"string"`. Optional.

\[Enum: AFTER\_HEAD|AFTER\_TITLE\_END|BEFORE\_SCRIPT\] All inside networks. Insert JavaScript after
&lt;HEAD&gt; tag Insert JavaScript after &lt;/title&gt; tag. Insert JavaScript before first tag.
Possible values are \`AFTER\_HEAD\`, \`AFTER\_TITLE\_END\`, \`BEFORE\_SCRIPT\`. Defaults to
\`AFTER\_HEAD\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "AFTER_HEAD",
  "enum": [
    "AFTER_HEAD",
    "AFTER_TITLE_END",
    "BEFORE_SCRIPT"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
