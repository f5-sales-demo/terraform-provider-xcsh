---
page_title: "xcsh_sensitive_data_policy reference"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_sensitive_data_policy reference."
---

# xcsh_sensitive_data_policy reference

<a id="canonical-0203133212301203-1132112323323302-1312032000023031-3213122123301323-0203001200123121-2232023331202233-1212100131131013-0000331232331001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-0021133210323100-2012112012231021-2330030003220132-2211002111322321-3222123123113011-2131230331102322-2133003312231310-0330022320311010)
- Property reference

<a id="canonical-2313012111321110-0212200302202123-1002012222203113-3230000002112312-3131010113303002-0202331103310200-1233020000023210-1031020211030310"></a>

### Direct properties for `xcsh_sensitive_data_policy`

<a id="canonical-0310223032011130-1010310203221232-1001023211231031-3322332103302123-3020201123320210-2303133303211233-0221333122210210-3022013031300323"></a>

#### `annotations` property

Type: `["map", "string"]`. Optional.

Annotations is an unstructured key-value map stored with a resource that may be set by external
tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when
modifying objects.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
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
      "ves.io.schema.rules.map.values.string.max_len": "1024",
      "ves.io.schema.rules.map.values.string.min_len": "1"
    },
    "values": {
      "maxLength": 1024,
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
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.map.keys.string.max_len": "64",
    "ves.io.schema.rules.map.keys.string.min_len": "1",
    "ves.io.schema.rules.map.values.string.max_len": "1024",
    "ves.io.schema.rules.map.values.string.min_len": "1"
  }
}
```

<a id="canonical-0111301323222301-2302310033002320-3001200101313221-0002301113000033-0330203031002303-0002312301111130-0122031303332122-3332103333211030"></a>

<a id="canonical-1131203303223011-3121230203123100-2022001300132231-2111333123031002-3012310333002232-3030103313001002-3111310311102320-0002211001100320"></a>

#### `compliances` property

Type: `["list", "string"]`. Optional, Computed.

\[Enum:
GDPR|CCPA|PIPEDA|LGPD|DPA\_UK|PDPA\_SG|APPI|HIPAA|CPRA\_2023|CPA\_CO|SOC2|PCI\_DSS|ISO\_IEC\_27001|ISO\_IEC\_27701|EPRIVACY\_DIRECTIVE|GLBA|SOX\]
Select relevant compliance frameworks, such as GDPR, HIPAA, or PCI-DSS, to ensure monitoring under
your sensitive data discovery. Defaults to \`\[\]\`. Server applies default when omitted. Possible
values are \`GDPR\`, \`CCPA\`, \`PIPEDA\`, \`LGPD\`, \`DPA\_UK\`, \`PDPA\_SG\`, \`APPI\`, \`HIPAA\`,
\`CPRA\_2023\`, \`CPA\_CO\`, \`SOC2\`, \`PCI\_DSS\`, \`ISO\_IEC\_27001\`, \`ISO\_IEC\_27701\`,
\`EPRIVACY\_DIRECTIVE\`, \`GLBA\`, \`SOX\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(17),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 17,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 17,
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
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.enum.defined_only": "true",
    "ves.io.schema.rules.repeated.items.enum.not_in": "[0]",
    "ves.io.schema.rules.repeated.max_items": "17",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

- [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-2100333133332323-2001302233101200-2010020300300200-0130311032100233-1233203010002032-3302003112301112-0311303331023330-1331100121002231): complete subsection reference.

<a id="canonical-3200010012202203-2233021122221330-0232221100110033-3213210330332231-0320012300320113-0002200002021123-3103000102012302-1221222323202033"></a>

<a id="canonical-1311203333032333-3130213000230332-2111100011332020-0210331312022230-2022221211131302-1233111332303211-2122202222322030-3211212303001311"></a>

#### `description` property

Type: `"string"`. Optional.

Human readable description for the object.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1200,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1200
    },
    "category": "discovery",
    "characterSet": {
      "description": "Free text with UTF-8 support"
    },
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1200,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 0
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "1200"
  }
}
```

<a id="canonical-2021102013133230-3220110223121332-2100130302201231-0033033101323101-1102203231222102-2122003333301110-3322200223221330-2331221120120002"></a>

<a id="canonical-3302323132033002-0132002001103022-1310201103011231-2322122322111033-2103313101223003-0103110111121112-3330012313231002-3013332100313211"></a>

#### `disable` property

Type: `"bool"`. Optional.

A value of true administratively disables the object.

Additional upstream details:

A value of true will administratively disable the object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2132101032002122-3332131200222302-0232130313320123-1021333120213200-1013101030233312-2013021023110233-2322211133131012-2311030010313212"></a>

