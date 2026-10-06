---
page_title: "waf_exclusion_rules"
subcategory: ""
description: "An ordered list of rules."
xcsh_docs: {"aliases": ["waf exclusion rules"], "body_bytes": 8825, "body_sha256": "sha256:62e4dd1f2edb450c2f5f15225b20b56194ff5dc09926635c91f57171efb6b79a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_domain", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_path", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:metadata", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:waf_skip_processing"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "parent_id": "xcsh-docs:data-sources:waf_exclusion_policy:reference", "path": "documentation/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230", "registry_path": "docs/guides/data-sources--waf_exclusion_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion_rules"], "schema_version": 1, "sections": [{"aliases": ["waf exclusion rules any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf exclusion rules any path"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "any_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf exclusion rules app firewall detection control"], "anchor": "section", "description": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf exclusion rules exact value"], "anchor": "schema-waf_exclusion_rules--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules expiration timestamp"], "anchor": "schema-waf_exclusion_rules--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["waf exclusion rules methods"], "anchor": "schema-waf_exclusion_rules--methods", "description": "Methods to be matched.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["waf exclusion rules path prefix"], "anchor": "schema-waf_exclusion_rules--path_prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "path_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules path regex"], "anchor": "schema-waf_exclusion_rules--path_regex", "description": "Exclusive with Define the regex for the path. For example, the regex ^/.*$ will match on all paths.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "path_regex"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules suffix value"], "anchor": "schema-waf_exclusion_rules--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "suffix_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf exclusion rules waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:waf_skip_processing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "An ordered list of rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# waf_exclusion_rules

Breadcrumbs:

- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/)
- waf_exclusion_rules

<a id="section"></a>

Type: `"list"`. Computed.

WAF Exclusion Rules. An ordered list of rules.

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
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "256",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

## Direct properties

- [any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/any_domain/): complete subsection reference.

- [any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/any_path/): complete subsection reference.

- [app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/): complete subsection reference.

<a id="schema-waf_exclusion_rules--exact_value"></a>

### exact_value property

Type: `"string"`. Computed.

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

<a id="schema-waf_exclusion_rules--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Computed.

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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/metadata/): complete subsection reference.

<a id="schema-waf_exclusion_rules--methods"></a>

### methods property

Type: `["list", "string"]`. Computed.

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

<a id="schema-waf_exclusion_rules--path_prefix"></a>

### path_prefix property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths).

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
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

<a id="schema-waf_exclusion_rules--path_regex"></a>

### path_regex property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

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
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

<a id="schema-waf_exclusion_rules--suffix_value"></a>

### suffix_value property

Type: `"string"`. Computed.

Exclusive with \[any\_domain exact\_value\] Suffix of domain name e.g 'xyz.com' will match
'\*.xyz.com' and 'xyz.com'.

Additional upstream details:

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

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/waf_skip_processing/): complete subsection reference.
