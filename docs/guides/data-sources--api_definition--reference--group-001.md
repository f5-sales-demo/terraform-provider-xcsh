---
page_title: "xcsh_api_definition reference"
subcategory: "API Management"
description: "Complete grouped canonical reference for xcsh_api_definition reference."
---

# xcsh_api_definition reference

<a id="canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- Property reference

<a id="canonical-1320111132003203-1110012300003310-3200113221323020-1032103303112333-1102313203003012-1330301212202323-2331013021222110-1231222032220220"></a>

### Direct properties for `xcsh_api_definition`

<a id="canonical-2202201313011010-3233012022102023-2231033323312223-1133031120232133-0010013132303033-0331222203303101-0022100223131121-2101300132120200"></a>

#### `annotations` property

Type: `["map", "string"]`. Computed.

Annotations applied to this resource.

Additional upstream details:

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

- [api_inventory_exclusion_list](data-sources--api_definition--reference--group-001.md#canonical-0232001312320302-3223003321232222-3021132221111131-2322000200330320-3223333103311211-3020223300332311-3010130312300222-3002230133012130): complete subsection reference.

- [api_inventory_inclusion_list](data-sources--api_definition--reference--group-001.md#canonical-0121203120223031-2200023230310231-3112201032210320-2133301331032212-2301001131321032-1011010030320320-1303121321020311-1201003111020230): complete subsection reference.

<a id="canonical-1311213322111021-1321001312213330-2323110123311311-0012222130132232-2333011022031332-1233031300132102-3302120132013221-1332320210023310"></a>

<a id="canonical-0321022312013202-0212312122112013-0211302303001022-2303103312222313-1021211300310301-2020322030203232-3010313211302232-1113312121011331"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the APIDefinition.

Additional upstream details:

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3232103102023122-3003032310201012-2010231323210212-3033302300021101-2120032211023300-2033033330133102-0013230111202031-0002020322122202"></a>

<a id="canonical-2120302311320203-1202312102211022-1200321301333021-2230002102211120-1001012022210110-3131223001321201-3102101021222123-3330333321333001"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2121013333001000-3322312230301010-0033113012001213-2021100331010333-0013013111333321-2030333123331330-0122100220003202-1120301331221333"></a>

<a id="canonical-2032233223231122-2210200022023103-1331103202222111-3020301301231110-1301033222221121-2003003310330122-3232130000113120-1303303103030031"></a>

#### `labels` property

Type: `["map", "string"]`. Computed.

Labels applied to this resource.

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

- [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-1113312301230031-0230200300100013-1130332120030100-0010201203220110-1001312011120333-1130232132203320-3220122122131201-2330322330002012): complete subsection reference.

<a id="canonical-1110122123122110-3312133320230320-3200120332332123-2013221133200030-0130111213002122-3233121032220303-3200021022133321-2022123012230022"></a>

<a id="canonical-0133321020110323-1210012123102310-1131200223030123-0003110313320230-1003013113232001-0130300122123030-0003113103330103-3101100101231023"></a>

#### `name` property

Type: `"string"`. Required.

Name of the APIDefinition.

Additional upstream details:

This is the name of configuration object. It has to be unique within the namespace. It can only be
specified during create API and cannot be changed during replace API. The value of name has to
follow DNS-1035 format.

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1232312233010103-2020210021003013-0002212032023013-2310113330012113-2010211320001132-3331333323132320-2121233130012310-2112112020012303"></a>

<a id="canonical-3300331001231333-3013022300322223-1210100031013312-3211122223330032-1332110000230010-3130323212303003-1001313212003210-2331121202031303"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the APIDefinition exists.

Additional upstream details:

This defines the workspace within which each the configuration object is to be created. Must be a
DNS\_LABEL format. For a namespace object itself, namespace value will be ""

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
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

- [non_api_endpoints](data-sources--api_definition--reference--group-001.md#canonical-3023113020212033-0022101202103220-3000133332010231-2311201321011310-0101302320201220-2121200211211033-3311211213101113-3323033212322200): complete subsection reference.

- [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-1303030311022303-2032313222313311-1221131323011003-1031102310113210-2313022033101221-0001300200113222-2020213030301212-2210322332103133): complete subsection reference.

<a id="canonical-1210301012131200-0032323033020231-1011133221130121-3110322130123223-2023312110031000-2312233112013101-2200001012031220-2121303112220020"></a>

<a id="canonical-1211000020100011-3303230011210303-0133201031231320-3111013322313221-2122021122202002-2131033333023023-2302223031320100-1233001111113310"></a>

#### `swagger_specs` property

Type: `["list", "string"]`. Computed.

URLs of versioned OpenAPI files uploaded through Web App &amp; API Protection &gt; Files &gt;
Swagger Files. The 512-byte item limit is a URL-length limit; inline string:/// OpenAPI content is
rejected and does not create an API-definition object. Defaults to \`\[\]\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 20,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 20,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.max_bytes": "512",
    "ves.io.schema.rules.repeated.items.string.pattern": "/api/object_store/namespaces/([a-z]([-a-z0-9]*[a-z0-9])?)/stored_objects/swagger/([a-z]([-a-z0-9]*[a-z0-9])?)/(v|V)[0-9]+(-[0-9]{2}){3}$",
    "ves.io.schema.rules.repeated.max_items": "20",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1020010131120302-3203020012012013-0031010330213023-2312300310103022-0012023223212010-3123010332222111-3000211122130203-0322101100220113"></a>

### All schema paths for `xcsh_api_definition`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--api_definition--reference--group-001.md#canonical-2202201313011010-3233012022102023-2231033323312223-1133031120232133-0010013132303033-0331222203303101-0022100223131121-2101300132120200) |
| `api_inventory_exclusion_list` | [api_inventory_exclusion_list](data-sources--api_definition--reference--group-001.md#canonical-3112121210312122-3313013301220203-2013031220320322-1200101203020313-3230313120032213-0103031303031132-0223032110231332-0213101223220231) |
| `api_inventory_exclusion_list.method` | [api_inventory_exclusion_list.method](data-sources--api_definition--reference--group-001.md#canonical-1331200331023031-0332120010313011-0230031102103230-2311032302103320-1313131131320111-3012221102203323-3302013130000232-1111330300313310) |
| `api_inventory_exclusion_list.path` | [api_inventory_exclusion_list.path](data-sources--api_definition--reference--group-001.md#canonical-3320320203211003-1222023010001312-2321222003311202-2313010230310021-0031222122232230-3331013112311222-0032221101200321-3130322201320131) |
| `api_inventory_inclusion_list` | [api_inventory_inclusion_list](data-sources--api_definition--reference--group-001.md#canonical-2320101200031010-3213121132000223-2302313101220100-2322122131212300-1313303010102201-0023111302121001-2323111213120010-0120302021003120) |
| `api_inventory_inclusion_list.method` | [api_inventory_inclusion_list.method](data-sources--api_definition--reference--group-001.md#canonical-2100032310232213-1333221032022223-0103202003232130-2010332112100333-1230332222312110-2222000212311113-0231023023113301-0033130133301302) |
| `api_inventory_inclusion_list.path` | [api_inventory_inclusion_list.path](data-sources--api_definition--reference--group-001.md#canonical-0302303102122330-2310303220203331-3320201232310233-2030231333332110-1020220020032030-1133130102101213-1121130200312121-3002223033031302) |
| `description` | [description](data-sources--api_definition--reference--group-001.md#canonical-1311213322111021-1321001312213330-2323110123311311-0012222130132232-2333011022031332-1233031300132102-3302120132013221-1332320210023310) |
| `id` | [ID](data-sources--api_definition--reference--group-001.md#canonical-3232103102023122-3003032310201012-2010231323210212-3033302300021101-2120032211023300-2033033330133102-0013230111202031-0002020322122202) |
| `labels` | [labels](data-sources--api_definition--reference--group-001.md#canonical-2121013333001000-3322312230301010-0033113012001213-2021100331010333-0013013111333321-2030333123331330-0122100220003202-1120301331221333) |
| `mixed_schema_origin` | [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-3311212003331003-2123303123123233-0100003201211320-3221113232132301-2132101132311002-2032300323311131-0312100233211113-1111033200212203) |
| `name` | [name](data-sources--api_definition--reference--group-001.md#canonical-1110122123122110-3312133320230320-3200120332332123-2013221133200030-0130111213002122-3233121032220303-3200021022133321-2022123012230022) |
| `namespace` | [namespace](data-sources--api_definition--reference--group-001.md#canonical-1232312233010103-2020210021003013-0002212032023013-2310113330012113-2010211320001132-3331333323132320-2121233130012310-2112112020012303) |
| `non_api_endpoints` | [non_api_endpoints](data-sources--api_definition--reference--group-001.md#canonical-1003103130211232-3030001000311032-0102021223102331-3102012312230103-1203130220103123-1002001010303223-1303010313111130-3213300232320003) |
| `non_api_endpoints.method` | [non_api_endpoints.method](data-sources--api_definition--reference--group-001.md#canonical-2321233012021100-3100323013033102-3012323121222321-1202131300321201-0210203223022121-0223112132220311-0312230200002331-0210133223013313) |
| `non_api_endpoints.path` | [non_api_endpoints.path](data-sources--api_definition--reference--group-001.md#canonical-0002233123101211-3101202311320300-0221100123301201-3101100211300102-2000332231122120-3211323302013322-1032210211003231-2030220113213331) |
| `strict_schema_origin` | [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-0111003222133310-3221331030021331-3020131323113232-1011210122211133-2121031031212132-2033320203003010-1223113202211123-1310301301111131) |
| `swagger_specs` | [swagger_specs](data-sources--api_definition--reference--group-001.md#canonical-1210301012131200-0032323033020231-1011133221130121-3110322130123223-2023312110031000-2312233112013101-2200001012031220-2121303112220020) |

<a id="canonical-0232001312320302-3223003321232222-3021132221111131-2322000200330320-3223333103311211-3020223300332311-3010130312300222-3002230133012130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_inventory_exclusion_list` properties

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123)
- api_inventory_exclusion_list

<a id="canonical-3112121210312122-3313013301220203-2013031220320322-1200101203020313-3230313120032213-0103031303031132-0223032110231332-0213101223220231"></a>

Type: `"list"`. Computed.

List of API Endpoints excluded from the API Inventory. Defaults to \`\[\]\`. Server applies default
when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0222110032032200-0303333011103221-3002012231201013-1012011121023130-2011230130331322-2110323333112101-1220120330122221-3321130303101130"></a>

### Direct properties for `api_inventory_exclusion_list`

<a id="canonical-1331200331023031-0332120010313011-0230031102103230-2311032302103320-1313131131320111-3012221102203323-3302013130000232-1111330300313310"></a>

#### `api_inventory_exclusion_list.method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-3320320203211003-1222023010001312-2321222003311202-2313010230310021-0031222122232230-3331013112311222-0032221101200321-3130322201320131"></a>

<a id="canonical-0011220222030010-3030132120103202-0213013113333032-0122223011233110-0030132020022300-3210210023001100-3130222003203131-2203003212232201"></a>

#### `api_inventory_exclusion_list.path` property

Type: `"string"`. Computed.

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-0121203120223031-2200023230310231-3112201032210320-2133301331032212-2301001131321032-1011010030320320-1303121321020311-1201003111020230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `api_inventory_inclusion_list` properties

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123)
- api_inventory_inclusion_list

<a id="canonical-2320101200031010-3213121132000223-2302313101220100-2322122131212300-1313303010102201-0023111302121001-2323111213120010-0120302021003120"></a>

Type: `"list"`. Computed.

List of API Endpoints included in the API Inventory. Typically, discovered API endpoints are added
to the API Inventory using this list. Defaults to \`\[\]\`. Server applies default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2210031221033033-2133201021131321-2233102122001023-3212011001130233-3021102100223203-3101211132013011-2322011112012012-0220310121011120"></a>

### Direct properties for `api_inventory_inclusion_list`

<a id="canonical-2100032310232213-1333221032022223-0103202003232130-2010332112100333-1230332222312110-2222000212311113-0231023023113301-0033130133301302"></a>

#### `api_inventory_inclusion_list.method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0302303102122330-2310303220203331-3320201232310233-2030231333332110-1020220020032030-1133130102101213-1121130200312121-3002223033031302"></a>

<a id="canonical-3323331013232313-1132323220323123-3322121313321333-2113110223300322-3303210112332212-0001221310211123-3302221231000102-2131023021120131"></a>

#### `api_inventory_inclusion_list.path` property

Type: `"string"`. Computed.

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-1113312301230031-0230200300100013-1130332120030100-0010201203220110-1001312011120333-1130232132203320-3220122122131201-2330322330002012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `mixed_schema_origin` properties

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123)
- mixed_schema_origin