<a id="canonical-0322111130222321-1201312110300310-2221332112322221-3122103101333300-0103111012000330-0320322023332321-1202231001123301-2113012111311012"></a>

#### `disabled_predefined_data_types` property

Type: `["list", "string"]`. Optional, Computed.

Select which pre-configured data types to disable, disabled data types will not be shown as
sensitive in the API discovery. Defaults to \`\[\]\`. Server applies default when omitted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.List{
  listvalidator.SizeAtMost(100),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 100,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 100,
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
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_len": "64",
    "ves.io.schema.rules.repeated.max_items": "100",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3312121223101233-3220112112013002-3100310003322300-3223333132303320-1303133030101030-1021221302302130-3213032133330103-0321212333121231"></a>

<a id="canonical-2301111220100111-1302130031022020-1033222011233003-0002200001311223-1200303001332223-0131230031133032-2233013130100321-1302230322330002"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0313233121002010-1112312030121331-2000100321111112-0013130201231310-0323030113112322-3321113021202012-2301233313133202-2203102211313022"></a>

<a id="canonical-0213113000130001-0303113313033020-3000232322120012-2332300000031313-0023333102312020-1023020020200202-2010202121321133-3101300011202103"></a>

#### `labels` property

Type: `["map", "string"]`. Optional.

Labels is a user defined key-value map that can be attached to resources for organization and
filtering.

Additional upstream details:

Map of string keys and values that can be used to organize and categorize (scope and select) objects
as chosen by the user. Values specified here will be used by selector expression.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2230022303322020-3013222032321212-1231100113330002-0222220123130031-3300112212022210-3233231001132233-0121322212133001-1300331210310222"></a>

<a id="canonical-0021330210033322-2330000212101003-2011311113332230-2323133231331302-1230311211132113-3111233302032330-0031120213000310-2323101200101212"></a>

#### `name` property

Type: `"string"`. Required.

Name of the Sensitive Data Policy. Must be unique within the namespace.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NameValidator(),
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
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

<a id="canonical-1320103100302013-2131230010112201-3032132311331231-2130212202011110-1203332012221311-2330021232331313-3332230330013111-0201231023121330"></a>

<a id="canonical-3032210211313333-0012323203132302-1132203322232200-0332313003321100-0022120310231032-2001003212032013-3320000332112330-2001303101002300"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the Sensitive Data Policy is created.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.String{
  validators.NamespaceValidator(),
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

- [timeouts](resources--sensitive_data_policy--reference--group-001.md#canonical-1221132322310121-0102330100311031-3320030111012032-2013213020333212-2133232303022122-3200122013033322-1133223032312132-1133030310333012): complete subsection reference.

<a id="canonical-2323130223211123-3021200110303010-3123032222130221-0022011111203312-3233131301303310-1220003310112012-2310120000012232-3113301323200302"></a>

### All schema paths for `xcsh_sensitive_data_policy`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](resources--sensitive_data_policy--reference--group-001.md#canonical-0310223032011130-1010310203221232-1001023211231031-3322332103302123-3020201123320210-2303133303211233-0221333122210210-3022013031300323) |
| `compliances` | [compliances](resources--sensitive_data_policy--reference--group-001.md#canonical-0111301323222301-2302310033002320-3001200101313221-0002301113000033-0330203031002303-0002312301111130-0122031303332122-3332103333211030) |
| `custom_data_types` | [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-2313223331030010-2222101033102223-2130330210231331-1001011103211303-2230033333231320-2130033332021330-3113030021123233-1132212230001230) |
| `custom_data_types.custom_data_type_ref` | [custom_data_types.custom_data_type_ref](resources--sensitive_data_policy--reference--group-001.md#canonical-2100332213220120-3000333012200323-2110002003111231-0322013123203203-2211323111221123-1121011222101202-1131323221113011-3310123221332021) |
| `custom_data_types.custom_data_type_ref.name` | [custom_data_types.custom_data_type_ref.name](resources--sensitive_data_policy--reference--group-001.md#canonical-2230322000130133-0112133021203233-3013123313231103-1230121310300223-2123322330333312-2013333230023111-2112332001200021-2113213020231233) |
| `custom_data_types.custom_data_type_ref.namespace` | [custom_data_types.custom_data_type_ref.namespace](resources--sensitive_data_policy--reference--group-001.md#canonical-0031310310331031-0101132030101330-2200221222023320-3211032202211312-3111101301301030-0222320310312100-1302113322312321-0132102221113111) |
| `custom_data_types.custom_data_type_ref.tenant` | [custom_data_types.custom_data_type_ref.tenant](resources--sensitive_data_policy--reference--group-001.md#canonical-0102110020301022-1233132321121213-3023002010320122-0012031233122003-3222311111131202-1000133302330301-2312321122301102-0211312322331302) |
| `description` | [description](resources--sensitive_data_policy--reference--group-001.md#canonical-3200010012202203-2233021122221330-0232221100110033-3213210330332231-0320012300320113-0002200002021123-3103000102012302-1221222323202033) |
| `disable` | [disable](resources--sensitive_data_policy--reference--group-001.md#canonical-2021102013133230-3220110223121332-2100130302201231-0033033101323101-1102203231222102-2122003333301110-3322200223221330-2331221120120002) |
| `disabled_predefined_data_types` | [disabled_predefined_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-2132101032002122-3332131200222302-0232130313320123-1021333120213200-1013101030233312-2013021023110233-2322211133131012-2311030010313212) |
| `id` | [ID](resources--sensitive_data_policy--reference--group-001.md#canonical-3312121223101233-3220112112013002-3100310003322300-3223333132303320-1303133030101030-1021221302302130-3213032133330103-0321212333121231) |
| `labels` | [labels](resources--sensitive_data_policy--reference--group-001.md#canonical-0313233121002010-1112312030121331-2000100321111112-0013130201231310-0323030113112322-3321113021202012-2301233313133202-2203102211313022) |
| `name` | [name](resources--sensitive_data_policy--reference--group-001.md#canonical-2230022303322020-3013222032321212-1231100113330002-0222220123130031-3300112212022210-3233231001132233-0121322212133001-1300331210310222) |
| `namespace` | [namespace](resources--sensitive_data_policy--reference--group-001.md#canonical-1320103100302013-2131230010112201-3032132311331231-2130212202011110-1203332012221311-2330021232331313-3332230330013111-0201231023121330) |
| `timeouts` | [timeouts](resources--sensitive_data_policy--reference--group-001.md#canonical-3320111330210132-1003020321123132-2220333231110212-1030112122100300-2210003100302333-1121310121131220-1020013310321301-1133130202011112) |
| `timeouts.create` | [timeouts.create](resources--sensitive_data_policy--reference--group-001.md#canonical-1112220111022032-2212312200221321-2301011232031030-0111033322110113-1300221313003132-0223333001020222-2102203101132012-3023203200330320) |
| `timeouts.delete` | [timeouts.delete](resources--sensitive_data_policy--reference--group-001.md#canonical-2213231003211010-2102223020131220-3310232310201122-3213201000113233-2133020001212033-0111310003323230-1100103203103332-0232102212103133) |
| `timeouts.read` | [timeouts.read](resources--sensitive_data_policy--reference--group-001.md#canonical-2122001113013133-1011101322130123-1303300130213130-0113322310310222-0030111033213110-2232303030030030-3332010300122131-3130031311110101) |
| `timeouts.update` | [timeouts.update](resources--sensitive_data_policy--reference--group-001.md#canonical-1220212313322112-1022032133212000-1111022230302133-1221132010020130-2323301012020121-2223320212213233-2133332230112211-3110222000113001) |

<a id="canonical-2100333133332323-2001302233101200-2010020300300200-0130311032100233-1233203010002032-3302003112301112-0311303331023330-1331100121002231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_data_types` properties

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-0021133210323100-2012112012231021-2330030003220132-2211002111322321-3222123123113011-2131230331102322-2133003312231310-0330022320311010)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-0203133212301203-1132112323323302-1312032000023031-3213122123301323-0203001200123121-2232023331202233-1212100131131013-0000331232331001)
- custom_data_types

