---
page_title: "xcsh_cloud_elastic_ip reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_elastic_ip reference."
---

# xcsh_cloud_elastic_ip reference

<a id="canonical-3233011131021232-3231202322001200-0221311312333302-0113233211312111-3032320102211210-0313320223303133-3011223131121200-0001033220313023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-0223023303003220-2123332302131330-3202021121113030-2021223333123231-2302313213301131-1233113113023011-3112030123323311-3322022223022102)
- Property reference

<a id="canonical-1302201320332011-1123333303023133-3330022000300131-2331210013112103-1111030032013113-0323200222112322-3322212223200100-1321131012312200"></a>

### Direct properties for `xcsh_cloud_elastic_ip`

<a id="canonical-1302111220013233-0012220111113310-0123322012330320-0213130223231120-1121303230023300-0112232121003321-2013313131100231-1100130120330212"></a>

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

<a id="canonical-0000030330033311-2113323122131011-1000320031310111-3123132311111031-2310223232210203-2222030000312102-1222212132012300-2000203113120020"></a>

<a id="canonical-1023311332010331-0300211033210232-1202220132033213-2023010313300032-3332300010023311-1203031121132020-0023112113013232-1300003221221133"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the CloudElasticIP.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2022230203223002-2133031031130121-3200331213111210-2232001120000021-2113021011033120-0131031210322123-2120211010233301-3033320022213210"></a>

<a id="canonical-0220213111012033-2013031121211111-3113200032320122-1321312202301202-2013002102323000-3121010133023202-0001023213032301-2333132002010103"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-2302121220023130-3032133201001221-3212033132132300-1133332131211123-2102233002312121-1120333101012103-1021301031200001-1221031300323022"></a>

<a id="canonical-3202223201032323-3033030022110002-1323000022210200-0012112002202123-1133220130220311-3312220310331231-0211013123303133-0033302113001120"></a>

#### `item_count` property

Type: `"number"`. Computed.

Number of Elastic Ips / Public Ips associated with this object per Node.

<a id="canonical-2200001232132310-2003332210223133-0333033303213132-2220331221200020-2030311011212013-1310212323332321-3131232203110000-2322230220233130"></a>

<a id="canonical-3303130122220231-2331312112202233-1302031221212211-3303312321012101-3333012120023113-2003000021010332-1112201222121220-3320201312123102"></a>

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

<a id="canonical-3322232131002301-1313321223330123-0102130111000313-3111111002331301-0211101232100323-0132312300021033-3033001213013120-0133001103101020"></a>

<a id="canonical-3003100322001200-2031223100313031-1230001002311232-0223212020223332-2232301102233023-2023233122133312-2120230221122221-2200002120113331"></a>

#### `name` property

Type: `"string"`. Required.

Name of the CloudElasticIP.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1312221101220321-1032003210011100-0321231002122121-1220031210033321-2231021033032322-3210300331231133-2023203102232023-2332033131010303"></a>

<a id="canonical-3212012221313232-1101032203331132-0003221313002222-0003013313330230-1002123220213001-1102010130333233-1110202100323331-1201111000130010"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the CloudElasticIP exists.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [site_ref](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-3031310220210312-2310331000021223-1230022123230030-3313120302233100-2002210121032133-1113103213300223-3312022301231330-0313320012003333): complete subsection reference.

<a id="canonical-3010202210213022-3130310010033132-1231021331311233-1323310233201112-1321021110230100-2003013131102130-3100011031320033-1112220030303012"></a>

