---
page_title: "xcsh_dns_compliance_checks reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_compliance_checks reference."
---

# xcsh_dns_compliance_checks reference

<a id="canonical-0330131002333110-1321122220010200-2131223002102021-1310111221323211-0220230233200220-1023112122322033-2333123220200320-2331221120111033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Property reference

Breadcrumbs:

- [xcsh_dns_compliance_checks](../data-sources/dns_compliance_checks.md#canonical-0031332033312303-1303103011111022-1022102332331120-2133012122303100-3003222020332210-1130031122302313-3111210203332032-1102020121010322)
- Property reference

<a id="canonical-0021331102012301-0030220132232300-2021321022033232-3111100300222011-2003320220213032-1301113132301112-2333112300333332-1313122002102321"></a>

### Direct properties for `xcsh_dns_compliance_checks`

<a id="canonical-0331211101033202-1001010213302313-3110302331231030-1232321230202203-2333033013202011-3031013002022210-1232120231023230-3123122123032103"></a>

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

<a id="canonical-3230212112001001-0012230311311311-3230330310133102-0030113000322331-0211113121120103-2030010233201031-0112331003330213-0213331030130022"></a>

<a id="canonical-2233202031121300-1200220311103311-0020013203231101-3133103113230332-0321131230103321-3221320313202011-1313001123000313-2203313331113200"></a>

#### `description` property

Type: `"string"`. Computed.

Description of the DNSComplianceChecks.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-3210021233302303-0212113002300032-0202211131130230-3320120222111212-3232231113233100-3132131202010320-1122200013303121-3302331010212012"></a>

<a id="canonical-2000233313321321-2111310121312022-0022200301320032-2301133021103211-3020112213133202-2003132303332120-1132210221331233-1233000003330220"></a>

#### `disallowed_query_type_list` property

Type: `["list", "string"]`. Computed.

\[Enum: QUERY|IQUERY|STATUS|NOTIFY|UPDATE\] Disallowed Query Type Values. Disallowed Query Type
Values. Possible values are \`QUERY\`, \`IQUERY\`, \`STATUS\`, \`NOTIFY\`, \`UPDATE\`. Defaults to
\`QUERY\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0330113113220130-1331212301030032-3232333013302121-0100002232121132-3023113332323320-0110131101121101-1233000231233111-2133320212232312"></a>

<a id="canonical-3231220223102001-2213330303300212-3021122030323133-0130210102300233-2133130303332102-1221313012330333-2312102130323302-1013200113102330"></a>

#### `disallowed_resource_record_type_list` property

Type: `["list", "string"]`. Computed.

\[Enum:
T|A|NS|MD|MF|CNAME|SOA|MB|MG|MR|NULL|WKS|PTR|HINFO|MINFO|MX|TXT|RP|AFSDB|X25|ISDN|RT|NSAP|NSAP\_PTR|SIG|KEY|PX|GPOS|AAAA|LOC|NXT|EID|NIMLOC|SRV|ATMA|NAPTR|KX|CERT|A6|DNAME|SINK|OPT|APL|DS|SSHFP|IPSECKEY|RRSIG|NSEC|DNSKEY|DHCID|NSEC3|NSEC3PARAM|TLSA|SMIMEA|HIP|NINFO|RKEY|TALINK|CDS|CDNSKEY|OPENPGPKEY|CSYNC|SPF|UINFO|UID|GID|UNSPEC|NID|L32|L64|LP|EUI48|EUI64|TKEY|TSIG|IXFR|AXFR|MAILB|MAILA|URI|CAA|TA|DLV\]
Disallowed Resource Record Types. Disallowed Resource Record Type List. Possible values are \`T\`,
\`A\`, \`NS\`, \`MD\`, \`MF\`, \`CNAME\`, \`SOA\`, \`MB\`, \`MG\`, \`MR\`, \`NULL\`, \`WKS\`,
\`PTR\`, \`HINFO\`, \`MINFO\`, \`MX\`, \`TXT\`, \`RP\`, \`AFSDB\`, \`X25\`, \`ISDN\`, \`RT\`,
\`NSAP\`, \`NSAP\_PTR\`, \`SIG\`, \`KEY\`, \`PX\`, \`GPOS\`, \`AAAA\`, \`LOC\`, \`NXT\`, \`EID\`,
\`NIMLOC\`, \`SRV\`, \`ATMA\`, \`NAPTR\`, \`KX\`, \`CERT\`, \`A6\`, \`DNAME\`, \`SINK\`, \`OPT\`,
\`APL\`, \`DS\`, \`SSHFP\`, \`IPSECKEY\`, \`RRSIG\`, \`NSEC\`, \`DNSKEY\`, \`DHCID\`, \`NSEC3\`,
\`NSEC3PARAM\`, \`TLSA\`, \`SMIMEA\`, \`HIP\`, \`NINFO\`, \`RKEY\`, \`TALINK\`, \`CDS\`,
\`CDNSKEY\`, \`OPENPGPKEY\`, \`CSYNC\`, \`SPF\`, \`UINFO\`, \`UID\`, \`GID\`, \`UNSPEC\`, \`NID\`,
\`L32\`, \`L64\`, \`LP\`, \`EUI48\`, \`EUI64\`, \`TKEY\`, \`TSIG\`, \`IXFR\`, \`AXFR\`, \`MAILB\`,
\`MAILA\`, \`URI\`, \`CAA\`, \`TA\`, \`DLV\`. Defaults to \`T\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 8,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 8,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "8",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0221031211232313-2230013210330130-2233000211000222-0120301023030333-3200311210123200-3203001110222301-2103200012133000-1312030011221020"></a>

<a id="canonical-0312023310223120-0122303002223230-3121220310333222-1003233113231210-3000301131022131-0210301211033013-1023302132203101-1301033001223010"></a>

#### `domain_denylist` property

Type: `["list", "string"]`. Computed.

List of domains to be denied by configuration object.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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
    "ves.io.schema.rules.repeated.items.string.etld_plus_one": "true",
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.etld_plus_one": "true",
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0130312010302310-3000001030103233-0232203123023022-2003121003111133-1100223121120132-3112112112100031-1023321231030232-3220203001003222"></a>

<a id="canonical-1132001101321211-2330231220203111-1322122120333112-2223323132302210-3002111013131201-1113210102230103-1300120213312330-0200321223332322"></a>

#### `id` property

Type: `"string"`. Computed.

Unique identifier for the resource.

<a id="canonical-0200123120321132-1023200020033113-2222231100221013-2131122032300322-0130300013121012-3022233113323123-2300303002012221-2320033010312001"></a>

<a id="canonical-3100112002012102-2123303011322012-1010200121331030-0031302000210012-3132121212232201-2001000312012300-1110212233000133-2120032001310000"></a>

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

<a id="canonical-2023020133231111-1012202231022210-1321312231132122-2110102111203231-2121320023020100-3233021133033013-2322010311330320-0101123222112100"></a>

<a id="canonical-2112231332210132-2133332330201123-3031330030221132-3021011130022100-1033213331010100-1003001313023023-0331222120201312-1133100212010133"></a>

#### `name` property

Type: `"string"`. Required.

Name of the DNSComplianceChecks.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2121001310230210-0003013320013320-1121001323001121-2303330000302002-2103023022200020-3203033133020310-3133113020322031-2031323230321323"></a>

<a id="canonical-1001111210233300-2301212231322203-2321122303230322-0123113201323300-0301133331011023-3330012331232101-3213200213032232-1332221012032000"></a>

#### `namespace` property

Type: `"string"`. Required.

Namespace where the DNSComplianceChecks exists.

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

<a id="canonical-2301013201100210-1323030333312323-1301301122003101-0012110101110201-3302012333121112-3123011323011133-0123112212001221-1213220323200211"></a>

### All schema paths for `xcsh_dns_compliance_checks`

Each exact path has one authoritative reference destination. Collection element indices are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `annotations` | [annotations](data-sources--dns_compliance_checks--reference--group-001.md#canonical-0331211101033202-1001010213302313-3110302331231030-1232321230202203-2333033013202011-3031013002022210-1232120231023230-3123122123032103) |
| `description` | [description](data-sources--dns_compliance_checks--reference--group-001.md#canonical-3230212112001001-0012230311311311-3230330310133102-0030113000322331-0211113121120103-2030010233201031-0112331003330213-0213331030130022) |
| `disallowed_query_type_list` | [disallowed_query_type_list](data-sources--dns_compliance_checks--reference--group-001.md#canonical-3210021233302303-0212113002300032-0202211131130230-3320120222111212-3232231113233100-3132131202010320-1122200013303121-3302331010212012) |
| `disallowed_resource_record_type_list` | [disallowed_resource_record_type_list](data-sources--dns_compliance_checks--reference--group-001.md#canonical-0330113113220130-1331212301030032-3232333013302121-0100002232121132-3023113332323320-0110131101121101-1233000231233111-2133320212232312) |
| `domain_denylist` | [domain_denylist](data-sources--dns_compliance_checks--reference--group-001.md#canonical-0221031211232313-2230013210330130-2233000211000222-0120301023030333-3200311210123200-3203001110222301-2103200012133000-1312030011221020) |
| `id` | [ID](data-sources--dns_compliance_checks--reference--group-001.md#canonical-0130312010302310-3000001030103233-0232203123023022-2003121003111133-1100223121120132-3112112112100031-1023321231030232-3220203001003222) |
| `labels` | [labels](data-sources--dns_compliance_checks--reference--group-001.md#canonical-0200123120321132-1023200020033113-2222231100221013-2131122032300322-0130300013121012-3022233113323123-2300303002012221-2320033010312001) |
| `name` | [name](data-sources--dns_compliance_checks--reference--group-001.md#canonical-2023020133231111-1012202231022210-1321312231132122-2110102111203231-2121320023020100-3233021133033013-2322010311330320-0101123222112100) |
| `namespace` | [namespace](data-sources--dns_compliance_checks--reference--group-001.md#canonical-2121001310230210-0003013320013320-1121001323001121-2303330000302002-2103023022200020-3203033133020310-3133113020322031-2031323230321323) |
