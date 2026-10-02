---
page_title: "waf_exclusion_rules"
subcategory: ""
description: "An ordered list of rules."
xcsh_docs: {"aliases": ["waf exclusion rules"], "body_bytes": 10513, "body_sha256": "sha256:6a2373b34c74703b6b9045054a3e53ae45045f1a7334b023decae08e91865762", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_domain", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_path", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:metadata", "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:waf_skip_processing"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_exclusion_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "parent_id": "xcsh-docs:data-sources:waf_exclusion_policy:reference", "path": "documentation/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/index.md", "product": "distributed-cloud", "provider_name": "waf_exclusion_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0200121001220321-3030303130212110-0210100301302001-2313310012132300-3321312133133210-2012201101133232-0222103011311302-3210120102133230", "registry_path": "docs/guides/data-sources--waf_exclusion_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["waf_exclusion_rules"], "schema_version": 1, "sections": [{"aliases": ["any domain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_domain", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "any_domain"], "syntax": "attribute", "type": "object"}, {"aliases": ["any path"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:any_path", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "any_path"], "syntax": "attribute", "type": "object"}, {"aliases": ["app firewall detection control"], "anchor": "section", "description": "Define the list of Signature IDs, Violations, Attack Types and Bot Names that should be excluded from triggering on the defined match criteria.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:app_firewall_detection_control", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion_rules", "app_firewall_detection_control"], "syntax": "attribute", "type": "object"}, {"aliases": ["exact value"], "anchor": "schema-waf_exclusion_rules--exact_value", "description": "Exclusive with Exact domain name.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "exact_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["expiration timestamp"], "anchor": "schema-waf_exclusion_rules--expiration_timestamp", "description": "The expiration_timestamp is the RFC 3339 format timestamp at which the containing rule is considered to be logically expired. The rule continues to exist in the configuration but is not applied anymore.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "expiration_timestamp"], "syntax": "attribute", "type": "string"}, {"aliases": ["metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:metadata", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["waf_exclusion_rules", "metadata"], "syntax": "attribute", "type": "object"}, {"aliases": ["methods"], "anchor": "schema-waf_exclusion_rules--methods", "description": "Methods to be matched.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "methods"], "syntax": "attribute", "type": "list"}, {"aliases": ["path prefix"], "anchor": "schema-waf_exclusion_rules--path_prefix", "description": "Exclusive with Path prefix to match (e.g. The value / will match on all paths)", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "path_prefix"], "syntax": "attribute", "type": "string"}, {"aliases": ["path regex"], "anchor": "schema-waf_exclusion_rules--path_regex", "description": "Exclusive with Define the regex for the path. For example, the regex ^/.*$ will match on all paths.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "path_regex"], "syntax": "attribute", "type": "string"}, {"aliases": ["suffix value"], "anchor": "schema-waf_exclusion_rules--suffix_value", "description": "Exclusive with Suffix of domain name e.g \"xyz.com\" will match \"*.xyz.com\" and \"xyz.com\"", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "suffix_value"], "syntax": "attribute", "type": "string"}, {"aliases": ["waf skip processing"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:waf_exclusion_policy:properties:waf_exclusion_rules:waf_skip_processing", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["waf_exclusion_rules", "waf_skip_processing"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "An ordered list of rules.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["waf_exclusion_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

An ordered list of rules.

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

Upstream description:

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

<a id="schema-waf_exclusion_rules--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Computed.

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

Upstream description:

Methods to be matched.

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

Upstream description:

Exclusive with \[any\_path path\_regex\] Path prefix to match (e.g. The value / will match on all
paths)

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

<a id="schema-waf_exclusion_rules--path_regex"></a>

### path_regex property

Type: `"string"`. Computed.

Exclusive with \[any\_path path\_prefix\] Define the regex for the path. For example, the regex
^/.\*$ will match on all paths.

Upstream description:

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

Upstream description:

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

- [waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/waf_skip_processing/): complete subsection reference.

## Next pages

- [waf_exclusion_rules.any_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/any_domain/)
- [waf_exclusion_rules.any_path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/any_path/)
- [waf_exclusion_rules.app_firewall_detection_control](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/app_firewall_detection_control/)
- [waf_exclusion_rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/metadata/)
- [waf_exclusion_rules.waf_skip_processing](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/waf_exclusion_rules/waf_skip_processing/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/properties/)
- [xcsh_waf_exclusion_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_exclusion_policy/)