### All schema paths for `xcsh_cloud_elastic_ip`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-1302111220013233-0012220111113310-0123322012330320-0213130223231120-1121303230023300-0112232121003321-2013313131100231-1100130120330212) |
| `description` | [description](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-0000030330033311-2113323122131011-1000320031310111-3123132311111031-2310223232210203-2222030000312102-1222212132012300-2000203113120020) |
| `id` | [ID](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2022230203223002-2133031031130121-3200331213111210-2232001120000021-2113021011033120-0131031210322123-2120211010233301-3033320022213210) |
| `item_count` | [item_count](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2302121220023130-3032133201001221-3212033132132300-1133332131211123-2102233002312121-1120333101012103-1021301031200001-1221031300323022) |
| `labels` | [labels](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2200001232132310-2003332210223133-0333033303213132-2220331221200020-2030311011212013-1310212323332321-3131232203110000-2322230220233130) |
| `name` | [name](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-3322232131002301-1313321223330123-0102130111000313-3111111002331301-0211101232100323-0132312300021033-3033001213013120-0133001103101020) |
| `namespace` | [namespace](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-1312221101220321-1032003210011100-0321231002122121-1220031210033321-2231021033032322-3210300331231133-2023203102232023-2332033131010303) |
| `site_ref` | [site_ref](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2133201013310232-1022122013032223-0002311110133013-0231210312032322-0130031120302031-0233032211311113-0222133210330023-1033210111203033) |
| `site_ref.kind` | [site_ref.kind](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2003201222030133-1122210332002300-3213021322100100-1101120013003131-2102132002312233-1203023113103000-1203122111300312-3131321002113333) |
| `site_ref.name` | [site_ref.name](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-3011311321100001-1131120200121311-2133322332021232-2231001211321132-1233203023013211-0221123000210311-2213311003333022-0232031103330332) |
| `site_ref.namespace` | [site_ref.namespace](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2102023011322030-1221032023003330-1223213300322103-1013120233031101-2021230100030123-0310303012003021-3101110210013130-0303113030010131) |
| `site_ref.tenant` | [site_ref.tenant](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-1003130011110131-3121222230323100-1031203021322012-1023211001320301-0130210212222103-1031232131102221-2231203133311001-3232002130133211) |
| `site_ref.uid` | [site_ref.uid](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-2113231011103103-3223001222001100-1003003101223023-3202110001021111-1021000223121311-1233222332103220-0012313002002132-0201313112312132) |

<a id="canonical-3031310220210312-2310331000021223-1230022123230030-3313120302233100-2002210121032133-1113103213300223-3312022301231330-0313320012003333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `site_ref` properties

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../data-sources/cloud_elastic_ip.md#canonical-0223023303003220-2123332302131330-3202021121113030-2021223333123231-2302313213301131-1233113113023011-3112030123323311-3322022223022102)
- [Property reference](data-sources--cloud_elastic_ip--reference--group-001.md#canonical-3233011131021232-3231202322001200-0221311312333302-0113233211312111-3032320102211210-0313320223303133-3011223131121200-0001033220313023)
- site_ref

<a id="canonical-2133201013310232-1022122013032223-0002311110133013-0231210312032322-0130031120302031-0233032211311113-0222133210330023-1033210111203033"></a>

Type: `"list"`. Computed.

Site to which this cloud elastic IP object is attached.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

<a id="canonical-0033302120130200-2223030333101203-3330132203123032-0333122302113011-2022320111032322-2032002311323323-2303220100221210-0113330033122220"></a>

### Direct properties for `site_ref`

<a id="canonical-2003201222030133-1122210332002300-3213021322100100-1101120013003131-2102132002312233-1203023113103000-1203122111300312-3131321002113333"></a>

#### `site_ref.kind` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-3011311321100001-1131120200121311-2133322332021232-2231001211321132-1233203023013211-0221123000210311-2213311003333022-0232031103330332"></a>

<a id="canonical-1023102030010311-3022000002210221-0321032321320123-2111002122032122-3102002312203122-0111122112302033-2232200023113201-2333132310113010"></a>

#### `site_ref.name` property

Type: `"string"`. Computed.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2102023011322030-1221032023003330-1223213300322103-1013120233031101-2021230100030123-0310303012003021-3101110210013130-0303113030010131"></a>

<a id="canonical-1032031330320310-0220002212322333-1101320230221301-2300330203322230-3103001103221021-0010113200013300-3331122103201122-2101232211311010"></a>

#### `site_ref.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-1003130011110131-3121222230323100-1031203021322012-1023211001320301-0130210212222103-1031232131102221-2231203133311001-3232002130133211"></a>

<a id="canonical-3011010132231323-1020223003233313-1233303201002003-2302210023022110-1130222102000100-0231130332111100-3322000312013112-2202021000032031"></a>

#### `site_ref.tenant` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

<a id="canonical-2113231011103103-3223001222001100-1003003101223023-3202110001021111-1021000223121311-1233222332103220-0012313002002132-0201313112312132"></a>

<a id="canonical-3001200210331311-2023220303312021-0010323321201120-0131003322022312-2310000323222011-3031311222332230-0220023101000312-1102231131200223"></a>

#### `site_ref.uid` property

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