<a id="canonical-3311212003331003-2123303123123233-0100003201211320-3221113232132301-2132101132311002-2032300323311131-0312100233211113-1111033200212203"></a>

Type: `["object", {}]`. Computed.

\[OneOf: mixed\_schema\_origin, strict\_schema\_origin\] Configuration parameter for mixed schema
origin.

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

- [mixed_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-3311212003331003-2123303123123233-0100003201211320-3221113232132301-2132101132311002-2032300323311131-0312100233211113-1111033200212203)
- [strict_schema_origin](data-sources--api_definition--reference--group-001.md#canonical-0111003222133310-3221331030021331-3020131323113232-1011210122211133-2121031031212132-2033320203003010-1223113202211123-1310301301111131)

Select alternatives according to the provider validators above.

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3023113020212033-0022101202103220-3000133332010231-2311201321011310-0101302320201220-2121200211211033-3311211213101113-3323033212322200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `non_api_endpoints` properties

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123)
- non_api_endpoints

<a id="canonical-1003103130211232-3030001000311032-0102021223102331-3102012312230103-1203130220103123-1002001010303223-1303010313111130-3213300232320003"></a>

Type: `"list"`. Computed.

API Discovery Exclusion List. List of Non-API Endpoints. Defaults to \`\[\]\`. Server applies
default when omitted.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 5000,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 5000,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "5000",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-2123111301223232-1211031313023123-1231321311302031-0323101110132210-0131210202112320-1210220321032013-2233001022321102-2022200202132300"></a>

### Direct properties for `non_api_endpoints`

<a id="canonical-2321233012021100-3100323013033102-3012323121222321-1202131300321201-0210203223022121-0223112132220311-0312230200002331-0210133223013313"></a>

#### `non_api_endpoints.method` property

Type: `"string"`. Computed.

\[Enum: ANY|GET|HEAD|POST|PUT|DELETE|CONNECT|OPTIONS|TRACE|PATCH|COPY\] Specifies the HTTP method
used to access a resource. Any HTTP Method. Possible values are \`ANY\`, \`GET\`, \`HEAD\`,
\`POST\`, \`PUT\`, \`DELETE\`, \`CONNECT\`, \`OPTIONS\`, \`TRACE\`, \`PATCH\`, \`COPY\`. Defaults to
\`ANY\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "ANY",
  "enum": [
    "ANY",
    "GET",
    "HEAD",
    "POST",
    "PUT",
    "DELETE",
    "CONNECT",
    "OPTIONS",
    "TRACE",
    "PATCH",
    "COPY"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-0002233123101211-3101202311320300-0221100123301201-3101100211300102-2000332231122120-3211323302013322-1032210211003231-2030220113213331"></a>

<a id="canonical-0123310322301121-2012120300231030-3020222322031102-1023232123113031-2120121130200212-3313031303113330-0312010210122201-1302101212111201"></a>

#### `non_api_endpoints.path` property

Type: `"string"`. Computed.

An endpoint path, as specified in OpenAPI, including parameters. The path should comply with RFC
3986 and may have parameters according to OpenAPI specification.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 1024,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 1024,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "1024",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.templated_http_path": "true"
  }
}
```

<a id="canonical-1303030311022303-2032313222313311-1221131323011003-1031102310113210-2313022033101221-0001300200113222-2020213030301212-2210322332103133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `strict_schema_origin` properties

Breadcrumbs:

- [xcsh_api_definition](../data-sources/api_definition.md#canonical-0131330321032121-2132003011332230-3320101012003002-0031333012123021-2120322202223330-2130200123310010-3012310330331120-1130223111032223)
- [Property reference](data-sources--api_definition--reference--group-001.md#canonical-3322310000110031-3303331030310132-3201123301111320-0223031203232010-1233333113000003-1023202201203132-0311103213220013-0232300212311123)
- strict_schema_origin

<a id="canonical-0111003222133310-3221331030021331-3020131323113232-1011210122211133-2121031031212132-2033320203003010-1223113202211123-1310301301111131"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for strict schema origin. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

This is an empty object or choice marker. It has no direct properties.