<a id="canonical-2313223331030010-2222101033102223-2130330210231331-1001011103211303-2230033333231320-2130033332021330-3113030021123233-1132212230001230"></a>

Type: `"object"`. list nested block, Optional.

Select your custom data types to be monitored in the API discovery. Defaults to \`\[\]\`. Server
applies default when omitted.

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
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
custom_data_types {
  # Configure direct properties listed below.
}
```

<a id="canonical-2130211033120232-1202321120223310-1132313333312102-0102100111002102-2310202232031232-2120312310020010-3211231032121113-3123132133332111"></a>

### Direct properties for `custom_data_types`

- [custom_data_type_ref](resources--sensitive_data_policy--reference--group-001.md#canonical-1001012313022121-1302233000031322-2330131011110312-2023002223232332-1311030210102230-1232310120113222-3113033003213333-2220310120302000): complete subsection reference.

<a id="canonical-1001012313022121-1302233000031322-2330131011110312-2023002223232332-1311030210102230-1232310120113222-3113033003213333-2220310120302000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `custom_data_types.custom_data_type_ref` properties

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-0021133210323100-2012112012231021-2330030003220132-2211002111322321-3222123123113011-2131230331102322-2133003312231310-0330022320311010)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-0203133212301203-1132112323323302-1312032000023031-3213122123301323-0203001200123121-2232023331202233-1212100131131013-0000331232331001)
- [custom_data_types](resources--sensitive_data_policy--reference--group-001.md#canonical-2100333133332323-2001302233101200-2010020300300200-0130311032100233-1233203010002032-3302003112301112-0311303331023330-1331100121002231)
- custom_data_types.custom_data_type_ref

<a id="canonical-2100332213220120-3000333012200323-2110002003111231-0322013123203203-2211323111221123-1121011222101202-1131323221113011-3310123221332021"></a>

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
custom_data_type_ref {
  # Configure direct properties listed below.
}
```

<a id="canonical-2202131301321333-2000131202322233-0212331022032021-0112130022311211-0010323323330300-2003100310331330-0023202110223222-2203230232310103"></a>

### Direct properties for `custom_data_types.custom_data_type_ref`

<a id="canonical-2230322000130133-0112133021203233-3013123313231103-1230121310300223-2123322330333312-2013333230023111-2112332001200021-2113213020231233"></a>

#### `custom_data_types.custom_data_type_ref.name` property

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

<a id="canonical-0031310310331031-0101132030101330-2200221222023320-3211032202211312-3111101301301030-0222320310312100-1302113322312321-0132102221113111"></a>

<a id="canonical-1230122330213013-1230130312213200-1023332021012213-1310221101012102-3030101130000101-1102113022303333-3231211222232011-2323302030031231"></a>

#### `custom_data_types.custom_data_type_ref.namespace` property

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

<a id="canonical-0102110020301022-1233132321121213-3023002010320122-0012031233122003-3222311111131202-1000133302330301-2312321122301102-0211312322331302"></a>

<a id="canonical-0213101203331303-1201111033331211-0221013132010330-2113020301333012-1321133211210012-0331310223113202-3213013321323101-1322133322220232"></a>

#### `custom_data_types.custom_data_type_ref.tenant` property

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

<a id="canonical-1221132322310121-0102330100311031-3320030111012032-2013213020333212-2133232303022122-3200122013033322-1133223032312132-1133030310333012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `timeouts` properties

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md#canonical-0021133210323100-2012112012231021-2330030003220132-2211002111322321-3222123123113011-2131230331102322-2133003312231310-0330022320311010)
- [Property reference](resources--sensitive_data_policy--reference--group-001.md#canonical-0203133212301203-1132112323323302-1312032000023031-3213122123301323-0203001200123121-2232023331202233-1212100131131013-0000331232331001)
- timeouts

<a id="canonical-3320111330210132-1003020321123132-2220333231110212-1030112122100300-2210003100302333-1121310121131220-1020013310321301-1133130202011112"></a>

Type: `"object"`. single nested block, Optional.

Terraform syntax:

```terraform
timeouts {
  # Configure direct properties listed below.
}
```

<a id="canonical-0320233033333000-3001032110023002-2233310221231132-1033010130300130-1322331201323122-0110110112312322-3310000230122321-1220111122031111"></a>

### Direct properties for `timeouts`

<a id="canonical-1112220111022032-2212312200221321-2301011232031030-0111033322110113-1300221313003132-0223333001020222-2102203101132012-3023203200330320"></a>

#### `timeouts.create` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).

<a id="canonical-2213231003211010-2102223020131220-3310232310201122-3213201000113233-2133020001212033-0111310003323230-1100103203103332-0232102212103133"></a>

<a id="canonical-0233211112223230-0320000031013331-2302330010313321-2231223000223321-0101300203302003-0233302032101012-0110221323123222-0100331122311331"></a>

#### `timeouts.delete` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are
saved into state before the destroy operation occurs.

<a id="canonical-2122001113013133-1011101322130123-1303300130213130-0113322310310222-0030111033213110-2232303030030030-3332010300122131-3130031311110101"></a>

<a id="canonical-0100320321233333-0112222210100123-3020112233330133-2300302030001310-1100111232231031-1333001120222330-3003123223322023-3103230013223000"></a>

#### `timeouts.read` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours). Read operations occur during any refresh or planning operation when refresh
is enabled.

<a id="canonical-1220212313322112-1022032133212000-1111022230302133-1221132010020130-2323301012020121-2223320212213233-2133332230112211-3110222000113001"></a>

<a id="canonical-2111332102320222-1132130012133233-2000212222222132-1123112032212010-0323113010020121-2130033222003212-1003020003113212-2232210220000330"></a>

#### `timeouts.update` property

Type: `"string"`. Optional.

A string that can be \[parsed as a duration\](https&#58;//pkg.go.dev/time\#ParseDuration) consisting
of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m"
(minutes), "h" (hours).
